// @File: repl.go
// @Author: Jason
// @Date: 2022/8/4

package main

import (
	"Swift-Demangle-in-Go-V1/demangle"
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	// Welcome 欢迎标语
	Welcome = "Hello! This is the swift demangler in go version!"
)

// Start 命令行函数入口
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		_, err := fmt.Fprintf(out, ">>")
		if err != nil {
			return
		}
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		result, err := demangle.ToString(line)
		if err != nil {
			fmt.Println("Opps, something goes wrong, err:", err)
			continue
		}
		msg := fmt.Sprintf("result: %s", result)
		_, err = fmt.Fprintln(out, msg)
		if err != nil {
			return
		}
	}
}

func main() {
	fmt.Println(Welcome)
	Start(os.Stdin, os.Stdout)
}
