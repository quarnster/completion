package java

import "testing"

func TestClasses(t *testing.T) {
	if p, err := DefaultClasspath(); err != nil {
		t.Error(err)
	} else {
		c, err := NewCompositeArchive(p)
		if err != nil {
			t.Log(err)
		}
		classes, err := c.Classes()
		if err != nil {
			t.Log("Error when getting all classes:", err)
		}
		if len(classes) < 1 {
			t.Error("No classes?")
		}
	}
}
