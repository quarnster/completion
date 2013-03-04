package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
)

var ignore = regexp.MustCompile(`\.git|build|testdata`)

func adddirs(pkg, path string, dirs []string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if fi, err := f.Readdir(0); err != nil {
		panic(err)
	} else {
		for _, f := range fi {
			if !f.IsDir() || ignore.MatchString(f.Name()) {
				continue
			}
			pkg2 := fmt.Sprintf("%s/%s", pkg, f.Name())
			dirs = append(dirs, pkg2)
			dirs = adddirs(pkg2, fmt.Sprintf("%s/%s", path, f.Name()), dirs)
		}
	}
	return dirs
}

func main() {
	c := exec.Command("go", "build", "-o", "parser_exe", "github.com/quarnster/parser/exe")
	if b, err := c.CombinedOutput(); err != nil {
		panic(err)
	} else if len(b) != 0 {
		fmt.Println(string(b))
	}
	c = exec.Command("./parser_exe", "-peg=../java/descriptors/descriptors.peg", "-notest", "-ignore", "Entry,ComponentType,FieldType,ObjectType", "-testfile", "none silly")
	if b, err := c.CombinedOutput(); err != nil {
		panic(err)
	} else if len(b) != 0 {
		fmt.Println(string(b))
	}
	f1, err := os.Open("descriptors/descriptors.go")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	f2, err := os.Create("../java/descriptors/descriptors.go")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	if _, err := io.Copy(f2, f1); err != nil {
		panic(err)
	}
	tests := []string{"test"}
	tests = adddirs("github.com/quarnster/completion", "..", tests)
	c = exec.Command("go", tests...)
	r, err := c.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err := c.Start(); err != nil {
		panic(err)
	}
	buf := make([]byte, 8)
	for {
		if n, err := r.Read(buf); err != nil && err != io.EOF {
			panic(err)
		} else if err == io.EOF {
			break
		} else {
			fmt.Print(string(buf[:n]))
		}
	}
}
