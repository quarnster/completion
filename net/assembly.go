package net

// references:
// http://msdn.microsoft.com/en-us/library/windows/hardware/gg463119.aspx
// http://www.codeproject.com/Articles/12585/The-NET-File-Format

import (
	"bytes"
	"common"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const (
	pe_signature_offset             = 0x3c
	pe32                            = 0x10b
	pe32p                           = 0x20b
	image_base_offset_pe32          = 28
	image_base_offset_pe32p         = 24
	clr_runtime_header_offset_pe32  = 208 - 2                             // 208 - 2
	clr_runtime_header_offset_pe32p = clr_runtime_header_offset_pe32 + 16 // 224 - 2
)

var (
	pe_magic = []byte("PE\u0000\u0000")
)

type Assembly struct {
}

func LoadAssembly(r io.ReadSeeker) (*Assembly, error) {
	var (
		br               = common.BinaryReader{r, binary.LittleEndian}
		ret              Assembly
		clr_runtime_seek = clr_runtime_header_offset_pe32
		clr_addr         uint32
		clr_size         uint32
		err              error
		isPe32           bool
		imageBase        uint64
		pe_offset        uint32
	)
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
	if _, err := r.Seek(20, 1); err != nil {
		return nil, err
	}
	if magic, err := br.Uint16(); err != nil {
		return nil, err
	} else if magic != pe32 && magic != pe32p {
		return nil, errors.New(fmt.Sprintf("Unsupported optional header magic: %x", magic))
	} else {
		isPe32 = magic == pe32
	}
	curr, err := r.Seek(0, 1)
	if isPe32 {
		if _, err := r.Seek(image_base_offset_pe32, 1); err != nil {
			return nil, err
		}
		if imageBase, err = br.Uint64(); err != nil {
			return nil, err
		}
	} else {
		if _, err := r.Seek(image_base_offset_pe32p, 1); err != nil {
			return nil, err
		}
		if u32, err := br.Uint32(); err != nil {
			return nil, err
		} else {
			imageBase = uint64(u32)
		}

		clr_runtime_seek = clr_runtime_header_offset_pe32p
	}
	if _, err := r.Seek(curr, 0); err != nil {
		return nil, err
	} else if _, err := r.Seek(int64(clr_runtime_seek), 1); err != nil {
		return nil, err
	}

	clr_addr, err = br.Uint32()
	if err != nil {
		return nil, err
	}
	clr_size, err = br.Uint32()
	if err != nil {
		return nil, err
	}
	fmt.Printf("%x, %x, %x, %x\n", imageBase, clr_addr, clr_size, pe_offset)
	if _, err := r.Seek(int64(clr_addr+pe_offset), 0); err != nil {
		return nil, err
	}
	if magic, err := br.Read(8); err != nil {
		return nil, err
	} else {
		fmt.Printf("%s\n", magic)
	}

	return &ret, nil
}
