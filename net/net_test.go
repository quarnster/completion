package net

import (
	//	"code.google.com/p/log4go"
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/util"
	"io/ioutil"
	"testing"
)

func TestNet(t *testing.T) {
	var (
		n    Net
		args content.CompleteAtArgs
		cmp  content.CompletionResult
	)

	tests := []struct {
		InFile       string
		Line, Column uint
	}{
		// TODO: this test is not platform independent as it depends on whatever framework you happen
		//       to have installed
		{"./testdata/CompleteSharp.cs", 40, 27},
	}
	args.SessionId = "a"
	for _, test := range tests {
		args.Location.File.Name = test.InFile
		args.Location.Line = test.Line
		args.Location.Column = test.Column
		ex := fmt.Sprintf("%s-%d-%d.cmp", test.InFile, test.Line, test.Column)
		if err := n.CompleteAt(&args, &cmp); err != nil {
			t.Error(err)
		} else if exp, err := ioutil.ReadFile(ex); err != nil {
			t.Logf("Couldn't read the expected output file %s (%s); it'll be created", ex, err)
			if err := ioutil.WriteFile(ex, []byte(cmp.String()), 0644); err != nil {
				t.Error(err)
			}
		} else if d := util.Diff(string(exp), cmp.String()); len(d) != 0 {
			t.Error(d)
		}
	}
}
