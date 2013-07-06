// +build !nolibclang

package libclang

import (
	"github.com/quarnster/completion/content"
	"testing"
)

func TestComplete(t *testing.T) {
	cache := Clang{*NewTranslationUnitCache()}
	var (
		args = content.CompleteAtArgs{
			Location: content.SourceLocation{
				File:   content.File{Name: "hello.cpp", Contents: "#include <stdio.h>\nint main(){ printf(); }"},
				Line:   2,
				Column: 12,
			},
		}
		res content.CompletionResult
	)
	if err := cache.CompleteAt(&args, &res); err != nil {
		t.Errorf("Error completing: %v", err)
	} else if l := len(res.Methods); l < 10 {
		t.Errorf("Expected more methods than %d", l)
	} else {
		t.Logf("%+v", res)
	}
}
