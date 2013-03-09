package content

import (
	"errors"
	"fmt"
	"strings"
)

type File struct {
	// The name of the file
	Name string `protocol:"required" json:",omitempty"`
	// The contents of the file (which might or might not correspond to the contents of the file saved on disc)
	Contents string `protocol:"optional" json:",omitempty"`
}

type SourceLocation struct {
	File   File `protocol:"required" json:",omitempty"`
	Line   uint `protocol:"required" json:",omitempty"`
	Column uint `protocol:"required" json:",omitempty"`
}

type FullyQualifiedName struct {
	Relative string `protocol:"required" json:",omitempty"`
	Absolute string `protocol:"optional" json:",omitempty"`
}

func (f *FullyQualifiedName) Validate() error {
	if strings.ContainsAny(string(f.Relative), "$.{}[]/*-+<>") {
		return errors.New(fmt.Sprintf("Relative name contains illegal characters: %s", f.Relative))
	}
	return nil
}
