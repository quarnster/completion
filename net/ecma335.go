package net

import (
	"errors"
	"fmt"
	"unsafe"
)

const (
	metadata_signature = 0x424A5342
)

const (
	id_nullTable              = -1
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

type MetadataUtil struct {
	HeapSizes uint8
	Tables    [64]MetadataTable
}

type MetadataTable struct {
	Ptr  uintptr
	Rows uint32
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
