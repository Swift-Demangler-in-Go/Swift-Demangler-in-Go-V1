// @File: Parser.go
// @Author: Jason
// @Date: 2022/8/8

package demangle

import (
	"fmt"
)

// Parser 解析类，封装输入的mangle字符，并记录demangle相关的信息
type Parser struct {
	input    string // mangle后的字符
	cur      int    // 当前解析位置的指针
	NodeInfo        // 记录当前结点的信息
}

// NewParser 构造函数
func NewParser(input string, offset int) Parser {
	p := Parser{
		input: input,
		cur:   offset,
	}
	p.S = make(map[int]string)
	return p
}

// GetAST 获取当前解析的AST结点
func (p *Parser) GetAST() AST {
	var a AST
	var name string
	for a == nil && !p.End() {
		b := p.GetByte()
		if IsDigit(b) {
			name = p.GetName()
			if !p.isPack {
				p.isPack = true
				a = &Pack{name: name}
				p.AddS(name)
				return a
			}
			a = p.GetAST()
			p.SetASTName(a, name)
		}
		// 关键字字符
		switch b {
		case 'C', 'V', 'O':
			a = p.GetClassAST()
		case 'F', 'f':
			a = p.GetFuncAST()
		case 'S':
			a = p.GetSAST()
		case 'T', 't':
			a = p.GetTupleAST()
		case 'B':
			a = p.GetBAST()
		default:
			if !p.onArg { // 在未开启接收函数前判断，减少开销
				typ := p.GetFuncType() // 获取函数的类型名称
				if typ != "" {         // 函数是否存在类型名称
					name := p.GetName()
					a = p.GetFuncAST()
					a.(*Func).typ = typ
					a.(*Func).SetName(name)
				}
			}
		}
	}
	return a
}

// GetFuncAST 解析获取函数AST结点
func (p *Parser) GetFuncAST() AST {
	if !p.onFun { // 判断是否正在解析函数，即是否已经遇到第一个F，如_TFCCxxxx
		p.onFun = true
		p.Next()
		return nil
	} else if !p.onArg {
		// 开启全局解析参数及返回值列表状态
		p.onArg = true
		defer func() { p.onArg = false }()
	}
	p.Next()
	f := &Func{}
	f.Arg = p.GetAST()
	f.Return = p.GetAST()
	return f
}

// GetSAST 解析获取以S开头的符号信息
func (p *Parser) GetSAST() AST {
	var a AST
	b := p.NextByte()
	if typ, is := SType[b]; is {
		a = &Name{name: string(typ)}
		p.Next()
	} else if b == '_' {
		p.Next()
		return &Name{name: p.GetS(-1)}
	} else if IsDigit(b) {
		si := p.GetNum()
		a = &Name{name: p.GetS(si)}
		p.Next()
	} else {
		a = &Err{fmt.Sprintf("Invalid symbol %c at %d", b, p.cur)}
	}
	return a
}

// GetClassAST 解析获取类AST结点
func (p *Parser) GetClassAST() AST {
	defer func() {
		p.isClass = true
		p.isPack = true
	}()
	var c Class
	var name string
	nClass := 0
	b := p.GetByte()
	for b == 'C' || b == 'O' || b == 'V' {
		nClass++
		b = p.NextByte()
	}
	name = p.GetName()
	if !p.isClass {
		p.AddS(name)
	}
	c = Class{pack: Pack{name: name}}

	for nClass > 0 {
		name = p.GetName()
		if !p.isClass {
			p.AddS(name)
		}
		a := Name{name: name}
		c.AddClass(a)
		nClass--
	}
	return &c
}

// GetTupleAST 解析获取元组结点
func (p *Parser) GetTupleAST() AST {
	var t Tuple
	if p.GetByte() == 't' {
		t.isVarLen = true
	}
	b := p.NextByte()
	for b != '_' {
		a := p.GetAST()
		t.Append(a)
		b = p.GetByte()
	}
	p.Next()
	return &t
}

// GetBAST 获取以B字母开头的类型
func (p *Parser) GetBAST() AST {
	b := p.NextByte()
	var a AST
	if v, is := BType[b]; is {
		a = &Name{name: string(v)}
	} else {
		a = &Err{err: fmt.Sprintf("Invalid symbol %c at %d", b, p.cur)}
	}
	p.Next()
	return a
}

// GetByte 获取当前解析的字符
func (p *Parser) GetByte() byte {
	if p.cur >= len(p.input) {
		return 0
	}
	return p.input[p.cur]
}

// NextByte 获取下一个解析的字符，并将peek后移一位。
func (p *Parser) NextByte() byte {
	p.cur++
	return p.input[p.cur]
}

// GetNextByte 获取下一个解析的字符，peek不移动
func (p *Parser) GetNextByte() byte {
	return p.input[p.cur+1]
}

// End 解析是否结束
func (p *Parser) End() bool {
	return p.cur >= len(p.input)
}

// Next cur指针后移
func (p *Parser) Next() {
	p.cur++
}

// GetNum 当前解析为数值时，获取相应的数值；cur、peek对应更新
func (p *Parser) GetNum() int {
	val := 0
	for IsDigit(p.input[p.cur]) {
		val = val*10 + int(p.input[p.cur]-'0')
		p.Next()
	}
	return val
}

// GetName 获取数值后的名称，如5Class，返回Class
func (p *Parser) GetName() string {
	n := p.GetNum()
	name := p.input[p.cur : p.cur+n]
	p.cur += n
	return name
}

// Info 打印Parser相关的信息（调试时使用）
func (p Parser) Info() {
	msg := fmt.Sprintf("Parser(len: %d,"+
		" isEnd: %v,"+
		" cur: %d,"+
		" left: '%s')\n", len(p.input), p.End(), p.cur, p.input[p.cur:])
	fmt.Println(msg)
}

// SetASTName 根据结点类型设置AST名称
func (p *Parser) SetASTName(a AST, name string) {
	if name == "" { // 名称为空不用设置
		return
	}
	switch a.(type) { // 添加参数外部名
	case *Name:
		a.(*Name).SetSubName(name)
	case *Func:
		if p.onArg { // 函数是参数，设置参数别名
			a.(*Func).SetSubName(name)
		} else { // 否则设置函数名字
			a.(*Func).SetName(name)
		}
	case *Tuple:
		a.(*Tuple).SetName(name)
	case *Class:
		a.(*Class).SetSubName(name)
	default:
	}
}

// GetFuncType 获取函数类型
func (p *Parser) GetFuncType() TypeName {

	b := p.GetByte()
	if v, is := AccessorL1[b]; is {
		p.Next()
		return v
	} else if v, is := AccessorL2[b]; is {
		b = p.NextByte()
		v2, is := Addressor[b]
		if !is {
			return ""
		} else {
			p.Next()
			return v2 + v
		}
	}
	return ""
}
