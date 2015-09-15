package main

import (
	"github.com/limetext/log4go"
	"encoding/json"
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/java"
	"github.com/quarnster/completion/net"
	"net/rpc/jsonrpc"
	"reflect"
	"strings"
	"testing"
	"time"
)

type testlogger struct {
	t *testing.T
}

func (t *testlogger) LogWrite(rec *log4go.LogRecord) {
	t.t.Log(strings.TrimSpace(log4go.FormatLogRecord(log4go.FORMAT_DEFAULT, rec)))
}
func (t *testlogger) Close() {
}

func TestRpc(t *testing.T) {
	log4go.Global.Close()
	log4go.Global.AddFilter("test", log4go.FINEST, &testlogger{t})
	var d Daemon
	if err := d.init(); err != nil {
		t.Fatal(err)
	}
	go d.serverloop()
	defer d.close()
	if c, err := jsonrpc.Dial(proto, port); err != nil {
		t.Error(err)
	} else {
		defer c.Close()
		tests := []struct {
			skip     func() error
			complete content.Complete
			rpcname  string
			abs      string
		}{
			{
				func() error {
					_, e := java.DefaultClasspath()
					return e
				},
				&java.Java{},
				"Java.Complete",
				"java://type/java/util/jar/JarEntry",
			},
			{
				func() error {
					if len(net.DefaultPaths()) == 0 {
						return fmt.Errorf("net paths are empty...")
					}
					return nil
				},
				&net.Net{},
				"Net.Complete",
				"net://type/System.String",
			},
		}
		for _, test := range tests {
			if err := test.skip(); err != nil {
				t.Logf("skipping: %v", err)
				continue
			}
			var a content.CompleteArgs
			var cmp1, cmp2 content.CompletionResult
			a.Location.Absolute = test.abs
			if err := test.complete.Complete(&a, &cmp1); err != nil {
				t.Error(err)
			}
			t.Log("calling", test.rpcname)
			if err := c.Call(test.rpcname, &a, &cmp2); err != nil {
				t.Error(err)
			} else if !reflect.DeepEqual(cmp1, cmp2) {
				t.Errorf("Results aren't equal: %v\n==============\n:%v", cmp1, cmp2)
			}
		}
	}
}

func TestRpcInvalid(t *testing.T) {
	log4go.Global.Close()
	log4go.Global.AddFilter("test", log4go.FINEST, &testlogger{t})

	var d Daemon
	if err := d.init(); err != nil {
		t.Fatal(err)
	}
	go d.serverloop()
	defer d.close()

	if c, err := jsonrpc.Dial(proto, port); err != nil {
		t.Error(err)
	} else {
		defer c.Close()
		tests := []struct {
			complete content.Complete
			rpcname  string
			abs      string
		}{
			{
				&java.Java{},
				"Java.Complete",
				"", // Intentional bad request
			},
		}
		for _, test := range tests {
			var a content.CompleteArgs
			var cmp2 content.CompletionResult
			a.Location.Absolute = test.abs
			t.Log("calling", test.rpcname)
			if err := c.Call(test.rpcname, &a, &cmp2); err == nil {
				t.Error("Expected an error, but didn't receive one")
			} else {
				t.Logf("As expected got an error: %s", err)
			}
		}
	}
}

func BenchmarkDirect(b *testing.B) {
	var n net.Net
	for i := 0; i < b.N; i++ {
		var a content.CompleteAtArgs
		a.Location.Line = 95
		a.Location.Column = 45
		a.Location.File.Name = "../net/testdata/CompleteSharp.cs"
		var cmp1 content.CompletionResult
		n.CompleteAt(&a, &cmp1)
	}
}

func BenchmarkRpc(b *testing.B) {
	var d Daemon
	if err := d.init(); err != nil {
		b.Fatal(err)
	}
	go d.serverloop()
	defer d.close()

	if c, err := jsonrpc.Dial(proto, port); err != nil {
		b.Error(err)
	} else {
		defer c.Close()

		s := time.Now()
		for i := 0; i < b.N; i++ {
			var a content.CompleteAtArgs
			a.Location.Line = 95
			a.Location.Column = 45
			a.Location.File.Name = "../net/testdata/CompleteSharp.cs"

			var cmp1 content.CompletionResult
			if err := c.Call("Net.CompleteAt", &a, &cmp1); err != nil {
				b.Error(err)
			}
		}
		e := time.Since(s).Seconds() * 1000
		b.Logf("%d calls in %f ms = %f ms/call", b.N, e, e/float64(b.N))
	}
}

func BenchmarkJsonMarshal(b *testing.B) {
	var n net.Net
	for i := 0; i < b.N; i++ {
		var a content.CompleteAtArgs
		a.Location.Line = 95
		a.Location.Column = 45
		a.Location.File.Name = "../net/testdata/CompleteSharp.cs"
		var cmp1 content.CompletionResult
		n.CompleteAt(&a, &cmp1)
		json.Marshal(&cmp1)
	}
}
