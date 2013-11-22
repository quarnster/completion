package net

// references:
// http://msdn.microsoft.com/en-us/library/windows/hardware/gg463119.aspx
// http://www.codeproject.com/Articles/12585/The-NET-File-Format
// http://msdn.microsoft.com/en-us/library/ms809762.aspx
// http://www.ecma-international.org/publications/standards/Ecma-335.htm

import (
	"code.google.com/p/log4go"
	"errors"
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/util/encoding/binary"
	"io"
)

var (
	ErrNotAssembly = errors.New("This does not appear to be a .net assembly")
)

type Assembly struct {
	MetadataUtil
	typelut map[string]*TypeDef
}

func (a *Assembly) Name() string {
	ci := ConcreteTableIndex{&a.MetadataUtil, 1, id_Module}
	if raw, err := ci.Data(); err == nil {
		mr := raw.(*ModuleRow)
		return string(mr.Name)
	}
	return ""
}

func (a *Assembly) Types() (types []content.Type, err error) {
	var (
		idx = ConcreteTableIndex{&a.MetadataUtil, 0, id_TypeDef}
	)
	for i := uint32(0); i < a.Tables[id_TypeDef].Rows; i++ {
		idx.index = 1 + i
		rawtype, err := idx.Data()
		if err != nil {
			return nil, err
		}
		var (
			tr = rawtype.(*TypeDefRow)
		)
		tc := ToContentType(tr)
		if err := check(&tc, tc.Name); err != nil {
			log4go.Debug("Skipping type %v, %s", tc, err)
			continue
		}
		types = append(types, tc)
	}
	return
}

func (a *Assembly) FindType(t content.FullyQualifiedName) (*TypeDef, error) {
	if t.Absolute == "" {
		return nil, errors.New("Can only look up types with a full absolute name")
	}
	return a.typelut[t.Absolute], nil
}

var err_tnf = errors.New(fmt.Sprintf("Type not found in assembly"))

func (a *Assembly) Complete(t *content.Type) (*content.CompletionResult, error) {
	if td, err := a.FindType(t.Name); err != nil {
		return nil, err
	} else if td == nil {
		return nil, err_tnf
	} else if ct, err := td.ToContentType(); err != nil {
		return nil, err
	} else {
		return (*content.CompletionResult)(ct), nil
	}
}

func LoadAssembly(r io.ReadSeeker) (*Assembly, error) {

	var (
		br        = binary.BinaryReader{r, binary.LittleEndian}
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
	md, err := t.MetadataUtil(&br)
	if err != nil {
		return nil, err
	}
	a := Assembly{*md, make(map[string]*TypeDef)}
	idx := ConcreteTableIndex{&a.MetadataUtil, 0, id_TypeDef}
	for i := uint32(0); i < a.Tables[id_TypeDef].Rows; i++ {
		idx.index = 1 + i
		if td, err := TypeDefFromIndex(&idx); err == nil {
			a.typelut[td.Name().Absolute] = td
		}
	}
	return &a, nil
}
