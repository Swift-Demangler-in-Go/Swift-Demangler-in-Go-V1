// @File: ast.go
// @Author: Jason
// @Date: 2022/8/4

package demangle

// AST 单个AST结点
type AST interface {
	Print(*PrintState) // 向PrintState中打印消息
}
