package backup

import (
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// VO 是用于数据导入/导出的 JSON 结构
type VO struct {
	WPM dto.Module `json:"wpm"`

	ExecutableTable       []*entity.Executable       `json:"executables"`
	IntentTemplateTable   []*entity.IntentTemplate   `json:"intenttemplates"`
	LocalRepositoryTable  []*entity.LocalRepository  `json:"repositories"`
	LocationTable         []*entity.Location         `json:"locations"`
	MediaTable            []*entity.Media            `json:"mediae"`
	NamespaceTable        []*entity.Namespace        `json:"namespaces"`
	ProjectTable          []*entity.Project          `json:"projects"`
	ContentTypeTable      []*entity.ContentType      `json:"contenttypes"`
	RemoteRepositoryTable []*entity.RemoteRepository `json:"remotes"`
	SettingTable          []*entity.Setting          `json:"settings"`
	UserTable             []*entity.User             `json:"users"`
	WorktreeTable         []*entity.Worktree         `json:"worktrees"`
	PackageTable          []*entity.SoftwarePackage  `json:"packages"`
}

// StartupVO 这个VO用来记录 wpm 的启动日志
type StartupVO struct {
	WPM      dto.Module                    `json:"wpm"`
	Versions map[string]*entity.Executable `json:"versions"` // map[version] exe
	Logs     []*entity.Executable          `json:"logs"`
}
