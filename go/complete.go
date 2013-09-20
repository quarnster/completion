// +build go_types_changed_again

package goo

import (
	"code.google.com/p/go.tools/go/types"
	"code.google.com/p/log4go"
	"github.com/quarnster/completion/content"
	//	"go/ast"
	//	"go/parser"
	//	"go/token"
	//	"io/ioutil"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var indent string

type Go struct {
	imports map[string]*types.Package
}

func (g *Go) method(m reflect.Method) (ret content.Method) {
	ret.Name.Relative = m.Name
	if pkg := m.PkgPath; pkg != "" {
		ret.Name.Absolute = m.PkgPath + "/" + m.Name
	}
	t := m.Func.Type()
	for i := 1; i < t.NumIn(); i++ {
		in := t.In(i)
		ret.Parameters = append(ret.Parameters, content.Variable{Type: g.ty(in)})
	}
	for i := 0; i < t.NumOut(); i++ {
		out := t.Out(i)
		ret.Returns = append(ret.Returns, content.Variable{Type: g.ty(out)})
	}
	return
}

func (g *Go) ty(t reflect.Type) (ret content.Type) {

	ret.Name.Relative = t.Name()
	if pkg := t.PkgPath(); pkg != "" {
		ret.Name.Absolute = pkg + "/" + t.Name()
	}
	if t.Kind() == reflect.Ptr {
		ret.Flags |= content.FLAG_TYPE_POINTER
		ret.Specialization = append(ret.Specialization, g.ty(t.Elem()))
	}
	return
}

func (g *Go) field(sf reflect.StructField) (ret content.Field) {
	ret.Name.Relative = sf.Name
	ret.Name.Absolute = sf.PkgPath + "/" + sf.Name
	ret.Type = g.ty(sf.Type)
	return
}

func (g *Go) complete(t reflect.Type, cmp *content.CompletionResult) error {
	if k := t.Kind(); k != reflect.Ptr {
		return fmt.Errorf("Expected a pointer to a type not %v", k)
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		cmp.Methods = append(cmp.Methods, g.method(m))
	}

	t = t.Elem()
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		cmp.Methods = append(cmp.Methods, g.method(m))
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		cmp.Fields = append(cmp.Fields, g.field(f))
	}

	return nil
}

var dotre = regexp.MustCompile(`·\d+`)

// TODO(): This does not return the right think for *, [], map[] types
func (g *Go) pkg_type(t types.Type) (ret content.Type) {
	n := t.String()
	if i := strings.LastIndex(n, "."); i > 0 {
		ret.Name.Relative = n[i+1:]
		ret.Name.Absolute = n
	} else {
		ret.Name.Relative = n
		ret.Name.Absolute = n
	}
	return
}

func (g *Go) pkg_var(v *types.Var) (ret content.Variable) {
	ret.Name.Relative = dotre.ReplaceAllString(v.Name(), "")
	ret.Type = g.pkg_type(v.Type())
	return
}

func (g *Go) complete_pkg(pkg string, cmp *content.CompletionResult) error {
	if g.imports == nil {
		g.imports = make(map[string]*types.Package)
	}

	if p, err := types.GcImport(g.imports, pkg); err != nil {
		return err
	} else {
		nn := p.Scope()
		for i := 0; i < nn.NumEntries(); i++ {
			t := nn.At(i)

			var flags content.Flags

			if n := t.Name(); n[0] != strings.ToUpper(n)[0] {
				flags = content.FLAG_ACC_PROTECTED
			} else {
				flags = content.FLAG_ACC_PUBLIC
			}
			switch t.(type) {
			case *types.Func:
				var m content.Method
				m.Flags |= flags
				m.Name.Relative = t.Name()
				sig := t.Type().Underlying().(*types.Signature)
				if sig.Recv() != nil {
					continue
				}
				par := sig.Params()
				for j := 0; j < par.Len(); j++ {
					m.Parameters = append(m.Parameters, g.pkg_var(par.At(j)))
				}
				ret := sig.Results()
				for j := 0; j < ret.Len(); j++ {
					m.Returns = append(m.Returns, g.pkg_var(ret.At(j)))
				}
				cmp.Methods = append(cmp.Methods, m)
			case *types.TypeName:
				var t2 content.Type
				t2.Flags |= flags
				t2.Name.Relative = t.Name()
				switch t.Type().Underlying().(type) {
				case *types.Interface:
					t2.Flags |= content.FLAG_TYPE_INTERFACE
				case *types.Struct:
					t2.Flags |= content.FLAG_TYPE_STRUCT
				}
				cmp.Types = append(cmp.Types, t2)
			case *types.Const, *types.Var:
				var f content.Field
				f.Name.Relative = t.Name()
				f.Type = g.pkg_type(t.Type())
				cmp.Fields = append(cmp.Fields, f)
			default:
				log4go.Warn("Unimplemented type in package completion: at: %+v, %v, %v", t, reflect.TypeOf(t), reflect.TypeOf(t.Type().Underlying()))
			}
		}
	}
	return nil
}

func (g *Go) Complete(args *content.CompleteArgs, cmp *content.CompletionResult) error {
	return nil
}

func (g *Go) CompleteAt(args *content.CompleteAtArgs, cmp *content.CompletionResult) error {
	return nil
}
