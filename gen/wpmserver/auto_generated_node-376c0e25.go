// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmserver

import (
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
	 markup0x23084a.Component `id:"GormDBAgent" initMethod:"Init"`
	Source datasource0x68a737.Source `inject:"#gorm-datasource-default"`
}


type pComExecutableDaoImpl struct {
	instance *impldao0x73998b.ExecutableDaoImpl
	 markup0x23084a.Component `id:"ExecutableDAO"`
	Agent impldao0x73998b.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComIntentDaoImpl struct {
	instance *impldao0x73998b.IntentDaoImpl
	 markup0x23084a.Component `id:"IntentDAO"`
	Agent impldao0x73998b.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
}


type pComRepositoryDaoImpl struct {
	instance *impldao0x73998b.RepositoryDaoImpl
	 markup0x23084a.Component `id:"RepositoryDAO"`
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


type pComIntentServiceImpl struct {
	instance *implservice0x22327c.IntentServiceImpl
	 markup0x23084a.Component `id:"IntentService"`
	IntentDAO dao0x5af8d0.IntentDAO `inject:"#IntentDAO"`
}


type pComRepositoryServiceImpl struct {
	instance *implservice0x22327c.RepositoryServiceImpl
	 markup0x23084a.Component `id:"RepositoryService"`
	UUIDGenService service0x3e063d.UUIDGenService `inject:"#UUIDGenService"`
	LocateService service0x3e063d.RepositoryLocateService `inject:"#RepositoryLocateService"`
	LocalRepositoryDAO dao0x5af8d0.LocalRepositoryDAO `inject:"#RepositoryDAO"`
}


type pComMainRepositoryServiceImpl struct {
	instance *implservice0x22327c.MainRepositoryServiceImpl
	 markup0x23084a.Component `id:"MainRepositoryService"`
}


type pComPipeServiceImpl struct {
	instance *implservice0x22327c.PipeServiceImpl
	 markup0x23084a.Component `id:"PipeService" initMethod:"Init"`
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
	LocateService service0x3e063d.RepositoryLocateService `inject:"#RepositoryLocateService"`
	SearchService service0x3e063d.RepositorySearchService `inject:"#RepositorySearchService"`
	RepositoryService service0x3e063d.LocalRepositoryService `inject:"#RepositoryService"`
}


type pComRepositoryLocateServiceImpl struct {
	instance *implservice0x22327c.RepositoryLocateServiceImpl
	 markup0x23084a.Component `id:"RepositoryLocateService"`
}


type pComRepositorySearchServiceImpl struct {
	instance *implservice0x22327c.RepositorySearchServiceImpl
	 markup0x23084a.Component `id:"RepositorySearchService"`
	LocateService service0x3e063d.RepositoryLocateService `inject:"#RepositoryLocateService"`
}


type pComRunIntentServiceImpl struct {
	instance *implservice0x22327c.RunIntentServiceImpl
	 markup0x23084a.Component `id:"RunIntentService"`
	ExecutableService service0x3e063d.ExecutableService `inject:"#ExecutableService"`
	PipeService service0x3e063d.PipeService `inject:"#PipeService"`
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


type pComIntentController struct {
	instance *controller0x9dc399.IntentController
	 markup0x23084a.RestController `class:"rest-controller"`
	IntentService service0x3e063d.IntentService `inject:"#IntentService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComLocalRepositoryController struct {
	instance *controller0x9dc399.LocalRepositoryController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoService service0x3e063d.LocalRepositoryService `inject:"#RepositoryService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComMainRepositoryController struct {
	instance *controller0x9dc399.MainRepositoryController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoService service0x3e063d.MainRepositoryService `inject:"#MainRepositoryService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComPipeController struct {
	instance *controller0x9dc399.PipeController
	 markup0x23084a.RestController `class:"rest-controller"`
	PipeService service0x3e063d.PipeService `inject:"#PipeService"`
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
	ImportService service0x3e063d.RepositoryImportService `inject:"#RepositoryImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComRunIntentController struct {
	instance *controller0x9dc399.RunIntentController
	 markup0x23084a.RestController `class:"rest-controller"`
	RunIntentService service0x3e063d.RunIntentService `inject:"#RunIntentService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}

