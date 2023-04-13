package utils

import (
	"fmt"
	"strings"
)

// PathBuilder ...
type PathBuilder struct {
	sep      string
	elements []string
}

// Init ...
func (inst *PathBuilder) Init(p string) *PathBuilder {
	inst.elements = inst.sep2elements(p)
	return inst
}

// Normalize 执行标准化
func (inst *PathBuilder) Normalize() error {
	src := inst.elements
	dst := make([]string, 0)
	for _, el := range src {
		part := strings.TrimSpace(el)
		if part == "" {
			// skip
		} else if part == "." {
			// skip
		} else if part == ".." {
			// pop one
			size := len(dst)
			if size < 1 {
				p := inst.Create()
				return fmt.Errorf("bad absolute path: " + p)
			}
			dst = dst[0 : size-1]
		} else {
			dst = append(dst, el)
		}
	}
	inst.elements = dst
	return nil
}

// Append ...
func (inst *PathBuilder) Append(list ...string) *PathBuilder {
	for _, item := range list {
		if strings.ContainsAny(item, "/\\") {
			sublist := inst.sep2elements(item)
			inst.elements = append(inst.elements, sublist...)
		} else {
			inst.elements = append(inst.elements, item)
		}
	}
	return inst
}

// Create ...
func (inst *PathBuilder) Create() string {
	b := strings.Builder{}
	src := inst.elements
	sep := inst.getSep()
	for idx, item := range src {
		if idx > 0 {
			b.WriteString(sep)
		}
		b.WriteString(item)
	}
	return b.String()
}

// CreateAbsolute ...
func (inst *PathBuilder) CreateAbsolute() string {
	sep := inst.getSep()
	path := inst.Create()
	if strings.HasPrefix(path, sep) {
		return path
	}
	return sep + path
}

func (inst *PathBuilder) getSep() string {
	sep := inst.sep
	if sep == "" {
		sep = "/"
	}
	return sep
}

func (inst *PathBuilder) sep2elements(p string) []string {
	const (
		sep1 = "\\"
		sep2 = "/"
	)
	p2 := strings.ReplaceAll(p, sep1, sep2)
	return strings.Split(p2, sep2)
}

// IsAbsolutePath 判断是否为（本地）绝对路径
func IsAbsolutePath(path string) bool {

	if strings.HasPrefix(path, "/") {
		return true
	}

	// for windows
	i := strings.Index(path, ":\\")
	if i == 1 {
		return true
	}

	return false
}
