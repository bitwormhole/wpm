package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// RepositoryImportAction 表示导入仓库的动作
type RepositoryImportAction string

// 定义导入仓库的动作
const (
	RepositoryImportActionFind   RepositoryImportAction = "find"
	RepositoryImportActionLocate RepositoryImportAction = "locate"
	RepositoryImportActionSave   RepositoryImportAction = "save"
)

// RepositoryImport ... 用于导入仓库
type RepositoryImport struct {
	Base

	Action RepositoryImportAction `json:"action"`
	Items  []*dto.LocalRepository `json:"local_repositories"`

	// Target *dto.Repository        `json:"target"`
}

////////////////////////////////////////////////////////////////////////////////

func (value RepositoryImportAction) String() string {
	return string(value)
}
