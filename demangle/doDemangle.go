// @File: doDemangle.go
// @Author: Jason
// @Date: 2022/8/4

package demangle

import (
	"errors"
)

// ToString 将mangle的符号输入，demangle后输出
func ToString(input string, options ...Option) (string, error) {
	var prefix string
	var is bool
	if prefix, is = HasSwiftPrefix(input); !is {
		return "", errors.New("get invalid swift mangle parser")
	}
	parser := NewParser(input, len(prefix))

	asts := ToAST(&parser)
	result, err := ASTToString(asts, options...)
	return result, err
}

// ASTToString 将AST转化为string
func ASTToString(asts []AST, options ...Option) (string, error) {
	ps := &PrintState{}
	for _, a := range asts {
		switch a.(type) {
		case *Err:
			e := a.(*Err)
			return "", e
		default:
			a.Print(ps)
		}
	}
	return ps.buf.String(), nil
}
