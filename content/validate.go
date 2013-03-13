package content

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

type Validateable interface {
	Validate() error
}

func isZero(value reflect.Value) (bool, error) {
	zeroValue := reflect.Zero(value.Type())
	switch value.Kind() {
	case reflect.String:
		return zeroValue.String() == value.String(), nil
	case reflect.Uint:
		return zeroValue.Uint() == value.Uint(), nil
	}
	return false, errors.New(fmt.Sprintf("Don't know how to validate kind %s", value.Kind()))
}

func Validate(v interface{}) error {
	if v2, ok := v.(Validateable); ok {
		if err := v2.Validate(); err != nil {
			return err
		}
	}
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("Expected a pointer not %s", t.Kind()))
	}
	v2 := t.Elem()
	errStr := ""
	switch v2.Kind() {
	case reflect.Struct:
		for i := 0; i < v2.NumField(); i++ {
			f := v2.Type().Field(i)
			options := f.Tag.Get("protocol")
			if options != "required" {
				continue
			}
			if err := Validate(v2.Field(i).Addr().Interface()); err != nil {
				errStr += fmt.Sprintf("Required field %s.%s is invalid: %s", v2.Type().Name(), f.Name, err)
			} else if f.Type.Kind() != reflect.Struct {
				if zero, err := isZero(v2.Field(i)); err != nil {
					return err
				} else if zero {
					if errStr != "" {
						errStr += "\n"
					}
					errStr += fmt.Sprintf("Required field %s.%s is the nil value", v2.Type().Name(), f.Name)
				}
			}
		}
	}
	if errStr != "" {
		return errors.New(errStr)
	}
	return nil
}

func (t *Type) Validate() error {
	switch t.Flags & FLAG_TYPE_MASK {
	case FLAG_TYPE_POINTER, FLAG_TYPE_ARRAY:
		if l := len(t.Specialization); l != 1 {
			return errors.New(fmt.Sprintf("Expected specialization length mismatch: %d != 1", l))
		}
	case FLAG_TYPE_METHOD:
		if l := len(t.Methods); l != 1 {
			return errors.New(fmt.Sprintf("Expected method length mismatch: %d != 1", l))
		}
	default:
		return Validate(&t.Name)
	}
	return nil
}

var badre = regexp.MustCompile(`(^[.$][^.]|[^\\][.$][^.])`)

func (f *FullyQualifiedName) Validate() error {
	if badre.MatchString(f.Relative) {
		str := badre.FindString(f.Relative)
		return errors.New(fmt.Sprintf("Relative name contains illegal characters: %s (%c)", f.Relative, str[0]))
	}
	return nil
}
