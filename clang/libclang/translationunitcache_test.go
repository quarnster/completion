// +build !nolibclang

package libclang

import (
	"github.com/quarnster/completion/content"
	"testing"
)

func TestTranslationUnitCache(t *testing.T) {
	cache := NewTranslationUnitCache()
	if tu := cache.GetTranslationUnit("../testdata/hello.cpp", []string{"-x", "c++"}, "", nil); tu == nil {
		t.Error("Didn't get a translation unit")
	}
	var (
		args = content.CompleteAtArgs{
			Location: content.SourceLocation{
				File:   content.File{Name: "../testdata/hello.cpp"},
				Line:   4,
				Column: 1,
			},
		}
		res content.CompletionResult
	)
	if err := cache.CompleteAt(&args, &res); err != nil {
		t.Errorf("Error completing: %v", err)
	} else {
		t.Logf("%+v", res)
	}
}
