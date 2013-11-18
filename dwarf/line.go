package dwarf

import (
	"fmt"
	"github.com/quarnster/util/encoding/binary"
)

type (
	lineEntry struct {
		file           int
		line           int
		column         int
		isa            int
		address        uint64
		is_stmt        bool
		basic_block    bool
		end_sequence   bool
		prologue_end   bool
		epilogue_begin bool
	}
	lineHeader struct {
		Header
		header_length                      int64
		minimum_instruction_length         uint8
		maximum_operations_per_instruction uint8
		default_is_stmt                    bool
		line_base                          int8
		line_range                         uint8
		opcode_base                        uint8
		standard_opcode_lengths            []uint8
		include_directories                []string
		file_names                         []fileEntry
		matrix                             []lineEntry
	}
	fileEntry struct {
		Name             string
		IncludeDirectory LEB128
		LastModified     LEB128
		Size             LEB128
	}
	state struct {
		lineEntry
		header        *lineHeader
		op_index      uint
		discriminator uint
	}
)

func (s *state) advance(advance uint) {
	addr := uint(s.header.minimum_instruction_length) * (uint(s.op_index) + advance)
	s.address += uint64(addr)
	s.op_index += uint(advance) % uint(s.header.maximum_operations_per_instruction)
}

func (s *state) execute(op DW_LNS, br *binary.BinaryReader) error {
	switch op {
	case DW_LNS_copy:
		s.header.matrix = append(s.header.matrix, s.lineEntry)
		s.discriminator = 0
		s.basic_block = false
		s.prologue_end = false
		s.epilogue_begin = false
	case DW_LNS_advance_pc:
		var advance LEB128
		if err := br.ReadInterface(&advance); err != nil {
			return err
		}
		s.advance(uint(advance))
	case DW_LNS_advance_line:
		var delta SLEB128
		if err := br.ReadInterface(&delta); err != nil {
			return err
		}
		s.line += int(delta)
	case DW_LNS_set_file:
		var arg LEB128
		if err := br.ReadInterface(&arg); err != nil {
			return err
		}
		s.file = int(arg)
	case DW_LNS_set_column:
		var arg LEB128
		if err := br.ReadInterface(&arg); err != nil {
			return err
		}
		s.column = int(arg)
	case DW_LNS_negate_stmt:
		s.is_stmt = !s.is_stmt
	case DW_LNS_set_basic_block:
		s.basic_block = true
	case DW_LNS_const_add_pc:
		// TODO???
	case DW_LNS_fixed_advance_pc:
		var arg uint16
		if err := br.ReadInterface(&arg); err != nil {
			return err
		}
		s.address += uint64(arg)
		s.op_index = 0
	case DW_LNS_set_prologue_end:
		s.prologue_end = true
	case DW_LNS_set_epilogue_begin:
		s.epilogue_begin = true
	case DW_LNS_set_isa:
		var arg LEB128
		if err := br.ReadInterface(&arg); err != nil {
			return err
		}
		s.isa = int(arg)
	default:
		if uint8(op) > s.header.opcode_base {
			op2 := int(uint8(op) - s.header.opcode_base)
			advance := op2 / int(s.header.line_range)
			line := (int(op2) % int(s.header.line_range)) + int(s.header.line_base)
			s.line += line
			s.advance(uint(advance))

			s.header.matrix = append(s.header.matrix, s.lineEntry)
			s.discriminator = 0
			s.epilogue_begin = false
			s.prologue_end = false
			s.basic_block = false
		}
	}
	return nil
}
func (s *state) executeLNE(op DW_LNE, length LEB128, br *binary.BinaryReader) error {
	switch op {
	case DW_LNE_end_sequence:
		s.end_sequence = true
		s.header.matrix = append(s.header.matrix, s.lineEntry)
		s.reset()
	case DW_LNE_set_address:
		var err error
		if length == 8 {
			s.address, err = br.Uint64()
		} else {
			var v uint32
			v, err = br.Uint32()
			s.address = uint64(v)
		}
		s.op_index = 0
		return err
	case DW_LNE_define_file:
		var fe fileEntry
		if err := br.ReadInterface(&fe); err != nil {
			return err
		} else {
			s.header.file_names = append(s.header.file_names, fe)
		}
	case DW_LNE_set_discriminator:
		var arg LEB128
		if err := br.ReadInterface(&arg); err != nil {
			return err
		}
		s.discriminator = uint(arg)
	default:
		panic(fmt.Errorf("%s", op))
	}
	return nil
}

func (s *state) reset() {
	s.address = 0
	s.op_index = 0
	s.file = 1
	s.line = 1
	s.column = 0
	s.is_stmt = s.header.default_is_stmt
	s.basic_block = false
	s.end_sequence = false
	s.prologue_end = false
	s.epilogue_begin = false
	s.isa = 0
	s.discriminator = 0
}

func (lh *lineHeader) Read(br *binary.BinaryReader) error {
	start, _ := br.Seek(0, 1)
	err := br.ReadInterface(&lh.Header)
	if err != nil {
		return err
	}
	if lh.is64 {
		lh.header_length, err = br.Int64()
	} else {
		var v uint32
		v, err = br.Uint32()
		lh.header_length = int64(v)
	}
	if err != nil {
		return err
	}
	if lh.minimum_instruction_length, err = br.Uint8(); err != nil {
		return err
	}
	lh.maximum_operations_per_instruction = 1
	// TODO:
	// if lh.maximum_operations_per_instruction, err = br.Uint8(); err != nil {
	// 	return err
	// }
	if err = br.ReadInterface(&lh.default_is_stmt); err != nil {
		return err
	} else if lh.line_base, err = br.Int8(); err != nil {
		return err
	} else if lh.line_range, err = br.Uint8(); err != nil {
		return err
	} else if lh.opcode_base, err = br.Uint8(); err != nil {
		return err
	}
	lh.standard_opcode_lengths = make([]uint8, lh.opcode_base-1)
	if err := br.ReadInterface(&lh.standard_opcode_lengths); err != nil {
		return err
	}
	for {
		var s string
		if err := br.ReadInterface(&s); err != nil {
			return err
		} else if s == "" {
			break
		} else {
			lh.include_directories = append(lh.include_directories, s)
		}
	}

	for {
		var f fileEntry
		pos, _ := br.Seek(0, 1)
		if v, err := br.Uint8(); err != nil {
			return err
		} else if v == '\u0000' {
			break
		} else {
			br.Seek(pos, 0)
		}
		if err := br.ReadInterface(&f); err != nil {
			return err
		} else {
			lh.file_names = append(lh.file_names, f)
		}
	}
	var s state
	s.header = lh
	s.reset()
	pos, _ := br.Seek(0, 1)
	for (pos - 4 - start) < int64(lh.Length) {
		var op DW_LNS
		if err := br.ReadInterface(&op); err != nil {
			return err
		}
		if op == 0 {
			var length LEB128
			if err := br.ReadInterface(&length); err != nil {
				return err
			}

			var op2 DW_LNE
			if err := br.ReadInterface(&op2); err != nil {
				return err
			}
			s.executeLNE(op2, length-1, br)
		} else {
			s.execute(op, br)
		}
		pos, _ = br.Seek(0, 1)
	}

	return nil
}
