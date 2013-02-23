package parse

type tokenType int

const (
	tokenImport tokenType = iota
	tokenAnnotation
	tokenClass
	tokenClassEnd
	tokenExtends
	tokenImplements
	tokenField
	tokenMethod
	tokenVar
	tokenVisitField
	tokenVisitMethod
	tokenVisitVar
)

var tokenString = map[tokenType]string{
	tokenImport:      "Import",
	tokenAnnotation:  "Annotation",
	tokenClass:       "Class",
	tokenClassEnd:    "ClassEnd",
	tokenExtends:     "Extends",
	tokenImplements:  "Implements",
	tokenField:       "Field",
	tokenMethod:      "Method",
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
