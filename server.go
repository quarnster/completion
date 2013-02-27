package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/quarnster/completion/content"
	"log"
	"net"
	"strconv"
	"time"
)

var (
	ln net.Listener
)

func handleConnection(conn net.Conn) {
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

	intent := content.Intent{}
	if err := json.Unmarshal(b, &intent); err != nil {
		log.Println(err)
		fmt.Fprintf(conn, "Failed to decode request: %v", err)
	}

	content.Handle(conn, intent)
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
			go handleConnection(conn)
		}
	}
}

func startServer() {
	listen()
	accept()
}
