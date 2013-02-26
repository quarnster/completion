package net

// references:
// http://msdn.microsoft.com/en-us/library/windows/hardware/gg463119.aspx
// http://www.codeproject.com/Articles/12585/The-NET-File-Format
// http://msdn.microsoft.com/en-us/library/ms809762.aspx
// http://www.ecma-international.org/publications/standards/Ecma-335.htm

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

type Assembly struct {
}

type Validateable interface {
	Validate() error
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

	sections := coff.SectionTable()
	for i := range sections {
		if err := sections[i].Validate(); err != nil {
			return nil, err
		}
		fmt.Println(sections[i].Name())
	}
	net := ids[14]
	off := net.VirtualAddress - sections[0].VirtualAddress + sections[0].PointerToRawData

	cor20 := (*image_cor20)(unsafe.Pointer(&data[off]))

	off = cor20.MetaData.VirtualAddress - sections[0].VirtualAddress + sections[0].PointerToRawData
	t := (*Metadata)(unsafe.Pointer(&data[off]))
	if err := t.Validate(); err != nil {
		return nil, err
	}

	fmt.Println(t.Version())
	for _, h := range t.StreamHeaders() {
		if err := h.Validate(); err != nil {
			return nil, err
		}
		fmt.Println(h.Name())
		if h.Name() == "#~" {
			off += h.Offset
		}
	}
	ht := (*hash_tilde_stream_header)(unsafe.Pointer(&data[off]))
	fmt.Printf("%#v\n", ht)
	if err := ht.Validate(); err != nil {
		return nil, err
	}

	return &ret, nil
}
