package net

import (
	"encoding/binary"
	"fmt"
)

type Guid []byte

func (g Guid) String() string {
	var (
		bytes  = []byte(g)
		native = binary.LittleEndian
	)
	if len(bytes) != 16 {
		return "nil"
	}

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
