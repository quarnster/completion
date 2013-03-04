package java

import "testing"

func TestDefaultClasspath(t *testing.T) {
	if p, err := DefaultClasspath(); err != nil {
		t.Skip(err)
	} else if len(p) < 1 {
		t.Error("Empty default classpath? How is that possible?")
	} else {
		t.Log("Default classpath: ", p)
	}
}
