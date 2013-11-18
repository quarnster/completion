package dwarf

type (
	DW_ACCESS     LEB128
	DW_AT         LEB128
	DW_ATE        LEB128
	DW_CC         LEB128
	DW_CHILDREN   LEB128
	DW_DS         LEB128
	DW_DSC        LEB128
	DW_END        LEB128
	DW_FORM       LEB128
	DW_ID         LEB128
	DW_INL        LEB128
	DW_LANG       LEB128
	DW_LNE        uint8
	DW_LNS        uint8
	DW_MACINFO    LEB128
	DW_ORD        LEB128
	DW_OP         LEB128
	DW_TAG        LEB128
	DW_VIRTUALITY LEB128
	DW_VIS        LEB128
)

const (
	DW_ATE_address DW_ATE = 1 + iota
	DW_ATE_boolean
	DW_ATE_complex_float
	DW_ATE_float
	DW_ATE_signed
	DW_ATE_signed_char
	DW_ATE_unsigned
	DW_ATE_unsigned_char
	DW_ATE_imaginary_float
	DW_ATE_packed_decimal
	DW_ATE_numeric_string
	DW_ATE_edited
	DW_ATE_signed_fixed
	DW_ATE_unsigned_fixed
	DW_ATE_decimal_float
	DW_ATE_UTF
	DW_ATE_lo_user DW_ATE = 0x80
	DW_ATE_hi_user DW_ATE = 0xff
)

const (
	DW_DS_unsigned DW_DS = 1 + iota
	DW_DS_leading_overpunch
	DW_DS_trailing_overpunch
	DW_DS_leading_separate
	DW_DS_trailing_separate
)

const (
	DW_END_default DW_END = iota
	DW_END_big
	DW_END_little
	DW_END_lo_user DW_END = 0x40
	DW_END_hi_user DW_END = 0xff
)

const (
	DW_ACCESS_public DW_ACCESS = 1 + iota
	DW_ACCESS_protected
	DW_ACCESS_private
)

const (
	DW_VIS_local DW_VIS = 1 + iota
	DW_VIS_exported
	DW_VIS_qualified
)

const (
	DW_VIRTUALITY_none DW_VIRTUALITY = iota
	DW_VIRTUALITY_virtual
	DW_VIRTUALITY_pure_virtual
)

const (
	DW_LANG_C89 DW_LANG = 1 + iota
	DW_LANG_C
	DW_LANG_Ada83
	DW_LANG_C_plus_plus
	DW_LANG_Cobol74
	DW_LANG_Cobol851
	DW_LANG_Fortran77
	DW_LANG_Fortran90
	DW_LANG_Pascal83
	DW_LANG_Modula2
	DW_LANG_Java
	DW_LANG_C99
	DW_LANG_Ada95
	DW_LANG_Fortran95
	DW_LANG_PLI
	DW_LANG_ObjC
	DW_LANG_ObjC_plus_plus
	DW_LANG_UPC
	DW_LANG_D
	DW_LANG_Python
	DW_LANG_OpenCL
	DW_LANG_Go
	DW_LANG_Modula3
	DW_LANG_Haskell
	DW_LANG_lo_user DW_LANG = 0x8000
	DW_LANG_hi_user DW_LANG = 0xffff
)

const (
	DW_LANG_Mips_Assembler DW_LANG = 0x8001
)

const (
	DW_ID_case_sensitive DW_ID = iota
	DW_ID_up_case
	DW_ID_down_case
	DW_ID_case_insensitive
)

const (
	DW_CC_normal DW_CC = 1 + iota
	DW_CC_program
	DW_CC_nocall
	DW_CC_lo_user DW_CC = 0x40
	DW_CC_hi_user DW_CC = 0xff
)

const (
	DW_INL_not_inlined DW_INL = iota
	DW_INL_inlined
	DW_INL_declared_not_inlined
	DW_INL_declared_inlined
)

const (
	DW_ORD_row_major DW_ORD = iota
	DW_ORD_col_major
)

const (
	DW_DSC_label DW_DSC = iota
	DW_DSC_range
)

