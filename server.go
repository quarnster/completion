package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type Handler func(io.Writer, Intent)

var (
	ln net.Listener

	handlers = map[string]Handler{
		"foo": foo,
	}

	flagAddr = flag.String("addr", ":8888", "Local address to listen on.")
)

func init() {
	flag.Parse()
}

func foo(w io.Writer, intent Intent) {
	log.Println(intent)
	if b, err := json.Marshal(&Response{1}); err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Failed to create response: %v", err)
	} else {
		w.Write(b)
	}
}

type Intent struct {
	Version   int64
	Operation string
}

type Response struct {
	Version int64 `json:"version"`
}

func handleIntent(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(time.Second * 3))

	r := bufio.NewReader(conn)

	size, err := r.ReadString(';')
	if err != nil {
		log.Println(err)
		fmt.Fprintf(conn, "Failed to read content length: %v\n", err)
		return
	}

	size = size[:len(size)-1]
	i, err := strconv.Atoi(size)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(conn, "Failed to get content length from %s: %v\n", size, err)
		return
	}

	b := make([]byte, i)
	r.Read(b)

	intent := Intent{}
	if err := json.Unmarshal(b, &intent); err != nil {
		log.Println(err)
		fmt.Fprintf(conn, "Failed to decode request: %v", err)
	}

	// TODO get registered handler for operation
	handlers["foo"](conn, intent)
}

func listen() {
	l, err := net.Listen("tcp", *flagAddr)
	if err != nil {
		log.Fatal(err)
	}
	ln = l
}

func accept() {
	for {
		if conn, err := ln.Accept(); err != nil {
			log.Print(err)
		} else {
			go handleIntent(conn)
		}
	}
}

func startServer() {
	listen()
	accept()
}

func main() {
	startServer()
}
