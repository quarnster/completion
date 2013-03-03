package net

import (
	"bytes"
	"encoding/binary"
	"testing"
)

var (
	unsigned_tests = [][2]uint32{
		{0x03, 0x03000000},
		{0x7F, 0x7F000000},
		{0x80, 0x80800000},
		{0x2E57, 0xAE570000},
		{0x3FFF, 0xBFFF0000},
		{0x4000, 0xC0004000},
		{0x1FFFFFFF, 0xDFFFFFFF},
	}

	signed_tests = [][2]int64{
		{3, 0x06000000},
		{-3, 0x7B000000},
		{64, 0x80800000},
		{-64, 0x01000000},
		{8192, 0xC0004000},
		{-8192, 0x80010000},
		{268435455, 0xDFFFFFFE},
		{-268435456, 0xC0000001},
	}
)

func TestSignedDecoding(t *testing.T) {
	var (
		data = make([]byte, 4)
	)

	for _, test := range signed_tests {
		binary.BigEndian.PutUint32(data, uint32(test[1]))
		if d, err := SignedDecode(bytes.NewReader(data)); err != nil {
			t.Error(err)
		} else if d != int32(test[0]) {
			d1 := make([]byte, 4)
			d2 := make([]byte, 4)
			binary.BigEndian.PutUint32(d1, uint32(d))
			binary.BigEndian.PutUint32(d2, uint32(test[0]))
			t.Errorf("Unexpected value: %x != %x, (%v, %v)", d, test[0], d1, d2)
		}
	}
}

func TestUnsignedDecoding(t *testing.T) {
	var (
		data = make([]byte, 4)
	)

	for _, test := range unsigned_tests {
		binary.BigEndian.PutUint32(data, test[1])
		if d, err := UnsignedDecode(bytes.NewReader(data)); err != nil {
			t.Error(err)
		} else if d != test[0] {
			d1 := make([]byte, 4)
			d2 := make([]byte, 4)
			binary.BigEndian.PutUint32(d1, uint32(d))
			binary.BigEndian.PutUint32(d2, uint32(test[0]))
			t.Errorf("Unexpected value: %x != %x, (%v, %v)", d, test[0], d1, d2)
		}
	}
}
