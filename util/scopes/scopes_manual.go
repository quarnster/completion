package scopes

import (
	"github.com/quarnster/completion/content"
	"github.com/quarnster/parser"
)

// Returns the RangeSet visible from the given source code location.
// Note that it only uses '{' and '}' to deduce scopes, and further
// processing is likely needed depending on language specifics.
func Visibility(loc content.SourceLocation) (ret parser.RangeSet) {
	ret.Add(parser.Range{0, len(loc.File.Contents)})
	var s SCOPES
	s.Parse(loc.File.Contents)
	var rec func(node *parser.Node)
	pos := int(loc.Offset())
	rec = func(node *parser.Node) {
		for _, child := range node.Children {
			if !child.Range.Inside(pos) {
				ret = ret.Cut(child.Range)
			} else {
				rec(child)
			}
		}
	}
	rec(s.RootNode())
	return
}
