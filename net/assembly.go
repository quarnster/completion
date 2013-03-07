package net

// references:
// http://msdn.microsoft.com/en-us/library/windows/hardware/gg463119.aspx
// http://www.codeproject.com/Articles/12585/The-NET-File-Format
// http://msdn.microsoft.com/en-us/library/ms809762.aspx
// http://www.ecma-international.org/publications/standards/Ecma-335.htm

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util"
	"io"
	"sort"
)

var (
	ErrInterface   = errors.New("TypeDef is an interface, not a class")
	ErrNotAssembly = errors.New("This does not appear to be a .net assembly")
)

type Assembly struct {
	MetadataUtil
}

type AbstractType interface {
	Name() string
	Namespace() string
}

func AbsoluteName(t AbstractType) string {
	n := t.Name()
	if ns := t.Namespace(); ns != "" {
		return ns + "." + n
	}
	return n
}

func ToContentType(t AbstractType) (t2 content.Type) {
	t2.Name.Relative = t.Name()
	if ns := t.Namespace(); ns != "" {
		t2.Name.Absolute = ns + "." + t2.Name.Relative
	}
	return
}

func (a *Assembly) ToContentType(index TypeDefIndex, t *Type) (t2 content.Type) {
	switch t.TypeId {
	case ELEMENT_TYPE_GENERICINST:
		t2 = ToContentType(t.Type)
		for i := range t.Instance {
			t2.Specialization = append(t2.Specialization, a.ToContentType(index, t.Instance[i]))
		}
	case ELEMENT_TYPE_VAR, ELEMENT_TYPE_MVAR:
		idx2 := a.Search(index, id_GenericParam, func(idx TableIndex) bool {
			if raw, err := idx.Data(); err == nil {
				gr := raw.(*GenericParamRow)
				return gr.Owner.Table() >= index.Table() && gr.Owner.Index() >= index.Index()
			}
			return false
		})
		if idx2.Table() != id_nullTable {
			ci := idx2.(*ConcreteTableIndex)
			ci.index += t.GenericNumber
			if raw, err := ci.Data(); err == nil {
				gr := raw.(*GenericParamRow)
				t2.Name.Relative = string(gr.Name)
			}
		}
	default:
		return ToContentType(t)
	}

	return
}

func (t *TypeDefRow) Name() string {
	return string(t.TypeName)
}

func (t *TypeDefRow) Namespace() string {
	return string(t.TypeNamespace)
}

func (t *TypeRefRow) Name() string {
	return string(t.TypeName)
}

func (t *TypeRefRow) Namespace() string {
	return string(t.TypeNamespace)
}

func (t *TypeSpecRow) Name() string {
	return string("spec")
}

func (t *TypeSpecRow) Namespace() string {
	return string("unknown")
}

func (a *Assembly) Name() string {
	ci := ConcreteTableIndex{&a.MetadataUtil, 1, id_Module}
	if raw, err := ci.Data(); err == nil {
		mr := raw.(*ModuleRow)
		return string(mr.Name)
	}
	return ""
}

func (a *Assembly) ListRange(index uint32, table, memberTable int, getindex func(interface{}) uint32) (startRow, endRow uint32) {
	var (
		idx      = ConcreteTableIndex{&a.MetadataUtil, index, table}
		tableEnd = a.Tables[memberTable].Rows + 1
	)
	if i, err := idx.Data(); err != nil {
		return 0, 0
	} else {
		startRow = getindex(i)
	}
	idx.index++
	if i, err := idx.Data(); err == nil {
		endRow = getindex(i)
	} else {
		endRow = tableEnd
	}
	if endRow < startRow {
		endRow = tableEnd
	}
	return
}

func (a *Assembly) Types() (types []content.Type, err error) {
	var (
		idx = ConcreteTableIndex{&a.MetadataUtil, 0, id_TypeDef}
	)
	for i := uint32(0); i < a.Tables[id_TypeDef].Rows; i++ {
		idx.index = 1 + i
		if rawtype, err := idx.Data(); err != nil {
			return nil, err
		} else {
			var (
				tr = rawtype.(*TypeDefRow)
			)
			types = append(types, ToContentType(tr))
		}
	}
	return
}

