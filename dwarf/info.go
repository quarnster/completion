package dwarf

import (
	"fmt"
	"github.com/quarnster/util/encoding/binary"
	"reflect"
)

type (
	Header struct {
		is64    bool
		Length  uint64
		Version uint16
	}
	InfoHeader struct {
		Header
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

func (ih *Header) Read(br *binary.BinaryReader) error {
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
	return err
}

func (ih *InfoHeader) Read(br *binary.BinaryReader) error {
	err := br.ReadInterface(&ih.Header)
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

func (ie *InfoEntry) data(form DW_FORM, br binary.BinaryReader) interface{} {
	if form == DW_FORM_ref_addr && ie.header.Version < 3 {
		form = DW_FORM_addr
	}
	switch form {
	case DW_FORM_flag_present:
		return true
	case DW_FORM_exprloc, DW_FORM_block:
		var size LEB128
		br.ReadInterface(&size)
		r, _ := br.Read(int(size))
		return r
	case DW_FORM_block1:
		size, _ := br.Uint8()
		r, _ := br.Read(int(size))
		return r
	case DW_FORM_block2:
		size, _ := br.Uint16()
		r, _ := br.Read(int(size))
		return r
	case DW_FORM_block4:
		size, _ := br.Uint32()
		r, _ := br.Read(int(size))
		return r
	case DW_FORM_addr:
		if ie.header.AddressSize == 8 {
			v, _ := br.Uint64()
			return v
		} else {
			v, _ := br.Uint32()
			return uint64(v)
		}
	case DW_FORM_ref_addr, DW_FORM_strp, DW_FORM_sec_offset:
		if ie.header.is64 {
			v, _ := br.Uint64()
			return v
		} else {
			v, _ := br.Uint32()
			return uint64(v)
		}
	case DW_FORM_ref1, DW_FORM_flag, DW_FORM_data1:
		v, _ := br.Uint8()
		return uint64(v)
	case DW_FORM_ref2, DW_FORM_data2:
		v, _ := br.Uint16()
		return uint64(v)
	case DW_FORM_ref4, DW_FORM_data4:
		v, _ := br.Uint32()
		return uint64(v)
	case DW_FORM_ref8, DW_FORM_data8:
		v, _ := br.Uint64()
		return v
	case DW_FORM_sdata, DW_FORM_udata:
		var r LEB128
		br.ReadInterface(&r)
		return uint64(r)
	case DW_FORM_string:
		buf := make([]byte, 4096)

		for i := range buf {
			if v, err := br.Uint8(); err != nil {
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
	panic(fmt.Errorf("Unimplemented format: %s", form))
}

func (ie *InfoEntry) Attribute(name DW_AT) interface{} {
	if _, err := ie.reader.Info.Seek(ie.debugInfoOffset, 0); err != nil {
		return nil
	}
	for _, attr := range ie.ae.Attributes {
		d := ie.data(attr.Form, ie.reader.Info)
		if d == nil {
			return nil
		}
		if name != attr.Name {
			continue
		}
		switch attr.Form {
		case DW_FORM_strp:
			ie.reader.Str.Seek(int64(reflect.ValueOf(d).Uint()), 0)
			return ie.data(DW_FORM_string, ie.reader.Str)
		}
		switch attr.Name {
		case DW_AT_encoding:
			return DW_ATE(reflect.ValueOf(d).Uint())
		case DW_AT_accessibility:
			return DW_ACCESS(reflect.ValueOf(d).Uint())
		case DW_AT_virtuality:
			return DW_VIRTUALITY(reflect.ValueOf(d).Uint())
		case DW_AT_language:
			return DW_LANG(reflect.ValueOf(d).Uint())
		}
		return d
	}
	return nil
}
