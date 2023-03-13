package service

// ContentTypeService ...
type ContentTypeService interface {

	// 根据文件名（或路径）查询对应的mime类型
	GetContentType(name string) (string, error)
}
