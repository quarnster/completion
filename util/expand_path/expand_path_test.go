package expand_path

import (
	"os"
	"testing"
)

func TestExpandPath(t *testing.T) {
	if err := os.Setenv("DUMMY_ROCK", "ABC"); err != nil {
		t.Fatal(err)
	} else if err := os.Setenv("DUMMY_PAPER", "DEF"); err != nil {
		t.Fatal(err)
	}
	var tests = []struct {
		in, out string
	}{
		{"-I${env:DUMMY_ROCK} -I${env:DUMMY_PAPER}", "-IABC -IDEF"},
		{"-I${env:DUMMY_ROCK}/${env:DUMMY_PAPER} -I${env:DUMMY_PAPER}", "-IABC/DEF -IDEF"},
		{"-I${env:DUMMY_ROCK}/${env:DUMMY_PAPER}.cpp -I${env:DUMMY_PAPER}", "-IABC/DEF.cpp -IDEF"},
		{"-I${folder:${env:DUMMY_ROCK}/${env:DUMMY_PAPER}.cpp} -I${env:DUMMY_PAPER}", "-IABC -IDEF"},
		{"-I${folder:${home}/${env:DUMMY_PAPER}.cpp} -I${env:DUMMY_PAPER}", ExpandPath("-I${home} -IDEF")},
	}
	for _, test := range tests {
		if out := ExpandPath(test.in); out != test.out {
			t.Errorf("With input: %s, expected %s, but got %s", test.in, test.out, out)
		}
	}
}
