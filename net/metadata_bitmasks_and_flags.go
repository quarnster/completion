package net

type (
	AssemblyFlags              uint32 //  II.23.1.2
	AssemblyHashAlgorithm      uint32 //  II.23.1.1
	EventAttributes            uint16 //  II.23.1.4
	FieldAttributes            uint16 //  II.23.1.5
	FileAttributes             uint32 //  II.23.1.6
	GenericParamAttributes     uint16 //  II.23.1.7
	ManifestResourceAttributes uint32 //  II.23.1.9
	MethodAttributes           uint16 //  II.23.1.10
	MethodImplAttributes       uint16 //  II.23.1.10
	MethodSemanticsAttributes  uint16 //  II.23.1.12
	ParamAttributes            uint16 //  II.23.1.13
	PInvokeAttributes          uint16 //  II.23.1.8
	PropertyAttributes         uint16 //  II.23.1.14
	TypeAttributes             uint32 //  II.23.1.15
)

// II.23.1.1 AssemblyHashAlgorithm
const (
	AssemblyHashAlgorithm_None         = 0x0000
	AssemblyHashAlgorithm_Reserved_MD5 = 0x8003
	AssemblyHashAlgorithm_SHA1         = 0x8004
)

func (h AssemblyHashAlgorithm) String() string {
	switch int(h) {
	case AssemblyHashAlgorithm_None:
		return "None"
	case AssemblyHashAlgorithm_SHA1:
		return "SHA1"
	case AssemblyHashAlgorithm_Reserved_MD5:
		return "Reserved (MD5)"
	}
	return "Unknown"
}

// II.23.1.2 AssemblyFlags
const (
	AssemblyFlags_PublicKey                  = 0x0001
	AssemblyFlags_Retargetable               = 0x0100
	AssemblyFlags_DisableJITcompileOptimizer = 0x4000
	AssemblyFlags_EnableJITcompileTracking   = 0x8000
)

// II.23.1.4 EventAttributes
const (
	EventAttributes_SpecialName   = 0x0200
	EventAttributes_RTSpecialName = 0x0400
)

// II.23.1.5 FieldAttributes
const (
	FieldAttributes_FieldAccessMask    = 0x0007
	FieldAttributes_CompilerControlled = 0x0000
	FieldAttributes_Private            = 0x0001
	FieldAttributes_FamANDAssem        = 0x0002
	FieldAttributes_Assembly           = 0x0003
	FieldAttributes_Family             = 0x0004
	FieldAttributes_FamORAssem         = 0x0005
	FieldAttributes_Public             = 0x0006
	FieldAttributes_Static             = 0x0010
	FieldAttributes_InitOnly           = 0x0020
	FieldAttributes_Literal            = 0x0040
	FieldAttributes_NotSerialized      = 0x0080
	FieldAttributes_SpecialName        = 0x0200
	FieldAttributes_PInvokeImpl        = 0x2000
	FieldAttributes_RTSpecialName      = 0x0400
	FieldAttributes_HasFieldMarshal    = 0x1000
	FieldAttributes_HasDefault         = 0x8000
	FieldAttributes_HasFieldRVA        = 0x0100
)

// II.23.1.6 FileAttributes
const (
	FileAttributes_ContainsMetaData   = 0x0000
	FileAttributes_ContainsNoMetaData = 0x0001
)

// II.23.1.7 GenericParamAttributes
const (
	GenericParamAttributes_VarianceMask                   = 0x0003
	GenericParamAttributes_None                           = 0x0000
	GenericParamAttributes_Covariant                      = 0x0001
	GenericParamAttributes_Contravariant                  = 0x0002
	GenericParamAttributes_SpecialConstraintMask          = 0x001C
	GenericParamAttributes_ReferenceTypeConstraint        = 0x0004
	GenericParamAttributes_NotNullableValueTypeConstraint = 0x0008
	GenericParamAttributes_DefaultConstructorConstraint   = 0x0010
)

// II.23.1.8 PInvokeAttributes
const (
	PInvokeAttributes_NoMangle            = 0x0001
	PInvokeAttributes_CharSetMask         = 0x0006
	PInvokeAttributes_CharSetNotSpec      = 0x0000
	PInvokeAttributes_CharSetAnsi         = 0x0002
	PInvokeAttributes_CharSetUnicode      = 0x0004
	PInvokeAttributes_CharSetAuto         = 0x0006
	PInvokeAttributes_SupportsLastError   = 0x0040
	PInvokeAttributes_CallConvMask        = 0x0700
	PInvokeAttributes_CallConvPlatformapi = 0x0100
	PInvokeAttributes_CallConvCdecl       = 0x0200
	PInvokeAttributes_CallConvStdcall     = 0x0300
	PInvokeAttributes_CallConvThiscall    = 0x0400
	PInvokeAttributes_CallConvFastcall    = 0x0500
)

