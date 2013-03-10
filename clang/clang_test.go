package clang

import (
	"github.com/quarnster/completion/content"
	"testing"
)

func TestClang(t *testing.T) {
	if _, err := RunClang("-v"); err != nil {
		t.Skipf("Couldn't launch clang: %s", err)
	}
	loc := content.SourceLocation{}
	loc.Column = 1
	loc.Line = 10
	loc.File.Name = "testdata/hello.cpp"
	t.Log(CompleteAt(loc))
}
