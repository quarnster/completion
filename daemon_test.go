package main

import (
	"encoding/json"
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/quarnster/completion/java"
	"github.com/quarnster/completion/net"
	"net/rpc/jsonrpc"
	"reflect"
	"testing"
	"time"
)

func TestRpc(t *testing.T) {
	var d Daemon
	if err := d.init(); err != nil {
		t.Fatal(err)
	}
	go d.serverloop()
	defer d.close()
	if c, err := jsonrpc.Dial("tcp", fmt.Sprintf("127.0.0.1%s", port)); err != nil {
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
				"java://type/java/util/jar/JarEntry",
			},
			{
				&net.Net{},
				"Net.Complete",
				"net://type/System.String",
			},
		}
		for _, test := range tests {
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
	var d Daemon
	if err := d.init(); err != nil {
		t.Fatal(err)
	}
	go d.serverloop()
	defer d.close()

	if c, err := jsonrpc.Dial("tcp", fmt.Sprintf("127.0.0.1%s", port)); err != nil {
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

	if c, err := jsonrpc.Dial("tcp", fmt.Sprintf("127.0.0.1%s", port)); err != nil {
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