// II.23.1.9 ManifestResourceAttributes
const (
	ManifestResourceAttributes_VisibilityMask = 0x0007
	ManifestResourceAttributes_Public         = 0x0001
	ManifestResourceAttributes_Private        = 0x0002
)

// II.23.1.10 MethodAttributes
const (
	MethodAttributes_MemberAccessMask   = 0x0007
	MethodAttributes_CompilerControlled = 0x0000
	MethodAttributes_Private            = 0x0001
	MethodAttributes_FamANDAssem        = 0x0002
	MethodAttributes_Assem              = 0x0003
	MethodAttributes_Family             = 0x0004
	MethodAttributes_FamORAssem         = 0x0005
	MethodAttributes_Public             = 0x0006
	MethodAttributes_Static             = 0x0010
	MethodAttributes_Final              = 0x0020
	MethodAttributes_Virtual            = 0x0040
	MethodAttributes_HideBySig          = 0x0080
	MethodAttributes_VtableLayoutMask   = 0x0100
	MethodAttributes_ReuseSlot          = 0x0000
	MethodAttributes_NewSlot            = 0x0100
	MethodAttributes_Strict             = 0x0200
	MethodAttributes_Abstract           = 0x0400
	MethodAttributes_SpecialName        = 0x0800
	MethodAttributes_PInvokeImpl        = 0x2000
	MethodAttributes_UnmanagedExport    = 0x0008
	MethodAttributes_RTSpecialName      = 0x1000
	MethodAttributes_HasSecurity        = 0x4000
	MethodAttributes_RequireSecObject   = 0x8000
)

// II.23.1.11 MethodImplAttributes
const (
	MethodImplAttributes_CodeTypeMask     = 0x0003
	MethodImplAttributes_IL               = 0x0000
	MethodImplAttributes_Native           = 0x0001
	MethodImplAttributes_OPTIL            = 0x0002
	MethodImplAttributes_Runtime          = 0x0003
	MethodImplAttributes_ManagedMask      = 0x0004
	MethodImplAttributes_Unmanaged        = 0x0004
	MethodImplAttributes_Managed          = 0x0000
	MethodImplAttributes_ForwardRef       = 0x0010
	MethodImplAttributes_PreserveSig      = 0x0080
	MethodImplAttributes_InternalCall     = 0x1000
	MethodImplAttributes_Synchronized     = 0x0020
	MethodImplAttributes_NoInlining       = 0x0008
	MethodImplAttributes_MaxMethodImplVal = 0xffff
	MethodImplAttributes_NoOptimization   = 0x0040
)

// II.23.1.12 MethodSemanticsAttributes
const (
	MethodSemanticsAttributes_Setter   = 0x0001
	MethodSemanticsAttributes_Getter   = 0x0002
	MethodSemanticsAttributes_Other    = 0x0004
	MethodSemanticsAttributes_AddOn    = 0x0008
	MethodSemanticsAttributes_RemoveOn = 0x0010
	MethodSemanticsAttributes_Fire     = 0x0020
)

// II.23.1.13 ParamAttributes
const (
	ParamAttributes_In              = 0x0001
	ParamAttributes_Out             = 0x0002
	ParamAttributes_Optional        = 0x0010
	ParamAttributes_HasDefault      = 0x1000
	ParamAttributes_HasFieldMarshal = 0x2000
	ParamAttributes_Unused          = 0xcfe0
)

// II.23.1.14 PropertyAttributes
const (
	PropertyAttributes_SpecialName   = 0x0200
	PropertyAttributes_RTSpecialName = 0x0400
	PropertyAttributes_HasDefault    = 0x1000
	PropertyAttributes_Unused        = 0xe9ff
)

