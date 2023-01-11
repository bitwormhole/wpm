package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// ProjectImportAction 表示导入仓库的动作
type ProjectImportAction string

// 定义导入仓库的动作
const (
	ProjectImportActionFind   ProjectImportAction = "find"
	ProjectImportActionLocate ProjectImportAction = "locate"
	ProjectImportActionSave   ProjectImportAction = "save"
)

// ProjectImport ... 用于导入项目
type ProjectImport struct {
	Base

	Action   ProjectImportAction `json:"action"`
	Projects []*dto.Project      `json:"projects"`
}

func (value ProjectImportAction) String() string {
	return string(value)
}
