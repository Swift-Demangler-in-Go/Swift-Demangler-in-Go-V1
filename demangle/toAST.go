// @File: toAST.go
// @Author: Jason
// @Date: 2022/8/5

package demangle

// ToAST *Parser解析中获取AST结点
func ToAST(p *Parser) []AST {
	asts := make([]AST, 0)
	for a := p.GetAST(); a != nil; a = p.GetAST() {
		asts = append(asts, a)
	}
	return asts
}
