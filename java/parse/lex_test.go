package parse

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestBasic(t *testing.T) {
	src, err := ioutil.ReadFile("testdata/Basic.java")
	if err != nil {
		t.Fatal(err)
	}

	l := lex(src)
	go l.run()
	for tkn := range l.tokens {
		var b []byte
		if tkn.start != tkn.end {
			b = bytes.TrimSpace(l.src[tkn.start : tkn.end-1])
		}
		fmt.Printf("type: %s, start: %v, end: %v, str: `%s`\n", tokenString[tkn.typ], tkn.start, tkn.end, b)
	}
}
