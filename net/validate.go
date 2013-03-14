package net

import (
	"bytes"
	"fmt"
	errors "github.com/quarnster/completion/util/errors"
)

func (coff *coff_file_header) Validate() error {
	if bytes.Compare(coff.Magic, pe_magic) != 0 {
		return errors.New(fmt.Sprintf("PE Magic mismatch: %v", coff.Magic))
	}
	return nil
}

func (o *optional_header) Validate() error {
	if o.Magic != pe32 && o.Magic != pe32p {
		return errors.New(fmt.Sprintf("Unkown optional header magic: %x", o.Magic))
	}
	if len(o.RVAS) != 16 || o.RVAS[14].VirtualAddress == 0 || o.RVAS[14].Size == 0 {
		return ErrNotAssembly
	}
	return nil
}

func (s *section_table) Validate() error {
	if s.Name[0] != '.' {
		return errors.New(fmt.Sprintf("This does not appear to be a valid section header: %#v", s))
	}
	return nil
}

func (cor20 *image_cor20) Validate() error {
	if cor20.MetaData.VirtualAddress == 0 || cor20.MetaData.Size == 0 {
		return ErrNotAssembly
	}
	return nil
}

func (m *MetadataHeader) Validate() error {
	if m.Signature != metadata_signature || m.MajorVersion != 1 || m.MinorVersion != 1 || m.Reserved != 0 {
		return ErrMetadata
	}
	return nil
}

func (s *stream_header) Validate() error {
	if s.Name[0] != '#' {
		return errors.New(fmt.Sprintf("This does not appear to be a valid stream header: %#v", s))
	}
	return nil
}

func (h *hash_tilde_stream_header) Validate() error {
	if h.Reserved != 0 || h.MajorVersion != 2 || h.MinorVersion != 0 /*|| h.Reserved2 != 1*/ { //TODO: Hmm spec says Reserved2 should be 1, but it appears to be 0x10?
		return errors.New(fmt.Sprintf("This does not appear to be a valid #~ stream header: %#v", h))
	}
	return nil
}
