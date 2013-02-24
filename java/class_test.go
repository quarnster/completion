package java

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"testing"
)

const (
	testdata_path = "./testdata/"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func rtJar(t *testing.T) string {
	substr := []byte("rt.jar")

	// -XshowSettings still exits with `1` returning an error, so use CombinedOutput.
	out, _ := exec.Command("java", "-XshowSettings:properties").CombinedOutput()

	for _, line := range bytes.Split(out, []byte("\n")) {
		if bytes.Contains(line, substr) {
			return string(bytes.TrimSpace(line))
		}
	}

	t.Fatal("Failed to locate rt.jar")
	panic("unreachable")
}

func TestLoadAllClasses(t *testing.T) {
	if z, err := zip.OpenReader(rtJar(t)); err != nil {
		t.Fatal(err)
	} else {
		defer z.Close()
		var (
			inChan  = make(chan []byte, len(z.File))
			outChan = make(chan error, 32)
			wg      sync.WaitGroup
			once    sync.Once
		)
		worker := func() {
			wg.Add(1)
			for zf := range inChan {
				f := bytes.NewReader(zf)
				if c, err := NewClass(f); err != nil {
					outChan <- err
				} else {
					t.Log("class", String(c.Constant_pool, c.This_class))
				}
			}
			wg.Done()
			wg.Wait()
			once.Do(func() {
				close(outChan)
			})
		}
		for i := 0; i < runtime.NumCPU(); i++ {
			go worker()
		}
		for _, zf := range z.File {
			if strings.HasSuffix(zf.Name, ".class") {
				if f, err := zf.Open(); err != nil {
					t.Error(err)
				} else {
					defer f.Close()
					if d, err := ioutil.ReadAll(f); err != nil {
						t.Error(err)
					} else {
						inChan <- d
					}
				}
			}
		}
		close(inChan)
		for o := range outChan {
			if o != nil {
				t.Error(o)
			}
		}
	}
}

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
		tests     = make(map[string][]byte)
		tests_run int
		err       error
	)

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
				tests[strings.Replace(fi[i].Name(), "_", "/", -1)] = data
			}
		}
	}

	z, err := zip.OpenReader(rtJar(t))
	if err != nil {
		t.Fatalf("Failed to read rt.jar: %s", err)
	}
	defer z.Close()

	for _, zf := range z.File {

		v, ok := tests[zf.Name]
		if !ok {
			continue
		}
		tests_run++
		t.Log("Testing", zf.Name)

		f, err := zf.Open()
		if err != nil {
			t.Error("Failed to open zip file: %s", err)
			continue
		}
		defer f.Close()

		d, err := ioutil.ReadAll(f)
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
			fn := testdata_path + strings.Replace(zf.Name, "/", "_", -1)
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
	if len(tests) != tests_run {
		t.Errorf("Didn't run all tests. Expected %d tests run but only have %d", len(tests), tests_run)
	}
}
