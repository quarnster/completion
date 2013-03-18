package java

import (
	"bytes"
	"fmt"
	"github.com/quarnster/completion/util"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	testdata_path = "./testdata/"
)

func TestSpecificClasses(t *testing.T) {
	var (
		tests = make(map[Classname]string)
		err   error
	)

	classpath, err := DefaultClasspath()
	if err != nil {
		t.Skip(err)
	}

	c, err := NewCompositeArchive(classpath)
	if c == nil {
		t.Fatal(err)
	} else if err != nil {
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
				tests[Filename(strings.Replace(fi[i].Name(), "_", "/", -1)).Classname()] = string(data)
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

		ret := c.String()
		if ct, err := c.ToContentType(); err != nil {
			t.Error(err)
		} else {
			ret += fmt.Sprintf("%s\n", ct)
		}

		if len(v) <= 1 {
			// Just if we want to add new tests, this will spit out the newly added
			// test data
			fn := testdata_path + strings.Replace(string(k.Filename()), "/", "_", -1)
			t.Logf("Creating new test data: %s", fn)
			if err := ioutil.WriteFile(fn, []byte(ret), 0666); err != nil {
				t.Errorf("Couldn't write test data to %s: %s", fn, err)
			}
			continue
		}

		if d := util.Diff(v, ret); len(d) != 0 {
			t.Error(d)
		}
	}
}

func BenchmarkLoadJavaLangString(b *testing.B) {
	b.StopTimer()
	classpath, err := DefaultClasspath()
	if err != nil {
		b.Fatal(err)
	}
	c, err := NewCompositeArchive(classpath)
	if c == nil {
		b.Fatal(err)
	} else if err != nil {
		b.Log(err)
	}
	defer c.Close()
	data, err := c.LoadClass(Classname("java.lang.String"))
	if err != nil {
		b.Fatal(err)
	}
	r := bytes.NewReader(data)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.Seek(0, 0)
		NewClass(r)
	}
}
