package content

import (
	"github.com/quarnster/completion/util"
	"reflect"
	"testing"
)

func TestFlags(t *testing.T) {
	type td struct {
		count int
		bits  uint
		name  string
	}
	tests := []td{
		{_FLAG_ACC_COUNT, FLAG_ACC_BITS, "Access"},
		{_FLAG_TYPE_COUNT, FLAG_TYPE_BITS, "Type"},
		{_FLAG_LANG_COUNT, FLAG_LANG_BITS, "Language"},
		{1 << _FLAG_BOOL_COUNT, _FLAG_BOOL_COUNT, "Bool"},
	}
	totSize := 0
	for _, test := range tests {
		if current, actual := test.bits, util.Bits(test.count); actual != current {
			t.Errorf("Bit mismatch for %s bits: %d != %d", test.name, current, actual)
		}
		totSize += int(test.bits)
	}
	var f Flags
	ty := reflect.TypeOf(f)
	if bits := ty.Bits(); totSize > bits {
		t.Errorf("The number of bits used does not fit in %s's size: %d > %d", ty.Name(), totSize, bits)
	}
	t.Logf("Bits used: %d", totSize)
}
