package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
)

var ignore = regexp.MustCompile(`\.git|build|testdata|3rdparty|editor|libclang|go`)
var (
	verbose bool
	test    bool
)

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

// TODO(q): skips libclang for now
func main() {
	flag.BoolVar(&verbose, "v", verbose, "Verbose output")
	flag.Parse()

	tests := []string{"test", "-tags", "nolibclang"}
	if verbose {
		tests = append(tests, "-v")
	}
	tests = append(tests, "github.com/quarnster/completion")
	tests = adddirs("github.com/quarnster/completion", "..", tests)
	c := exec.Command("go", tests...)
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
