package net

import (
	"code.google.com/p/log4go"
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/net/csharp"
	"github.com/quarnster/parser"
	"regexp"
)

type Net struct {
}

func typeresolve(td *TypeDef, node *parser.Node) (*content.Type, error) {
	switch n := node.Name; n {
	case "MethodCall":
		if methods, err := td.Methods(); err != nil {
			return nil, err
		} else {
			for _, method := range methods {
				if method.Name.Relative == node.Children[0].Data() {
					return &method.Returns[0].Type, nil
				}
			}
		}
	case "Identifier":
		if fields, err := td.Fields(); err != nil {
			return nil, err
		} else {
			for _, field := range fields {
				if field.Name.Relative == node.Data() {
					return &field.Type, nil
				}
			}
		}

		// Is it a Property then perhaps?
		name := "get_" + node.Data()
		if methods, err := td.Methods(); err != nil {
			return nil, err
		} else {
			for _, method := range methods {
				if method.Name.Relative == name {
					return &method.Returns[0].Type, nil
				}
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
		paths := DefaultPaths()
		s := args.Settings()
		if v, ok := s.Get("net_paths").([]string); ok {
			paths = append(paths, v...)
		}
		c := Cache{paths: paths}
		std := []string{"mscorlib.dll", "System.dll"}
		for _, lib := range std {
			if asm, err := c.Load(lib); err != nil {
				return nil, err
			} else {
				log4go.Debug("Found %s (%s)", lib, asm.Name())
			}
		}
		cache = &c
		session.Set("net_cache", cache)
	}
	return cache, nil
}

func (c *Net) variable(n *parser.Node) string {
	switch n.Name {
	case "ReturnType":
		return n.Data()
	default:
		if len(n.Children) > 0 {
			return c.variable(n.Children[0])
		}
	}
	return ""
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
	line := args.Location.File.Line(args.Location.Offset())
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
		switch n := node.Name; n {
		case "Op":
			n := len(node.Children)
			if n < 2 {
				return fmt.Errorf("Not enough children in Op node: %d < 2: %s", n, node)
			}
			switch op := node.Children[1].Data(); op {
			case ".":
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
					v := node.Children[0].Data()
					if re, err := regexp.Compile(fmt.Sprintf(`[\s,]%s[,;=\)\s]`, v)); err != nil {
						return err
					} else {
						idx := re.FindAllStringIndex(contents, -1)
						for _, i := range idx {
							line := args.Location.File.Line(content.Offset(i[0]))
							var p csharp.CSHARP
							p.SetData(line)
							if !p.Complete() {
								return fmt.Errorf("%s, %s, %s", line, p.Error(), p.RootNode())
							} else {
								if t := c.variable(p.RootNode()); t != "" {
									// TODO: Need to figure out which variables are actually visible
									//       rather than just returning the first variable we stumble upon
									td = findtype(cache, using, t)
									if td != nil {
										break
									}
								}
							}
						}
					}
				}
				if n > 2 {
					return complete(node.Children[2])
				} else if td == nil {
					return tdnil
				}
				return c.complete(cache, &content.Type{Name: td.Name()}, cmp)
			default:
				if n < 3 {
					return fmt.Errorf("Not enough children in Op node: %d < 3: %s", n, node)
				}
				td = nil
				return complete(node.Children[2])
			}
		default:
			return fmt.Errorf("Don't know how to complete %s", node)
		}
		return nil
	}

	if err := complete(r.Children[0]); err == tdnil {
		return err
	} else {
		return err
	}

}

func (c *Net) complete(cache *Cache, t *content.Type, cmp *content.CompletionResult) error {
	if ct, err := cache.Complete(t); err != nil {
		return err
	} else {
		cmp.Fields = ct.Fields
		cmp.Types = ct.Types
		cmp.Methods = ct.Methods
	}
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
