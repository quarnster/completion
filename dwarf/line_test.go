package dwarf

import (
	"io"
	"testing"
)

func TestLine(t *testing.T) {
	for _, test := range []string{"./testdata/hello4", "./testdata/hello", "./testdata/game.bz2", "./testdata/listener.o.bz2", "./testdata/completion.bz2", "./testdata/hello4"} {
		t.Logf("\n%s", test)
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
		for {
			var v lineHeader
			if err := bri.ReadInterface(&v); err != nil {
				if err == io.EOF {
					break
				}
				t.Error(err)
			}
			//			t.Logf("%+v", v)
			for i, v := range v.standard_opcode_lengths {
				t.Logf("standard_opcode_lengths[%s] = %d", DW_LNS(i+1), v)
			}
			for i, v := range v.file_names {
				t.Logf("%d: %+v", (i + 1), v)
			}
			for i, v := range v.matrix {
				t.Logf("%d: %+v", (i + 1), v)
			}
		}
	}
}
