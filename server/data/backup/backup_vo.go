package backup

import (
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// VO 是用于数据导入/导出的 JSON 结构
type VO struct {
	WPM dto.Module

	ExecutableTable       []*entity.Executable
	IntentTemplateTable   []*entity.IntentTemplate
	LocalRepositoryTable  []*entity.LocalRepository
	LocationTable         []*entity.Location
	MediaTable            []*entity.Media
	NamespaceTable        []*entity.Namespace
	ProjectTable          []*entity.Project
	ContentTypeTable      []*entity.ContentType
	RemoteRepositoryTable []*entity.RemoteRepository
	SettingTable          []*entity.Setting
	UserTable             []*entity.User
	WorktreeTable         []*entity.Worktree
}

// StartupVO 这个VO用来记录 wpm 的启动日志
type StartupVO struct {
	WPM dto.Module

	Versions map[string]*entity.Executable // map[version] exe

	Logs []*entity.Executable
}
