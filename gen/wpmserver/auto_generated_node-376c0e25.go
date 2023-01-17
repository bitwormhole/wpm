// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmserver

import (
	store0x8467b3 "github.com/bitwormhole/gitlib/git/store"
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	markup0x23084a "github.com/bitwormhole/starter/markup"
	dao0x5af8d0 "github.com/bitwormhole/wpm/server/data/dao"
	service0x3e063d "github.com/bitwormhole/wpm/server/service"
	impldao0x73998b "github.com/bitwormhole/wpm/server/support/impldao"
	implservice0x22327c "github.com/bitwormhole/wpm/server/support/implservice"
	controller0x9dc399 "github.com/bitwormhole/wpm/server/web/controller"
)

type pComGormDBAgentImpl struct {
	instance *impldao0x73998b.GormDBAgentImpl
	 markup0x23084a.Component `id:"GormDBAgent" class:"life"`
	Sources datasource0x68a737.SourceManager `inject:"#starter-gorm-source-manager"`
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


type pComIntentTemplateDaoImpl struct {
	instance *impldao0x73998b.IntentTemplateDaoImpl
	 markup0x23084a.Component `id:"IntentTemplateDAO"`
	Agent impldao0x73998b.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComRepositoryDaoImpl struct {
	instance *impldao0x73998b.RepositoryDaoImpl
	 markup0x23084a.Component `id:"LocalRepositoryDAO"`
	Agent impldao0x73998b.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComMediaDaoImpl struct {
	instance *impldao0x73998b.MediaDaoImpl
	 markup0x23084a.Component `id:"MediaDAO"`
	Agent impldao0x73998b.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComProjectDaoImpl struct {
	instance *impldao0x73998b.ProjectDaoImpl
	 markup0x23084a.Component `id:"ProjectDAO"`
	Agent impldao0x73998b.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
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


type pComExecutableServiceImpl struct {
	instance *implservice0x22327c.ExecutableServiceImpl
	 markup0x23084a.Component `id:"ExecutableService"`
	ExecutableDAO dao0x5af8d0.ExecutableDAO `inject:"#ExecutableDAO"`
}


type pComIntentTemplateServiceImpl struct {
	instance *implservice0x22327c.IntentTemplateServiceImpl
	 markup0x23084a.Component `id:"IntentTemplateService"`
	IntentTempDAO dao0x5af8d0.IntentTemplateDAO `inject:"#IntentTemplateDAO"`
}


type pComIntentHandlerImpl struct {
	instance *implservice0x22327c.IntentHandlerImpl
	 markup0x23084a.Component `id:"IntentHandlerService"`
}


type pComLocalRepositoryFinderImpl struct {
	instance *implservice0x22327c.LocalRepositoryFinderImpl
	 markup0x23084a.Component `id:"LocalRepositoryFinder"`
	GitLibAgent store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComRepositoryServiceImpl struct {
	instance *implservice0x22327c.RepositoryServiceImpl
	 markup0x23084a.Component `id:"RepositoryService"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
	RepoFinder service0x3e063d.LocalRepositoryFinder `inject:"#LocalRepositoryFinder"`
	LocalRepositoryDAO dao0x5af8d0.LocalRepositoryDAO `inject:"#LocalRepositoryDAO"`
}


type pComLocalRepositoryStateLoaderImpl struct {
	instance *implservice0x22327c.LocalRepositoryStateLoaderImpl
	 markup0x23084a.Component `id:"LocalRepositoryStateLoader"`
	Dao dao0x5af8d0.LocalRepositoryDAO `inject:"#LocalRepositoryDAO"`
	GitLibAgent store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComMediaServiceImpl struct {
	instance *implservice0x22327c.MediaServiceImpl
	 markup0x23084a.Component `id:"MediaService"`
	MediaDAO dao0x5af8d0.MediaDAO `inject:"#MediaDAO"`
	SysMainRepoService service0x3e063d.SystemMainRepositoryService `inject:"#SystemMainRepositoryService"`
}


type pComProjectImportServiceImpl struct {
	instance *implservice0x22327c.ProjectImportServiceImpl
	 markup0x23084a.Component `id:"ProjectImportService"`
}


type pComProjectServiceImpl struct {
	instance *implservice0x22327c.ProjectServiceImpl
	 markup0x23084a.Component `id:"ProjectService"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
	ProjectDAO dao0x5af8d0.ProjectDAO `inject:"#ProjectDAO"`
}


type pComRemoteRepositoryServiceImpl struct {
	instance *implservice0x22327c.RemoteRepositoryServiceImpl
	 markup0x23084a.Component `id:"RemoteRepositoryService"`
}


type pComRepositoryImportServiceImpl struct {
	instance *implservice0x22327c.RepositoryImportServiceImpl
	 markup0x23084a.Component `id:"RepositoryImportService"`
	RepositoryService service0x3e063d.LocalRepositoryService `inject:"#RepositoryService"`
	RepoFinder service0x3e063d.LocalRepositoryFinder `inject:"#LocalRepositoryFinder"`
}


type pComRunIntentServiceImpl struct {
	instance *implservice0x22327c.RunIntentServiceImpl
	 markup0x23084a.Component `id:"IntentService"`
	ExecutableService service0x3e063d.ExecutableService `inject:"#ExecutableService"`
	IntentHandlerService service0x3e063d.IntentHandlerService `inject:"#IntentHandlerService"`
}


type pComSystemMainRepositoryServiceImpl struct {
	instance *implservice0x22327c.SystemMainRepositoryServiceImpl
	 markup0x23084a.Component `id:"SystemMainRepositoryService"`
	SystemMainRepoPath string `inject:"${wpm.system-main-repository.path}"`
	GitLA store0x8467b3.LibAgent `inject:"#git-lib-agent"`
}


type pComUserMainRepositoryServiceImpl struct {
	instance *implservice0x22327c.UserMainRepositoryServiceImpl
	 markup0x23084a.Component `id:"UserMainRepositoryService"`
}


type pComUUIDGenServiceImpl struct {
	instance *implservice0x22327c.UUIDGenServiceImpl
	 markup0x23084a.Component `id:"UUIDGenService" initMethod:"Init"`
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


type pComExecutableController struct {
	instance *controller0x9dc399.ExecutableController
	 markup0x23084a.RestController `class:"rest-controller"`
	ExecutableService service0x3e063d.ExecutableService `inject:"#ExecutableService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComExecutableImportController struct {
	instance *controller0x9dc399.ExecutableImportController
	 markup0x23084a.RestController `class:"rest-controller"`
	ExecutableImportService service0x3e063d.ExecutableImportService `inject:"#ExecutableImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComRunIntentController struct {
	instance *controller0x9dc399.RunIntentController
	 markup0x23084a.RestController `class:"rest-controller"`
	IntentService service0x3e063d.IntentService `inject:"#IntentService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComIntentTemplateController struct {
	instance *controller0x9dc399.IntentTemplateController
	 markup0x23084a.RestController `class:"rest-controller"`
	IntentTemplateService service0x3e063d.IntentTemplateService `inject:"#IntentTemplateService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComLocalRepositoryController struct {
	instance *controller0x9dc399.LocalRepositoryController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoService service0x3e063d.LocalRepositoryService `inject:"#RepositoryService"`
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
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComProjectController struct {
	instance *controller0x9dc399.ProjectController
	 markup0x23084a.RestController `class:"rest-controller"`
	ProjectService service0x3e063d.ProjectService `inject:"#ProjectService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComProjectImportController struct {
	instance *controller0x9dc399.ProjectImportController
	 markup0x23084a.RestController `class:"rest-controller"`
	ProjectImportService service0x3e063d.ProjectImportService `inject:"#ProjectImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComRemoteRepositoryController struct {
	instance *controller0x9dc399.RemoteRepositoryController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoService service0x3e063d.RemoteRepositoryService `inject:"#RemoteRepositoryService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComRepositoryImportController struct {
	instance *controller0x9dc399.RepositoryImportController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoStateLoader service0x3e063d.LocalRepositoryStateLoader `inject:"#LocalRepositoryStateLoader"`
	ImportService service0x3e063d.RepositoryImportService `inject:"#RepositoryImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComUserMainRepositoryController struct {
	instance *controller0x9dc399.UserMainRepositoryController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoService service0x3e063d.UserMainRepositoryService `inject:"#UserMainRepositoryService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}

