package content

import (
	"errors"
	"fmt"
	"reflect"
)

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
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("Expected a pointer not %s", t.Kind()))
	}
	v2 := t.Elem()
	if v2.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("Expected an element type of struct not %s", v2.Kind()))
	}
	errStr := ""
	for i := 0; i < v2.NumField(); i++ {
		f := v2.Type().Field(i)
		options := f.Tag.Get("protocol")
		if options != "required" {
			continue
		}
		if f.Type.Kind() == reflect.Struct {
			if err := Validate(v2.Field(i).Addr().Interface()); err != nil {
				errStr += fmt.Sprintf("Required field %s.%s of is invalid: %s", v2.Type().Name(), f.Name, err)
			}
		} else if zero, err := isZero(v2.Field(i)); err != nil {
			return err
		} else if zero {
			if errStr != "" {
				errStr += "\n"
			}
			errStr += fmt.Sprintf("Required field %s.%s is the nil value", v2.Type().Name(), f.Name)
		}
	}
	if errStr != "" {
		return errors.New(errStr)
	}
	return nil
}
