package net

// references:
// http://msdn.microsoft.com/en-us/library/windows/hardware/gg463119.aspx
// http://www.codeproject.com/Articles/12585/The-NET-File-Format
// http://msdn.microsoft.com/en-us/library/ms809762.aspx

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/quarnster/completion/common"
	"io"
	"reflect"
	"unsafe"
)

const (
	pe_signature_offset = 0x3c
	pe32                = 0x10b
	pe32p               = 0x20b
)

var (
	pe_magic = []byte("PE\u0000\u0000")
)

type Assembly struct {
}

type coff_file_header struct {
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
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
	Name                 [8]byte
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

func (s *section_table) String() string {
	return fmt.Sprintf("Name: %s, VirtualAddress: %d, PointerToRawData: %d", string(s.Name[0:8]), s.VirtualAddress, s.PointerToRawData)
}

func assertEndianness() error {
	test := uint32(0x01020304)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, test)
	if test2 := *(*uint32)(unsafe.Pointer(&bytes[0])); test != test2 {
		return errors.New("This module assumes a little endian system")
	}
	return nil
}

func goArray(a unsafe.Pointer, ptr uintptr, l int) {
	sliceHeader := (*reflect.SliceHeader)(a)
	sliceHeader.Cap = l
	sliceHeader.Len = l
	sliceHeader.Data = ptr
}

func LoadAssembly(r io.ReadSeeker) (*Assembly, error) {
	if err := assertEndianness(); err != nil {
		return nil, err
	}

	var (
		br        = common.BinaryReader{r, binary.LittleEndian}
		ret       Assembly
		err       error
		isPe32    bool
		data      []byte
		pe_offset uint32
	)

	if size, err := r.Seek(0, 2); err != nil {
		return nil, err
	} else if _, err := r.Seek(0, 0); err != nil {
		return nil, err
	} else if data, err = br.Read(int(size)); err != nil {
		return nil, err
	}

	r = bytes.NewReader(data)
	br.Reader = r

	if _, err := r.Seek(pe_signature_offset, 0); err != nil {
		return nil, err
	}

	if pe_offset, err = br.Uint32(); err != nil {
		return nil, err
	} else if _, err := r.Seek(int64(pe_offset), 0); err != nil {
		return nil, err
	}

	if check, err := br.Read(len(pe_magic)); err != nil {
		return nil, err
	} else if bytes.Compare(check, pe_magic) != 0 {
		return nil, errors.New(fmt.Sprintf("PE Magic mismatch: %v", check))
	}

	coff := (*coff_file_header)(unsafe.Pointer(&data[int(pe_offset)+len(pe_magic)]))

	var opt_t reflect.Type
	ptr := uintptr(unsafe.Pointer(coff))
	ptr += unsafe.Sizeof(*coff)
	opt_pe32 := (*optional_header_pe32)(unsafe.Pointer(ptr))
	opt_pe32p := (*optional_header_pe32p)(unsafe.Pointer(ptr))

	if opt_pe32.Magic != pe32 && opt_pe32.Magic != pe32p {
		return nil, errors.New(fmt.Sprintf("Unsupported optional header magic: %x", opt_pe32.Magic))
	} else {
		isPe32 = opt_pe32.Magic == pe32
	}
	rvas := 0
	if isPe32 {
		rvas = int(opt_pe32.NumberOfRvaAndSizes)
		opt_t = reflect.TypeOf(*opt_pe32)
	} else {
		rvas = int(opt_pe32p.NumberOfRvaAndSizes)
		opt_t = reflect.TypeOf(*opt_pe32p)
	}
	ptr += opt_t.Size()
	var ids []image_data_directory
	goArray(unsafe.Pointer(&ids), ptr, rvas)
	for i := range ids {
		fmt.Printf("%d, %d\n", ids[i].VirtualAddress, ids[i].Size)
	}

	ptr = uintptr(unsafe.Pointer(coff)) + unsafe.Sizeof(*coff) + uintptr(coff.SizeOfOptionalHeader)
	var sections []section_table
	goArray(unsafe.Pointer(&sections), ptr, int(coff.NumberOfSections))
	for i := range sections {
		fmt.Printf("%d, %s\n", i, &sections[i])
	}
	net := ids[14]
	off := net.VirtualAddress - sections[0].VirtualAddress + sections[0].PointerToRawData

	type image_cor20 struct {
		Size         uint32
		MajorVersion uint16
		MinorVersion uint16
		MetaData     image_data_directory
		Flags        uint32
	}
	cor20 := (*image_cor20)(unsafe.Pointer(&data[off]))

	type test struct {
		Signature    uint32
		MajorVersion uint16
		MinorVersion uint16
		Reserved     uint32
		Length       uint32
	}
	off = cor20.MetaData.VirtualAddress - sections[0].VirtualAddress + sections[0].PointerToRawData
	t := (*test)(unsafe.Pointer(&data[off]))

	r.Seek(int64(off+uint32(unsafe.Sizeof(*t))), 0)
	d, _ := br.Read(int((t.Length + 3) &^ 3))
	fmt.Println(string(d))
	flags, _ := br.Uint16()
	fmt.Println("flags:", flags)
	streams, _ := br.Uint16()
	fmt.Println("streams:", streams)
	type stream_header struct {
		Offset uint32
		Size   uint32
		Name   [128]byte
	}
	off2, _ := r.Seek(0, 1)
	ptr = uintptr(unsafe.Pointer(&data[off2]))
	for i := uint16(0); i < streams; i++ {
		sh := (*stream_header)(unsafe.Pointer(ptr))
		l := 0
		for i, v := range sh.Name {
			if v == 0 {
				l = i
				break
			}
		}
		// +1 to include the terminating 0
		ptr += 8 + uintptr((l+1+3)&^3)
		fmt.Printf("%d, %d, %d, %s\n", l, sh.Offset, sh.Size, string(sh.Name[:l]))
	}
	return &ret, nil
}
