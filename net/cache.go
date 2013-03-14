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
	"sync"
)

type (
	entry struct {
		name     string
		assembly *Assembly
	}
	loadreq struct {
		name   string
		reload bool
	}
	Cache struct {
		entries []entry
		paths   []string
		watch   *fsnotify.Watcher
		mutex   sync.Mutex
		load    chan loadreq
	}
)

func (c *Cache) reload(e *entry) error {
	if data, err := ioutil.ReadFile(e.name); err != nil {
		return err
	} else if asm, err := LoadAssembly(bytes.NewReader(data)); err != nil {
		return err
	} else {
		e.assembly = asm
		ci := ConcreteTableIndex{metadataUtil: &asm.MetadataUtil, table: id_AssemblyRef, index: 0}
		for i := uint32(0); i < asm.Tables[id_AssemblyRef].Rows; i++ {
			ci.index = i + 1
			if raw, err := ci.Data(); err != nil {
				return err
			} else {
				row := raw.(*AssemblyRefRow)
				c.load <- loadreq{name: string(row.Name)}
			}
		}
	}
	return nil
}

func (c *Cache) loaderthread() {
	for req := range c.load {
		loaded := false
		func() {
			c.mutex.Lock()
			defer c.mutex.Unlock()
			for i := range c.entries {
				if c.entries[i].name == req.name {
					loaded = true
					if req.reload {
						if err := c.reload(&c.entries[i]); err != nil {
							log.Println("Error reloading assembly:", err)
						}
					}
					break
				}
			}
		}()
		if !loaded {
			exts := []string{".dll", ".exe"}
			if _, err := c.Load(req.name); err != nil {
				for _, ext := range exts {
					if _, err := c.Load(req.name + ext); err == nil {
						break
					}
				}
			}
		}
	}
}

func (c *Cache) watchthread() {
	for {
		select {
		case ev := <-c.watch.Event:
			c.load <- loadreq{ev.Name, true}
		case err := <-c.watch.Error:
			log.Println("error:", err)
		}
	}
}

func (c *Cache) Load(name string) (*Assembly, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for i := range c.entries {
		if c.entries[i].assembly.Name() == name {
			return c.entries[i].assembly, nil
		}
	}
	return c.cache(name)
}

func (c *Cache) cache(assembly string) (*Assembly, error) {
	errs := ""
	if c.watch == nil {
		c.load = make(chan loadreq, 64)
		var err error
		c.watch, err = fsnotify.NewWatcher()
		if err != nil {
			return nil, err
		}
		go c.watchthread()
		go c.loaderthread()
	}
	for _, p := range c.paths {
		path := filepath.Join(p, assembly)
		if _, err := os.Stat(path); err != nil {
			continue
		} else {
			if err := c.watch.Watch(path); err != nil {
				return nil, err
			} else if err := c.watch.WatchFlags(path, fsnotify.FSN_MODIFY); err != nil {
				return nil, err
			}
			e := entry{name: path}
			if err := c.reload(&e); err != nil {
				if len(errs) != 0 {
					errs += "\n"
				}
				errs += fmt.Sprintf("\t%s", err.Error())
				c.watch.RemoveWatch(path)
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
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for i := range c.entries {
		if ret, err := c.entries[i].assembly.Complete(t); err == nil {
			return ret, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Couldn't find type %s", t))
}
