import sys
import re
import time
try:
    import completion.jsonrpc as jsonrpc
except:
    import jsonrpc
import sublime
import sublime_plugin
import subprocess
import threading

def log(a):
    settings = sublime.load_settings("completion.sublime-settings")
    if settings.get("debug", False):
        print(a)

proxy = jsonrpc.ServerProxy(jsonrpc.JsonRpc10(), jsonrpc.TransportTcpIp(addr=("127.0.0.1", 12345), logfunc=log, timeout=2.0))
language_regex = re.compile("(?<=source\.)[\w+#]+")
daemon = None

def pipe_reader(pipe):
    global daemon
    while True and daemon != None:
        try:
            line = pipe.readline()
            if len(line) == 0:
                break
        except:
            traceback.print_exc()
    daemon = None

def plugin_unloaded():
    global daemon
    if daemon != None:
        daemon.kill()
        daemon = None


def do_query(context, callback, driver, args, launch_daemon, daemon_command, debug=False):
    for i in range(2):
        # The reason for the loop here is to try and start the daemon if it wasn't running and the user settings allows this
        s = time.time()
        try:
            response = driver.CompleteAt(args)
            break
        except jsonrpc.RPCFault as e:
            print(e.error_data)
            return
        except jsonrpc.RPCTransportError as e2:
            global daemon
            if daemon == None and launch_daemon:
                daemon = subprocess.Popen(daemon_command, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
                t = threading.Thread(target=pipe_reader, args=(daemon.stdout,))
                t.start()
                t = threading.Thread(target=pipe_reader, args=(daemon.stderr,))
                t.start()
            else:
                return

    e = time.time()
    print("Perform: %f ms" % ((e - s) * 1000))
    s = time.time()
    completions = []
    if debug:
        print("response:", response)

    def relname(dict):
        return dict["Relative"] if "Relative" in dict else ""

    if "Methods" in response:
        for m in response["Methods"]:
            n = m["Name"]["Relative"] + "("
            ins = n
            res = n
            if "Parameters" in m:
                c = 1
                for p in m["Parameters"]:
                    if c > 1:
                        ins += ", "
                        res += ", "
                    tn = relname(p["Type"]["Name"])
                    vn = relname(p["Name"])
                    ins += "${%d:%s %s}" % (c, tn, vn)
                    res += "%s %s" % (tn, vn)
                    c += 1
            ins += ")"
            res += ")"
            if "Returns" in m:
                # TODO: multiple returns
                res += "\t" + relname(m["Returns"][0]["Type"]["Name"])
            completions.append((res, ins))
    if "Fields" in response:
        for f in response["Fields"]:
            tn = relname(f["Type"]["Name"])
            vn = relname(f["Name"])
            ins = "%s" % (vn)
            res = "%s\t%s" % (vn, tn)
            completions.append((res, ins))
    if "Types" in response:
        # TODO: "Types"
        print(response["Types"])

    e = time.time()
    print("Post processing: %f ms" % ((e - s) * 1000))
    if debug:
        print(completions)
    callback(context, completions)


class Ev(sublime_plugin.EventListener):

    def __init__(self):
        self.request_context = None
        self.response_context = None
        self.response = None

    def get_language(self, view, caret):
        language = language_regex.search(view.scope_name(caret))
        if language == None:
            return None
        return language.group(0)

    def prepare_request(self, view, prefix, locations, settings):
        s = time.time()
        row, col = view.rowcol(locations[0])

        # TODO: detecting which "driver" is to be used should at some point be possible (but not required) to delegate to the server
        drivers = {
            "c++": "Clang",
            "c": "Clang",
            "java": "Java",
            "cs": "Net"
        }

        lang = self.get_language(view, locations[0])
        if not lang in drivers:
            return (None, None)
        else:
            driver = getattr(proxy, drivers[lang])

        # TODO: Make the request async
        args = {
            "Location": {
                "File": {
                    "Name": view.file_name(),
                },
                "Column": col + 1,
                "Line": row + 1
            },
            "SessionOverrides": {
                # TODO: what would be a good way to handle this? Query the "driver" for which options are configurable?
                # TODO: Sessions should be used when possible to avoid sending the same configuration all the time.
                "compiler_flags": view.settings().get("sublimeclang_options", []),
                "net_paths":view.settings().get("net_paths", []),
                "net_assemblies":view.settings().get("net_assemblies", []),
            }
        }
        if view.is_dirty():
            args["Location"]["File"]["Contents"] = view.substr(sublime.Region(0, view.size()))

        e = time.time()
        print("Prepare: %f ms" % ((e - s) * 1000))

        return (driver, args)

    def got_response(self, context, completions):
        if self.request_context != context:
            print("Context out of date.. Discarding.")
            return
        self.response_context = context
        self.response = completions
        if len(completions) != 0:
            sublime.set_timeout(self.hack)

    def hack(self):
        sublime.active_window().run_command("hide_auto_complete")
        def hack2():
            sublime.active_window().run_command("auto_complete")
        sublime.set_timeout(hack2, 1)

    def get_context(self, view, prefix, locations):
        return "%s:%s" % (view.file_name(), locations[0]-len(prefix))

    def on_query_completions(self, view, prefix, locations):
        context = self.get_context(view, prefix, locations)

        if self.response_context == context:
            # Have a finished response from before
            return self.response
        elif self.request_context == context:
            # Currently being processed already...
            return

        # No active query or the context of the current query is not what we want.
        # Start a new request.
        self.request_context = context

        settings = sublime.load_settings("completion.sublime-settings")
        launch_daemon = settings.get("launch_daemon", False)
        daemon_command = settings.get("daemon_command")
        debug = settings.get("debug", False)

        driver, args = self.prepare_request(view, prefix, locations, settings)
        if driver == None or args == None:
            return

        t = threading.Thread(target=do_query, args=(context, self.got_response, driver, args, launch_daemon, daemon_command, debug))
        t.start()
