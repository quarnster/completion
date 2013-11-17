package dwarf

import (
	"bytes"
	"compress/bzip2"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

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
			if filename == "./testdata/listener.o.bz2" {
				data = append(data, 0)
			}
			return bytes.NewReader(data), nil
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
