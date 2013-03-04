package net

type (

	// Spec II.22.2
	AssemblyRow struct {
		HashAlgId      AssemblyHashAlgorithm
		MajorVersion   uint16
		MinorVersion   uint16
		BuildNumber    uint16
		RevisionNumber uint16
		Flags          AssemblyFlags
		PublicKey      BlobIndex
		Name           StringIndex
		Culture        StringIndex
	}

	// Spec II.22.3
	AssemblyOSRow struct {
		OSPlatformID   uint32
		OSMajorVersion uint32
		OSMinorVersion uint32
	}

	// Spec II.22.4
	AssemblyProcessorRow struct {
		Processor uint32
	}

	// Spec II.22.5
	AssemblyRefRow struct {
		MajorVersion     uint16
		MinorVersion     uint16
		BuildNumber      uint16
		RevisionNumber   uint16
		Flags            AssemblyFlags
		PublicKeyOrToken BlobIndex
		Name             StringIndex
		Culture          StringIndex
		HashValue        BlobIndex
	}

	// Spec II.22.6
	AssemblyRefOSRow struct {
		OSPlatformId   uint32
		OSMajorVersion uint32
		OSMinorVersion uint32
		AssemblyRef    AssemblyRefIndex
	}

	// Spec II.22.7
	AssemblyRefProcessorRow struct {
		Processor   uint32
		AssemblyRef AssemblyRefIndex
	}

	// Spec II.22.8
	ClassLayoutRow struct {
		PackingSize uint16
		ClassSize   uint32
		Parent      TypeDefIndex
	}

	// Spec II.22.9
	ConstantRow struct {
		Type    uint8
		Padding uint8
		Parent  HasConstantEncodedIndex
		Value   BlobIndex
	}

	// Spec II.22.10
	CustomAttributeRow struct {
		Parent HasCustomAttributeEncodedIndex
		Type   CustomAttributeTypeEncodedIndex
		Value  BlobIndex
	}

	// Spec II.22.11
	DeclSecurityRow struct {
		Action        uint16
		Parent        HasDeclSecurityEncodedIndex
		PermissionSet BlobIndex
	}

	// Spec II.22.12
	EventMapRow struct {
		Parent    TypeDefIndex
		EventList EventIndex
	}

	// Spec II.22.13
	EventRow struct {
		EventFlags EventAttributes
		Name       StringIndex
		EventType  TypeDefOrRefEncodedIndex
	}

	// Spec II.22.14
	ExportedTypeRow struct {
		Flags          TypeAttributes
		TypeDefId      uint32
		TypeName       StringIndex
		TypeNamespace  StringIndex
		Implementation ImplementationEncodedIndex
	}

	// Spec II.22.15
	FieldRow struct {
		Flags     FieldAttributes
		Name      StringIndex
		Signature BlobIndex
	}

	// Spec II.22.16
	FieldLayoutRow struct {
		Offset uint32
		Field  FieldIndex
	}

	// Spec II.22.17
	FieldMarshalRow struct {
		Parent     HasFieldMarshalEncodedIndex
		NativeType BlobIndex
	}

	// Spec II.22.18
	FieldRVARow struct {
		RVA   uint32
		Field FieldIndex
	}

	// Spec II.22.19
	FileRow struct {
		Flags     FileAttributes
		Name      StringIndex
		HashValue BlobIndex
	}

	// Spec II.22.20
	GenericParamRow struct {
		Number uint16
		Flags  GenericParamAttributes
		Owner  TypeOrMethodDefEncodedIndex
		Name   StringIndex
	}

	// Spec II.22.21
	GenericParamConstraintRow struct {
		Owner      GenericParamIndex
		Constraint TypeDefOrRefEncodedIndex
	}

	// Spec II.22.22
	ImplMapRow struct {
		MappingFlags    PInvokeAttributes
		MemberForwarded MemberForwardedEncodedIndex
		ImportName      StringIndex
		ImportScope     ModuleRefIndex
	}

	// Spec II.22.23
	InterfaceImplRow struct {
		Class     TypeDefIndex
		Interface TypeDefOrRefEncodedIndex
	}

	// Spec II.22.24
	ManifestResourceRow struct {
		Offset         uint32
		Flags          ManifestResourceAttributes
		Name           StringIndex
		Implementation ImplementationEncodedIndex
	}

	// Spec II.22.25
	MemberRefRow struct {
		Class     MemberRefParentEncodedIndex
		Name      StringIndex
		Signature BlobIndex
	}

	// Spec II.22.26
	MethodDefRow struct {
		RVA       uint32
		ImplFlags MethodImplAttributes
		Flags     MethodAttributes
		Name      StringIndex
		Signature BlobIndex
		ParamList ParamIndex
	}

	// Spec II.22.27
	MethodImplRow struct {
		Class             TypeDefIndex
		MethodBody        MethodDefOrRefEncodedIndex
		MethodDeclaration MethodDefOrRefEncodedIndex
	}

	// Spec II.22.28
	MethodSemanticsRow struct {
		Semantics   MethodSemanticsAttributes
		Method      MethodDefIndex
		Association HasSemanticsEncodedIndex
	}

	// Spec II.22.29
	MethodSpecRow struct {
		Method        MethodDefOrRefEncodedIndex
		Instantiation BlobIndex
	}

	// Spec II.22.30
	ModuleRow struct {
		Generation uint16
		Name       StringIndex
		Mvid       Guid
		EncId      Guid
		EncBaseId  Guid
	}

	// Spec II.22.31
	ModuleRefRow struct {
		Name StringIndex
	}

	// Spec II.22.32
	NestedClassRow struct {
		NestedClass    TypeDefIndex
		EnclosingClass TypeDefIndex
	}

	// Spec II.22.33
	ParamRow struct {
		Flags    ParamAttributes
		Sequence uint16
		Name     StringIndex
	}

	// Spec II.22.34
	PropertyRow struct {
		Flags PropertyAttributes
		Name  StringIndex
		Type  BlobIndex
	}

	// Spec II.22.35
	PropertyMapRow struct {
		Parent       TypeDefIndex
		PropertyList PropertyIndex
	}

	// Spec II.22.36
	StandAloneSigRow struct {
		Signature BlobIndex
	}

	// Spec II.22.37
	TypeDefRow struct {
		Flags         TypeAttributes
		TypeName      StringIndex
		TypeNamespace StringIndex
		Extends       TypeDefOrRefEncodedIndex
		FieldList     FieldIndex
		MethodList    MethodDefIndex
	}

	// Spec II.22.38
	TypeRefRow struct {
		ResolutionScope ResolutionScopeEncodedIndex
		TypeName        StringIndex
		TypeNamespace   StringIndex
	}

	// Spec II.22.39
	TypeSpecRow struct {
		Signature BlobIndex
	}
)
