package dwarf

import (
	"bytes"
	"debug/elf"
	"debug/macho"
	"fmt"
	"github.com/quarnster/completion/util"
	"github.com/quarnster/util/encoding/binary"
	"io"
	"io/ioutil"
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

func (s *sectionReader) Reader(name string) (ret binary.BinaryReader) {
	var data []byte
	ret.Endianess = binary.LittleEndian
	if sec := s.Section(name); sec != nil {
		data, _ = sec.Data()
	}
	ret.Reader = bytes.NewReader(data)
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
		bri := f.Reader("debug_info")
		bra := f.Reader("debug_abbrev")
		brs := f.Reader("debug_str")
		var ih InfoHeader
		nextHeader, _ := bri.Seek(0, 1)
		var abbr_entries []AbbrevEntry
		buf := bytes.NewBuffer(nil)
		var indent = ""
		for {
			off, _ := bri.Seek(0, 1)
			if off >= nextHeader {
				if err := bri.ReadInterface(&ih); err != nil {
					if err != io.EOF {
						t.Error(err)
					}
					break
				} else {
					buf.WriteString(fmt.Sprintf("%+v\n", ih))
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
				indent = indent[:len(indent)-2]
				continue
			}
			//			buf.WriteString(fmt.Sprintf("%v, %d\n", ie.id, len(abbr_entries)))
			ie.ae = &abbr_entries[ie.id-1]
			ie.reader.Info = bri
			ie.reader.Abbrev = bra
			ie.reader.Str = brs
			buf.WriteString(fmt.Sprintf("%s%s [%d] ", indent, ie.Tag().String(), ie.id))
			if ie.ae.Children == DW_CHILDREN_yes {
				buf.WriteRune('*')
			}
			buf.WriteRune('\n')
			indent += "  "
			for _, attr := range ie.ae.Attributes {
				pos, _ := bri.Seek(0, 1)
				va := ie.Attribute(attr.Name)
				switch v := va.(type) {
				case uint64:
					buf.WriteString(fmt.Sprintf("%s0x%08x %s 0x%08x\n", indent, pos, attr, v))
				default:
					buf.WriteString(fmt.Sprintf("%s0x%08x %s %v\n", indent, pos, attr, v))
				}
			}

			if ie.ae.Children == DW_CHILDREN_no {
				indent = indent[:len(indent)-2]
			}
			buf.WriteRune('\n')
			off, _ = bri.Seek(0, 1)
			//			buf.WriteString(fmt.Sprintf("0x%x 0x%x", off, nextHeader))
		}
		res := buf.String()
		expected := ""
		fn := test + ".info"
		if data, err := ioutil.ReadFile(fn); err == nil {
			expected = string(data)
		}
		if len(expected) <= 1 {
			t.Logf("Creating new test data: %s", fn)
			if err := ioutil.WriteFile(fn, []byte(res), 0644); err != nil {
				t.Errorf("Couldn't write test data to %s: %s", fn, err)
			}
		} else if d := util.Diff(expected, res); len(d) != 0 {
			t.Error(d)
		}
	}
}
