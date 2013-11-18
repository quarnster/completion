package dwarf

import (
	"github.com/quarnster/util/encoding/binary"
)

type (
	LEB128  uint
	SLEB128 int
)

func (l *LEB128) Read(br *binary.BinaryReader) error {
	curr := LEB128(0)
	shift := LEB128(0)
	for {
		v, err := br.Uint8()
		if err != nil {
			return err
		}
		br := v&0x80 == 0
		v = v & 0x7f
		curr |= LEB128(v) << shift
		if br {
			break
		}
		shift += 7
	}
	*l = curr
	return nil
}

func (l *SLEB128) Read(br *binary.BinaryReader) error {
	var (
		curr  LEB128
		shift LEB128
		v     uint8
		err   error
	)
	for {
		v, err = br.Uint8()
		if err != nil {
			return err
		}
		br := v&0x80 == 0
		v = v & 0x7f
		curr |= LEB128(v) << shift
		shift += 7
		if br {
			break
		}
	}
	const size = 64
	if (shift < size) && v&0x40 != 0 {
		curr |= -(1 << shift)
	}
	*l = SLEB128(curr)

	return nil
}

func (dw *DW_ACCESS) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_ACCESS(v)
	return nil
}
func (dw *DW_AT) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_AT(v)
	return nil
}
func (dw *DW_ATE) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_ATE(v)
	return nil
}
func (dw *DW_CC) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_CC(v)
	return nil
}
func (dw *DW_CHILDREN) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_CHILDREN(v)
	return nil
}
func (dw *DW_DS) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_DS(v)
	return nil
}
func (dw *DW_DSC) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_DSC(v)
	return nil
}
func (dw *DW_END) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_END(v)
	return nil
}
func (dw *DW_FORM) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_FORM(v)
	return nil
}
func (dw *DW_ID) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_ID(v)
	return nil
}
func (dw *DW_INL) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_INL(v)
	return nil
}
func (dw *DW_LANG) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_LANG(v)
	return nil
}

func (dw *DW_MACINFO) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_MACINFO(v)
	return nil
}
func (dw *DW_ORD) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_ORD(v)
	return nil
}
func (dw *DW_OP) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_OP(v)
	return nil
}
func (dw *DW_TAG) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_TAG(v)
	return nil
}
func (dw *DW_VIRTUALITY) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_VIRTUALITY(v)
	return nil
}
func (dw *DW_VIS) Read(br *binary.BinaryReader) error {
	var v LEB128
	if err := br.ReadInterface(&v); err != nil {
		return err
	}
	*dw = DW_VIS(v)
	return nil
}
