package clang

import (
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestClang(t *testing.T) {
	if _, err := RunClang("-v"); err != nil {
		t.Skipf("Couldn't launch clang: %s", err)
	}
	loc := content.SourceLocation{}
	loc.Column = 1
	loc.Line = 10
	loc.File.Name = "testdata/hello.cpp"
	t.Log(CompleteAt(loc))
}

func TestParseResult(t *testing.T) {
	var tests = make(map[string][]string)
	if dir, err := os.Open("./testdata"); err != nil {
		t.Fatal(err)
	} else if fi, err := dir.Readdir(0); err != nil {
		t.Fatal(err)
	} else {
		for _, f := range fi {
			if strings.HasSuffix(f.Name(), ".in") {
				path := "./testdata/" + f.Name()
				if input, err := ioutil.ReadFile(path); err != nil {
					t.Fatal(err)
				} else {
					data, err := ioutil.ReadFile(path + ".res")
					if err != nil {
						t.Logf("Error reading test data: %s. Testdata will be created", err)
					}
					tests[path] = []string{string(input), string(data)}
				}
			}
		}
	}
	for k, data := range tests {
		if strings.HasSuffix(k, "mm5.in") {
			// With almost 30000 lines this one is way to slow to parse and diff at the moment
			continue
		}
		s := time.Now()
		input, expected := data[0], data[1]
		res := ""
		if cmp, err := parseresult(input); err != nil {
			t.Error(err)
			continue
		} else {
			res = fmt.Sprintf("%s", cmp)
		}
		e := time.Now()
		t.Logf("Parse %s %s", k, e.Sub(s))
		if len(expected) <= 1 {
			// Just if we want to add new tests, this will spit out the newly added
			// test data
			fn := k + ".res"
			t.Logf("Creating new test data: %s", fn)
			if err := ioutil.WriteFile(fn, []byte(res), 0644); err != nil {
				t.Errorf("Couldn't write test data to %s: %s", fn, err)
			}
			continue
		}
		s = time.Now()
		if d := util.Diff(expected, res); len(d) != 0 {
			t.Error(d)
		}
		e = time.Now()

		t.Logf("Diff %s %s", k, e.Sub(s))
	}

}
