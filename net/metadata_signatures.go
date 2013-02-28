package net

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// II.23.2

func encodedWidth(first byte) int {
	if first&(1<<7) == 0 {
		// one byte value
		return 1
	} else if first&(1<<6) == 0 {
		// two byte value
		return 2
	} else if first&(1<<5) == 0 {
		// four byte value
		return 4
	}
	return -1
}

func UnsignedDecode(data []byte) (uint32, error) {
	switch encodedWidth(data[0]) {
	case 1:
		return uint32(data[0]), nil
	case 2:
		return uint32(binary.BigEndian.Uint16(data)) &^ 0x8000, nil
	case 4:
		return binary.BigEndian.Uint32(data) &^ 0xc0000000, nil
	}
	return 0, errors.New(fmt.Sprintf("Invalid bitmask signature: %x", data[0]))
}

func SignedDecode(data []byte) (ret int32, err error) {
	if d, err := UnsignedDecode(data); err != nil {
		return 0, err
	} else {
		sign := (d & 1)
		switch encodedWidth(data[0]) {
		case 1:
			sign <<= 7
			sign |= sign >> 1
			return int32(int8(d>>1 | sign)), nil
		case 2:
			sign <<= 15
			sign |= sign >> 1
			sign |= sign >> 1
			return int32(int16(d>>1 | sign)), nil
		case 4:
			sign <<= 31
			sign |= sign >> 1
			sign |= sign >> 1
			sign |= sign >> 1
			return int32(int32(d>>1 | sign)), nil
		}
	}
	panic("unreachable")
}
