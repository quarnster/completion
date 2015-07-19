package clang

import (
	"fmt"
	"github.com/quarnster/completion/clang/parser"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

var skip = false

func TestClang(t *testing.T) {
	if _, _, err := RunClang("", "-v"); err != nil {
		skip = true
		t.Skipf("Couldn't launch clang: %s", err)
	}
	var (
		a content.CompleteAtArgs
		b content.CompletionResult
		c Clang
	)

	a.Location.Column = 1
	a.Location.Line = 4
	a.Location.File.Name = "testdata/hello.cpp"
	if err := c.CompleteAt(&a, &b); err != nil {
		t.Error(err)
	} else {
		t.Log(b)
	}
}

func TestClangUnsaved(t *testing.T) {
	if skip {
		t.Skipf("Clang not installed, skipping")
	}
	var (
		a        content.CompleteAtArgs
		b        content.CompletionResult
		c        Clang
		expected string
	)
	a.Location.File.Name = "test.cpp"
	a.Location.File.Contents = `
class Test {
public:
	int A;
	int B;
	int C;
};

void main() {
	Test t;
	t.
}

`
	a.Location.Line = 11
	a.Location.Column = 4
	if err := c.CompleteAt(&a, &b); err == nil {
		// Need to provide -xc++ or similar, so it should complain about this if we didn't
		t.Error("Expected an error, but didn't get one")
	}
	a.SessionOverrides.Set("compiler_flags", []string{"-x", "c++", "-fno-exceptions"})
	if err := c.CompleteAt(&a, &b); err != nil {
		t.Errorf("Unexpected error: %s", err)
	} else if strings.Index(b.Messages, "void main") == -1 {
		t.Errorf("Expected an error about main returning void but got: \"%s\"", b.Messages)
	}
	res := fmt.Sprintf("%s", b)
	fn := "./testdata/unsaved.cpp.res"
	if data, err := ioutil.ReadFile(fn); err != nil {
		t.Logf("Error reading test data: %s. Testdata will be created", err)
	} else {
		expected = string(data)
	}
	if len(expected) <= 1 {
		// Just if we want to add new tests, this will spit out the newly added
		// test data
		t.Logf("Creating new test data: %s", fn)
		if err := ioutil.WriteFile(fn, []byte(res), 0644); err != nil {
			t.Errorf("Couldn't write test data to %s: %s", fn, err)
		}
	} else if d := util.Diff(expected, res); len(d) != 0 {
		t.Error(d)
	}

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
		if !strings.HasSuffix(k, "j.in") {
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

func BenchmarkParse(b *testing.B) {
	if raw, err := ioutil.ReadFile("./testdata/mm5.in"); err != nil {
		b.Fatal(err)
	} else {
		var p parser.PARSER
		data := string(raw)
		for i := 0; i < b.N; i++ {
			p.Parse(data)
		}
	}
}

func BenchmarkParseResult(b *testing.B) {
	if raw, err := ioutil.ReadFile("./testdata/mm5.in"); err != nil {
		b.Fatal(err)
	} else {
		data := string(raw)
		for i := 0; i < b.N; i++ {
			parseresult(data)
		}
	}
}

func TestGetDefinition(t *testing.T) {
	if skip {
		t.Skipf("Clang not installed, skipping")
	}
	var (
		a content.GetDefinitionArgs
		b content.SourceLocation
		c Clang
	)
	a.Identifier = "printf"
	a.Location.Column = 2
	a.Location.Line = 4
	a.Location.File.Name = "testdata/hello.cpp"
	if err := c.GetDefinition(&a, &b); err != nil {
		t.Error(err)
	} else if b.File.Name == "" || b.Line == 0 || b.Column == 0 {
		t.Error(b)
	} else {
		t.Log(b)
	}
}
