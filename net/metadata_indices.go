package net

import (
	"math"
	"unsafe"
)

type (
	TableIndex struct {
		t     *MetadataUtil
		ptr   uintptr
		table int
	}
	AssemblyRefIndex  TableIndex
	BlobIndex         TableIndex
	EventIndex        TableIndex
	FieldIndex        TableIndex
	GenericParamIndex TableIndex
	GuidIndex         TableIndex
	MethodDefIndex    TableIndex
	ModuleRefIndex    TableIndex
	ParamIndex        TableIndex
	PropertyIndex     TableIndex
	StringIndex       TableIndex
	TypeDefIndex      TableIndex
)

func (t *TableIndex) Size() uint8 {
	if t.t.Tables[t.table-1].Rows < 1<<16 {
		return 2
	}
	return 4
}

func (t *TableIndex) Index() uint32 {
	if t.Size() == 2 {
		return uint32(*(*uint16)(unsafe.Pointer(t.ptr)))
	}
	return *(*uint32)(unsafe.Pointer(t.ptr))
}

// II.24.2.6
type (
	EncodedIndex struct {
		t      *MetadataUtil
		ptr    uintptr
		encIdx int
	}
	CustomAttributeTypeEncodedIndex EncodedIndex
	HasConstantEncodedIndex         EncodedIndex
	HasCustomAttributeEncodedIndex  EncodedIndex
	HasDeclSecurityEncodedIndex     EncodedIndex
	HasFieldMarshalEncodedIndex     EncodedIndex
	HasSemanticsEncodedIndex        EncodedIndex
	ImplementationEncodedIndex      EncodedIndex
	MemberForwardedEncodedIndex     EncodedIndex
	MemberRefParentEncodedIndex     EncodedIndex
	MethodDefOrRefEncodedIndex      EncodedIndex
	ResolutionScopeEncodedIndex     EncodedIndex
	TypeDefOrRefEncodedIndex        EncodedIndex
	TypeOrMethodDefEncodedIndex     EncodedIndex
)

const (
	idx_TypeDefOrRef = iota
	idx_HasConstant
	idx_HasCustomAttribute
	idx_HasFieldMarshal
	idx_HasDeclSecurity
	idx_MemberRefParent
	idx_HasSemantics
	idx_MethodDefOrRef
	idx_MemberForwarded
	idx_Implementation
	idx_CustomAttributeType
	idx_ResolutionScope
	idx_TypeOrMethodDef
)

var enc_lut = [][]int{
	// TypeDefOrRef
	[]int{
		id_TypeDef,
		id_TypeRef,
		id_TypeSpec},
	// HasConstant
	[]int{
		id_Field,
		id_Param,
		id_Property},
	// HasCustomAttribute
	[]int{
		id_MethodDef,
		id_Field,
		id_TypeRef,
		id_TypeDef,
		id_Param,
		id_InterfaceImpl,
		id_MemberRef,
		id_Module,
		id_DeclSecurity,
		id_Property,
		id_Event,
		id_StandAloneSig,
		id_ModuleRef,
		id_TypeSpec,
		id_Assembly,
		id_AssemblyRef,
		id_File,
		id_ExportedType,
		id_ManifestResource,
		id_GenericParam,
		id_GenericParamConstraint,
		id_MethodSpec},
	// HasFieldMarshal
	[]int{
		id_Field,
		id_Param},
	// HasDeclSecurity
	[]int{
		id_TypeDef,
		id_MethodDef,
		id_Assembly},
	// MemberRefParent
	[]int{
		id_TypeDef,
		id_TypeRef,
		id_ModuleRef,
		id_MethodDef,
		id_TypeSpec},
	// HasSemantics
	[]int{
		id_Event,
		id_Property},
	// MethodDefOrRef
	[]int{
		id_MethodDef,
		id_MemberRef},
	// MemberForwarded
	[]int{
		id_Field,
		id_MethodDef},
	// Implementation
	[]int{
		id_File,
		id_AssemblyRef,
		id_ExportedType},
	// CustomAttributeType
	[]int{
		id_nullTable,
		id_nullTable,
		id_MethodDef,
		id_MemberRef,
		id_nullTable},
	// ResolutionScope
	[]int{
		id_Module,
		id_ModuleRef,
		id_AssemblyRef,
		id_TypeRef},
	// TypeOrMethodDef
	[]int{
		id_TypeDef,
		id_MethodDef,
	},
}

func bits(values int) uint {
	return uint(math.Ceil(math.Log2(float64(values))))
}

func (t *EncodedIndex) RawIndex() uint32 {
	if t.Size() == 2 {
		return uint32(*(*uint16)(unsafe.Pointer(t.ptr)))
	}
	return *(*uint32)(unsafe.Pointer(t.ptr))
}

func (id *EncodedIndex) Table() int {
	var (
		ret    = id.RawIndex()
		tables = enc_lut[id.encIdx]
		b      = bits(len(tables))
		mask   = uint32(0xffff << b)
	)
	return tables[ret&^mask]
}

func (id *EncodedIndex) Index() uint32 {
	var (
		ret    = id.RawIndex()
		tables = enc_lut[id.encIdx]
		b      = bits(len(tables))
	)
	return ret >> b
}

func (id *EncodedIndex) Size() uint8 {
	var (
		tables = enc_lut[id.encIdx]
		rows   uint32
	)
	for _, t := range tables {
		if s2 := id.t.Tables[t].Rows; s2 > rows {
			rows = s2
		}
	}
	if rows<<bits(len(tables)) < 1<<16 {
		return 2
	}
	return 4
}
