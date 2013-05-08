package expand_path

import (
	"bytes"
	"code.google.com/p/log4go"
	"github.com/quarnster/parser"
	"os"
	"os/user"
	"path/filepath"
)

func eval(node *parser.Node) string {
	switch node.Name {
	case "Env":
		name := eval(node.Children[0])
		if ret := os.Getenv(name); ret == "" {
			return name + "_NOT_SET"
		} else {
			return ret
		}
	case "Folder":
		return filepath.Dir(eval(node.Children[0]))
	case "String", "NonOperation", "RecNonOp":
		return node.Data()
	case "Home":
		if u, err := user.Current(); err != nil {
			log4go.Warn("Couldn't lookup current user: %s", err)
		} else {
			return u.HomeDir
		}
	case "EXPAND_PATH", "RecOp":
		buf := bytes.NewBuffer(nil)
		for _, c := range node.Children {
			buf.WriteString(eval(c))
		}
		return buf.String()
	}
	return ""
}

func ExpandPath(path string) string {
	var p EXPAND_PATH
	if !p.Parse(path) {
		log4go.Error("Failed to parse expand_path expression. Input data: %s, Error: %s, Parsed tree: %s", path, p.Error(), p.RootNode())
		return path
	}
	return eval(p.RootNode())
}
