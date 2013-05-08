package java

import (
	"github.com/quarnster/completion/content"
	"testing"
)

func TestJavaComplete(t *testing.T) {
	var a content.CompleteArgs
	var cmp content.CompletionResult
	a.Location.Absolute = "java://type/java/util/jar/JarEntry"
	c := Java{}
	if err := c.Complete(&a, &cmp); err != nil {
		t.Error(err)
	} else {
		t.Log(cmp)
	}
}
