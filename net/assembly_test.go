package net

import (
	"os"
	"reflect"
	"testing"
)

const System_dll = "/Library/Frameworks/Mono.framework/Libraries/mono/4.0/System.dll"

func TestLoadAssembly(t *testing.T) {
	if f, err := os.Open(System_dll); err != nil {
		t.Error(err)
	} else {
		defer f.Close()
		if asm, err := LoadAssembly(f); err != nil {
			t.Error(err)
		} else {
			td := asm.Tables[id_TypeDef]
			ptr := td.Ptr
			row := reflect.New(td.RowType).Interface().(*TypeDefRow)
			size, err := asm.Size(td.RowType)
			if err != nil {
				t.Fatal(err)
			}
			for i := uint32(0); i < td.Rows; i++ {
				if _, err := asm.Create(ptr, row); err != nil {
					t.Error(err)
				} else {
					t.Logf("%#v", row)
				}
				ptr += uintptr(size)
			}
			t.Log(td)
		}
	}
}
