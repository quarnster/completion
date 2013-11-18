package dwarf

import (
	"fmt"
	"github.com/quarnster/completion/util"
	"io"
	"io/ioutil"
	"testing"
)

func TestLine(t *testing.T) {
	for _, test := range []string{"./testdata/8", "./testdata/hello", "./testdata/game.bz2", "./testdata/listener.o.bz2", "./testdata/completion.bz2", "./testdata/hello4"} {
		rf, err := readFile(test)
		if err != nil {
			t.Error(err)
			continue
		}
		f, err := newSectionReader(rf)
		if err != nil {
			t.Error(err)
			continue
		}
		defer f.Close()
		bri := f.Reader("debug_line")
		res := ""
		for {
			var v lineHeader
			if err := bri.ReadInterface(&v); err != nil {
				if err == io.EOF {
					break
				}
				t.Error(err)
			}
			res += fmt.Sprintf("%s\n", v)
		}
		expected := ""
		fn := test + ".line"
		if data, err := ioutil.ReadFile(fn); err == nil {
			expected = string(data)
		}
		if len(expected) <= 1 {
			t.Logf("Creating new test data: %s", fn)
			if err := ioutil.WriteFile(fn, []byte(res), 0644); err != nil {
				t.Errorf("Couldn't write test data to %s: %s", fn, err)
			}
		} else if d := util.Diff(expected, res); len(d) != 0 {
			t.Error(d)
		}
	}
}
