package java

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
	"unsafe"
)

/*
http://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html
http://en.wikipedia.org/wiki/Java_class_file
*/
type u1 uint8
type u2 uint16
type u4 uint32
type cp_info int
type field_info int
type method_info int
type attribute_info int

type Constant struct {
}

type Class struct {
	Magic         u4
	Minor_version u2
	Major_version u2
	Constant_pool []Constant
	Access_flags  u2
	This_class    u2
	Super_class   u2
	Interfaces    []u2
	Fields        []field_info
	Methods       []method_info
	Attributes    []attribute_info
}

const (
	magic                       = 0xcafebabe
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ClassDecoder struct {
	reader io.Reader
	err    error
}

func (r *ClassDecoder) Read(size int) ([]byte, error) {
	data := make([]byte, size)
	if n, err := r.reader.Read(data); err != nil {
		return nil, err
	} else if n != len(data) {
		return nil, errors.New("Didn't read the expected number of bytes")
	}
	return data, nil
}

func (dec *ClassDecoder) Decode(v interface{}) error {
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("Expected a pointer not %s", t.Kind()))
	}
	v2 := t.Elem()

	switch v2.Type().Name() {
	case "Constant":
		var tag u1
		if err := dec.Decode(&tag); err != nil {
			return err
		} else {
			switch tag {
			case CONSTANT_Class:
				var name_index u2
				return dec.Decode(&name_index)
			case CONSTANT_Fieldref:
				fallthrough
			case CONSTANT_Methodref:
				fallthrough
			case CONSTANT_InterfaceMethodref:
				var class_index, name_and_type_index u2
				if err := dec.Decode(&class_index); err != nil {
					return err
				} else {
					return dec.Decode(&name_and_type_index)
				}
			default:
				return errors.New(fmt.Sprintf("Unimplemented tag: %d", tag))
			case CONSTANT_String:
				var string_index u2
				return dec.Decode(&string_index)
			case CONSTANT_Integer:
				var v int32
				return dec.Decode(&v)
			case CONSTANT_Float:
				var v float32
				return dec.Decode(&v)
			case CONSTANT_Long:
				var v int64
				return dec.Decode(&v)
			case CONSTANT_Double:
				var v float64
				return dec.Decode(&v)
			case CONSTANT_NameAndType:
				var name_index, type_index u2
				if err := dec.Decode(&name_index); err != nil {
					return err
				} else {
					return dec.Decode(&type_index)
				}
			case CONSTANT_Utf8:
				var length u2
				if err := dec.Decode(&length); err != nil {
					return err
				} else if d, err := dec.Read(int(length)); err != nil {
					return err
				} else {
					data := string(d)
					fmt.Println(data)
				}
			case CONSTANT_MethodHandle:
				type MH struct {
					Reference_kind  u1
					Reference_index u2
				}
				var m MH
				return dec.Decode(&m)
			case CONSTANT_MethodType:
				var desc_index u2
				return dec.Decode(&desc_index)
			case CONSTANT_InvokeDynamic:
				var bootstrap_method_attr_index, name_and_type_index u2
				if err := dec.Decode(&bootstrap_method_attr_index); err != nil {
					return err
				} else {
					return dec.Decode(&name_and_type_index)
				}
			}
		}
	default:
		switch v2.Kind() {
		case reflect.Struct:
			for i := 0; i < v2.NumField(); i++ {
				f := v2.Field(i)
				fmt.Println("Field: ", f.Kind(), f.Type())
				a := f.Addr()
				if err := dec.Decode(a.Interface()); err != nil {
					return err
				}
			}
		case reflect.Slice:
			var count uint16
			if err := dec.Decode(&count); err != nil {
				return err
			} else {
				ic := int(count) - 1
				v2.Set(reflect.MakeSlice(v2.Type(), ic, ic))

				for i := 0; i < v2.Len(); i++ {
					f := v2.Index(i)
					a := f.Addr()
					if err := dec.Decode(a.Interface()); err != nil {
						return err
					}
				}
			}
		case reflect.Int32:
			if d, err := dec.Read(4); err != nil {
				return err
			} else {
				v2.SetInt(int64(binary.BigEndian.Uint32(d)))
			}
		case reflect.Float32:
			if d, err := dec.Read(4); err != nil {
				return err
			} else {
				i32 := binary.BigEndian.Uint32(d)
				f32 := *(*float32)(unsafe.Pointer(&i32))
				v2.SetFloat(float64(f32))
				fmt.Printf("%f\n", v2.Float())
			}
		case reflect.Uint32:
			if d, err := dec.Read(4); err != nil {
				return err
			} else {
				v2.SetUint(uint64(binary.BigEndian.Uint32(d)))
			}
		case reflect.Uint16:
			if d, err := dec.Read(2); err != nil {
				return err
			} else {
				v2.SetUint(uint64(binary.BigEndian.Uint16(d)))
			}
		case reflect.Uint8:
			if d, err := dec.Read(1); err != nil {
				return err
			} else {
				v2.SetUint(uint64(d[0]))
			}
		default:
			return errors.New(fmt.Sprintf("Unknown kind: %s", v2.Kind()))
		}
	}
	return nil
}

func NewClass(reader io.Reader) (*Class, error) {
	r := ClassDecoder{reader, nil}
	var c Class
	if err := r.Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
