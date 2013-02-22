package java

import (
	"os"
	"testing"
)

func TestNewClass(t *testing.T) {
	if f, err := os.Open("/Users/quarnster/Library/Application Support/Sublime Text 3/Packages/SublimeJava/SublimeJava.class"); err != nil {
		t.Fatal(err)
	} else {
		defer f.Close()
		if _, err := NewClass(f); err != nil {
			t.Error(err)
		}
	}
}
