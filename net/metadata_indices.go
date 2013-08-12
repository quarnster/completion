package net

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/quarnster/util/encoding/binary"
	"reflect"
)

type (
	ConcreteTableIndex struct {
		metadataUtil *MetadataUtil
		index        uint32
		table        int
	}
	TableIndex interface {
		Table() int
		Index() uint32
		Data() (interface{}, error)
		String() string
	}
	AssemblyRefIndex  TableIndex
	BlobIndex         TableIndex
	EventIndex        TableIndex
	FieldIndex        TableIndex
	GenericParamIndex TableIndex
	MethodDefIndex    TableIndex
	ModuleRefIndex    TableIndex
	ParamIndex        TableIndex
	PropertyIndex     TableIndex
	StringIndex       string
	TypeDefIndex      TableIndex
)

func (t *ConcreteTableIndex) Index() uint32 {
	return t.index
}

func (t *ConcreteTableIndex) Table() int {
	return t.table
}

func (t *ConcreteTableIndex) Data() (interface{}, error) {
	switch t.table {
	case id_nullTable:
		return nil, errors.New("This is a null table")
	case id_Blob:
		return t.metadataUtil.BlobHeap.Index(t.index)
	}
	var (
		err   error
		table = t.metadataUtil.Tables[t.table]
		data  []byte
	)
	if data, err = table.Index(t.index); err != nil {
		return nil, err
	}
	ret := reflect.New(table.RowType).Interface()
	if err := t.metadataUtil.Create(&binary.BinaryReader{Reader: bytes.NewReader(data), Endianess: binary.LittleEndian}, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (t *ConcreteTableIndex) String() string {
	switch t.table {
	case id_nullTable:
		return "nulltable"
	case id_Blob:
		return fmt.Sprintf("Blob[%d]", t.index)
	}

	return fmt.Sprintf("%s[%d]", table_row_type_lut[t.table].Name(), t.index)
}

// II.24.2.6
type (
	EncodedIndex                    TableIndex
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

var idx_name_lut = map[string]int{
	"AssemblyRefIndex":                id_AssemblyRef,
	"EventIndex":                      id_Event,
	"FieldIndex":                      id_Field,
	"GenericParamIndex":               id_GenericParam,
	"MethodDefIndex":                  id_MethodDef,
	"ModuleRefIndex":                  id_ModuleRef,
	"ParamIndex":                      id_Param,
	"PropertyIndex":                   id_Property,
	"TypeDefIndex":                    id_TypeDef,
	"BlobIndex":                       id_Blob,
	"TypeDefOrRefEncodedIndex":        idx_TypeDefOrRef,
	"HasConstantEncodedIndex":         idx_HasConstant,
	"HasCustomAttributeEncodedIndex":  idx_HasCustomAttribute,
	"HasFieldMarshalEncodedIndex":     idx_HasFieldMarshal,
	"HasDeclSecurityEncodedIndex":     idx_HasDeclSecurity,
	"MemberRefParentEncodedIndex":     idx_MemberRefParent,
	"HasSemanticsEncodedIndex":        idx_HasSemantics,
	"MethodDefOrRefEncodedIndex":      idx_MethodDefOrRef,
	"MemberForwardedEncodedIndex":     idx_MemberForwarded,
	"ImplementationEncodedIndex":      idx_Implementation,
	"CustomAttributeTypeEncodedIndex": idx_CustomAttributeType,
	"ResolutionScopeEncodedIndex":     idx_ResolutionScope,
	"TypeOrMethodDefEncodedIndex":     idx_TypeOrMethodDef,
}

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
