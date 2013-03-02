package net

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/quarnster/completion/common"
	"io"
	"reflect"
)

// II.23.2
func encodedWidth(first uint8) int {
	if first&(1<<7) == 0 {
		// one byte value
		return 1
	} else if first&(1<<6) == 0 {
		// two byte value
		return 2
	} else if first&(1<<5) == 0 {
		// four byte value
		return 4
	}
	return -1
}

func read4pop(reader io.ReadSeeker) ([]byte, error) {
	var data = make([]byte, 4)
	if currpos, err := reader.Seek(0, 1); err != nil {
		return nil, err
	} else if n, err := reader.Read(data); err != nil {
		return nil, err
	} else if n != 4 {
		return nil, errors.New(fmt.Sprintf("Didn't read the expected amount: %d != %d", n, 4))
	} else if _, err = reader.Seek(currpos, 0); err != nil {
		return nil, err
	}
	return data, nil
}

func UnsignedDecode(reader io.ReadSeeker) (uint32, error) {
	data, err := read4pop(reader)
	if err != nil {
		return 0, err
	}
	w := encodedWidth(data[0])
	if _, err := reader.Seek(int64(w), 1); err != nil {
		return 0, err
	}
	switch w {
	case 1:
		return uint32(data[0]), nil
	case 2:
		return uint32(binary.BigEndian.Uint16(data)) &^ 0x8000, nil
	case 4:
		return binary.BigEndian.Uint32(data) &^ 0xc0000000, nil
	}
	return 0, errors.New(fmt.Sprintf("Invalid bitmask signature: %x", data[0]))
}

func SignedDecode(reader io.ReadSeeker) (ret int32, err error) {
	data, err := read4pop(reader)
	if err != nil {
		return 0, err
	}

	if d, err := UnsignedDecode(reader); err != nil {
		return 0, err
	} else {
		sign := (d & 1)
		switch encodedWidth(data[0]) {
		case 1:
			sign <<= 7
			sign |= sign >> 1
			return int32(int8(d>>1 | sign)), nil
		case 2:
			sign <<= 15
			sign |= sign >> 1
			sign |= sign >> 1
			return int32(int16(d>>1 | sign)), nil
		case 4:
			sign <<= 31
			sign |= sign >> 1
			sign |= sign >> 1
			sign |= sign >> 1
			return int32(int32(d>>1 | sign)), nil
		}
	}
	panic("unreachable")
}

const (
	SIGNATURE_HASTHIS      = 0x20
	SIGNATURE_EXPLICITTHIS = 0x40
	SIGNATURE_DEFAULT      = 0x0
	SIGNATURE_VARARG       = 0x5
	SIGNATURE_GENERIC      = 0x10
	SIGNATURE_C            = 0x1
	SIGNATURE_STDCALL      = 0x2
	SIGNATURE_THISCALL     = 0x3
	SIGNATURE_FASTCALL     = 0x4
	SIGNATURE_SENTINEL     = 0x41
	SIGNATURE_FIELD        = 0x6
)

type EncUint uint32
type EncInt int32

type Decoder struct {
	Reader common.BinaryReader
}

func (d *Decoder) Decode(v interface{}) error {
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("Expected a pointer not %s", t.Kind()))
	}
	v2 := t.Elem()
	switch v2.Type().Name() {
	case "MethodDefSig":
		var err error
		m := v.(*MethodDefSig)
		if m.Id, err = d.Reader.Uint8(); err != nil {
			return err
		}
		if m.Id&SIGNATURE_GENERIC != 0 {
			if err := d.Decode(&m.GenParamCount); err != nil {
				return err
			}
		}
		var paramCount EncUint
		if err := d.Decode(&paramCount); err != nil {
			return err
		}
		if err := d.Decode(&m.RetType); err != nil {
			return err
		}
		m.Params = make([]Param, paramCount)
		for i := range m.Params {
			if err = d.Decode(&m.Params[i]); err != nil {
				return err
			}
		}
		return nil
	case "FieldSig":
		if id, err := d.Reader.Uint8(); err != nil {
			return err
		} else if id != SIGNATURE_FIELD {
			return errors.New(fmt.Sprintf("Expected type SIGNATURE_FIELD (%x) not %x", SIGNATURE_FIELD, id))
		}
	}

	switch v2.Kind() {
	case reflect.Int32:
		if v, err := SignedDecode(d.Reader.Reader); err != nil {
			return err
		} else {
			v2.SetInt(int64(v))
		}
	case reflect.Uint32:
		if v, err := UnsignedDecode(d.Reader.Reader); err != nil {
			return err
		} else {
			v2.SetUint(uint64(v))
		}
	}
	return nil
}

// II.23.2.1 MethodDefSig
type MethodDefSig struct {
	Id            uint8
	GenParamCount EncUint
	RetType       RetType
	Params        []Param `count:"ParamCount"`
}

// II.23.2.4 FieldSig
type FieldSig struct {
	CustomMod []CustomMod
	Type      Type
}

// II.23.2.7 CustomMod
type CustomMod struct {
	ModType uint32 `values:"ELEMENT_TYPE_CMOD_OPT / ELEMENT_TYPE_CMOD_REQD"`
	Index   TypeDefOrRefOrSpecEncoded
}

// II.23.2.8 TypeDefOrRefOrSpecEncoded
type TypeDefOrRefOrSpecEncoded TypeDefOrRefEncodedIndex

// II.23.2.10 Param
type Param struct {
	Mod CustomMod
}

// II.23.2.11 RetType
type RetType struct {
}

type Type struct {
}
