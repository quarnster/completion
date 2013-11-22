package net

import (
	"bytes"
	"code.google.com/p/log4go"
	"errors"
	"fmt"
	"github.com/howeyc/fsnotify"
	"github.com/quarnster/completion/content"
	"io/ioutil"
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
	data, err := ioutil.ReadFile(e.name)
	if err != nil {
		return err
	}
	asm, err := LoadAssembly(bytes.NewReader(data))
	if err != nil {
		return err
	}
	e.assembly = asm
	ci := ConcreteTableIndex{metadataUtil: &asm.MetadataUtil, table: id_AssemblyRef, index: 0}
	for i := uint32(0); i < asm.Tables[id_AssemblyRef].Rows; i++ {
		ci.index = i + 1
		raw, err := ci.Data()
		if err != nil {
			return err
		}
		row := raw.(*AssemblyRefRow)
		c.load <- loadreq{name: string(row.Name)}
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
				if c.entries[i].name != req.name {
					continue
				}
				loaded = true
				if req.reload {
					if err := c.reload(&c.entries[i]); err != nil {
						log4go.Warn("Error reloading assembly:", err)
					}
				}
				break
			}
		}()
		if loaded {
			continue
		}
		exts := []string{".dll", ".exe"}
		_, err := c.Load(req.name)
		if err == nil {
			continue
		}
		for _, ext := range exts {
			if _, err := c.Load(req.name + ext); err == nil {
				break
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
			log4go.Error("error:", err)
		}
	}
}

func (c *Cache) AddPath(path string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	path = filepath.Clean(path)
	for _, p := range c.paths {
		if p == path {
			return
		}
	}
	c.paths = append(c.paths, path)
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
		}
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
	if len(errs) == 0 {
		errs = fmt.Sprintf("Assembly %s is not in path.", assembly)
	} else {
		errs = fmt.Sprintf("Could not load assembly %s:\n%s", assembly, errs)
	}
	return nil, errors.New(errs)
}

func (c *Cache) FindType(t content.FullyQualifiedName) (*TypeDef, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for i := range c.entries {
		if t2, _ := c.entries[i].assembly.FindType(t); t2 != nil {
			return t2, nil
		}
	}
	return nil, nil
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
