package java

import "strings"

func FileToClass(path string) string {
	return strings.Replace(path[:len(path)-len(".class")], "/", ".", -1)
}

func ClassToFile(class string) string {
	return strings.Replace(class, ".", "/", -1) + ".class"
}
