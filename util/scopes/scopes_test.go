package scopes

import (
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util"
	"github.com/quarnster/util/text"
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
			res  []text.Region
		}{
			{32, []text.Region{{872, 1090}, {1096, 1158}, {34796, 34846}, {35017, 35067}, {37046, 37049}}},
			{116, []text.Region{{872, 1090}, {1096, 1248}, {1337, 1414}, {2803, 2897}, {3812, 3990}, {3991, 4012}, {4108, 4321}, {4533, 4565}, {4565, 4639}, {4988, 5114}, {6210, 6293}, {6672, 6740}, {6927, 7035}, {8232, 8286}, {8500, 8568}, {10533, 10641}, {11209, 11307}, {12510, 12570}, {12645, 12709}, {13010, 13091}, {32050, 32228}, {32343, 32392}, {33297, 33371}, {33636, 33718}, {34790, 34846}, {35017, 35067}, {37046, 37049}}},
			{730, []text.Region{{872, 1090}, {1096, 1208}, {32060, 32228}, {32343, 32483}, {32499, 32597}, {32608, 32705}, {32735, 32880}, {33287, 33371}, {33636, 33718}, {34790, 34846}, {35017, 35067}, {37046, 37049}}},
		}
		loc := content.SourceLocation{File: content.File{Contents: string(d)}, Column: 1}
		for i := range tests {
			loc.Line = tests[i].line
			vis := Visibility(loc)
			if !reflect.DeepEqual(vis.Regions(), tests[i].res) {
				t.Errorf("Output differs.\nExpected\n%v, got\n%v", tests[i].res, vis.Regions())
			}
			if testing.Verbose() {
				t.Log(Substr(loc.File.Contents, vis))
			}
		}
	}
}
