package net

import (
	"code.google.com/p/log4go"
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/net/csharp"
	"github.com/quarnster/completion/util/scopes"
	"github.com/quarnster/parser"
	"reflect"
	"regexp"
	"strings"
)

func init() {
	if err := content.RegisterType("net_paths", reflect.TypeOf([]string{})); err != nil {
		panic(err)
	} else if err := content.RegisterType("net_assemblies", reflect.TypeOf([]string{})); err != nil {
		panic(err)
	}
}

type Net struct {
}

func typeresolve(td *TypeDef, node *parser.Node) (*content.Type, error) {
	switch n := node.Name; n {
	case "MethodCall":
		methods, err := td.Methods()
		if err != nil {
			return nil, err
		}
		for _, method := range methods {
			if method.Name.Relative == node.Children[0].Data() {
				return &method.Returns[0].Type, nil
			}
		}
	case "Identifier":
		fields, err := td.Fields()
		if err != nil {
			return nil, err
		}
		for _, field := range fields {
			if field.Name.Relative == node.Data() {
				return &field.Type, nil
			}
		}

		// Is it a Property then perhaps?
		name := "get_" + node.Data()
		methods, err := td.Methods()
		if err != nil {
			return nil, err
		}
		for _, method := range methods {
			if method.Name.Relative == name {
				return &method.Returns[0].Type, nil
			}
		}
		// TODO: could also be an inner class
	}

	// Found nothing.
	// TODO: Try parents
	return nil, fmt.Errorf("No such type found: %s, %s", td.Name(), node)
}

var tm = map[string]string{
	"string": "System.String",
	"int":    "System.Int32",
}

func findtype(cache *Cache, using *parser.Node, name string) *TypeDef {
	if n, ok := tm[name]; ok {
		name = n
	}
	n := content.FullyQualifiedName{Absolute: fmt.Sprintf("net://type/%s", name)}
	if td, _ := cache.FindType(n); td != nil {
		return td
	}

	for _, child := range using.Children {
		n := content.FullyQualifiedName{Absolute: fmt.Sprintf("net://type/%s.%s", child.Children[0].Data(), name)}
		if td, _ := cache.FindType(n); td != nil {
			return td
		}
	}

	return nil
}

func (c *Net) cache(args *content.Args) (*Cache, error) {
	var cache *Cache
	session := args.SessionOrCreate(".internal.net.default")
	if a, ok := session.Get("net_cache").(*Cache); ok {
		cache = a
	}
	if cache == nil {
		log4go.Debug("cache is nil..")
		paths := DefaultPaths()
		c := Cache{paths: paths}
		std := []string{"mscorlib.dll", "System.dll"}
		for _, lib := range std {
			asm, err := c.Load(lib)
			if err != nil {
				return nil, err
			}
			log4go.Debug("Found %s (%s)", lib, asm.Name())
		}
		cache = &c
		session.Set("net_cache", cache)
	}
	s := args.Settings()
	if v, ok := s.Get("net_paths").([]string); ok {
		for _, p := range v {
			cache.AddPath(p)
		}
	}
	if v, ok := s.Get("net_assemblies").([]string); ok {
		for _, lib := range v {
			asm, err := cache.Load(lib)
			if err != nil {
				return nil, err
			}
			log4go.Debug("Found %s (%s)", lib, asm.Name())
		}
	}

	return cache, nil
}

func (c *Net) variable(n *parser.Node) string {
	switch n.Name {
	case "Access":
	default:
		if len(n.Children) == 0 {
			return n.Data()
		}
		for _, child := range n.Children {
			if v := c.variable(child); v != "" {
				return v
			}
		}
	}
	return ""
}

func (c *Net) classes(path string, cs []string, n *parser.Node) []string {
	switch n.Name {
	case "Namespace", "Class":
		var p2 = ""
		for _, child := range n.Children {
			switch child.Name {
			case "Identifier", "SpacedIdentifier":
				p2 = path + child.Data()
				cs = append(cs, p2)
			case "ClassScope":
				cs = c.classes(p2+"$", cs, child)
			case "Namespace", "Class":
				cs = c.classes(p2+".", cs, child)
			}
		}
		return cs
	case "CSHARP", "ClassScope":
		for _, child := range n.Children {
			cs = c.classes(path, cs, child)
		}
	}
	return cs
}

