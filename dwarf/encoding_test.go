package dwarf

import (
	"bytes"
	"github.com/quarnster/util/encoding/binary"
	"testing"
)

func TestLEB128Read(t *testing.T) {
	tests := []struct {
		output LEB128
		input  []byte
	}{
		{2, []byte{2}},
		{127, []byte{127}},
		{128, []byte{0 + 0x80, 1}},
		{129, []byte{1 + 0x80, 1}},
		{130, []byte{2 + 0x80, 1}},
		{12857, []byte{57 + 0x80, 100}},
	}
	for _, test := range tests {
		br := binary.BinaryReader{Reader: bytes.NewReader(test.input)}
		var v LEB128
		if err := br.ReadInterface(&v); err != nil {
			t.Error(err)
		} else if v != test.output {
			t.Errorf("%d != %d", v, test.output)
		}
	}
}
