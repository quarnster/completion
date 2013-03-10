package clang

import (
	"fmt"
	cp "github.com/quarnster/completion/clang/parser"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/parser"
	"os/exec"
)

func RunClang(args ...string) ([]byte, error) {
	cmd := exec.Command("clang", args...)
	fmt.Println("cmd:", cmd.Args)
	return cmd.CombinedOutput()
}

func data(n *parser.Node) string {
	switch n.Name {
	default:
		return n.Data()
	case "InheritedName":
		return data(n.Children[1])
	case "OptionalArgument":
		return data(n.Children[0])
	}
}

func parseresult(in string) (ret content.CompletionResult, err error) {
	var p cp.PARSER
	p.Parse(in)
	n := p.RootNode()
	for i := range n.Children {
		child := n.Children[i]
		switch child.Name {
		case "Variable":
			v := content.Field{}
			v.Type.Name.Relative = child.Children[0].Data()
			v.Name.Relative = data(child.Children[1])
			ret.Fields = append(ret.Fields, v)
		case "Function":
			f := content.Method{}
			f.Returns = append(f.Returns, content.Variable{Type: content.Type{Name: content.FullyQualifiedName{Relative: child.Children[0].Data()}}})
			f.Name.Relative = data(child.Children[1])
			args := child.Children[2:]
			for j := range args {
				if args[j].Name == "ConstQualifier" {
					// TODO
					break
				}
				p := content.Variable{}
				p.Type.Name.Relative = data(args[j])
				f.Parameters = append(f.Parameters, p)
			}
			ret.Methods = append(ret.Methods, f)
		}
	}
	return
}

func CompleteAt(loc content.SourceLocation) (ret content.CompletionResult, err error) {
	if out, err := RunClang("-fsyntax-only", "-Xclang", fmt.Sprintf("-code-completion-at=%s:%d:%d", loc.File.Name, loc.Line, loc.Column), loc.File.Name); err != nil {
		return ret, err
	} else {
		return parseresult(string(out))
	}
}
