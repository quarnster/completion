package content

import (
	"compress/gzip"
	"encoding/json"
	"github.com/quarnster/completion/util"
	"io/ioutil"
	"os"
	"testing"
)

func loadData() (cmp CompletionResult, err error) {
	if f, err := os.Open("./testdata/big.json.gz"); err != nil {
		return cmp, err
	} else {
		defer f.Close()
		if r, err := gzip.NewReader(f); err != nil {
			return cmp, err
		} else {
			defer r.Close()
			dec := json.NewDecoder(r)
			if err := dec.Decode(&cmp); err != nil {
				return cmp, err
			}
		}
	}
	return
}

func TestString(t *testing.T) {
	if cmp, err := loadData(); err != nil {
		t.Fatal(err)
	} else {
		expfile := "./testdata/big.json.gz.in"
		expected := ""
		if d, err := ioutil.ReadFile(expfile); err != nil {
			t.Log("Expected result file doesn't exist and will be created")
		} else {
			expected = string(d)
		}
		res := cmp.String()
		if len(expected) <= 1 {
			t.Logf("Creating new test data: %s", expfile)
			if err := ioutil.WriteFile(expfile, []byte(res), 0644); err != nil {
				t.Errorf("Couldn't write test data to %s: %s", expfile, err)
			}
		} else if d := util.Diff(expected, res); len(d) != 0 {
			t.Error(d)
		}
	}
}

func BenchmarkString(b *testing.B) {
	b.StopTimer()
	if cmp, err := loadData(); err != nil {
		b.Fatal(err)
	} else {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			cmp.String()
		}
	}
}
