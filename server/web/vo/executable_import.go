package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// ExecutableImportAction 表示导入应用程序的动作
type ExecutableImportAction string

// 定义导入应用程序的动作
const (
	ExecutableImportActionFind   ExecutableImportAction = "find"
	ExecutableImportActionLocate ExecutableImportAction = "locate"
	ExecutableImportActionSave   ExecutableImportAction = "save"
)

// ExecutableImport ...
type ExecutableImport struct {
	Base

	Action      ExecutableImportAction `json:"action"`
	Executables []*dto.Executable      `json:"executables"`
}

func (value ExecutableImportAction) String() string {
	return string(value)
}
