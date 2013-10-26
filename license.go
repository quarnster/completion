package main

import (
	"bytes"
	"fmt"
	"github.com/quarnster/completion/editor"
	"github.com/robmerrell/comandante"
	"io/ioutil"
)

func init() {
	license := comandante.NewCommand("license", "Print licenses used", printLicenses)
	bin.RegisterCommand(license)
}

func printLicenses() error {
	f, err := editor.Open("README.md")
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	sep := []byte("# License")
	idx := bytes.Index(data, sep) + len(sep)
	fmt.Println(string(data[idx:]))
	return err
}
