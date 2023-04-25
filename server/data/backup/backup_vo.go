package backup

import (
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// StoreVO 是用于数据导入/导出的 JSON 结构
type StoreVO struct {
	WPM dto.Module `json:"WPM"`

	Tables StoreTables `json:"Tables"`
}

// StoreTables 承载所有的数据表
type StoreTables struct {
	ExecutableTable       []*entity.Executable       `json:"Executables"`
	IntentTemplateTable   []*entity.IntentTemplate   `json:"IntentTemplates"`
	LocalRepositoryTable  []*entity.LocalRepository  `json:"Repositories"`
	LocationTable         []*entity.Location         `json:"Locations"`
	MediaTable            []*entity.Media            `json:"Mediae"`
	SourceTable           []*entity.Namespace        `json:"Sources"`
	ProjectTable          []*entity.Project          `json:"Projects"`
	ContentTypeTable      []*entity.ContentType      `json:"ContentTypes"`
	RemoteRepositoryTable []*entity.RemoteRepository `json:"Remotes"`
	SettingTable          []*entity.Setting          `json:"Settings"`
	UserTable             []*entity.User             `json:"Users"`
	WorktreeTable         []*entity.Worktree         `json:"Worktrees"`
	PackageTable          []*entity.SoftwarePackage  `json:"Packages"`
}

// StartupVO 这个VO用来记录 wpm 的启动日志
type StartupVO struct {
	WPM      dto.Module                    `json:"WPM"`
	Versions map[string]*entity.Executable `json:"Versions"` // map[version] exe
	// Logs     []*entity.Executable          `json:"logs"`
}
