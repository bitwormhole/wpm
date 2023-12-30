package dxo

import "strings"

// StringListCRLF 表示一个包含多个词组的字符串，各个元素之间以 CR|LF 分隔
type StringListCRLF string

func (value StringListCRLF) String() string {
	return string(value)
}

// Array ...
func (value StringListCRLF) Array() []string {
	const (
		c1 = "\r"
		c2 = "\n"
	)
	text := value.String()
	text = strings.ReplaceAll(text, c1, c2)
	src := strings.Split(text, c2)
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

// NewStringListCRLF ... 根据 src 创建 StringListCRLF
func NewStringListCRLF(src []string) StringListCRLF {
	builder := strings.Builder{}
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item != "" {
			builder.WriteString(item)
			builder.WriteString("\n")
		}
	}
	text := builder.String()
	return StringListCRLF(text)
}
