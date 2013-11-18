package dwarf

import (
	"io"
	"testing"
)

func TestAbbrev(t *testing.T) {
	for _, test := range []string{"./testdata/8", "./testdata/hello", "./testdata/game.bz2", "./testdata/listener.o.bz2", "./testdata/completion.bz2"} {
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
		br := f.Reader("debug_line")
		for {
			var v AbbrevEntry
			before, _ := br.Seek(0, 1)
			if err := br.ReadInterface(&v); err != nil {
				if after, _ := br.Seek(0, 1); !(err == io.EOF && before+1 == after) {
					t.Errorf("%s, %d, %d, %v", err, before, after, before+1 == after)
				}
				break
			} else {
				t.Log("%+v", v)
			}
		}
	}
}
