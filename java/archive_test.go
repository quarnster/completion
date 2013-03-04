package java

import "testing"

func TestClasses(t *testing.T) {
	if p, err := DefaultClasspath(); err != nil {
		t.Skip(err)
	} else {
		c, err := NewCompositeArchive(p)

		// yes, c not err as err can be filled out with errors opening individual classpaths
		if c == nil {
			t.Fatal(err)
		} else if err != nil {
			t.Log(err)
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

func BenchmarkCompositeArchiveLoadClass(b *testing.B) {
	b.StopTimer()
	if p, err := DefaultClasspath(); err != nil {
		b.Skip(err)
	} else {
		classes := []Classname{"java.lang.String", "javax.swing.JLabel"}
		c, err := NewCompositeArchive(p)
		// yes, c not err as err can be filled out with errors opening individual classpaths
		if c == nil {
			b.Fatal(err)
		} else if err != nil {
			b.Log(err)
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
