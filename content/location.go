package content

type (
	File struct {
		// The name of the file
		Name string `protocol:"required" json:",omitempty"`
		// The contents of the file (which might or might not correspond to the contents of the file saved on disc)
		Contents string `protocol:"optional" json:",omitempty"`
	}

	SourceLocation struct {
		File   File `protocol:"required" json:",omitempty"`
		Line   uint `protocol:"required" json:",omitempty"`
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
		// presented to the user.
		Relative string `protocol:"optional" json:",omitempty"`

		// The Absolute string contains all information needed
		// by a backend "driver" to
		Absolute string `protocol:"optional" json:",omitempty"`
	}
)
