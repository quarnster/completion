package net

import (
	"code.google.com/p/log4go"
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
		if r := node.Data(); strings.HasSuffix(r, ".") {
			return r[:len(r)-1]
		} else {
			return r
		}
	default:
		if len(node.Children) > 0 {
			return resolve(node.Children[0])
		} else {
			return ""
		}
	}
}

func (c *Net) CompleteAt(args *content.CompleteAtArgs, cmp *content.CompletionResult) error {
	contents := args.Location.File.Contents
	if contents == "" {
		if d, err := ioutil.ReadFile(args.Location.File.Name); err != nil {
			return err
		} else {
			contents = string(d)
		}
	}
	lines := strings.Split(contents, "\n")

	line := lines[args.Location.Line-1]

	var p csharp.CSHARP
	p.SetData(line[:args.Location.Column-1])
	if !p.Expression() {
		return p.Error()
	}

	exp := resolve(p.RootNode())
	p.SetData(contents)
	p.UsingDirectives()
	using := p.RootNode()
	log4go.Debug("Parsed file: %s, %s, %s", exp, using)
	for _, child := range using.Children {
		var a2 content.CompleteArgs
		a2.SessionId = args.SessionId
		a2.SessionOverrides = args.SessionOverrides
		a2.Location.Absolute = "net://type/" + child.Children[0].Data() + "." + exp
		if err := c.Complete(&a2, cmp); err == nil {
			return nil
		}
	}
	return nil
}

func (c *Net) Complete(args *content.CompleteArgs, cmp *content.CompletionResult) error {
	var cache *Cache
	session := args.Session()
	if session != nil {
		if a, ok := session.Get("net_cache").(*Cache); ok {
			cache = a
		}
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
				return err
			} else {
				log4go.Debug("Found %s (%s)", lib, asm.Name())
			}
		}
		cache = &c
		if session != nil {
			session.Set("net_cache", cache)
		}
	}
	if ct, err := cache.Complete(&content.Type{Name: args.Location}); err != nil {
		return err
	} else {
		cmp.Fields = ct.Fields
		cmp.Types = ct.Types
		cmp.Methods = ct.Methods
	}
	return nil
}