const (
	DW_LNS_copy DW_LNS = 1 + iota
	DW_LNS_advance_pc
	DW_LNS_advance_line
	DW_LNS_set_file
	DW_LNS_set_column
	DW_LNS_negate_stmt
	DW_LNS_set_basic_block
	DW_LNS_const_add_pc
	DW_LNS_fixed_advance_pc
	DW_LNS_set_prologue_end
	DW_LNS_set_epilogue_begin
	DW_LNS_set_isa
)

const (
	DW_LNE_end_sequence DW_LNE = 1 + iota
	DW_LNE_set_address
	DW_LNE_define_file
	DW_LNE_set_discriminator
	DW_LNE_lo_user DW_LNE = 0x80
	DW_LNE_hi_user DW_LNE = 0xff
)

const (
	DW_MACINFO_define     DW_MACINFO = 0x01
	DW_MACINFO_undef      DW_MACINFO = 0x02
	DW_MACINFO_start_file DW_MACINFO = 0x03
	DW_MACINFO_end_file   DW_MACINFO = 0x04
	DW_MACINFO_vendor_ext DW_MACINFO = 0xff
)

const (
	DW_TAG_array_type               DW_TAG = 0x01
	DW_TAG_class_type               DW_TAG = 0x02
	DW_TAG_entry_point              DW_TAG = 0x03
	DW_TAG_enumeration_type         DW_TAG = 0x04
	DW_TAG_formal_parameter         DW_TAG = 0x05
	DW_TAG_imported_declaration     DW_TAG = 0x08
	DW_TAG_label                    DW_TAG = 0x0a
	DW_TAG_lexical_block            DW_TAG = 0x0b
	DW_TAG_member                   DW_TAG = 0x0d
	DW_TAG_pointer_type             DW_TAG = 0x0f
	DW_TAG_reference_type           DW_TAG = 0x10
	DW_TAG_compile_unit             DW_TAG = 0x11
	DW_TAG_string_type              DW_TAG = 0x12
	DW_TAG_structure_type           DW_TAG = 0x13
	DW_TAG_subroutine_type          DW_TAG = 0x15
	DW_TAG_typedef                  DW_TAG = 0x16
	DW_TAG_union_type               DW_TAG = 0x17
	DW_TAG_unspecified_parameters   DW_TAG = 0x18
	DW_TAG_variant                  DW_TAG = 0x19
	DW_TAG_common_block             DW_TAG = 0x1a
	DW_TAG_common_inclusion         DW_TAG = 0x1b
	DW_TAG_inheritance              DW_TAG = 0x1c
	DW_TAG_inlined_subroutine       DW_TAG = 0x1d
	DW_TAG_module                   DW_TAG = 0x1e
	DW_TAG_ptr_to_member_type       DW_TAG = 0x1f
	DW_TAG_set_type                 DW_TAG = 0x20
	DW_TAG_subrange_type            DW_TAG = 0x21
	DW_TAG_with_stmt                DW_TAG = 0x22
	DW_TAG_access_declaration       DW_TAG = 0x23
	DW_TAG_base_type                DW_TAG = 0x24
	DW_TAG_catch_block              DW_TAG = 0x25
	DW_TAG_const_type               DW_TAG = 0x26
	DW_TAG_constant                 DW_TAG = 0x27
	DW_TAG_enumerator               DW_TAG = 0x28
	DW_TAG_file_type                DW_TAG = 0x29
	DW_TAG_friend                   DW_TAG = 0x2a
	DW_TAG_namelist                 DW_TAG = 0x2b
	DW_TAG_namelist_item            DW_TAG = 0x2c
	DW_TAG_packed_type              DW_TAG = 0x2d
	DW_TAG_subprogram               DW_TAG = 0x2e
	DW_TAG_template_type_parameter  DW_TAG = 0x2f
	DW_TAG_template_value_parameter DW_TAG = 0x30
	DW_TAG_thrown_type              DW_TAG = 0x31
	DW_TAG_try_block                DW_TAG = 0x32
	DW_TAG_variant_part             DW_TAG = 0x33
	DW_TAG_variable                 DW_TAG = 0x34
	DW_TAG_volatile_type            DW_TAG = 0x35
	DW_TAG_dwarf_procedure          DW_TAG = 0x36
	DW_TAG_restrict_type            DW_TAG = 0x37
	DW_TAG_interface_type           DW_TAG = 0x38
	DW_TAG_namespace                DW_TAG = 0x39
	DW_TAG_imported_module          DW_TAG = 0x3a
	DW_TAG_unspecified_type         DW_TAG = 0x3b
	DW_TAG_partial_unit             DW_TAG = 0x3c
	DW_TAG_imported_unit            DW_TAG = 0x3d
	DW_TAG_condition                DW_TAG = 0x3f
	DW_TAG_shared_type              DW_TAG = 0x40
	DW_TAG_type_unit                DW_TAG = 0x41
	DW_TAG_rvalue_reference_type    DW_TAG = 0x42
	DW_TAG_template_alias           DW_TAG = 0x43
	DW_TAG_lo_user                  DW_TAG = 0x4080
	DW_TAG_hi_user                  DW_TAG = 0xffff
)

