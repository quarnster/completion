package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBinaryReaderSimple(t *testing.T) {
	type Test struct {
		A uint32
		B uint16
		C uint8
		D uint32
	}
	var (
		t1 = Test{1, 2, 3, 4}
		t2 Test
		b  = bytes.NewBuffer(nil)
	)
	if err := binary.Write(b, binary.LittleEndian, t1); err != nil {
		t.Fatal(err)
	}
	br := BinaryReader{Reader: bytes.NewReader(b.Bytes()), Endianess: binary.LittleEndian}

	if err := br.ReadInterface(&t2); err != nil {
		t.Error(err)
	} else if t1 != t2 {
		t.Error(t1, t2)
	}
}

func TestBinaryReaderAlign(t *testing.T) {
	type Test struct {
		A uint32
		B uint16
		C uint8 `align:"2"`
		D uint32
	}
	var (
		t1 = Test{1, 2, 3, 4}
		t2 Test
		b  = bytes.NewBuffer(nil)
	)
	if err := binary.Write(b, binary.LittleEndian, t1.A); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(b, binary.LittleEndian, t1.B); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(b, binary.LittleEndian, t1.C); err != nil {
		t.Fatal(err)
	}
	if err := b.WriteByte('\u0000'); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(b, binary.LittleEndian, t1.D); err != nil {
		t.Fatal(err)
	}

	br := BinaryReader{Reader: bytes.NewReader(b.Bytes()), Endianess: binary.LittleEndian}

	if err := br.ReadInterface(&t2); err != nil {
		t.Error(err)
	} else if t1 != t2 {
		t.Error(t1, t2)
	}
}

func TestBinaryReaderAlign2(t *testing.T) {
	type Test struct {
		A uint32
		B uint16 `align:"4"`
		C uint8  `align:"4"`
		D uint32
	}
	var (
		t1 = Test{1, 2, 3, 4}
		t2 Test
		b  = bytes.NewBuffer(nil)
	)
	if err := binary.Write(b, binary.LittleEndian, []uint32{1, 2, 3, 4}); err != nil {
		t.Fatal(err)
	}
	br := BinaryReader{Reader: bytes.NewReader(b.Bytes()), Endianess: binary.LittleEndian}

	if err := br.ReadInterface(&t2); err != nil {
		t.Error(err)
	} else if t1 != t2 {
		t.Error(t1, t2)
	}
}

func TestBinaryReaderStringLength(t *testing.T) {
	type Test struct {
		Length uint32
		Data   string `length:"Length"`
		C      uint32
	}
	var (
		str = "Hello World!"
		t1  = Test{uint32(len(str)), str, 1337}
		t2  Test
		b   = bytes.NewBuffer(nil)
	)
	if err := binary.Write(b, binary.LittleEndian, t1.Length); err != nil {
		t.Fatal(err)
	}
	if _, err := b.Write([]byte(t1.Data)); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(b, binary.LittleEndian, t1.C); err != nil {
		t.Fatal(err)
	}

	br := BinaryReader{Reader: bytes.NewReader(b.Bytes()), Endianess: binary.LittleEndian}

	if err := br.ReadInterface(&t2); err != nil {
		t.Error(err)
	} else if t1 != t2 {
		t.Error(t1, t2)
	}
}

func TestBinaryReaderStringLengthAlign(t *testing.T) {
	type Test struct {
		Length uint32
		Data   string `length:"Length" align:"16"`
		C      uint32
	}
	var (
		str = "Hello World!"
		t1  = Test{uint32(len(str)), str, 1337}
		t2  Test
		b   = bytes.NewBuffer(nil)
	)
	if t1.Length == 16 {
		t.Fatal("It's not supposed to be aligned already")
	}
	if err := binary.Write(b, binary.LittleEndian, t1.Length); err != nil {
		t.Fatal(err)
	}
	if _, err := b.Write([]byte(t1.Data)); err != nil {
		t.Fatal(err)
	}
	if _, err := b.Write(make([]byte, 16-len(t1.Data))); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(b, binary.LittleEndian, t1.C); err != nil {
		t.Fatal(err)
	}

	br := BinaryReader{Reader: bytes.NewReader(b.Bytes()), Endianess: binary.LittleEndian}

	if err := br.ReadInterface(&t2); err != nil {
		t.Error(err)
	} else if t1 != t2 {
		t.Error(t1, t2)
	}
}

func TestBinaryReaderString(t *testing.T) {
	type Test struct {
		Data string `max:"32"`
		C    uint32
	}
	var (
		str = "Hello World!"
		t1  = Test{str, 1337}
		t2  Test
		b   = bytes.NewBuffer(nil)
	)
	if _, err := b.Write([]byte(t1.Data)); err != nil {
		t.Fatal(err)
	}
	if err := b.WriteByte('\u0000'); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(b, binary.LittleEndian, t1.C); err != nil {
		t.Fatal(err)
	}

	br := BinaryReader{Reader: bytes.NewReader(b.Bytes()), Endianess: binary.LittleEndian}

	if err := br.ReadInterface(&t2); err != nil {
		t.Error(err)
	} else if t1 != t2 {
		t.Error(t1, t2)
	}
}

func TestBinaryReaderStringMax(t *testing.T) {
	type Test struct {
		Data string `max:"12"`
		C    uint32
	}
	var (
		str = "Hello World!"
		t1  = Test{str, 1337}
		t2  Test
		b   = bytes.NewBuffer(nil)
	)
	if _, err := b.Write([]byte(t1.Data)); err != nil {
		t.Fatal(err)
	}
	if err := binary.Write(b, binary.LittleEndian, t1.C); err != nil {
		t.Fatal(err)
	}

	br := BinaryReader{Reader: bytes.NewReader(b.Bytes()), Endianess: binary.LittleEndian}

	if err := br.ReadInterface(&t2); err != nil {
		t.Error(err)
	} else if t1 != t2 {
		t.Error(t1, t2)
	}
}

func TestBinaryReaderComplex(t *testing.T) {
	type inner struct {
		Data string `max:"32" align:"4"`
	}
	type Test struct {
		Count    uint32
		Contents []inner `length:"Count"`
	}
	var (
		t1 = Test{
			2,
			[]inner{
				inner{"Hello"},
				inner{"World"},
			}}
		t2 Test
		b  = bytes.NewBuffer(nil)
	)
	if err := binary.Write(b, binary.LittleEndian, t1.Count); err != nil {
		t.Fatal(err)
	}
	for i := range t1.Contents {
		if _, err := b.Write([]byte(t1.Contents[i].Data)); err != nil {
			t.Fatal(err)
		}
		if _, err := b.Write([]byte("\u0000\u0000\u0000")); err != nil {
			t.Fatal(err)
		}
	}

	br := BinaryReader{Reader: bytes.NewReader(b.Bytes()), Endianess: binary.LittleEndian}
	if err := br.ReadInterface(&t2); err != nil {
		t.Error(err)
	} else {
		s1 := fmt.Sprintf("%#v", t1)
		s2 := fmt.Sprintf("%#v", t2)
		if s1 != s2 {
			t.Error(t1, t2)
		}
	}
}
