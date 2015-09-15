package util

import "github.com/limetext/log4go"

func init() {
	log4go.AddFilter("stdout", log4go.WARNING, log4go.NewConsoleLogWriter())
}
