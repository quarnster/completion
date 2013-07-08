package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files := os.Args[1:]
	fmt.Println("Creating archive with the contents of", files)
	out := bytes.NewBuffer(nil)
	z := zip.NewWriter(out)
	for i := range files {
		if w, err := z.Create(files[i][3:]); err != nil {
			panic(err)
		} else if d, err := ioutil.ReadFile(files[i]); err != nil {
			panic(err)
		} else if n, err := w.Write(d); err != nil {
			panic(err)
		} else if n != len(d) {
			panic(fmt.Errorf("%d != %d", n != len(d)))
		}
	}
	if err := z.Close(); err != nil {
		panic(err)
	}
	b2 := bytes.NewBuffer(nil)
	b2.WriteString(`package editor

var archive_data = []byte{`)
	for i, b := range out.Bytes() {
		if i > 0 {
			b2.WriteRune(',')
		}
		b2.WriteString(fmt.Sprintf("%v", b))
	}
	b2.WriteString("}")
	if err := ioutil.WriteFile("../editor/archive_data.go", b2.Bytes(), 0644); err != nil {
		panic(err)
	}
}
