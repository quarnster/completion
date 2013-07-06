package goo

import (
	"github.com/quarnster/completion/content"
	//	"go/ast"
	//	"go/parser"
	//	"go/token"
	//	"io/ioutil"
	"fmt"
	"reflect"
)

var indent string

type Go struct {
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

func (g *Go) Complete(args *content.CompleteArgs, cmp *content.CompletionResult) error {
	return nil
}

func (g *Go) CompleteAt(args *content.CompleteAtArgs, cmp *content.CompletionResult) error {
	return nil
}
