package csharp

import (
	"github.com/quarnster/completion/util"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const testdata = "../testdata/"

func TestParse(t *testing.T) {
	if f, err := os.Open(testdata); err != nil {
		t.Fatal(err)
	} else if fi, err := f.Readdir(-1); err != nil {
		t.Fatal(err)
	} else {
		for i := range fi {
			if n := testdata + fi[i].Name(); strings.HasSuffix(n, ".cs") {
				var p CSHARP
				if d, err := ioutil.ReadFile(n); err != nil {
					t.Error(err)
				} else if !p.Parse(string(d)) {
					t.Errorf("%s: %s", n, p.Error())
				} else {
					back := p.RootNode().Children[len(p.RootNode().Children)-1]
					if back.Name != "EndOfFile" {
						t.Log(p.RootNode())
						t.Errorf("Didn't finish parsing %s: %s", n, p.Error())
					}
					root := p.RootNode()
					root.Simplify()
					res := root.String()
					if exp, err := ioutil.ReadFile(n + ".pt"); err != nil {
						t.Logf("Expected result for parsing %s does not exist and will be created", n)
						ioutil.WriteFile(n+".pt", []byte(res), 0644)
					} else if d := util.Diff(string(exp), res); len(d) > 0 {
						t.Error(d)
					}
				}
			}
		}
	}
}

func TestParse2(t *testing.T) {
	var p CSHARP
	p.SetData("if (fix && toadd.Trim().")
	p.Complete()
	t.Log(p.RootNode())
}
