// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmserver

import (
	store0x8467b3 "github.com/bitwormhole/gitlib/git/store"
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
	dao0x5af8d0 "github.com/bitwormhole/wpm/server/data/dao"
	service0x3e063d "github.com/bitwormhole/wpm/server/service"
	support0xf47d7f "github.com/bitwormhole/wpm/server/support"
	impldao0x73998b "github.com/bitwormhole/wpm/server/support/impldao"
	implservice0x22327c "github.com/bitwormhole/wpm/server/support/implservice"
	filequery0xca51d2 "github.com/bitwormhole/wpm/server/utils/filequery"
	handlers0x162741 "github.com/bitwormhole/wpm/server/utils/filequery/handlers"
	intents0xec84e7 "github.com/bitwormhole/wpm/server/utils/intents"
	filters0x5d53d8 "github.com/bitwormhole/wpm/server/utils/intents/filters"
	controller0x9dc399 "github.com/bitwormhole/wpm/server/web/controller"
	filter0x8aa8f6 "github.com/bitwormhole/wpm/server/web/filter"
)

type pComLocalRepositoryFinderImpl struct {
	instance *implservice0x22327c.LocalRepositoryFinderImpl
	 markup0x23084a.Component `id:"LocalRepositoryFinder"`
	GitLibAgent store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComMainRepositoryServiceImpl struct {
	instance *implservice0x22327c.MainRepositoryServiceImpl
	 markup0x23084a.Component `id:"MainRepositoryService"`
	AppDataService service0x3e063d.AppDataService `inject:"#AppDataService"`
	GitLA store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComAboutServiceImpl struct {
	instance *implservice0x22327c.AboutServiceImpl
	 markup0x23084a.Component `id:"AboutService"`
	Profile string `inject:"${application.profiles.active}"`
	Name string `inject:"${application.about.name}"`
	Title string `inject:"${application.about.title}"`
	Copyright string `inject:"${application.about.copyright}"`
	PlatformService service0x3e063d.PlatformService `inject:"#PlatformService"`
	ProfileService service0x3e063d.ProfileService `inject:"#ProfileService"`
}


type pComMediaServiceImpl struct {
	instance *implservice0x22327c.MediaServiceImpl
	 markup0x23084a.Component `id:"MediaService"`
	MediaDAO dao0x5af8d0.MediaDAO `inject:"#MediaDAO"`
	SysMainRepoService service0x3e063d.MainRepositoryService `inject:"#MainRepositoryService"`
}


type pComFileQueryServiceImpl struct {
	instance *implservice0x22327c.FileQueryServiceImpl
	 markup0x23084a.Component `id:"FileQueryService"`
	HandlerRegistryList []filequery0xca51d2.HandlerRegistry `inject:".filequery-handler-registry"`
}


type pComProjectServiceImpl struct {
	instance *implservice0x22327c.ProjectServiceImpl
	 markup0x23084a.Component `id:"ProjectService"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
	ProjectDAO dao0x5af8d0.ProjectDAO `inject:"#ProjectDAO"`
}


type pComIntentTemplateServiceImpl struct {
	instance *implservice0x22327c.IntentTemplateServiceImpl
	 markup0x23084a.Component `id:"IntentTemplateService"`
	IntentTempDAO dao0x5af8d0.IntentTemplateDAO `inject:"#IntentTemplateDAO"`
}


type pComLocalRepositoryServiceImpl struct {
	instance *implservice0x22327c.LocalRepositoryServiceImpl
	 markup0x23084a.Component `id:"LocalRepositoryService"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
	RepoFinder service0x3e063d.LocalRepositoryFinder `inject:"#LocalRepositoryFinder"`
	LocalRepositoryDAO dao0x5af8d0.LocalRepositoryDAO `inject:"#LocalRepositoryDAO"`
}


type pComRemoteRepositoryServiceImpl struct {
	instance *implservice0x22327c.RemoteRepositoryServiceImpl
	 markup0x23084a.Component `id:"RemoteRepositoryService"`
}


type pComRunIntentServiceImpl struct {
	instance *implservice0x22327c.RunIntentServiceImpl
	 markup0x23084a.Component `id:"IntentService"`
	GitLibAgent store0x8467b3.LibAgent `inject:"#git-lib-agent"`
	IntentFilterManager intents0xec84e7.FilterManager `inject:"#intent-filter-manager"`
	LocalRepositoryService service0x3e063d.LocalRepositoryService `inject:"#LocalRepositoryService"`
	ExecutableService service0x3e063d.ExecutableService `inject:"#ExecutableService"`
	IntentHandlerService service0x3e063d.IntentHandlerService `inject:"#IntentHandlerService"`
}


type pComAppDataServiceImpl struct {
	instance *implservice0x22327c.AppDataServiceImpl
	 markup0x23084a.Component `id:"AppDataService"`
	ProfileService service0x3e063d.ProfileService `inject:"#ProfileService"`
}


type pComUUIDGenServiceImpl struct {
	instance *implservice0x22327c.UUIDGenServiceImpl
	 markup0x23084a.Component `id:"UUIDGenService" initMethod:"Init"`
}


type pComRepositoryImportServiceImpl struct {
	instance *implservice0x22327c.RepositoryImportServiceImpl
	 markup0x23084a.Component `id:"RepositoryImportService"`
	RepositoryService service0x3e063d.LocalRepositoryService `inject:"#LocalRepositoryService"`
	RepoFinder service0x3e063d.LocalRepositoryFinder `inject:"#LocalRepositoryFinder"`
}


type pComExampleServiceImpl struct {
	instance *implservice0x22327c.ExampleServiceImpl
	 markup0x23084a.Component `id:"ExampleService"`
}


type pComExecutableImportServiceImpl struct {
	instance *implservice0x22327c.ExecutableImportServiceImpl
	 markup0x23084a.Component `id:"ExecutableImportService"`
	ExecutableService service0x3e063d.ExecutableService `inject:"#ExecutableService"`
}


type pComProjectImportServiceImpl struct {
	instance *implservice0x22327c.ProjectImportServiceImpl
	 markup0x23084a.Component `id:"ProjectImportService"`
}


type pComLocalRepositoryStateLoaderImpl struct {
	instance *implservice0x22327c.LocalRepositoryStateLoaderImpl
	 markup0x23084a.Component `id:"LocalRepositoryStateLoader"`
	LocalRepoService service0x3e063d.LocalRepositoryService `inject:"#LocalRepositoryService"`
	Dao dao0x5af8d0.LocalRepositoryDAO `inject:"#LocalRepositoryDAO"`
	GitLibAgent store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComExecutableServiceImpl struct {
	instance *implservice0x22327c.ExecutableServiceImpl
	 markup0x23084a.Component `id:"ExecutableService"`
	ExecutableDAO dao0x5af8d0.ExecutableDAO `inject:"#ExecutableDAO"`
	IconService service0x3e063d.AppIconService `inject:"#AppIconService"`
}


type pComAppIconServiceImpl struct {
	instance *implservice0x22327c.AppIconServiceImpl
	 markup0x23084a.Component `id:"AppIconService" class:"life"`
	PropsName string `inject:"${wpm.exe-icons.properties}"`
	Context application0x67f6c5.Context `inject:"context"`
}


type pComIntentHandlerImpl struct {
	instance *implservice0x22327c.IntentHandlerImpl
	 markup0x23084a.Component `id:"IntentHandlerService"`
}


type pComWpmDataSource struct {
	instance *support0xf47d7f.WpmDataSource
	 markup0x23084a.Component `class:"starter-gorm-source-registry"`
	DM datasource0x68a737.DriverManager `inject:"#starter-gorm-driver-manager"`
	AppDataService service0x3e063d.AppDataService `inject:"#AppDataService"`
	Driver string `inject:"${datasource.wpm.driver}"`
	Host string `inject:"${datasource.wpm.host}"`
	Port int `inject:"${datasource.wpm.port}"`
	UserName string `inject:"${datasource.wpm.username}"`
	Password string `inject:"${datasource.wpm.password}"`
	Database string `inject:"${datasource.wpm.database}"`
}


type pComIntentTemplateDaoImpl struct {
	instance *impldao0x73998b.IntentTemplateDaoImpl
	 markup0x23084a.Component `id:"IntentTemplateDAO"`
	Agent impldao0x73998b.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComMediaDaoImpl struct {
	instance *impldao0x73998b.MediaDaoImpl
	 markup0x23084a.Component `id:"MediaDAO"`
	Agent impldao0x73998b.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComExampleDaoImpl struct {
	instance *impldao0x73998b.ExampleDaoImpl
	 markup0x23084a.Component `id:"ExampleDAO"`
	Agent impldao0x73998b.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComExecutableDaoImpl struct {
	instance *impldao0x73998b.ExecutableDaoImpl
	 markup0x23084a.Component `id:"ExecutableDAO"`
	Agent impldao0x73998b.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComProjectDaoImpl struct {
	instance *impldao0x73998b.ProjectDaoImpl
	 markup0x23084a.Component `id:"ProjectDAO"`
	Agent impldao0x73998b.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComRepositoryDaoImpl struct {
	instance *impldao0x73998b.RepositoryDaoImpl
	 markup0x23084a.Component `id:"LocalRepositoryDAO"`
	Agent impldao0x73998b.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComGormDBAgentImpl struct {
	instance *impldao0x73998b.GormDBAgentImpl
	 markup0x23084a.Component `id:"GormDBAgent" class:"life"`
	Sources datasource0x68a737.SourceManager `inject:"#starter-gorm-source-manager"`
}


type pComFileSystemHandler struct {
	instance *handlers0x162741.FileSystemHandler
	 markup0x23084a.Component `class:"filequery-handler-registry"`
}


type pComIntentFilterForNautilus struct {
	instance *filters0x5d53d8.IntentFilterForNautilus
	 markup0x23084a.Component `class:"intent-filter-registry"`
}


type pComIntentFilterForExplorer struct {
	instance *filters0x5d53d8.IntentFilterForExplorer
	 markup0x23084a.Component `class:"intent-filter-registry"`
}


type pComIntentFilterForVscode struct {
	instance *filters0x5d53d8.IntentFilterForVscode
	 markup0x23084a.Component `class:"intent-filter-registry"`
}


type pComIntentFilterForChrome struct {
	instance *filters0x5d53d8.IntentFilterForChrome
	 markup0x23084a.Component `class:"intent-filter-registry"`
}


type pComExecuteIntentFilter struct {
	instance *filters0x5d53d8.ExecuteIntentFilter
	 markup0x23084a.Component `class:"intent-filter-registry"`
	GitLibAgent store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComIntentFilterForCmd struct {
	instance *filters0x5d53d8.IntentFilterForCmd
	 markup0x23084a.Component `class:"intent-filter-registry"`
}


type pComIntentFilterForMsEdge struct {
	instance *filters0x5d53d8.IntentFilterForMsEdge
	 markup0x23084a.Component `class:"intent-filter-registry"`
}


type pComIntentFilterForFirefox struct {
	instance *filters0x5d53d8.IntentFilterForFirefox
	 markup0x23084a.Component `class:"intent-filter-registry"`
}


type pComIntentFilterFor struct {
	instance *filters0x5d53d8.IntentFilterFor
	 markup0x23084a.Component `class:"intent-filter-registry"`
}


type pComIntentFilterForPowerShell struct {
	instance *filters0x5d53d8.IntentFilterForPowerShell
	 markup0x23084a.Component `class:"intent-filter-registry"`
}


type pComFilterManagerImpl struct {
	instance *intents0xec84e7.FilterManagerImpl
	 markup0x23084a.Component `id:"intent-filter-manager"`
	Filters []intents0xec84e7.FilterRegistry `inject:".intent-filter-registry"`
}


type pComHTTP404Filter struct {
	instance *filter0x8aa8f6.HTTP404Filter
	 markup0x23084a.RestController `class:"rest-controller"`
	Context application0x67f6c5.Context `inject:"context"`
}


type pComHostFilter struct {
	instance *filter0x8aa8f6.HostFilter
	 markup0x23084a.RestController `class:"rest-controller"`
}


type pComExampleController struct {
	instance *controller0x9dc399.ExampleController
	 markup0x23084a.RestController `class:"rest-controller"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComExecutableController struct {
	instance *controller0x9dc399.ExecutableController
	 markup0x23084a.RestController `class:"rest-controller"`
	ExecutableService service0x3e063d.ExecutableService `inject:"#ExecutableService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComRunIntentController struct {
	instance *controller0x9dc399.RunIntentController
	 markup0x23084a.RestController `class:"rest-controller"`
	IntentService service0x3e063d.IntentService `inject:"#IntentService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComProjectImportController struct {
	instance *controller0x9dc399.ProjectImportController
	 markup0x23084a.RestController `class:"rest-controller"`
	ProjectImportService service0x3e063d.ProjectImportService `inject:"#ProjectImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComExecutableImportController struct {
	instance *controller0x9dc399.ExecutableImportController
	 markup0x23084a.RestController `class:"rest-controller"`
	ExecutableImportService service0x3e063d.ExecutableImportService `inject:"#ExecutableImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComRepositoryImportController struct {
	instance *controller0x9dc399.RepositoryImportController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoStateLoader service0x3e063d.LocalRepositoryStateLoader `inject:"#LocalRepositoryStateLoader"`
	ImportService service0x3e063d.RepositoryImportService `inject:"#RepositoryImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComAboutController struct {
	instance *controller0x9dc399.AboutController
	 markup0x23084a.RestController `class:"rest-controller"`
	AboutService service0x3e063d.AboutService `inject:"#AboutService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComIntentTemplateController struct {
	instance *controller0x9dc399.IntentTemplateController
	 markup0x23084a.RestController `class:"rest-controller"`
	IntentTemplateService service0x3e063d.IntentTemplateService `inject:"#IntentTemplateService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComAuthController struct {
	instance *controller0x9dc399.AuthController
	 markup0x23084a.RestController `class:"rest-controller"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComMediaController struct {
	instance *controller0x9dc399.MediaController
	 markup0x23084a.RestController `class:"rest-controller"`
	MediaService service0x3e063d.MediaService `inject:"#MediaService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComPlatformController struct {
	instance *controller0x9dc399.PlatformController
	 markup0x23084a.RestController `class:"rest-controller"`
	PlatformService service0x3e063d.PlatformService `inject:"#PlatformService"`
	ProfileService service0x3e063d.ProfileService `inject:"#ProfileService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComRemoteRepositoryController struct {
	instance *controller0x9dc399.RemoteRepositoryController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoService service0x3e063d.RemoteRepositoryService `inject:"#RemoteRepositoryService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComFileQueryController struct {
	instance *controller0x9dc399.FileQueryController
	 markup0x23084a.RestController `class:"rest-controller"`
	FileQueryService service0x3e063d.FileQueryService `inject:"#FileQueryService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComLocalRepositoryController struct {
	instance *controller0x9dc399.LocalRepositoryController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoService service0x3e063d.LocalRepositoryService `inject:"#LocalRepositoryService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComProjectController struct {
	instance *controller0x9dc399.ProjectController
	 markup0x23084a.RestController `class:"rest-controller"`
	ProjectService service0x3e063d.ProjectService `inject:"#ProjectService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}

