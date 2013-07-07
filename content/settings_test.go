package content

import (
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

func TestParent(t *testing.T) {
	s1 := &Settings{data: map[string]interface{}{
		"hello": "world",
		"yes":   true,
		"no":    3,
	}}
	s2 := &Settings{data: map[string]interface{}{
		"no":  false,
		"yes": false,
		"new": 14,
	},
		parent: s1,
	}
	exp := map[string]interface{}{
		"hello": "world",
		"yes":   false,
		"no":    false,
		"new":   14,
	}
	for k, v := range exp {
		if a, b := v, s2.Get(k); a != b {
			t.Errorf("%s != %s", a, b)
		}
	}
	s2.Set("hello", "you")
	if v := s2.Get("hello").(string); v != "you" {
		t.Errorf("Expected 'you' got: %s", v)
	}
	if v := s1.Get("hello").(string); v != "world" {
		t.Errorf("Expected 'world' got: %s", v)
	}

}
