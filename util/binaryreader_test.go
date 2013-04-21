package util

import (
	"bytes"
	"encoding/binary"
	"errors"
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

func TestBinaryReaderIf(t *testing.T) {
	type A struct {
		A, B, C, D uint8
	}
	type C struct {
		A, B, C, D uint16
	}
	type Test struct {
		Magic uint32
		A     A `if:"Magic == 0x1"`
		B     uint8
		C     C `if:"Magic == 0x2"`
		D     uint8
	}

	var (
		t1 = Test{0, A{1, 2, 3, 4},
			5,
			C{6, 7, 8, 9},
			6}

		t2    Test
		tests = []uint32{1, 2, 0}
	)
	for _, test := range tests {
		b := bytes.NewBuffer(nil)
		t1.Magic = test
		if err := binary.Write(b, binary.LittleEndian, t1.Magic); err != nil {
			t.Fatal(err)
		}
		if t1.Magic == 1 {
			if err := binary.Write(b, binary.LittleEndian, t1.A); err != nil {
				t.Fatal(err)
			}
		}
		if err := b.WriteByte(t1.B); err != nil {
			t.Fatal(err)
		}
		if t1.Magic == 2 {
			if err := binary.Write(b, binary.LittleEndian, t1.C); err != nil {
				t.Fatal(err)
			}
		}
		if err := b.WriteByte(t1.D); err != nil {
			t.Fatal(err)
		}
		br := BinaryReader{Reader: bytes.NewReader(b.Bytes()), Endianess: binary.LittleEndian}
		if err := br.ReadInterface(&t2); err != nil {
			t.Error(err)
		} else {
			if t1.Magic != t2.Magic || t1.B != t2.B || t1.D != t2.D || (t1.Magic == 1 && (t1.A != t2.A)) || (t1.Magic == 2 && (t1.C != t2.C)) {
				t.Error(t1, t2)
			}
		}

	}
}

type ValidateableTest struct {
	Magic uint32
}

func (v *ValidateableTest) Validate() error {
	if v.Magic != 0x1337 {
		return errors.New(fmt.Sprintf("Not valid: %x", v.Magic))
	}
	return nil
}

func TestBinaryReaderValidateableFail(t *testing.T) {
	var (
		t1 = ValidateableTest{1337}
		t2 ValidateableTest
		b  = bytes.NewBuffer(nil)
	)
	if err := binary.Write(b, binary.LittleEndian, t1); err != nil {
		t.Fatal(err)
	}
	br := BinaryReader{Reader: bytes.NewReader(b.Bytes()), Endianess: binary.LittleEndian}
	if err := br.ReadInterface(&t2); err == nil {
		t.Error("Expected validation failure, but didn't get one")
	}
}

func TestBinaryReaderValidateablePass(t *testing.T) {
	var (
		t1 = ValidateableTest{0x1337}
		t2 ValidateableTest
		b  = bytes.NewBuffer(nil)
	)
	if err := binary.Write(b, binary.LittleEndian, t1); err != nil {
		t.Fatal(err)
	}
	br := BinaryReader{Reader: bytes.NewReader(b.Bytes()), Endianess: binary.LittleEndian}
	if err := br.ReadInterface(&t2); err != nil {
		t.Error(err)
	}
}
