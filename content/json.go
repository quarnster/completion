package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var registered_types map[string]reflect.Type

func init() {
	type reg struct {
		k string
		t reflect.Type
	}
	types := []reg{
		{"settings", reflect.TypeOf(Settings{})},
		{"location", reflect.TypeOf(SourceLocation{})},
		{"completion", reflect.TypeOf(CompletionResult{})},
	}
	for _, t := range types {
		if err := RegisterType(t.k, t.t); err != nil {
			panic(err)
		}
	}
}

// Register a type to make it possible to correctly deserialize it.
func RegisterType(key string, t reflect.Type) error {
	if registered_types == nil {
		registered_types = make(map[string]reflect.Type)
	}
	key = strings.ToLower(key)
	if v, ok := registered_types[key]; ok && v.Name() != t.Name() {
		return errors.New(fmt.Sprintf("Can't register type %s with key %s, as it is already used by type %s. ", t.Name(), key, v.Name()))
	}
	registered_types[key] = t
	return nil
}

func (s Settings) MarshalJSON() ([]byte, error) {
	for k, v := range s.data {
		registered, ok := registered_types[k]
		if !ok {
			continue
		}
		actual := reflect.TypeOf(v)
		if ak, rk := actual.Kind(), registered.Kind(); ak != rk {
			return nil, errors.New(fmt.Sprintf("Unable to marshal object %s with key %s, as it's registered as type %s != %s", actual, k, rk, ak))
		}
	}
	return json.Marshal(s.data)
}

func (s *Settings) UnmarshalJSON(data []byte) error {
	var tmp map[string]json.RawMessage

	s.data = make(map[string]interface{})

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	for k, v := range tmp {
		var i interface{}
		k = strings.ToLower(k)
		if t, ok := registered_types[k]; ok {
			ptr := reflect.New(t)
			i = ptr.Interface()
			if err := json.Unmarshal([]byte(v), i); err != nil {
				return err
			}
			i = ptr.Elem().Interface()
		} else if err := json.Unmarshal(v, &i); err != nil {
			return err
		}

		s.data[k] = i
	}
	return nil
}
