package java

import (
	"github.com/quarnster/completion/content"
	"testing"
)

func TestCompletionIntent(t *testing.T) {
	it := content.NewIntent("completion.complete.fqn")
	it.Data.Set("fqn", content.FullyQualifiedName{Absolute: "java://type/java/util/jar/JarEntry"})
	c := CompletionHandler{}
	if !c.CanHandle(&it) {
		t.Fatal("Can't handle the intent")
	} else {
		r := c.Handle(&it)
		if err, ok := r.Data.Get("error").(string); ok {
			t.Error(err)
		} else if cmp, ok := r.Data.Get("completions").(content.CompletionResult); !ok {
			t.Error("No completions entry")
		} else {
			t.Log(cmp)
		}
	}
}
