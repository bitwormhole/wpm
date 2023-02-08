package dto

// File 表示一个文件或者文件夹的信息
type File struct {
	Name string `json:"name"`

	Base

	Exists bool `json:"exists"`
	IsFile bool `json:"is_file"`
	IsDir  bool `json:"is_dir"`

	Size int64  `json:"size"`
	Type string `json:"content_type"`
}
