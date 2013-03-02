package net

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

type Guid uintptr

func (g Guid) String() string {
	if uintptr(g) == 0 {
		return "nil"
	}
	var (
		bytes  []byte
		native = binary.LittleEndian
	)
	goArray(unsafe.Pointer(&bytes), uintptr(g), 16)

	return fmt.Sprintf("%08X-%04X-%04X-%04X-%04X%04X%04X",
		native.Uint32(bytes[:4]),
		native.Uint16(bytes[4:6]),
		native.Uint16(bytes[6:8]),
		binary.BigEndian.Uint16(bytes[8:10]),
		binary.BigEndian.Uint16(bytes[10:12]),
		binary.BigEndian.Uint16(bytes[12:14]),
		binary.BigEndian.Uint16(bytes[14:16]),
	)
}
