package expression

import (
	"fmt"
	"github.com/quarnster/parser"
	"reflect"
	"strconv"
)

func Eval(v *reflect.Value, node *parser.Node) (int, error) {
	switch node.Name {
	case "EXPRESSION":
		if l := len(node.Children); l != 2 {
			return 0, fmt.Errorf("Unexpected child length: %d, %s", l, node)
		}
		return Eval(v, node.Children[0])
	case "Identifier":
		if f := v.FieldByName(node.Data()); !f.IsValid() {
			return 0, fmt.Errorf("No field by name %s in struct %s", node.Data(), v)
		} else {
			switch f.Kind() {
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				return int(f.Uint()), nil
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return int(f.Int()), nil
			default:
				return 0, fmt.Errorf("Unexpected identifier kind: %v %v", f, f.Kind())
			}
		}
	case "Constant":
		i, err := strconv.ParseInt(node.Data(), 0, 32)
		return int(i), err
	case "EndOfFile":
		return 0, nil
	default:
		if l := len(node.Children); l != 2 {
			return 0, fmt.Errorf("Unexpected child length: %d, %s", l, node)
		}
		if a, err := Eval(v, node.Children[0]); err != nil {
			return 0, err
		} else if b, err := Eval(v, node.Children[1]); err != nil {
			return 0, err
		} else {
			switch node.Name {
			case "Eq":
				if a == b {
					return 1, nil
				} else {
					return 0, nil
				}
			case "Lt":
				if a < b {
					return 1, nil
				} else {
					return 0, nil
				}
			case "Gt":
				if a > b {
					return 1, nil
				} else {
					return 0, nil
				}
			case "Le":
				if a <= b {
					return 1, nil
				} else {
					return 0, nil
				}
			case "Ge":
				if a >= b {
					return 1, nil
				} else {
					return 0, nil
				}
			case "Add":
				return a + b, nil
			case "Sub":
				return a - b, nil
			case "Mul":
				return a * b, nil
			case "ShiftLeft":
				return int(uint(a) << uint(b)), nil
			case "ShiftRight":
				return int(uint(a) >> uint(b)), nil
			case "Mask":
				return a & b, nil
			case "AndNot":
				return a &^ b, nil
			default:
				return 0, fmt.Errorf("Unimplemented operation: %s", node.Name)
			}
		}
	}
}
