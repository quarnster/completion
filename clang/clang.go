package clang

import (
	"fmt"
	"github.com/quarnster/completion/clang/parser"
	"github.com/quarnster/completion/content"
	"os/exec"
)

func RunClang(args ...string) ([]byte, error) {
	cmd := exec.Command("clang", args...)
	fmt.Println("cmd:", cmd.Args)
	return cmd.CombinedOutput()
}

func parseresult(in string) (ret content.CompletionResult, err error) {
	var p parser.PARSER
	p.Parse(in)
	n := p.RootNode()
	for i := range n.Children {
		child := n.Children[i]
		switch child.Name {
		case "Variable":
			v := content.Field{}
			v.Type.Name.Relative = child.Children[0].Data()
			v.Name.Relative = child.Children[1].Data()
			ret.Fields = append(ret.Fields, v)
		case "Function":
			f := content.Method{}
			f.Returns = append(f.Returns, content.Variable{Type: content.Type{Name: content.FullyQualifiedName{Relative: child.Children[0].Data()}}})
			f.Name.Relative = child.Children[1].Data()
			args := child.Children[2:]
			for j := range args {
				p := content.Variable{}
				p.Type.Name.Relative = args[j].Data()
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
