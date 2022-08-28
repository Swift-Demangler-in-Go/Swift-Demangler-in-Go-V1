// @File: Pack.go
// @Author: Jason
// @Date: 2022/8/9

package demangle

// Pack 包名AST结点
type Pack struct {
	name string
}

// Print 打印规则
func (pk *Pack) Print(ps *PrintState) {
	ps.WriteString(pk.name)
	//ps.WriteByte('.')
}
