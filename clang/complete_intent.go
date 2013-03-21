package clang

import (
	"github.com/quarnster/completion/content"
	"regexp"
)

var extre = regexp.MustCompile(`\.(h|hpp|c|cc|cpp|m|mm)$`)

type CompletionHandler struct {
}

func (c *CompletionHandler) Handle(it *content.Intent) *content.Response {
	loc := it.Data.Get("loc").(content.SourceLocation)
	args, _ := it.Settings().Get("compiler_flags").([]string)
	if cmp, err := CompleteAt(args, loc); err != nil {
		return content.NewErrorResponse(err.Error())
	} else {
		r := content.NewResponse()
		r.Data.Set("completions", cmp)
		return &r
	}
}

func (c *CompletionHandler) CanHandle(it *content.Intent) bool {
	if it.Operation != content.CompleteSourceLocation {
		return false
	} else if loc, ok := it.Data.Get("loc").(content.SourceLocation); !ok {
		return false
	} else {
		return extre.MatchString(loc.File.Name)
	}
}
