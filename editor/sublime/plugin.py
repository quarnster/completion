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
import os
import os.path
import signal

def log(a):
    settings = sublime.load_settings("completion.sublime-settings")
    if settings.get("debug", False):
        print(a)

def make_transport():
    settings = sublime.load_settings("completion.sublime-settings")
    proto = settings.get("proto", "unix")
    port = settings.get("port", os.path.join(sublime.packages_path(), "User", "completion.rpc"))

    if proto == "tcp":
        tp = jsonrpc.TransportTcpIp
    elif proto == "unix":
        tp = jsonrpc.TransportUnixSocket
        if not os.path.exists(port):
            start_daemon()
    return tp(addr=port, logfunc=log, timeout=2.0)

def make_proxy():
    global proxy
    proxy = jsonrpc.ServerProxy(jsonrpc.JsonRpc10(),make_transport())

proxy = None
language_regex = re.compile("(?<=source\.)[\w+#]+")
daemon = None
last_launch = None

def pipe_reader(name, pipe, output=False):
    global daemon
    while True and daemon != None:
        try:
            line = pipe.readline()
            if len(line) == 0:
                break
            if output:
                log("%s: %s" % (name, line))
        except:
            traceback.print_exc()
            break
    daemon = None

def start_daemon(daemon_command=None):
    global daemon
    global last_launch
    if daemon_command == None:
        settings = sublime.load_settings("completion.sublime-settings")
        daemon_command = settings.get("daemon_command")

    now = time.time()
    output = False
    if last_launch != None and (now - last_launch) < 5:
        # Last tried to launch 5 seconds ago, enable debug output in the
        # pipe threads
        output = True
    last_launch = now
    daemon = subprocess.Popen(daemon_command, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    t = threading.Thread(target=pipe_reader, args=("stdout", daemon.stdout,output,))
    t.start()
    t = threading.Thread(target=pipe_reader, args=("stderr", daemon.stderr,output,))
    t.start()


def plugin_unloaded():
    global daemon
    if daemon != None:
        daemon.send_signal(signal.SIGINT)
        time.sleep(2)
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
            if daemon == None and launch_daemon:
                start_daemon(daemon_command)
            else:
                return

    e = time.time()
    print("Perform: %f ms" % ((e - s) * 1000))
    s = time.time()
    completions = []
    if debug:
        print("response:", response)
    if "Messages" in response:
        print(response["Messages"])

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

def get_language(view, caret):
    language = language_regex.search(view.scope_name(caret))
    if language == None:
        return None
    return language.group(0)

def prepare_request(view, prefix, locations, settings):
    if proxy == None:
        make_proxy()
    s = time.time()
    row, col = view.rowcol(locations[0])

    # TODO: detecting which "driver" is to be used should at some point be possible (but not required) to delegate to the server
    drivers = {
        "c++": "Clang",
        "c": "Clang",
        "java": "Java",
        "cs": "Net"
    }

    lang = get_language(view, locations[0])
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

def get_context(view, prefix, locations):
    return "%s:%s" % (view.file_name(), locations[0]-len(prefix))

def exec_goto(driver, args):
    response = driver.GetDefinition(args)

    try:
        sublime.active_window().open_file("%s:%d:%d" % (response["File"]["Name"], response["Line"], response["Column"]), sublime.ENCODED_POSITION|sublime.TRANSIENT)
    except:
        sublime.status_message("definition not found!")
        print(response)


class CompletionGotoDefinitionCommand(sublime_plugin.TextCommand):
    def __init__(self, view):
        self.view = view

    def run(self, edit):
        settings = sublime.load_settings("completion.sublime-settings")
        launch_daemon = settings.get("launch_daemon", False)
        daemon_command = settings.get("daemon_command")
        debug = settings.get("debug", False)

        locations = [a.b for a in self.view.sel()]
        prefix = self.view.substr(self.view.word(locations[0]))
        driver, args = prepare_request(self.view, prefix, locations, settings)
        if driver == None or args == None:
            return
        args["Identifier"] = prefix
        sublime.status_message("looking for definition of %s" % prefix)
        t = threading.Thread(target=exec_goto, args=(driver, args))
        t.start()

class Ev(sublime_plugin.EventListener):

    def __init__(self):
        self.request_context = None
        self.response_context = None
        self.response = None

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


    def on_query_completions(self, view, prefix, locations):
        context = get_context(view, prefix, locations)

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

        driver, args = prepare_request(view, prefix, locations, settings)
        if driver == None or args == None:
            return

        t = threading.Thread(target=do_query, args=(context, self.got_response, driver, args, launch_daemon, daemon_command, debug))
        t.start()
