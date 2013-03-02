package net

import (
	"os"
	"testing"
)

// The expected value was cut'n'pasted from the output of: monodis $filename | grep GUID | awk '{print $6}'
var tests = [][2]string{
	{"CompleteSharp.exe", "3401283C-C835-429E-BD63-BDA728AB549F"},
	{"7zip.dll", "36DCBA29-69C8-4E99-B590-8BD2082695D5"},
	{"JsonParser.dll", "FA75A08B-59E8-4D4A-B10E-2114294958F4"},
	{"Png4BitIndexed.dll", "DC74EA34-D62F-4136-93BF-797DF4217B8E"},
	{"Png8BitIndexed.dll", "2B5CA898-BB41-4B8C-9925-988FA7F2F57A"},
	{"QRCodeGenerator.dll", "EF0BCC38-B8C8-444F-B6A8-96D99B4AF645"},
}

func TestGuid(t *testing.T) {
	for i := range tests {
		fn, expect := tests[i][0], tests[i][1]
		if f, err := os.Open(testdata_path + fn); err != nil {
			t.Error(err)
		} else {
			defer f.Close()
			if a, err := LoadAssembly(f); err != nil {
				t.Error(err)
			} else {
				idx := ConcreteTableIndex{&a.MetadataUtil, 1, id_Module}
				if raw, err := idx.Data(); err != nil {
					t.Error(err)
				} else {
					m := raw.(*ModuleRow)
					if m.Mvid.String() != expect {
						t.Errorf("Unexpected Guid: %s != %s", m.Mvid, expect)
					}
				}
			}
		}
	}
}
