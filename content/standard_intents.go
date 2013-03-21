package content

var (
	// The CompleteFullyQualifiedName operation expects a "fqn" key in the intent's data field
	// pointing to the FullyQualifiedName of the type we want to complete
	CompleteFullyQualifiedName = "completion.complete.fqn"

	// The CompleteSourceLocation operation expects a "loc" key in the intent's data field
	// pointing to the SourceLocation we want to complete
	CompleteSourceLocation = "complete.loc"
)
