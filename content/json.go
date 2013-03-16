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
	if err := RegisterType("settings", reflect.TypeOf(Settings{})); err != nil {
		panic(err)
	}
}

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

func (s *Settings) MarshalJSON() ([]byte, error) {
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
			val := reflect.New(t)
			i = val.Interface()
			if err := json.Unmarshal(v, i); err != nil {
				return err
			}
			i = val.Elem().Interface()
		} else if err := json.Unmarshal(v, &i); err != nil {
			return err
		}

		s.data[k] = i
	}
	return nil
}
