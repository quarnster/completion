package dwarf

import (
	"code.google.com/p/log4go"
	"debug/dwarf"
	"debug/macho"
	"errors"
	"fmt"
	"github.com/quarnster/completion/content"
	"io"
)

type DWARFHelper struct {
	df *dwarf.Data
	id string
}

func NewDWARFHelper(id string, r io.ReaderAt) (*DWARFHelper, error) {
	// TODO: detect file format..
	f, err := macho.NewFile(r)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	df, err := f.DWARF()
	if err != nil {
		for i := range f.Sections {
			log4go.Debug(f.Sections[i])
		}
		return nil, err
	}
	var ret DWARFHelper
	ret.df = df
	ret.id = id
	return &ret, nil
}
func (d *DWARFHelper) Flags(e *dwarf.Entry) (ret content.Flags) {
	if v, ok := e.Val(dwarf.AttrAccessibility).(int64); ok {
		switch v {
		case 1:
			ret |= content.FLAG_ACC_PUBLIC
		case 2:
			ret |= content.FLAG_ACC_PROTECTED
		case 3:
			ret |= content.FLAG_ACC_PRIVATE
		}
	}
	return
}

func (d *DWARFHelper) Complete(id content.FullyQualifiedName) (content.CompletionResult, error) {
	var cmp content.CompletionResult
	r := d.df.Reader()
	for {
		e, err := r.Next()
		if err != nil {
			return cmp, err
		} else if e == nil {
			break
		}
		switch e.Tag {
		case dwarf.TagCompileUnit:
			continue
		case dwarf.TagClassType, dwarf.TagTypedef, dwarf.TagStructType:
			t, err := d.GetType(e.Offset)
			if err != nil {
				return cmp, err
			}
			if t.Name != id || !e.Children {
				r.SkipChildren()
				continue
			}
			r2 := d.df.Reader()
			r2.Seek(e.Offset)
			for {
				e, err := r.Next()
				if err != nil {
					return cmp, err
				} else if e == nil || e.Tag == 0 {
					break
				}
				switch e.Tag {
				case dwarf.TagMember:
					if f, err := d.toContentField(e); err != nil {
						return cmp, err
					} else {
						f.Name.Absolute = fmt.Sprintf("dwarf://field/%s;%d", d.id, e.Offset)
						cmp.Fields = append(cmp.Fields, f)
					}
				case dwarf.TagSubprogram:
					if m, err := d.toContentMethod(e); err != nil {
						return cmp, err
					} else {
						m.Name.Absolute = fmt.Sprintf("dwarf://method/%s;%d", d.id, e.Offset)
						cmp.Methods = append(cmp.Methods, m)
					}
					r.SkipChildren()
				}
			}
			return cmp, nil
		default:
			r.SkipChildren()
		}
	}
	return cmp, errors.New(fmt.Sprintf("Unable to find type: %s", id))
}

func (d *DWARFHelper) toContentField(e *dwarf.Entry) (content.Field, error) {
	var f content.Field
	if v, ok := e.Val(dwarf.AttrName).(string); ok {
		f.Name.Relative = v
	}
	if v, ok := e.Val(dwarf.AttrType).(dwarf.Offset); ok {
		if t, err := d.GetType(v); err != nil {
			return f, err
		} else {
			f.Type = t
		}
	}
	f.Flags = d.Flags(e)
	return f, nil
}

var i = 0

func (d *DWARFHelper) toContentMethod(e *dwarf.Entry) (m content.Method, err error) {
	if v, ok := e.Val(dwarf.AttrName).(string); ok {
		m.Name.Relative = v
	}
	if v, ok := e.Val(dwarf.AttrType).(dwarf.Offset); ok {
		if t, err := d.GetType(v); err != nil {
			return m, err
		} else {
			m.Returns = append(m.Returns, content.Variable{Type: t})
		}
	} else {
		m.Returns = append(m.Returns, content.Variable{Type: content.Type{Name: content.FullyQualifiedName{Relative: "void"}}})
	}
	m.Flags = d.Flags(e)
	if !e.Children {
		return
	}
	r := d.df.Reader()
	r.Seek(e.Offset)
	for {
		e, err := r.Next()
		if err != nil {
			return m, err
		} else if e == nil || e.Tag == 0 {
			break
		}
		if e.Tag != dwarf.TagFormalParameter {
			continue
		}
		var p content.Variable
		if v, ok := e.Val(dwarf.AttrType).(dwarf.Offset); ok {
			if t, err := d.GetType(v); err != nil {
				return m, err
			} else {
				p.Type = t
			}
		}
		if v, ok := e.Val(dwarf.AttrName).(string); ok {
			p.Name.Relative = v
		}
		if v, ok := e.Val(dwarf.AttrArtificial).(bool); ok && v {
			if p.Type.Flags&content.FLAG_TYPE_MASK != content.FLAG_TYPE_POINTER {
				m.Parameters = append(m.Parameters, p)
				continue
			}
			// C++ "this" pointer
			t := p.Type.Specialization[0]
			if t.Flags&content.FLAG_CONST != 0 {
				m.Flags |= content.FLAG_CONST
			}
			//m.Name.Absolute = t.Name.Relative + "::" + m.Name.Relative
		} else {
			m.Parameters = append(m.Parameters, p)
		}
	}
	return
}

