package clang

import (
	"code.google.com/p/log4go"
	"fmt"
	cp "github.com/quarnster/completion/clang/parser"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/parser"
	"io/ioutil"
	"os"
	"os/exec"
)

func RunClang(args ...string) ([]byte, error) {
	cmd := exec.Command("clang", args...)
	log4go.Debug("Running clang command: %v", cmd)
	return cmd.Output()
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
		case "CFunction":
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
		case "ObjCFunction":
			f := content.Method{}
			f.Returns = append(f.Returns, content.Variable{Type: content.Type{Name: content.FullyQualifiedName{Relative: child.Children[1].Data()}}})
			args := child.Children[2:]
			for j := range args {
				p := content.Variable{}
				p.Type.Name.Relative = data(args[j].Children[1])
				p.Name.Relative = data(args[j].Children[0])

				f.Parameters = append(f.Parameters, p)
			}
			ret.Methods = append(ret.Methods, f)

		}
	}
	return
}

type Clang struct {
}

func (c *Clang) CompleteAt(a *content.CompleteAtArgs, ret *content.CompletionResult) error {
	args, _ := a.Settings().Get("compiler_flags").([]string)
	fn := a.Location.File.Name

	if a.Location.File.Contents != "" {
		if f, err := ioutil.TempFile("", "completion_clang"); err != nil {
			return err
		} else {
			defer os.Remove(f.Name())
			fn = f.Name()
		}
	}

	args = append([]string{"-fsyntax-only", "-Xclang", fmt.Sprintf("-code-completion-at=%s:%d:%d", fn, a.Location.Line, a.Location.Column)}, args...)
	args = append(args, fn)
	if out, err := RunClang(args...); len(out) == 0 {
		return err
	} else if r, err := parseresult(string(out)); err != nil {
		return err
	} else {
		*ret = r
		return nil
	}
}
