package net

import (
	"errors"
	"fmt"
	"math"
	"unsafe"
)

const (
	metadata_signature = 0x424A5342
)

const (
	id_nullTable              = 0
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
)

// ECMA-335 II.24.2.1
type Metadata struct {
	Signature    uint32
	MajorVersion uint16
	MinorVersion uint16
	Reserved     uint32
	Length       uint32
	version      [256]byte
}

func (m *Metadata) Version() string {
	return string(m.version[:m.Length])
}

func (m *Metadata) StreamHeaders() []*stream_header {
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

func (m *Metadata) Validate() error {
	if m.Signature != metadata_signature || m.MajorVersion != 1 || m.MinorVersion != 1 || m.Reserved != 0 {
		return errors.New(fmt.Sprintf("Metadata header isn't in the expected format: %#v", m))
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
	Rows         [64]uint32
}

func (h *hash_tilde_stream_header) StringHeapIndexSize() uintptr {
	if h.HeapSizes&0x01 != 0 {
		return 4
	}
	return 2
}

func (h *hash_tilde_stream_header) GUIDHeapIndexSize() uintptr {
	if h.HeapSizes&0x02 != 0 {
		return 4
	}
	return 2
}

func (h *hash_tilde_stream_header) BlobHeapIndexSize() uintptr {
	if h.HeapSizes&0x04 != 0 {
		return 4
	}
	return 2
}

func (h *hash_tilde_stream_header) TableIndexSize(table int) uintptr {
	return 0 // TODO
}

func (h *hash_tilde_stream_header) TableSize(table int) uintptr {
	return 0 // TODO
}

func (h *hash_tilde_stream_header) Validate() error {
	if h.Reserved != 0 || h.MajorVersion != 2 || h.MinorVersion != 0 /*|| h.Reserved2 != 1*/ { //TODO: Hmm spec says Reserved2 should be 1, but it appears to be 0x10?
		return errors.New(fmt.Sprintf("This does not appear to be a valid #~ stream header: %#v", h))
	}
	return nil
}

const (
	encidx_TypeDefOrRef = iota
	encidx_HasConstant
	encidx_HasCustomAttribute
	encidx_HasFieldMarshall
	encidx_HasDeclSecurity
	encidx_MemberRefParent
	encidx_HasSemantics
	encidx_MethodDefOrRef
	encidx_MemberForwarded
	encidx_Implementation
	encidx_CustomAttributeType
	encidx_ResolutionScope
	encidx_TypeOrMethodDef
)

var encidx_tables = map[int][]int{
	encidx_TypeDefOrRef: []int{
		id_TypeDef,
		id_TypeRef,
		id_TypeSpec},
	encidx_HasConstant: []int{
		id_Field,
		id_Param,
		id_Property},
	encidx_HasCustomAttribute: []int{
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
	encidx_HasFieldMarshall: []int{
		id_Field,
		id_Param},
	encidx_HasDeclSecurity: []int{
		id_TypeDef,
		id_MethodDef,
		id_Assembly},
	encidx_MemberRefParent: []int{
		id_TypeDef,
		id_TypeRef,
		id_ModuleRef,
		id_MethodDef,
		id_TypeSpec},
	encidx_HasSemantics: []int{
		id_Event,
		id_Property},
	encidx_MethodDefOrRef: []int{
		id_MethodDef,
		id_MemberRef},
	encidx_MemberForwarded: []int{
		id_Field,
		id_MethodDef},
	encidx_Implementation: []int{
		id_File,
		id_AssemblyRef,
		id_ExportedType},
	encidx_CustomAttributeType: []int{
		id_nullTable,
		id_nullTable,
		id_MethodDef,
		id_MemberRef,
		id_nullTable},
	encidx_ResolutionScope: []int{
		id_Module,
		id_ModuleRef,
		id_AssemblyRef,
		id_TypeRef},
	encidx_TypeOrMethodDef: []int{
		id_TypeDef,
		id_MethodDef,
	},
}

type row_size_func func(*hash_tilde_stream_header) uintptr

func bits(values int) uint {
	return uint(math.Ceil(math.Log2(float64(values))))
}

func encoded_index_size(h *hash_tilde_stream_header, tables []int) uintptr {
	s := uintptr(0)
	for _, t := range tables {
		if s2 := h.TableSize(t); s2 > s {
			s = s2
		}
	}
	if s < (1 << (16 - bits(len(tables)))) {
		return 2
	}
	return 4
}

var row_size_lut = map[int]row_size_func{
	id_Assembly: func(h *hash_tilde_stream_header) uintptr {
		return 4 + 4*2 + 4 + h.BlobHeapIndexSize() + h.StringHeapIndexSize()
	},
	id_AssemblyOS:        func(h *hash_tilde_stream_header) uintptr { return 3 * 4 },
	id_AssemblyProcessor: func(h *hash_tilde_stream_header) uintptr { return 4 },
	id_AssemblyRef: func(h *hash_tilde_stream_header) uintptr {
		return 4*2 + 4 + 2*h.BlobHeapIndexSize() + 2*h.StringHeapIndexSize()
	},
	id_AssemblyRefOS:        func(h *hash_tilde_stream_header) uintptr { return 3*4 + h.TableIndexSize(id_AssemblyRef) },
	id_AssemblyRefProcessor: func(h *hash_tilde_stream_header) uintptr { return 4 + h.TableIndexSize(id_AssemblyRef) },
	id_ClassLayout:          func(h *hash_tilde_stream_header) uintptr { return 2 + 4 + h.TableIndexSize(id_TypeDef) },
	id_Constant: func(h *hash_tilde_stream_header) uintptr {
		return 1 + 1 + encoded_index_size(h, encidx_tables[encidx_HasConstant]) + h.BlobHeapIndexSize()
	},
	id_CustomAttribute:        func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_DeclSecurity:           func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_EventMap:               func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_Event:                  func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_ExportedType:           func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_Field:                  func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_FieldLayout:            func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_FieldMarshal:           func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_FieldRVA:               func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_File:                   func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_GenericParam:           func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_GenericParamConstraint: func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_ImplMap:                func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_InterfaceImpl:          func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_ManifestResource:       func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_MemberRef:              func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_MethodDef:              func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_MethodImpl:             func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_MethodSemantics:        func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_MethodSpec:             func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_Module:                 func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_ModuleRef:              func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_NestedClass:            func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_Param:                  func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_Property:               func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_PropertyMap:            func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_StandAloneSig:          func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_TypeDef:                func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_TypeRef:                func(h *hash_tilde_stream_header) uintptr { return 0 },
	id_TypeSpec:               func(h *hash_tilde_stream_header) uintptr { return 0 }}

// II.22.2 Assembly : 0x20
type table_row_Assembly struct {
	HashAlgId                                               uint32
	MajorVersion, MinorVersion, BuildNumber, RevisionNumber uint16
	Flags                                                   uint32
	// PublicKey blobIndex
	// Name stringIndex
}

func (t *table_row_Assembly) Size(s *hash_tilde_stream_header) uintptr {
	return unsafe.Sizeof(*t) + s.BlobHeapIndexSize() + s.StringHeapIndexSize()
}

// II.22.3 AssemblyOS : 0x22
type table_row_AssemblyOS struct {
	OSPlatformID   uint32
	OSMajorVersion uint32
	OSMinorVersion uint32
}

func (t *table_row_AssemblyOS) Size(s *hash_tilde_stream_header) uintptr {
	return unsafe.Sizeof(*t)
}

// II.22.4 AssemblyProcessor : 0x21
type table_row_AssemblyProcessor struct {
	Processor uint32
}
