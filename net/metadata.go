package net

import (
	"fmt"
	"github.com/quarnster/completion/util"
	errors "github.com/quarnster/completion/util/errors"
	"github.com/quarnster/util/encoding/binary"
	"reflect"
	"sort"
	"strings"
)

const (
	metadata_signature = 0x424A5342
)

const (
	id_Assembly               = 0x20 // II.22.2
	id_AssemblyOS             = 0x22 // II.22.3
	id_AssemblyProcessor      = 0x21 // II.22.4
	id_AssemblyRef            = 0x23 // II.22.5
	id_AssemblyRefOS          = 0x25 // II.22.6
	id_AssemblyRefProcessor   = 0x24 // II.22.7
	id_ClassLayout            = 0x0F // II.22.8
	id_Constant               = 0x0B // II.22.9
	id_CustomAttribute        = 0x0C // II.22.10
	id_DeclSecurity           = 0x0E // II.22.11
	id_EventMap               = 0x12 // II.22.12
	id_Event                  = 0x14 // II.22.13
	id_ExportedType           = 0x27 // II.22.14
	id_Field                  = 0x04 // II.22.15
	id_FieldLayout            = 0x10 // II.22.16
	id_FieldMarshal           = 0x0D // II.22.17
	id_FieldRVA               = 0x1D // II.22.18
	id_File                   = 0x26 // II.22.19
	id_GenericParam           = 0x2A // II.22.20
	id_GenericParamConstraint = 0x2C // II.22.21
	id_ImplMap                = 0x1C // II.22.22
	id_InterfaceImpl          = 0x09 // II.22.23
	id_ManifestResource       = 0x28 // II.22.24
	id_MemberRef              = 0x0A // II.22.25
	id_MethodDef              = 0x06 // II.22.26
	id_MethodImpl             = 0x19 // II.22.27
	id_MethodSemantics        = 0x18 // II.22.28
	id_MethodSpec             = 0x2B // II.22.29
	id_Module                 = 0x00 // II.22.30
	id_ModuleRef              = 0x1A // II.22.31
	id_NestedClass            = 0x29 // II.22.32
	id_Param                  = 0x08 // II.22.33
	id_Property               = 0x17 // II.22.34
	id_PropertyMap            = 0x15 // II.22.35
	id_StandAloneSig          = 0x11 // II.22.36
	id_TypeDef                = 0x02 // II.22.37
	id_TypeRef                = 0x01 // II.22.38
	id_TypeSpec               = 0x1B // II.22.39
	id_nullTable              = 0x100
	id_Blob                   = 0x101
)

const (
	bit_stringHeapIndexSize = (1 << iota)
	bit_guidHeapIndexSize
	bit_blobHeapIndexSize
)

var (
	ErrMetadata        = errors.New("Metadata header isn't in the expected format")
	table_row_type_lut = map[int]reflect.Type{
		id_Assembly:               reflect.TypeOf(AssemblyRow{}),
		id_AssemblyOS:             reflect.TypeOf(AssemblyOSRow{}),
		id_AssemblyProcessor:      reflect.TypeOf(AssemblyProcessorRow{}),
		id_AssemblyRef:            reflect.TypeOf(AssemblyRefRow{}),
		id_AssemblyRefOS:          reflect.TypeOf(AssemblyRefOSRow{}),
		id_AssemblyRefProcessor:   reflect.TypeOf(AssemblyRefProcessorRow{}),
		id_ClassLayout:            reflect.TypeOf(ClassLayoutRow{}),
		id_Constant:               reflect.TypeOf(ConstantRow{}),
		id_CustomAttribute:        reflect.TypeOf(CustomAttributeRow{}),
		id_DeclSecurity:           reflect.TypeOf(DeclSecurityRow{}),
		id_EventMap:               reflect.TypeOf(EventMapRow{}),
		id_Event:                  reflect.TypeOf(EventRow{}),
		id_ExportedType:           reflect.TypeOf(ExportedTypeRow{}),
		id_Field:                  reflect.TypeOf(FieldRow{}),
		id_FieldLayout:            reflect.TypeOf(FieldLayoutRow{}),
		id_FieldMarshal:           reflect.TypeOf(FieldMarshalRow{}),
		id_FieldRVA:               reflect.TypeOf(FieldRVARow{}),
		id_File:                   reflect.TypeOf(FileRow{}),
		id_GenericParam:           reflect.TypeOf(GenericParamRow{}),
		id_GenericParamConstraint: reflect.TypeOf(GenericParamConstraintRow{}),
		id_ImplMap:                reflect.TypeOf(ImplMapRow{}),
		id_InterfaceImpl:          reflect.TypeOf(InterfaceImplRow{}),
		id_ManifestResource:       reflect.TypeOf(ManifestResourceRow{}),
		id_MemberRef:              reflect.TypeOf(MemberRefRow{}),
		id_MethodDef:              reflect.TypeOf(MethodDefRow{}),
		id_MethodImpl:             reflect.TypeOf(MethodImplRow{}),
		id_MethodSemantics:        reflect.TypeOf(MethodSemanticsRow{}),
		id_MethodSpec:             reflect.TypeOf(MethodSpecRow{}),
		id_Module:                 reflect.TypeOf(ModuleRow{}),
		id_ModuleRef:              reflect.TypeOf(ModuleRefRow{}),
		id_NestedClass:            reflect.TypeOf(NestedClassRow{}),
		id_Param:                  reflect.TypeOf(ParamRow{}),
		id_Property:               reflect.TypeOf(PropertyRow{}),
		id_PropertyMap:            reflect.TypeOf(PropertyMapRow{}),
		id_StandAloneSig:          reflect.TypeOf(StandAloneSigRow{}),
		id_TypeDef:                reflect.TypeOf(TypeDefRow{}),
		id_TypeRef:                reflect.TypeOf(TypeRefRow{}),
		id_TypeSpec:               reflect.TypeOf(TypeSpecRow{}),
	}
)

