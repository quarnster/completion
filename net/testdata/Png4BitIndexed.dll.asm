ModuleRow
	&{Generation:0 Name:Png4BitIndexed.dll Mvid:DC74EA34-D62F-4136-93BF-797DF4217B8E EncId:nil EncBaseId:nil}
TypeRefRow
	&{ResolutionScope:AssemblyRefRow[1] TypeName:Object TypeNamespace:System}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:ArgumentOutOfRangeException TypeNamespace:System}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:Byte TypeNamespace:System}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:RuntimeHelpers TypeNamespace:System.Runtime.CompilerServices}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:Array TypeNamespace:System}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:RuntimeFieldHandle TypeNamespace:System}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:UInt32 TypeNamespace:System}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:ArgumentNullException TypeNamespace:System}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:ArgumentException TypeNamespace:System}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:FileStream TypeNamespace:System.IO}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:FileMode TypeNamespace:System.IO}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:Stream TypeNamespace:System.IO}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:MemoryStream TypeNamespace:System.IO}
	&{ResolutionScope:AssemblyRefRow[2] TypeName:DeflateStream TypeNamespace:System.IO.Compression}
	&{ResolutionScope:AssemblyRefRow[2] TypeName:CompressionMode TypeNamespace:System.IO.Compression}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:IDisposable TypeNamespace:System}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:CompilerGeneratedAttribute TypeNamespace:System.Runtime.CompilerServices}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:ValueType TypeNamespace:System}
	&{ResolutionScope:AssemblyRefRow[1] TypeName:RuntimeCompatibilityAttribute TypeNamespace:System.Runtime.CompilerServices}
TypeDefRow
	&{Flags:0 TypeName:<Module> TypeNamespace: Extends:TypeDefRow[0] FieldList:FieldRow[1] MethodList:MethodDefRow[1]}
	&{Flags:1048833 TypeName:Png4BitIndexed TypeNamespace:PeterO Extends:TypeRefRow[1] FieldList:FieldRow[1] MethodList:MethodDefRow[1]}
	&{Flags:1048960 TypeName:<PrivateImplementationDetails>{dc74ea34-d62f-4136-93bf-797df4217b8e} TypeNamespace: Extends:TypeRefRow[1] FieldList:FieldRow[14] MethodList:MethodDefRow[14]}
	&{Flags:1048851 TypeName:$ArrayType=32 TypeNamespace: Extends:TypeRefRow[18] FieldList:FieldRow[22] MethodList:MethodDefRow[14]}
	&{Flags:1048851 TypeName:$ArrayType=4 TypeNamespace: Extends:TypeRefRow[18] FieldList:FieldRow[22] MethodList:MethodDefRow[14]}
	&{Flags:1048851 TypeName:$ArrayType=12 TypeNamespace: Extends:TypeRefRow[18] FieldList:FieldRow[22] MethodList:MethodDefRow[14]}
FieldRow
	&{Flags:1 Name:subdata1 Signature:Blob[1]}
	&{Flags:1 Name:crcTable Signature:Blob[5]}
	&{Flags:1 Name:width Signature:Blob[9]}
	&{Flags:1 Name:height Signature:Blob[9]}
	&{Flags:1 Name:realRowSize Signature:Blob[9]}
	&{Flags:1 Name:blockSize Signature:Blob[9]}
	&{Flags:1 Name:rowSize Signature:Blob[9]}
	&{Flags:1 Name:imageData Signature:Blob[1]}
	&{Flags:1 Name:data Signature:Blob[1]}
	&{Flags:1 Name:subdata2 Signature:Blob[1]}
	&{Flags:1 Name:colors Signature:Blob[1]}
	&{Flags:1 Name:idatCrc Signature:Blob[12]}
	&{Flags:1 Name:transparent Signature:Blob[9]}
	&{Flags:275 Name:$field-0 Signature:Blob[24]}
	&{Flags:275 Name:$field-1 Signature:Blob[36]}
	&{Flags:275 Name:$field-2 Signature:Blob[40]}
	&{Flags:275 Name:$field-3 Signature:Blob[36]}
	&{Flags:275 Name:$field-4 Signature:Blob[36]}
	&{Flags:275 Name:$field-5 Signature:Blob[36]}
	&{Flags:275 Name:$field-6 Signature:Blob[36]}
	&{Flags:275 Name:$field-7 Signature:Blob[36]}
