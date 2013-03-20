package java

import (
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/java/descriptors"
)

func (af AccessFlags) ToContentFlags() (ret content.Flags) {
	if af&ACC_PUBLIC != 0 {
		ret |= content.FLAG_ACC_PUBLIC
	}
	if af&ACC_PRIVATE != 0 {
		ret |= content.FLAG_ACC_PRIVATE
	}
	if af&ACC_PROTECTED != 0 {
		ret |= content.FLAG_ACC_PROTECTED
	}
	if af&ACC_STATIC != 0 {
		ret |= content.FLAG_STATIC
	}
	if af&ACC_FINAL != 0 {
		ret |= content.FLAG_FINAL
	}
	return
}

func (c *Class) ToContentFQN(index u2) (ret content.FullyQualifiedName) {
	return descriptors.ToContentFQN(c.Constant_pool.Lut(index).String())
}

func (c *Class) Fields() (fields []content.Field, err error) {
	cn := c.Constant_pool.Lut(c.This_class).String()
	for i, inf := range c.RawFields {
		var p descriptors.DESCRIPTORS
		desc := c.Constant_pool.Lut(inf.Descriptor_index).String()
		p.Parse(desc)
		outf := descriptors.ToContentField(p.RootNode().Children[0])
		outf.Flags = inf.Access_flags.ToContentFlags()
		outf.Name.Relative = c.Constant_pool.Lut(inf.Name_index).String()
		outf.Name.Absolute = fmt.Sprintf("java://field/%s;%d", cn, i)
		fields = append(fields, outf)
	}
	return
}

func (c *Class) Methods() (methods []content.Method, err error) {
	cn := c.Constant_pool.Lut(c.This_class).String()
	for i, inf := range c.RawMethods {
		var p descriptors.DESCRIPTORS
		desc := c.Constant_pool.Lut(inf.Descriptor_index).String()
		p.Parse(desc)
		outf := descriptors.ToContentMethod(p.RootNode().Children[0])
		outf.Flags = inf.Access_flags.ToContentFlags()
		outf.Name.Relative = c.Constant_pool.Lut(inf.Name_index).String()
		outf.Name.Absolute = fmt.Sprintf("java://method/%s;%d", cn, i)
		methods = append(methods, outf)
	}
	return
}

func (c *Class) Implements() (ret []content.Type, err error) {
	for _, i := range c.Interfaces {
		ret = append(ret, content.Type{Name: c.ToContentFQN(i)})
	}
	return
}

func (c *Class) Extends() (ret []content.Type, err error) {
	ret = append(ret, content.Type{Name: c.ToContentFQN(c.Super_class)})
	return
}

func (c *Class) ToContentType() (ret content.Type, err error) {
	ret.Name = c.ToContentFQN(c.This_class)
	ret.Flags = c.Access_flags.ToContentFlags()
	if f, err := c.Fields(); err != nil {
		return ret, err
	} else {
		ret.Fields = f
	}
	if m, err := c.Methods(); err != nil {
		return ret, err
	} else {
		ret.Methods = m
	}
	if i, err := c.Implements(); err != nil {
		return ret, err
	} else {
		ret.Implements = i
	}
	if e, err := c.Extends(); err != nil {
		return ret, err
	} else {
		ret.Extends = e
	}

	return
}
