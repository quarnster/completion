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
	"unsafe"
)

type Assembly struct {
	MetadataUtil
}

func TypeToContentType(t *Type, t2 *content.Type) {
	t2.Name.Relative = t.Name()
	if ns := t.Namespace(); ns != "" {
		t2.Name.Absolute = ns + "." + t2.Name.Relative
	}
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
		endRow = startRow
	}
	if endRow < startRow {
		endRow = tableEnd
	}
	if endRow > tableEnd {
		endRow = tableEnd
	}
	return
}

func (a *Assembly) Fields(index uint32) (fields []content.Field, err error) {
	var (
		startRow, endRow = a.ListRange(index, id_TypeDef, id_Field, func(i interface{}) uint32 { return i.(*TypeDefRow).FieldList.Index() })
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
				TypeToContentType(&sig.Type, &f.Type)
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

func (a *Assembly) Parameters(index uint32) (params []content.Variable, err error) {
	var (
		startRow, endRow = a.ListRange(index, id_MethodDef, id_Param, func(i interface{}) uint32 { return i.(*MethodDefRow).ParamList.Index() })
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

func (a *Assembly) Methods(index uint32) (methods []content.Method, err error) {
	var (
		startRow, endRow = a.ListRange(index, id_TypeDef, id_MethodDef, func(i interface{}) uint32 { return i.(*TypeDefRow).MethodList.Index() })
		idx              = ConcreteTableIndex{&a.MetadataUtil, startRow, id_MethodDef}
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
			if m.Parameters, err = a.Parameters(i); err != nil {
				return nil, err
			}
			if dec, err = NewSignatureDecoder(method.Signature); err != nil {
				return nil, err
			} else if err = dec.Decode(&sig); err != nil {
				return nil, err
			} else if a, b := len(sig.Params), len(m.Parameters); a != b {
				return nil, errors.New(fmt.Sprintf("Mismatched parameter count: %d != %d (%v, %v)", a, b, m, sig))
			} else {

				for i := range sig.Params {
					TypeToContentType(&sig.Params[i].Type, &m.Parameters[i].Type)
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
				TypeToContentType(&sig.RetType.Type, &m.Returns[0].Type)
			}
			methods = append(methods, m)
		}
	}
	return methods, nil
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
	off := net.VirtualAddress - sections[0].VirtualAddress + sections[0].PointerToRawData

	cor20 := (*image_cor20)(unsafe.Pointer(&data[off]))

	off = cor20.MetaData.VirtualAddress - sections[0].VirtualAddress + sections[0].PointerToRawData
	t := (*MetadataHeader)(unsafe.Pointer(&data[off]))
	if err := t.Validate(); err != nil {
		return nil, err
	}
	if md, err := t.MetadataUtil(); err != nil {
		return nil, err
	} else {
		return &Assembly{*md}, nil
	}
	panic("Unreachable")
}
