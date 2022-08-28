// @File: Tuple.go
// @Author: Jason
// @Date: 2022/8/15

package demangle

// Tuple 元组类型AST
type Tuple struct {
	Name
	isVarLen bool  // 是否为变长元组
	ele      []AST // 元组内的元素
}

// Print 向PrintState内输入信息
func (t *Tuple) Print(ps *PrintState) {
	if t.subName != "" {
		ps.WriteString(t.subName)
		ps.WriteString(": ")
	}
	ps.WriteByte('(')
	for i, a := range t.ele {
		a.Print(ps)
		if i != len(t.ele)-1 {
			ps.WriteString(", ")
		}
	}
	if t.isVarLen && len(t.ele) > 0 {
		ps.WriteString("...")
	}
	ps.WriteByte(')')
}

// SetName 设置别名
func (t *Tuple) SetName(name string) {
	t.subName = name
}

// Append 向元组元素内添加元素
func (t *Tuple) Append(ele AST) {
	t.ele = append(t.ele, ele)
}
