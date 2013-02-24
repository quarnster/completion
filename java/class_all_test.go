package java

import (
	"bytes"
	"runtime"
	"sync"
	"testing"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func TestLoadAllClasses(t *testing.T) {
	classpath, err := DefaultClasspath()
	if err != nil {
		t.Fatal(err)
	}
	var (
		inChan  = make(chan []byte, 128)
		outChan = make(chan error, 32)
		wg      sync.WaitGroup
		once    sync.Once
	)
	worker := func() {
		wg.Add(1)

		for d := range inChan {
			f := bytes.NewReader(d)
			if c, err := NewClass(f); err != nil {
				outChan <- err
			} else {
				t.Log("class", String(c.Constant_pool, c.This_class))
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

	for _, cp := range classpath {
		classes := LoadAllClassesInPath(cp)
		for i := 0; i < len(classes); {
			if len(inChan)+1 >= cap(inChan) {
				for len(outChan) > 0 {
					o := <-outChan
					if o != nil {
						t.Error(o)
					}
				}
				time.Sleep(time.Microsecond * 100)
			} else {
				inChan <- classes[i]
				i++
			}
		}
	}
	close(inChan)
	for o := range outChan {
		if o != nil {
			t.Error(o)
		}
	}
}
