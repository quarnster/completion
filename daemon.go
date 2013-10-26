package main

import (
	"code.google.com/p/log4go"
	"flag"
	"github.com/quarnster/completion/clang"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/java"
	cnet "github.com/quarnster/completion/net"
	"github.com/robmerrell/comandante"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"runtime/debug"
	"time"
)

var (
	port = ":12345"
)

func init() {
	daemon := comandante.NewCommand("daemon", "Run rpc daemon", runRpcDaemon)
	daemon.FlagInit = daemonFlagInit
	bin.RegisterCommand(daemon)
}

func daemonFlagInit(fs *flag.FlagSet) {
	fs.StringVar(&port, "port", port, "TCP port the server will listen on")
}

type Daemon struct {
	l      net.Listener
	server *rpc.Server
}

func (d *Daemon) init() error {
	ifs := []interface{}{
		&content.Session{},
		&clang.Clang{},
		&java.Java{},
		&cnet.Net{},
	}
	d.server = rpc.NewServer()

	for _, i := range ifs {
		if err := d.server.Register(i); err != nil {
			return err
		}
	}
	var err error
	d.l, err = net.Listen("tcp", port)
	if err != nil {
		return err
	}
	return nil
}

func (d *Daemon) close() error {
	return d.l.Close()
}

func runRpcDaemon() error {
	var (
		d   Daemon
		err error
	)
	if err = d.init(); err != nil {
		return err
	}

	return d.serverloop()
}

func (d *Daemon) serverloop() error {
	errorcount := 0
	for {
		if conn, err := d.l.Accept(); err != nil {
			errorcount++
			log4go.Error("Error accepting connection: %s", err)
			if errorcount > 10 {
				return log4go.Error("Too many errors, shutting server down")
			}
		} else {
			go func() {
				s := time.Now()
				conn.SetDeadline(time.Time{})
				codec := jsonrpc.NewServerCodec(conn)
				defer func() {
					codec.Close()
					if r := recover(); r != nil {
						log4go.Error("Recovered from panic: %v, stack: %s", r, string(debug.Stack()))
					}
					log4go.Debug("Serviced in %f milliseconds", time.Since(s).Seconds()*1000)
				}()
				for {
					if err := d.server.ServeRequest(codec); err != nil {
						log4go.Error("Error handling request: %v", err)
						break
					}
				}
			}()
		}
	}
}
