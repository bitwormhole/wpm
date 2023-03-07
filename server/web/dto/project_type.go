package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// ProjectType 表示一个命令模板
type ProjectType struct {
	ID dxo.ProjectTypeID `json:"id"`
	Base

	Name        string `json:"name"`
	Type        string `json:"type"`
	Label       string `json:"label"`
	Description string `json:"description"`

	AsFile bool `json:"as_file"`
	AsDir  bool `json:"as_dir"`
}
