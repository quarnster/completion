package clang

import (
	"github.com/limetext/log4go"
	"fmt"
	cp "github.com/quarnster/completion/clang/parser"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util/expand_path"
	"github.com/quarnster/parser"
	"io/ioutil"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
)

func init() {
	if err := content.RegisterType("compiler_flags", reflect.TypeOf([]string{})); err != nil {
		panic(err)
	}
}

func RunClang(stdin string, args ...string) ([]byte, []byte, error) {
	cmd := exec.Command("clang", args...)
	log4go.Debug("Running clang command: %v", cmd)

	if in, err := cmd.StdinPipe(); err != nil {
		return nil, nil, err
	} else if e, err := cmd.StderrPipe(); err != nil {
		return nil, nil, err
	} else if s, err := cmd.StdoutPipe(); err != nil {
		return nil, nil, err
	} else if err := cmd.Start(); err != nil {
		return nil, nil, err
	} else {
		in.Write([]byte(stdin))
		in.Close()
		so, serr := ioutil.ReadAll(s)
		eo, eerr := ioutil.ReadAll(e)
		// We ignore the output error here as a non-zero exit
		// code doesn't necessarily mean that the output isn't
		// useful
		cmd.Wait()

		log4go.Fine("stdout: %s\n", string(so))
		log4go.Fine("stderr: %s\n", string(eo))
		if serr != nil {
			return so, eo, serr
		} else if eerr != nil {
			return so, eo, eerr
		}
		return so, eo, nil
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

func (c *Clang) prepare(a *content.CompleteAtArgs) (fn string, args []string, err error) {
	origargs, _ := a.Settings().Get("compiler_flags").([]string)
	args = make([]string, len(origargs))
	for i := range origargs {
		args[i] = expand_path.ExpandPath(origargs[i])
	}
	fn = a.Location.File.Name
	if a.Location.File.Contents != "" {
		// File is unsaved, so use stdin as the filename
		fn = "-"
	}

	return fn, args, nil
}

func (c *Clang) GetDefinition(a *content.GetDefinitionArgs, ret *content.SourceLocation) error {
	fn, args, err := c.prepare(&a.CompleteAtArgs)
	if err != nil {
		return err
	}
	args = append([]string{"-fsyntax-only", "-Xclang", "-ast-dump", "-Xclang", "-ast-dump-filter", "-Xclang", a.Identifier}, args...)
	args = append(args, fn)
	out, oute, err := RunClang(a.Location.File.Contents, args...)
	if len(out) == 0 {
		if err != nil {
			return err
		} else {
			return fmt.Errorf("%s", oute)
		}
	}
	re, err := regexp.Compile(`\w+Decl[^<]+<(..[^:,]+):?(\d+)?:?(\d+)?.*?\s` + a.Identifier + `\s`)
	if err != nil {
		return err
	}
	res := re.FindAllStringSubmatch(string(out), -1)
	if len(res) == 0 {
		return fmt.Errorf("Not found")
	}
	ret.File.Name = res[0][1]
	i, _ := strconv.ParseInt(res[0][2], 10, 32)
	ret.Line = uint(i)
	i, _ = strconv.ParseInt(res[0][3], 10, 32)
	ret.Column = uint(i)
	return nil
}

func (c *Clang) CompleteAt(a *content.CompleteAtArgs, ret *content.CompletionResult) error {
	fn, args, err := c.prepare(a)
	if err != nil {
		return err
	}

	args = append([]string{"-fsyntax-only", "-Xclang", fmt.Sprintf("-code-completion-at=%s:%d:%d", fn, a.Location.Line, a.Location.Column)}, args...)
	args = append(args, fn)
	if out, oute, err := RunClang(a.Location.File.Contents, args...); len(out) == 0 {
		if err != nil {
			return err
		} else {
			return fmt.Errorf("%s", oute)
		}
	} else if r, err := parseresult(string(out)); err != nil {
		return err
	} else {
		*ret = r
		ret.Messages = string(oute)
		return nil
	}
}
