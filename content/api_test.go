package content

import (
	"bytes"
	// "encoding/binary"
	"encoding/gob"
	"fmt"
	"testing"
)

func greet(s string) (ret string) {
	return fmt.Sprintf("Hello, %s!", s)
}

func TestNewAPI(t *testing.T) {
	NewAPI(greet)

	var data []Data
	data = append(data, Data{"s", []byte("World")})
	intent := Intent{1, "content.greet", data}

	// represents client connection
	buf := new(bytes.Buffer)

	//
	if err := Handle(buf, intent); err != nil {
		t.Fatal(err)
	}

	// represents serialization to client and deserialization
	var resp Response
	gob.NewDecoder(buf).Decode(&resp)

	/* not working due to complications in api.go

	binary.Read(buf, binary.LittleEndian, &resp)
	*/

	if len(resp.Data) != 1 {
		t.Fatalf("Expected one named returned but received %v.", len(resp.Data))
	}

	var s string
	s = string(resp.Data[0].Value)
	// bufResp := bytes.NewBuffer(resp.Data[0].Value)
	// binary.Read(bufResp, binary.LittleEndian, &s)

	//
	if s != "Hello, World!" {
		t.Fatalf(`Expected "Hello, World!" but received "%s".`, s)
	}
}
