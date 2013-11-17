package dwarf

import (
	"github.com/quarnster/util/encoding/binary"
)

type (
	InfoHeader struct {
		is64              bool
		Length            uint64
		Version           uint16
		DebugAbbrevOffset int64
		AddressSize       uint8
	}

	InfoEntry struct {
		header          *InfoHeader
		debugInfoOffset int64
		id              LEB128
		ae              *AbbrevEntry
		reader          dwarfReader
	}

	dwarfReader struct {
		Abbrev binary.BinaryReader
		Info   binary.BinaryReader
		Str    binary.BinaryReader
	}
)

func (ih *InfoHeader) Read(br *binary.BinaryReader) error {
	if v, err := br.Uint32(); err != nil {
		return err
	} else if v != 0xffffffff {
		ih.Length = uint64(v)
	} else if v, err := br.Uint64(); err != nil {
		return err
	} else {
		ih.is64 = true
		ih.Length = v
	}
	var err error
	ih.Version, err = br.Uint16()
	if err != nil {
		return err
	}
	if ih.is64 {
		ih.DebugAbbrevOffset, err = br.Int64()
	} else {
		var v uint32
		v, err = br.Uint32()
		ih.DebugAbbrevOffset = int64(v)
	}
	if err != nil {
		return err
	}
	ih.AddressSize, err = br.Uint8()
	return err
}

func (ie *InfoEntry) Tag() DW_TAG {
	return ie.ae.Tag
}

func (ie *InfoEntry) Read(br *binary.BinaryReader) error {
	ret := br.ReadInterface(&ie.id)
	ie.debugInfoOffset, _ = br.Seek(0, 1)
	return ret
}

func (ie *InfoEntry) data(form DW_FORM) interface{} {
	switch form {
	case DW_FORM_block:
		var size LEB128
		ie.reader.Info.ReadInterface(&size)
		r, _ := ie.reader.Info.Read(int(size))
		return r
	case DW_FORM_block1:
		size, _ := ie.reader.Info.Uint8()
		r, _ := ie.reader.Info.Read(int(size))
		return r
	case DW_FORM_block2:
		size, _ := ie.reader.Info.Uint16()
		r, _ := ie.reader.Info.Read(int(size))
		return r
	case DW_FORM_block4:
		size, _ := ie.reader.Info.Uint32()
		r, _ := ie.reader.Info.Read(int(size))
		return r
	case DW_FORM_addr:
		if ie.header.AddressSize == 8 {
			v, _ := ie.reader.Info.Uint64()
			return v
		} else {
			v, _ := ie.reader.Info.Uint32()
			return v
		}
	case DW_FORM_ref_addr, DW_FORM_strp:
		if ie.header.is64 {
			v, _ := ie.reader.Info.Uint64()
			return v
		} else {
			v, _ := ie.reader.Info.Uint32()
			return v
		}
	case DW_FORM_ref1, DW_FORM_flag, DW_FORM_data1:
		v, _ := ie.reader.Info.Uint8()
		return v
	case DW_FORM_ref2, DW_FORM_data2:
		v, _ := ie.reader.Info.Uint16()
		return v
	case DW_FORM_ref4, DW_FORM_data4:
		v, _ := ie.reader.Info.Uint32()
		return v
	case DW_FORM_ref8, DW_FORM_data8:
		v, _ := ie.reader.Info.Uint64()
		return v
	case DW_FORM_sdata, DW_FORM_udata:
		var r LEB128
		ie.reader.Info.ReadInterface(&r)
		return r
	case DW_FORM_string:
		buf := make([]byte, 4096)

		for i := range buf {
			if v, err := ie.reader.Info.Uint8(); err != nil {
				return err
			} else if v == 0 {
				buf = buf[:i]
				break
			} else {
				buf[i] = byte(v)
			}
		}
		return string(buf)
	}

	return nil
}

func (ie *InfoEntry) Attribute(name DW_AT) interface{} {
	if _, err := ie.reader.Info.Seek(ie.debugInfoOffset, 0); err != nil {
		return nil
	}
	for _, attr := range ie.ae.Attributes {
		d := ie.data(attr.Form)
		if d == nil || name == attr.Name {
			return d
		}
	}
	return nil
}
