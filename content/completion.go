package content

const (
	FLAG_BOOL_SHIFT = 0
)

// Zero or more of these flags can be set at a time
const (
	FLAG_STATIC = 1 << (FLAG_BOOL_SHIFT + iota)
	FLAG_FINAL
	FLAG_CONSTRUCTOR
	FLAG_CONST
	FLAG_RESTRICT
	FLAG_REFERENCE
	FLAG_VOLATILE
	_FLAG_BOOL_COUNT = iota
	FLAG_BOOL_BITS   = _FLAG_BOOL_COUNT
	FLAG_BOOL_MASK   = ((1 << FLAG_BOOL_BITS) - 1) << FLAG_BOOL_SHIFT
	FLAG_ACC_SHIFT   = FLAG_BOOL_SHIFT + FLAG_BOOL_BITS
)

// Access flags, only one of these can be set
const (
	FLAG_ACC_NOTSET = iota << FLAG_ACC_SHIFT
	FLAG_ACC_PUBLIC
	FLAG_ACC_PRIVATE
	FLAG_ACC_PROTECTED
	_FLAG_ACC_COUNT = iota
	FLAG_ACC_BITS   = 2
	FLAG_ACC_MASK   = ((1 << FLAG_ACC_BITS) - 1) << FLAG_ACC_SHIFT
	FLAG_TYPE_SHIFT = FLAG_ACC_SHIFT + FLAG_ACC_BITS
)

// Type flags, only one of these can be set
const (
	FLAG_TYPE_NOTSET = iota << FLAG_TYPE_SHIFT
	FLAG_TYPE_CLASS
	FLAG_TYPE_INTERFACE
	FLAG_TYPE_PACKAGE
	FLAG_TYPE_ARRAY
	FLAG_TYPE_POINTER
	FLAG_TYPE_METHOD
	FLAG_TYPE_STRUCT
	FLAG_TYPE_ENUM
	_FLAG_TYPE_COUNT = iota
	FLAG_TYPE_BITS   = 4
	FLAG_TYPE_MASK   = ((1 << FLAG_TYPE_BITS) - 1) << FLAG_TYPE_SHIFT
	FLAG_LANG_SHIFT  = FLAG_TYPE_SHIFT + FLAG_TYPE_BITS
)

const (
	FLAG_LANG_NOTSET = iota << FLAG_LANG_SHIFT
	FLAG_LANG_C
	FLAG_LANG_CXX
	FLAG_LANG_OBJC
	FLAG_LANG_OBJCXX
	FLAG_LANG_JAVA
	FLAG_LANG_NET
	_FLAG_LANG_COUNT = iota
	FLAG_LANG_BITS   = 3
	FLAG_LANG_MASK   = ((1 << FLAG_LANG_BITS) - 1) << FLAG_LANG_SHIFT
)

type (
	Flags uint32

	Variable struct {
		Name FullyQualifiedName `protocol:"required" json:",omitempty"`
		Type Type               `protocol:"required" json:",omitempty"`
	}

	Field struct {
		Variable `protocol:"required" json:",omitempty"`
		Flags    Flags `protocol:"optional" json:",omitempty"`
	}

	Method struct {
		Name           FullyQualifiedName `protocol:"required" json:",omitempty"`
		Flags          Flags              `protocol:"optional" json:",omitempty"`
		Returns        []Variable         `protocol:"optional" json:",omitempty"`
		Parameters     []Variable         `protocol:"optional" json:",omitempty"`
		Specialization []Type             `protocol:"optional" json:",omitempty"`
	}

	// TODO: Is there any instance where the completion result would need to contain
	// 	     fields *not* in Type?
	CompletionResult Type
	Type             struct {
		Name           FullyQualifiedName `protocol:"optional" json:",omitempty"`
		Specialization []Type             `protocol:"optional" json:",omitempty"`
		Flags          Flags              `protocol:"optional" json:",omitempty"`
		Methods        []Method           `protocol:"optional" json:",omitempty"`
		Fields         []Field            `protocol:"optional" json:",omitempty"`
		Types          []Type             `protocol:"optional" json:",omitempty"`
		Extends        []Type             `protocol:"optional" json:",omitempty"`
		Implements     []Type             `protocol:"optional" json:",omitempty"`
	}

	Available interface {
		Available(*Args, *bool) error
	}

	Complete interface {
		Complete(*CompleteArgs, *CompletionResult) error
	}

	CompleteAt interface {
		CompleteAt(*CompleteAtArgs, *CompletionResult) error
	}
)
