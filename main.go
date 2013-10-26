package main

import (
	"code.google.com/p/log4go"
	"github.com/robmerrell/comandante"
	"os"
)

var bin = comandante.New(os.Args[0], os.Args[0]+" code utilities")

func init() {
	// TODO: should really log to stderr and not stdout
	log4go.Global.AddFilter("stdout", log4go.FINE, log4go.NewConsoleLogWriter())
}

func main() {
	bin.IncludeHelp()

	if err := bin.Run(); err != nil {
		log4go.Error("Error running command: %v", err)
	}
}
