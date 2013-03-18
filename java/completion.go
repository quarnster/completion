package java

import (
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/java/descriptors"
	"strings"
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
	ret.Absolute = strings.Replace(String(c.Constant_pool, index), "/", ".", -1)
	ret.Relative = ret.Absolute
	if i := strings.LastIndex(ret.Absolute, "."); i > 0 {
		ret.Relative = ret.Relative[i+1:]
	}
	return
}

func (c *Class) Fields() (fields []content.Field, err error) {
	for _, inf := range c.RawFields {
		var p descriptors.DESCRIPTORS
		p.Parse(String(c.Constant_pool, inf.Descriptor_index))
		outf := descriptors.ToContentField(p.RootNode().Children[0])
		outf.Flags = inf.Access_flags.ToContentFlags()
		outf.Name.Relative = String(c.Constant_pool, inf.Name_index)
		fields = append(fields, outf)
	}
	return
}

func (c *Class) Methods() (methods []content.Method, err error) {
	for _, inf := range c.RawMethods {
		var p descriptors.DESCRIPTORS
		p.Parse(String(c.Constant_pool, inf.Descriptor_index))
		outf := descriptors.ToContentMethod(p.RootNode().Children[0])
		outf.Flags = inf.Access_flags.ToContentFlags()
		outf.Name.Relative = String(c.Constant_pool, inf.Name_index)
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
