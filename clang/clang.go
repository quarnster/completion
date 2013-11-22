package clang

import (
	"code.google.com/p/log4go"
	"fmt"
	cp "github.com/quarnster/completion/clang/parser"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util/expand_path"
	"github.com/quarnster/parser"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
)

func init() {
	if err := content.RegisterType("compiler_flags", reflect.TypeOf([]string{})); err != nil {
		panic(err)
	}
}

func RunClang(args ...string) ([]byte, error) {
	cmd := exec.Command("clang", args...)
	log4go.Debug("Running clang command: %v", cmd)

	if e, err := cmd.StderrPipe(); err != nil {
		return nil, err
	} else if s, err := cmd.StdoutPipe(); err != nil {
		return nil, err
	} else if err := cmd.Start(); err != nil {
		return nil, err
	} else {
		so, serr := ioutil.ReadAll(s)
		eo, eerr := ioutil.ReadAll(e)
		// We ignore the output error here as a non-zero exit
		// code doesn't necessarily mean that the output isn't
		// useful
		cmd.Wait()

		log4go.Fine("stdout: %s\n", string(so))
		log4go.Fine("stderr: %s\n", string(eo))
		if serr != nil {
			return nil, serr
		} else if eerr != nil {
			return nil, eerr
		}
		return so, nil
	}
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
	origargs, _ := a.Settings().Get("compiler_flags").([]string)
	args := make([]string, len(origargs))
	for i := range origargs {
		args[i] = expand_path.ExpandPath(origargs[i])
	}
	fn := a.Location.File.Name

	if cnt := a.Location.File.Contents; cnt != "" {
		f, err := ioutil.TempFile("", "completion_clang")
		if err != nil {
			return err
		}
		fn = f.Name()
		defer os.Remove(fn)
		fn += filepath.Ext(a.Location.File.Name)
		defer os.Remove(fn)
		if err := ioutil.WriteFile(fn, []byte(cnt), 0644); err != nil {
			return err
		}
		args = append(args, "-I"+filepath.Dir(a.Location.File.Name))
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
