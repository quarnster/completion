package net

import (
	"bytes"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func TestLoadMonoAssemblies(t *testing.T) {
	mp := MonoDefaultPath()
	if mp == "" {
		t.Skip("Unable to find mono")
	}
	if f, err := os.Open(mp); err != nil {
		t.Error(err)
	} else {
		fi, err := f.Readdir(0)
		if err != nil {
			t.Fatal(err)
		}
		var (
			inChan  = make(chan string, 64)
			outChan = make(chan error, 32)
			wg      sync.WaitGroup
			once    sync.Once
		)
		worker := func() {
			wg.Add(1)

			for fn := range inChan {
				data, err := ioutil.ReadFile(mp + fn)
				if err != nil {
					outChan <- err
					continue
				}
				if asm, err := LoadAssembly(bytes.NewReader(data)); err != nil {
					outChan <- err
				} else {
					ci := ConcreteTableIndex{metadataUtil: &asm.MetadataUtil, table: id_Module, index: 1}
					if d, err := ci.Data(); err != nil {
						outChan <- err
					} else {
						mr := d.(*ModuleRow)
						t.Logf("Successfully loaded module %50s {%s}", mr.Name, mr.Mvid)
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

		for i := 0; i < len(fi); {
			if fn := fi[i].Name(); strings.HasSuffix(fn, ".dll") {
				if len(inChan)+1 >= cap(inChan) {
					for len(outChan) > 0 {
						o := <-outChan
						if o != nil {
							t.Error(o)
						}
					}
					time.Sleep(time.Microsecond * 100)
				} else {
					inChan <- fn
					i++
				}
			} else {
				i++
			}
		}
		close(inChan)
		for o := range outChan {
			if o != nil {
				t.Error(o)
			}
		}
	}
}
