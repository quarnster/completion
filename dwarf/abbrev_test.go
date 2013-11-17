package dwarf

import (
	"bytes"
	"debug/macho"
	"github.com/quarnster/util/encoding/binary"
	"io"
	"testing"
)

func TestAbbrev(t *testing.T) {
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
		for {
			var v AbbrevEntry
			before, _ := br.Seek(0, 1)
			if err := br.ReadInterface(&v); err != nil {
				if after, _ := br.Seek(0, 1); !(err == io.EOF && before+1 == after) {
					t.Errorf("%s, %d, %d, %v", err, before, after, before+1 == after)
				}
				break
			} else {
				t.Logf("%+v", v)
			}
		}
	}
}
