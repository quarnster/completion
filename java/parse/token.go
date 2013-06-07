package parse

type tokenType int

const (
	tokenPackage tokenType = iota
	tokenImport
	tokenAnnotation
	tokenClass
	tokenClassEnd
	tokenExtends
	tokenImplements
	tokenInit
	tokenField
	tokenMethod
	tokenParameter
	tokenVar
	tokenVisitField
	tokenVisitMethod
	tokenVisitVar
)

var tokenString = map[tokenType]string{
	tokenPackage:     "Package",
	tokenImport:      "Import",
	tokenAnnotation:  "Annotation",
	tokenClass:       "Class",
	tokenClassEnd:    "ClassEnd",
	tokenExtends:     "Extends",
	tokenImplements:  "Implements",
	tokenInit:        "Initializer",
	tokenField:       "Field",
	tokenMethod:      "Method",
	tokenParameter:   "Parameter",
	tokenVar:         "Var",
	tokenVisitField:  "VisitField",
	tokenVisitMethod: "VisitMethod",
	tokenVisitVar:    "VisitVar",
}

type token struct {
	typ   tokenType
	start int
	end   int
}
