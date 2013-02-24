package java

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	testdata_path = "./testdata/"
)

// Diff cut'n'paste from http://golang.org/src/cmd/gofmt/gofmt.go
func diff(b1, b2 []byte) (data []byte, err error) {
	f1, err := ioutil.TempFile("", "parser")
	if err != nil {
		return
	}
	defer os.Remove(f1.Name())
	defer f1.Close()

	f2, err := ioutil.TempFile("", "parser")
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

func TestSpecificClasses(t *testing.T) {
	var (
		tests = make(map[Classname][]byte)
		err   error
	)

	classpath, err := DefaultClasspath()
	if err != nil {
		t.Fatal(err)
	}

	c, err := NewCompositeArchive(classpath)
	if err != nil {
		t.Log(err)
	}
	defer c.Close()

	testdata, err := os.Open(testdata_path)
	if err != nil {
		t.Fatalf("Failed to open %s: %s", testdata_path, err)
	}
	defer testdata.Close()

	fi, err := testdata.Readdir(0)
	if err != nil {
		t.Errorf("Failed to fully read %s: %s. Got %d useable entries.", testdata_path, err, len(fi))
	}

	for i := range fi {
		if strings.HasSuffix(fi[i].Name(), ".class") {
			if data, err := ioutil.ReadFile(testdata_path + fi[i].Name()); err != nil {
				t.Errorf("Error reading test data: %s", err)
			} else {
				tests[Filename(strings.Replace(fi[i].Name(), "_", "/", -1)).Classname()] = data
			}
		}
	}

	for k, v := range tests {
		t.Log("Testing", k)

		d, err := c.LoadClass(k)
		if err != nil {
			t.Errorf("Failed to read file contents: %s", err)
			continue
		}

		b := bytes.NewReader(d)
		c, err := NewClass(b)
		if err != nil {
			t.Errorf("Failed to create Class struct: %s", err)
			continue
		}

		if len(v) <= 1 {
			// Just if we want to add new tests, this will spit out the newly added
			// test data
			fn := testdata_path + strings.Replace(string(k.Filename()), "/", "_", -1)
			t.Logf("Creating new test data: %s", fn)
			if err := ioutil.WriteFile(fn, []byte(c.String()), 0666); err != nil {
				t.Errorf("Couldn't write test data to %s: %s", fn, err)
			}
			continue
		}

		if d, err := diff([]byte(v), []byte(c.String())); err != nil {
			t.Error(err)
		} else if len(d) != 0 {
			t.Error(string(d))
		}
	}
}
