package net

import (
	"github.com/quarnster/completion/content"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

func TestCacheComplete(t *testing.T) {
	paths := DefaultPaths()
	if len(paths) == 0 {
		t.Skip("No default paths available")
	}

	tests := []string{"mscorlib.dll", "System.dll"}

	t.Log(paths)
	c := Cache{paths: paths}
	for _, test := range tests {
		if asm, err := c.Load(test); err != nil {
			t.Error(err)
		} else {
			t.Logf("Found %s (%s)", test, asm.Name())
		}
	}

	tests2 := []content.Type{
		content.Type{Name: content.FullyQualifiedName{Absolute: "net://type/System.String"}},
	}
	for _, test := range tests2 {
		if res, err := c.Complete(&test); err != nil {
			t.Error(err)
		} else {
			t.Log(res)
		}
	}
}

func mcs(args ...string) ([]byte, error) {
	cmd := exec.Command("mcs", args...)
	return cmd.CombinedOutput()
}

func TestCacheReload(t *testing.T) {
	if out, err := mcs(); len(out) == 0 {
		t.Skip("It does not appear that mcs is installed:", err)
	}
	source, err := ioutil.TempFile("", "Cache")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(source.Name())

	binary, err := ioutil.TempFile("", "Cache")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(binary.Name())

	if err := ioutil.WriteFile(source.Name(), []byte("class Test1 {}"), 0644); err != nil {
		t.Fatal(err)
	} else if out, err := mcs("-target:library", "-out:"+binary.Name(), source.Name()); err != nil {
		t.Fatal(string(out), err)
	}

	c := Cache{paths: []string{filepath.Dir(binary.Name())}}
	if _, err := c.Load(filepath.Base(binary.Name())); err != nil {
		t.Error(err)
	} else {
		ty1 := content.Type{Name: content.FullyQualifiedName{Absolute: "net://type/Test1"}}
		ty2 := content.Type{Name: content.FullyQualifiedName{Absolute: "net://type/Test2"}}
		if _, err := c.Complete(&ty1); err != nil {
			t.Fatal(err)
		} else if _, err := c.Complete(&ty2); err == nil {
			t.Fatal("Should not be possible to complete Test2")
		}

		if err := ioutil.WriteFile(source.Name(), []byte("class Test2 {}"), 0644); err != nil {
			t.Fatal(err)
		} else if out, err := mcs("-target:library", "-out:"+binary.Name(), source.Name()); err != nil {
			t.Fatal(string(out), err)
		}

		// Give it some time to reload
		time.Sleep(time.Millisecond * 100)

		if _, err := c.Complete(&ty1); err == nil {
			t.Fatal("Should not be possible to complete Test1")
		} else if _, err := c.Complete(&ty2); err != nil {
			t.Fatal(err)
		}
	}
}

func TestCacheLoadDep(t *testing.T) {
	paths := DefaultPaths()
	if len(paths) == 0 {
		t.Skip("No default paths available")
	}

	t.Log(paths)
	c := Cache{paths: paths}
	if asm, err := c.Load("System.dll"); err != nil {
		t.Error(err)
	} else {
		t.Logf("Found %s", asm.Name())
		time.Sleep(time.Millisecond * 100)
		if len(c.entries) <= 1 {
			t.Error("Dependencies not loaded:", c.entries)
		}
	}
}

func BenchmarkCacheComplete(b *testing.B) {
	b.StopTimer()

	paths := DefaultPaths()
	if len(paths) == 0 {
		b.Skip("No default paths available")
	}

	tests := []string{"mscorlib.dll", "System.dll"}

	c := Cache{paths: paths}
	for _, test := range tests {
		if _, err := c.Load(test); err != nil {
			b.Error(err)
		}
	}

	tests2 := []content.Type{
		content.Type{Name: content.FullyQualifiedName{Absolute: "net://type/System.String"}},
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range tests2 {
			if _, err := c.Complete(&test); err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkCacheFindType(b *testing.B) {
	b.StopTimer()
	paths := DefaultPaths()
	if len(paths) == 0 {
		b.Skip("No default paths available")
	}

	tests := []string{"mscorlib.dll", "System.dll"}

	c := Cache{paths: paths}
	for _, test := range tests {
		if _, err := c.Load(test); err != nil {
			b.Error(err)
		}
	}

	tests2 := []content.FullyQualifiedName{
		content.FullyQualifiedName{Absolute: "net://type/string"},
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range tests2 {
			if _, err := c.FindType(test); err != nil {
				b.Error(err)
			}
		}
	}
}
