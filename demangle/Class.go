// @File: Class.go
// @Author: Jason
// @Date: 2022/8/9

package demangle

// Class 类AST结点
type Class struct {
	Name         // 组合命名类型
	pack  Pack   // 类的包名
	child []Name // 内部类的类名
}

// Print Class结点的打印规则
func (c *Class) Print(ps *PrintState) {
	if c.subName != "" {
		ps.WriteString(c.subName)
		ps.WriteString(": ")
	}
	c.pack.Print(ps)
	if len(c.child) > 0 {
		ps.WriteByte('.')
	}
	for i, a := range c.child {
		ps.WriteString(a.name)
		if i != len(c.child)-1 { // len(c.child) > 0 &&
			ps.WriteByte('.')
		}
	}
}

// AddClass 添加子类结点
func (c *Class) AddClass(child Name) {
	//if child.name != "" {
	c.child = append(c.child, child)
	//}
}
