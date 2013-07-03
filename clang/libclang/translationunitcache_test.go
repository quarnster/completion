// +build !nolibclang

package libclang

import (
	"testing"
)

func TestTranslationUnitCache(t *testing.T) {
	cache := NewTranslationUnitCache()
	if tu := cache.GetTranslationUnit("../testdata/hello.cpp", []string{"-x", "c++"}, "", nil); tu == nil {
		t.Error("Didn't get a translation unit")
	}
}
