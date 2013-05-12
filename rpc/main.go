package main

import (
	"code.google.com/p/log4go"
	"flag"
	"github.com/quarnster/completion/clang"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/java"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"runtime/debug"
)

func main() {
	var (
		port = ":12345"
		ifs  = []interface{}{
			&content.Session{},
			&clang.Clang{},
			&java.Java{},
		}
	)
	flag.StringVar(&port, "port", port, "TCP port the server will listen on")
	flag.Parse()

	log4go.Global.AddFilter("stdout", log4go.FINE, log4go.NewConsoleLogWriter())

	server := rpc.NewServer()

	for _, i := range ifs {
		if err := server.Register(i); err != nil {
			log4go.Crash(err)
		}
	}

	l, err := net.Listen("tcp", port)
	if err != nil {
		log4go.Crashf("Could not open port: %s", err)
	}
	for {
		if conn, err := l.Accept(); err != nil {
			log4go.Error("Error accepting connection: %s", err)
		} else {
			go func() {
				codec := jsonrpc.NewServerCodec(conn)
				defer func() {
					codec.Close()
					if r := recover(); r != nil {
						log4go.Error("Recovered from panic: %v, stack: %s", r, string(debug.Stack()))
					}
				}()
				server.ServeRequest(codec)
			}()
		}
	}
}
