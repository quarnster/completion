package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
)

var ignore = regexp.MustCompile(`\.git|build|testdata|3rdparty|editor`)
var verbose bool

func createPluginArchive() {
	files := []string{
		"3rdparty/jsonrpc.py",
		"editor/sublime/plugin.py",
	}
	out := bytes.NewBuffer(nil)
	z := zip.NewWriter(out)
	for i := range files {
		if w, err := z.Create(files[i]); err != nil {
			panic(err)
		} else if d, err := ioutil.ReadFile("../" + files[i]); err != nil {
			panic(err)
		} else if n, err := w.Write(d); err != nil {
			panic(err)
		} else if n != len(d) {
			panic(fmt.Errorf("%d != %d", n != len(d)))
		}
	}
	if err := z.Close(); err != nil {
		panic(err)
	}
	b2 := bytes.NewBuffer(nil)
	b2.WriteString(`package editor

var archive_data = []byte{`)
	for i, b := range out.Bytes() {
		if i > 0 {
			b2.WriteRune(',')
		}
		b2.WriteString(fmt.Sprintf("%v", b))
	}
	b2.WriteString("}")
	if err := ioutil.WriteFile("../editor/archive_data.go", b2.Bytes(), 0644); err != nil {
		panic(err)
	}
}

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

func buildPeg(pegFile, ignore string) {
	c := exec.Command("./parser_exe", "-peg="+pegFile, "-notest", "-ignore", ignore, "-testfile", "none silly")
	if verbose {
		fmt.Println(c.Args)
	}
	if b, err := c.CombinedOutput(); err != nil {
		panic(fmt.Sprint(string(b), err))
	} else if len(b) != 0 {
		fmt.Println(string(b))
	}
	outname := filepath.Base(pegFile)
	outname = outname[:len(outname)-4]
	f1, err := os.Open(fmt.Sprintf("%s/%s.go", outname, outname))
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	f2, err := os.Create(pegFile[:len(pegFile)-4] + ".go")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	if verbose {
		fmt.Printf("Copying %s to %s\n", f1.Name(), f2.Name())
	}
	if _, err := io.Copy(f2, f1); err != nil {
		panic(err)
	}
}

func readthread(r io.Reader, out chan string) {
	buf := make([]byte, 2048)
	for {
		if n, err := r.Read(buf); err != nil && err != io.EOF {
			panic(err)
		} else if err == io.EOF {
			break
		} else {
			out <- string(buf[:n])
		}
	}
	close(out)
}

func main() {
	flag.BoolVar(&verbose, "v", verbose, "Verbose output")
	flag.Parse()
	createPluginArchive()
	var extension = ""
	if runtime.GOOS == "windows" {
		extension += ".exe"
	}
	c := exec.Command("go", "build", "-o", "parser_exe"+extension, "github.com/quarnster/parser/exe")
	if verbose {
		fmt.Println(c.Args)
	}
	if b, err := c.CombinedOutput(); err != nil {
		panic(fmt.Sprint(string(b), err))
	} else if len(b) != 0 {
		fmt.Println(string(b))
	}
	buildPeg("../java/signatures/signatures.peg", "TypeSignature,TypeArgument,Entry,FormalTypeParameter")
	buildPeg("../java/descriptors/descriptors.peg", "Entry,ComponentType,FieldType,ObjectType")
	buildPeg("../clang/parser/parser.peg", "Spacing,Pattern,Completion,Completions,ReturnType,Arguments,Argument,UnhandledStuff,KeyedStuff,Type,TemplateName,Function")
	buildPeg("../util/expression/expression.peg", "Spacing,Primary,Op,Expression,Grouping,BooleanOp")
	buildPeg("../util/expand_path/expand_path.peg", "Operation,File")
	buildPeg("../net/csharp/csharp.peg", "Complete,Junk,UsingDirectives,Primary,BOM,Spacing,Reference,Grouping,InnerScope,EndOfLine,Comment,LineComment,BlockComment,Loops,Code,SpacedIdentifier,SpacedAccess")
	tests := []string{"test"}
	if verbose {
		tests = append(tests, "-v")
	}
	tests = adddirs("github.com/quarnster/completion", "..", tests)
	c = exec.Command("go", tests...)
	r, err := c.StdoutPipe()
	if err != nil {
		panic(err)
	}
	r2, err := c.StderrPipe()
	if err != nil {
		panic(err)
	}
	if err := c.Start(); err != nil {
		panic(err)
	}
	sc := make(chan string)
	ec := make(chan string)
	go readthread(r, sc)
	go readthread(r2, ec)
	so, eo := true, true
	for so && eo {
		line := ""
		select {
		case line, so = <-sc:
		case line, eo = <-ec:
		}
		fmt.Print(line)
	}
}
