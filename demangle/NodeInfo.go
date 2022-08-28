// @File: NodeInfo.go
// @Author: Jason
// @Date: 2022/8/9

package demangle

// NodeInfo 使用NodeInfo记录当前解析的进度和信息。
type NodeInfo struct {
	isPack  bool           // 是否完成包名的解析
	isClass bool           // 是否完成类的解析
	onFun   bool           // 是否正在解析函数
	onArg   bool           // 是否正在解析参数/返回值
	nS      int            // 记录Si_替代的个数
	S       map[int]string // 替代符号表，如S_: Pack; S0_: Class;
}

// AddS 向替代符号表中添加替代对应的符号
func (n *NodeInfo) AddS(name string) {
	if n.nS == 0 {
		n.S[-1] = name
	} else {
		n.S[n.nS-1] = n.S[n.nS-2] + "." + name
	}
	n.nS++
}

// GetS 获取Si_替代
func (n *NodeInfo) GetS(i int) string {
	return n.S[i]
}
