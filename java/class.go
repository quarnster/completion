package java

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/quarnster/completion/util"
	"io"
	"reflect"
	"unsafe"
)

/*
http://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html
http://en.wikipedia.org/wiki/Java_class_file
*/
type (
	u1 uint8
	u2 uint16
	u4 uint32

	Constant struct {
		tag   u1
		index [2]u2
		value string
	}
	AccessFlags u2

	member_info struct {
		Access_flags     AccessFlags
		Name_index       u2
		Descriptor_index u2
		Attributes       []attribute_info
	}
	attribute_info struct {
		Attribute_name_index u2
		Info                 []byte
	}

	Class struct {
		Magic         u4
		Minor_version u2
		Major_version u2
		Constant_pool []Constant
		Access_flags  AccessFlags
		This_class    u2
		Super_class   u2
		Interfaces    []u2
		RawFields     []member_info
		RawMethods    []member_info
		Attributes    []attribute_info
	}

	ClassDecoder struct {
		reader util.BinaryReader
		err    error
	}
)

func String(lut []Constant, idx u2) string {
	if idx > 0 && int(idx) <= len(lut) {
		return lut[idx-1].String(lut)
	}
	return fmt.Sprintf("Invalid index: %d", idx)
}

func (c Constant) String(lut []Constant) string {
	switch c.tag {
	case CONSTANT_Utf8:
		return c.value
	case CONSTANT_String:
		fallthrough
	case CONSTANT_MethodType:
		fallthrough
	case CONSTANT_Class:
		return String(lut, c.index[0])
	}
	return fmt.Sprintf("Non-stringed tag: %d", c.tag)
}

func (a AccessFlags) String() (ret string) {
	if a&ACC_PUBLIC != 0 {
		ret += "public "
	}
	if a&ACC_PRIVATE != 0 {
		ret += "private "
	}
	if a&ACC_PROTECTED != 0 {
		ret += "protected "
	}
	if a&ACC_STATIC != 0 {
		ret += "static "
	}
	if a&ACC_FINAL != 0 {
		ret += "final "
	}
	if a&ACC_VOLATILE != 0 {
		ret += "volatile "
	}
	if a&ACC_TRANSIENT != 0 {
		ret += "transient "
	}
	if a&ACC_SYNTHETIC != 0 {
		ret += "synthetic "
	}
	if a&ACC_ENUM != 0 {
		ret += "enum "
	}
	return ret
}

func (mi *member_info) String(c []Constant) string {
	ret := mi.Access_flags.String()
	ret += String(c, mi.Name_index)
	ret += " " + String(c, mi.Descriptor_index)
	return ret
}

