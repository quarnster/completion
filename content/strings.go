package content

import "fmt"

func (f File) String() string {
	return f.Name
}

func (s SourceLocation) String() string {
	return fmt.Sprintf("%s:%d:%d", s.File, s.Line, s.Column)
}

func (f FullyQualifiedName) String() string {
	if f.Absolute != "" {
		return f.Absolute
	}
	return f.Relative
}

func (t Type) String() string {
	return t.Name.String()
}

func (v Variable) String() string {
	return fmt.Sprintf("%s %s", v.Type, v.Name)
}

func (f Field) String() (ret string) {
	return ret + fmt.Sprintf("%s%s %s", f.Flags, f.Type, f.Name)
}

func (m Method) String() (ret string) {
	ret = fmt.Sprintf("%s", m.Flags)
	for i := range m.Returns {
		if i > 0 {
			ret += ", "
		}
		ret += fmt.Sprintf("%s", m.Returns[i])
	}
	ret += fmt.Sprintf(" %s(", m.Name)
	for i := range m.Parameters {
		if i > 0 {
			ret += ", "
		}
		ret += fmt.Sprintf("%s", m.Parameters[i])
	}
	ret += ")"
	return
}

func (a Flags) String() (ret string) {
	switch a & FLAG_ACC_MASK {
	case FLAG_ACC_PUBLIC:
		ret += "public "
	case FLAG_ACC_PRIVATE:
		ret += "private "
	case FLAG_ACC_PROTECTED:
		ret += "protected "
	}
	if a&FLAG_STATIC != 0 {
		ret += "static "
	}
	if a&FLAG_FINAL != 0 {
		ret += "final "
	}
	return
}
