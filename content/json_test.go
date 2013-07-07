package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func comparesettings(s1, s2 *Settings) error {
	for k, v := range s1.data {
		if v2, ok := s2.data[k]; !ok {
			return errors.New(fmt.Sprintf("%s was in s1, but not s2", k))
		} else {
			t1, t2 := reflect.TypeOf(v), reflect.TypeOf(v2)
			if n1, n2 := t1.Name(), t2.Name(); n1 != n2 {
				return errors.New(fmt.Sprintf("Types don't match: %s != %s", n1, n2))
			} else if k1, k2 := t1.Kind(), t2.Kind(); k1 != k2 {
				return errors.New(fmt.Sprintf("Kind mismatch: %s != %s", k1, k2))
			} else if t := reflect.TypeOf(v); t == reflect.TypeOf(Settings{}) {
				ss1 := v.(Settings)
				ss2 := v2.(Settings)
				if err := comparesettings(&ss1, &ss2); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func TestJson(t *testing.T) {
	source := Settings{data: map[string]interface{}{
		"hello": "world",
		"yes":   true,
		"no":    3.0,
		"complex": CompletionResult{Fields: []Field{
			{Variable: Variable{Name: FullyQualifiedName{Relative: "Hello"}}},
		}},
	}}
	if err := RegisterType("complex", reflect.TypeOf(source.Get("complex"))); err != nil {
		t.Fatal(err)
	} else if err = RegisterType("complex", reflect.TypeOf(source.Get("complex"))); err != nil {
		t.Error(err)
	} else if err = RegisterType("complex", reflect.TypeOf(source)); err == nil {
		t.Error("Expected an error but didn't get one")
	}
	s1 := NewSettings()
	s1.Set("settings", source)
	if d, err := json.MarshalIndent(&s1, "", "\t"); err != nil {
		t.Error(err)
	} else {
		s2 := &Settings{}
		if err := json.Unmarshal(d, &s2); err != nil {
			t.Error(err)
		} else {
			t.Log(string(d))
			t.Log(s1)
			t.Log(s2)
			if err := comparesettings(s1, s2); err != nil {
				t.Error(err)
			}
		}
	}

	s1.Set("settings", &source)
	if _, err := json.Marshal(&s1); err == nil {
		t.Error("It shouldn't be possible to marshal a type that was registered as a struct, but is now a pointer to a struct")
	}
}
