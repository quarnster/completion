package dwarf

import "github.com/quarnster/util/encoding/binary"

type (
	AttributeSpecification struct {
		Name DW_AT
		Form DW_FORM
	}

	AbbrevEntry struct {
		Code       LEB128
		Tag        DW_TAG
		Children   DW_CHILDREN
		Attributes []AttributeSpecification
	}
)

func (ae *AbbrevEntry) Read(br *binary.BinaryReader) error {
	if err := br.ReadInterface(&ae.Code); err != nil {
		return err
	} else if err := br.ReadInterface(&ae.Tag); err != nil {
		return err
	} else if err := br.ReadInterface(&ae.Children); err != nil {
		return err
	}
	for {
		var v AttributeSpecification
		if err := br.ReadInterface(&v); err != nil {
			return err
		}
		if v.Name == 0 && v.Form == 0 {
			break
		}
		ae.Attributes = append(ae.Attributes, v)
	}
	return nil
}
