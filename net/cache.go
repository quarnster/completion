package net

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/quarnster/completion/content"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type entry struct {
	name     string
	modtime  time.Time
	assembly *Assembly
}

func (e *entry) Reload() error {
	if data, err := ioutil.ReadFile(e.name); err != nil {
		return err
	} else if asm, err := LoadAssembly(bytes.NewReader(data)); err != nil {
		return err
	} else {
		e.modtime = time.Now()
		e.assembly = asm
	}
	return nil
}

func (e *entry) Update() error {
	if fi, err := os.Stat(e.name); err != nil {
		return err
	} else if fi.ModTime() != e.modtime {
		return e.Reload()
	}
	return nil
}

type Cache struct {
	entries []entry
	paths   []string
}

func (c *Cache) Load(name string) (*Assembly, error) {
	for i := range c.entries {
		if c.entries[i].assembly.Name() == name {
			if err := c.entries[i].Update(); err != nil {
				return nil, err
			} else {
				return c.entries[i].assembly, nil
			}
		}
	}
	return c.Cache(name)
}

func (c *Cache) Cache(assembly string) (*Assembly, error) {
	errs := ""
	for _, p := range c.paths {
		path := filepath.Join(p, assembly)
		if _, err := os.Stat(path); err != nil {
			continue
		} else {
			e := entry{name: path}
			if err := e.Reload(); err != nil {
				if len(errs) != 0 {
					errs += "\n"
				}
				errs += err.Error()
				continue
			}
			c.entries = append(c.entries, e)
			return e.assembly, nil
		}
	}
	if len(errs) == 0 {
		errs = fmt.Sprintf("Assembly %s is not in path.", assembly)
	} else {
		errs = fmt.Sprintf("Could not load assembly %s:\n%s", assembly, errs)
	}
	return nil, errors.New(errs)
}

func (c *Cache) Complete(t *content.Type) (*content.CompletionResult, error) {
	for i := range c.entries {
		if ret, err := c.entries[i].assembly.Complete(t); err == nil {
			return ret, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Couldn't find type %s", t))
}