func (a *Assembly) Fields(index TypeDefIndex) (fields []content.Field, err error) {
	var (
		startRow, endRow = a.ListRange(index.Index(), id_TypeDef, id_Field, func(i interface{}) uint32 { return i.(*TypeDefRow).FieldList.Index() })
		idx              = ConcreteTableIndex{&a.MetadataUtil, startRow, id_Field}
	)
	for i := startRow; i < endRow; i++ {
		idx.index = i
		if rawfield, err := idx.Data(); err != nil {
			return nil, err
		} else {
			var (
				field = rawfield.(*FieldRow)
				f     content.Field
				dec   *SignatureDecoder
				sig   FieldSig
			)
			f.Name.Relative = string(field.Name)
			if dec, err = NewSignatureDecoder(field.Signature); err != nil {
				return nil, err
			} else if err = dec.Decode(&sig); err != nil {
				return nil, err
			} else {
				f.Type = ToContentType(&sig.Type)
			}
			if field.Flags&FieldAttributes_Static != 0 {
				f.Flags |= content.FLAG_STATIC
			}
			if field.Flags&FieldAttributes_Public != 0 {
				f.Flags |= content.FLAG_ACC_PUBLIC
			} else if field.Flags&FieldAttributes_Private != 0 {
				f.Flags |= content.FLAG_ACC_PRIVATE
			} else if field.Flags&FieldAttributes_Family != 0 {
				f.Flags |= content.FLAG_ACC_PROTECTED
			}

			fields = append(fields, f)
		}
	}
	return fields, nil
}

func (a *Assembly) Parameters(index MethodDefIndex) (params []content.Variable, err error) {
	var (
		startRow, endRow = a.ListRange(index.Index(), id_MethodDef, id_Param, func(i interface{}) uint32 { return i.(*MethodDefRow).ParamList.Index() })
		idx              = ConcreteTableIndex{&a.MetadataUtil, startRow, id_Param}
	)
	for i := startRow; i < endRow; i++ {
		idx.index = i
		if rawparam, err := idx.Data(); err != nil {
			return nil, err
		} else {
			param := rawparam.(*ParamRow)
			var f content.Variable
			f.Name.Relative = string(param.Name)
			params = append(params, f)
		}
	}
	return params, nil
}

func (a *Assembly) Methods(index TypeDefIndex) (methods []content.Method, err error) {
	var (
		startRow, endRow = a.ListRange(index.Index(), id_TypeDef, id_MethodDef, func(i interface{}) uint32 { return i.(*TypeDefRow).MethodList.Index() })
		idx              = &ConcreteTableIndex{&a.MetadataUtil, startRow, id_MethodDef}
	)
	for i := startRow; i < endRow; i++ {
		idx.index = i
		if rawmethod, err := idx.Data(); err != nil {
			return nil, err
		} else {
			var (
				m      content.Method
				method = rawmethod.(*MethodDefRow)
				dec    *SignatureDecoder
				sig    MethodDefSig
			)
			m.Name.Relative = string(method.Name)
			if m.Parameters, err = a.Parameters(idx); err != nil {
				return nil, err
			}
			if dec, err = NewSignatureDecoder(method.Signature); err != nil {
				return nil, err
			} else if err = dec.Decode(&sig); err != nil {
				return nil, err
			} else {
				// TODO: need to figure out why this mismatch happens
				l := len(sig.Params)
				if l2 := len(m.Parameters); l2 < l {
					l = l2
				}
				for i := range sig.Params[:l] {
					m.Parameters[i].Type = a.ToContentType(index, &sig.Params[i].Type)
				}
				if method.Flags&MethodAttributes_Final != 0 {
					m.Flags |= content.FLAG_FINAL
				}
				if method.Flags&MethodAttributes_Static != 0 {
					m.Flags |= content.FLAG_STATIC
				}
				if method.Flags&MethodAttributes_Public != 0 {
					m.Flags |= content.FLAG_ACC_PUBLIC
				} else if method.Flags&MethodAttributes_Private != 0 {
					m.Flags |= content.FLAG_ACC_PRIVATE
				} else if method.Flags&MethodAttributes_Family != 0 {
					m.Flags |= content.FLAG_ACC_PROTECTED
				}

				m.Returns = make([]content.Variable, 1)
				m.Returns[0].Type = a.ToContentType(index, &sig.RetType.Type)
			}
			methods = append(methods, m)
		}
	}
	return methods, nil
}

func (a *Assembly) Extends(index TypeDefIndex) (t content.Type, err error) {
	if raw, err := index.Data(); err != nil {
		return t, err
	} else {
		row := raw.(*TypeDefRow)
		if (row.Flags & TypeAttributes_ClassSemanticsMask) != TypeAttributes_Class {
			return t, ErrInterface
		}
		if row.Extends.Index() != 0 {
			if raw, err := row.Extends.Data(); err != nil {
				return t, err
			} else {
				t = ToContentType(raw.(AbstractType))
			}
		}
	}
	return
}

