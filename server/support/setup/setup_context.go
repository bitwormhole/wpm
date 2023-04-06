package setup

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/service"
)

// Context ...
type Context struct {
	StartAt util.Time

	AppDataService           service.AppDataService
	FileSystemService        service.FileSystemService
	ExecutableImportService  service.ExecutableImportService
	MediaService             service.MediaService
	ProjectTypeImportService service.ProjectTypeImportService
	SettingService           service.SettingService
	IntentTemplateService    service.IntentTemplateService
	SoftwarePackageService   service.SoftwarePackageService
}
