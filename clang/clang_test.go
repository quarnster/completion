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
	if _, err := RunClang("-v"); err != nil {
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
	a.SessionOverrides.Set("clang_language", "c++")
	if err := c.CompleteAt(&a, &b); err != nil {
		t.Error(err)
	} else {
		t.Log(b)
	}
}

func TestClangCompleteatNotIdentifierBegin(t *testing.T) {
	if skip {
		t.Skipf("Clang not installed, skipping")
	}
	var (
		a content.CompleteAtArgs
		b content.CompletionResult
		c Clang
	)

	a.Location.Column = 10
	a.Location.Line = 5
	a.Location.File.Name = "testdata/set.cpp"
	a.SessionOverrides.Set("clang_language", "c++")
	if err := c.CompleteAt(&a, &b); err != nil {
		t.Error(err)
	} else {
		contain_clear := false
		contain_printf := false
		for _, method := range b.Methods {
			if method.Name.Relative == "clear" {
				contain_clear = true
			}
			if method.Name.Relative == "printf" {
				contain_printf = true
			}
		}
		if !contain_clear || contain_printf {
			t.Error("Completion result contains incorrect data")
		}
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
	a.SessionOverrides.Set("compiler_flags", []string{"-fno-exceptions"})
	a.SessionOverrides.Set("clang_language", "c++")
	if err := c.CompleteAt(&a, &b); err != nil {
		t.Error(err)
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
		fmt.Print("res:")
		fmt.Print(res)
		fmt.Print("\n")
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
	a.SessionOverrides.Set("clang_language", "c++")
	if err := c.GetDefinition(&a, &b); err != nil {
		t.Error(err)
	} else if b.File.Name == "" || b.Line == 0 || b.Column == 0 {
		t.Error(b)
	} else {
		t.Log(b)
	}
}

func TestIdentificatorBegin(t *testing.T) {
	if getIdentificatorBegin(" ~asd", 3) != 2 { t.Error("failed example 1") }
	if getIdentificatorBegin("   ~   asd", 8) != 7 { t.Error("failed example 2") }
	if getIdentificatorBegin(" a:: ~ asd", 8) != 5 { t.Error("failed example 3") }
	if getIdentificatorBegin(" a ->~as", 7) != 5 { t.Error("failed example 4") }
	if getIdentificatorBegin("a. ~s", 4) != 3 { t.Error("failed example 5") }
	if getIdentificatorBegin(" asdasd", 5) != 1 { t.Error("failed example 6") }
	if getIdentificatorBegin(".", 1) != 1 { t.Error("failed example 7") }
	if getIdentificatorBegin("a", 1) != 0 { t.Error("failed example 8") }
}
