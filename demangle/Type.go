// @File: Type.go
// @Author: Jason
// @Date: 2022/8/8

package demangle

// TypeName 类型名类型，底层为string
type TypeName string

// SType S开头的字符类型
var SType = map[byte]TypeName{
	'i': "Swift.Int",
	'b': "Swift.Bool",
	'f': "Swift.Float",
	'S': "Swift.String",
	'a': "Swift.Array",
	'q': "Swift.Optional",
	'Q': "Swift.ImplicitlyUnwrappedOptional",
	'r': "Swift.UnsafeMutableBufferPointer",
	'R': "Swift.UnsafeBufferPointer",
	'u': "Swift.UInt",
	'o': "__C",
	'p': "Swift.UnsafeMutablePointer",
	'P': "Swift.UnsafePointer",
	'd': "Swift.Double",
	'c': "Swift.UnicodeScalar",
	'C': "__C_Synthesized",
	'v': "Swift.UnsafeMutableRawPointer",
	'V': "Swift.UnsafeRawPointer",
}

// BType B开头的类型
var BType = map[byte]TypeName{
	'O': "Builtin.UnknownObject",
	'B': "Builtin.UnsafeValueBuffer",
	'w': "Builtin.Word",
	't': "Builtin.SILToken",
	'o': "Builtin.NativeObject",
	'p': "Builtin.RawPointer",
	'b': "Builtin.BridgeObject",
}

// AccessorL1 长度为1的符号的函数类型
var AccessorL1 = map[byte]TypeName{
	's': "setter",
	'm': "materializeForSet",
	'g': "getter",
	'G': "getter",
	'w': "willset",
	'W': "didset",
}

// AccessorL2 长度为2的符号的函数类型前缀
var AccessorL2 = map[byte]TypeName{
	'a': "MutableAddressor",
	'l': "Addressor",
}

// Addressor 长度为2的符号的函数类型后缀
var Addressor = map[byte]TypeName{
	'u': "unsafe",        // unsafe addressor (no owner)
	'O': "owning",        // owning addressor (non-native owner), not used anymore
	'o': "nativeOwning",  // owning addressor (native owner), not used anymore
	'p': "nativePinning", // pinning addressor (native owner), not used anymore
}
