package dwarf

import (
	"bytes"
	"debug/macho"
	"github.com/quarnster/util/encoding/binary"
	"io"
	"testing"
)

func TestInfo(t *testing.T) {
	for _, test := range []string{"./testdata/8", "./testdata/hello", "./testdata/game.bz2", "./testdata/listener.o.bz2", "./testdata/completion.bz2"} {
		t.Logf("\n%s", test)
		rf, err := readFile(test)
		if err != nil {
			t.Fatal(err)
		}
		f, err := macho.NewFile(rf)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		info := f.Section("__debug_info")
		if info == nil {
			for _, sec := range f.Sections {
				t.Log(sec.Name)
			}
			t.Fatal("No such section")
		}
		abbrev := f.Section("__debug_abbrev")
		if abbrev == nil {
			for _, sec := range f.Sections {
				t.Log(sec.Name)
			}
			t.Fatal("No such section")
		}

		data, err := info.Data()
		if err != nil {
			t.Fatal(err)
		}
		data2, err := abbrev.Data()
		if err != nil {
			t.Fatal(err)
		}

		bri := binary.BinaryReader{Reader: bytes.NewReader(data), Endianess: binary.LittleEndian}
		bra := binary.BinaryReader{Reader: bytes.NewReader(data2), Endianess: binary.LittleEndian}
		var ih InfoHeader
		if err := bri.ReadInterface(&ih); err != nil {
			t.Error(err)
			break
		} else {
			//				t.Logf("%+v", ih)
		}
		var abbr_entries []AbbrevEntry
		for {
			var abr AbbrevEntry
			if err := bra.ReadInterface(&abr); err != nil {
				break
			} else {
				abbr_entries = append(abbr_entries, abr)
			}
		}
		//		bra.Seek(ih.DebugAbbrevOffset, 0)

		for {
			var (
				ie InfoEntry
			)
			if err := bri.ReadInterface(&ie); err != nil {
				if err == io.EOF {
					break
				}
				t.Error(err)
				break
			} else {
				//				t.Logf("%+v", ih)
			}

			// if err := bra.ReadInterface(&abr); err != nil {
			// 	t.Error(err)
			// 	break
			// } else {
			// 	//				t.Logf("%+v", abr)
			// }
			ie.header = &ih
			if ie.id == 0 {
				continue
			}
			ie.ae = &abbr_entries[ie.id-1]
			ie.reader.Info = bri
			ie.reader.Abbrev = bra
			t.Logf("%s", ie.Tag())
			for _, attr := range ie.ae.Attributes {
				pos, _ := bri.Seek(0, 1)
				va := ie.Attribute(attr.Name)
				switch v := va.(type) {
				case uint:
					t.Logf("\t0x%08x %s 0x%08x", pos, attr, v)
				case uint8:
					t.Logf("\t0x%08x %s 0x%08x", pos, attr, v)
				case uint16:
					t.Logf("\t0x%08x %s 0x%08x", pos, attr, v)
				case uint32:
					t.Logf("\t0x%08x %s 0x%08x", pos, attr, v)
				case uint64:
					t.Logf("\t0x%08x %s 0x%08x", pos, attr, v)
				default:
					t.Logf("\t0x%08x %s %v", pos, attr, v)
				}
			}
			off, _ := bri.Seek(0, 1)
			t.Logf("0x%x", off)
			//			t.Log(bri.Read(16))
			//			break
			//bri.Seek(int64(ie.Length), 1)
		}
		break
	}
}