MethodDefRow
	&{RVA:8272 ImplFlags:0 Flags:6278 Name:.ctor Signature:Blob[83] ParamList:ParamRow[1]}
	&{RVA:8584 ImplFlags:0 Flags:129 Name:Adler32 Signature:Blob[89] ParamList:ParamRow[3]}
	&{RVA:8788 ImplFlags:0 Flags:129 Name:Crc32 Signature:Blob[109] ParamList:ParamRow[6]}
	&{RVA:8988 ImplFlags:0 Flags:134 Name:SetPixel Signature:Blob[126] ParamList:ParamRow[10]}
	&{RVA:9180 ImplFlags:0 Flags:134 Name:GetPixel Signature:Blob[137] ParamList:ParamRow[13]}
	&{RVA:9304 ImplFlags:0 Flags:134 Name:GetColor Signature:Blob[143] ParamList:ParamRow[15]}
	&{RVA:9400 ImplFlags:0 Flags:134 Name:SetColor Signature:Blob[149] ParamList:ParamRow[16]}
	&{RVA:9519 ImplFlags:0 Flags:129 Name:GetBE Signature:Blob[156] ParamList:ParamRow[18]}
	&{RVA:9580 ImplFlags:0 Flags:134 Name:Save Signature:Blob[19] ParamList:ParamRow[19]}
	&{RVA:10336 ImplFlags:0 Flags:2182 Name:get_Height Signature:Blob[179] ParamList:ParamRow[20]}
	&{RVA:10344 ImplFlags:0 Flags:2182 Name:get_Width Signature:Blob[179] ParamList:ParamRow[20]}
	&{RVA:10352 ImplFlags:0 Flags:2182 Name:get_Transparent Signature:Blob[179] ParamList:ParamRow[20]}
	&{RVA:10360 ImplFlags:0 Flags:2182 Name:set_Transparent Signature:Blob[183] ParamList:ParamRow[20]}
ParamRow
	&{Flags:0 Sequence:1 Name:width}
	&{Flags:0 Sequence:2 Name:height}
	&{Flags:0 Sequence:1 Name:stream}
	&{Flags:0 Sequence:2 Name:offset}
	&{Flags:0 Sequence:3 Name:length}
	&{Flags:0 Sequence:1 Name:stream}
	&{Flags:0 Sequence:2 Name:offset}
	&{Flags:0 Sequence:3 Name:length}
	&{Flags:0 Sequence:4 Name:crc}
	&{Flags:0 Sequence:1 Name:x}
	&{Flags:0 Sequence:2 Name:y}
	&{Flags:0 Sequence:3 Name:pixel}
	&{Flags:0 Sequence:1 Name:x}
	&{Flags:0 Sequence:2 Name:y}
	&{Flags:0 Sequence:1 Name:index}
	&{Flags:0 Sequence:1 Name:index}
	&{Flags:0 Sequence:2 Name:color}
	&{Flags:0 Sequence:1 Name:crc}
	&{Flags:0 Sequence:1 Name:filename}
	&{Flags:0 Sequence:1 Name:value}
MemberRefRow
	&{Class:TypeRefRow[1] Name:.ctor Signature:Blob[15]}
	&{Class:TypeRefRow[2] Name:.ctor Signature:Blob[19]}
	&{Class:TypeRefRow[4] Name:InitializeArray Signature:Blob[28]}
	&{Class:TypeRefRow[8] Name:.ctor Signature:Blob[19]}
	&{Class:TypeRefRow[9] Name:.ctor Signature:Blob[19]}
	&{Class:TypeRefRow[10] Name:.ctor Signature:Blob[44]}
	&{Class:TypeRefRow[12] Name:Write Signature:Blob[51]}
	&{Class:TypeRefRow[13] Name:.ctor Signature:Blob[15]}
	&{Class:TypeRefRow[12] Name:WriteByte Signature:Blob[59]}
	&{Class:TypeRefRow[14] Name:.ctor Signature:Blob[64]}
	&{Class:TypeRefRow[16] Name:Dispose Signature:Blob[15]}
	&{Class:TypeRefRow[13] Name:ToArray Signature:Blob[73]}
	&{Class:TypeRefRow[17] Name:.ctor Signature:Blob[15]}
	&{Class:TypeRefRow[19] Name:.ctor Signature:Blob[15]}
CustomAttributeRow
	&{Parent:AssemblyRow[1] Type:MemberRefRow[14] Value:Blob[192]}
	&{Parent:TypeDefRow[3] Type:MemberRefRow[13] Value:Blob[78]}
