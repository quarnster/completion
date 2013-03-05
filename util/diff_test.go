package util

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
)

func TestDiff(t *testing.T) {
	const testfile = "./testdata/diff.json"
	var tests [][]string
	if d, err := ioutil.ReadFile(testfile); err != nil {
		t.Fatal(err)
	} else if err := json.Unmarshal(d, &tests); err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		if r := Diff(test[0], test[1]); !strings.EqualFold(test[2], r) {
			t.Errorf("Expected: %s\nGot: %s", test[2], r)
		} else {
			t.Logf("diff\n%s", r)
		}
	}
}