const (
	DW_TAG_format_label DW_TAG = 0x4101 + iota
	DW_TAG_function_template
	DW_TAG_class_template
	DW_TAG_GNU_BINCL
	DW_TAG_GNU_EINCL
	DW_TAG_GNU_template_template_param
	DW_TAG_GNU_template_parameter_pack
	DW_TAG_GNU_formal_parameter_pack
	DW_TAG_GNU_call_site
	DW_TAG_GNU_call_site_parameter
)

const (
	DW_TAG_APPLE_property DW_TAG = 0x4200
)

const (
	DW_CHILDREN_no DW_CHILDREN = iota
	DW_CHILDREN_yes
)

const (
	DW_OP_addr        DW_OP = 0x03
	DW_OP_deref       DW_OP = 0x06
	DW_OP_const1u     DW_OP = 0x08
	DW_OP_const1s     DW_OP = 0x09
	DW_OP_const2u     DW_OP = 0x0a
	DW_OP_const2s     DW_OP = 0x0b
	DW_OP_const4u     DW_OP = 0x0c
	DW_OP_const4s     DW_OP = 0x0d
	DW_OP_const8u     DW_OP = 0x0e
	DW_OP_const8s     DW_OP = 0x0f
	DW_OP_constu      DW_OP = 0x10
	DW_OP_consts      DW_OP = 0x11
	DW_OP_dup         DW_OP = 0x12
	DW_OP_drop        DW_OP = 0x13
	DW_OP_over        DW_OP = 0x14
	DW_OP_pick        DW_OP = 0x15
	DW_OP_swap        DW_OP = 0x16
	DW_OP_rot         DW_OP = 0x17
	DW_OP_xderef      DW_OP = 0x18
	DW_OP_abs         DW_OP = 0x19
	DW_OP_and         DW_OP = 0x1a
	DW_OP_div         DW_OP = 0x1b
	DW_OP_minus       DW_OP = 0x1c
	DW_OP_mod         DW_OP = 0x1d
	DW_OP_mul         DW_OP = 0x1e
	DW_OP_neg         DW_OP = 0x1f
	DW_OP_not         DW_OP = 0x20
	DW_OP_or          DW_OP = 0x21
	DW_OP_plus        DW_OP = 0x22
	DW_OP_plus_uconst DW_OP = 0x23
	DW_OP_shl         DW_OP = 0x24
	DW_OP_shr         DW_OP = 0x25
	DW_OP_shra        DW_OP = 0x26
	DW_OP_xor         DW_OP = 0x27
	DW_OP_skip        DW_OP = 0x2f
	DW_OP_bra         DW_OP = 0x28
	DW_OP_eq          DW_OP = 0x29
	DW_OP_ge          DW_OP = 0x2a
	DW_OP_gt          DW_OP = 0x2b
	DW_OP_le          DW_OP = 0x2c
	DW_OP_lt          DW_OP = 0x2d
	DW_OP_ne          DW_OP = 0x2e
)

