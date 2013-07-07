package content

import (
	"testing"
)

func TestLine(t *testing.T) {
	tests := []struct {
		sl  SourceLocation
		exp string
	}{
		{
			SourceLocation{
				File:   File{Name: "location_test.go"},
				Line:   7,
				Column: 10,
			},
			"func TestLine(t *testing.T) {",
		},
		{
			SourceLocation{
				File:   File{Name: "location_test.go"},
				Line:   1,
				Column: 1,
			},
			"package content",
		},
	}
	for _, test := range tests {
		if err := test.sl.File.Load(); err != nil {
			t.Fatal(err)
		} else if l := test.sl.File.Line(test.sl.Offset()); l != test.exp {
			t.Errorf("Expected \"%s\", but got: %s", test.exp, l)
		}
	}
}
