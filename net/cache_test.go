package net

import (
	"github.com/quarnster/completion/content"
	"testing"
)

func TestCache(t *testing.T) {
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
		content.Type{Name: content.FullyQualifiedName{Absolute: "System.String"}},
	}
	for _, test := range tests2 {
		if res, err := c.Complete(&test); err != nil {
			t.Error(err)
		} else {
			t.Log(res)
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
		content.Type{Name: content.FullyQualifiedName{Absolute: "System.String"}},
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