const (
	DW_OP_lit0 DW_OP = 0x30 + iota
	DW_OP_lit1
	DW_OP_lit2
	DW_OP_lit3
	DW_OP_lit4
	DW_OP_lit5
	DW_OP_lit6
	DW_OP_lit7
	DW_OP_lit8
	DW_OP_lit9
	DW_OP_lit10
	DW_OP_lit11
	DW_OP_lit12
	DW_OP_lit13
	DW_OP_lit14
	DW_OP_lit15
	DW_OP_lit16
	DW_OP_lit17
	DW_OP_lit18
	DW_OP_lit19
	DW_OP_lit20
	DW_OP_lit21
	DW_OP_lit22
	DW_OP_lit23
	DW_OP_lit24
	DW_OP_lit25
	DW_OP_lit26
	DW_OP_lit27
	DW_OP_lit28
	DW_OP_lit29
	DW_OP_lit30
	DW_OP_lit31
	DW_OP_reg0
	DW_OP_reg1
	DW_OP_reg2
	DW_OP_reg3
	DW_OP_reg4
	DW_OP_reg5
	DW_OP_reg6
	DW_OP_reg7
	DW_OP_reg8
	DW_OP_reg9
	DW_OP_reg10
	DW_OP_reg11
	DW_OP_reg12
	DW_OP_reg13
	DW_OP_reg14
	DW_OP_reg15
	DW_OP_reg16
	DW_OP_reg17
	DW_OP_reg18
	DW_OP_reg19
	DW_OP_reg20
	DW_OP_reg21
	DW_OP_reg22
	DW_OP_reg23
	DW_OP_reg24
	DW_OP_reg25
	DW_OP_reg26
	DW_OP_reg27
	DW_OP_reg28
	DW_OP_reg29
	DW_OP_reg30
	DW_OP_reg31
	DW_OP_breg1
	DW_OP_breg2
	DW_OP_breg3
	DW_OP_breg4
	DW_OP_breg5
	DW_OP_breg6
	DW_OP_breg7
	DW_OP_breg8
	DW_OP_breg9
	DW_OP_breg10
	DW_OP_breg11
	DW_OP_breg12
	DW_OP_breg13
	DW_OP_breg14
	DW_OP_breg15
	DW_OP_breg16
	DW_OP_breg17
	DW_OP_breg18
	DW_OP_breg19
	DW_OP_breg20
	DW_OP_breg21
	DW_OP_breg22
	DW_OP_breg23
	DW_OP_breg24
	DW_OP_breg25
	DW_OP_breg26
	DW_OP_breg27
	DW_OP_breg28
	DW_OP_breg29
	DW_OP_breg30
	DW_OP_breg31
)

const (
	DW_OP_regx                DW_OP = 0x90
	DW_OP_fbreg               DW_OP = 0x91
	DW_OP_bregx               DW_OP = 0x92
	DW_OP_piece               DW_OP = 0x93
	DW_OP_deref_size          DW_OP = 0x94
	DW_OP_xderef_size         DW_OP = 0x95
	DW_OP_nop                 DW_OP = 0x96
	DW_OP_push_object_address DW_OP = 0x97
	DW_OP_call2               DW_OP = 0x98
	DW_OP_call4               DW_OP = 0x99
	DW_OP_call_ref            DW_OP = 0x9a
	DW_OP_form_tls_address    DW_OP = 0x9b
	DW_OP_call_frame_cfa      DW_OP = 0x9c
	DW_OP_bit_piece           DW_OP = 0x9d
	DW_OP_implicit_value      DW_OP = 0x9e
	DW_OP_stack_value         DW_OP = 0x9f
	DW_OP_lo_user             DW_OP = 0xe0
	DW_OP_hi_user             DW_OP = 0xff
)

const (
	DW_FORM_addr         DW_FORM = 0x01
	DW_FORM_block2       DW_FORM = 0x03
	DW_FORM_block4       DW_FORM = 0x04
	DW_FORM_data2        DW_FORM = 0x05
	DW_FORM_data4        DW_FORM = 0x06
	DW_FORM_data8        DW_FORM = 0x07
	DW_FORM_string       DW_FORM = 0x08
	DW_FORM_block        DW_FORM = 0x09
	DW_FORM_block1       DW_FORM = 0x0a
	DW_FORM_data1        DW_FORM = 0x0b
	DW_FORM_flag         DW_FORM = 0x0c
	DW_FORM_sdata        DW_FORM = 0x0d
	DW_FORM_strp         DW_FORM = 0x0e
	DW_FORM_udata        DW_FORM = 0x0f
	DW_FORM_ref_addr     DW_FORM = 0x10
	DW_FORM_ref1         DW_FORM = 0x11
	DW_FORM_ref2         DW_FORM = 0x12
	DW_FORM_ref4         DW_FORM = 0x13
	DW_FORM_ref8         DW_FORM = 0x14
	DW_FORM_ref_udata    DW_FORM = 0x15
	DW_FORM_indirect     DW_FORM = 0x16
	DW_FORM_sec_offset   DW_FORM = 0x17
	DW_FORM_exprloc      DW_FORM = 0x18
	DW_FORM_flag_present DW_FORM = 0x19
	DW_FORM_ref_sig8     DW_FORM = 0x20
)