type (
	// ECMA-335 II.24.2.1
	MetadataHeader struct {
		Signature     uint32
		MajorVersion  uint16
		MinorVersion  uint16
		Reserved      uint32
		Version       string `length:"uint32" align:"4"`
		Flags         uint16
		StreamHeaders []stream_header `length:"uint16"`
	}

	// ECMA-335 II.24.2.2
	stream_header struct {
		Offset uint32
		Size   uint32
		Name   string `max:"32" align:"4"`
	}

	// ECMA-335 II.24.2.6 #~ stream
	hash_tilde_stream_header struct {
		Reserved     uint32
		MajorVersion uint8
		MinorVersion uint8
		HeapSizes    uint8
		Reserved2    uint8
		Valid        uint64
		Sorted       uint64
	}

	MetadataUtil struct {
		Tables     [64]MetadataTable
		StringHeap MetadataTable
		BlobHeap   MetadataTable
		GuidHeap   MetadataTable
	}

	MetadataTable struct {
		data    []byte
		Rows    uint32
		RowType reflect.Type
		RowSize uint32
	}
)

func (mh *MetadataHeader) MetadataUtil(br *binary.BinaryReader) (*MetadataUtil, error) {
	var (
		ret MetadataUtil
	)

	off, err := br.Seek(0, 1)
	if err != nil {
		return nil, err
	}
	base := off
	for _, h := range mh.StreamHeaders {
		switch h.Name {
		case "#~":
			off += int64(h.Offset)
		case "#Strings":
			if _, err := br.Seek(base+int64(h.Offset), 0); err != nil {
				return nil, err
			} else if ret.StringHeap.data, err = br.Read(int(h.Size)); err != nil {
				return nil, err
			}
			ret.StringHeap.Rows = h.Size
		case "#Blob":
			if _, err := br.Seek(base+int64(h.Offset), 0); err != nil {
				return nil, err
			} else if ret.BlobHeap.data, err = br.Read(int(h.Size)); err != nil {
				return nil, err
			}
			ret.BlobHeap.Rows = h.Size
		case "#GUID":
			if _, err := br.Seek(base+int64(h.Offset), 0); err != nil {
				return nil, err
			} else if ret.GuidHeap.data, err = br.Read(int(h.Size)); err != nil {
				return nil, err
			}
			ret.GuidHeap.Rows = h.Size
		}
	}
	if _, err := br.Seek(off, 0); err != nil {
		return nil, err
	}
	h := hash_tilde_stream_header{}
	if err := br.ReadInterface(&h); err != nil {
		return nil, err
	}

	if h.HeapSizes&bit_stringHeapIndexSize != 0 {
		ret.StringHeap.RowSize = 4
	} else {
		ret.StringHeap.RowSize = 2
	}
	if h.HeapSizes&bit_blobHeapIndexSize != 0 {
		ret.BlobHeap.RowSize = 4
	} else {
		ret.BlobHeap.RowSize = 2
	}
	if h.HeapSizes&bit_guidHeapIndexSize != 0 {
		ret.GuidHeap.RowSize = 4
	} else {
		ret.GuidHeap.RowSize = 2
	}

	for i := range ret.Tables {
		if valid := (h.Valid >> uint(i)) & 1; valid == 0 {
			continue
		}
		if ret.Tables[i].Rows, err = br.Uint32(); err != nil {
			return nil, err
		}
		ret.Tables[i].RowType = table_row_type_lut[i]
	}

	for i := range ret.Tables {
		if ret.Tables[i].Rows == 0 {
			continue
		}
		size, err := ret.Size(ret.Tables[i].RowType)
		if err != nil {
			return nil, err
		}
		ret.Tables[i].RowSize = uint32(size)
		if ret.Tables[i].data, err = br.Read(int(ret.Tables[i].RowSize * ret.Tables[i].Rows)); err != nil {
			return nil, err
		}
	}
	return &ret, nil
}

func (m *MetadataUtil) ReadIndex(br *binary.BinaryReader, size uint) (uint32, error) {
	if size != 2 {
		return br.Uint32()
	} else if v, e := br.Uint16(); e != nil {
		return 0, e
	} else {
		return uint32(v), nil
	}
}

