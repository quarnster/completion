package dwarf

import (
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util"
	"io/ioutil"
	"testing"
)

func TestBlah(t *testing.T) {
	tests := []string{
		"./testdata/hello",
		"./testdata/8",
	}
	for _, testfile := range tests {
		if dh, err := NewDWARFHelper(testfile); err != nil {
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
	if dh, err := NewDWARFHelper(testfile); err != nil {
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
