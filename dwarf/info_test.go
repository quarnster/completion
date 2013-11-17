package dwarf

import (
	"bytes"
	"debug/elf"
	"debug/macho"
	"github.com/quarnster/util/encoding/binary"
	"io"
	"testing"
)

type (
	Section interface {
		Data() ([]byte, error)
	}
	sectionReader struct {
		elf   *elf.File
		macho *macho.File
	}
)

func newSectionReader(rf io.ReaderAt) (ret sectionReader, err error) {
	ret.macho, err = macho.NewFile(rf)
	if err != nil {
		ret.elf, err = elf.NewFile(rf)
	}
	return
}

func (s *sectionReader) Section(name string) Section {
	if s.elf != nil {
		if r := s.elf.Section("." + name); r == nil {
			return nil
		} else {
			return r
		}
	} else {
		if r := s.macho.Section("__" + name); r == nil {
			return nil
		} else {
			return r
		}
	}
}

func (s *sectionReader) Close() error {
	if s.elf != nil {
		return s.elf.Close()
	}
	return s.macho.Close()
}

func TestInfo(t *testing.T) {
outer:
	for _, test := range []string{"./testdata/8", "./testdata/hello", "./testdata/game.bz2", "./testdata/listener.o.bz2", "./testdata/completion.bz2", "./testdata/hello4"} {
		t.Logf("\n%s", test)
		rf, err := readFile(test)
		if err != nil {
			t.Error(err)
			continue
		}
		f, err := newSectionReader(rf)
		if err != nil {
			t.Error(err)
			continue
		}
		defer f.Close()
		info := f.Section("debug_info")
		if info == nil {
			t.Error("No info section")
			continue
		}
		data, err := info.Data()
		if err != nil {
			t.Error(err)
			continue
		}
		abbrev := f.Section("debug_abbrev")
		if abbrev == nil {
			t.Error("No abbrev section")
			continue
		}
		data2, err := abbrev.Data()
		if err != nil {
			t.Error(err)
			continue
		}
		var (
			datastr []byte
		)
		str := f.Section("debug_str")
		if str == nil {
			t.Log("No str section")
		} else {
			datastr, err = str.Data()
			if err != nil {
				t.Log(err)
			}
		}

		bri := binary.BinaryReader{Reader: bytes.NewReader(data), Endianess: binary.LittleEndian}
		bra := binary.BinaryReader{Reader: bytes.NewReader(data2), Endianess: binary.LittleEndian}
		brs := binary.BinaryReader{Reader: bytes.NewReader(datastr), Endianess: binary.LittleEndian}
		var ih InfoHeader
		nextHeader, _ := bri.Seek(0, 1)
		var abbr_entries []AbbrevEntry
		for {
			off, _ := bri.Seek(0, 1)
			if off >= nextHeader {
				bri.Seek(nextHeader, 0)
				t.Log(bri.Read(16))
				bri.Seek(nextHeader, 0)
				if err := bri.ReadInterface(&ih); err != nil {
					if err != io.EOF {
						t.Error(err)
					}
					continue outer
				} else {
					t.Logf("%+v", ih)
					nextHeader += int64(ih.Length) + 4
					abbr_entries = nil
					bra.Seek(ih.DebugAbbrevOffset, 0)
					for {
						var abr AbbrevEntry
						if err := bra.ReadInterface(&abr); err != nil {
							break
						} else {
							abbr_entries = append(abbr_entries, abr)
						}
					}
				}
			}
			var (
				ie InfoEntry
			)
			if err := bri.ReadInterface(&ie); err != nil {
				if err == io.EOF {
					break
				}
				t.Error(err)
				break
			}

			ie.header = &ih
			if ie.id <= 0 {
				continue
			}
			t.Log(ie.id, len(abbr_entries))
			ie.ae = &abbr_entries[ie.id-1]
			ie.reader.Info = bri
			ie.reader.Abbrev = bra
			ie.reader.Str = brs
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
			off, _ = bri.Seek(0, 1)
			t.Logf("0x%x 0x%x", off, nextHeader)
		}
		break
	}
}
