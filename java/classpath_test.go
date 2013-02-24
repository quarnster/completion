package java

import "testing"

func TestDefaultClasspath(t *testing.T) {
	if p, err := DefaultClasspath(); err != nil {
		t.Error(err)
	} else if len(p) < 1 {
		t.Error("Empty default classpath? How is that possible?")
	} else {
		t.Log("Default classpath: ", p)
	}
}

func TestClasses(t *testing.T) {
	if p, err := DefaultClasspath(); err != nil {
		t.Error(err)
	} else {
		classes, err := Classes(p)
		if err != nil {
			t.Log("Error when getting all classes:", err)
		}
		if len(classes) < 1 {
			t.Error("No classes?")
		}
	}
}

func TestLoadClassInMap(t *testing.T) {
	if p, err := DefaultClasspath(); err != nil {
		t.Error(err)
	} else if m := ClasspathMap(p); len(m) < 1 {
		t.Error("No classes?")
	}
}
