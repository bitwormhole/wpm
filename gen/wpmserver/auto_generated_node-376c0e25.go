// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmserver

import (
	store0x8467b3 "github.com/bitwormhole/gitlib/git/store"
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
	caches0x9d186b "github.com/bitwormhole/wpm/server/components/caches"
	intents0x8557f3 "github.com/bitwormhole/wpm/server/components/intents"
	v3filters0xa6552a "github.com/bitwormhole/wpm/server/components/intents/v3filters"
	dao0x5af8d0 "github.com/bitwormhole/wpm/server/data/dao"
	dbagent0x9f90fb "github.com/bitwormhole/wpm/server/data/dbagent"
	service0x3e063d "github.com/bitwormhole/wpm/server/service"
	support0xf47d7f "github.com/bitwormhole/wpm/server/support"
	appruntime0x8dfe0a "github.com/bitwormhole/wpm/server/support/appruntime"
	backups0xe44d54 "github.com/bitwormhole/wpm/server/support/backups"
	caches0xd7996d "github.com/bitwormhole/wpm/server/support/caches"
	checkupdate0xea1855 "github.com/bitwormhole/wpm/server/support/checkupdate"
	contenttypes0x61ca37 "github.com/bitwormhole/wpm/server/support/contenttypes"
	executables0xd3773a "github.com/bitwormhole/wpm/server/support/executables"
	httpclient0xf20fe2 "github.com/bitwormhole/wpm/server/support/httpclient"
	impldao0x73998b "github.com/bitwormhole/wpm/server/support/impldao"
	implservice0x22327c "github.com/bitwormhole/wpm/server/support/implservice"
	init0xc984bc "github.com/bitwormhole/wpm/server/support/init"
	intents0x8ee0e0 "github.com/bitwormhole/wpm/server/support/intents"
	intenttemplates0x2e3dcf "github.com/bitwormhole/wpm/server/support/intenttemplates"
	locations0xb36349 "github.com/bitwormhole/wpm/server/support/locations"
	mediae0xf005e2 "github.com/bitwormhole/wpm/server/support/mediae"
	namespaces0xceefcf "github.com/bitwormhole/wpm/server/support/namespaces"
	platforms0xb539c0 "github.com/bitwormhole/wpm/server/support/platforms"
	plugins0x82e34b "github.com/bitwormhole/wpm/server/support/plugins"
	presets0x875f8b "github.com/bitwormhole/wpm/server/support/presets"
	projects0x4d85c7 "github.com/bitwormhole/wpm/server/support/projects"
	repositories0x637d5e "github.com/bitwormhole/wpm/server/support/repositories"
	repositoryworktreeproject0x399028 "github.com/bitwormhole/wpm/server/support/repositoryworktreeproject"
	settings0x19237d "github.com/bitwormhole/wpm/server/support/settings"
	setup0xd9ff02 "github.com/bitwormhole/wpm/server/support/setup"
	worktrees0xa762f3 "github.com/bitwormhole/wpm/server/support/worktrees"
	filequery0xca51d2 "github.com/bitwormhole/wpm/server/utils/filequery"
	handlers0x162741 "github.com/bitwormhole/wpm/server/utils/filequery/handlers"
	controller0x9dc399 "github.com/bitwormhole/wpm/server/web/controller"
	filter0x8aa8f6 "github.com/bitwormhole/wpm/server/web/filter"
)

type pComCLIMakerFilter struct {
	instance *v3filters0xa6552a.CLIMakerFilter
	 markup0x23084a.Component `class:"wpm-intent-filter"`
}


type pComCLIRunnerFilter struct {
	instance *v3filters0xa6552a.CLIRunnerFilter
	 markup0x23084a.Component `class:"wpm-intent-filter"`
	IntentHandlerService service0x3e063d.IntentHandlerService `inject:"#IntentHandlerService"`
}


type pComExampleFilter struct {
	instance *v3filters0xa6552a.ExampleFilter
	 markup0x23084a.Component `class:"wpm-intent-filter"`
}


type pComPrepareActionFilter struct {
	instance *v3filters0xa6552a.PrepareActionFilter
	 markup0x23084a.Component `class:"wpm-intent-filter"`
}


