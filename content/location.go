package content

type File struct {
	// The name of the file
	Name string `protocol:required`
	// The contents of the file (which might or might not correspond to the contents of the file saved on disc)
	Contents string `protocol:optional`
}

type SourceLocation struct {
	File   File `protocol:required`
	Line   uint `protocol:required`
	Column uint `protocol:required`
}

type FullyQualifiedName struct {
	// TODO
	// Relative would be "lastIndexOf" whereas the absolute would be something like
	// java/lang/String.class/lastIndexOf/(Ljava/lang/String;I)I or whatever
	// we decide on. We should probably define this better so that there's some standard in how
	// it's encoded across languages and to make it harder to write a path that's "wrong".
	Relative string `protocol:required`
	Absolute string `protocol:optional`
}
