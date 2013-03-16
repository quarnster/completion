package content

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJson(t *testing.T) {
	source := Settings{map[string]interface{}{
		"hello": "world",
		"yes":   true,
		"no":    3.0,
		"complex": CompletionResult{Fields: []Field{
			{Variable: Variable{Name: FullyQualifiedName{Relative: "Hello"}}},
		}},
	}}
	RegisterType("complex", reflect.TypeOf(source.Get("complex")))

	s1 := NewSettings()
	s1.Set("settings", source)
	if d, err := json.MarshalIndent(&s1, "", "\t"); err != nil {
		t.Error(err)
	} else {
		s2 := &Settings{}
		if err := json.Unmarshal(d, &s2); err != nil {
			t.Error(err)
		} else {
			for k, v := range s1.data {
				if v2, ok := s2.data[k]; !ok {
					t.Errorf("%s was in s1, but not s2", k)
				} else if n1, n2 := reflect.TypeOf(v).Name(), reflect.TypeOf(v2).Name(); n1 != n2 {
					t.Errorf("Types don't match: %s != %s", n1, n2)
				}
			}
		}
	}
}
