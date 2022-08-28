// @File: utils.go
// @Author: Jason
// @Date: 2022/8/5

package demangle

import "strings"

// Option 参数的设置
type Option int

// HasSwiftPrefix 是否有Swift demangle所需的前缀
func HasSwiftPrefix(input string) (string, bool) {
	prefix := []string{"$s", "_T", "$S", "_T0"}
	for _, pre := range prefix {
		if strings.HasPrefix(input, pre) {
			return pre, true
		}
	}
	return "", false
}

// IsDigit 是否为数字字节
func IsDigit(c byte) bool {
	return c <= '9' && c >= '0'
}
