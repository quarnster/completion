import sys
import jsonrpc
import sublime
import sublime_plugin
def log(a):
    print(a)

proxy = jsonrpc.ServerProxy( jsonrpc.JsonRpc10(), jsonrpc.TransportTcpIp(addr=("127.0.0.1",12345), logfunc=log))

class Ev(sublime_plugin.EventListener):
    def on_query_completions(self, view, prefix, locations):
        row, col = view.rowcol(locations[0])

        # TODO: detect which "driver" should be used
        # TODO: Make the request async
        args = {
            "Location": {
                "File": {
                    "Name": view.file_name(),
                },
                "Column": col+1,
                "Line": row+1
            },
            "SessionOverrides": {
                # TODO: what would be a good way to handle this? Query the "driver" for which options are configurable?
                # TODO: Sessions should be used when possible to avoid sending the same configuration all the time.
                "compiler_flags": view.settings().get("sublimeclang_options", []),
            }
        }
        if view.is_dirty():
            args["Location"]["File"]["Contents"] = view.substr(sublime.Region(0, view.size()))

        res = proxy.Clang.CompleteAt(args)

        completions = []
        if "Methods" in res:
            for m in res["Methods"]:
                n = m["Name"]["Relative"] + "("
                ins = n
                res = n
                if "Parameters" in m:
                    c = 1
                    for p in m["Parameters"]:
                        if c > 1:
                            ins += ", "
                            res += ", "
                        tn = p["Type"]["Name"]["Relative"]
                        vn = p["Name"]["Relative"] if "Relative" in p["Name"] else ""
                        ins += "${%d:%s %s}" % (c, tn, vn)
                        res += "%s %s" % (tn, vn)
                        c += 1
                ins += ")"
                res += ")\t" + m["Returns"][0]["Type"]["Name"]["Relative"]
                completions.append((res, ins))
        if "Fields" in res:
            for f in res["Fields"]:
                tn = f["Type"]["Name"]["Relative"]
                vn = f["Name"]["Relative"] if "Relative" in f["Name"] else ""
                ins = "%s" % (vn)
                res = "%s\t%s" % (vn, tn)
                completions.append((res, ins))
        if "Types" in res:
            # TODO: "Types"
            print(res["Types"])

        return completions