func (a *Assembly) Search(index TypeDefIndex, tableId int, equal func(TableIndex) bool) TableIndex {
	table := a.Tables[tableId]
	ci := ConcreteTableIndex{metadataUtil: &a.MetadataUtil, index: 0, table: tableId}
	idx := sort.Search(int(table.Rows), func(in int) bool {
		i := uint32(in)
		ci.index = i + 1
		return equal(&ci)
	})
	if uint32(idx) == table.Rows {
		return nil
	}
	ci.index = uint32(idx) + 1
	return &ci
}

func (a *Assembly) Implements(index TypeDefIndex) (interfaces []content.Type, err error) {
	table := a.Tables[id_InterfaceImpl]
	rawidx := a.Search(index, id_InterfaceImpl, func(ti TableIndex) bool {
		if raw, err := ti.Data(); err == nil {
			c := raw.(*InterfaceImplRow)
			return c.Class.Table() >= index.Table() && c.Class.Index() >= index.Index()
		}
		return false
	})
	if rawidx == nil {
		return nil, nil
	}
	ci := rawidx.(*ConcreteTableIndex)
	for i := uint32(ci.index); i < table.Rows+1; i++ {
		ci.index = i
		if raw, err := ci.Data(); err != nil {
			return nil, err
		} else {
			c := raw.(*InterfaceImplRow)
			if c.Class.Index() != index.Index() {
				break
			}
			if raw, err := c.Interface.Data(); err != nil {
				return nil, err
			} else {
				interfaces = append(interfaces, ToContentType(raw.(AbstractType)))
			}
		}
	}
	return
}

func (a *Assembly) FindType(t content.FullyQualifiedName) (TypeDefIndex, error) {
	if t.Absolute == "" {
		return nil, errors.New("Can only look up types with a full absolute name")
	}
	idx := ConcreteTableIndex{&a.MetadataUtil, 0, id_TypeDef}
	for i := uint32(0); i < a.Tables[id_TypeDef].Rows; i++ {
		idx.index = 1 + i
		if rawtype, err := idx.Data(); err != nil {
			return nil, err
		} else {
			var (
				tr = rawtype.(*TypeDefRow)
			)
			if AbsoluteName(tr) == t.Absolute {
				return &idx, nil
			}
		}
	}
	return nil, nil
}

func (a *Assembly) Complete(t content.FullyQualifiedName) (*content.CompletionResult, error) {
	var ret content.CompletionResult
	if idx, err := a.FindType(t); err != nil {
		return nil, err
	} else if idx == nil {
		return nil, errors.New(fmt.Sprintf("Type not found in assembly: %s", t.Absolute))
	} else if ret.Fields, err = a.Fields(idx); err != nil {
		return nil, err
	} else if ret.Methods, err = a.Methods(idx); err != nil {
		return nil, err
	}

	return &ret, nil
}

func LoadAssembly(r io.ReadSeeker) (*Assembly, error) {

	var (
		br        = util.BinaryReader{r, binary.LittleEndian}
		err       error
		pe_offset uint32
		coff      coff_file_header
		cor20     image_cor20
		t         MetadataHeader
	)

	if _, err := r.Seek(pe_signature_offset, 0); err != nil {
		return nil, err
	}

	if pe_offset, err = br.Uint32(); err != nil {
		return nil, err
	} else if _, err := r.Seek(int64(pe_offset), 0); err != nil {
		return nil, err
	}

	if err := br.ReadInterface(&coff); err != nil {
		return nil, err
	}
	net := coff.OptionalHeader.RVAS[14]
	off := coff.VirtualToFileOffset(net.VirtualAddress)
	if _, err := br.Seek(int64(off), 0); err != nil {
		return nil, err
	}

	if err := br.ReadInterface(&cor20); err != nil {
		return nil, err
	}
	off = coff.VirtualToFileOffset(cor20.MetaData.VirtualAddress)
	if _, err := br.Seek(int64(off), 0); err != nil {
		return nil, err
	}
	if err := br.ReadInterface(&t); err != nil {
		return nil, err
	}
	if _, err := br.Seek(int64(off), 0); err != nil {
		return nil, err
	}
	if md, err := t.MetadataUtil(&br); err != nil {
		return nil, err
	} else {
		return &Assembly{*md}, nil
	}
	panic("Unreachable")
}