const (
	DW_AT_sibling              DW_AT = 0x01
	DW_AT_location             DW_AT = 0x02
	DW_AT_name                 DW_AT = 0x03
	DW_AT_ordering             DW_AT = 0x09
	DW_AT_byte_size            DW_AT = 0x0b
	DW_AT_bit_offset           DW_AT = 0x0c
	DW_AT_bit_size             DW_AT = 0x0d
	DW_AT_stmt_list            DW_AT = 0x10
	DW_AT_low_pc               DW_AT = 0x11
	DW_AT_high_pc              DW_AT = 0x12
	DW_AT_language             DW_AT = 0x13
	DW_AT_discr                DW_AT = 0x15
	DW_AT_discr_value          DW_AT = 0x16
	DW_AT_visibility           DW_AT = 0x17
	DW_AT_import               DW_AT = 0x18
	DW_AT_string_length        DW_AT = 0x19
	DW_AT_common_reference     DW_AT = 0x1a
	DW_AT_comp_dir             DW_AT = 0x1b
	DW_AT_const_value          DW_AT = 0x1c
	DW_AT_containing_type      DW_AT = 0x1d
	DW_AT_default_value        DW_AT = 0x1e
	DW_AT_inline               DW_AT = 0x20
	DW_AT_is_optional          DW_AT = 0x21
	DW_AT_lower_bound          DW_AT = 0x22
	DW_AT_producer             DW_AT = 0x25
	DW_AT_prototyped           DW_AT = 0x27
	DW_AT_return_addr          DW_AT = 0x2a
	DW_AT_start_scope          DW_AT = 0x2c
	DW_AT_bit_stride           DW_AT = 0x2e
	DW_AT_upper_bound          DW_AT = 0x2f
	DW_AT_abstract_origin      DW_AT = 0x31
	DW_AT_accessibility        DW_AT = 0x32
	DW_AT_address_class        DW_AT = 0x33
	DW_AT_artificial           DW_AT = 0x34
	DW_AT_base_types           DW_AT = 0x35
	DW_AT_calling_convention   DW_AT = 0x36
	DW_AT_count                DW_AT = 0x37
	DW_AT_data_member_location DW_AT = 0x38
	DW_AT_decl_column          DW_AT = 0x39
	DW_AT_decl_file            DW_AT = 0x3a
	DW_AT_decl_line            DW_AT = 0x3b
	DW_AT_declaration          DW_AT = 0x3c
	DW_AT_discr_list           DW_AT = 0x3d
	DW_AT_encoding             DW_AT = 0x3e
	DW_AT_external             DW_AT = 0x3f
	DW_AT_frame_base           DW_AT = 0x40
	DW_AT_friend               DW_AT = 0x41
	DW_AT_identifier_case      DW_AT = 0x42
	DW_AT_macro_info           DW_AT = 0x43
	DW_AT_namelist_item        DW_AT = 0x44
	DW_AT_priority             DW_AT = 0x45
	DW_AT_segment              DW_AT = 0x46
	DW_AT_specification        DW_AT = 0x47
	DW_AT_static_link          DW_AT = 0x48
	DW_AT_type                 DW_AT = 0x49
	DW_AT_use_location         DW_AT = 0x4a
	DW_AT_variable_parameter   DW_AT = 0x4b
	DW_AT_virtuality           DW_AT = 0x4c
	DW_AT_vtable_elem_location DW_AT = 0x4d
	DW_AT_allocated            DW_AT = 0x4e
	DW_AT_associated           DW_AT = 0x4f
	DW_AT_data_location        DW_AT = 0x50
	DW_AT_byte_stride          DW_AT = 0x51
	DW_AT_entry_pc             DW_AT = 0x52
	DW_AT_use_UTF8             DW_AT = 0x53
	DW_AT_extension            DW_AT = 0x54
	DW_AT_ranges               DW_AT = 0x55
	DW_AT_trampoline           DW_AT = 0x56
	DW_AT_call_column          DW_AT = 0x57
	DW_AT_call_file            DW_AT = 0x58
	DW_AT_call_line            DW_AT = 0x59
	DW_AT_description          DW_AT = 0x5a
	DW_AT_binary_scale         DW_AT = 0x5b
	DW_AT_decimal_scale        DW_AT = 0x5c
	DW_AT_small                DW_AT = 0x5d
	DW_AT_decimal_sign         DW_AT = 0x5e
	DW_AT_digit_count          DW_AT = 0x5f
	DW_AT_picture_string       DW_AT = 0x60
	DW_AT_mutable              DW_AT = 0x61
	DW_AT_threads_scaled       DW_AT = 0x62
	DW_AT_explicit             DW_AT = 0x63
	DW_AT_object_pointer       DW_AT = 0x64
	DW_AT_endianity            DW_AT = 0x65
	DW_AT_elemental            DW_AT = 0x66
	DW_AT_pure                 DW_AT = 0x67
	DW_AT_recursive            DW_AT = 0x68
	DW_AT_signature            DW_AT = 0x69
	DW_AT_main_subprogram      DW_AT = 0x6a
	DW_AT_data_bit_offset      DW_AT = 0x6b
	DW_AT_const_expr           DW_AT = 0x6c
	DW_AT_enum_class           DW_AT = 0x6d
	DW_AT_linkage_name         DW_AT = 0x6e
	DW_AT_lo_user              DW_AT = 0x2000
	DW_AT_hi_user              DW_AT = 0x3fff
)

