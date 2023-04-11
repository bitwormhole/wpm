package utils

import "strings"

// URLResolver 是一个简单的URL解析器，如果可以，把相对URL解析为绝对URL
type URLResolver struct {
	docURL string
}

// ForDocument 设置参考文档的URL
func (inst *URLResolver) ForDocument(url string) *URLResolver {
	inst.docURL = url
	return inst
}

// Resolve 解析传入的URL
func (inst *URLResolver) Resolve(url string) string {

	doc := inst.docURL
	if inst.isAbsoluteURL(url) {
		return url
	} else if !inst.isAbsoluteURL(doc) {
		return url
	}

	b := &URLBuilder{}
	b.Init(doc)
	b.Query = ""
	b.Fragment = ""

	if inst.isAbsolutePath(url) {
		b.Path = url
	} else {
		// i1 := strings.IndexRune(url, '?')
		// i2 := strings.IndexRune(url, '#')
		path := b.Path + "/../" + url
		pb := PathBuilder{}
		pb.Init(path).Normalize()
		b.Path = pb.CreateAbsolute()
	}

	return b.Create()
}

func (inst *URLResolver) isAbsoluteURL(url string) bool {
	return strings.Contains(url, "://")
}

func (inst *URLResolver) isAbsolutePath(url string) bool {
	return strings.HasPrefix(url, "/")
}
