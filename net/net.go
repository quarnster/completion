package net

import (
	"code.google.com/p/log4go"
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/net/csharp"
	"github.com/quarnster/parser"
	"io/ioutil"
	"strings"
)

type Net struct {
}

func resolve(node *parser.Node) string {
	switch n := node.Name; n {
	case "Identifier":
		return node.Data()
	default:
		if len(node.Children) > 0 {
			return resolve(node.Children[0])
		} else {
			return ""
		}
	}
}

func base(node *parser.Node) *parser.Node {
	switch n := node.Name; n {
	case "DotIdentifier":
		return node.Children[0]
	}
	return node
}

func findtype(cache *Cache, using *parser.Node, base *parser.Node) *TypeDef {
	bn := base.Data()
	for _, child := range using.Children {
		n := content.FullyQualifiedName{Absolute: fmt.Sprintf("net://type/%s.%s", child.Children[0].Data(), bn)}
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

func (c *Net) CompleteAt(args *content.CompleteAtArgs, cmp *content.CompletionResult) error {
	cache, err := c.cache(&args.Args)
	if err != nil {
		return err
	}
	contents := args.Location.File.Contents
	if contents == "" {
		if d, err := ioutil.ReadFile(args.Location.File.Name); err != nil {
			return err
		} else {
			contents = string(d)
		}
	}

	var up csharp.CSHARP
	up.SetData(contents)
	if !up.UsingDirectives() {
		return up.Error()
	}
	using := up.RootNode()

	lines := strings.Split(contents, "\n")
	line := lines[args.Location.Line-1]

	var p csharp.CSHARP
	line = line[:args.Location.Column-1]
	p.SetData(line)
	if !p.Complete() {
		return fmt.Errorf("%s, %s, %s", line, p.Error(), p.RootNode())
	}
	r := p.RootNode()
	r.Simplify()

	var complete func(node *parser.Node) error
	complete = func(node *parser.Node) error {
		//		log4go.Debug("Complete: %s", node)
		switch n := node.Name; n {
		case "Op":
			n := len(node.Children)
			if n < 2 {
				return fmt.Errorf("Not enough children in Op node: %d < 2: %s", n, node)
			}
			switch op := node.Children[1].Data(); op {
			case ".":
				return complete(node.Children[0])
			default:
				if n < 3 {
					return fmt.Errorf("Not enough children in Op node: %d < 3: %s", n, node)
				} else {

				}
				return complete(node.Children[2])
			}
		case "MethodCall":
			bn := base(node.Children[0])
			if td := findtype(cache, using, bn); td == nil {
				return fmt.Errorf("Couldn't find base type with name %s", bn)
			} else if methods, err := td.Methods(); err != nil {
				return err
			} else {
				for _, method := range methods {
					if method.Name.Relative == node.Children[0].Children[1].Data() {
						// TODO: a.b.c.d.e.f. etc completions...
						if ct, err := cache.Complete(&method.Returns[0].Type); err != nil {
							return err
						} else {
							cmp.Fields = ct.Fields
							cmp.Types = ct.Types
							cmp.Methods = ct.Methods
							return nil
						}
					}
				}
				// Found nothing. Try parents
				// TODO
			}
		case "Identifier":
			if td := findtype(cache, using, node); td == nil {
				return fmt.Errorf("Couldn't find base type with name %s", node)
			} else {
				log4go.Debug("Performing Identifier completion: %s", node)
				if ct, err := td.ToContentType(); err != nil {
					return err
				} else {
					cmp.Fields = ct.Fields
					cmp.Methods = ct.Methods
					cmp.Types = ct.Types
				}
			}
		case "DotIdentifier":
			bn := base(node)
			if td := findtype(cache, using, bn); td == nil {
				return fmt.Errorf("Couldn't find base type with name %s", bn)
			} else if fields, err := td.Fields(); err != nil {
				return err
			} else {
				for _, field := range fields {
					if field.Name.Relative == node.Children[1].Data() {
						// TODO: a.b.c.d.e.f. etc completions...
						if ct, err := cache.Complete(&field.Type); err != nil {
							return err
						} else {
							cmp.Fields = ct.Fields
							cmp.Types = ct.Types
							cmp.Methods = ct.Methods
							return nil
						}
					}
				}
				// Found nothing. Try parents
				// TODO
			}
		default:
			return fmt.Errorf("Don't know how to complete %s", node)
		}
		return nil
	}

	//	log4go.Debug("Want to complete: %s", r.Children[0])

	return complete(r.Children[0])
}

func (c *Net) Complete(args *content.CompleteArgs, cmp *content.CompletionResult) error {
	cache, err := c.cache(&args.Args)
	if err != nil {
		return err
	}
	log4go.Debug("Trying to complete: %s", args.Location)
	if ct, err := cache.Complete(&content.Type{Name: args.Location}); err != nil {
		return err
	} else {
		cmp.Fields = ct.Fields
		cmp.Types = ct.Types
		cmp.Methods = ct.Methods
	}
	return nil
}
