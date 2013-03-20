package descriptors

import (
	"github.com/quarnster/completion/content"
	"github.com/quarnster/parser"
	"strings"
)

// http://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.3.2-210
var lut = map[string]string{
	"B": "byte",
	"C": "char",
	"D": "double",
	"F": "float",
	"I": "int",
	"J": "long",
	"S": "short",
	"Z": "boolean",
	"V": "void"}

func ToContentFQN(str string) (ret content.FullyQualifiedName) {
	ret.Relative = str
	if i := strings.LastIndex(ret.Relative, "/"); i > 0 {
		ret.Relative = ret.Relative[i+1:]
	}
	ret.Absolute = "java://type/" + str
	return
}

func ToContentType(node *parser.Node) (ret content.Type) {
	if node.Name == "ArrayType" {
		ret.Flags |= content.FLAG_TYPE_ARRAY
		ret.Specialization = append(ret.Specialization, ToContentType(node.Children[0]))
		return
	}
	switch node.Name {
	case "BaseType":
		fallthrough
	case "VoidDescriptor":
		ret.Name.Relative = lut[node.Data()]
		ret.Name.Absolute = ret.Name.Relative
	case "Classname":
		ret.Name = ToContentFQN(node.Data())
	default:
		return ToContentType(node.Children[0])
	}

	return
}

func ToContentField(node *parser.Node) (ret content.Field) {
	ret.Type = ToContentType(node.Children[0])
	return
}

func ToContentMethod(node *parser.Node) (ret content.Method) {
	ri := len(node.Children) - 1
	ret.Returns = append(ret.Returns, content.Variable{Type: ToContentType(node.Children[ri])})
	for i := range node.Children[:ri] {
		ret.Parameters = append(ret.Parameters, content.Variable{Type: ToContentType(node.Children[i])})
	}
	return
}

func convert(node *parser.Node) string {
	ret := ""
	switch node.Name {
	case "MethodDescriptor":
		ri := len(node.Children) - 1
		ret += convert(node.Children[ri]) + " ("
		for i := range node.Children[:ri] {
			if i > 0 {
				ret += ", "
			}
			ret += convert(node.Children[i])
		}
		ret += ")"
		return ret
	case "BaseType":
		fallthrough
	case "VoidDescriptor":
		return lut[node.Data()]
	case "Classname":
		return strings.Replace(node.Data(), "/", ".", -1)
	}
	for i := range node.Children {
		if ret != "" {
			ret += " "
		}
		ret += convert(node.Children[i])
	}
	if node.Name == "ArrayType" {
		ret += "[]"
	}
	return ret
}
