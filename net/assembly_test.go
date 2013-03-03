package net

import (
	"fmt"
	"github.com/quarnster/completion/util"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

const (
	testdata_path = "./testdata/"
)

func TestLoadAssembly(t *testing.T) {
	var (
		tests = make(map[string][]byte)
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
			tests[path] = data
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
					ptr, err := td.Index(i + 1)
					if err != nil {
						t.Error(err)
						continue
					}
					if _, err := asm.Create(ptr, row); err != nil {
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
			td := asm.Tables[id_TypeDef]
			ty := reflect.New(td.RowType).Interface().(*TypeDefRow)
			for i := uint32(0); i < td.Rows; i++ {
				ptr, _ := td.Index(i + 1)
				asm.Create(ptr, ty)
				res += fmt.Sprintln(ty.TypeName)
				if fields, err := asm.Fields(i + 1); err != nil {
					t.Error(err)
				} else {
					for j := range fields {
						res += fmt.Sprintf("\t%s\n", fields[j])
					}
				}
				if methods, err := asm.Methods(i + 1); err != nil {
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

			if d, err := util.Diff([]byte(v), []byte(res)); err != nil {
				t.Error(err)
			} else if len(d) != 0 {
				t.Error(string(d))
			}
		}
	}
}
