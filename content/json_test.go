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
	source := Settings{map[string]interface{}{
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

func TestJson2(t *testing.T) {
	it := NewIntent("completion.complete.location")
	it.Data.Set("location", SourceLocation{File{Name: "hello.cpp"}, 10, 2})
	it.Data.Set("sessionid", 1337)
	settings := NewSettings()
	settings.Set("compilation_options", []string{"-Iincludepath1", "-Iincludepath2"})
	it.Data.Set("settings", *settings)

	if d, err := json.MarshalIndent(it, "", "\t"); err != nil {
		t.Error(err)
	} else {
		t.Log(string(d))
	}

	// obviously the proper action would be taken in the "real" code
	if cmp, err := loadData(); err != nil {
		t.Error(err)
	} else {
		cmp.Methods = cmp.Methods[:3] // Just to make the result smaller in this mock up use case
		cmp.Fields = cmp.Fields[:3]   // Just to make the result smaller in this mock up use case
		r := NewResponse()
		r.Data.Set("completions", cmp)
		if d, err := json.MarshalIndent(r, "", "\t"); err != nil {
			t.Error(err)
		} else {
			t.Log(string(d))
		}
	}

	// and in the case of an error
	{
		r := NewResponse()
		r.Data.Set("error", "something") // Could include a stacktrace etc
		if d, err := json.MarshalIndent(r, "", "\t"); err != nil {
			t.Error(err)
		} else {
			t.Log(string(d))
		}
	}
}
