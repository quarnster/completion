package net

import (
	"fmt"
	errors "github.com/quarnster/completion/util/errors"
	"math"
	"reflect"
	"strings"
	"unsafe"
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
var (
	ErrMetadata = errors.New("Metadata header isn't in the expected format")
)
var table_row_type_lut = map[int]reflect.Type{
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

// ECMA-335 II.24.2.1
type MetadataHeader struct {
	Signature    uint32
	MajorVersion uint16
	MinorVersion uint16
	Reserved     uint32
	Length       uint32
	version      [256]byte
}

func (m *MetadataHeader) Version() string {
	return string(m.version[:m.Length])
}

func (m *MetadataHeader) StreamHeaders() []*stream_header {
	ptr := uintptr(unsafe.Pointer(&m.version))
	ptr += uintptr(m.Length+3) &^ 3
	ptr += 2 // flags
	streams := *(*uint16)(unsafe.Pointer(ptr))
	ptr += 2
	ret := make([]*stream_header, streams)
	for i := range ret {
		ret[i] = (*stream_header)(unsafe.Pointer(ptr))
		// +1 to include the terminating 0
		ptr += 8 + uintptr((len(ret[i].Name())+1+3)&^3)
	}
	return ret
}

func (m *MetadataHeader) Validate() error {
	if m.Signature != metadata_signature || m.MajorVersion != 1 || m.MinorVersion != 1 || m.Reserved != 0 {
		return ErrMetadata
	}
	return nil
}

// ECMA-335 II.24.2.2
type stream_header struct {
	Offset uint32
	Size   uint32
	name   [32]byte
}

func (s *stream_header) Name() string {
	l := 0
	for j, v := range s.name {
		if v == 0 {
			l = j
			break
		}
	}
	return string(s.name[:l])
}

func (s *stream_header) Validate() error {
	if s.name[0] != '#' {
		return errors.New(fmt.Sprintf("This does not appear to be a valid stream header: %#v", s))
	}
	return nil
}

// ECMA-335 II.24.2.6 #~ stream
type hash_tilde_stream_header struct {
	Reserved     uint32
	MajorVersion uint8
	MinorVersion uint8
	HeapSizes    uint8
	Reserved2    uint8
	Valid        uint64
	Sorted       uint64
}

type MetadataUtil struct {
	Tables     [64]MetadataTable
	StringHeap MetadataTable
	BlobHeap   MetadataTable
	GuidHeap   MetadataTable
}

func (mh *MetadataHeader) MetadataUtil() (*MetadataUtil, error) {
	var ret MetadataUtil
	off := uintptr(unsafe.Pointer(mh))

	base := off
	for _, h := range mh.StreamHeaders() {
		if err := h.Validate(); err != nil {
			return nil, err
		}
		switch h.Name() {
		case "#~":
			off += uintptr(h.Offset)
		case "#Strings":
			ret.StringHeap.Ptr = uintptr(unsafe.Pointer(base + uintptr(h.Offset)))
			ret.StringHeap.Rows = h.Size
		case "#Blob":
			ret.BlobHeap.Ptr = uintptr(unsafe.Pointer(base + uintptr(h.Offset)))
			ret.BlobHeap.Rows = h.Size
		case "#GUID":
			ret.GuidHeap.Ptr = uintptr(unsafe.Pointer(base + uintptr(h.Offset)))
			ret.GuidHeap.Rows = h.Size
		}
	}

	h := (*hash_tilde_stream_header)(unsafe.Pointer(off))
	if err := h.Validate(); err != nil {
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

	ptr := uintptr(unsafe.Pointer(h))
	ptr += reflect.TypeOf(*h).Size()
	for i := range ret.Tables {
		if valid := (h.Valid >> uint(i)) & 1; valid != 0 {
			ret.Tables[i].Rows = *(*uint32)(unsafe.Pointer(ptr))
			ret.Tables[i].RowType = table_row_type_lut[i]
			ptr += 4
		}
	}

	for i := range ret.Tables {
		if ret.Tables[i].Rows != 0 {
			ret.Tables[i].Ptr = ptr
			if size, err := ret.Size(ret.Tables[i].RowType); err != nil {
				return nil, err
			} else {
				ret.Tables[i].RowSize = uintptr(size)
				ptr += uintptr(size * uint(ret.Tables[i].Rows))
			}
		}
	}
	return &ret, nil
}

const (
	bit_stringHeapIndexSize = (1 << iota)
	bit_guidHeapIndexSize
	bit_blobHeapIndexSize
)

func (m *MetadataUtil) ReadIndex(ptr uintptr, size uint) uint32 {
	if size == 2 {
		return uint32(*(*uint16)(unsafe.Pointer(ptr)))
	}
	return *(*uint32)(unsafe.Pointer(ptr))
}

func bits(values int) uint {
	return uint(math.Ceil(math.Log2(float64(values))))
}

func (m *MetadataUtil) Create(ptr uintptr, v interface{}) (uintptr, error) {
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.Ptr {
		return 0, errors.New(fmt.Sprintf("Expected a pointer not %s", t.Kind()))
	}
	v2 := t.Elem()
	name := v2.Type().Name()

	if name == "StringIndex" {
		size := m.StringHeap.RowSize
		index := m.ReadIndex(ptr, uint(size))
		start := m.StringHeap.Ptr + uintptr(index)
		end := m.StringHeap.Ptr + uintptr(m.StringHeap.Rows)
		var data []byte

		goArray(unsafe.Pointer(&data), start, int(end))
		for i := range data {
			if data[i] == '\u0000' {
				end = uintptr(i)
				break
			}
		}
		v2.SetString(string(data[:end]))
		ptr += uintptr(size)
	} else if name == "Guid" {
		size := m.GuidHeap.RowSize
		index := m.ReadIndex(ptr, uint(size))
		if index != 0 {
			idx := m.GuidHeap.Ptr + uintptr((index-1)*16)
			v2.SetUint(uint64(idx))
		}
		ptr += uintptr(size)
	} else if strings.HasSuffix(name, "EncodedIndex") {
		if size, err := m.Size(v2.Type()); err != nil {
			return 0, err
		} else {
			var (
				idx = m.ReadIndex(ptr, size)

				tables = enc_lut[idx_name_lut[name]]
				b      = bits(len(tables))
				mask   = uint32(0xffff << b)
				tbl    = idx &^ mask
				ti     ConcreteTableIndex
			)
			idx = idx >> b
			ti.index = idx
			ti.table = tables[int(tbl)]
			ti.metadataUtil = m
			v2.Set(reflect.ValueOf(&ti))
			ptr += uintptr(size)
		}
	} else if strings.HasSuffix(name, "Index") {
		if size, err := m.Size(v2.Type()); err != nil {
			return 0, err
		} else {
			var ti ConcreteTableIndex
			ti.index = m.ReadIndex(ptr, size)
			if name == "BlobIndex" {
				ti.table = id_Blob
			} else {
				ti.table = idx_name_lut[name]
			}
			ti.metadataUtil = m
			v2.Set(reflect.ValueOf(&ti))
			ptr += uintptr(size)
		}
	} else {
		switch v2.Kind() {
		case reflect.Struct:
			for i := 0; i < v2.NumField(); i++ {
				f := v2.Field(i)
				a := f.Addr()
				if p2, err := m.Create(ptr, a.Interface()); err != nil {
					return 0, err
				} else {
					ptr = p2
				}
			}
		case reflect.Uint32:
			v2.SetUint(uint64(*(*uint32)(unsafe.Pointer(ptr))))
			ptr += 4
		case reflect.Uint16:
			v2.SetUint(uint64(*(*uint16)(unsafe.Pointer(ptr))))
			ptr += 2
		case reflect.Uint8:
			v2.SetUint(uint64(*(*uint8)(unsafe.Pointer(ptr))))
			ptr += 1
		default:
			return 0, errors.New(fmt.Sprintf("Don't know how to create %s (%s)", v2.Kind(), v2.Type()))
		}
	}
	return ptr, nil
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
			if rows<<bits(len(tables)) < 1<<16 {
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

type MetadataTable struct {
	Ptr     uintptr
	Rows    uint32
	RowType reflect.Type
	RowSize uintptr
}

func (mt *MetadataTable) Index(index uint32) (uintptr, error) {
	if mt.Ptr == 0 {
		return 0, errors.New("Trying to dereference a nil table")
	}
	index--
	if index >= mt.Rows {
		return 0, errors.New(fmt.Sprintf("Index outside of bounds: %x >= %x", index, mt.Rows))
	}
	return mt.Ptr + uintptr(index)*mt.RowSize, nil
}

func (h *hash_tilde_stream_header) Validate() error {
	if h.Reserved != 0 || h.MajorVersion != 2 || h.MinorVersion != 0 /*|| h.Reserved2 != 1*/ { //TODO: Hmm spec says Reserved2 should be 1, but it appears to be 0x10?
		return errors.New(fmt.Sprintf("This does not appear to be a valid #~ stream header: %#v", h))
	}
	return nil
}
