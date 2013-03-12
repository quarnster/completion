package dwarf

import (
	"testing"
)

const testfile = "./testdata/hello"

func TestBlah(t *testing.T) {
	if dh, err := NewDWARFHelper(testfile); err != nil {
		t.Error(err)
	} else {
		t.Log(dh.Load())
	}
}