ClassLayoutRow
	&{PackingSize:1 ClassSize:32 Parent:TypeDefRow[4]}
	&{PackingSize:1 ClassSize:4 Parent:TypeDefRow[5]}
	&{PackingSize:1 ClassSize:12 Parent:TypeDefRow[6]}
StandAloneSigRow
	&{Signature:Blob[98]}
	&{Signature:Blob[118]}
	&{Signature:Blob[133]}
	&{Signature:Blob[162]}
PropertyMapRow
	&{Parent:TypeDefRow[2] PropertyList:PropertyRow[1]}
PropertyRow
	&{Flags:0 Name:Height Type:Blob[188]}
	&{Flags:0 Name:Width Type:Blob[188]}
	&{Flags:0 Name:Transparent Type:Blob[188]}
MethodSemanticsRow
	&{Semantics:2 Method:MethodDefRow[10] Association:PropertyRow[1]}
	&{Semantics:2 Method:MethodDefRow[11] Association:PropertyRow[2]}
	&{Semantics:1 Method:MethodDefRow[13] Association:PropertyRow[3]}
	&{Semantics:2 Method:MethodDefRow[12] Association:PropertyRow[3]}
FieldRVARow
	&{RVA:16384 Field:FieldRow[14]}
	&{RVA:16416 Field:FieldRow[15]}
	&{RVA:16420 Field:FieldRow[16]}
	&{RVA:16432 Field:FieldRow[17]}
	&{RVA:16436 Field:FieldRow[18]}
	&{RVA:16440 Field:FieldRow[19]}
	&{RVA:16444 Field:FieldRow[20]}
	&{RVA:16448 Field:FieldRow[21]}
AssemblyRow
	&{HashAlgId:SHA1 MajorVersion:0 MinorVersion:0 BuildNumber:0 RevisionNumber:0 Flags:0 PublicKey:Blob[0] Name:Png4BitIndexed}
AssemblyRefRow
	&{MajorVersion:0 MinorVersion:2 BuildNumber:0 RevisionNumber:0 Flags:0 PublicKeyOrToken:Blob[0] Name:et_Width Culture:mscorlib HashValue:Blob[0]}
	&{MajorVersion:0 MinorVersion:2 BuildNumber:0 RevisionNumber:0 Flags:0 PublicKeyOrToken:Blob[0] Name:et_Width Culture:System HashValue:Blob[0]}
NestedClassRow
	&{NestedClass:TypeDefRow[0] EnclosingClass:TypeDefRow[4]}
	&{NestedClass:TypeDefRow[3] EnclosingClass:TypeDefRow[5]}
	&{NestedClass:TypeDefRow[3] EnclosingClass:TypeDefRow[6]}
	<Module>
	PeterO.Png4BitIndexed
	<PrivateImplementationDetails>{dc74ea34-d62f-4136-93bf-797df4217b8e}
	$ArrayType=32
	$ArrayType=4
	$ArrayType=12
<Module>
Png4BitIndexed
	extends System.Object
	private byte[] subdata1
	private uint[] crcTable
	private int width
	private int height
	private int realRowSize
	private int blockSize
	private int rowSize
	private byte[] imageData
	private byte[] data
	private byte[] subdata2
	private byte[] colors
	private uint idatCrc
	private int transparent
	public void  .ctor(int width, int height)
	private byte[]  Adler32(byte[] stream, int offset, int length)
	private uint  Crc32(byte[] stream, int offset, int length, uint crc)
	public void  SetPixel(int x, int y, int pixel)
	public int  GetPixel(int x, int y)
	public byte[]  GetColor(int index)
	public void  SetColor(int index, byte[] color)
	private byte[]  GetBE(uint crc)
	public void  Save(string filename)
	public int  get_Height()
	public int  get_Width()
	public int  get_Transparent()
	public void  set_Transparent(int value)
<PrivateImplementationDetails>{dc74ea34-d62f-4136-93bf-797df4217b8e}
	extends System.Object
	public static $ArrayType=32 $field-0
	public static $ArrayType=4 $field-1
	public static $ArrayType=12 $field-2
	public static $ArrayType=4 $field-3
	public static $ArrayType=4 $field-4
	public static $ArrayType=4 $field-5
	public static $ArrayType=4 $field-6
	public static $ArrayType=4 $field-7
$ArrayType=32
	extends System.ValueType
$ArrayType=4
	extends System.ValueType
$ArrayType=12
	extends System.ValueType
