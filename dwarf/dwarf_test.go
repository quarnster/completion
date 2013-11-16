package dwarf

import (
	"bytes"
	"compress/bzip2"
	"debug/macho"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util"
	"github.com/quarnster/util/encoding/binary"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestAbbr(t *testing.T) {
	for _, test := range []string{"./testdata/8", "./testdata/hello", "./testdata/game.bz2", "./testdata/listener.o.bz2", "./testdata/completion.bz2"} {
		t.Logf("\n%s", test)
		rf, err := readFile(test)
		if err != nil {
			t.Fatal(err)
		}
		f, err := macho.NewFile(rf)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		sec := f.Section("__debug_abbrev")
		if sec == nil {
			for _, sec := range f.Sections {
				t.Log(sec.Name)
			}
			t.Fatal("No such section")
		}
		data, err := sec.Data()
		if err != nil {
			t.Fatal(err)
		}
		br := binary.BinaryReader{Reader: bytes.NewReader(data)}
		var (
			tag   DW_TAG
			child DW_CHILDREN
		)

		for {
			v, err := br.Uint8()
			if err != nil {
				t.Fatal(err)
			} else if v == 0 {
				break
			}
			if err := br.ReadInterface(&tag); err != nil {
				t.Fatal(err)
			}
			if err := br.ReadInterface(&child); err != nil {
				t.Fatal(err)
			}
			t.Logf("[%d] %-30s %s", v, tag, child)
			for {
				var (
					at   DW_AT
					form DW_FORM
				)
				if err := br.ReadInterface(&at); err != nil {
					t.Fatal(err)
				}
				if err := br.ReadInterface(&form); err != nil {
					t.Fatal(err)
				}
				if at == 0 {
					t.Log("\n")
					break
				}
				t.Logf("\t%-30s %s", at, form)
			}
		}
	}
}

func readFile(filename string) (io.ReaderAt, error) {
	if fs, err := os.Open(filename); err != nil {
		return nil, err
	} else {
		defer fs.Close()
		var r io.Reader = fs
		if filepath.Ext(filename) == ".bz2" {
			r = bzip2.NewReader(fs)
		}
		if data, err := ioutil.ReadAll(r); err != nil {
			return nil, err
		} else {
			// TODO: why, oh why do I need to append a byte?
			return bytes.NewReader(append(data, 0)), nil
		}
	}
}

func TestBlah(t *testing.T) {
	tests := []string{
		"./testdata/hello",
		"./testdata/8",
	}
	for _, testfile := range tests {
		if r, err := readFile(testfile); err != nil {
			t.Error(err)
		} else if dh, err := NewDWARFHelper(testfile, r); err != nil {
			t.Error(err)
		} else {
			t.Log(dh.Load())
		}
	}
}

func TestCompleteGame(t *testing.T) {
	testfile := "./testdata/game.bz2"
	tests := []string{
		"asCObjectType",
		"asCScriptFunction",
	}
	if r, err := readFile(testfile); err != nil {
		t.Error(err)
	} else if dh, err := NewDWARFHelper(testfile, r); err != nil {
		t.Error(err)
	} else {
		res := ""
		expected := ""
		if data, err := ioutil.ReadFile(testfile + ".res"); err == nil {
			expected = string(data)
		}

		for _, test := range tests {
			res += test + "\n"
			if cmp, err := dh.Complete(content.FullyQualifiedName{Relative: test}); err != nil {
				t.Error(err)
			} else {
				res += cmp.String()
			}
		}

		if len(expected) <= 1 {
			fn := testfile + ".res"
			t.Logf("Creating new test data: %s", fn)
			if err := ioutil.WriteFile(fn, []byte(res), 0644); err != nil {
				t.Errorf("Couldn't write test data to %s: %s", fn, err)
			}
		} else if d := util.Diff(expected, res); len(d) != 0 {
			t.Error(d)
		}
	}
}
