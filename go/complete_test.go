package goo

import (
	"github.com/quarnster/completion/content"
	"reflect"
	"testing"
)

func TestGoo(t *testing.T) {
	var (
		g   Go
		cmp content.CompletionResult
	)
	if err := g.complete(reflect.TypeOf(&g), &cmp); err != nil {
		t.Error(err)
	} else {
		t.Log(cmp)
	}
}
