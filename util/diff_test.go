package util

import (
	"strings"
	"testing"
)

var tests = [][]string{
	{"a\nb\nc\nd", "a\nb\nd", "- c"},
	{"a\nb\nc\nd", "a\nb\nd\ne", "- c\n+ e"},
	{"a\nb\nc\nd\ne\nf", "a\nb\nd\ne\nf\ng", "- c\n+ g"},
	{"a\nb\nc\nd\ne\nf", "b\nd\ne\nf\ng", "- a\n- c\n+ g"},
	{"a\n\rb\n\rc\n\rd\n\re\n\rf", "b\nd\ne\nf\ng", "- a\n- c\n+ g"},
	{"a\nb\nc\nd\ne\nf", "b\n\rd\n\re\n\rf\n\rg", "- a\n- c\n+ g"},
	{"a\r\nb\r\nc\r\nd\r\ne\r\nf", "b\nd\ne\nf\ng", "- a\n- c\n+ g"},
}

func TestDiff(t *testing.T) {
	for _, test := range tests {
		if r := Diff(test[0], test[1]); !strings.EqualFold(test[2], r) {
			t.Errorf("Expected: %s\nGot: %s", test[2], r)
		} else {
			t.Logf("diff %#v %#v == %#v", test[0], test[1], r)
		}
	}
}
