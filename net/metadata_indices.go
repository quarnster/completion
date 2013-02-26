package net

type (
	TableIndex        uint32
	AssemblyRefIndex  TableIndex
	BlobIndex         TableIndex
	EventIndex        TableIndex
	FieldIndex        TableIndex
	GenericParamIndex TableIndex
	GuidIndex         TableIndex
	MethodDefIndex    TableIndex
	ModuleRefIndex    TableIndex
	ParamIndex        TableIndex
	PropertyIndex     TableIndex
	StringIndex       TableIndex
	TypeDefIndex      TableIndex
)

// II.24.2.6
type (
	EncodedIndex                    TableIndex
	CustomAttributeTypeEncodedIndex EncodedIndex
	HasConstantEncodedIndex         EncodedIndex
	HasCustomAttributeEncodedIndex  EncodedIndex
	HasDeclSecurityEncodedIndex     EncodedIndex
	HasFieldMarshalEncodedIndex     EncodedIndex
	HasSemanticsEncodedIndex        EncodedIndex
	ImplementationEncodedIndex      EncodedIndex
	MemberForwardedEncodedIndex     EncodedIndex
	MemberRefParentEncodedIndex     EncodedIndex
	MethodDefOrRefEncodedIndex      EncodedIndex
	ResolutionScopeEncodedIndex     EncodedIndex
	TypeDefOrRefEncodedIndex        EncodedIndex
	TypeOrMethodDefEncodedIndex     EncodedIndex
)

func (t TypeDefOrRefEncodedIndex) Tables() []int {
	return []int{
		id_TypeDef,
		id_TypeRef,
		id_TypeSpec}
}
func (t HasConstantEncodedIndex) Tables() []int {
	return []int{
		id_Field,
		id_Param,
		id_Property}
}
func (t HasCustomAttributeEncodedIndex) Tables() []int {
	return []int{
		id_MethodDef,
		id_Field,
		id_TypeRef,
		id_TypeDef,
		id_Param,
		id_InterfaceImpl,
		id_MemberRef,
		id_Module,
		id_DeclSecurity,
		id_Property,
		id_Event,
		id_StandAloneSig,
		id_ModuleRef,
		id_TypeSpec,
		id_Assembly,
		id_AssemblyRef,
		id_File,
		id_ExportedType,
		id_ManifestResource,
		id_GenericParam,
		id_GenericParamConstraint,
		id_MethodSpec}
}
func (t HasFieldMarshalEncodedIndex) Tables() []int {
	return []int{
		id_Field,
		id_Param}
}
func (t HasDeclSecurityEncodedIndex) Tables() []int {
	return []int{
		id_TypeDef,
		id_MethodDef,
		id_Assembly}
}
func (t MemberRefParentEncodedIndex) Tables() []int {
	return []int{
		id_TypeDef,
		id_TypeRef,
		id_ModuleRef,
		id_MethodDef,
		id_TypeSpec}
}
func (t HasSemanticsEncodedIndex) Tables() []int {
	return []int{
		id_Event,
		id_Property}
}
func (t MethodDefOrRefEncodedIndex) Tables() []int {
	return []int{
		id_MethodDef,
		id_MemberRef}
}
func (t MemberForwardedEncodedIndex) Tables() []int {
	return []int{
		id_Field,
		id_MethodDef}
}
func (t ImplementationEncodedIndex) Tables() []int {
	return []int{
		id_File,
		id_AssemblyRef,
		id_ExportedType}
}
func (t CustomAttributeTypeEncodedIndex) Tables() []int {
	return []int{
		id_nullTable,
		id_nullTable,
		id_MethodDef,
		id_MemberRef,
		id_nullTable}
}
func (t ResolutionScopeEncodedIndex) Tables() []int {
	return []int{
		id_Module,
		id_ModuleRef,
		id_AssemblyRef,
		id_TypeRef}
}
func (t TypeOrMethodDefEncodedIndex) Tables() []int {
	return []int{
		id_TypeDef,
		id_MethodDef,
	}
}

// Todo:
// func bits(values int) uint {
// 	return uint(math.Ceil(math.Log2(float64(values))))
// }

// func (id EncodedIndex) Size(h *hash_tilde_stream_header) uintptr {
// 	var (
// 		tables = id.Tables()
// 		rows   uintptr
// 	)
// 	for _, t := range tables {
// 		if s2 := h.TableSize(t); s2 > rows {
// 			rows = s2
// 		}
// 	}
// 	if rows<<bits(len(tables)) < 1<<16 {
// 		return 2
// 	}
// 	return 4
// }
