package content

import (
	"encoding/json"
	"testing"
)

func TestSetGet(t *testing.T) {
	s := NewSettings()
	s.Set("hello", "world")
	s.Set("yes", true)
	s.Set("no", false)
	s.Set("cool", -1337)

	if v, ok := s.Get("hello").(string); !ok || v != "world" {
		t.Error("Should be ok")
	}
	if _, ok := s.Get("hello").(int); ok {
		t.Error("Should not be ok")
	}
	if _, ok := s.Get("notset").(int); ok {
		t.Error("Should not be ok")
	}
	if _, ok := s.Get("notset").(string); ok {
		t.Error("Should not be ok")
	}
}

func TestClone(t *testing.T) {
	s := NewSettings()
	s.Set("hello", "world")
	s.Set("yes", true)
	s2 := s.Clone()
	s2.Set("no", false)
	for k := range s.data {
		if a, b := s.Get(k), s2.Get(k); a != b {
			t.Errorf("%s != %s", a, b)
		}
	}
	if a, b := len(s.data), len(s2.data); a == b {
		t.Errorf("Expected different data lengths: %d == %d", a, b)
	}
}

func TestMerge(t *testing.T) {
	s1 := &Settings{map[string]interface{}{
		"hello": "world",
		"yes":   true,
		"no":    3,
	}}
	s2 := &Settings{map[string]interface{}{
		"no":  false,
		"yes": false,
		"new": 14,
	}}
	exp := map[string]interface{}{
		"hello": "world",
		"yes":   false,
		"no":    false,
		"new":   14,
	}
	s1.Merge(s2)
	if a, b := len(s1.data), len(exp); a != b {
		t.Errorf("%d != %d", a, b)
	}
	for k, v := range exp {
		if a, b := v, s1.Get(k); a != b {
			t.Errorf("%s != %s", a, b)
		}
	}

}

func TestJson(t *testing.T) {
	s1 := &Settings{map[string]interface{}{
		"hello": "world",
		"yes":   true,
		"no":    3,
		"complex": CompletionResult{Fields: []Field{
			{Variable: Variable{Name: FullyQualifiedName{Relative: "Hello"}}},
		}},
	}}
	if d, err := json.MarshalIndent(&s1, "", "\t"); err != nil {
		t.Error(err)
	} else {
		t.Log(string(d))
		s2 := &Settings{}
		if err := json.Unmarshal(d, &s2); err != nil {
			t.Error(err)
		} else {
			t.Log(s1)
			t.Log(s2)
		}
	}
}
