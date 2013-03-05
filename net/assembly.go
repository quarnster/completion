package net

// references:
// http://msdn.microsoft.com/en-us/library/windows/hardware/gg463119.aspx
// http://www.codeproject.com/Articles/12585/The-NET-File-Format
// http://msdn.microsoft.com/en-us/library/ms809762.aspx
// http://www.ecma-international.org/publications/standards/Ecma-335.htm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util"
	"io"
	"reflect"
	"sort"
	"unsafe"
)

var (
	ErrInterface   = errors.New("TypeDef is an interface, not a class")
	ErrNotAssembly = errors.New("This does not appear to be a .net assembly")
)

type Assembly struct {
	MetadataUtil
	rawdata []byte
}

type AbstractType interface {
	Name() string
	Namespace() string
}

func ToContentType(t AbstractType) (t2 content.Type) {
	t2.Name.Relative = t.Name()
	if ns := t.Namespace(); ns != "" {
		t2.Name.Absolute = ns + "." + t2.Name.Relative
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
					m.Parameters[i].Type = ToContentType(&sig.Params[i].Type)
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
				m.Returns[0].Type = ToContentType(&sig.RetType.Type)
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

func (a *Assembly) Implements(index TypeDefIndex) (interfaces []content.Type, err error) {
	table := a.Tables[id_InterfaceImpl]
	ci := ConcreteTableIndex{metadataUtil: &a.MetadataUtil, index: 0, table: id_InterfaceImpl}
	idx := sort.Search(int(table.Rows), func(in int) bool {
		i := uint32(in)
		ci.index = i + 1
		if raw, err := ci.Data(); err == nil {
			c := raw.(*InterfaceImplRow)
			return c.Class.Index() == index.Index()
		}
		return false
	})

	for i := uint32(idx); i < table.Rows; i++ {
		ci.index = i + 1
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

type Validateable interface {
	Validate() error
}

func assertEndianness() error {
	test := uint32(0x01020304)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, test)
	if test2 := *(*uint32)(unsafe.Pointer(&bytes[0])); test != test2 {
		return errors.New("This module assumes a little endian system")
	}
	return nil
}

func goArray(a unsafe.Pointer, ptr uintptr, l int) {
	sliceHeader := (*reflect.SliceHeader)(a)
	sliceHeader.Cap = l
	sliceHeader.Len = l
	sliceHeader.Data = ptr
}

func LoadAssembly(r io.ReadSeeker) (*Assembly, error) {
	if err := assertEndianness(); err != nil {
		return nil, err
	}

	var (
		br        = util.BinaryReader{r, binary.LittleEndian}
		err       error
		isPe32    bool
		data      []byte
		pe_offset uint32
	)

	if size, err := r.Seek(0, 2); err != nil {
		return nil, err
	} else if _, err := r.Seek(0, 0); err != nil {
		return nil, err
	} else if data, err = br.Read(int(size)); err != nil {
		return nil, err
	}

	r = bytes.NewReader(data)
	br.Reader = r

	if _, err := r.Seek(pe_signature_offset, 0); err != nil {
		return nil, err
	}

	if pe_offset, err = br.Uint32(); err != nil {
		return nil, err
	} else if _, err := r.Seek(int64(pe_offset), 0); err != nil {
		return nil, err
	}

	if check, err := br.Read(len(pe_magic)); err != nil {
		return nil, err
	} else if bytes.Compare(check, pe_magic) != 0 {
		return nil, errors.New(fmt.Sprintf("PE Magic mismatch: %v", check))
	}

	coff := (*coff_file_header)(unsafe.Pointer(&data[int(pe_offset)+len(pe_magic)]))

	var opt_t reflect.Type
	ptr := uintptr(unsafe.Pointer(coff))
	ptr += unsafe.Sizeof(*coff)
	opt_pe32 := (*optional_header_pe32)(unsafe.Pointer(ptr))
	opt_pe32p := (*optional_header_pe32p)(unsafe.Pointer(ptr))

	if opt_pe32.Magic != pe32 && opt_pe32.Magic != pe32p {
		return nil, errors.New(fmt.Sprintf("Unsupported optional header magic: %x", opt_pe32.Magic))
	} else {
		isPe32 = opt_pe32.Magic == pe32
	}
	rvas := 0
	if isPe32 {
		rvas = int(opt_pe32.NumberOfRvaAndSizes)
		opt_t = reflect.TypeOf(*opt_pe32)
	} else {
		rvas = int(opt_pe32p.NumberOfRvaAndSizes)
		opt_t = reflect.TypeOf(*opt_pe32p)
	}
	ptr += opt_t.Size()
	var ids []image_data_directory
	goArray(unsafe.Pointer(&ids), ptr, rvas)

	sections := coff.SectionTable()
	for i := range sections {
		if err := sections[i].Validate(); err != nil {
			return nil, err
		}
	}
	net := ids[14]
	if net.VirtualAddress == 0 || net.Size == 0 {
		return nil, ErrNotAssembly
	}
	off := uint32(0)
	sec := 0
	for net.VirtualAddress > sections[sec+1].VirtualAddress && sec < len(sections)-2 {
		sec++
	}
	off = net.VirtualAddress - sections[sec].VirtualAddress + sections[sec].PointerToRawData

	cor20 := (*image_cor20)(unsafe.Pointer(&data[off]))
	if cor20.MetaData.VirtualAddress == 0 || cor20.MetaData.Size == 0 {
		return nil, ErrNotAssembly
	}
	sec = 0
	for cor20.MetaData.VirtualAddress > sections[sec+1].VirtualAddress && sec < len(sections)-2 {
		sec++
	}

	off = cor20.MetaData.VirtualAddress - sections[sec].VirtualAddress + sections[sec].PointerToRawData
	t := (*MetadataHeader)(unsafe.Pointer(&data[off]))
	if err := t.Validate(); err != nil {
		return nil, err
	}
	if md, err := t.MetadataUtil(); err != nil {
		return nil, err
	} else {
		return &Assembly{*md, data}, nil
	}
	panic("Unreachable")
}
