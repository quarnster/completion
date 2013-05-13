package java

import (
	"github.com/quarnster/completion/content"
	"testing"
)

func TestJavaComplete(t *testing.T) {
	if p, err := DefaultClasspath(); err != nil || len(p) == 0 {
		t.Skipf("Default java classpath not available: %v, %d", err, p)
	}

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
