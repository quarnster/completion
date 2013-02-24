// +build extensive_testing

package java

import (
	"bytes"
	"runtime"
	"sync"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func TestLoadAllClasses(t *testing.T) {
	classpath, err := DefaultClasspath()
	if err != nil {
		t.Fatal(err)
	}
	type lookup struct {
		path string
		cn   Classname
	}
	var (
		classmap = ClasspathMap(classpath)
		inChan   = make(chan lookup, 5000)
		outChan  = make(chan error, 32)
		wg       sync.WaitGroup
		once     sync.Once
	)
	worker := func() {
		wg.Add(1)

		for l := range inChan {
			if d, err := LoadClassEx(l.path, l.cn); err != nil {
				outChan <- err
			} else {
				f := bytes.NewReader(d)
				if c, err := NewClass(f); err != nil {
					outChan <- err
				} else {
					t.Log("class", String(c.Constant_pool, c.This_class))
				}
			}
		}
		wg.Done()
		wg.Wait()
		once.Do(func() {
			close(outChan)
		})
	}
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker()
	}

	for cp, classes := range classmap {
		for i := range classes {
			inChan <- lookup{cp, classes[i]}
		}
	}
	close(inChan)
	for o := range outChan {
		if o != nil {
			t.Error(o)
		}
	}
}
