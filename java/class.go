// TODO(d) LocalVariableTable and LocalVariableTypeTable are unaccounted for causing tests to fail
package java

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/quarnster/util/encoding/binary"
	"io"
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
		cp    *ConstantPool
		Tag   u1
		Index [2]u2
		Value string
	}
	AccessFlags u2

	member_info struct {
		Access_flags     AccessFlags
		Name_index       u2
		Descriptor_index u2
		Attributes       []attribute_info `length:"uint16"`
	}
	attribute_info struct {
		Attribute_name_index u2
		Info                 []byte `length:"uint32"`
	}

	ConstantPool struct {
		constants []Constant
	}

	Class struct {
		Magic         u4
		Minor_version u2
		Major_version u2
		Constant_pool ConstantPool
		Access_flags  AccessFlags
		This_class    u2
		Super_class   u2
		Interfaces    []u2             `length:"uint16"`
		RawFields     []member_info    `length:"uint16"`
		RawMethods    []member_info    `length:"uint16"`
		Attributes    []attribute_info `length:"uint16"`
	}

	// http://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.3
	exception_table struct {
		Start_pc   u2
		End_pc     u2
		Handler_pc u2
		Catch_type u2
	}
	Code_attribute struct {
		Max_stack       u2
		Max_locals      u2
		Code            []u1              `length:"uint32"`
		Exception_table []exception_table `length:"uint16"`
		Attributes      []attribute_info  `length:"uint16"`
	}
)

func (cp *ConstantPool) readConstant(c *Constant, br *binary.BinaryReader) error {
	err := br.ReadInterface(&c.Tag)
	if err != nil {
		return err
	}
	switch c.Tag {
	case CONSTANT_String:
		fallthrough
	case CONSTANT_MethodType:
		fallthrough
	case CONSTANT_Class:
		return br.ReadInterface(&c.Index[0])
	case CONSTANT_Fieldref:
		fallthrough
	case CONSTANT_Methodref:
		fallthrough
	case CONSTANT_NameAndType:
		fallthrough
	case CONSTANT_InvokeDynamic:
		fallthrough
	case CONSTANT_InterfaceMethodref:
		if err := br.ReadInterface(&c.Index[0]); err != nil {
			return err
		}
		return br.ReadInterface(&c.Index[1])
	case CONSTANT_Integer:
		var v int32
		return br.ReadInterface(&v)
	case CONSTANT_Float:
		var v float32
		return br.ReadInterface(&v)
	case CONSTANT_Long:
		var v int64
		return br.ReadInterface(&v)
	case CONSTANT_Double:
		var v float64
		return br.ReadInterface(&v)
	case CONSTANT_Utf8:
		var length u2
		if err := br.ReadInterface(&length); err != nil {
			return err
		} else if d, err := br.Read(int(length)); err != nil {
			return err
		} else {
			c.Value = string(d)
		}
	case CONSTANT_MethodHandle:
		var ref_kind u1
		if err := br.ReadInterface(&ref_kind); err != nil {
			return err
		}
		c.Index[0] = u2(ref_kind)
		return br.ReadInterface(&c.Index[1])
	default:
		return errors.New(fmt.Sprintf("Unimplemented tag: %d", c.Tag))
	}
	return nil
}

func (c *ConstantPool) Read(br *binary.BinaryReader) error {
	var count uint16
	if err := br.ReadInterface(&count); err != nil {
		return err
	}
	ic := int(count)
	ic--
	c.constants = make([]Constant, ic, ic)

	for i := 0; i < len(c.constants); i++ {
		cc := &c.constants[i]
		cc.cp = c
		if err := c.readConstant(cc, br); err != nil {
			return err
		}

		if cc.Tag == CONSTANT_Double || cc.Tag == CONSTANT_Long {
			// All 8-byte constants take up two entries in the constant_pool table of the class file.
			i++
		}
	}
	return nil
}

func (c *ConstantPool) Lut(idx u2) *Constant {
	if idx > 0 && int(idx) <= len(c.constants) {
		return &c.constants[idx-1]
	}
	return nil
}

