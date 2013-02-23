package common

import (
	"encoding/binary"
	"errors"
	"io"
	"unsafe"
)

type BinaryReader struct {
	Reader    io.Reader
	Endianess binary.ByteOrder
}

func (r *BinaryReader) Read(size int) ([]byte, error) {
	data := make([]byte, size)
	if size == 0 {
		return data, nil
	}
	if n, err := r.Reader.Read(data); err != nil {
		return nil, err
	} else if n != len(data) {
		return nil, errors.New("Didn't read the expected number of bytes")
	}
	return data, nil
}

func (r *BinaryReader) Uint64() (uint64, error) {
	if data, err := r.Read(8); err != nil {
		return 0, err
	} else {
		return r.Endianess.Uint64(data), nil
	}
	panic("Unreachable")
}

func (r *BinaryReader) Uint32() (uint32, error) {
	if data, err := r.Read(4); err != nil {
		return 0, err
	} else {
		return r.Endianess.Uint32(data), nil
	}
	panic("Unreachable")
}

func (r *BinaryReader) Uint16() (uint16, error) {
	if data, err := r.Read(2); err != nil {
		return 0, err
	} else {
		return r.Endianess.Uint16(data), nil
	}
	panic("Unreachable")
}

func (r *BinaryReader) Int64() (int64, error) {
	if data, err := r.Uint64(); err != nil {
		return 0, err
	} else {
		return int64(data), nil
	}
	panic("Unreachable")
}

func (r *BinaryReader) Int32() (int32, error) {
	if data, err := r.Uint32(); err != nil {
		return 0, err
	} else {
		return int32(data), nil
	}
	panic("Unreachable")
}

func (r *BinaryReader) Int16() (int16, error) {
	if data, err := r.Uint16(); err != nil {
		return 0, err
	} else {
		return int16(data), nil
	}
	panic("Unreachable")
}

func (r *BinaryReader) Float32() (float32, error) {
	if i32, err := r.Int32(); err != nil {
		return 0, err
	} else {
		f32 := *(*float32)(unsafe.Pointer(&i32))
		return f32, nil
	}
	panic("Unreachable")
}

func (r *BinaryReader) Float64() (float64, error) {
	if i64, err := r.Int64(); err != nil {
		return 0, err
	} else {
		f64 := *(*float64)(unsafe.Pointer(&i64))
		return f64, nil
	}
	panic("Unreachable")
}
