package net

import (
	"errors"
	"fmt"
	"unsafe"
)

const (
	metadata_signature = 0x424A5342
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
	Reserverd2   uint8
	Valid        uint64
	Sorted       uint64
	Rows         [64]uint32
}
