package java

import "testing"

func TestClasses(t *testing.T) {
	if p, err := DefaultClasspath(); err != nil {
		t.Skip(err)
	} else {
		c, err := NewCompositeArchive(p)

		// yes, c not err as err can be filled out with errors opening individual classpaths
		if err != nil {
			t.Fatal(err)
		}
		defer c.Close()
		classes, err := c.Classes()
		if err != nil {
			t.Log("Error when getting all classes:", err)
		}
		if len(classes) < 1 {
			t.Error("No classes?")
		}
	}
}

func TestNewCompositeArchive(t *testing.T) {
	if c, err := NewCompositeArchive(nil); err != nil {
		t.Error(err)
	} else if c == nil {
		t.Error("Composite is nil")
	}
}

func BenchmarkCompositeArchiveLoadClass(b *testing.B) {
	b.StopTimer()
	if p, err := DefaultClasspath(); err != nil {
		b.Skip(err)
	} else {
		classes := []Classname{"java.lang.String", "javax.swing.JLabel"}
		c, err := NewCompositeArchive(p)
		// yes, c not err as err can be filled out with errors opening individual classpaths
		if err != nil {
			b.Fatal(err)
		}
		defer c.Close()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			for j := range classes {
				c.LoadClass(classes[j])
			}
		}
	}
}
