package main

import (
	"code.google.com/p/log4go"
	"flag"
	"github.com/quarnster/completion/clang"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/java"
	cnet "github.com/quarnster/completion/net"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"runtime/debug"
)

var (
	port    = ":12345"
	l       net.Listener
	server  *rpc.Server
	running = make(chan bool)
)

func init() {
	var (
		ifs = []interface{}{
			&content.Session{},
			&clang.Clang{},
			&java.Java{},
			&cnet.Net{},
		}
		err error
	)

	flag.StringVar(&port, "port", port, "TCP port the server will listen on")
	flag.Parse()
	log4go.Global.AddFilter("stdout", log4go.FINE, log4go.NewConsoleLogWriter())
	server = rpc.NewServer()

	for _, i := range ifs {
		if err := server.Register(i); err != nil {
			log4go.Crash(err)
		}
	}

	l, err = net.Listen("tcp", port)
	if err != nil {
		log4go.Crashf("Could not open port: %s", err)
	}
	go serverloop()
}

func serverloop() {
	errorcount := 0
	for {
		if conn, err := l.Accept(); err != nil {
			errorcount++
			log4go.Error("Error accepting connection: %s", err)
			if errorcount > 10 {
				log4go.Error("Too many errors, shutting server down")
				break
			}
		} else {
			go func() {
				codec := jsonrpc.NewServerCodec(conn)
				defer func() {
					log4go.Debug("In defer")
					codec.Close()
				}()
				handle := func() error {
					defer func() {
						if r := recover(); r != nil {
							log4go.Error("Recovered from panic: %v, stack: %s", r, string(debug.Stack()))
							// TODO(q): what would be the proper way to handle this?
							codec.Close()
						}
					}()
					return server.ServeRequest(codec)
				}
				for {
					if err := handle(); err != nil {
						log4go.Error("Error handling request: %v", err)
						break
					}
				}
			}()
		}
	}
	running <- false
}

func main() {
	<-running
}
