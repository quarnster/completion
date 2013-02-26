package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"testing"
)

func init() {
	listen()
	go accept()
}

func write(conn net.Conn, s string) {
	fmt.Fprintf(conn, "%v;%s", len(s), s)
}

func TestHandler(t *testing.T) {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		t.Fatal(err)
	}

	write(conn, `{"version": 1, "operation": "goto_doc"}`)

	if reply, err := ioutil.ReadAll(conn); err != nil {
		t.Fatal(err)
	} else {
		log.Println(string(reply))
	}

}
