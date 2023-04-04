package utils

import "strings"

// Strings  一组面向字符串的工具
type Strings struct{}

// TrimToLower  移除两端空白，并准换为小写字符
func (inst *Strings) TrimToLower(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}
