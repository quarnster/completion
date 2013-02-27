package util

import (
	"io/ioutil"
	"os"
	"os/Exec"
)

// Diff cut'n'paste from http://golang.org/src/cmd/gofmt/gofmt.go
func Diff(b1, b2 []byte) (data []byte, err error) {
	f1, err := ioutil.TempFile("", "completion_diff")
	if err != nil {
		return
	}
	defer os.Remove(f1.Name())
	defer f1.Close()

	f2, err := ioutil.TempFile("", "completion_diff")
	if err != nil {
		return
	}
	defer os.Remove(f2.Name())
	defer f2.Close()

	f1.Write(b1)
	f2.Write(b2)

	data, err = exec.Command("diff", "-u", f1.Name(), f2.Name()).CombinedOutput()
	if len(data) > 0 {
		// diff exits with a non-zero status when the files don't match.
		// Ignore that failure as long as we get output.
		err = nil
	}
	return

}
