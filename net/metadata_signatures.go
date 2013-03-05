package net

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/quarnster/completion/util"
	"io"
	"reflect"
	"unsafe"
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
	} else if _, err := reader.Read(data); err != nil {
		return nil, err
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

var NotACModError = errors.New("Not a cmod")

type EncUint uint32
type EncInt int32

type SignatureDecoder struct {
	Reader       util.BinaryReader
	metadataUtil *MetadataUtil
}

func NewSignatureDecoder(idx BlobIndex) (*SignatureDecoder, error) {
	if idx.Index() == 0 {
		return nil, errors.New("BlobIndex is 0")
	}
	var (
		ci     = idx.(*ConcreteTableIndex)
		table  = ci.metadataUtil.BlobHeap
		ptr    = table.Ptr + uintptr(ci.index)
		end    = table.Ptr + uintptr(table.Rows)
		length = end - ptr
		data   []byte
	)
	if ptr > end {
		return nil, errors.New(fmt.Sprintf("Indexing beyond the end of the table: %x > %x", ptr, end))
	}

	goArray(unsafe.Pointer(&data), ptr, int(length))
	r := bytes.NewReader(data)
	if l, err := UnsignedDecode(r); err != nil {
		return nil, err
	} else {
		pos, _ := r.Seek(0, 1)
		data = data[pos : uint32(pos)+l]
	}
	return &SignatureDecoder{metadataUtil: ci.metadataUtil, Reader: util.BinaryReader{Reader: bytes.NewReader(data), Endianess: binary.LittleEndian}}, nil
}

func (d *SignatureDecoder) Decode(v interface{}) error {
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("Expected a pointer not %s", t.Kind()))
	}
	v2 := t.Elem()
	switch raw := v.(type) {
	case *MethodDefSig:
		var (
			err error
			pc  EncUint
		)
		if err = d.Decode(&raw.Id); err != nil {
			return err
		}
		if raw.Id&SIGNATURE_GENERIC != 0 {
			if err := d.Decode(&raw.GenParamCount); err != nil {
				return err
			}
		}
		if err := d.Decode(&pc); err != nil {
			return err
		}
		if err := d.Decode(&raw.RetType); err != nil {
			return err
		}
		raw.Params = make([]Param, pc)
		for i := range raw.Params {
			if err = d.Decode(&raw.Params[i]); err != nil {
				return err
			}
		}
		// if raw.Id&SIGNATURE_HASTHIS != 0 {
		// 	raw.Params = raw.Params[1:] // pop "this"
		// }
		// if raw.Id&SIGNATURE_EXPLICITTHIS != 0 {
		// 	raw.Params = raw.Params[1:] // Pop its type
		// }
		return nil
	case *FieldSig:
		var id EncUint
		if err := d.Decode(&id); err != nil {
			return err
		} else if id != SIGNATURE_FIELD {
			return errors.New(fmt.Sprintf("Expected type SIGNATURE_FIELD (%x) not %x", SIGNATURE_FIELD, id))
		}
		// Intentionally not returning here
	case *CustomMod:
		p, _ := d.Reader.Seek(0, 1)
		if err := d.Decode(&raw.ModType); err != nil && err != io.EOF {
			return err
		} else if raw.ModType != ELEMENT_TYPE_CMOD_OPT && raw.ModType != ELEMENT_TYPE_CMOD_REQD {
			d.Reader.Seek(p, 0)
			return NotACModError
		} else {
			return d.Decode(&raw.Index)
		}
	case *[]CustomMod:
		for {
			var cm CustomMod
			if err := d.Decode(&cm); err != nil && err != NotACModError {
				return err
			} else if err == NotACModError {
				//fmt.Println("err:", err)
				break
			} else {
				*raw = append(*raw, cm)
			}
		}
		return nil
	case *TypeDefOrRefEncodedIndex:
		if val, err := UnsignedDecode(d.Reader.Reader); err != nil {
			return err
		} else {
			table := val & 3
			index := val >> 2
			*raw = &ConcreteTableIndex{table: enc_lut[idx_TypeDefOrRef][table], index: index, metadataUtil: d.metadataUtil}
		}
		return nil
	case *Type:
		if val, err := UnsignedDecode(d.Reader.Reader); err != nil {
			return err
		} else {
			if val == ELEMENT_TYPE_BYREF {
				return d.Decode(raw)
			}
			raw.TypeId = val
			switch raw.TypeId {
			case ELEMENT_TYPE_VALUETYPE:
				fallthrough
			case ELEMENT_TYPE_CLASS:
				if err := d.Decode(&raw.Class); err != nil {
					return err
				}
			case ELEMENT_TYPE_SZARRAY:
				raw.Type = &Type{}
				if err := d.Decode(&raw.Mod); err != nil {
					return err
				} else if err := d.Decode(raw.Type); err != nil {
					return err
				}
			case ELEMENT_TYPE_VAR:
				// Generic parameter type
				if _, err := UnsignedDecode(d.Reader.Reader); err != nil {
					return err
				} else {
					// what to do with the read value?
				}
			}

			return nil
		}
	}

	switch v2.Kind() {
	case reflect.Struct:
		for i := 0; i < v2.NumField(); i++ {
			if err := d.Decode(v2.Field(i).Addr().Interface()); err != nil {
				return err
			}
		}
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
	case reflect.Uint8:
		if v, err := d.Reader.Uint8(); err != nil {
			return err
		} else {
			v2.SetUint(uint64(v))
		}
	default:
		return errors.New(fmt.Sprintf("Don't know how to decode type %s, %s", v2.Type().Name(), v2.Kind()))
	}
	return nil
}

