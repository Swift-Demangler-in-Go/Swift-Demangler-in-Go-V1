// @File: Name.go
// @Author: Jason
// @Date: 2022/8/10

package demangle

// Name 命名AST结点，封装名字信息
type Name struct {
	name    string // 变量名
	subName string // 作为参数时，节点的外部参数名
}

// Print 实现AST接口，定义打印规则
func (n *Name) Print(ps *PrintState) {
	if n.subName != "" {
		ps.WriteString(n.subName + ": ")
	}
	ps.WriteString(n.name)
}

// SetName 设置节点名
func (n *Name) SetName(name string) {
	n.name = name
}

// SetSubName 设置外部参数名
func (n *Name) SetSubName(name string) {
	n.subName = name
}
