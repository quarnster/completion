package main

import (
	"flag"
)

var (
	flagAddr = flag.String("addr", ":8888", "Local address to listen on.")
)

func init() {
	flag.Parse()
}

func main() {
	startServer()
}
