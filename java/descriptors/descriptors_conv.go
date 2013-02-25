package descriptors

import (
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
