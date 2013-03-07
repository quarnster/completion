package net

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/howeyc/fsnotify"
	"github.com/quarnster/completion/content"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type entry struct {
	name     string
	assembly *Assembly
}

func (e *entry) Reload() error {
	if data, err := ioutil.ReadFile(e.name); err != nil {
		return err
	} else if asm, err := LoadAssembly(bytes.NewReader(data)); err != nil {
		return err
	} else {
		e.assembly = asm
	}
	return nil
}

type Cache struct {
	entries []entry
	paths   []string
	watch   *fsnotify.Watcher
}

func (c *Cache) watchthread() {
	for {
		select {
		case ev := <-c.watch.Event:
			for i := range c.entries {
				if c.entries[i].assembly.Name() == ev.Name {
					log.Println("Reloading", ev.Name)
					if err := c.entries[i].Reload(); err != nil {
						log.Println("Error reloading assembly:", err)
					}
					break
				}
			}
		case err := <-c.watch.Error:
			log.Println("error:", err)
		}
	}
}

func (c *Cache) Load(name string) (*Assembly, error) {
	for i := range c.entries {
		if c.entries[i].assembly.Name() == name {
			if err := c.entries[i].Reload(); err != nil {
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
	if c.watch == nil {
		var err error
		c.watch, err = fsnotify.NewWatcher()
		if err != nil {
			return nil, err
		}
		go c.watchthread()
	}
	for _, p := range c.paths {
		path := filepath.Join(p, assembly)
		if _, err := os.Stat(path); err != nil {
			continue
		} else {
			if err := c.watch.Watch(path); err != nil {
				return nil, err
			}
			e := entry{name: path}
			if err := e.Reload(); err != nil {
				if len(errs) != 0 {
					errs += "\n"
				}
				errs += fmt.Sprintf("\t%s", err.Error())
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