var tdnil = fmt.Errorf("Typedef is nil")

func (c *Net) CompleteAt(args *content.CompleteAtArgs, cmp *content.CompletionResult) error {
	cache, err := c.cache(&args.Args)
	if err != nil {
		return err
	}
	if err := args.Location.File.Load(); err != nil {
		return err
	}
	contents := args.Location.File.Contents
	var up csharp.CSHARP
	up.SetData(contents)
	if !up.UsingDirectives() {
		return up.Error()
	}
	using := up.RootNode()

	var p csharp.CSHARP
	off := args.Location.Offset()
	line := args.Location.File.Line(off)
	line = line[:args.Location.Column-1]
	p.SetData(line)
	if !p.Complete() {
		return fmt.Errorf("%s, %s, %s", line, p.Error(), p.RootNode())
	}
	r := p.RootNode()
	r.Simplify()
	var td *TypeDef
	var complete func(node *parser.Node) error
	var ns string

	complete = func(node *parser.Node) error {
		if node.Name != "CompleteOp" {
			return fmt.Errorf("Don't know how to complete %s", node)
		}
		n := len(node.Children)
		if n < 2 {
			return fmt.Errorf("Not enough children in Op node: %d < 2: %s", n, node)
		}
		if op := node.Children[1].Data(); op != "." {
			if n < 3 {
				return fmt.Errorf("Not enough children in Op node: %d < 3: %s", n, node)
			}
			td = nil
			return complete(node.Children[2])
		}
		if td == nil {
			tn := ns + node.Children[0].Data()
			if td = findtype(cache, using, tn); td == nil {
				ns = tn + "."
			}
		} else if t2, err := typeresolve(td, node.Children[0]); err != nil {
			return err
		} else if td, err = cache.FindType(t2.Name); err != nil {
			return err
		} else if td == nil {
			return fmt.Errorf("Couldn't find type: %s", node.Children[0])
		}
		if td == nil {
			// Probably a variable completion then?
			v := c.variable(node.Children[0])
			loc := content.File{Contents: scopes.Substr(args.Location.File.Contents, scopes.Visibility(args.Location))}
			re, err := regexp.Compile(fmt.Sprintf(`[=\(,;}\s](\w+(\s*\[\])*\s+)*%s[;\s=,)\[]`, v))
			if err != nil {
				return err
			}
			idx := re.FindAllStringIndex(loc.Contents, -1)
			for _, i := range idx {
				// TODO: It's better at getting the right variable, but still not 100% correct
				line := loc.Contents[i[0]+1 : i[1]-1]
				var p csharp.CSHARP
				p.SetData(line)
				if !p.CompleteVariable() {
					continue
				}
				t := c.variable(p.RootNode())
				if t == "" {
					continue
				}
				if td = findtype(cache, using, t); td != nil {
					break
				}
				// Internal class perhaps?
				var p2 csharp.CSHARP
				if !p2.Parse(loc.Contents) {
					continue
				}
				for _, t2 := range c.classes("", nil, p2.RootNode()) {
					if !strings.HasSuffix(t2, t) {
						continue
					}
					if td = findtype(cache, using, t2); td != nil {
						break
					}
				}
				if td != nil {
					break
				}
			}
		}
		if n > 2 {
			return complete(node.Children[2])
		} else if td == nil {
			return tdnil
		}
		return c.complete(cache, &content.Type{Name: td.Name()}, cmp)
	}

	if err := complete(r.Children[0]); err == tdnil {
		return err
	} else {
		return err
	}
}

func (c *Net) complete(cache *Cache, t *content.Type, cmp *content.CompletionResult) error {
	ct, err := cache.Complete(t)
	if err != nil {
		return err
	}
	cmp.Fields = ct.Fields
	cmp.Types = ct.Types
	cmp.Methods = ct.Methods
	return nil
}

func (c *Net) Complete(args *content.CompleteArgs, cmp *content.CompletionResult) error {
	cache, err := c.cache(&args.Args)
	if err != nil {
		return err
	}
	log4go.Debug("Trying to complete: %s", args.Location)
	return c.complete(cache, &content.Type{Name: args.Location}, cmp)
}
