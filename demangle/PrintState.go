// @File: PrintState.go
// @Author: Jason
// @Date: 2022/8/5

package demangle

import "strings"

// PrintState 信息打印上下文
type PrintState struct {
	buf strings.Builder // strings.Builder
}

// WriteString 向buf中写入字符
func (p *PrintState) WriteString(s string) {
	p.buf.WriteString(s)
}

// WriteByte 向buf中写入字节
func (p *PrintState) WriteByte(b byte) {
	p.buf.WriteByte(b)
}
