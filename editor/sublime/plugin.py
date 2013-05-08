import sys
import jsonrpc
import sublime
import sublime_plugin

proxy = jsonrpc.ServerProxy( jsonrpc.JsonRpc10(), jsonrpc.TransportTcpIp(addr=("127.0.0.1",12345)) )

class Ev(sublime_plugin.EventListener):
    def on_query_completions(self, view, prefix, locations):
        row, col = view.rowcol(locations[0])

        # TODO: detect which "driver" should be used
        # TODO: save the file to a temporary location if it's unsaved
        # TODO: Make the request async
        res = proxy.Clang.CompleteAt({"Location": {"File": {"Name": view.file_name()}, "Column": col+1, "Line": row+1}})
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
        # TODO: "Fields"
        # TODO: "Types"
        return completions

