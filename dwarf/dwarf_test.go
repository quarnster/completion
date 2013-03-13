package dwarf

import (
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
