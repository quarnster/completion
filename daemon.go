package main

import (
	"github.com/limetext/log4go"
	"flag"
	"github.com/quarnster/completion/clang"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/java"
	cnet "github.com/quarnster/completion/net"
	"github.com/robmerrell/comandante"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/signal"
	"runtime/debug"
	"time"
)

var (
	port   = "/tmp/completion.rpc"
	proto  = "unix"
	rmFile = false
)

func init() {
	daemon := comandante.NewCommand("daemon", "Run rpc daemon", runRpcDaemon)
	daemon.FlagInit = daemonFlagInit
	bin.RegisterCommand(daemon)
}

func daemonFlagInit(fs *flag.FlagSet) {
	fs.StringVar(&port, "port", port, "TCP port the server will listen on")
	fs.StringVar(&proto, "proto", proto, "Network protocol type (tcp, udp, unix, see http://golang.org/pkg/net/)")
	fs.BoolVar(&rmFile, "remove", rmFile, "Remove unix socket file if it already exists (when using unix protocol)")
}

type Daemon struct {
	quit   bool
	l      net.Listener
	server *rpc.Server
}

func (d *Daemon) init() error {
	d.quit = false
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
	if proto == "unix" && rmFile {
		os.Remove(port)
	}
	var err error
	d.l, err = net.Listen(proto, port)
	if err != nil {
		return err
	}
	return nil
}

func (d *Daemon) close() error {
	d.quit = true
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
	defer d.close()

	return d.serverloop()
}

func (d *Daemon) handleConn(conn net.Conn) {
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
}

func (d *Daemon) serverloop() error {
	errorcount := 0
	sigchan := make(chan os.Signal)
	conchan := make(chan net.Conn)
	errchan := make(chan error)
	signal.Notify(sigchan, os.Interrupt, os.Kill)
	go func() {
		for {
			if conn, err := d.l.Accept(); err != nil {
				if d.quit {
					return
				}
				errchan <- log4go.Error("Error accepting connection: %s", err)
			} else {
				conchan <- conn
			}
		}
	}()
outer:
	for {
		select {
		case s := <-sigchan:
			log4go.Debug("Exiting due to signal: %s", s)
			d.quit = true
			break outer
		case conn := <-conchan:
			go d.handleConn(conn)
		case <-errchan:
			errorcount++
			if errorcount > 10 {
				return log4go.Error("Too many errors, shutting server down")
			}
		}
	}
	return nil
}
