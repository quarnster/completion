package content

import (
	"testing"
)

type test struct {
	data        interface{}
	expectError bool
}

var tests = []test{
	{File{}, true},
	{&File{}, true},
	{&File{Contents: "Something"}, true},
	{&File{Name: "Something"}, false},
	{&File{Name: "Something", Contents: "Something"}, false},
	{&SourceLocation{}, true},
	{&SourceLocation{Line: 1, Column: 1}, true},
	{&SourceLocation{File: File{Name: "Something"}, Line: 1, Column: 1}, false},
	{&FullyQualifiedName{}, true},
	{&FullyQualifiedName{Absolute: "Something"}, true},
	{&FullyQualifiedName{Relative: "Something"}, false},
}

func TestValidate(t *testing.T) {
	for _, test := range tests {
		err := Validate(test.data)
		t.Log(err)
		if err == nil && test.expectError {
			t.Error("Expected an error")
		} else if err != nil && !test.expectError {
			t.Errorf("Didn't expected an error: %s", err)
		}
	}
}
