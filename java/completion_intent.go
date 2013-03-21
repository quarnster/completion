package java

import (
	"bytes"
	"code.google.com/p/log4go"
	"github.com/quarnster/completion/content"
	"strings"
)

type CompletionHandler struct {
}

func fqnToClassname(fqn content.FullyQualifiedName) Classname {
	return Classname(strings.Replace(fqn.Absolute[len("java://type/"):], "/", ".", -1))
}

func (c *CompletionHandler) Complete(archive Archive, fqn content.FullyQualifiedName) (cmp content.CompletionResult, err error) {
	if data, err := archive.LoadClass(fqnToClassname(fqn)); err != nil {
		return cmp, err
	} else if class, err := NewClass(bytes.NewReader(data)); err != nil {
		return cmp, err
	} else if ct, err := class.ToContentType(); err != nil {
		return cmp, err
	} else {
		// TODO: inherited fields and methods
		cmp.Fields = ct.Fields
		cmp.Types = ct.Types
		cmp.Methods = ct.Methods
	}
	return
}

func (c *CompletionHandler) Handle(it *content.Intent) *content.Response {
	cp, err := DefaultClasspath()
	if err != nil {
		// We don't really care about not being able to get the default classpath as it could be provided manually by the user
		log4go.Warn("Couldn't get the default classpath: %s", err)
	}
	settings := it.Settings()
	if cp2, ok := settings.Get("classpath").([]string); ok {
		// TODO: do we ever want to override rather than append to the classpath?
		cp = append(cp, cp2...)
	}
	// TODO: this should probably be cached in the session...
	if archive, err := NewCompositeArchive(cp); err != nil {
		return content.NewErrorResponse(err.Error())
	} else {
		if cmp, err := c.Complete(archive, it.Data.Get("fqn").(content.FullyQualifiedName)); err != nil {
			return content.NewErrorResponse(err.Error())
		} else {
			r := content.NewResponse()
			r.Data.Set("completions", cmp)
			return &r
		}
	}
}

func (c *CompletionHandler) CanHandle(it *content.Intent) bool {
	// TODO
	if it.Operation != "completion.complete.fqn" {
		return false
	}
	fqn, ok := it.Data.Get("fqn").(content.FullyQualifiedName)
	return ok && strings.HasPrefix(fqn.Absolute, "java://type/")
}