const (
	DW_AT_MIPS_loop_begin DW_AT = 0x2002 + iota
	DW_AT_MIPS_tail_loop_begin
	DW_AT_MIPS_epilog_begin
	DW_AT_MIPS_loop_unroll_factor
	DW_AT_MIPS_software_pipeline_depth
	DW_AT_MIPS_linkage_name
	DW_AT_MIPS_stride
	DW_AT_MIPS_abstract_name
	DW_AT_MIPS_clone_origin
	DW_AT_MIPS_has_inlines
	DW_AT_MIPS_stride_byte
	DW_AT_MIPS_stride_elem
	DW_AT_MIPS_ptr_dopetype
	DW_AT_MIPS_allocatable_dopetype
	DW_AT_MIPS_assumed_shape_dopetype
	DW_AT_MIPS_assumed_size
)
const (
	DW_AT_sf_names DW_AT = 0x2101 + iota
	DW_AT_src_info
	DW_AT_mac_info
	DW_AT_src_coords
	DW_AT_body_begin
	DW_AT_body_end
	DW_AT_GNU_vector
	DW_AT_GNU_guarded_by
	DW_AT_GNU_pt_guarded_by
	DW_AT_GNU_guarded
	DW_AT_GNU_pt_guarded
	DW_AT_GNU_locks_excluded
	DW_AT_GNU_exclusive_locks_required
	DW_AT_GNU_shared_locks_required
	DW_AT_GNU_odr_signature
	DW_AT_GNU_template_name
	DW_AT_GNU_call_site_value
	DW_AT_GNU_call_site_data_value
	DW_AT_GNU_call_site_target
	DW_AT_GNU_call_site_target_clobbered
	DW_AT_GNU_tail_call
	DW_AT_GNU_all_tail_call_sites
	DW_AT_GNU_all_call_sites
	DW_AT_GNU_all_source_call_sites
	DW_AT_GNU_macros
)

const (
	DW_AT_APPLE_optimized DW_AT = 0x3fe1 + iota
	DW_AT_APPLE_flags
	DW_AT_APPLE_isa
	DW_AT_APPLE_block
	DW_AT_APPLE_major_runtime_vers
	DW_AT_APPLE_runtime_class
	DW_AT_APPLE_omit_frame_ptr
	DW_AT_APPLE_property_name
	DW_AT_APPLE_property_getter
	DW_AT_APPLE_property_setter
	DW_AT_APPLE_property_attribute
	DW_AT_APPLE_objc_complete_type
	DW_AT_APPLE_property
)
