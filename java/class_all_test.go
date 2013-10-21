package java

import (
	"bytes"
	"github.com/quarnster/completion/java/descriptors"
	"github.com/quarnster/completion/java/signatures"
	"github.com/quarnster/util/encoding/binary"
	"runtime"
	"sync"
	"testing"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func testparse(c *Class, members []member_info, method bool, t *testing.T) {
	for i := range members {
		var p descriptors.DESCRIPTORS
		desc := c.Constant_pool.Lut(members[i].Descriptor_index).String()
		if !p.Parse(desc) || p.RootNode().Range.End() != len(desc) {
			t.Errorf("Failed to parse descriptor: %s\n%s\n%s", p.Error(), desc, p.RootNode())
		}
		var p2 signatures.SIGNATURES
		for _, attr := range members[i].Attributes {
			if c.Constant_pool.Lut(attr.Attribute_name_index).String() == "Signature" {
				br := binary.BinaryReader{bytes.NewReader(attr.Info), binary.BigEndian}
				if i16, err := br.Uint16(); err != nil {
					t.Error(err)
				} else {
					sign := c.Constant_pool.Lut(u2(i16)).String()
					p2.SetData(sign)
					var ret bool
					if method {
						ret = p2.MethodTypeSignature()
					} else {
						ret = p2.FieldTypeSignature()
					}
					p2.RootNode().UpdateRange()
					if !ret || p2.RootNode().Range.End() != len(sign) {
						t.Errorf("Failed to parse signature: %s\n%s\n%s", p2.Error(), desc, p2.RootNode())
					}
				}
			}
		}
	}
}

func TestLoadAllClasses(t *testing.T) {
	classpath, err := DefaultClasspath()
	if err != nil {
		t.Skip(err)
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
				t.Log("class", c.Constant_pool.Lut(c.This_class))
				testparse(c, c.RawFields, false, t)
				testparse(c, c.RawMethods, true, t)
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
		a, err := NewArchive(cp)
		if err != nil {
			t.Logf("Failed to open archive: %s", err)
			continue
		}
		defer a.Close()
		classes := a.LoadAllClasses()
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
