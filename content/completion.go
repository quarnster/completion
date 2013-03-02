package content

// AKA generic parameter in some languages
type TemplateParameter struct {
	// TODO
}

type Type struct {
	Name           FullyQualifiedName  `protocol:"required" json:",omitempty"`
	Specialization []TemplateParameter `protocol:"optional" json:",omitempty"`
}

type Field struct {
	Name FullyQualifiedName `protocol:"required" json:",omitempty"`
	Type Type               `protocol:"required" json:",omitempty"`
}

type Method struct {
	Name           FullyQualifiedName  `protocol:"required" json:",omitempty"`
	Returns        []Field             `protocol:"optional" json:",omitempty"`
	Parameters     []Field             `protocol:"optional" json:",omitempty"`
	Specialization []TemplateParameter `protocol:"optional" json:",omitempty"`
}

type CompletionResult struct {
	Types   []Type   `protocol:"optional" json:",omitempty"`
	Fields  []Field  `protocol:"optional" json:",omitempty"`
	Methods []Method `protocol:"optional" json:",omitempty"`
}

// TODO: maybe these are two different calls?
type CompletionRequest struct {
	// Either we want to complete this specific type/package/whatever
	Type FullyQualifiedName `protocol:"optional" json:",omitempty"`

	// Or we want to perform the complete operation at this specific location
	Location SourceLocation `protocol:"optional" json:",omitempty"`
}
