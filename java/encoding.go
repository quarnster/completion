package java

import "strings"

type (
	Filename  string
	Classname string
)

func (fn Filename) Classname() Classname {
	path := string(fn)
	return Classname(strings.Replace(path[:len(path)-len(".class")], "/", ".", -1))
}

func (cn Classname) Filename() Filename {
	class := string(cn)
	return Filename(strings.Replace(class, ".", "/", -1) + ".class")
}
