package net

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

const (
	testdata_path = "./testdata/"
	findtype_test = testdata_path + "7zip.dll"
)

func TestLoadAssembly(t *testing.T) {
	var (
		tests = make(map[string]string)
		err   error
	)

	testdata, err := os.Open(testdata_path)
	if err != nil {
		t.Fatalf("Failed to open %s: %s", testdata_path, err)
	}
	defer testdata.Close()

	fi, err := testdata.Readdir(0)
	if err != nil {
		t.Errorf("Failed to fully read %s: %s. Got %d useable entries.", testdata_path, err, len(fi))
	}

	for i := range fi {
		if strings.HasSuffix(fi[i].Name(), ".exe") || strings.HasSuffix(fi[i].Name(), ".dll") {
			path := testdata_path + fi[i].Name()
			data, err := ioutil.ReadFile(path + ".asm")
			if err != nil {
				t.Logf("Error reading test data: %s. Testdata will be created", err)
			}
			tests[path] = string(data)
		}
	}

	for k, v := range tests {
		t.Log("Testing", k)
		f, err := os.Open(k)
		if err != nil {
			t.Error(err)
			continue
		}
		defer f.Close()
		if asm, err := LoadAssembly(f); err != nil {
			t.Error(err)
			continue
		} else {
			var res string
			for i := range asm.Tables {
				if asm.Tables[i].Rows == 0 {
					continue
				}
				td := asm.Tables[i]
				res += fmt.Sprintln(td.RowType.Name())
				row := reflect.New(td.RowType).Interface()
				for i := uint32(0); i < td.Rows; i++ {
					data, err := td.Index(i + 1)
					if err != nil {
						t.Error(err)
						continue
					}
					if err := asm.Create(&util.BinaryReader{bytes.NewReader(data), binary.LittleEndian}, row); err != nil {
						t.Error(err)
					} else {
						res += fmt.Sprintf("\t%+v\n", row)
					}
				}
			}

			if types, err := asm.Types(); err != nil {
				t.Error(err)
			} else {
				for i := range types {
					res += fmt.Sprintf("\t%s\n", types[i])
				}
			}
			var idx = &ConcreteTableIndex{metadataUtil: &asm.MetadataUtil, index: 0, table: id_TypeDef}
			td := asm.Tables[id_TypeDef]
			//ty := reflect.New(td.RowType).Interface().(*TypeDefRow)
			for i := uint32(0); i < td.Rows; i++ {
				idx.index = i + 1
				var ty *TypeDefRow
				if tr, err := idx.Data(); err != nil {
					t.Error(err)
					continue
				} else {
					ty = tr.(*TypeDefRow)
				}
				res += fmt.Sprintln(ty.TypeName)
				if i > 0 && (ty.Flags&TypeAttributes_ClassSemanticsMask) != TypeAttributes_Interface {
					if ext, err := asm.Extends(idx); err != nil {
						t.Error(err)
					} else {
						res += fmt.Sprintf("\textends %s\n", ext)
					}
					if impl, err := asm.Implements(idx); err != nil {
						t.Error(err)
					} else {
						if len(impl) != 0 {
							res += "\timplements:\n"
						}
						for j := range impl {
							res += fmt.Sprintf("\t\t%s\n", impl[j])
						}
					}
				}
				if fields, err := asm.Fields(idx); err != nil {
					t.Error(err)
				} else {
					for j := range fields {
						res += fmt.Sprintf("\t%s\n", fields[j])
					}
				}
				if methods, err := asm.Methods(idx); err != nil {
					t.Error(err)
				} else {
					for j := range methods {
						res += fmt.Sprintf("\t%s\n", methods[j])
					}
				}
			}

			if len(v) <= 1 {
				// Just if we want to add new tests, this will spit out the newly added
				// test data
				fn := k + ".asm"
				t.Logf("Creating new test data: %s", fn)
				if err := ioutil.WriteFile(fn, []byte(res), 0644); err != nil {
					t.Errorf("Couldn't write test data to %s: %s", fn, err)
				}
				continue
			}

			if d := util.Diff(v, res); len(d) != 0 {
				t.Error(d)
			}
		}
	}
}

func TestFindType(t *testing.T) {
	type Test struct {
		in         string
		expectFind bool
	}
	var tests = []Test{
		{"SevenZip.CommandLineParser.SwitchType", true},
		{"SevenZip.Compression.LZ.OutWindow", true},
		{"SevenZip.Compression.LZ.OutWindow.Something", false},
		{"SevenZip.ISetDecoderProperties", true},
		{"Key", true},
		{"Key.Something", false},
	}
	f, err := os.Open(findtype_test)
	if err != nil {
		t.Fatalf("Failed to open %s: %s", findtype_test, err)
	}
	defer f.Close()
	if asm, err := LoadAssembly(f); err != nil {
		t.Error(err)
	} else {
		for _, test := range tests {
			if ty, err := asm.FindType(content.FullyQualifiedName{Absolute: test.in}); err != nil {
				t.Error(err)
			} else if ty == nil {
				if test.expectFind {
					t.Errorf("Expected to find type %s, but didn't", test.in)
				}
			} else if raw, err := ty.Data(); err != nil {
				t.Error(err)
			} else {
				tr := raw.(*TypeDefRow)
				if n := AbsoluteName(tr); n != test.in {
					t.Errorf("Type name mismatch: %s != %s", n, test.in)
				}
			}
		}
	}
}
