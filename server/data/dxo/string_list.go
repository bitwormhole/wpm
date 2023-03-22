package dxo

import "strings"

// StringList 表示一个包含多个词组的字符串，各个元素之间以逗号分隔
type StringList string

func (value StringList) String() string {
	return string(value)
}

// Array ...
func (value StringList) Array() []string {
	const sep = ","
	text := value.String()
	src := strings.Split(text, sep)
	dst := make([]string, 0)
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item != "" {
			dst = append(dst, item)
		}
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

// NewStringList ... 根据 src 创建 StringList
func NewStringList(src []string) StringList {
	const sep = ","
	builder := strings.Builder{}
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item != "" {
			builder.WriteString(item)
			builder.WriteString(sep)
		}
	}
	text := builder.String()
	return StringList(text)
}
