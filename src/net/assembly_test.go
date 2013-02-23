package net

import (
	"os"
	"testing"
)

const System_dll = "/Library/Frameworks/Mono.framework/Libraries/mono/4.0/System.dll"

func TestLoadAssembly(t *testing.T) {
	if f, err := os.Open(System_dll); err != nil {
		t.Error(err)
	} else {
		defer f.Close()
		if _, err := LoadAssembly(f); err != nil {
			t.Error(err)
		}
	}
}
