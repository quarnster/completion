package java

import (
	"archive/zip"
	//"os"
	"bytes"
	"io/ioutil"
	"runtime"
	"strings"
	"sync"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(8)
}

const rt_jar = "/Library/Java/JavaVirtualMachines/jdk1.7.0_09.jdk/Contents/Home/jre/lib/rt.jar"

func TestLoadAllClasses(t *testing.T) {
	if z, err := zip.OpenReader(rt_jar); err != nil {
		t.Fatal(err)
	} else {
		defer z.Close()
		var (
			inChan  = make(chan []byte, len(z.File))
			outChan = make(chan error, 32)
			wg      sync.WaitGroup
			once    sync.Once
		)
		worker := func() {
			wg.Add(1)
			for zf := range inChan {
				f := bytes.NewReader(zf)
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
		for _, zf := range z.File {
			if strings.HasSuffix(zf.Name, ".class") {
				if f, err := zf.Open(); err != nil {
					t.Error(err)
				} else {
					defer f.Close()
					if d, err := ioutil.ReadAll(f); err != nil {
						t.Error(err)
					} else {
						inChan <- d
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
}
