package setup

import "github.com/bitwormhole/wpm/server/service"

// Context ...
type Context struct {
	ExecutableImportService  service.ExecutableImportService
	MediaService             service.MediaService
	ProjectTypeImportService service.ProjectTypeImportService
	SettingService           service.SettingService
}
