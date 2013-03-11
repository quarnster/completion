package content

type Type struct {
	Name           FullyQualifiedName `protocol:"optional" json:",omitempty"`
	Specialization []Type             `protocol:"optional" json:",omitempty"`
	Flags          Flags              `protocol:"optional" json:",omitempty"`
	Methods        []Method           `protocol:"optional" json:",omitempty"`
	Fields         []Field            `protocol:"optional" json:",omitempty"`
	Types          []Type             `protocol:"optional" json:",omitempty"`
	Extends        []Type             `protocol:"optional" json:",omitempty"`
	Implements     []Type             `protocol:"optional" json:",omitempty"`
}

const (
	FLAG_ACC_SHIFT  = 0
	FLAG_ACC_BITS   = 2
	FLAG_ACC_MASK   = 0x3 << FLAG_ACC_SHIFT
	FLAG_TYPE_SHIFT = FLAG_ACC_BITS
	FLAG_TYPE_BITS  = 3
	FLAG_TYPE_MASK  = 0x7 << FLAG_TYPE_SHIFT
)
const (
	FLAG_ACC_NOTSET = iota << FLAG_ACC_SHIFT
	FLAG_ACC_PUBLIC
	FLAG_ACC_PRIVATE
	FLAG_ACC_PROTECTED
)
const (
	FLAG_TYPE_NOTSET = iota << FLAG_TYPE_SHIFT
	FLAG_TYPE_CLASS
	FLAG_TYPE_INTERFACE
	FLAG_TYPE_PACKAGE
	FLAG_TYPE_ARRAY
	FLAG_TYPE_POINTER
)
const (
	FLAG_STATIC = 1 << (FLAG_ACC_BITS + FLAG_TYPE_BITS + iota)
	FLAG_FINAL
	FLAG_CONSTRUCTOR
)

type Flags uint32

type Variable struct {
	Name FullyQualifiedName `protocol:"required" json:",omitempty"`
	Type Type               `protocol:"required" json:",omitempty"`
}

type Field struct {
	Variable `protocol:"required" json:",omitempty"`
	Flags    Flags `protocol:"optional" json:",omitempty"`
}

type Method struct {
	Name           FullyQualifiedName `protocol:"required" json:",omitempty"`
	Flags          Flags              `protocol:"optional" json:",omitempty"`
	Returns        []Variable         `protocol:"optional" json:",omitempty"`
	Parameters     []Variable         `protocol:"optional" json:",omitempty"`
	Specialization []Type             `protocol:"optional" json:",omitempty"`
}

type CompletionResult struct {
	Types   []Type   `protocol:"optional" json:",omitempty"`
	Fields  []Field  `protocol:"optional" json:",omitempty"`
	Methods []Method `protocol:"optional" json:",omitempty"`
}
