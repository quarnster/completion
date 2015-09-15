package main

import (
	"github.com/limetext/log4go"
	"github.com/robmerrell/comandante"
	"os"
	"time"
)

var bin = comandante.New(os.Args[0], os.Args[0]+" code utilities")

func init() {
	// TODO: should really log to stderr and not stdout
	log4go.Global.AddFilter("stdout", log4go.FINEST, log4go.NewConsoleLogWriter())
}

func main() {
	bin.IncludeHelp()
	err := bin.Run()
	// TODO: HACK! See https://code.google.com/p/log4go/issues/detail?id=8
	// Give log4go a chance to catch up and print all its messages
	time.Sleep(time.Second)
	if err != nil {
		log4go.Exit(err)
	}
}
