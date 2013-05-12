package net

import (
	"code.google.com/p/log4go"
	"github.com/quarnster/completion/content"
)

type Net struct {
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
