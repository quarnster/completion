package content

// AKA generic parameter in some languages
type TemplateParameter struct {
	// TODO
}

type Type struct {
	Name           FullyQualifiedName `protocol:required`
	Specialization TemplateParameter  `protocol:repeated`
}

type Field struct {
	Name FullyQualifiedName `protocol:required`
	Type Type               `protocol:required`
}

type Method struct {
	Name           FullyQualifiedName `protocol:required`
	Returns        Field              `protocol:optional`
	Parameters     Field              `protocol:optional`
	Specialization TemplateParameter  `protocol:optional`
}

type CompletionResult struct {
	Types   []Type   `protocol:optional`
	Fields  []Field  `protocol:optional`
	Methods []Method `protocol:optional`
}

// TODO: maybe these are two different calls?
type CompletionRequest struct {
	// Either we want to complete this specific type/package/whatever
	Type FullyQualifiedName `protocol:optional`

	// Or we want to perform the complete operation at this specific location
	Location SourceLocation `protocol:optional`
}
