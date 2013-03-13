package dwarf

import (
	"bytes"
	"code.google.com/p/log4go"
	"compress/bzip2"
	"debug/dwarf"
	"debug/macho"
	"errors"
	"fmt"
	"github.com/quarnster/completion/content"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DWARFHelper struct {
	df *dwarf.Data
}

func NewDWARFHelper(filename string) (*DWARFHelper, error) {
	// TODO: detect file format..
	var f *macho.File
	var err error
	if filepath.Ext(filename) == ".bz2" {
		if fs, err2 := os.Open(filename); err2 != nil {
			return nil, err2
		} else {
			defer fs.Close()
			if data, err2 := ioutil.ReadAll(bzip2.NewReader(fs)); err2 != nil {
				return nil, err2
			} else {
				f, err = macho.NewFile(bytes.NewReader(data))
			}
		}
	} else {
		f, err = macho.Open(filename)
	}
	if err != nil {
		return nil, err
	} else {
		defer f.Close()
		if df, err := f.DWARF(); err != nil {
			for i := range f.Sections {
				log4go.Debug(f.Sections[i])
			}
			return nil, err
		} else {
			var ret DWARFHelper
			ret.df = df
			return &ret, nil
		}
	}
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

func (d *DWARFHelper) Load() (content.CompletionResult, error) {
	var cmp content.CompletionResult
	r := d.df.Reader()
	for {
		if e, err := r.Next(); err != nil {
			return cmp, err
			break
		} else if e == nil {
			break
		} else {
			switch e.Tag {
			case dwarf.TagMember:
				var f content.Field
				if v, ok := e.Val(dwarf.AttrName).(string); ok {
					f.Name.Relative = v
				}
				if v, ok := e.Val(dwarf.AttrType).(dwarf.Offset); ok {
					if t, err := d.GetType(v); err != nil {
						return cmp, err
					} else {
						f.Type = t
					}
				}
				f.Flags = d.Flags(e)
				cmp.Fields = append(cmp.Fields, f)
			case dwarf.TagSubprogram:
				var m content.Method
				if v, ok := e.Val(dwarf.AttrName).(string); ok {
					m.Name.Relative = v
				}
				if v, ok := e.Val(dwarf.AttrType).(dwarf.Offset); ok {
					if t, err := d.GetType(v); err != nil {
						return cmp, err
					} else {
						m.Returns = append(m.Returns, content.Variable{Type: t})
					}
				} else {
					m.Returns = append(m.Returns, content.Variable{Type: content.Type{Name: content.FullyQualifiedName{Relative: "void"}}})
				}
				m.Flags = d.Flags(e)
				if e.Children {
					for {
						if e, err := r.Next(); err != nil {
							return cmp, err
						} else if e == nil || e.Tag == 0 {
							break
						} else if e.Tag == dwarf.TagFormalParameter {
							var p content.Variable
							if v, ok := e.Val(dwarf.AttrType).(dwarf.Offset); ok {
								if t, err := d.GetType(v); err != nil {
									return cmp, err
								} else {
									p.Type = t
								}
							}
							if v, ok := e.Val(dwarf.AttrName).(string); ok {
								p.Name.Relative = v
							}
							m.Parameters = append(m.Parameters, p)
						}
					}
				}
				cmp.Methods = append(cmp.Methods, m)
			case dwarf.TagClassType, dwarf.TagTypedef, dwarf.TagStructType:
				if t, err := d.GetType(e.Offset); err != nil {
					return cmp, err
				} else {
					cmp.Types = append(cmp.Types, t)
				}
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
	if e, err := r.Next(); err != nil {
		return t, err
	} else {
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
				} else if e.Tag == dwarf.TagFormalParameter {
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
			}
			t.Methods = append(t.Methods, m)
		}

		return t, nil
	}
}
