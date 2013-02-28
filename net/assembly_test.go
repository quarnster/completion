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

const testdata_path = "./testdata/"

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
		if strings.HasSuffix(fi[i].Name(), ".asm") {
			path := testdata_path + fi[i].Name()
			if data, err := ioutil.ReadFile(path); err != nil {
				t.Errorf("Error reading test data: %s", err)
			} else {
				tests[strings.Replace(path, ".asm", "", -1)] = data
			}
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
