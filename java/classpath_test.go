package java

import (
	"reflect"
	"testing"
)

func TestDefaultClasspath(t *testing.T) {
	if p, err := DefaultClasspath(); err != nil {
		t.Skip(err)
	} else if len(p) < 1 {
		t.Error("Empty default classpath? How is that possible?")
	} else {
		t.Log("Default classpath: ", p)
	}
}

func TestClasspathReading(t *testing.T) {
	var cases = []struct {
		test     string
		expected []string
	}{
		{
			"\n\n  some.jar \nwith spaces.jar \n\nmore.jar",
			[]string{"some.jar", "with spaces.jar", "more.jar"},
		},
		{"", []string{}},
	}
	for _, c := range cases {
		cp, err := classpathFromProps(map[string]string{"sun.boot.class.path": c.test})
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(c.expected, cp) {
			t.Errorf("Expected %v, got %v", c.expected, cp)
		}
	}
}
