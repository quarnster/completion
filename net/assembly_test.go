package net

import (
	"bytes"
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util"
	"github.com/quarnster/util/encoding/binary"
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
					if err := asm.Create(&binary.BinaryReader{bytes.NewReader(data), binary.LittleEndian}, row); err != nil {
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
			tdTable := asm.Tables[id_TypeDef]
			//ty := reflect.New(td.RowType).Interface().(*TypeDefRow)
			for i := uint32(0); i < tdTable.Rows; i++ {
				idx.index = i + 1
				td, err := TypeDefFromIndex(idx)
				if err != nil {
					t.Error(err)
					continue
				} else if n := td.Name(); check(&n, n) != nil {
					continue
				}

				if ct, err := td.ToContentType(); err != nil {
					t.Error(err)
					continue
				} else {
					res += fmt.Sprintln(ct)
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
		{"net://type/SevenZip.CommandLineParser.SwitchType", true},
		{"net://type/SevenZip.Compression.LZ.OutWindow", true},
		{"net://type/SevenZip.Compression.LZ.OutWindow.Something", false},
		{"net://type/SevenZip.ISetDecoderProperties", true},
		{"net://type/Key", false},
		{"net://type/LzmaAlone$Key", true},
		{"net://type/LzmaAlone$Key.Something", false},
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
			} else {
				if n := ty.Name().Absolute; n != test.in || !test.expectFind {
					t.Errorf("Type name mismatch: %s != %s", n, test.in)
				}
			}
		}
	}
}

func BenchmarkComplete(b *testing.B) {
	b.StopTimer()
	f, err := os.Open(findtype_test)
	if err != nil {
		b.Fatalf("Failed to open %s: %s", findtype_test, err)
	}
	defer f.Close()
	if asm, err := LoadAssembly(f); err != nil {
		b.Error(err)
	} else {
		b.StartTimer()
		tn := content.Type{Name: content.FullyQualifiedName{Absolute: "net://type/SevenZip.Compression.LZ.OutWindow"}}
		for i := 0; i < b.N; i++ {
			if _, err := asm.Complete(&tn); err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkMetatableLookup(b *testing.B) {
	b.StopTimer()
	f, err := os.Open(findtype_test)
	if err != nil {
		b.Fatalf("Failed to open %s: %s", findtype_test, err)
	}
	defer f.Close()
	if asm, err := LoadAssembly(f); err != nil {
		b.Error(err)
	} else if ty, err := asm.FindType(content.FullyQualifiedName{Absolute: "net://type/SevenZip.Compression.LZ.OutWindow"}); err != nil {
		b.Error(err)
	} else {
		idx := ty.index
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			if raw, err := idx.Data(); err != nil {
				b.Error(err)
			} else {
				if _, ok := raw.(*TypeDefRow); !ok {
					b.Error("It's not a TypeDefRow")
				}
			}
		}
	}
}
