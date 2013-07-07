package content

import (
	"io/ioutil"
	"strings"
)

type (
	File struct {
		// The name of the file
		Name string `protocol:"required" json:",omitempty"`
		// The contents of the file (which might or might not correspond to the contents of the file saved on disc)
		Contents string `protocol:"optional" json:",omitempty"`
	}

	SourceLocation struct {
		File File `protocol:"required" json:",omitempty"`
		// The line number starting at line 1
		Line uint `protocol:"required" json:",omitempty"`
		// The column number starting at line 1
		Column uint `protocol:"required" json:",omitempty"`
	}

	FullyQualifiedName struct {
		// While strongly encouraged to be provided, the relative name
		// is optional rather than required by the protocol.
		// This is because the required constraint cannot be satisfied
		// in all different languages. For example C does not have named
		// returns, and python does not type a functions parameters.
		//
		// When available, the Relative name would be the identifier
		// presented to the user, for example "String" rather than
		// "java.lang.String".
		Relative string `protocol:"optional" json:",omitempty"`

		// The Absolute string contains all information needed
		// by a backend "driver" to be able to perform a "Complete"
		// (rather than "CompleteAt") operation and is usually
		// treated as an internal driver specific identifier.
		//
		// TODO():	What about presenting "packages" or "namespaces" to the user?
		//			That's not always redundant information.
		// 			We may want to change this field to be the "java.lang.String",
		//			or perhaps we have a "Path" field containing "java.lang".
		//			The information needed by a driver for non-"At" operations
		//			(i.e going to the documentation/implementation of a type)
		// 			may be specified in another field then.
		Absolute string `protocol:"optional" json:",omitempty"`
	}
	Offset int
)

func (f *File) Load() error {
	if f.Contents != "" {
		return nil
	}
	if d, err := ioutil.ReadFile(f.Name); err != nil {
		return err
	} else {
		f.Contents = string(d)
	}
	return nil
}

// Returns the line at the given offset
func (f *File) Line(offset Offset) string {
	const linebreaks = "\n\r"
	if offset > Offset(len(f.Contents)) || offset < 0 {
		return ""
	}
	end := strings.IndexAny(f.Contents[offset:], linebreaks)
	if end == -1 {
		end = len(f.Contents)
	} else {
		end += int(offset)
	}
	start := strings.LastIndexAny(f.Contents[:end], linebreaks)
	return f.Contents[start+1 : end]
}

// Returns the 0 based offset that the SourceLocation represents.
func (s *SourceLocation) Offset() Offset {
	line, col := uint(1), uint(1)
	for i, c := range s.File.Contents {
		if line == s.Line && col == s.Column {
			return Offset(i)
		}
		// TODO(): do we need to consider \r too?
		if c == '\n' {
			line++
			col = 1
		} else {
			col++
		}
	}
	return -1
}
