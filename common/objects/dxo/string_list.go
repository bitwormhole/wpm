package dxo

import (
	"sort"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

type StringArray []string

func (value StringArray) String() string {
	return value.StringList().String()
}

func (value StringArray) Array() []string {
	return value
}

func (value StringArray) StringList() StringList {
	return NewStringList(value)
}

// Deduplication 排重
func (value StringArray) Deduplication() StringArray {
	src := value
	dst := make([]string, 0)
	tab := make(map[string]bool)
	for _, item := range src {
		if tab[item] {
			continue
		}
		tab[item] = true
		dst = append(dst, item)
	}
	return dst
}

// Sort 排序
func (value StringArray) Sort() StringArray {
	src := value
	dst := make([]string, 0)
	for _, item := range src {
		dst = append(dst, item)
	}
	sort.Strings(dst)
	return dst
}

// Trim 去除空白项
func (value StringArray) Trim() StringArray {
	src := value
	dst := make([]string, 0)
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		dst = append(dst, item)
	}
	return dst
}

// Normalize 标准化 (排空 -> 排序 -> 排重)
func (value StringArray) Normalize() StringArray {
	value = value.Trim()
	value = value.Sort()
	value = value.Deduplication()
	return value
}

////////////////////////////////////////////////////////////////////////////////

// StringList 表示一个包含多个词组的字符串，各个元素之间以逗号分隔
type StringList string

func (value StringList) String() string {
	return string(value)
}

func (value StringList) StringArray() StringArray {
	return value.Array()
}

// Array ...
func (value StringList) Array() []string {
	const sep = ","
	text := value.String()
	src := strings.Split(text, sep)
	dst := make([]string, 0)
	for _, item := range src {
		item = strings.TrimSpace(item)
		dst = append(dst, item)
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

// NewStringList ... 根据 src 创建 StringList
func NewStringList(src []string) StringList {
	const sep = ","
	builder := strings.Builder{}
	for i, item := range src {
		if i > 0 {
			builder.WriteString(sep)
		}
		builder.WriteString(item)
	}
	text := builder.String()
	return StringList(text)
}
