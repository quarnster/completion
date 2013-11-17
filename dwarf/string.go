package dwarf

import "fmt"

func (dw DW_ACCESS) String() string {
	if v, ok := lut_DW_ACCESS[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_ACCESS value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_AT) String() string {
	if v, ok := lut_DW_AT[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_AT value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_ATE) String() string {
	if v, ok := lut_DW_ATE[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_ATE value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_CC) String() string {
	if v, ok := lut_DW_CC[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_CC value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_CHILDREN) String() string {
	if v, ok := lut_DW_CHILDREN[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_CHILDREN value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_DS) String() string {
	if v, ok := lut_DW_DS[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_DS value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_DSC) String() string {
	if v, ok := lut_DW_DSC[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_DSC value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_END) String() string {
	if v, ok := lut_DW_END[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_END value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_FORM) String() string {
	if v, ok := lut_DW_FORM[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_FORM value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_ID) String() string {
	if v, ok := lut_DW_ID[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_ID value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_INL) String() string {
	if v, ok := lut_DW_INL[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_INL value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_LANG) String() string {
	if v, ok := lut_DW_LANG[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_LANG value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_LNE) String() string {
	if v, ok := lut_DW_LNE[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_LNE value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_LNS) String() string {
	if v, ok := lut_DW_LNS[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_LNS value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_MACINFO) String() string {
	if v, ok := lut_DW_MACINFO[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_MACINFO value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_ORD) String() string {
	if v, ok := lut_DW_ORD[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_ORD value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_OP) String() string {
	if v, ok := lut_DW_OP[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_OP value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_TAG) String() string {
	if v, ok := lut_DW_TAG[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_TAG value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_VIRTUALITY) String() string {
	if v, ok := lut_DW_VIRTUALITY[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_VIRTUALITY value: %d (0x%x)", int(dw), int(dw))
}

func (dw DW_VIS) String() string {
	if v, ok := lut_DW_VIS[dw]; ok {
		return v
	}
	return fmt.Sprintf("Unknown DW_VIS value: %d (0x%x)", int(dw), int(dw))
}

var (
	lut_DW_ATE = map[DW_ATE]string{
		DW_ATE_address:         "DW_ATE_address",
		DW_ATE_boolean:         "DW_ATE_boolean",
		DW_ATE_complex_float:   "DW_ATE_complex_float",
		DW_ATE_float:           "DW_ATE_float",
		DW_ATE_signed:          "DW_ATE_signed",
		DW_ATE_signed_char:     "DW_ATE_signed_char",
		DW_ATE_unsigned:        "DW_ATE_unsigned",
		DW_ATE_unsigned_char:   "DW_ATE_unsigned_char",
		DW_ATE_imaginary_float: "DW_ATE_imaginary_float",
		DW_ATE_packed_decimal:  "DW_ATE_packed_decimal",
		DW_ATE_numeric_string:  "DW_ATE_numeric_string",
		DW_ATE_edited:          "DW_ATE_edited",
		DW_ATE_signed_fixed:    "DW_ATE_signed_fixed",
		DW_ATE_unsigned_fixed:  "DW_ATE_unsigned_fixed",
		DW_ATE_decimal_float:   "DW_ATE_decimal_float",
		DW_ATE_UTF:             "DW_ATE_UTF",
		DW_ATE_lo_user:         "DW_ATE_lo_user",
		DW_ATE_hi_user:         "DW_ATE_hi_user",
	}
	lut_DW_DS = map[DW_DS]string{
		DW_DS_unsigned:           "DW_DS_unsigned",
		DW_DS_leading_overpunch:  "DW_DS_leading_overpunch",
		DW_DS_trailing_overpunch: "DW_DS_trailing_overpunch",
		DW_DS_leading_separate:   "DW_DS_leading_separate",
		DW_DS_trailing_separate:  "DW_DS_trailing_separate",
	}
	lut_DW_END = map[DW_END]string{
		DW_END_default: "DW_END_default",
		DW_END_big:     "DW_END_big",
		DW_END_little:  "DW_END_little",
		DW_END_lo_user: "DW_END_lo_user",
		DW_END_hi_user: "DW_END_hi_user",
	}
	lut_DW_ACCESS = map[DW_ACCESS]string{
		DW_ACCESS_public:    "DW_ACCESS_public",
		DW_ACCESS_protected: "DW_ACCESS_protected",
		DW_ACCESS_private:   "DW_ACCESS_private",
	}
	lut_DW_VIS = map[DW_VIS]string{
		DW_VIS_local:     "DW_VIS_local",
		DW_VIS_exported:  "DW_VIS_exported",
		DW_VIS_qualified: "DW_VIS_qualified",
	}
	lut_DW_VIRTUALITY = map[DW_VIRTUALITY]string{
		DW_VIRTUALITY_none:         "DW_VIRTUALITY_none",
		DW_VIRTUALITY_virtual:      "DW_VIRTUALITY_virtual",
		DW_VIRTUALITY_pure_virtual: "DW_VIRTUALITY_pure_virtual",
	}
	lut_DW_LANG = map[DW_LANG]string{
		DW_LANG_C89:            "DW_LANG_C89",
		DW_LANG_C:              "DW_LANG_C",
		DW_LANG_Ada83:          "DW_LANG_Ada83",
		DW_LANG_C_plus_plus:    "DW_LANG_C_plus_plus",
		DW_LANG_Cobol74:        "DW_LANG_Cobol74",
		DW_LANG_Cobol851:       "DW_LANG_Cobol851",
		DW_LANG_Fortran77:      "DW_LANG_Fortran77",
		DW_LANG_Fortran90:      "DW_LANG_Fortran90",
		DW_LANG_Pascal83:       "DW_LANG_Pascal83",
		DW_LANG_Modula2:        "DW_LANG_Modula2",
		DW_LANG_Java:           "DW_LANG_Java",
		DW_LANG_C99:            "DW_LANG_C99",
		DW_LANG_Ada95:          "DW_LANG_Ada95",
		DW_LANG_Fortran95:      "DW_LANG_Fortran95",
		DW_LANG_PLI:            "DW_LANG_PLI",
		DW_LANG_ObjC:           "DW_LANG_ObjC",
		DW_LANG_ObjC_plus_plus: "DW_LANG_ObjC_plus_plus",
		DW_LANG_UPC:            "DW_LANG_UPC",
		DW_LANG_D:              "DW_LANG_D",
		DW_LANG_Python:         "DW_LANG_Python",
		DW_LANG_OpenCL:         "DW_LANG_OpenCL",
		DW_LANG_Go:             "DW_LANG_Go",
		DW_LANG_Modula3:        "DW_LANG_Modula3",
		DW_LANG_Haskell:        "DW_LANG_Haskell",
		DW_LANG_lo_user:        "DW_LANG_lo_user",
		DW_LANG_hi_user:        "DW_LANG_hi_user",
		DW_LANG_Mips_Assembler: "DW_LANG_Mips_Assembler",
	}
	lut_DW_ID = map[DW_ID]string{
		DW_ID_case_sensitive:   "DW_ID_case_sensitive",
		DW_ID_up_case:          "DW_ID_up_case",
		DW_ID_down_case:        "DW_ID_down_case",
		DW_ID_case_insensitive: "DW_ID_case_insensitive",
	}
	lut_DW_CC = map[DW_CC]string{
		DW_CC_normal:  "DW_CC_normal",
		DW_CC_program: "DW_CC_program",
		DW_CC_nocall:  "DW_CC_nocall",
		DW_CC_lo_user: "DW_CC_lo_user",
		DW_CC_hi_user: "DW_CC_hi_user",
	}
	lut_DW_INL = map[DW_INL]string{
		DW_INL_not_inlined:          "DW_INL_not_inlined",
		DW_INL_inlined:              "DW_INL_inlined",
		DW_INL_declared_not_inlined: "DW_INL_declared_not_inlined",
		DW_INL_declared_inlined:     "DW_INL_declared_inlined",
	}
	lut_DW_ORD = map[DW_ORD]string{
		DW_ORD_row_major: "DW_ORD_row_major",
		DW_ORD_col_major: "DW_ORD_col_major",
	}
	lut_DW_DSC = map[DW_DSC]string{
		DW_DSC_label: "DW_DSC_label",
		DW_DSC_range: "DW_DSC_range",
	}
	lut_DW_LNS = map[DW_LNS]string{
		DW_LNS_copy:               "DW_LNS_copy",
		DW_LNS_advance_pc:         "DW_LNS_advance_pc",
		DW_LNS_advance_line:       "DW_LNS_advance_line",
		DW_LNS_set_file:           "DW_LNS_set_file",
		DW_LNS_set_column:         "DW_LNS_set_column",
		DW_LNS_negate_stmt:        "DW_LNS_negate_stmt",
		DW_LNS_set_basic_block:    "DW_LNS_set_basic_block",
		DW_LNS_const_add_pc:       "DW_LNS_const_add_pc",
		DW_LNS_fixed_advance_pc:   "DW_LNS_fixed_advance_pc",
		DW_LNS_set_prologue_end:   "DW_LNS_set_prologue_end",
		DW_LNS_set_epilogue_begin: "DW_LNS_set_epilogue_begin",
		DW_LNS_set_isa:            "DW_LNS_set_isa",
	}
	lut_DW_LNE = map[DW_LNE]string{
		DW_LNE_end_sequence:      "DW_LNE_end_sequence",
		DW_LNE_set_address:       "DW_LNE_set_address",
		DW_LNE_define_file:       "DW_LNE_define_file",
		DW_LNE_set_discriminator: "DW_LNE_set_discriminator",
		DW_LNE_lo_user:           "DW_LNE_lo_user",
		DW_LNE_hi_user:           "DW_LNE_hi_user",
	}
	lut_DW_MACINFO = map[DW_MACINFO]string{
		DW_MACINFO_define:     "DW_MACINFO_define",
		DW_MACINFO_undef:      "DW_MACINFO_undef",
		DW_MACINFO_start_file: "DW_MACINFO_start_file",
		DW_MACINFO_end_file:   "DW_MACINFO_end_file",
		DW_MACINFO_vendor_ext: "DW_MACINFO_vendor_ext",
	}
	lut_DW_TAG = map[DW_TAG]string{
		DW_TAG_array_type:                  "DW_TAG_array_type",
		DW_TAG_class_type:                  "DW_TAG_class_type",
		DW_TAG_entry_point:                 "DW_TAG_entry_point",
		DW_TAG_enumeration_type:            "DW_TAG_enumeration_type",
		DW_TAG_formal_parameter:            "DW_TAG_formal_parameter",
		DW_TAG_imported_declaration:        "DW_TAG_imported_declaration",
		DW_TAG_label:                       "DW_TAG_label",
		DW_TAG_lexical_block:               "DW_TAG_lexical_block",
		DW_TAG_member:                      "DW_TAG_member",
		DW_TAG_pointer_type:                "DW_TAG_pointer_type",
		DW_TAG_reference_type:              "DW_TAG_reference_type",
		DW_TAG_compile_unit:                "DW_TAG_compile_unit",
		DW_TAG_string_type:                 "DW_TAG_string_type",
		DW_TAG_structure_type:              "DW_TAG_structure_type",
		DW_TAG_subroutine_type:             "DW_TAG_subroutine_type",
		DW_TAG_typedef:                     "DW_TAG_typedef",
		DW_TAG_union_type:                  "DW_TAG_union_type",
		DW_TAG_unspecified_parameters:      "DW_TAG_unspecified_parameters",
		DW_TAG_variant:                     "DW_TAG_variant",
		DW_TAG_common_block:                "DW_TAG_common_block",
		DW_TAG_common_inclusion:            "DW_TAG_common_inclusion",
		DW_TAG_inheritance:                 "DW_TAG_inheritance",
		DW_TAG_inlined_subroutine:          "DW_TAG_inlined_subroutine",
		DW_TAG_module:                      "DW_TAG_module",
		DW_TAG_ptr_to_member_type:          "DW_TAG_ptr_to_member_type",
		DW_TAG_set_type:                    "DW_TAG_set_type",
		DW_TAG_subrange_type:               "DW_TAG_subrange_type",
		DW_TAG_with_stmt:                   "DW_TAG_with_stmt",
		DW_TAG_access_declaration:          "DW_TAG_access_declaration",
		DW_TAG_base_type:                   "DW_TAG_base_type",
		DW_TAG_catch_block:                 "DW_TAG_catch_block",
		DW_TAG_const_type:                  "DW_TAG_const_type",
		DW_TAG_constant:                    "DW_TAG_constant",
		DW_TAG_enumerator:                  "DW_TAG_enumerator",
		DW_TAG_file_type:                   "DW_TAG_file_type",
		DW_TAG_friend:                      "DW_TAG_friend",
		DW_TAG_namelist:                    "DW_TAG_namelist",
		DW_TAG_namelist_item:               "DW_TAG_namelist_item",
		DW_TAG_packed_type:                 "DW_TAG_packed_type",
		DW_TAG_subprogram:                  "DW_TAG_subprogram",
		DW_TAG_template_type_parameter:     "DW_TAG_template_type_parameter",
		DW_TAG_template_value_parameter:    "DW_TAG_template_value_parameter",
		DW_TAG_thrown_type:                 "DW_TAG_thrown_type",
		DW_TAG_try_block:                   "DW_TAG_try_block",
		DW_TAG_variant_part:                "DW_TAG_variant_part",
		DW_TAG_variable:                    "DW_TAG_variable",
		DW_TAG_volatile_type:               "DW_TAG_volatile_type",
		DW_TAG_dwarf_procedure:             "DW_TAG_dwarf_procedure",
		DW_TAG_restrict_type:               "DW_TAG_restrict_type",
		DW_TAG_interface_type:              "DW_TAG_interface_type",
		DW_TAG_namespace:                   "DW_TAG_namespace",
		DW_TAG_imported_module:             "DW_TAG_imported_module",
		DW_TAG_unspecified_type:            "DW_TAG_unspecified_type",
		DW_TAG_partial_unit:                "DW_TAG_partial_unit",
		DW_TAG_imported_unit:               "DW_TAG_imported_unit",
		DW_TAG_condition:                   "DW_TAG_condition",
		DW_TAG_shared_type:                 "DW_TAG_shared_type",
		DW_TAG_type_unit:                   "DW_TAG_type_unit",
		DW_TAG_rvalue_reference_type:       "DW_TAG_rvalue_reference_type",
		DW_TAG_template_alias:              "DW_TAG_template_alias",
		DW_TAG_lo_user:                     "DW_TAG_lo_user",
		DW_TAG_hi_user:                     "DW_TAG_hi_user",
		DW_TAG_format_label:                "DW_TAG_format_label",
		DW_TAG_function_template:           "DW_TAG_function_template",
		DW_TAG_class_template:              "DW_TAG_class_template",
		DW_TAG_GNU_BINCL:                   "DW_TAG_GNU_BINCL",
		DW_TAG_GNU_EINCL:                   "DW_TAG_GNU_EINCL",
		DW_TAG_GNU_template_template_param: "DW_TAG_GNU_template_template_param",
		DW_TAG_GNU_template_parameter_pack: "DW_TAG_GNU_template_parameter_pack",
		DW_TAG_GNU_formal_parameter_pack:   "DW_TAG_GNU_formal_parameter_pack",
		DW_TAG_GNU_call_site:               "DW_TAG_GNU_call_site",
		DW_TAG_GNU_call_site_parameter:     "DW_TAG_GNU_call_site_parameter",
		DW_TAG_APPLE_property:              "DW_TAG_APPLE_property",
	}
	lut_DW_CHILDREN = map[DW_CHILDREN]string{
		DW_CHILDREN_no:  "DW_CHILDREN_no",
		DW_CHILDREN_yes: "DW_CHILDREN_yes",
	}
	lut_DW_OP = map[DW_OP]string{
		DW_OP_addr:                "DW_OP_addr",
		DW_OP_deref:               "DW_OP_deref",
		DW_OP_const1u:             "DW_OP_const1u",
		DW_OP_const1s:             "DW_OP_const1s",
		DW_OP_const2u:             "DW_OP_const2u",
		DW_OP_const2s:             "DW_OP_const2s",
		DW_OP_const4u:             "DW_OP_const4u",
		DW_OP_const4s:             "DW_OP_const4s",
		DW_OP_const8u:             "DW_OP_const8u",
		DW_OP_const8s:             "DW_OP_const8s",
		DW_OP_constu:              "DW_OP_constu",
		DW_OP_consts:              "DW_OP_consts",
		DW_OP_dup:                 "DW_OP_dup",
		DW_OP_drop:                "DW_OP_drop",
		DW_OP_over:                "DW_OP_over",
		DW_OP_pick:                "DW_OP_pick",
		DW_OP_swap:                "DW_OP_swap",
		DW_OP_rot:                 "DW_OP_rot",
		DW_OP_xderef:              "DW_OP_xderef",
		DW_OP_abs:                 "DW_OP_abs",
		DW_OP_and:                 "DW_OP_and",
		DW_OP_div:                 "DW_OP_div",
		DW_OP_minus:               "DW_OP_minus",
		DW_OP_mod:                 "DW_OP_mod",
		DW_OP_mul:                 "DW_OP_mul",
		DW_OP_neg:                 "DW_OP_neg",
		DW_OP_not:                 "DW_OP_not",
		DW_OP_or:                  "DW_OP_or",
		DW_OP_plus:                "DW_OP_plus",
		DW_OP_plus_uconst:         "DW_OP_plus_uconst",
		DW_OP_shl:                 "DW_OP_shl",
		DW_OP_shr:                 "DW_OP_shr",
		DW_OP_shra:                "DW_OP_shra",
		DW_OP_xor:                 "DW_OP_xor",
		DW_OP_skip:                "DW_OP_skip",
		DW_OP_bra:                 "DW_OP_bra",
		DW_OP_eq:                  "DW_OP_eq",
		DW_OP_ge:                  "DW_OP_ge",
		DW_OP_gt:                  "DW_OP_gt",
		DW_OP_le:                  "DW_OP_le",
		DW_OP_lt:                  "DW_OP_lt",
		DW_OP_ne:                  "DW_OP_ne",
		DW_OP_lit0:                "DW_OP_lit0",
		DW_OP_lit1:                "DW_OP_lit1",
		DW_OP_lit2:                "DW_OP_lit2",
		DW_OP_lit3:                "DW_OP_lit3",
		DW_OP_lit4:                "DW_OP_lit4",
		DW_OP_lit5:                "DW_OP_lit5",
		DW_OP_lit6:                "DW_OP_lit6",
		DW_OP_lit7:                "DW_OP_lit7",
		DW_OP_lit8:                "DW_OP_lit8",
		DW_OP_lit9:                "DW_OP_lit9",
		DW_OP_lit10:               "DW_OP_lit10",
		DW_OP_lit11:               "DW_OP_lit11",
		DW_OP_lit12:               "DW_OP_lit12",
		DW_OP_lit13:               "DW_OP_lit13",
		DW_OP_lit14:               "DW_OP_lit14",
		DW_OP_lit15:               "DW_OP_lit15",
		DW_OP_lit16:               "DW_OP_lit16",
		DW_OP_lit17:               "DW_OP_lit17",
		DW_OP_lit18:               "DW_OP_lit18",
		DW_OP_lit19:               "DW_OP_lit19",
		DW_OP_lit20:               "DW_OP_lit20",
		DW_OP_lit21:               "DW_OP_lit21",
		DW_OP_lit22:               "DW_OP_lit22",
		DW_OP_lit23:               "DW_OP_lit23",
		DW_OP_lit24:               "DW_OP_lit24",
		DW_OP_lit25:               "DW_OP_lit25",
		DW_OP_lit26:               "DW_OP_lit26",
		DW_OP_lit27:               "DW_OP_lit27",
		DW_OP_lit28:               "DW_OP_lit28",
		DW_OP_lit29:               "DW_OP_lit29",
		DW_OP_lit30:               "DW_OP_lit30",
		DW_OP_lit31:               "DW_OP_lit31",
		DW_OP_reg0:                "DW_OP_reg0",
		DW_OP_reg1:                "DW_OP_reg1",
		DW_OP_reg2:                "DW_OP_reg2",
		DW_OP_reg3:                "DW_OP_reg3",
		DW_OP_reg4:                "DW_OP_reg4",
		DW_OP_reg5:                "DW_OP_reg5",
		DW_OP_reg6:                "DW_OP_reg6",
		DW_OP_reg7:                "DW_OP_reg7",
		DW_OP_reg8:                "DW_OP_reg8",
		DW_OP_reg9:                "DW_OP_reg9",
		DW_OP_reg10:               "DW_OP_reg10",
		DW_OP_reg11:               "DW_OP_reg11",
		DW_OP_reg12:               "DW_OP_reg12",
		DW_OP_reg13:               "DW_OP_reg13",
		DW_OP_reg14:               "DW_OP_reg14",
		DW_OP_reg15:               "DW_OP_reg15",
		DW_OP_reg16:               "DW_OP_reg16",
		DW_OP_reg17:               "DW_OP_reg17",
		DW_OP_reg18:               "DW_OP_reg18",
		DW_OP_reg19:               "DW_OP_reg19",
		DW_OP_reg20:               "DW_OP_reg20",
		DW_OP_reg21:               "DW_OP_reg21",
		DW_OP_reg22:               "DW_OP_reg22",
		DW_OP_reg23:               "DW_OP_reg23",
		DW_OP_reg24:               "DW_OP_reg24",
		DW_OP_reg25:               "DW_OP_reg25",
		DW_OP_reg26:               "DW_OP_reg26",
		DW_OP_reg27:               "DW_OP_reg27",
		DW_OP_reg28:               "DW_OP_reg28",
		DW_OP_reg29:               "DW_OP_reg29",
		DW_OP_reg30:               "DW_OP_reg30",
		DW_OP_reg31:               "DW_OP_reg31",
		DW_OP_breg1:               "DW_OP_breg1",
		DW_OP_breg2:               "DW_OP_breg2",
		DW_OP_breg3:               "DW_OP_breg3",
		DW_OP_breg4:               "DW_OP_breg4",
		DW_OP_breg5:               "DW_OP_breg5",
		DW_OP_breg6:               "DW_OP_breg6",
		DW_OP_breg7:               "DW_OP_breg7",
		DW_OP_breg8:               "DW_OP_breg8",
		DW_OP_breg9:               "DW_OP_breg9",
		DW_OP_breg10:              "DW_OP_breg10",
		DW_OP_breg11:              "DW_OP_breg11",
		DW_OP_breg12:              "DW_OP_breg12",
		DW_OP_breg13:              "DW_OP_breg13",
		DW_OP_breg14:              "DW_OP_breg14",
		DW_OP_breg15:              "DW_OP_breg15",
		DW_OP_breg16:              "DW_OP_breg16",
		DW_OP_breg17:              "DW_OP_breg17",
		DW_OP_breg18:              "DW_OP_breg18",
		DW_OP_breg19:              "DW_OP_breg19",
		DW_OP_breg20:              "DW_OP_breg20",
		DW_OP_breg21:              "DW_OP_breg21",
		DW_OP_breg22:              "DW_OP_breg22",
		DW_OP_breg23:              "DW_OP_breg23",
		DW_OP_breg24:              "DW_OP_breg24",
		DW_OP_breg25:              "DW_OP_breg25",
		DW_OP_breg26:              "DW_OP_breg26",
		DW_OP_breg27:              "DW_OP_breg27",
		DW_OP_breg28:              "DW_OP_breg28",
		DW_OP_breg29:              "DW_OP_breg29",
		DW_OP_breg30:              "DW_OP_breg30",
		DW_OP_breg31:              "DW_OP_breg31",
		DW_OP_regx:                "DW_OP_regx",
		DW_OP_fbreg:               "DW_OP_fbreg",
		DW_OP_bregx:               "DW_OP_bregx",
		DW_OP_piece:               "DW_OP_piece",
		DW_OP_deref_size:          "DW_OP_deref_size",
		DW_OP_xderef_size:         "DW_OP_xderef_size",
		DW_OP_nop:                 "DW_OP_nop",
		DW_OP_push_object_address: "DW_OP_push_object_address",
		DW_OP_call2:               "DW_OP_call2",
		DW_OP_call4:               "DW_OP_call4",
		DW_OP_call_ref:            "DW_OP_call_ref",
		DW_OP_form_tls_address:    "DW_OP_form_tls_address",
		DW_OP_call_frame_cfa:      "DW_OP_call_frame_cfa",
		DW_OP_bit_piece:           "DW_OP_bit_piece",
		DW_OP_implicit_value:      "DW_OP_implicit_value",
		DW_OP_stack_value:         "DW_OP_stack_value",
		DW_OP_lo_user:             "DW_OP_lo_user",
		DW_OP_hi_user:             "DW_OP_hi_user",
	}
	lut_DW_FORM = map[DW_FORM]string{
		DW_FORM_addr:         "DW_FORM_addr",
		DW_FORM_block2:       "DW_FORM_block2",
		DW_FORM_block4:       "DW_FORM_block4",
		DW_FORM_data2:        "DW_FORM_data2",
		DW_FORM_data4:        "DW_FORM_data4",
		DW_FORM_data8:        "DW_FORM_data8",
		DW_FORM_string:       "DW_FORM_string",
		DW_FORM_block:        "DW_FORM_block",
		DW_FORM_block1:       "DW_FORM_block1",
		DW_FORM_data1:        "DW_FORM_data1",
		DW_FORM_flag:         "DW_FORM_flag",
		DW_FORM_sdata:        "DW_FORM_sdata",
		DW_FORM_strp:         "DW_FORM_strp",
		DW_FORM_udata:        "DW_FORM_udata",
		DW_FORM_ref_addr:     "DW_FORM_ref_addr",
		DW_FORM_ref1:         "DW_FORM_ref1",
		DW_FORM_ref2:         "DW_FORM_ref2",
		DW_FORM_ref4:         "DW_FORM_ref4",
		DW_FORM_ref8:         "DW_FORM_ref8",
		DW_FORM_ref_udata:    "DW_FORM_ref_udata",
		DW_FORM_indirect:     "DW_FORM_indirect",
		DW_FORM_sec_offset:   "DW_FORM_sec_offset",
		DW_FORM_exprloc:      "DW_FORM_exprloc",
		DW_FORM_flag_present: "DW_FORM_flag_present",
		DW_FORM_ref_sig8:     "DW_FORM_ref_sig8",
	}
	lut_DW_AT = map[DW_AT]string{
		DW_AT_sibling:                        "DW_AT_sibling",
		DW_AT_location:                       "DW_AT_location",
		DW_AT_name:                           "DW_AT_name",
		DW_AT_ordering:                       "DW_AT_ordering",
		DW_AT_byte_size:                      "DW_AT_byte_size",
		DW_AT_bit_offset:                     "DW_AT_bit_offset",
		DW_AT_bit_size:                       "DW_AT_bit_size",
		DW_AT_stmt_list:                      "DW_AT_stmt_list",
		DW_AT_low_pc:                         "DW_AT_low_pc",
		DW_AT_high_pc:                        "DW_AT_high_pc",
		DW_AT_language:                       "DW_AT_language",
		DW_AT_discr:                          "DW_AT_discr",
		DW_AT_discr_value:                    "DW_AT_discr_value",
		DW_AT_visibility:                     "DW_AT_visibility",
		DW_AT_import:                         "DW_AT_import",
		DW_AT_string_length:                  "DW_AT_string_length",
		DW_AT_common_reference:               "DW_AT_common_reference",
		DW_AT_comp_dir:                       "DW_AT_comp_dir",
		DW_AT_const_value:                    "DW_AT_const_value",
		DW_AT_containing_type:                "DW_AT_containing_type",
		DW_AT_default_value:                  "DW_AT_default_value",
		DW_AT_inline:                         "DW_AT_inline",
		DW_AT_is_optional:                    "DW_AT_is_optional",
		DW_AT_lower_bound:                    "DW_AT_lower_bound",
		DW_AT_producer:                       "DW_AT_producer",
		DW_AT_prototyped:                     "DW_AT_prototyped",
		DW_AT_return_addr:                    "DW_AT_return_addr",
		DW_AT_start_scope:                    "DW_AT_start_scope",
		DW_AT_bit_stride:                     "DW_AT_bit_stride",
		DW_AT_upper_bound:                    "DW_AT_upper_bound",
		DW_AT_abstract_origin:                "DW_AT_abstract_origin",
		DW_AT_accessibility:                  "DW_AT_accessibility",
		DW_AT_address_class:                  "DW_AT_address_class",
		DW_AT_artificial:                     "DW_AT_artificial",
		DW_AT_base_types:                     "DW_AT_base_types",
		DW_AT_calling_convention:             "DW_AT_calling_convention",
		DW_AT_count:                          "DW_AT_count",
		DW_AT_data_member_location:           "DW_AT_data_member_location",
		DW_AT_decl_column:                    "DW_AT_decl_column",
		DW_AT_decl_file:                      "DW_AT_decl_file",
		DW_AT_decl_line:                      "DW_AT_decl_line",
		DW_AT_declaration:                    "DW_AT_declaration",
		DW_AT_discr_list:                     "DW_AT_discr_list",
		DW_AT_encoding:                       "DW_AT_encoding",
		DW_AT_external:                       "DW_AT_external",
		DW_AT_frame_base:                     "DW_AT_frame_base",
		DW_AT_friend:                         "DW_AT_friend",
		DW_AT_identifier_case:                "DW_AT_identifier_case",
		DW_AT_macro_info:                     "DW_AT_macro_info",
		DW_AT_namelist_item:                  "DW_AT_namelist_item",
		DW_AT_priority:                       "DW_AT_priority",
		DW_AT_segment:                        "DW_AT_segment",
		DW_AT_specification:                  "DW_AT_specification",
		DW_AT_static_link:                    "DW_AT_static_link",
		DW_AT_type:                           "DW_AT_type",
		DW_AT_use_location:                   "DW_AT_use_location",
		DW_AT_variable_parameter:             "DW_AT_variable_parameter",
		DW_AT_virtuality:                     "DW_AT_virtuality",
		DW_AT_vtable_elem_location:           "DW_AT_vtable_elem_location",
		DW_AT_allocated:                      "DW_AT_allocated",
		DW_AT_associated:                     "DW_AT_associated",
		DW_AT_data_location:                  "DW_AT_data_location",
		DW_AT_byte_stride:                    "DW_AT_byte_stride",
		DW_AT_entry_pc:                       "DW_AT_entry_pc",
		DW_AT_use_UTF8:                       "DW_AT_use_UTF8",
		DW_AT_extension:                      "DW_AT_extension",
		DW_AT_ranges:                         "DW_AT_ranges",
		DW_AT_trampoline:                     "DW_AT_trampoline",
		DW_AT_call_column:                    "DW_AT_call_column",
		DW_AT_call_file:                      "DW_AT_call_file",
		DW_AT_call_line:                      "DW_AT_call_line",
		DW_AT_description:                    "DW_AT_description",
		DW_AT_binary_scale:                   "DW_AT_binary_scale",
		DW_AT_decimal_scale:                  "DW_AT_decimal_scale",
		DW_AT_small:                          "DW_AT_small",
		DW_AT_decimal_sign:                   "DW_AT_decimal_sign",
		DW_AT_digit_count:                    "DW_AT_digit_count",
		DW_AT_picture_string:                 "DW_AT_picture_string",
		DW_AT_mutable:                        "DW_AT_mutable",
		DW_AT_threads_scaled:                 "DW_AT_threads_scaled",
		DW_AT_explicit:                       "DW_AT_explicit",
		DW_AT_object_pointer:                 "DW_AT_object_pointer",
		DW_AT_endianity:                      "DW_AT_endianity",
		DW_AT_elemental:                      "DW_AT_elemental",
		DW_AT_pure:                           "DW_AT_pure",
		DW_AT_recursive:                      "DW_AT_recursive",
		DW_AT_signature:                      "DW_AT_signature",
		DW_AT_main_subprogram:                "DW_AT_main_subprogram",
		DW_AT_data_bit_offset:                "DW_AT_data_bit_offset",
		DW_AT_const_expr:                     "DW_AT_const_expr",
		DW_AT_enum_class:                     "DW_AT_enum_class",
		DW_AT_linkage_name:                   "DW_AT_linkage_name",
		DW_AT_MIPS_loop_begin:                "DW_AT_MIPS_loop_begin",
		DW_AT_MIPS_tail_loop_begin:           "DW_AT_MIPS_tail_loop_begin",
		DW_AT_MIPS_epilog_begin:              "DW_AT_MIPS_epilog_begin",
		DW_AT_MIPS_loop_unroll_factor:        "DW_AT_MIPS_loop_unroll_factor",
		DW_AT_MIPS_software_pipeline_depth:   "DW_AT_MIPS_software_pipeline_depth",
		DW_AT_MIPS_linkage_name:              "DW_AT_MIPS_linkage_name",
		DW_AT_MIPS_stride:                    "DW_AT_MIPS_stride",
		DW_AT_MIPS_abstract_name:             "DW_AT_MIPS_abstract_name",
		DW_AT_MIPS_clone_origin:              "DW_AT_MIPS_clone_origin",
		DW_AT_MIPS_has_inlines:               "DW_AT_MIPS_has_inlines",
		DW_AT_MIPS_stride_byte:               "DW_AT_MIPS_stride_byte",
		DW_AT_MIPS_stride_elem:               "DW_AT_MIPS_stride_elem",
		DW_AT_MIPS_ptr_dopetype:              "DW_AT_MIPS_ptr_dopetype",
		DW_AT_MIPS_allocatable_dopetype:      "DW_AT_MIPS_allocatable_dopetype",
		DW_AT_MIPS_assumed_shape_dopetype:    "DW_AT_MIPS_assumed_shape_dopetype",
		DW_AT_MIPS_assumed_size:              "DW_AT_MIPS_assumed_size",
		DW_AT_sf_names:                       "DW_AT_sf_names",
		DW_AT_src_info:                       "DW_AT_src_info",
		DW_AT_mac_info:                       "DW_AT_mac_info",
		DW_AT_src_coords:                     "DW_AT_src_coords",
		DW_AT_body_begin:                     "DW_AT_body_begin",
		DW_AT_body_end:                       "DW_AT_body_end",
		DW_AT_GNU_vector:                     "DW_AT_GNU_vector",
		DW_AT_GNU_guarded_by:                 "DW_AT_GNU_guarded_by",
		DW_AT_GNU_pt_guarded_by:              "DW_AT_GNU_pt_guarded_by",
		DW_AT_GNU_guarded:                    "DW_AT_GNU_guarded",
		DW_AT_GNU_pt_guarded:                 "DW_AT_GNU_pt_guarded",
		DW_AT_GNU_locks_excluded:             "DW_AT_GNU_locks_excluded",
		DW_AT_GNU_exclusive_locks_required:   "DW_AT_GNU_exclusive_locks_required",
		DW_AT_GNU_shared_locks_required:      "DW_AT_GNU_shared_locks_required",
		DW_AT_GNU_odr_signature:              "DW_AT_GNU_odr_signature",
		DW_AT_GNU_template_name:              "DW_AT_GNU_template_name",
		DW_AT_GNU_call_site_value:            "DW_AT_GNU_call_site_value",
		DW_AT_GNU_call_site_data_value:       "DW_AT_GNU_call_site_data_value",
		DW_AT_GNU_call_site_target:           "DW_AT_GNU_call_site_target",
		DW_AT_GNU_call_site_target_clobbered: "DW_AT_GNU_call_site_target_clobbered",
		DW_AT_GNU_tail_call:                  "DW_AT_GNU_tail_call",
		DW_AT_GNU_all_tail_call_sites:        "DW_AT_GNU_all_tail_call_sites",
		DW_AT_GNU_all_call_sites:             "DW_AT_GNU_all_call_sites",
		DW_AT_GNU_all_source_call_sites:      "DW_AT_GNU_all_source_call_sites",
		DW_AT_GNU_macros:                     "DW_AT_GNU_macros",
		DW_AT_APPLE_optimized:                "DW_AT_APPLE_optimized",
		DW_AT_APPLE_flags:                    "DW_AT_APPLE_flags",
		DW_AT_APPLE_isa:                      "DW_AT_APPLE_isa",
		DW_AT_APPLE_block:                    "DW_AT_APPLE_block",
		DW_AT_APPLE_major_runtime_vers:       "DW_AT_APPLE_major_runtime_vers",
		DW_AT_APPLE_runtime_class:            "DW_AT_APPLE_runtime_class",
		DW_AT_APPLE_omit_frame_ptr:           "DW_AT_APPLE_omit_frame_ptr",
		DW_AT_APPLE_property_name:            "DW_AT_APPLE_property_name",
		DW_AT_APPLE_property_getter:          "DW_AT_APPLE_property_getter",
		DW_AT_APPLE_property_setter:          "DW_AT_APPLE_property_setter",
		DW_AT_APPLE_property_attribute:       "DW_AT_APPLE_property_attribute",
		DW_AT_APPLE_objc_complete_type:       "DW_AT_APPLE_objc_complete_type",
		DW_AT_APPLE_property:                 "DW_AT_APPLE_property",
		DW_AT_lo_user:                        "DW_AT_lo_user",
		DW_AT_hi_user:                        "DW_AT_hi_user",
	}
)
