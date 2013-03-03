package java

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestFileToClassAndClassToFile(t *testing.T) {
	var tests map[string]string
	if data, err := ioutil.ReadFile(testdata_path + "encoding_file_to_class.json"); err != nil {
		t.Fatalf("Failed to open testdata file: %s", err)
	} else if err := json.Unmarshal(data, &tests); err != nil {
		t.Fatalf("Failed to parse testdata file: %s", err)
	} else {
		for _k, _v := range tests {
			k, v := Filename(_k), Classname(_v)
			t.Logf("Testing conversion to and from %s", k)
			if res := k.Classname(); res != v {
				t.Errorf("Didn't properly convert file name %s to class name %s: %s", k, v, res)
			}
			if res := v.Filename(); res != k {
				t.Errorf("Didn't properly convert class name %s to file name %s: %s", v, k, res)
			}
		}
	}
}
