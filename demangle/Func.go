// @File: Func.go
// @Author: Jason
// @Date: 2022/8/10

package demangle

// Func 函数AST结点
type Func struct {
	Name            // 组合命名类型
	typ    TypeName // 函数类型，getter、setter
	Arg    AST      // 参数列表
	Return AST      // 返回值列表
}

// Print 函数结点的打印规则
func (f *Func) Print(ps *PrintState) {
	if f.subName != "" {
		ps.WriteString(f.subName)
		ps.WriteString(": ")
	} else if f.name != "" {
		ps.WriteByte('.')
		ps.WriteString(f.name)
	}
	if f.typ != "" {
		ps.WriteByte('.')
		ps.WriteString(string(f.typ))
		ps.WriteString(" : ")
	}

	switch f.Arg.(type) {
	case *Tuple:
		f.Arg.Print(ps)
	case nil:
	default:
		ps.WriteByte('(')
		f.Arg.Print(ps)
		ps.WriteByte(')')
	}

	ps.WriteString(" -> ")

	switch f.Return.(type) {
	case *Tuple:
		f.Return.Print(ps)
	case nil:
		ps.WriteString("()")
	default:
		f.Return.Print(ps)
	}

}