type MethodDefSigId uint8

// II.23.2.1 MethodDefSig
type MethodDefSig struct {
	Id            MethodDefSigId
	GenParamCount EncUint
	RetType       RetType
	Params        []Param
}

// II.23.2.4 FieldSig
type FieldSig struct {
	CustomMod []CustomMod
	Type      Type
}

// II.23.2.7 CustomMod
type CustomMod struct {
	ModType uint32
	Index   TypeDefOrRefEncodedIndex
}

// II.23.2.10 Param
type Param struct {
	Mod  []CustomMod
	Type Type
}

// II.23.2.11 RetType
type RetType struct {
	Mod  []CustomMod
	Type Type
}

type Type struct {
	TypeId uint32
	Class  TypeDefOrRefEncodedIndex
	Type   *Type
	Mod    []CustomMod
}

var lut_element_type = map[int]string{
	ELEMENT_TYPE_VOID:        "void",
	ELEMENT_TYPE_BOOLEAN:     "boolean",
	ELEMENT_TYPE_CHAR:        "char",
	ELEMENT_TYPE_I1:          "i1",
	ELEMENT_TYPE_U1:          "byte",
	ELEMENT_TYPE_I2:          "short",
	ELEMENT_TYPE_U2:          "ushort",
	ELEMENT_TYPE_I4:          "int",
	ELEMENT_TYPE_U4:          "uint",
	ELEMENT_TYPE_I8:          "long",
	ELEMENT_TYPE_U8:          "ulong",
	ELEMENT_TYPE_R4:          "r4",
	ELEMENT_TYPE_R8:          "r8",
	ELEMENT_TYPE_STRING:      "string",
	ELEMENT_TYPE_PTR:         "ptr",
	ELEMENT_TYPE_BYREF:       "byref",
	ELEMENT_TYPE_VALUETYPE:   "valuetype",
	ELEMENT_TYPE_CLASS:       "class",
	ELEMENT_TYPE_VAR:         "var",
	ELEMENT_TYPE_ARRAY:       "array",
	ELEMENT_TYPE_GENERICINST: "genericinst",
	ELEMENT_TYPE_TYPEDBYREF:  "typedbyref",
	ELEMENT_TYPE_I:           "i",
	ELEMENT_TYPE_U:           "u",
	ELEMENT_TYPE_FNPTR:       "fnptr",
	ELEMENT_TYPE_OBJECT:      "object",
	ELEMENT_TYPE_SZARRAY:     "szarray",
	ELEMENT_TYPE_MVAR:        "mvar",
	ELEMENT_TYPE_CMOD_REQD:   "cmod_reqd",
	ELEMENT_TYPE_CMOD_OPT:    "cmod_opt",
	ELEMENT_TYPE_INTERNAL:    "internal",
	ELEMENT_TYPE_MODIFIER:    "modifier",
	ELEMENT_TYPE_SENTINEL:    "sentinel",
	ELEMENT_TYPE_PINNED:      "pinned",
}

func (t *Type) Name() string {
	switch t.TypeId {
	case ELEMENT_TYPE_VALUETYPE:
		fallthrough
	case ELEMENT_TYPE_CLASS:
		v, _ := t.Class.Data()
		switch t := v.(type) {
		case *TypeDefRow:
			return string(t.TypeName)
		case *TypeRefRow:
			return string(t.TypeName)
		default:
			fmt.Sprintf("Unhandled type: %s", t)
		}
	case ELEMENT_TYPE_SZARRAY:
		return fmt.Sprintf("%s[]", t.Type)
	}
	return lut_element_type[int(t.TypeId)]
}

func (t *Type) Namespace() string {
	switch t.TypeId {
	case ELEMENT_TYPE_VALUETYPE:
		fallthrough
	case ELEMENT_TYPE_CLASS:
		v, _ := t.Class.Data()
		switch t := v.(type) {
		case *TypeDefRow:
			return string(t.TypeNamespace)
		case *TypeRefRow:
			return string(t.TypeNamespace)
		default:
			fmt.Sprintf("Unhandled type: %s", t)
		}
	}
	return ""
}

func (t *Type) String() string {
	return t.Name()
}
