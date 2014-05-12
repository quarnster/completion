package simplify

import (
	"github.com/quarnster/parser"
	"regexp"
)

var (
	replacements = []struct {
		re  *regexp.Regexp
		rep string
	}{
		{regexp.MustCompile(`;[^=;{}]+;`), ";"},
		{regexp.MustCompile(`\s*\{\s*\}`), ";"},
		{regexp.MustCompile(`\s(if|else|while|catch)\s*([^{;]+;|;)`), ""},
		{regexp.MustCompile(`\sfor\s?\([^)]*\);`), ""},
		{regexp.MustCompile(`\{\s*`), "{"},
		{regexp.MustCompile(`\s+`), " "},
		{regexp.MustCompile(`(;|{|})\s?`), "$1\n"},
		{regexp.MustCompile(`\s{`), "\n{"},
		{regexp.MustCompile(`^\s+`), ""},
	}
)

// Performs a series of transforms on the input code string
// trying to make it as simple as possible to parse.
//
// The output is a compact representation of the original
// source code input, where each line in the output
// ends with either a ';', a '{', or a '}', and only one
// of these characters may be present in each line.
func Simplify(data string) string {
	for _, op := range replacements {
		data = op.re.ReplaceAllString(data, op.rep)
	}
	Simplified(data)
	return data
}

func Simplified(data string) (*parser.Node, error) {
	var s SIMPLIFY
	s.SetData(data)
	if !s.Simplified() {
		return nil, s.Error()
	}
	return s.RootNode(), nil
}
