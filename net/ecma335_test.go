package net

import "testing"

func TestBits(t *testing.T) {
	tests := map[int]uint{
		encidx_TypeDefOrRef:        2,
		encidx_HasConstant:         2,
		encidx_HasCustomAttribute:  5,
		encidx_HasFieldMarshall:    1,
		encidx_HasDeclSecurity:     2,
		encidx_MemberRefParent:     3,
		encidx_HasSemantics:        1,
		encidx_MethodDefOrRef:      1,
		encidx_MemberForwarded:     1,
		encidx_Implementation:      2,
		encidx_CustomAttributeType: 3,
		encidx_ResolutionScope:     2,
		encidx_TypeOrMethodDef:     1,
	}

	for k, v := range tests {
		if b := bits(len(encidx_tables[k])); b != v {
			t.Errorf("Mismatched bitcount for table %d: %d != %d (%d possible tables)", k, b, v, len(encidx_tables[k]))
		}
	}
}