type pComPreparePropertiesFilter struct {
	instance *v3filters0xa6552a.PreparePropertiesFilter
	 markup0x23084a.Component `class:"wpm-intent-filter"`
}


type pComCheckTemplateFilter struct {
	instance *v3filters0xa6552a.CheckTemplateFilter
	 markup0x23084a.Component `class:"wpm-intent-filter"`
}


type pComFindTemplateFilter struct {
	instance *v3filters0xa6552a.FindTemplateFilter
	 markup0x23084a.Component `class:"wpm-intent-filter"`
	IntentTemplateService service0x3e063d.IntentTemplateService `inject:"#IntentTemplateService"`
}


type pComGormDBAgentImpl struct {
	instance *dbagent0x9f90fb.GormDBAgentImpl
	 markup0x23084a.Component `id:"GormDBAgent" class:"life"`
	Sources datasource0x68a737.SourceManager `inject:"#starter-gorm-source-manager"`
}


type pComImpAppRuntimeService struct {
	instance *appruntime0x8dfe0a.ImpAppRuntimeService
	 markup0x23084a.Component `id:"AppRuntimeService" class:"life"`
	FileSystemService service0x3e063d.FileSystemService `inject:"#FileSystemService"`
	AppDataService service0x3e063d.AppDataService `inject:"#AppDataService"`
	MediaService service0x3e063d.MediaService `inject:"#MediaService"`
}


