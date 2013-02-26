package net

import (
	"errors"
	"fmt"
	"unsafe"
)

const (
	pe_signature_offset = 0x3c
	pe32                = 0x10b
	pe32p               = 0x20b
)

type coff_file_header struct {
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
}

func (coff *coff_file_header) SectionTable() (sections []section_table) {
	ptr := uintptr(unsafe.Pointer(coff)) + unsafe.Sizeof(*coff) + uintptr(coff.SizeOfOptionalHeader)
	goArray(unsafe.Pointer(&sections), ptr, int(coff.NumberOfSections))
	return sections
}

type optional_header_common_1 struct {
	Magic                   uint16
	MajorLinkerVersion      uint8
	MinorLinkerVersion      uint8
	SizeOfCode              uint32
	SizeOfInitializedData   uint32
	SizeOfUninitializedData uint32
	AddressOfEntryPoint     uint32
	BaseOfCode              uint32
}
type optional_header_common_2 struct {
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
}
type optional_header_common_3 struct {
	LoaderFlags         uint32
	NumberOfRvaAndSizes uint32
}

type optional_header_pe32 struct {
	optional_header_common_1
	BaseOfData uint32
	ImageBase  uint32
	optional_header_common_2
	SizeOfStackReserve uint32
	SizeOfStackCommit  uint32
	SizeOfHeapReserve  uint32
	SizeOfHeapCommit   uint32
	optional_header_common_3
}

type optional_header_pe32p struct {
	optional_header_common_1
	ImageBase uint64
	optional_header_common_2
	SizeOfStackReserve uint64
	SizeOfStackCommit  uint64
	SizeOfHeapReserve  uint64
	SizeOfHeapCommit   uint64
	optional_header_common_3
}

type image_data_directory struct {
	VirtualAddress uint32
	Size           uint32
}

type section_table struct {
	name                 [8]byte
	VirtualSize          uint32
	VirtualAddress       uint32
	SizeOfRawData        uint32
	PointerToRawData     uint32
	PointerToRelocations uint32
	PointerToLineNumbers uint32
	NumberOfRelocations  uint16
	NumberOfLinenumbers  uint16
	Characteristics      uint32
}

func (s *section_table) Validate() error {
	if s.name[0] != '.' {
		return errors.New(fmt.Sprintf("This does not appear to be a valid section header: %#v", s))
	}
	return nil
}

func (s *section_table) Name() string {
	return string(s.name[:])
}

type image_cor20 struct {
	Size         uint32
	MajorVersion uint16
	MinorVersion uint16
	MetaData     image_data_directory
	Flags        uint32
}

func (s *section_table) String() string {
	return fmt.Sprintf("Name: %s, VirtualAddress: %d, PointerToRawData: %d", s.Name(), s.VirtualAddress, s.PointerToRawData)
}

var (
	pe_magic = []byte("PE\u0000\u0000")
)
