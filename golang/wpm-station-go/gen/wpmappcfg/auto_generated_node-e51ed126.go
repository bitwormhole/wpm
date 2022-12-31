// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmappcfg

import (
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	markup0x23084a "github.com/bitwormhole/starter/markup"
	dao0xda26a8 "github.com/bitwormhole/wpm/app/data/dao"
	service0xe6cb35 "github.com/bitwormhole/wpm/app/service"
	impldao0x204069 "github.com/bitwormhole/wpm/app/support/impldao"
	implservice0x5e39ce "github.com/bitwormhole/wpm/app/support/implservice"
	controller0x906cab "github.com/bitwormhole/wpm/app/web/controller"
)

type pComGormDBAgentImpl struct {
	instance *impldao0x204069.GormDBAgentImpl
	 markup0x23084a.Component `id:"GormDBAgent" initMethod:"Init"`
	Source datasource0x68a737.Source `inject:"#gorm-datasource-default"`
}


type pComExecutableDaoImpl struct {
	instance *impldao0x204069.ExecutableDaoImpl
	 markup0x23084a.Component `id:"ExecutableDAO"`
	Agent impldao0x204069.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0xe6cb35.UUIDGenService `inject:"#UUIDGenService"`
}


type pComIntentDaoImpl struct {
	instance *impldao0x204069.IntentDaoImpl
	 markup0x23084a.Component `id:"IntentDAO"`
	Agent impldao0x204069.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0xe6cb35.UUIDGenService `inject:"#UUIDGenService"`
}


type pComProjectDaoImpl struct {
	instance *impldao0x204069.ProjectDaoImpl
	 markup0x23084a.Component `id:"ProjectDAO"`
	Agent impldao0x204069.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0xe6cb35.UUIDGenService `inject:"#UUIDGenService"`
}


type pComRepositoryDaoImpl struct {
	instance *impldao0x204069.RepositoryDaoImpl
	 markup0x23084a.Component `id:"RepositoryDAO"`
	Agent impldao0x204069.GormDBAgent `inject:"#GormDBAgent"`
	UUIDGenService service0xe6cb35.UUIDGenService `inject:"#UUIDGenService"`
}


type pComExecutableImportServiceImpl struct {
	instance *implservice0x5e39ce.ExecutableImportServiceImpl
	 markup0x23084a.Component `id:"ExecutableImportService"`
	ExecutableService service0xe6cb35.ExecutableService `inject:"#ExecutableService"`
}


type pComExecutableServiceImpl struct {
	instance *implservice0x5e39ce.ExecutableServiceImpl
	 markup0x23084a.Component `id:"ExecutableService"`
	ExecutableDAO dao0xda26a8.ExecutableDAO `inject:"#ExecutableDAO"`
}


type pComIntentServiceImpl struct {
	instance *implservice0x5e39ce.IntentServiceImpl
	 markup0x23084a.Component `id:"IntentService"`
	IntentDAO dao0xda26a8.IntentDAO `inject:"#IntentDAO"`
}


type pComProjectImportServiceImpl struct {
	instance *implservice0x5e39ce.ProjectImportServiceImpl
	 markup0x23084a.Component `id:"ProjectImportService"`
}


type pComProjectServiceImpl struct {
	instance *implservice0x5e39ce.ProjectServiceImpl
	 markup0x23084a.Component `id:"ProjectService"`
	UUIDGenService service0xe6cb35.UUIDGenService `inject:"#UUIDGenService"`
	ProjectDAO dao0xda26a8.ProjectDAO `inject:"#ProjectDAO"`
}


type pComRepositoryImportServiceImpl struct {
	instance *implservice0x5e39ce.RepositoryImportServiceImpl
	 markup0x23084a.Component `id:"RepositoryImportService"`
	LocateService service0xe6cb35.RepositoryLocateService `inject:"#RepositoryLocateService"`
	SearchService service0xe6cb35.RepositorySearchService `inject:"#RepositorySearchService"`
	RepositoryService service0xe6cb35.RepositoryService `inject:"#RepositoryService"`
}


type pComRepositoryLocateServiceImpl struct {
	instance *implservice0x5e39ce.RepositoryLocateServiceImpl
	 markup0x23084a.Component `id:"RepositoryLocateService"`
}


type pComRepositorySearchServiceImpl struct {
	instance *implservice0x5e39ce.RepositorySearchServiceImpl
	 markup0x23084a.Component `id:"RepositorySearchService"`
	LocateService service0xe6cb35.RepositoryLocateService `inject:"#RepositoryLocateService"`
}


type pComRepositoryServiceImpl struct {
	instance *implservice0x5e39ce.RepositoryServiceImpl
	 markup0x23084a.Component `id:"RepositoryService"`
	UUIDGenService service0xe6cb35.UUIDGenService `inject:"#UUIDGenService"`
	LocateService service0xe6cb35.RepositoryLocateService `inject:"#RepositoryLocateService"`
	RepositoryDAO dao0xda26a8.RepositoryDAO `inject:"#RepositoryDAO"`
}


type pComUUIDGenServiceImpl struct {
	instance *implservice0x5e39ce.UUIDGenServiceImpl
	 markup0x23084a.Component `id:"UUIDGenService" initMethod:"Init"`
}


type pComAuthController struct {
	instance *controller0x906cab.AuthController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoService service0xe6cb35.RepositoryService `inject:"#RepositoryService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComExecutableController struct {
	instance *controller0x906cab.ExecutableController
	 markup0x23084a.RestController `class:"rest-controller"`
	ExecutableService service0xe6cb35.ExecutableService `inject:"#ExecutableService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComExecutableImportController struct {
	instance *controller0x906cab.ExecutableImportController
	 markup0x23084a.RestController `class:"rest-controller"`
	ExecutableImportService service0xe6cb35.ExecutableImportService `inject:"#ExecutableImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComIntentController struct {
	instance *controller0x906cab.IntentController
	 markup0x23084a.RestController `class:"rest-controller"`
	IntentService service0xe6cb35.IntentService `inject:"#IntentService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComProjectController struct {
	instance *controller0x906cab.ProjectController
	 markup0x23084a.RestController `class:"rest-controller"`
	ProjectService service0xe6cb35.ProjectService `inject:"#ProjectService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComProjectImportController struct {
	instance *controller0x906cab.ProjectImportController
	 markup0x23084a.RestController `class:"rest-controller"`
	ProjectImportService service0xe6cb35.ProjectImportService `inject:"#ProjectImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComRepositoryController struct {
	instance *controller0x906cab.RepositoryController
	 markup0x23084a.RestController `class:"rest-controller"`
	RepoService service0xe6cb35.RepositoryService `inject:"#RepositoryService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComRepositoryImportController struct {
	instance *controller0x906cab.RepositoryImportController
	 markup0x23084a.RestController `class:"rest-controller"`
	ImportService service0xe6cb35.RepositoryImportService `inject:"#RepositoryImportService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}