type pComWpmBackupController struct {
	instance *backups0xe44d54.WpmBackupController
	 markup0x23084a.Component `id:"" class:"rest-controller"`
	BackupService service0x3e063d.DatabaseBackupService `inject:"#DatabaseBackupService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComImpBackupServiceDAO struct {
	instance *backups0xe44d54.ImpBackupServiceDAO
	 markup0x23084a.Component `id:"wpm-database-backup-dao" class:""`
	Agent dbagent0x9f90fb.GormDBAgent `inject:"#GormDBAgent"`
}


type pComImpBackupService struct {
	instance *backups0xe44d54.ImpBackupService
	 markup0x23084a.Component `id:"DatabaseBackupService" class:"life"`
	AppDataService service0x3e063d.AppDataService `inject:"#AppDataService"`
	FilesysService service0x3e063d.FileSystemService `inject:"#FileSystemService"`
	BackupDao dao0x5af8d0.Backup `inject:"#wpm-database-backup-dao"`
	DoDump bool `inject:"${wpm.db.dump.enabled}"`
}


type pComImpCacheManager struct {
	instance *caches0xd7996d.ImpCacheManager
	 markup0x23084a.Component `id:"CacheService" class:"life"`
	ProviderRegistryList []caches0x9d186b.ProviderRegistry `inject:".wpm-cache-provider"`
}


type pComTheCheckUpdateServiceImpl struct {
	instance *checkupdate0xea1855.TheCheckUpdateServiceImpl
	 markup0x23084a.Component `id:"CheckUpdateService"`
	PackagesURL string `inject:"${wpm.check-update.url}"`
	AboutService service0x3e063d.AboutService `inject:"#AboutService"`
	SettingService service0x3e063d.SettingService `inject:"#SettingService"`
}


type pComContentTypeController struct {
	instance *contenttypes0x61ca37.ContentTypeController
	 markup0x23084a.RestController `class:"rest-controller"`
	ContentTypeService service0x3e063d.ContentTypeService `inject:"#ContentTypeService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComProjectTypeDaoImpl struct {
	instance *contenttypes0x61ca37.ProjectTypeDaoImpl
	 markup0x23084a.Component `id:"ContentTypeDAO"`
	Agent dbagent0x9f90fb.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComProjectTypeImportController struct {
	instance *contenttypes0x61ca37.ProjectTypeImportController
	 markup0x23084a.RestController `class:"rest-controller"`
	ProjectTypeImportService service0x3e063d.ProjectTypeImportService `inject:"#ProjectTypeImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComProjectTypeImportServiceImpl struct {
	instance *contenttypes0x61ca37.ProjectTypeImportServiceImpl
	 markup0x23084a.Component `id:"ProjectTypeImportService"`
	AC application0x67f6c5.Context `inject:"context"`
	ProjectTypeService service0x3e063d.ContentTypeService `inject:"#ContentTypeService"`
	PresetService service0x3e063d.PresetService `inject:"#PresetService"`
}


type pComProjectTypeServiceImpl struct {
	instance *contenttypes0x61ca37.ProjectTypeServiceImpl
	 markup0x23084a.Component `id:"ContentTypeService"`
	ProjectTypeDAO dao0x5af8d0.ContentTypeDAO `inject:"#ContentTypeDAO"`
	FileSystemService service0x3e063d.FileSystemService `inject:"#FileSystemService"`
}


type pComExecutableController struct {
	instance *executables0xd3773a.ExecutableController
	 markup0x23084a.RestController `class:"rest-controller"`
	ExecutableService service0x3e063d.ExecutableService `inject:"#ExecutableService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComExecutableDaoImpl struct {
	instance *executables0xd3773a.ExecutableDaoImpl
	 markup0x23084a.Component `id:"ExecutableDAO"`
	Agent dbagent0x9f90fb.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComExecutableImportController struct {
	instance *executables0xd3773a.ExecutableImportController
	 markup0x23084a.RestController `class:"rest-controller"`
	ExecutableImportService service0x3e063d.ExecutableImportService `inject:"#ExecutableImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComExecutableImportServiceImpl struct {
	instance *executables0xd3773a.ExecutableImportServiceImpl
	 markup0x23084a.Component `id:"ExecutableImportService"`
	AC application0x67f6c5.Context `inject:"context"`
	ExecutableService service0x3e063d.ExecutableService `inject:"#ExecutableService"`
	FileSystemService service0x3e063d.FileSystemService `inject:"#FileSystemService"`
}


type pComExecutableServiceImpl struct {
	instance *executables0xd3773a.ExecutableServiceImpl
	 markup0x23084a.Component `id:"ExecutableService"`
	ExecutableDAO dao0x5af8d0.ExecutableDAO `inject:"#ExecutableDAO"`
	IconService service0x3e063d.AppIconService `inject:"#AppIconService"`
	FileSystemService service0x3e063d.FileSystemService `inject:"#FileSystemService"`
	LocationService service0x3e063d.LocationService `inject:"#LocationService"`
}


type pComImpHTTPClientService struct {
	instance *httpclient0xf20fe2.ImpHTTPClientService
	 markup0x23084a.Component `id:"HTTPClientService"`
	MaxContentLength int `inject:"${wpm.httpclient.max-content-length}"`
}


type pComExampleDaoImpl struct {
	instance *impldao0x73998b.ExampleDaoImpl
	 markup0x23084a.Component `id:"ExampleDAO"`
	Agent dbagent0x9f90fb.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComAboutServiceImpl struct {
	instance *implservice0x22327c.AboutServiceImpl
	 markup0x23084a.Component `id:"AboutService"`
	Profile string `inject:"${application.profiles.active}"`
	Name string `inject:"${application.about.name}"`
	Title string `inject:"${application.about.title}"`
	Copyright string `inject:"${application.about.copyright}"`
	ServerPort int `inject:"${server.port}"`
	EnableDebug bool `inject:"${wpm.debug.enabled}"`
	PlatformService service0x3e063d.PlatformService `inject:"#PlatformService"`
	ProfileService service0x3e063d.ProfileService `inject:"#ProfileService"`
	AppRuntimeService service0x3e063d.AppRuntimeService `inject:"#AppRuntimeService"`
}


type pComAppDataServiceImpl struct {
	instance *implservice0x22327c.AppDataServiceImpl
	 markup0x23084a.Component `id:"AppDataService"`
	ProfileService service0x3e063d.ProfileService `inject:"#ProfileService"`
	AppRuntimeService service0x3e063d.AppRuntimeService `inject:"#AppRuntimeService"`
	SQLiteDatabaseNameWithAppVersion bool `inject:"${datasource.wpm.database-name-with-version}"`
}


type pComAppIconServiceImpl struct {
	instance *implservice0x22327c.AppIconServiceImpl
	 markup0x23084a.Component `id:"AppIconService" class:"life"`
	PropsName string `inject:"${wpm.exe-icons.properties}"`
	Context application0x67f6c5.Context `inject:"context"`
}


type pComExampleServiceImpl struct {
	instance *implservice0x22327c.ExampleServiceImpl
	 markup0x23084a.Component `id:"ExampleService"`
}


type pComFileQueryServiceImpl struct {
	instance *implservice0x22327c.FileQueryServiceImpl
	 markup0x23084a.Component `id:"FileQueryService"`
	HandlerRegistryList []filequery0xca51d2.HandlerRegistry `inject:".filequery-handler-registry"`
}


type pComFileSystemServiceImpl struct {
	instance *implservice0x22327c.FileSystemServiceImpl
	 markup0x23084a.Component `id:"FileSystemService"`
}


type pComUUIDGenServiceImpl struct {
	instance *implservice0x22327c.UUIDGenServiceImpl
	 markup0x23084a.Component `id:"UUIDGenService" initMethod:"Init"`
}


type pComWpmInitController struct {
	instance *init0xc984bc.WpmInitController
	 markup0x23084a.RestController `class:"rest-controller"`
	InitService service0x3e063d.InitService `inject:"#InitService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComImpInitService struct {
	instance *init0xc984bc.ImpInitService
	 markup0x23084a.Component `id:"InitService" class:"life"`
	AboutService service0x3e063d.AboutService `inject:"#AboutService"`
	ProjectTypeService service0x3e063d.ContentTypeService `inject:"#ContentTypeService"`
	ExecutableService service0x3e063d.ExecutableService `inject:"#ExecutableService"`
	CheckUpdateService service0x3e063d.CheckUpdateService `inject:"#CheckUpdateService"`
	SetupService service0x3e063d.SetupService `inject:"#SetupService"`
}


type pComRunIntentController struct {
	instance *intents0x8ee0e0.RunIntentController
	 markup0x23084a.RestController `class:"rest-controller"`
	IntentService service0x3e063d.IntentService `inject:"#IntentService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComFilterManagerImpl struct {
	instance *intents0x8ee0e0.FilterManagerImpl
	 markup0x23084a.Component `id:"wpm-intent-filter-manager"`
	FilterRegistryList []intents0x8557f3.FilterRegistry `inject:".wpm-intent-filter"`
}


type pComIntentHandlerImpl struct {
	instance *intents0x8ee0e0.IntentHandlerImpl
	 markup0x23084a.Component `id:"IntentHandlerService"`
}


type pComRunIntentServiceImpl struct {
	instance *intents0x8ee0e0.RunIntentServiceImpl
	 markup0x23084a.Component `id:"IntentService"`
	GitLibAgent store0x8467b3.LibAgent `inject:"#git-lib-agent"`
	IntentFilterManager intents0x8557f3.FilterManager `inject:"#wpm-intent-filter-manager"`
	LocalRepositoryService service0x3e063d.LocalRepositoryService `inject:"#LocalRepositoryService"`
	ExecutableService service0x3e063d.ExecutableService `inject:"#ExecutableService"`
	IntentHandlerService service0x3e063d.IntentHandlerService `inject:"#IntentHandlerService"`
}


type pComIntentTemplateController struct {
	instance *intenttemplates0x2e3dcf.IntentTemplateController
	 markup0x23084a.RestController `class:"rest-controller"`
	IntentTemplateService service0x3e063d.IntentTemplateService `inject:"#IntentTemplateService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComIntentTemplateDaoImpl struct {
	instance *intenttemplates0x2e3dcf.IntentTemplateDaoImpl
	 markup0x23084a.Component `id:"IntentTemplateDAO"`
	Agent dbagent0x9f90fb.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComIntentTemplateServiceImpl struct {
	instance *intenttemplates0x2e3dcf.IntentTemplateServiceImpl
	 markup0x23084a.Component `id:"IntentTemplateService"`
	AC application0x67f6c5.Context `inject:"context"`
	IntentTempDAO dao0x5af8d0.IntentTemplateDAO `inject:"#IntentTemplateDAO"`
	IntentFilterManager intents0x8557f3.FilterManager `inject:"#wpm-intent-filter-manager"`
}


type pComImpLocationDao struct {
	instance *locations0xb36349.ImpLocationDao
	 markup0x23084a.Component `id:"LocationDAO"`
	Agent dbagent0x9f90fb.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComImpLocationService struct {
	instance *locations0xb36349.ImpLocationService
	 markup0x23084a.Component `id:"LocationService"`
	DAO dao0x5af8d0.LocationDAO `inject:"#LocationDAO"`
}


type pComMediaController struct {
	instance *mediae0xf005e2.MediaController
	 markup0x23084a.RestController `class:"rest-controller"`
	MediaService service0x3e063d.MediaService `inject:"#MediaService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComMediaDaoImpl struct {
	instance *mediae0xf005e2.MediaDaoImpl
	 markup0x23084a.Component `id:"MediaDAO"`
	Agent dbagent0x9f90fb.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComMediaServiceImpl struct {
	instance *mediae0xf005e2.MediaServiceImpl
	 markup0x23084a.Component `id:"MediaService"`
	AC application0x67f6c5.Context `inject:"context"`
	MediaDAO dao0x5af8d0.MediaDAO `inject:"#MediaDAO"`
	SysMainRepoService service0x3e063d.MainRepositoryService `inject:"#MainRepositoryService"`
	FileSystemService service0x3e063d.FileSystemService `inject:"#FileSystemService"`
	ContentTypeService service0x3e063d.ContentTypeService `inject:"#ContentTypeService"`
	ResPathPrefix string `inject:"${wpm.presets.res-path-prefix}"`
	WebPathPrefix string `inject:"${wpm.presets.web-path-prefix}"`
}


type pComNamespaceController struct {
	instance *namespaces0xceefcf.NamespaceController
	 markup0x23084a.RestController `class:"rest-controller"`
	NamespaceService service0x3e063d.NamespaceService `inject:"#NamespaceService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComImpNamespaceDao struct {
	instance *namespaces0xceefcf.ImpNamespaceDao
	 markup0x23084a.Component `id:"NamespaceDAO"`
	AC application0x67f6c5.Context `inject:"context"`
}


type pComImpNamespaceService struct {
	instance *namespaces0xceefcf.ImpNamespaceService
	 markup0x23084a.Component `id:"NamespaceService"`
	MyDAO dao0x5af8d0.NamespaceDAO `inject:"#NamespaceDAO"`
}


type pComLinuxPlatformServiceImpl struct {
	instance *platforms0xb539c0.LinuxPlatformServiceImpl
	 markup0x23084a.Component `class:"PlatformProviderRegistry"`
}


type pComPlatformServiceImpl struct {
	instance *platforms0xb539c0.PlatformServiceImpl
	 markup0x23084a.Component `id:"PlatformService"`
	Providers []service0x3e063d.PlatformProviderRegistry `inject:".PlatformProviderRegistry"`
}


type pComProfileServiceImpl struct {
	instance *platforms0xb539c0.ProfileServiceImpl
	 markup0x23084a.Component `id:"ProfileService"`
	PlatformService service0x3e063d.PlatformService `inject:"#PlatformService"`
}


type pComWindowsPlatformServiceImpl struct {
	instance *platforms0xb539c0.WindowsPlatformServiceImpl
	 markup0x23084a.Component `class:"PlatformProviderRegistry"`
}


type pComSoftwarePackageController struct {
	instance *plugins0x82e34b.SoftwarePackageController
	 markup0x23084a.RestController `class:"rest-controller"`
	SoftwarePackageService service0x3e063d.SoftwarePackageService `inject:"#SoftwarePackageService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComPluginDaoImpl struct {
	instance *plugins0x82e34b.PluginDaoImpl
	 markup0x23084a.Component `id:"SoftwarePackageDAO"`
	Agent dbagent0x9f90fb.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComPluginServiceImpl struct {
	instance *plugins0x82e34b.PluginServiceImpl
	 markup0x23084a.Component `id:"SoftwarePackageService"`
	SoftwarePackageDAO dao0x5af8d0.SoftwarePackageDAO `inject:"#SoftwarePackageDAO"`
	NamespaceService service0x3e063d.NamespaceService `inject:"#NamespaceService"`
	HTTPClientService service0x3e063d.HTTPClientService `inject:"#HTTPClientService"`
}


type pComSoftwareSetController struct {
	instance *plugins0x82e34b.SoftwareSetController
	 markup0x23084a.RestController `class:"rest-controller"`
	SoftwareSetService service0x3e063d.SoftwareSetService `inject:"#SoftwareSetService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComImpSoftwareSetService struct {
	instance *plugins0x82e34b.ImpSoftwareSetService
	 markup0x23084a.Component `id:"SoftwareSetService"`
	SoftwarePackageService service0x3e063d.SoftwarePackageService `inject:"#SoftwarePackageService"`
}


type pComCacheProvider struct {
	instance *presets0x875f8b.CacheProvider
	 markup0x23084a.Component ` id:"PresetCache"  class:"wpm-cache-provider"`
	AC application0x67f6c5.Context `inject:"context"`
	CS service0x3e063d.CacheService `inject:"#CacheService"`
	ListFileName string `inject:"${wpm.presets.list-file-name}"`
}


type pComImpPresetService struct {
	instance *presets0x875f8b.ImpPresetService
	 markup0x23084a.Component `id:"PresetService"`
	Cache presets0x875f8b.Cache `inject:"#PresetCache"`
}


type pComProjectController struct {
	instance *projects0x4d85c7.ProjectController
	 markup0x23084a.RestController `class:"rest-controller"`
	ProjectService service0x3e063d.ProjectService `inject:"#ProjectService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComProjectDaoImpl struct {
	instance *projects0x4d85c7.ProjectDaoImpl
	 markup0x23084a.Component `id:"ProjectDAO"`
	Agent dbagent0x9f90fb.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComProjectImportController struct {
	instance *projects0x4d85c7.ProjectImportController
	 markup0x23084a.RestController `class:"rest-controller"`
	ProjectImportService service0x3e063d.ProjectImportService `inject:"#ProjectImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComProjectImportServiceImpl struct {
	instance *projects0x4d85c7.ProjectImportServiceImpl
	 markup0x23084a.Component `id:"ProjectImportService"`
}


type pComProjectServiceImpl struct {
	instance *projects0x4d85c7.ProjectServiceImpl
	 markup0x23084a.Component `id:"ProjectService"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
	LocalRepoService service0x3e063d.LocalRepositoryService `inject:"#LocalRepositoryService"`
	ProjectTypeService service0x3e063d.ContentTypeService `inject:"#ContentTypeService"`
	FileSystemService service0x3e063d.FileSystemService `inject:"#FileSystemService"`
	LocationService service0x3e063d.LocationService `inject:"#LocationService"`
	ProjectDAO dao0x5af8d0.ProjectDAO `inject:"#ProjectDAO"`
	LocalRepoDAO dao0x5af8d0.LocalRepositoryDAO `inject:"#LocalRepositoryDAO"`
	GitLibAgent store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComLocalRepositoryController struct {
	instance *repositories0x637d5e.LocalRepositoryController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoService service0x3e063d.LocalRepositoryService `inject:"#LocalRepositoryService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComRepositoryDaoImpl struct {
	instance *repositories0x637d5e.RepositoryDaoImpl
	 markup0x23084a.Component `id:"LocalRepositoryDAO"`
	Agent dbagent0x9f90fb.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComLocalRepositoryFinderImpl struct {
	instance *repositories0x637d5e.LocalRepositoryFinderImpl
	 markup0x23084a.Component `id:"LocalRepositoryFinder"`
	GitLibAgent store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComLocalRepositoryServiceImpl struct {
	instance *repositories0x637d5e.LocalRepositoryServiceImpl
	 markup0x23084a.Component `id:"LocalRepositoryService"`
	GitLibAgent store0x8467b3.LibAgent `inject:"#git-lib-agent"`
	LocalRepositoryDAO dao0x5af8d0.LocalRepositoryDAO `inject:"#LocalRepositoryDAO"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
	RepoFinder service0x3e063d.LocalRepositoryFinder `inject:"#LocalRepositoryFinder"`
	LrStateLoader service0x3e063d.LocalRepositoryStateLoader `inject:"#LocalRepositoryStateLoader"`
	FileSystemService service0x3e063d.FileSystemService `inject:"#FileSystemService"`
	LocationService service0x3e063d.LocationService `inject:"#LocationService"`
	ProjectService service0x3e063d.ProjectService `inject:"#ProjectService"`
}


type pComLocalRepositoryStateLoaderImpl struct {
	instance *repositories0x637d5e.LocalRepositoryStateLoaderImpl
	 markup0x23084a.Component `id:"LocalRepositoryStateLoader"`
	LocalRepoService service0x3e063d.LocalRepositoryService `inject:"#LocalRepositoryService"`
	Dao dao0x5af8d0.LocalRepositoryDAO `inject:"#LocalRepositoryDAO"`
	GitLibAgent store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComMainRepositoryServiceImpl struct {
	instance *repositories0x637d5e.MainRepositoryServiceImpl
	 markup0x23084a.Component `id:"MainRepositoryService"`
	AppDataService service0x3e063d.AppDataService `inject:"#AppDataService"`
	GitLA store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComRemoteRepositoryController struct {
	instance *repositories0x637d5e.RemoteRepositoryController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoService service0x3e063d.RemoteRepositoryService `inject:"#RemoteRepositoryService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComRemoteRepositoryServiceImpl struct {
	instance *repositories0x637d5e.RemoteRepositoryServiceImpl
	 markup0x23084a.Component `id:"RemoteRepositoryService"`
}


type pComRepositoryImportController struct {
	instance *repositories0x637d5e.RepositoryImportController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoStateLoader service0x3e063d.LocalRepositoryStateLoader `inject:"#LocalRepositoryStateLoader"`
	ImportService service0x3e063d.RepositoryImportService `inject:"#RepositoryImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComRepositoryImportServiceImpl struct {
	instance *repositories0x637d5e.RepositoryImportServiceImpl
	 markup0x23084a.Component `id:"RepositoryImportService"`
	RepositoryService service0x3e063d.LocalRepositoryService `inject:"#LocalRepositoryService"`
	RepoFinder service0x3e063d.LocalRepositoryFinder `inject:"#LocalRepositoryFinder"`
}


type pComRepoWorktreeProjectController struct {
	instance *repositoryworktreeproject0x399028.RepoWorktreeProjectController
	 markup0x23084a.RestController `class:"rest-controller"`
	RWPService service0x3e063d.RepositoryWorktreeProjectService `inject:"#RepositoryWorktreeProjectService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComRWPServiceImpl struct {
	instance *repositoryworktreeproject0x399028.RWPServiceImpl
	 markup0x23084a.Component `id:"RepositoryWorktreeProjectService"`
	Repositories service0x3e063d.LocalRepositoryService `inject:"#LocalRepositoryService"`
	Worktrees service0x3e063d.WorktreeService `inject:"#WorktreeService"`
	Projects service0x3e063d.ProjectService `inject:"#ProjectService"`
	GitLib store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComSettingController struct {
	instance *settings0x19237d.SettingController
	 markup0x23084a.RestController `class:"rest-controller"`
	SettingService service0x3e063d.SettingService `inject:"#SettingService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComSettingDaoImpl struct {
	instance *settings0x19237d.SettingDaoImpl
	 markup0x23084a.Component `id:"SettingDAO"`
	Agent dbagent0x9f90fb.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComSettingServiceImpl struct {
	instance *settings0x19237d.SettingServiceImpl
	 markup0x23084a.Component `id:"SettingService" class:"life" `
	SettingDAO dao0x5af8d0.SettingDAO `inject:"#SettingDAO"`
}


type pComWpmSetupController struct {
	instance *setup0xd9ff02.WpmSetupController
	 markup0x23084a.Component `class:"rest-controller"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
	SetupService service0x3e063d.SetupService `inject:"#SetupService"`
}


type pComImpSetupService struct {
	instance *setup0xd9ff02.ImpSetupService
	 markup0x23084a.Component `id:"SetupService"`
	AppDataService service0x3e063d.AppDataService `inject:"#AppDataService"`
	FileSystemService service0x3e063d.FileSystemService `inject:"#FileSystemService"`
	ExecutableService service0x3e063d.ExecutableService `inject:"#ExecutableService"`
	ExecutableImportService service0x3e063d.ExecutableImportService `inject:"#ExecutableImportService"`
	IntentTemplateService service0x3e063d.IntentTemplateService `inject:"#IntentTemplateService"`
	MediaService service0x3e063d.MediaService `inject:"#MediaService"`
	ProjectTypeImportService service0x3e063d.ProjectTypeImportService `inject:"#ProjectTypeImportService"`
	SettingService service0x3e063d.SettingService `inject:"#SettingService"`
}


type pComImpWorktreeDao struct {
	instance *worktrees0xa762f3.ImpWorktreeDao
	 markup0x23084a.Component `id:"WorktreeDAO"`
	Agent dbagent0x9f90fb.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComImpWorktreeService struct {
	instance *worktrees0xa762f3.ImpWorktreeService
	 markup0x23084a.Component `id:"WorktreeService"`
	FileSystemService service0x3e063d.FileSystemService `inject:"#FileSystemService"`
	LocationService service0x3e063d.LocationService `inject:"#LocationService"`
	RepoFinder service0x3e063d.LocalRepositoryFinder `inject:"#LocalRepositoryFinder"`
	DAO dao0x5af8d0.WorktreeDAO `inject:"#WorktreeDAO"`
}


type pComWpmDataSource struct {
	instance *support0xf47d7f.WpmDataSource
	 markup0x23084a.Component `class:"starter-gorm-source-registry"`
	DM datasource0x68a737.DriverManager `inject:"#starter-gorm-driver-manager"`
	AppDataService service0x3e063d.AppDataService `inject:"#AppDataService"`
	AboutService service0x3e063d.AboutService `inject:"#AboutService"`
	Driver string `inject:"${datasource.wpm.driver}"`
	Host string `inject:"${datasource.wpm.host}"`
	Port int `inject:"${datasource.wpm.port}"`
	UserName string `inject:"${datasource.wpm.username}"`
	Password string `inject:"${datasource.wpm.password}"`
	Database string `inject:"${datasource.wpm.database}"`
}


type pComFileSystemHandler struct {
	instance *handlers0x162741.FileSystemHandler
	 markup0x23084a.Component `class:"filequery-handler-registry"`
}


type pComAboutController struct {
	instance *controller0x9dc399.AboutController
	 markup0x23084a.RestController `class:"rest-controller"`
	AboutService service0x3e063d.AboutService `inject:"#AboutService"`
	UpdateService service0x3e063d.CheckUpdateService `inject:"#CheckUpdateService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
	Profile string `inject:"${application.profiles.active}"`
}


type pComAuthController struct {
	instance *controller0x9dc399.AuthController
	 markup0x23084a.RestController `class:"rest-controller"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComExampleController struct {
	instance *controller0x9dc399.ExampleController
	 markup0x23084a.RestController `class:"rest-controller"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComFileQueryController struct {
	instance *controller0x9dc399.FileQueryController
	 markup0x23084a.RestController `class:"rest-controller"`
	FileQueryService service0x3e063d.FileQueryService `inject:"#FileQueryService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComPlatformController struct {
	instance *controller0x9dc399.PlatformController
	 markup0x23084a.RestController `class:"rest-controller"`
	PlatformService service0x3e063d.PlatformService `inject:"#PlatformService"`
	ProfileService service0x3e063d.ProfileService `inject:"#ProfileService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComUploadController struct {
	instance *controller0x9dc399.UploadController
	 markup0x23084a.RestController `class:"rest-controller"`
	FileSystemService service0x3e063d.FileSystemService `inject:"#FileSystemService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComHostFilter struct {
	instance *filter0x8aa8f6.HostFilter
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComHTTP404Filter struct {
	instance *filter0x8aa8f6.HTTP404Filter
	 markup0x23084a.RestController `class:"rest-controller"`
	Context application0x67f6c5.Context `inject:"context"`
}

