package net

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/quarnster/completion/content"
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

func TestLoadAllAssemblies(t *testing.T) {
	paths := DefaultPaths()
	if len(paths) == 0 {
		t.Skip("Neither mono nor Windows .NET Framework paths were possible to get")
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
			start := time.Now()
			data, err := ioutil.ReadFile(fn)
			if err != nil {
				outChan <- errors.New(fmt.Sprintf("%s: %s\n", fn, err))
				continue
			}
			if asm, err := LoadAssembly(bytes.NewReader(data)); err != nil {
				if err != ErrNotAssembly && err != ErrMetadata {
					outChan <- errors.New(fmt.Sprintf("%s: %s\n", fn, err))
				} else {
					t.Logf("%s is not a .net assembly", fn)
				}
			} else {
				ci := ConcreteTableIndex{metadataUtil: &asm.MetadataUtil, table: id_TypeDef, index: 1}
				for i := uint32(0); i < asm.Tables[id_TypeDef].Rows; i++ {
					ci.index = i + 1
					if td, err := TypeDefFromIndex(&ci); err != nil {
						outChan <- errors.New(fmt.Sprintf("%s: %s\n", fn, err))
					} else {
						if n := td.Name(); content.Validate(&n) != nil {
							continue
						} else if _, err := td.ToContentType(); err != nil {
							outChan <- errors.New(fmt.Sprintf("%s: %s\n", fn, err))
						}
					}
				}
				ci.table = id_Module
				ci.index = 1
				if d, err := ci.Data(); err != nil {
					outChan <- errors.New(fmt.Sprintf("%s: %s\n", fn, err))
				} else {
					mr := d.(*ModuleRow)
					ci.table = id_Assembly
					if d2, err := ci.Data(); err != nil {
						if asm.Tables[id_Assembly].Rows == 0 {
							// It's ok for an assembly to not have an assembly table, although
							// I've only ever seen this with System.EnterpriseServices.Wrapper.dll
							t.Logf("Successfully loaded module %50s {%s}", mr.Name, mr.Mvid)
						} else {
							outChan <- errors.New(fmt.Sprintf("%s: %s\n", fn, err))
						}
					} else {
						ar := d2.(*AssemblyRow)
						if mn, an := string(mr.Name), string(ar.Name); !strings.HasPrefix(mn, an) && (an != "mscorlib" && mn != "CommonLanguageRuntimeLibrary") {
							outChan <- errors.New(fmt.Sprintf("The assembly name isn't the prefix of the module name: %s, %s", an, mn))
						} else {
							t.Logf("Successfully loaded module %50s {%s} %s (%s)", mn, mr.Mvid, an, time.Now().Sub(start))
						}
					}
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

	for _, path := range paths {
		if f, err := os.Open(path); err != nil {
			t.Error(err)
		} else {
			fi, err := f.Readdir(0)
			if err != nil {
				t.Fatal(err)
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
						inChan <- path + fn
						i++
					}
				} else {
					i++
				}
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