func (d *DWARFHelper) Load() (content.CompletionResult, error) {
	var cmp content.CompletionResult
	r := d.df.Reader()
	for {
		e, err := r.Next()
		if err != nil {
			return cmp, err
			break
		} else if e == nil {
			break
		}
		switch e.Tag {
		case dwarf.TagMember:
			if f, err := d.toContentField(e); err != nil {
				return cmp, err
			} else {
				cmp.Fields = append(cmp.Fields, f)
			}
		case dwarf.TagSubprogram:
			if m, err := d.toContentMethod(e); err != nil {
				return cmp, err
			} else {
				cmp.Methods = append(cmp.Methods, m)
			}
		case dwarf.TagClassType, dwarf.TagTypedef, dwarf.TagStructType:
			if t, err := d.GetType(e.Offset); err != nil {
				return cmp, err
			} else {
				cmp.Types = append(cmp.Types, t)
			}
		}
	}
	return cmp, nil
}

func (d *DWARFHelper) GetType(off dwarf.Offset) (content.Type, error) {
	var t content.Type
	r := d.df.Reader()
	r.Seek(off)
	recurse := false
	e, err := r.Next()
	if err != nil {
		return t, err
	}
	switch e.Tag {
	case dwarf.TagVolatileType, dwarf.TagReferenceType, dwarf.TagRestrictType, dwarf.TagConstType, dwarf.TagSubroutineType:
		recurse = true
	case dwarf.TagPointerType:
		t.Flags |= content.FLAG_TYPE_POINTER
		recurse = true
	case dwarf.TagArrayType:
		recurse = true
		t.Flags |= content.FLAG_TYPE_ARRAY
	case dwarf.TagClassType, dwarf.TagTypedef, dwarf.TagBaseType, dwarf.TagEnumerationType, dwarf.TagStructType:
		if v, ok := e.Val(dwarf.AttrName).(string); ok {
			t.Name.Relative = v
		}
	case dwarf.TagUnionType:
		t.Name.Relative = "union"
	default:
		return t, errors.New(fmt.Sprintf("Don't know how to handle %+v", e))
	}
	if recurse {
		if v, ok := e.Val(dwarf.AttrType).(dwarf.Offset); ok {
			if t2, err := d.GetType(v); err != nil {
				return t, err
			} else {
				t.Specialization = append(t.Specialization, t2)
			}
		}
	}
	switch e.Tag {
	case dwarf.TagVolatileType, dwarf.TagReferenceType, dwarf.TagRestrictType, dwarf.TagConstType:
		if len(t.Specialization) == 0 {
			t.Name.Relative = "void"
		} else {
			t = t.Specialization[0]
		}
	}

	switch e.Tag {
	case dwarf.TagPointerType:
		if len(t.Specialization) == 0 {
			t.Specialization = append(t.Specialization, content.Type{Name: content.FullyQualifiedName{Relative: "void"}})
		}
	case dwarf.TagVolatileType:
		t.Flags |= content.FLAG_VOLATILE
	case dwarf.TagReferenceType:
		t.Flags |= content.FLAG_REFERENCE
	case dwarf.TagRestrictType:
		t.Flags |= content.FLAG_RESTRICT
	case dwarf.TagConstType:
		t.Flags |= content.FLAG_CONST
	case dwarf.TagSubroutineType:
		var m content.Method
		if len(t.Specialization) > 0 {
			m.Returns = append(m.Returns, content.Variable{Type: t.Specialization[0]})
		} else {
			m.Returns = append(m.Returns, content.Variable{Type: content.Type{Name: content.FullyQualifiedName{Relative: "void"}}})
		}
		t.Specialization = t.Specialization[0:0]
		t.Flags = content.FLAG_TYPE_METHOD
		for {
			if e, err := r.Next(); err != nil {
				return t, err
			} else if e == nil || e.Tag == 0 {
				break
			}
			if e.Tag != dwarf.TagFormalParameter {
				continue
			}
			var p content.Variable
			if v, ok := e.Val(dwarf.AttrType).(dwarf.Offset); ok {
				if t2, err := d.GetType(v); err != nil {
					return t, err
				} else {
					p.Type = t2
				}
			}
			if v, ok := e.Val(dwarf.AttrName).(string); ok {
				p.Name.Relative = v
			}
			m.Parameters = append(m.Parameters, p)
		}
		t.Methods = append(t.Methods, m)
	}

	return t, nil
}