func (c *Constant) String() string {
	switch c.Tag {
	case CONSTANT_Utf8:
		return c.Value
	case CONSTANT_String:
		fallthrough
	case CONSTANT_MethodType:
		fallthrough
	case CONSTANT_Class:
		return fmt.Sprintf("%s", c.cp.Lut(c.Index[0]))
	}
	return fmt.Sprintf("Non-stringed tag: %d", c.Tag)
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

func (a *attribute_info) String(c *ConstantPool) (ret string) {
	ret = c.Lut(a.Attribute_name_index).String()
	switch n := c.Lut(a.Attribute_name_index).String(); n {
	case "Signature", "SourceFile":
		ret += "="
		br := binary.BinaryReader{bytes.NewReader(a.Info), binary.BigEndian}
		if i16, err := br.Uint16(); err != nil {
			ret += err.Error()
		} else {
			ret += c.Lut(u2(i16)).String()
		}
	case "Code":
		ret += " ("
		var cl Code_attribute
		br := binary.BinaryReader{bytes.NewReader(a.Info), binary.BigEndian}
		if err := br.ReadInterface(&cl); err != nil {
			ret += err.Error()
		} else {
			for _, a2 := range cl.Attributes {
				ret += fmt.Sprintf(" %s", c.Lut(a2.Attribute_name_index))
			}
		}
		ret += " )"
	}
	return ret
}

func (mi *member_info) String(c *ConstantPool) string {
	ret := mi.Access_flags.String()
	ret += fmt.Sprintf("%s %s", c.Lut(mi.Name_index), c.Lut(mi.Descriptor_index))
	for i, a := range mi.Attributes {
		if i == 0 {
			ret += "\n\t\t"
		}
		ret += fmt.Sprintf(" %s", a.String(c))
	}
	return ret
}

func (c *Class) String() (ret string) {
	ret += fmt.Sprintf("%sclass %s\n", c.Access_flags, c.Constant_pool.Lut(c.This_class))
	ret += fmt.Sprintln("extends", c.Constant_pool.Lut(c.Super_class))
	ret += fmt.Sprintln("implements")
	for _, i := range c.Interfaces {
		ret += fmt.Sprintf("\t%s\n", c.Constant_pool.Lut(i))
	}
	ret += fmt.Sprintln("attributes")
	for _, a := range c.Attributes {
		ret += fmt.Sprintf("\t%s\n", a.String(&c.Constant_pool))
	}

	ret += fmt.Sprintln("Fields")
	for _, f := range c.RawFields {
		ret += fmt.Sprintf("\t%s\n", f.String(&c.Constant_pool))
	}
	ret += fmt.Sprintln("Methods")
	for _, m := range c.RawMethods {
		ret += fmt.Sprintf("\t%s\n", m.String(&c.Constant_pool))
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
	// public Declared may be accessed from outside its package.
	ACC_PUBLIC = 0x0001
	// private Declared usable only within the defining class.
	ACC_PRIVATE = 0x0002
	// protected Declared may be accessed within subclasses.
	ACC_PROTECTED = 0x0004
	// Declared static.
	ACC_STATIC = 0x0008
	// final Declared never directly assigned to after object construction (JLS §17.5).
	ACC_FINAL = 0x0010
	// volatile Declared cannot be cached.
	ACC_VOLATILE = 0x0040
	// transient Declared not written or read by a persistent object manager.
	ACC_TRANSIENT = 0x0080
	// synthetic Declared not present in the source code.
	ACC_SYNTHETIC = 0x1000
	// Declared as an element of an enum.
	ACC_ENUM = 0x4000
)

func NewClass(reader io.ReadSeeker) (*Class, error) {
	r := binary.BinaryReader{reader, binary.BigEndian}
	var c Class
	if err := r.ReadInterface(&c); err != nil {
		return nil, err
	} else if c.Magic != magic {
		return nil, errors.New(fmt.Sprintf("Magic isn't what's expected: %x", c.Magic))
	}
	return &c, nil
}
