package descriptors

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestParser(t *testing.T) {
	var p DESCRIPTORS
	var tests map[string]string
	if data, err := ioutil.ReadFile("./testdata/encoding_descriptors.json"); err != nil {
		t.Fatalf("Failed to open testdata file: %s", err)
	} else if err := json.Unmarshal(data, &tests); err != nil {
		t.Fatalf("Failed to parse testdata file: %s", err)
	} else {
		for k, v := range tests {
			t.Logf("Testing conversion to and from %s", k)
			if !p.Parse(k) {
				t.Errorf("Didn't parse test properly: %s", p.Error())
			} else if v2 := convert(p.RootNode()); v != v2 {
				t.Errorf("Didn't convert properly: Expected \"%s\", but got \"%s\"", v, v2)
			}
		}
	}
}

func BenchmarkParser(b *testing.B) {
	b.StopTimer()
	var p DESCRIPTORS
	var tests map[string]string
	if data, err := ioutil.ReadFile("./testdata/encoding_descriptors.json"); err != nil {
		b.Fatalf("Failed to open testdata file: %s", err)
	} else if err := json.Unmarshal(data, &tests); err != nil {
		b.Fatalf("Failed to parse testdata file: %s", err)
	} else {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			for k, _ := range tests {
				p.Parse(k)
				convert(p.RootNode())
			}
		}
	}
}