func (m *MetadataUtil) Create(br *binary.BinaryReader, v interface{}) error {
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("Expected a pointer not %s", t.Kind()))
	}
	v2 := t.Elem()
	name := v2.Type().Name()

	if name == "StringIndex" {
		size := m.StringHeap.RowSize
		index, err := m.ReadIndex(br, uint(size))
		if err != nil {
			return err
		}
		data := m.StringHeap.data[index:m.StringHeap.Rows]

		for i := range data {
			if data[i] == '\u0000' {
				data = data[:i]
				break
			}
		}
		v2.SetString(string(data))
	} else if name == "Guid" {
		size := m.GuidHeap.RowSize
		if index, err := m.ReadIndex(br, uint(size)); err != nil {
			return err
		} else if index != 0 {
			index = (index - 1) * 16
			g := Guid(m.GuidHeap.data[index : index+16])
			v2.Set(reflect.ValueOf(g))
		}
	} else if strings.HasSuffix(name, "EncodedIndex") {
		size, err := m.Size(v2.Type())
		if err != nil {
			return err
		}
		idx, err := m.ReadIndex(br, size)
		if err != nil {
			return err
		}
		var (
			tables = enc_lut[idx_name_lut[name]]
			b      = util.Bits(len(tables))
			mask   = uint32(0xffff << b)
			tbl    = idx &^ mask
			ti     ConcreteTableIndex
		)
		idx = idx >> b
		ti.index = idx
		ti.table = tables[int(tbl)]
		ti.metadataUtil = m
		v2.Set(reflect.ValueOf(&ti))
	} else if strings.HasSuffix(name, "Index") {
		size, err := m.Size(v2.Type())
		if err != nil {
			return err
		}
		var ti ConcreteTableIndex
		if ti.index, err = m.ReadIndex(br, size); err != nil {
			return err
		}
		if name == "BlobIndex" {
			ti.table = id_Blob
		} else {
			ti.table = idx_name_lut[name]
		}
		ti.metadataUtil = m
		v2.Set(reflect.ValueOf(&ti))
	} else {
		if v2.Kind() != reflect.Struct {
			return br.ReadInterface(v)
		}
		for i := 0; i < v2.NumField(); i++ {
			f := v2.Field(i)
			a := f.Addr()
			if err := m.Create(br, a.Interface()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *MetadataUtil) Size(t reflect.Type) (uint, error) {
	size := uint(0)
	name := t.Name()
	switch name {
	case "StringIndex":
		size = uint(m.StringHeap.RowSize)
	case "Guid":
		size = uint(m.GuidHeap.RowSize)
	case "BlobIndex":
		size = uint(m.BlobHeap.RowSize)
	default:
		if strings.HasSuffix(name, "EncodedIndex") {
			id := idx_name_lut[name]
			var (
				tables = enc_lut[id]
				rows   uint32
			)
			for _, t := range tables {
				if t == id_nullTable {
					continue
				}
				if s2 := m.Tables[t].Rows; s2 > rows {
					rows = s2
				}
			}
			if rows<<util.Bits(len(tables)) < 1<<16 {
				size = 2
			} else {
				size = 4
			}
		} else if strings.HasSuffix(name, "Index") {
			if m.Tables[idx_name_lut[name]].Rows < 1<<16 {
				size = 2
			} else {
				size = 4
			}
		} else {
			switch t.Kind() {
			case reflect.Struct:
				for i := 0; i < t.NumField(); i++ {
					f := t.Field(i)
					if s, err := m.Size(f.Type); err != nil {
						return 0, err
					} else {
						size += s
					}
				}
			case reflect.Uint8:
				return 1, nil
			case reflect.Uint16:
				return 2, nil
			case reflect.Uint32:
				return 4, nil
			case reflect.Uint64:
				return 8, nil
			default:
				return 0, errors.New(fmt.Sprintf("Don't know the size of: %s, %s", t.Name(), t.Kind()))
			}
		}
	}
	return size, nil
}

func (mu *MetadataUtil) Search(tableId int, equal func(TableIndex) bool) TableIndex {
	mt := mu.Tables[tableId]
	ci := ConcreteTableIndex{metadataUtil: mu, index: 0, table: tableId}
	idx := sort.Search(int(mt.Rows), func(in int) bool {
		i := uint32(in)
		ci.index = i + 1
		return equal(&ci)
	})
	if uint32(idx) == mt.Rows {
		return nil
	}
	ci.index = uint32(idx) + 1
	return &ci
}

func (mt *MetadataTable) Index(index uint32) ([]byte, error) {
	if mt.Rows == 0 {
		return nil, errors.New("Trying to dereference a nil table")
	}
	index--
	if index >= mt.Rows {
		return nil, errors.New(fmt.Sprintf("Index outside of bounds: %x >= %x", index, mt.Rows))
	}
	return mt.data[index*mt.RowSize:], nil
}