func (c *Class) String() (ret string) {
	ret += fmt.Sprintf("%sclass %s\n", c.Access_flags, String(c.Constant_pool, c.This_class))
	ret += fmt.Sprintln("extends", String(c.Constant_pool, c.Super_class))
	ret += fmt.Sprintln("implements")
	for _, i := range c.Interfaces {
		ret += fmt.Sprintf("\t%s\n", String(c.Constant_pool, i))
	}
	ret += fmt.Sprintln("Fields")
	for _, f := range c.RawFields {
		ret += fmt.Sprintf("\t%s\n", f.String(c.Constant_pool))
	}
	ret += fmt.Sprintln("Methods")
	for _, m := range c.RawMethods {
		ret += fmt.Sprintf("\t%s\n", m.String(c.Constant_pool))
	}
	return ret
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

const (
	// Declared public; may be accessed from outside its package.
	ACC_PUBLIC = 0x0001
	// Declared private; usable only within the defining class.
	ACC_PRIVATE = 0x0002
	// Declared protected; may be accessed within subclasses.
	ACC_PROTECTED = 0x0004
	// Declared static.
	ACC_STATIC = 0x0008
	// Declared final; never directly assigned to after object construction (JLS §17.5).
	ACC_FINAL = 0x0010
	// Declared volatile; cannot be cached.
	ACC_VOLATILE = 0x0040
	// Declared transient; not written or read by a persistent object manager.
	ACC_TRANSIENT = 0x0080
	// Declared synthetic; not present in the source code.
	ACC_SYNTHETIC = 0x1000
	// Declared as an element of an enum.
	ACC_ENUM = 0x4000
)

func (dec *ClassDecoder) Decode(v interface{}) error {
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("Expected a pointer not %s", t.Kind()))
	}
	v2 := t.Elem()

	switch v2.Type().Name() {
	case "Constant":
		c := (*Constant)(unsafe.Pointer(t.Pointer()))
		if err := dec.Decode(&c.tag); err != nil {
			return err
		} else {
			switch c.tag {
			case CONSTANT_String:
				fallthrough
			case CONSTANT_MethodType:
				fallthrough
			case CONSTANT_Class:
				return dec.Decode(&c.index[0])
			case CONSTANT_Fieldref:
				fallthrough
			case CONSTANT_Methodref:
				fallthrough
			case CONSTANT_NameAndType:
				fallthrough
			case CONSTANT_InvokeDynamic:
				fallthrough
			case CONSTANT_InterfaceMethodref:
				if err := dec.Decode(&c.index[0]); err != nil {
					return err
				} else {
					return dec.Decode(&c.index[1])
				}
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
			case CONSTANT_Utf8:
				var length u2
				if err := dec.Decode(&length); err != nil {
					return err
				} else if d, err := dec.reader.Read(int(length)); err != nil {
					return err
				} else {
					c.value = string(d)
				}
			case CONSTANT_MethodHandle:
				var ref_kind u1
				if err := dec.Decode(&ref_kind); err != nil {
					return err
				} else {
					c.index[0] = u2(ref_kind)
					return dec.Decode(&c.index[1])
				}
			default:
				return errors.New(fmt.Sprintf("Unimplemented tag: %d", c.tag))
			}
		}
	default:
		switch v2.Kind() {
		case reflect.Struct:
			for i := 0; i < v2.NumField(); i++ {
				f := v2.Field(i)
				if v2.Type().Name() == "attribute_info" && f.Kind() == reflect.Slice {
					var length u4
					if err := dec.Decode(&length); err != nil {
						return err
					} else {
						if d, err := dec.reader.Read(int(length)); err != nil {
							return err
						} else {
							f.SetBytes(d)
						}
					}
				} else {
					a := f.Addr()
					if err := dec.Decode(a.Interface()); err != nil {
						return err
					}
				}
			}
		case reflect.Slice:
			var count uint16
			if err := dec.Decode(&count); err != nil {
				return err
			} else {
				ic := int(count)
				isConstantPool := v2.Type().String() == "[]java.Constant"
				if isConstantPool {
					ic--
				}
				v2.Set(reflect.MakeSlice(v2.Type(), ic, ic))

				for i := 0; i < v2.Len(); i++ {
					f := v2.Index(i)
					a := f.Addr()

					if err := dec.Decode(a.Interface()); err != nil {
						return err
					}
					if isConstantPool {
						c := (*Constant)(unsafe.Pointer(a.Pointer()))
						if c.tag == CONSTANT_Double || c.tag == CONSTANT_Long {
							// All 8-byte constants take up two entries in the constant_pool table of the class file.
							i++
						}
					}
				}
			}
		default:
			return dec.reader.ReadInterface(v)
		}
	}
	return nil
}

func NewClass(reader io.ReadSeeker) (*Class, error) {
	r := ClassDecoder{util.BinaryReader{reader, binary.BigEndian}, nil}
	var c Class
	if err := r.Decode(&c); err != nil {
		return nil, err
	} else if c.Magic != magic {
		return nil, errors.New(fmt.Sprintf("Magic isn't what's expected: %x", c.Magic))
	}
	return &c, nil
}
