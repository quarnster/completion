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
			{32, []parser.Range{{0, 1157}, {34801, 34849}, {35022, 35070}, {37051, 37055}}},
			{116, []parser.Range{{0, 1247}, {1338, 1413}, {2804, 2896}, {3813, 4011}, {4109, 4321}, {4535, 4639}, {4990, 5114}, {6212, 6293}, {6674, 6740}, {6929, 7035}, {8235, 8287}, {8503, 8569}, {10536, 10642}, {11212, 11308}, {12513, 12571}, {12648, 12710}, {13013, 13092}, {32053, 32230}, {32347, 32394}, {33301, 33374}, {33641, 33721}, {34795, 34849}, {35022, 35070}, {37051, 37055}}},
			{730, []parser.Range{{0, 1207}, {32063, 32230}, {32347, 32882}, {33291, 33374}, {33641, 33721}, {34795, 34849}, {35022, 35070}, {37051, 37055}}},
		}
		loc := content.SourceLocation{File: content.File{Contents: string(d)}, Column: 1}
		for i := range tests {
			loc.Line = tests[i].line
			vis := Visibility(loc)
			if !reflect.DeepEqual(vis, tests[i].res) {
				t.Errorf("Output differs. Expected %v, got %v", tests[i].res, vis)
			}
			if testing.Verbose() {
				for _, r := range vis {
					t.Log(loc.File.Contents[r.Start:r.End])
				}
			}
		}
	}
}
