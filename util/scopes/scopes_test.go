package scopes

import (
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util"
	"github.com/quarnster/parser"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
)

func TestScopes(t *testing.T) {
	tests := []struct {
		in string
	}{
		{"../../clang/testdata/hello.cpp"},
		{"../../net/testdata/CompleteSharp.cs"},
	}
	for _, test := range tests {
		if d, err := ioutil.ReadFile(test.in); err != nil {
			t.Error(err)
		} else {
			var p SCOPES
			if !p.Parse(string(d)) {
				t.Errorf("Failed to parse: %s\n%v", p.Error(), p.RootNode())
				continue
			}
			ex := "testdata/" + filepath.Base(test.in) + ".exp"
			cmp := p.RootNode().String()
			if exp, err := ioutil.ReadFile(ex); err != nil {
				t.Logf("Couldn't read the expected output file %s (%s); it'll be created", ex, err)
				if err := ioutil.WriteFile(ex, []byte(cmp), 0644); err != nil {
					t.Error(err)
				}
			} else if d := util.Diff(string(exp), cmp); len(d) != 0 {
				t.Error(d)
			}
		}
	}
}

func TestVisibility(t *testing.T) {
	if d, err := ioutil.ReadFile("../../net/testdata/CompleteSharp.cs"); err != nil {
		t.Error(err)
	} else {
		var p SCOPES
		if !p.Parse(string(d)) {
			t.Fatalf("Failed to parse: %s\n%v", p.Error(), p.RootNode())
		}
		tests := []struct {
			line uint
			res  parser.RangeSet
		}{
			{32, []parser.Range{{0, 0}, {872, 1090}, {1096, 1158}, {34800, 34850}, {35021, 35071}, {37050, 37053}}},
			{116, []parser.Range{{0, 0}, {872, 1090}, {1096, 1248}, {1337, 1414}, {2803, 2897}, {3812, 3990}, {3991, 4012}, {4108, 4322}, {4534, 4640}, {4989, 5115}, {6211, 6294}, {6673, 6741}, {6928, 7036}, {8234, 8288}, {8502, 8570}, {10535, 10643}, {11211, 11309}, {12512, 12572}, {12647, 12711}, {13012, 13093}, {32052, 32231}, {32346, 32395}, {33300, 33375}, {33640, 33722}, {34794, 34850}, {35021, 35071}, {37050, 37053}}},
			{730, []parser.Range{{0, 0}, {872, 1090}, {1096, 1208}, {32062, 32231}, {32346, 32486}, {32502, 32600}, {32611, 32708}, {32738, 32883}, {33290, 33375}, {33640, 33722}, {34794, 34850}, {35021, 35071}, {37050, 37053}}},
		}
		loc := content.SourceLocation{File: content.File{Contents: string(d)}, Column: 1}
		for i := range tests {
			loc.Line = tests[i].line
			vis := Visibility(loc)
			if !reflect.DeepEqual(vis, tests[i].res) {
				t.Errorf("Output differs. Expected %v, got %v", tests[i].res, vis)
			}
			if testing.Verbose() {
				t.Log(Substr(loc.File.Contents, vis))
			}
		}
	}
}
