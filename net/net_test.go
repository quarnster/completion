package net

import (
	"code.google.com/p/log4go"
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/net/csharp"
	"github.com/quarnster/completion/util"
	"io/ioutil"
	"testing"
)

func init() {
	log4go.Close()
	log4go.AddFilter("test", log4go.DEBUG, log4go.NewConsoleLogWriter())
}
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
		{"./testdata/CompleteSharp.cs", 40, 41},
		{"./testdata/CompleteSharp.cs", 47, 47},
		//{"./testdata/CompleteSharp.cs", 28, 14},
		{"./testdata/CompleteSharp.cs", 211, 46},
		{"./testdata/CompleteSharp.cs", 761, 83},
		{"./testdata/CompleteSharp.cs", 761, 27},
		{"./testdata/CompleteSharp.cs", 857, 29},
		{"./testdata/CompleteSharp.cs", 737, 15},
		{"./testdata/CompleteSharp.cs", 95, 38},
		{"./testdata/CompleteSharp.cs", 95, 45},
		{"./testdata/CompleteSharp.cs", 776, 39},
	}
	args.SessionId = "a"
	args.Settings().Set("net_paths", []string{"./testdata/"})
	args.Settings().Set("net_assemblies", []string{"CompleteSharp.exe"})
	for _, test := range tests {
		args.Location.File.Name = test.InFile
		args.Location.Line = test.Line
		args.Location.Column = test.Column
		ex := fmt.Sprintf("%s-%d-%d.cmp", test.InFile, test.Line, test.Column)
		if err := n.CompleteAt(&args, &cmp); err != nil {
			t.Errorf("Unable to complete %v: %s", test, err)
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

func BenchmarkFindtype(b *testing.B) {
	var (
		n    Net
		args content.CompleteAtArgs
	)
	args.Location.File.Name = "./testdata/CompleteSharp.cs"
	args.Location.File.Load()
	args.Location.Line = 95
	args.Location.Column = 45
	cache, _ := n.cache(&args.Args)

	var up csharp.CSHARP
	up.SetData(args.Location.File.Contents)
	if !up.UsingDirectives() {
		b.Fatal(up.Error())
	}
	using := up.RootNode()

	for i := 0; i < b.N; i++ {
		findtype(cache, using, "string")
	}
}

func BenchmarkNetComplete(b *testing.B) {
	var (
		n    Net
		args content.CompleteAtArgs
	)
	args.Location.File.Name = "./testdata/CompleteSharp.cs"
	args.Location.File.Load()
	args.Location.Line = 95
	args.Location.Column = 45
	for i := 0; i < b.N; i++ {
		var cmp content.CompletionResult
		n.CompleteAt(&args, &cmp)
	}
}