// II.23.1.15 TypeAttributes
const (
	TypeAttributes_VisibilityMask         = 0x00000007
	TypeAttributes_NotPublic              = 0x00000000
	TypeAttributes_Public                 = 0x00000001
	TypeAttributes_NestedPublic           = 0x00000002
	TypeAttributes_NestedPrivate          = 0x00000003
	TypeAttributes_NestedFamily           = 0x00000004
	TypeAttributes_NestedAssembly         = 0x00000005
	TypeAttributes_NestedFamANDAssem      = 0x00000006
	TypeAttributes_NestedFamORAssem       = 0x00000007
	TypeAttributes_LayoutMask             = 0x00000018
	TypeAttributes_AutoLayout             = 0x00000000
	TypeAttributes_SequentialLayout       = 0x00000008
	TypeAttributes_ExplicitLayout         = 0x00000010
	TypeAttributes_ClassSemanticsMask     = 0x00000020
	TypeAttributes_Class                  = 0x00000000
	TypeAttributes_Interface              = 0x00000020
	TypeAttributes_Abstract               = 0x00000080
	TypeAttributes_Sealed                 = 0x00000100
	TypeAttributes_SpecialName            = 0x00000400
	TypeAttributes_Import                 = 0x00001000
	TypeAttributes_Serializable           = 0x00002000
	TypeAttributes_StringFormatMask       = 0x00030000
	TypeAttributes_AnsiClass              = 0x00000000
	TypeAttributes_UnicodeClass           = 0x00010000
	TypeAttributes_AutoClass              = 0x00020000
	TypeAttributes_CustomFormatClass      = 0x00030000
	TypeAttributes_CustomStringFormatMask = 0x00C00000
	TypeAttributes_BeforeFieldInit        = 0x00100000
	TypeAttributes_RTSpecialName          = 0x00000800
	TypeAttributes_HasSecurity            = 0x00040000
	TypeAttributes_IsTypeForwarder        = 0x00200000
)

// II.23.1.16 Element types used in signatures
const (
	ELEMENT_TYPE_END         = 0x00
	ELEMENT_TYPE_VOID        = 0x01
	ELEMENT_TYPE_BOOLEAN     = 0x02
	ELEMENT_TYPE_CHAR        = 0x03
	ELEMENT_TYPE_I1          = 0x04
	ELEMENT_TYPE_U1          = 0x05
	ELEMENT_TYPE_I2          = 0x06
	ELEMENT_TYPE_U2          = 0x07
	ELEMENT_TYPE_I4          = 0x08
	ELEMENT_TYPE_U4          = 0x09
	ELEMENT_TYPE_I8          = 0x0a
	ELEMENT_TYPE_U8          = 0x0b
	ELEMENT_TYPE_R4          = 0x0c
	ELEMENT_TYPE_R8          = 0x0d
	ELEMENT_TYPE_STRING      = 0x0e
	ELEMENT_TYPE_PTR         = 0x0f
	ELEMENT_TYPE_BYREF       = 0x10
	ELEMENT_TYPE_VALUETYPE   = 0x11
	ELEMENT_TYPE_CLASS       = 0x12
	ELEMENT_TYPE_VAR         = 0x13
	ELEMENT_TYPE_ARRAY       = 0x14
	ELEMENT_TYPE_GENERICINST = 0x15
	ELEMENT_TYPE_TYPEDBYREF  = 0x16
	ELEMENT_TYPE_I           = 0x18
	ELEMENT_TYPE_U           = 0x19
	ELEMENT_TYPE_FNPTR       = 0x1b
	ELEMENT_TYPE_OBJECT      = 0x1c
	ELEMENT_TYPE_SZARRAY     = 0x1d
	ELEMENT_TYPE_MVAR        = 0x1e
	ELEMENT_TYPE_CMOD_REQD   = 0x1f
	ELEMENT_TYPE_CMOD_OPT    = 0x20
	ELEMENT_TYPE_INTERNAL    = 0x21
	ELEMENT_TYPE_MODIFIER    = 0x40
	ELEMENT_TYPE_SENTINEL    = 0x41
	ELEMENT_TYPE_PINNED      = 0x45

	ELEMENT_TYPE_SYSTEM           = 0x50 // Indicates an argument of type System.Type
	CUSTOM_ATTRIBUTE_BOXED_OBJECT = 0x51 // Used in custom attributes to specify a boxed object (§II.23.3).
	CUSTOM_ATTRIBUTE_FIELD        = 0x53 // Used in custom attributes to indicate a FIELD (§II.22.10, II.23.3).
	CUSTOM_ATTRIBUTE_PROPERTY     = 0x54 // Used in custom attributes to indicate a PROPERTY (§II.22.10, II.23.3).
	CUSTOM_ATTRIBUTE_ENUM         = 0x55 // Used in custom attributes to specify an enum (§II.23.3).
)
