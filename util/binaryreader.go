package util

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

type Validateable interface {
	Validate() error
}

type BinaryReader struct {
	Reader    io.ReadSeeker
	Endianess binary.ByteOrder
}

func (r *BinaryReader) ReadInterface(v interface{}) error {
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("Expected a pointer not %s", t.Kind()))
	}
	v2 := t.Elem()
	switch v2.Kind() {
	case reflect.Uint64:
		if d, err := r.Uint64(); err != nil {
			return err
		} else {
			v2.SetUint(uint64(d))
		}
	case reflect.Uint32:
		if d, err := r.Uint32(); err != nil {
			return err
		} else {
			v2.SetUint(uint64(d))
		}
	case reflect.Uint16:
		if d, err := r.Uint16(); err != nil {
			return err
		} else {
			v2.SetUint(uint64(d))
		}
	case reflect.Uint8:
		if d, err := r.Uint8(); err != nil {
			return err
		} else {
			v2.SetUint(uint64(d))
		}
	case reflect.Int64:
		if d, err := r.Int64(); err != nil {
			return err
		} else {
			v2.SetInt(int64(d))
		}
	case reflect.Int32:
		if d, err := r.Int32(); err != nil {
			return err
		} else {
			v2.SetInt(int64(d))
		}
	case reflect.Int16:
		if d, err := r.Int16(); err != nil {
			return err
		} else {
			v2.SetInt(int64(d))
		}
	case reflect.Int8:
		if d, err := r.Int8(); err != nil {
			return err
		} else {
			v2.SetInt(int64(d))
		}
	case reflect.Float32:
		if f, err := r.Float32(); err != nil {
			return err
		} else {
			v2.SetFloat(float64(f))
		}
	case reflect.Float64:
		if f, err := r.Float64(); err != nil {
			return err
		} else {
			v2.SetFloat(f)
		}
	case reflect.Struct:
		for i := 0; i < v2.NumField(); i++ {
			var (
				f    = v2.Field(i)
				f2   = v2.Type().Field(i)
				size = -1
				err  error
			)
			if fi := f2.Tag.Get("if"); fi != "" {
				arr := strings.Split(fi, ",")
				if f3 := v2.FieldByName(arr[0]); !f3.IsValid() {
					return errors.New(fmt.Sprintf("No such field: %s in %s", arr[0], v2.Type().Name()))
				} else if test, err := strconv.ParseUint(arr[1], 0, 64); err != nil {
					return err
				} else {
					if f3.Uint() != test {
						continue
					}
				}
			}
			if l := f2.Tag.Get("length"); l != "" {
				if v3 := v2.FieldByName(l); v3.IsValid() {
					size = int(v3.Uint())
				} else if size, err = strconv.Atoi(l); err != nil {
					return err
				}
			}

			switch f.Type().Kind() {
			case reflect.String:
				var data []byte
				if size >= 0 {
					if data, err = r.Read(size); err != nil {
						return err
					}
				} else {
					var max = math.MaxInt32
					if m := f2.Tag.Get("max"); m != "" {
						if max, err = strconv.Atoi(m); err != nil {
							return err
						}
					}

					for i := 0; i < max; i++ {
						if u, err := r.Uint8(); err != nil {
							return err
						} else if u == '\u0000' {
							size = i + 1
							break
						} else {
							data = append(data, u)
						}
					}
				}
				f.SetString(string(data))
			case reflect.Slice:
				if size == -1 {
					return errors.New("SliceHeader require a known length")
				}
				var v3 = reflect.MakeSlice(f.Type(), size, size)
				for i := 0; i < size; i++ {
					if err = r.ReadInterface(v3.Index(i).Addr().Interface()); err != nil {
						return err
					}
				}
				f.Set(v3)
			default:
				if err := r.ReadInterface(f.Addr().Interface()); err != nil {
					return err
				} else {
					size = int(f.Type().Size())
				}
			}

			if al := f2.Tag.Get("align"); al != "" {
				if a, err := strconv.Atoi(al); err != nil {
					return err
				} else if seek := ((size + (a - 1)) &^ (a - 1)) - size; seek > 0 {
					if _, err := r.Seek(int64(seek), 1); err != nil {
						return err
					}
				}
			}
		}
	default:
		return errors.New(fmt.Sprintf("Don't know how to read type %s", v2.Kind()))
	}
	if val, ok := v.(Validateable); ok {
		return val.Validate()
	}
	return nil
}

func (r *BinaryReader) Seek(offset int64, whence int) (int64, error) {
	return r.Reader.Seek(offset, whence)
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
}

func (r *BinaryReader) Uint32() (uint32, error) {
	if data, err := r.Read(4); err != nil {
		return 0, err
	} else {
		return r.Endianess.Uint32(data), nil
	}
}

func (r *BinaryReader) Uint16() (uint16, error) {
	if data, err := r.Read(2); err != nil {
		return 0, err
	} else {
		return r.Endianess.Uint16(data), nil
	}
}

func (r *BinaryReader) Uint8() (uint8, error) {
	if data, err := r.Read(1); err != nil {
		return 0, err
	} else {
		return uint8(data[0]), nil
	}
}

func (r *BinaryReader) Int64() (int64, error) {
	if data, err := r.Uint64(); err != nil {
		return 0, err
	} else {
		return int64(data), nil
	}
}

func (r *BinaryReader) Int32() (int32, error) {
	if data, err := r.Uint32(); err != nil {
		return 0, err
	} else {
		return int32(data), nil
	}
}

func (r *BinaryReader) Int16() (int16, error) {
	if data, err := r.Uint16(); err != nil {
		return 0, err
	} else {
		return int16(data), nil
	}
}

func (r *BinaryReader) Int8() (int8, error) {
	if data, err := r.Read(1); err != nil {
		return 0, err
	} else {
		return int8(data[0]), nil
	}
}

func (r *BinaryReader) Float32() (float32, error) {
	if i32, err := r.Int32(); err != nil {
		return 0, err
	} else {
		f32 := *(*float32)(unsafe.Pointer(&i32))
		return f32, nil
	}
}

func (r *BinaryReader) Float64() (float64, error) {
	if i64, err := r.Int64(); err != nil {
		return 0, err
	} else {
		f64 := *(*float64)(unsafe.Pointer(&i64))
		return f64, nil
	}
}
