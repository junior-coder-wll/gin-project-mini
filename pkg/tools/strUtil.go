//@Author: wulinlin
//@Description:
//@File:  strUtil
//@Version: 1.0.0
//@Date: 2023/03/20 14:41

package tools

import (
	"bytes"
	"strings"
)

// 判断字符串是否为空或空白
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// 判断字符串是否不为空或空白
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// 拼接多个字符串为一个字符串
func Join(sep string, strs ...string) string {
	var buffer bytes.Buffer
	for i, s := range strs {
		if i > 0 {
			buffer.WriteString(sep)
		}
		buffer.WriteString(s)
	}
	return buffer.String()
}

// 反转字符串
func Reverse(str string) string {
	var buffer bytes.Buffer
	for i := len(str) - 1; i >= 0; i-- {
		buffer.WriteByte(str[i])
	}
	return buffer.String()
}

// 判断字符串是否包含指定子串
func Contains(str string, substr string) bool {
	return strings.Contains(str, substr)
}
