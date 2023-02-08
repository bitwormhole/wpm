// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmserver

import (
	store0x8467b3 "github.com/bitwormhole/gitlib/git/store"
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
	implservice0x5a8f41 "github.com/bitwormhole/wpm/common/implservice"
	dao0x5af8d0 "github.com/bitwormhole/wpm/server/data/dao"
	service0x3e063d "github.com/bitwormhole/wpm/server/service"
	support0xf47d7f "github.com/bitwormhole/wpm/server/support"
	impldao0x73998b "github.com/bitwormhole/wpm/server/support/impldao"
	implservice0x22327c "github.com/bitwormhole/wpm/server/support/implservice"
	intents0xec84e7 "github.com/bitwormhole/wpm/server/utils/intents"
	filters0x5d53d8 "github.com/bitwormhole/wpm/server/utils/intents/filters"
	controller0x9dc399 "github.com/bitwormhole/wpm/server/web/controller"
	filter0x8aa8f6 "github.com/bitwormhole/wpm/server/web/filter"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: GormDBAgent
	cominfobuilder.Next()
	cominfobuilder.ID("GormDBAgent").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComGormDBAgentImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ExampleDAO
	cominfobuilder.Next()
	cominfobuilder.ID("ExampleDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ExecutableDAO
	cominfobuilder.Next()
	cominfobuilder.ID("ExecutableDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: IntentTemplateDAO
	cominfobuilder.Next()
	cominfobuilder.ID("IntentTemplateDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentTemplateDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: LocalRepositoryDAO
	cominfobuilder.Next()
	cominfobuilder.ID("LocalRepositoryDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: MediaDAO
	cominfobuilder.Next()
	cominfobuilder.ID("MediaDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMediaDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ProjectDAO
	cominfobuilder.Next()
	cominfobuilder.ID("ProjectDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: AboutService
	cominfobuilder.Next()
	cominfobuilder.ID("AboutService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAboutServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: AppDataService
	cominfobuilder.Next()
	cominfobuilder.ID("AppDataService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAppDataServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: AppIconService
	cominfobuilder.Next()
	cominfobuilder.ID("AppIconService").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAppIconServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ExampleService
	cominfobuilder.Next()
	cominfobuilder.ID("ExampleService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ExecutableImportService
	cominfobuilder.Next()
	cominfobuilder.ID("ExecutableImportService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableImportServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ExecutableService
	cominfobuilder.Next()
	cominfobuilder.ID("ExecutableService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: FileQueryService
	cominfobuilder.Next()
	cominfobuilder.ID("FileQueryService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFileQueryServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: IntentTemplateService
	cominfobuilder.Next()
	cominfobuilder.ID("IntentTemplateService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentTemplateServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: IntentHandlerService
	cominfobuilder.Next()
	cominfobuilder.ID("IntentHandlerService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentHandlerImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: LocalRepositoryFinder
	cominfobuilder.Next()
	cominfobuilder.ID("LocalRepositoryFinder").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalRepositoryFinderImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: LocalRepositoryService
	cominfobuilder.Next()
	cominfobuilder.ID("LocalRepositoryService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalRepositoryServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: LocalRepositoryStateLoader
	cominfobuilder.Next()
	cominfobuilder.ID("LocalRepositoryStateLoader").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalRepositoryStateLoaderImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: MainRepositoryService
	cominfobuilder.Next()
	cominfobuilder.ID("MainRepositoryService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMainRepositoryServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: MediaService
	cominfobuilder.Next()
	cominfobuilder.ID("MediaService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMediaServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ProjectImportService
	cominfobuilder.Next()
	cominfobuilder.ID("ProjectImportService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectImportServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ProjectService
	cominfobuilder.Next()
	cominfobuilder.ID("ProjectService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: RemoteRepositoryService
	cominfobuilder.Next()
	cominfobuilder.ID("RemoteRepositoryService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRemoteRepositoryServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: RepositoryImportService
	cominfobuilder.Next()
	cominfobuilder.ID("RepositoryImportService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryImportServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: IntentService
	cominfobuilder.Next()
	cominfobuilder.ID("IntentService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRunIntentServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: UUIDGenService
	cominfobuilder.Next()
	cominfobuilder.ID("UUIDGenService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUUIDGenServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com27-support0xf47d7f.WpmDataSource
	cominfobuilder.Next()
	cominfobuilder.ID("com27-support0xf47d7f.WpmDataSource").Class("starter-gorm-source-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComWpmDataSource{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com28-filters0x5d53d8.ExecuteIntentFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com28-filters0x5d53d8.ExecuteIntentFilter").Class("intent-filter-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecuteIntentFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com29-filters0x5d53d8.IntentFilterFor
	cominfobuilder.Next()
	cominfobuilder.ID("com29-filters0x5d53d8.IntentFilterFor").Class("intent-filter-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentFilterFor{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com30-filters0x5d53d8.IntentFilterForChrome
	cominfobuilder.Next()
	cominfobuilder.ID("com30-filters0x5d53d8.IntentFilterForChrome").Class("intent-filter-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentFilterForChrome{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com31-filters0x5d53d8.IntentFilterForCmd
	cominfobuilder.Next()
	cominfobuilder.ID("com31-filters0x5d53d8.IntentFilterForCmd").Class("intent-filter-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentFilterForCmd{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com32-filters0x5d53d8.IntentFilterForExplorer
	cominfobuilder.Next()
	cominfobuilder.ID("com32-filters0x5d53d8.IntentFilterForExplorer").Class("intent-filter-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentFilterForExplorer{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com33-filters0x5d53d8.IntentFilterForFirefox
	cominfobuilder.Next()
	cominfobuilder.ID("com33-filters0x5d53d8.IntentFilterForFirefox").Class("intent-filter-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentFilterForFirefox{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com34-filters0x5d53d8.IntentFilterForMsEdge
	cominfobuilder.Next()
	cominfobuilder.ID("com34-filters0x5d53d8.IntentFilterForMsEdge").Class("intent-filter-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentFilterForMsEdge{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com35-filters0x5d53d8.IntentFilterForNautilus
	cominfobuilder.Next()
	cominfobuilder.ID("com35-filters0x5d53d8.IntentFilterForNautilus").Class("intent-filter-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentFilterForNautilus{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com36-filters0x5d53d8.IntentFilterForPowerShell
	cominfobuilder.Next()
	cominfobuilder.ID("com36-filters0x5d53d8.IntentFilterForPowerShell").Class("intent-filter-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentFilterForPowerShell{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com37-filters0x5d53d8.IntentFilterForVscode
	cominfobuilder.Next()
	cominfobuilder.ID("com37-filters0x5d53d8.IntentFilterForVscode").Class("intent-filter-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentFilterForVscode{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: intent-filter-manager
	cominfobuilder.Next()
	cominfobuilder.ID("intent-filter-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFilterManagerImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com39-controller0x9dc399.AboutController
	cominfobuilder.Next()
	cominfobuilder.ID("com39-controller0x9dc399.AboutController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAboutController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com40-controller0x9dc399.AuthController
	cominfobuilder.Next()
	cominfobuilder.ID("com40-controller0x9dc399.AuthController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAuthController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com41-controller0x9dc399.ExampleController
	cominfobuilder.Next()
	cominfobuilder.ID("com41-controller0x9dc399.ExampleController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com42-controller0x9dc399.ExecutableController
	cominfobuilder.Next()
	cominfobuilder.ID("com42-controller0x9dc399.ExecutableController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com43-controller0x9dc399.ExecutableImportController
	cominfobuilder.Next()
	cominfobuilder.ID("com43-controller0x9dc399.ExecutableImportController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableImportController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com44-controller0x9dc399.FileQueryController
	cominfobuilder.Next()
	cominfobuilder.ID("com44-controller0x9dc399.FileQueryController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFileQueryController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com45-controller0x9dc399.RunIntentController
	cominfobuilder.Next()
	cominfobuilder.ID("com45-controller0x9dc399.RunIntentController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRunIntentController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com46-controller0x9dc399.IntentTemplateController
	cominfobuilder.Next()
	cominfobuilder.ID("com46-controller0x9dc399.IntentTemplateController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentTemplateController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com47-controller0x9dc399.LocalRepositoryController
	cominfobuilder.Next()
	cominfobuilder.ID("com47-controller0x9dc399.LocalRepositoryController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalRepositoryController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com48-controller0x9dc399.MediaController
	cominfobuilder.Next()
	cominfobuilder.ID("com48-controller0x9dc399.MediaController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMediaController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com49-controller0x9dc399.PlatformController
	cominfobuilder.Next()
	cominfobuilder.ID("com49-controller0x9dc399.PlatformController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPlatformController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com50-controller0x9dc399.ProjectController
	cominfobuilder.Next()
	cominfobuilder.ID("com50-controller0x9dc399.ProjectController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com51-controller0x9dc399.ProjectImportController
	cominfobuilder.Next()
	cominfobuilder.ID("com51-controller0x9dc399.ProjectImportController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectImportController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com52-controller0x9dc399.RemoteRepositoryController
	cominfobuilder.Next()
	cominfobuilder.ID("com52-controller0x9dc399.RemoteRepositoryController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRemoteRepositoryController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com53-controller0x9dc399.RepositoryImportController
	cominfobuilder.Next()
	cominfobuilder.ID("com53-controller0x9dc399.RepositoryImportController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryImportController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com54-filter0x8aa8f6.HostFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com54-filter0x8aa8f6.HostFilter").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHostFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com55-filter0x8aa8f6.HTTP404Filter
	cominfobuilder.Next()
	cominfobuilder.ID("com55-filter0x8aa8f6.HTTP404Filter").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHTTP404Filter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: PlatformService
	cominfobuilder.Next()
	cominfobuilder.ID("PlatformService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPlatformServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com57-implservice0x5a8f41.LinuxPlatformServiceImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com57-implservice0x5a8f41.LinuxPlatformServiceImpl").Class("PlatformProviderRegistry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLinuxPlatformServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com58-implservice0x5a8f41.WindowsPlatformServiceImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com58-implservice0x5a8f41.WindowsPlatformServiceImpl").Class("PlatformProviderRegistry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComWindowsPlatformServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ProfileService
	cominfobuilder.Next()
	cominfobuilder.ID("ProfileService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProfileServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGormDBAgentImpl : the factory of component: GormDBAgent
type comFactory4pComGormDBAgentImpl struct {

    mPrototype * impldao0x73998b.GormDBAgentImpl

	
	mSourcesSelector config.InjectionSelector

}

func (inst * comFactory4pComGormDBAgentImpl) init() application.ComponentFactory {

	
	inst.mSourcesSelector = config.NewInjectionSelector("#starter-gorm-source-manager",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGormDBAgentImpl) newObject() * impldao0x73998b.GormDBAgentImpl {
	return & impldao0x73998b.GormDBAgentImpl {}
}

func (inst * comFactory4pComGormDBAgentImpl) castObject(instance application.ComponentInstance) * impldao0x73998b.GormDBAgentImpl {
	return instance.Get().(*impldao0x73998b.GormDBAgentImpl)
}

func (inst * comFactory4pComGormDBAgentImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComGormDBAgentImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComGormDBAgentImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComGormDBAgentImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGormDBAgentImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGormDBAgentImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Sources = inst.getterForFieldSourcesSelector(context)
	return context.LastError()
}

//getterForFieldSourcesSelector
func (inst * comFactory4pComGormDBAgentImpl) getterForFieldSourcesSelector (context application.InstanceContext) datasource0x68a737.SourceManager {

	o1 := inst.mSourcesSelector.GetOne(context)
	o2, ok := o1.(datasource0x68a737.SourceManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "GormDBAgent")
		eb.Set("field", "Sources")
		eb.Set("type1", "?")
		eb.Set("type2", "datasource0x68a737.SourceManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExampleDaoImpl : the factory of component: ExampleDAO
type comFactory4pComExampleDaoImpl struct {

    mPrototype * impldao0x73998b.ExampleDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComExampleDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExampleDaoImpl) newObject() * impldao0x73998b.ExampleDaoImpl {
	return & impldao0x73998b.ExampleDaoImpl {}
}

func (inst * comFactory4pComExampleDaoImpl) castObject(instance application.ComponentInstance) * impldao0x73998b.ExampleDaoImpl {
	return instance.Get().(*impldao0x73998b.ExampleDaoImpl)
}

func (inst * comFactory4pComExampleDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExampleDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExampleDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExampleDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComExampleDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) impldao0x73998b.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(impldao0x73998b.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExampleDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "impldao0x73998b.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComExampleDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExampleDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableDaoImpl : the factory of component: ExecutableDAO
type comFactory4pComExecutableDaoImpl struct {

    mPrototype * impldao0x73998b.ExecutableDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableDaoImpl) newObject() * impldao0x73998b.ExecutableDaoImpl {
	return & impldao0x73998b.ExecutableDaoImpl {}
}

func (inst * comFactory4pComExecutableDaoImpl) castObject(instance application.ComponentInstance) * impldao0x73998b.ExecutableDaoImpl {
	return instance.Get().(*impldao0x73998b.ExecutableDaoImpl)
}

func (inst * comFactory4pComExecutableDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExecutableDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExecutableDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExecutableDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComExecutableDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) impldao0x73998b.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(impldao0x73998b.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "impldao0x73998b.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComExecutableDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentTemplateDaoImpl : the factory of component: IntentTemplateDAO
type comFactory4pComIntentTemplateDaoImpl struct {

    mPrototype * impldao0x73998b.IntentTemplateDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComIntentTemplateDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentTemplateDaoImpl) newObject() * impldao0x73998b.IntentTemplateDaoImpl {
	return & impldao0x73998b.IntentTemplateDaoImpl {}
}

func (inst * comFactory4pComIntentTemplateDaoImpl) castObject(instance application.ComponentInstance) * impldao0x73998b.IntentTemplateDaoImpl {
	return instance.Get().(*impldao0x73998b.IntentTemplateDaoImpl)
}

func (inst * comFactory4pComIntentTemplateDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentTemplateDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentTemplateDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentTemplateDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentTemplateDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentTemplateDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComIntentTemplateDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) impldao0x73998b.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(impldao0x73998b.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentTemplateDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "impldao0x73998b.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComIntentTemplateDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentTemplateDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryDaoImpl : the factory of component: LocalRepositoryDAO
type comFactory4pComRepositoryDaoImpl struct {

    mPrototype * impldao0x73998b.RepositoryDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryDaoImpl) newObject() * impldao0x73998b.RepositoryDaoImpl {
	return & impldao0x73998b.RepositoryDaoImpl {}
}

func (inst * comFactory4pComRepositoryDaoImpl) castObject(instance application.ComponentInstance) * impldao0x73998b.RepositoryDaoImpl {
	return instance.Get().(*impldao0x73998b.RepositoryDaoImpl)
}

func (inst * comFactory4pComRepositoryDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositoryDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositoryDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositoryDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComRepositoryDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) impldao0x73998b.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(impldao0x73998b.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "impldao0x73998b.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComRepositoryDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMediaDaoImpl : the factory of component: MediaDAO
type comFactory4pComMediaDaoImpl struct {

    mPrototype * impldao0x73998b.MediaDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComMediaDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMediaDaoImpl) newObject() * impldao0x73998b.MediaDaoImpl {
	return & impldao0x73998b.MediaDaoImpl {}
}

func (inst * comFactory4pComMediaDaoImpl) castObject(instance application.ComponentInstance) * impldao0x73998b.MediaDaoImpl {
	return instance.Get().(*impldao0x73998b.MediaDaoImpl)
}

func (inst * comFactory4pComMediaDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMediaDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMediaDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMediaDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMediaDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMediaDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComMediaDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) impldao0x73998b.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(impldao0x73998b.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MediaDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "impldao0x73998b.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComMediaDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MediaDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectDaoImpl : the factory of component: ProjectDAO
type comFactory4pComProjectDaoImpl struct {

    mPrototype * impldao0x73998b.ProjectDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectDaoImpl) newObject() * impldao0x73998b.ProjectDaoImpl {
	return & impldao0x73998b.ProjectDaoImpl {}
}

func (inst * comFactory4pComProjectDaoImpl) castObject(instance application.ComponentInstance) * impldao0x73998b.ProjectDaoImpl {
	return instance.Get().(*impldao0x73998b.ProjectDaoImpl)
}

func (inst * comFactory4pComProjectDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComProjectDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) impldao0x73998b.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(impldao0x73998b.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "impldao0x73998b.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComProjectDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAboutServiceImpl : the factory of component: AboutService
type comFactory4pComAboutServiceImpl struct {

    mPrototype * implservice0x22327c.AboutServiceImpl

	
	mProfileSelector config.InjectionSelector
	mNameSelector config.InjectionSelector
	mTitleSelector config.InjectionSelector
	mCopyrightSelector config.InjectionSelector
	mPlatformServiceSelector config.InjectionSelector
	mProfileServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComAboutServiceImpl) init() application.ComponentFactory {

	
	inst.mProfileSelector = config.NewInjectionSelector("${application.profiles.active}",nil)
	inst.mNameSelector = config.NewInjectionSelector("${application.about.name}",nil)
	inst.mTitleSelector = config.NewInjectionSelector("${application.about.title}",nil)
	inst.mCopyrightSelector = config.NewInjectionSelector("${application.about.copyright}",nil)
	inst.mPlatformServiceSelector = config.NewInjectionSelector("#PlatformService",nil)
	inst.mProfileServiceSelector = config.NewInjectionSelector("#ProfileService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAboutServiceImpl) newObject() * implservice0x22327c.AboutServiceImpl {
	return & implservice0x22327c.AboutServiceImpl {}
}

func (inst * comFactory4pComAboutServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.AboutServiceImpl {
	return instance.Get().(*implservice0x22327c.AboutServiceImpl)
}

func (inst * comFactory4pComAboutServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAboutServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAboutServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAboutServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAboutServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAboutServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Profile = inst.getterForFieldProfileSelector(context)
	obj.Name = inst.getterForFieldNameSelector(context)
	obj.Title = inst.getterForFieldTitleSelector(context)
	obj.Copyright = inst.getterForFieldCopyrightSelector(context)
	obj.PlatformService = inst.getterForFieldPlatformServiceSelector(context)
	obj.ProfileService = inst.getterForFieldProfileServiceSelector(context)
	return context.LastError()
}

//getterForFieldProfileSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldProfileSelector (context application.InstanceContext) string {
    return inst.mProfileSelector.GetString(context)
}

//getterForFieldNameSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldNameSelector (context application.InstanceContext) string {
    return inst.mNameSelector.GetString(context)
}

//getterForFieldTitleSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldTitleSelector (context application.InstanceContext) string {
    return inst.mTitleSelector.GetString(context)
}

//getterForFieldCopyrightSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldCopyrightSelector (context application.InstanceContext) string {
    return inst.mCopyrightSelector.GetString(context)
}

//getterForFieldPlatformServiceSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldPlatformServiceSelector (context application.InstanceContext) service0x3e063d.PlatformService {

	o1 := inst.mPlatformServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.PlatformService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "AboutService")
		eb.Set("field", "PlatformService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.PlatformService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProfileServiceSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldProfileServiceSelector (context application.InstanceContext) service0x3e063d.ProfileService {

	o1 := inst.mProfileServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProfileService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "AboutService")
		eb.Set("field", "ProfileService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProfileService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAppDataServiceImpl : the factory of component: AppDataService
type comFactory4pComAppDataServiceImpl struct {

    mPrototype * implservice0x22327c.AppDataServiceImpl

	
	mProfileServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComAppDataServiceImpl) init() application.ComponentFactory {

	
	inst.mProfileServiceSelector = config.NewInjectionSelector("#ProfileService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAppDataServiceImpl) newObject() * implservice0x22327c.AppDataServiceImpl {
	return & implservice0x22327c.AppDataServiceImpl {}
}

func (inst * comFactory4pComAppDataServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.AppDataServiceImpl {
	return instance.Get().(*implservice0x22327c.AppDataServiceImpl)
}

func (inst * comFactory4pComAppDataServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAppDataServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAppDataServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAppDataServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAppDataServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAppDataServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ProfileService = inst.getterForFieldProfileServiceSelector(context)
	return context.LastError()
}

//getterForFieldProfileServiceSelector
func (inst * comFactory4pComAppDataServiceImpl) getterForFieldProfileServiceSelector (context application.InstanceContext) service0x3e063d.ProfileService {

	o1 := inst.mProfileServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProfileService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "AppDataService")
		eb.Set("field", "ProfileService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProfileService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAppIconServiceImpl : the factory of component: AppIconService
type comFactory4pComAppIconServiceImpl struct {

    mPrototype * implservice0x22327c.AppIconServiceImpl

	
	mPropsNameSelector config.InjectionSelector
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComAppIconServiceImpl) init() application.ComponentFactory {

	
	inst.mPropsNameSelector = config.NewInjectionSelector("${wpm.exe-icons.properties}",nil)
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAppIconServiceImpl) newObject() * implservice0x22327c.AppIconServiceImpl {
	return & implservice0x22327c.AppIconServiceImpl {}
}

func (inst * comFactory4pComAppIconServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.AppIconServiceImpl {
	return instance.Get().(*implservice0x22327c.AppIconServiceImpl)
}

func (inst * comFactory4pComAppIconServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAppIconServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAppIconServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAppIconServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAppIconServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAppIconServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.PropsName = inst.getterForFieldPropsNameSelector(context)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldPropsNameSelector
func (inst * comFactory4pComAppIconServiceImpl) getterForFieldPropsNameSelector (context application.InstanceContext) string {
    return inst.mPropsNameSelector.GetString(context)
}

//getterForFieldContextSelector
func (inst * comFactory4pComAppIconServiceImpl) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExampleServiceImpl : the factory of component: ExampleService
type comFactory4pComExampleServiceImpl struct {

    mPrototype * implservice0x22327c.ExampleServiceImpl

	

}

func (inst * comFactory4pComExampleServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExampleServiceImpl) newObject() * implservice0x22327c.ExampleServiceImpl {
	return & implservice0x22327c.ExampleServiceImpl {}
}

func (inst * comFactory4pComExampleServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.ExampleServiceImpl {
	return instance.Get().(*implservice0x22327c.ExampleServiceImpl)
}

func (inst * comFactory4pComExampleServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExampleServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExampleServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExampleServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableImportServiceImpl : the factory of component: ExecutableImportService
type comFactory4pComExecutableImportServiceImpl struct {

    mPrototype * implservice0x22327c.ExecutableImportServiceImpl

	
	mExecutableServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableImportServiceImpl) init() application.ComponentFactory {

	
	inst.mExecutableServiceSelector = config.NewInjectionSelector("#ExecutableService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableImportServiceImpl) newObject() * implservice0x22327c.ExecutableImportServiceImpl {
	return & implservice0x22327c.ExecutableImportServiceImpl {}
}

func (inst * comFactory4pComExecutableImportServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.ExecutableImportServiceImpl {
	return instance.Get().(*implservice0x22327c.ExecutableImportServiceImpl)
}

func (inst * comFactory4pComExecutableImportServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExecutableImportServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExecutableImportServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExecutableImportServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableImportServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableImportServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ExecutableService = inst.getterForFieldExecutableServiceSelector(context)
	return context.LastError()
}

//getterForFieldExecutableServiceSelector
func (inst * comFactory4pComExecutableImportServiceImpl) getterForFieldExecutableServiceSelector (context application.InstanceContext) service0x3e063d.ExecutableService {

	o1 := inst.mExecutableServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ExecutableService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableImportService")
		eb.Set("field", "ExecutableService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ExecutableService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableServiceImpl : the factory of component: ExecutableService
type comFactory4pComExecutableServiceImpl struct {

    mPrototype * implservice0x22327c.ExecutableServiceImpl

	
	mExecutableDAOSelector config.InjectionSelector
	mIconServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableServiceImpl) init() application.ComponentFactory {

	
	inst.mExecutableDAOSelector = config.NewInjectionSelector("#ExecutableDAO",nil)
	inst.mIconServiceSelector = config.NewInjectionSelector("#AppIconService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableServiceImpl) newObject() * implservice0x22327c.ExecutableServiceImpl {
	return & implservice0x22327c.ExecutableServiceImpl {}
}

func (inst * comFactory4pComExecutableServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.ExecutableServiceImpl {
	return instance.Get().(*implservice0x22327c.ExecutableServiceImpl)
}

func (inst * comFactory4pComExecutableServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExecutableServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExecutableServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExecutableServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ExecutableDAO = inst.getterForFieldExecutableDAOSelector(context)
	obj.IconService = inst.getterForFieldIconServiceSelector(context)
	return context.LastError()
}

//getterForFieldExecutableDAOSelector
func (inst * comFactory4pComExecutableServiceImpl) getterForFieldExecutableDAOSelector (context application.InstanceContext) dao0x5af8d0.ExecutableDAO {

	o1 := inst.mExecutableDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.ExecutableDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableService")
		eb.Set("field", "ExecutableDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.ExecutableDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldIconServiceSelector
func (inst * comFactory4pComExecutableServiceImpl) getterForFieldIconServiceSelector (context application.InstanceContext) service0x3e063d.AppIconService {

	o1 := inst.mIconServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AppIconService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableService")
		eb.Set("field", "IconService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AppIconService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFileQueryServiceImpl : the factory of component: FileQueryService
type comFactory4pComFileQueryServiceImpl struct {

    mPrototype * implservice0x22327c.FileQueryServiceImpl

	

}

func (inst * comFactory4pComFileQueryServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFileQueryServiceImpl) newObject() * implservice0x22327c.FileQueryServiceImpl {
	return & implservice0x22327c.FileQueryServiceImpl {}
}

func (inst * comFactory4pComFileQueryServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.FileQueryServiceImpl {
	return instance.Get().(*implservice0x22327c.FileQueryServiceImpl)
}

func (inst * comFactory4pComFileQueryServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFileQueryServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFileQueryServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFileQueryServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileQueryServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileQueryServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentTemplateServiceImpl : the factory of component: IntentTemplateService
type comFactory4pComIntentTemplateServiceImpl struct {

    mPrototype * implservice0x22327c.IntentTemplateServiceImpl

	
	mIntentTempDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComIntentTemplateServiceImpl) init() application.ComponentFactory {

	
	inst.mIntentTempDAOSelector = config.NewInjectionSelector("#IntentTemplateDAO",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentTemplateServiceImpl) newObject() * implservice0x22327c.IntentTemplateServiceImpl {
	return & implservice0x22327c.IntentTemplateServiceImpl {}
}

func (inst * comFactory4pComIntentTemplateServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.IntentTemplateServiceImpl {
	return instance.Get().(*implservice0x22327c.IntentTemplateServiceImpl)
}

func (inst * comFactory4pComIntentTemplateServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentTemplateServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentTemplateServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentTemplateServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentTemplateServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentTemplateServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.IntentTempDAO = inst.getterForFieldIntentTempDAOSelector(context)
	return context.LastError()
}

//getterForFieldIntentTempDAOSelector
func (inst * comFactory4pComIntentTemplateServiceImpl) getterForFieldIntentTempDAOSelector (context application.InstanceContext) dao0x5af8d0.IntentTemplateDAO {

	o1 := inst.mIntentTempDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.IntentTemplateDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentTemplateService")
		eb.Set("field", "IntentTempDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.IntentTemplateDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentHandlerImpl : the factory of component: IntentHandlerService
type comFactory4pComIntentHandlerImpl struct {

    mPrototype * implservice0x22327c.IntentHandlerImpl

	

}

func (inst * comFactory4pComIntentHandlerImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentHandlerImpl) newObject() * implservice0x22327c.IntentHandlerImpl {
	return & implservice0x22327c.IntentHandlerImpl {}
}

func (inst * comFactory4pComIntentHandlerImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.IntentHandlerImpl {
	return instance.Get().(*implservice0x22327c.IntentHandlerImpl)
}

func (inst * comFactory4pComIntentHandlerImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentHandlerImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentHandlerImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentHandlerImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentHandlerImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentHandlerImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalRepositoryFinderImpl : the factory of component: LocalRepositoryFinder
type comFactory4pComLocalRepositoryFinderImpl struct {

    mPrototype * implservice0x22327c.LocalRepositoryFinderImpl

	
	mGitLibAgentSelector config.InjectionSelector

}

func (inst * comFactory4pComLocalRepositoryFinderImpl) init() application.ComponentFactory {

	
	inst.mGitLibAgentSelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) newObject() * implservice0x22327c.LocalRepositoryFinderImpl {
	return & implservice0x22327c.LocalRepositoryFinderImpl {}
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.LocalRepositoryFinderImpl {
	return instance.Get().(*implservice0x22327c.LocalRepositoryFinderImpl)
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.GitLibAgent = inst.getterForFieldGitLibAgentSelector(context)
	return context.LastError()
}

//getterForFieldGitLibAgentSelector
func (inst * comFactory4pComLocalRepositoryFinderImpl) getterForFieldGitLibAgentSelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mGitLibAgentSelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryFinder")
		eb.Set("field", "GitLibAgent")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalRepositoryServiceImpl : the factory of component: LocalRepositoryService
type comFactory4pComLocalRepositoryServiceImpl struct {

    mPrototype * implservice0x22327c.LocalRepositoryServiceImpl

	
	mUUIDGenServiceSelector config.InjectionSelector
	mRepoFinderSelector config.InjectionSelector
	mLocalRepositoryDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComLocalRepositoryServiceImpl) init() application.ComponentFactory {

	
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)
	inst.mRepoFinderSelector = config.NewInjectionSelector("#LocalRepositoryFinder",nil)
	inst.mLocalRepositoryDAOSelector = config.NewInjectionSelector("#LocalRepositoryDAO",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) newObject() * implservice0x22327c.LocalRepositoryServiceImpl {
	return & implservice0x22327c.LocalRepositoryServiceImpl {}
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.LocalRepositoryServiceImpl {
	return instance.Get().(*implservice0x22327c.LocalRepositoryServiceImpl)
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	obj.RepoFinder = inst.getterForFieldRepoFinderSelector(context)
	obj.LocalRepositoryDAO = inst.getterForFieldLocalRepositoryDAOSelector(context)
	return context.LastError()
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComLocalRepositoryServiceImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryService")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldRepoFinderSelector
func (inst * comFactory4pComLocalRepositoryServiceImpl) getterForFieldRepoFinderSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryFinder {

	o1 := inst.mRepoFinderSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryFinder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryService")
		eb.Set("field", "RepoFinder")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryFinder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLocalRepositoryDAOSelector
func (inst * comFactory4pComLocalRepositoryServiceImpl) getterForFieldLocalRepositoryDAOSelector (context application.InstanceContext) dao0x5af8d0.LocalRepositoryDAO {

	o1 := inst.mLocalRepositoryDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.LocalRepositoryDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryService")
		eb.Set("field", "LocalRepositoryDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.LocalRepositoryDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalRepositoryStateLoaderImpl : the factory of component: LocalRepositoryStateLoader
type comFactory4pComLocalRepositoryStateLoaderImpl struct {

    mPrototype * implservice0x22327c.LocalRepositoryStateLoaderImpl

	
	mLocalRepoServiceSelector config.InjectionSelector
	mDaoSelector config.InjectionSelector
	mGitLibAgentSelector config.InjectionSelector

}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) init() application.ComponentFactory {

	
	inst.mLocalRepoServiceSelector = config.NewInjectionSelector("#LocalRepositoryService",nil)
	inst.mDaoSelector = config.NewInjectionSelector("#LocalRepositoryDAO",nil)
	inst.mGitLibAgentSelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) newObject() * implservice0x22327c.LocalRepositoryStateLoaderImpl {
	return & implservice0x22327c.LocalRepositoryStateLoaderImpl {}
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.LocalRepositoryStateLoaderImpl {
	return instance.Get().(*implservice0x22327c.LocalRepositoryStateLoaderImpl)
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.LocalRepoService = inst.getterForFieldLocalRepoServiceSelector(context)
	obj.Dao = inst.getterForFieldDaoSelector(context)
	obj.GitLibAgent = inst.getterForFieldGitLibAgentSelector(context)
	return context.LastError()
}

//getterForFieldLocalRepoServiceSelector
func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) getterForFieldLocalRepoServiceSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryService {

	o1 := inst.mLocalRepoServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryStateLoader")
		eb.Set("field", "LocalRepoService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldDaoSelector
func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) getterForFieldDaoSelector (context application.InstanceContext) dao0x5af8d0.LocalRepositoryDAO {

	o1 := inst.mDaoSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.LocalRepositoryDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryStateLoader")
		eb.Set("field", "Dao")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.LocalRepositoryDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldGitLibAgentSelector
func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) getterForFieldGitLibAgentSelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mGitLibAgentSelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryStateLoader")
		eb.Set("field", "GitLibAgent")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMainRepositoryServiceImpl : the factory of component: MainRepositoryService
type comFactory4pComMainRepositoryServiceImpl struct {

    mPrototype * implservice0x22327c.MainRepositoryServiceImpl

	
	mAppDataServiceSelector config.InjectionSelector
	mGitLASelector config.InjectionSelector

}

func (inst * comFactory4pComMainRepositoryServiceImpl) init() application.ComponentFactory {

	
	inst.mAppDataServiceSelector = config.NewInjectionSelector("#AppDataService",nil)
	inst.mGitLASelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMainRepositoryServiceImpl) newObject() * implservice0x22327c.MainRepositoryServiceImpl {
	return & implservice0x22327c.MainRepositoryServiceImpl {}
}

func (inst * comFactory4pComMainRepositoryServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.MainRepositoryServiceImpl {
	return instance.Get().(*implservice0x22327c.MainRepositoryServiceImpl)
}

func (inst * comFactory4pComMainRepositoryServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMainRepositoryServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMainRepositoryServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMainRepositoryServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMainRepositoryServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMainRepositoryServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AppDataService = inst.getterForFieldAppDataServiceSelector(context)
	obj.GitLA = inst.getterForFieldGitLASelector(context)
	return context.LastError()
}

//getterForFieldAppDataServiceSelector
func (inst * comFactory4pComMainRepositoryServiceImpl) getterForFieldAppDataServiceSelector (context application.InstanceContext) service0x3e063d.AppDataService {

	o1 := inst.mAppDataServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AppDataService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MainRepositoryService")
		eb.Set("field", "AppDataService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AppDataService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldGitLASelector
func (inst * comFactory4pComMainRepositoryServiceImpl) getterForFieldGitLASelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mGitLASelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MainRepositoryService")
		eb.Set("field", "GitLA")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMediaServiceImpl : the factory of component: MediaService
type comFactory4pComMediaServiceImpl struct {

    mPrototype * implservice0x22327c.MediaServiceImpl

	
	mMediaDAOSelector config.InjectionSelector
	mSysMainRepoServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComMediaServiceImpl) init() application.ComponentFactory {

	
	inst.mMediaDAOSelector = config.NewInjectionSelector("#MediaDAO",nil)
	inst.mSysMainRepoServiceSelector = config.NewInjectionSelector("#MainRepositoryService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMediaServiceImpl) newObject() * implservice0x22327c.MediaServiceImpl {
	return & implservice0x22327c.MediaServiceImpl {}
}

func (inst * comFactory4pComMediaServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.MediaServiceImpl {
	return instance.Get().(*implservice0x22327c.MediaServiceImpl)
}

func (inst * comFactory4pComMediaServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMediaServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMediaServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMediaServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMediaServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMediaServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.MediaDAO = inst.getterForFieldMediaDAOSelector(context)
	obj.SysMainRepoService = inst.getterForFieldSysMainRepoServiceSelector(context)
	return context.LastError()
}

//getterForFieldMediaDAOSelector
func (inst * comFactory4pComMediaServiceImpl) getterForFieldMediaDAOSelector (context application.InstanceContext) dao0x5af8d0.MediaDAO {

	o1 := inst.mMediaDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.MediaDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MediaService")
		eb.Set("field", "MediaDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.MediaDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSysMainRepoServiceSelector
func (inst * comFactory4pComMediaServiceImpl) getterForFieldSysMainRepoServiceSelector (context application.InstanceContext) service0x3e063d.MainRepositoryService {

	o1 := inst.mSysMainRepoServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.MainRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MediaService")
		eb.Set("field", "SysMainRepoService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.MainRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectImportServiceImpl : the factory of component: ProjectImportService
type comFactory4pComProjectImportServiceImpl struct {

    mPrototype * implservice0x22327c.ProjectImportServiceImpl

	

}

func (inst * comFactory4pComProjectImportServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectImportServiceImpl) newObject() * implservice0x22327c.ProjectImportServiceImpl {
	return & implservice0x22327c.ProjectImportServiceImpl {}
}

func (inst * comFactory4pComProjectImportServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.ProjectImportServiceImpl {
	return instance.Get().(*implservice0x22327c.ProjectImportServiceImpl)
}

func (inst * comFactory4pComProjectImportServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectImportServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectImportServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectImportServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectImportServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectImportServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectServiceImpl : the factory of component: ProjectService
type comFactory4pComProjectServiceImpl struct {

    mPrototype * implservice0x22327c.ProjectServiceImpl

	
	mUUIDGenServiceSelector config.InjectionSelector
	mProjectDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectServiceImpl) init() application.ComponentFactory {

	
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)
	inst.mProjectDAOSelector = config.NewInjectionSelector("#ProjectDAO",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectServiceImpl) newObject() * implservice0x22327c.ProjectServiceImpl {
	return & implservice0x22327c.ProjectServiceImpl {}
}

func (inst * comFactory4pComProjectServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.ProjectServiceImpl {
	return instance.Get().(*implservice0x22327c.ProjectServiceImpl)
}

func (inst * comFactory4pComProjectServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	obj.ProjectDAO = inst.getterForFieldProjectDAOSelector(context)
	return context.LastError()
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComProjectServiceImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectService")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProjectDAOSelector
func (inst * comFactory4pComProjectServiceImpl) getterForFieldProjectDAOSelector (context application.InstanceContext) dao0x5af8d0.ProjectDAO {

	o1 := inst.mProjectDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.ProjectDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectService")
		eb.Set("field", "ProjectDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.ProjectDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRemoteRepositoryServiceImpl : the factory of component: RemoteRepositoryService
type comFactory4pComRemoteRepositoryServiceImpl struct {

    mPrototype * implservice0x22327c.RemoteRepositoryServiceImpl

	

}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) newObject() * implservice0x22327c.RemoteRepositoryServiceImpl {
	return & implservice0x22327c.RemoteRepositoryServiceImpl {}
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.RemoteRepositoryServiceImpl {
	return instance.Get().(*implservice0x22327c.RemoteRepositoryServiceImpl)
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryImportServiceImpl : the factory of component: RepositoryImportService
type comFactory4pComRepositoryImportServiceImpl struct {

    mPrototype * implservice0x22327c.RepositoryImportServiceImpl

	
	mRepositoryServiceSelector config.InjectionSelector
	mRepoFinderSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryImportServiceImpl) init() application.ComponentFactory {

	
	inst.mRepositoryServiceSelector = config.NewInjectionSelector("#LocalRepositoryService",nil)
	inst.mRepoFinderSelector = config.NewInjectionSelector("#LocalRepositoryFinder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryImportServiceImpl) newObject() * implservice0x22327c.RepositoryImportServiceImpl {
	return & implservice0x22327c.RepositoryImportServiceImpl {}
}

func (inst * comFactory4pComRepositoryImportServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.RepositoryImportServiceImpl {
	return instance.Get().(*implservice0x22327c.RepositoryImportServiceImpl)
}

func (inst * comFactory4pComRepositoryImportServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositoryImportServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositoryImportServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositoryImportServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryImportServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryImportServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.RepositoryService = inst.getterForFieldRepositoryServiceSelector(context)
	obj.RepoFinder = inst.getterForFieldRepoFinderSelector(context)
	return context.LastError()
}

//getterForFieldRepositoryServiceSelector
func (inst * comFactory4pComRepositoryImportServiceImpl) getterForFieldRepositoryServiceSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryService {

	o1 := inst.mRepositoryServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryImportService")
		eb.Set("field", "RepositoryService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldRepoFinderSelector
func (inst * comFactory4pComRepositoryImportServiceImpl) getterForFieldRepoFinderSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryFinder {

	o1 := inst.mRepoFinderSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryFinder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryImportService")
		eb.Set("field", "RepoFinder")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryFinder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRunIntentServiceImpl : the factory of component: IntentService
type comFactory4pComRunIntentServiceImpl struct {

    mPrototype * implservice0x22327c.RunIntentServiceImpl

	
	mGitLibAgentSelector config.InjectionSelector
	mIntentFilterManagerSelector config.InjectionSelector
	mLocalRepositoryServiceSelector config.InjectionSelector
	mExecutableServiceSelector config.InjectionSelector
	mIntentHandlerServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComRunIntentServiceImpl) init() application.ComponentFactory {

	
	inst.mGitLibAgentSelector = config.NewInjectionSelector("#git-lib-agent",nil)
	inst.mIntentFilterManagerSelector = config.NewInjectionSelector("#intent-filter-manager",nil)
	inst.mLocalRepositoryServiceSelector = config.NewInjectionSelector("#LocalRepositoryService",nil)
	inst.mExecutableServiceSelector = config.NewInjectionSelector("#ExecutableService",nil)
	inst.mIntentHandlerServiceSelector = config.NewInjectionSelector("#IntentHandlerService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRunIntentServiceImpl) newObject() * implservice0x22327c.RunIntentServiceImpl {
	return & implservice0x22327c.RunIntentServiceImpl {}
}

func (inst * comFactory4pComRunIntentServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.RunIntentServiceImpl {
	return instance.Get().(*implservice0x22327c.RunIntentServiceImpl)
}

func (inst * comFactory4pComRunIntentServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRunIntentServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRunIntentServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRunIntentServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunIntentServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunIntentServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.GitLibAgent = inst.getterForFieldGitLibAgentSelector(context)
	obj.IntentFilterManager = inst.getterForFieldIntentFilterManagerSelector(context)
	obj.LocalRepositoryService = inst.getterForFieldLocalRepositoryServiceSelector(context)
	obj.ExecutableService = inst.getterForFieldExecutableServiceSelector(context)
	obj.IntentHandlerService = inst.getterForFieldIntentHandlerServiceSelector(context)
	return context.LastError()
}

//getterForFieldGitLibAgentSelector
func (inst * comFactory4pComRunIntentServiceImpl) getterForFieldGitLibAgentSelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mGitLibAgentSelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentService")
		eb.Set("field", "GitLibAgent")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldIntentFilterManagerSelector
func (inst * comFactory4pComRunIntentServiceImpl) getterForFieldIntentFilterManagerSelector (context application.InstanceContext) intents0xec84e7.FilterManager {

	o1 := inst.mIntentFilterManagerSelector.GetOne(context)
	o2, ok := o1.(intents0xec84e7.FilterManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentService")
		eb.Set("field", "IntentFilterManager")
		eb.Set("type1", "?")
		eb.Set("type2", "intents0xec84e7.FilterManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLocalRepositoryServiceSelector
func (inst * comFactory4pComRunIntentServiceImpl) getterForFieldLocalRepositoryServiceSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryService {

	o1 := inst.mLocalRepositoryServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentService")
		eb.Set("field", "LocalRepositoryService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldExecutableServiceSelector
func (inst * comFactory4pComRunIntentServiceImpl) getterForFieldExecutableServiceSelector (context application.InstanceContext) service0x3e063d.ExecutableService {

	o1 := inst.mExecutableServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ExecutableService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentService")
		eb.Set("field", "ExecutableService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ExecutableService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldIntentHandlerServiceSelector
func (inst * comFactory4pComRunIntentServiceImpl) getterForFieldIntentHandlerServiceSelector (context application.InstanceContext) service0x3e063d.IntentHandlerService {

	o1 := inst.mIntentHandlerServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.IntentHandlerService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentService")
		eb.Set("field", "IntentHandlerService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.IntentHandlerService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUUIDGenServiceImpl : the factory of component: UUIDGenService
type comFactory4pComUUIDGenServiceImpl struct {

    mPrototype * implservice0x22327c.UUIDGenServiceImpl

	

}

func (inst * comFactory4pComUUIDGenServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUUIDGenServiceImpl) newObject() * implservice0x22327c.UUIDGenServiceImpl {
	return & implservice0x22327c.UUIDGenServiceImpl {}
}

func (inst * comFactory4pComUUIDGenServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.UUIDGenServiceImpl {
	return instance.Get().(*implservice0x22327c.UUIDGenServiceImpl)
}

func (inst * comFactory4pComUUIDGenServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComUUIDGenServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComUUIDGenServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComUUIDGenServiceImpl) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComUUIDGenServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUUIDGenServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComWpmDataSource : the factory of component: com27-support0xf47d7f.WpmDataSource
type comFactory4pComWpmDataSource struct {

    mPrototype * support0xf47d7f.WpmDataSource

	
	mDMSelector config.InjectionSelector
	mAppDataServiceSelector config.InjectionSelector
	mDriverSelector config.InjectionSelector
	mHostSelector config.InjectionSelector
	mPortSelector config.InjectionSelector
	mUserNameSelector config.InjectionSelector
	mPasswordSelector config.InjectionSelector
	mDatabaseSelector config.InjectionSelector

}

func (inst * comFactory4pComWpmDataSource) init() application.ComponentFactory {

	
	inst.mDMSelector = config.NewInjectionSelector("#starter-gorm-driver-manager",nil)
	inst.mAppDataServiceSelector = config.NewInjectionSelector("#AppDataService",nil)
	inst.mDriverSelector = config.NewInjectionSelector("${datasource.wpm.driver}",nil)
	inst.mHostSelector = config.NewInjectionSelector("${datasource.wpm.host}",nil)
	inst.mPortSelector = config.NewInjectionSelector("${datasource.wpm.port}",nil)
	inst.mUserNameSelector = config.NewInjectionSelector("${datasource.wpm.username}",nil)
	inst.mPasswordSelector = config.NewInjectionSelector("${datasource.wpm.password}",nil)
	inst.mDatabaseSelector = config.NewInjectionSelector("${datasource.wpm.database}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComWpmDataSource) newObject() * support0xf47d7f.WpmDataSource {
	return & support0xf47d7f.WpmDataSource {}
}

func (inst * comFactory4pComWpmDataSource) castObject(instance application.ComponentInstance) * support0xf47d7f.WpmDataSource {
	return instance.Get().(*support0xf47d7f.WpmDataSource)
}

func (inst * comFactory4pComWpmDataSource) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComWpmDataSource) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComWpmDataSource) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComWpmDataSource) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWpmDataSource) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWpmDataSource) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.DM = inst.getterForFieldDMSelector(context)
	obj.AppDataService = inst.getterForFieldAppDataServiceSelector(context)
	obj.Driver = inst.getterForFieldDriverSelector(context)
	obj.Host = inst.getterForFieldHostSelector(context)
	obj.Port = inst.getterForFieldPortSelector(context)
	obj.UserName = inst.getterForFieldUserNameSelector(context)
	obj.Password = inst.getterForFieldPasswordSelector(context)
	obj.Database = inst.getterForFieldDatabaseSelector(context)
	return context.LastError()
}

//getterForFieldDMSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldDMSelector (context application.InstanceContext) datasource0x68a737.DriverManager {

	o1 := inst.mDMSelector.GetOne(context)
	o2, ok := o1.(datasource0x68a737.DriverManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com27-support0xf47d7f.WpmDataSource")
		eb.Set("field", "DM")
		eb.Set("type1", "?")
		eb.Set("type2", "datasource0x68a737.DriverManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldAppDataServiceSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldAppDataServiceSelector (context application.InstanceContext) service0x3e063d.AppDataService {

	o1 := inst.mAppDataServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AppDataService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com27-support0xf47d7f.WpmDataSource")
		eb.Set("field", "AppDataService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AppDataService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldDriverSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldDriverSelector (context application.InstanceContext) string {
    return inst.mDriverSelector.GetString(context)
}

//getterForFieldHostSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldHostSelector (context application.InstanceContext) string {
    return inst.mHostSelector.GetString(context)
}

//getterForFieldPortSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldPortSelector (context application.InstanceContext) int {
    return inst.mPortSelector.GetInt(context)
}

//getterForFieldUserNameSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldUserNameSelector (context application.InstanceContext) string {
    return inst.mUserNameSelector.GetString(context)
}

//getterForFieldPasswordSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldPasswordSelector (context application.InstanceContext) string {
    return inst.mPasswordSelector.GetString(context)
}

//getterForFieldDatabaseSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldDatabaseSelector (context application.InstanceContext) string {
    return inst.mDatabaseSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecuteIntentFilter : the factory of component: com28-filters0x5d53d8.ExecuteIntentFilter
type comFactory4pComExecuteIntentFilter struct {

    mPrototype * filters0x5d53d8.ExecuteIntentFilter

	
	mGitLibAgentSelector config.InjectionSelector

}

func (inst * comFactory4pComExecuteIntentFilter) init() application.ComponentFactory {

	
	inst.mGitLibAgentSelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecuteIntentFilter) newObject() * filters0x5d53d8.ExecuteIntentFilter {
	return & filters0x5d53d8.ExecuteIntentFilter {}
}

func (inst * comFactory4pComExecuteIntentFilter) castObject(instance application.ComponentInstance) * filters0x5d53d8.ExecuteIntentFilter {
	return instance.Get().(*filters0x5d53d8.ExecuteIntentFilter)
}

func (inst * comFactory4pComExecuteIntentFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExecuteIntentFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExecuteIntentFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExecuteIntentFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecuteIntentFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecuteIntentFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.GitLibAgent = inst.getterForFieldGitLibAgentSelector(context)
	return context.LastError()
}

//getterForFieldGitLibAgentSelector
func (inst * comFactory4pComExecuteIntentFilter) getterForFieldGitLibAgentSelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mGitLibAgentSelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com28-filters0x5d53d8.ExecuteIntentFilter")
		eb.Set("field", "GitLibAgent")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentFilterFor : the factory of component: com29-filters0x5d53d8.IntentFilterFor
type comFactory4pComIntentFilterFor struct {

    mPrototype * filters0x5d53d8.IntentFilterFor

	

}

func (inst * comFactory4pComIntentFilterFor) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentFilterFor) newObject() * filters0x5d53d8.IntentFilterFor {
	return & filters0x5d53d8.IntentFilterFor {}
}

func (inst * comFactory4pComIntentFilterFor) castObject(instance application.ComponentInstance) * filters0x5d53d8.IntentFilterFor {
	return instance.Get().(*filters0x5d53d8.IntentFilterFor)
}

func (inst * comFactory4pComIntentFilterFor) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentFilterFor) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentFilterFor) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentFilterFor) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterFor) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterFor) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentFilterForChrome : the factory of component: com30-filters0x5d53d8.IntentFilterForChrome
type comFactory4pComIntentFilterForChrome struct {

    mPrototype * filters0x5d53d8.IntentFilterForChrome

	

}

func (inst * comFactory4pComIntentFilterForChrome) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentFilterForChrome) newObject() * filters0x5d53d8.IntentFilterForChrome {
	return & filters0x5d53d8.IntentFilterForChrome {}
}

func (inst * comFactory4pComIntentFilterForChrome) castObject(instance application.ComponentInstance) * filters0x5d53d8.IntentFilterForChrome {
	return instance.Get().(*filters0x5d53d8.IntentFilterForChrome)
}

func (inst * comFactory4pComIntentFilterForChrome) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentFilterForChrome) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentFilterForChrome) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentFilterForChrome) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForChrome) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForChrome) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentFilterForCmd : the factory of component: com31-filters0x5d53d8.IntentFilterForCmd
type comFactory4pComIntentFilterForCmd struct {

    mPrototype * filters0x5d53d8.IntentFilterForCmd

	

}

func (inst * comFactory4pComIntentFilterForCmd) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentFilterForCmd) newObject() * filters0x5d53d8.IntentFilterForCmd {
	return & filters0x5d53d8.IntentFilterForCmd {}
}

func (inst * comFactory4pComIntentFilterForCmd) castObject(instance application.ComponentInstance) * filters0x5d53d8.IntentFilterForCmd {
	return instance.Get().(*filters0x5d53d8.IntentFilterForCmd)
}

func (inst * comFactory4pComIntentFilterForCmd) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentFilterForCmd) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentFilterForCmd) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentFilterForCmd) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForCmd) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForCmd) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentFilterForExplorer : the factory of component: com32-filters0x5d53d8.IntentFilterForExplorer
type comFactory4pComIntentFilterForExplorer struct {

    mPrototype * filters0x5d53d8.IntentFilterForExplorer

	

}

func (inst * comFactory4pComIntentFilterForExplorer) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentFilterForExplorer) newObject() * filters0x5d53d8.IntentFilterForExplorer {
	return & filters0x5d53d8.IntentFilterForExplorer {}
}

func (inst * comFactory4pComIntentFilterForExplorer) castObject(instance application.ComponentInstance) * filters0x5d53d8.IntentFilterForExplorer {
	return instance.Get().(*filters0x5d53d8.IntentFilterForExplorer)
}

func (inst * comFactory4pComIntentFilterForExplorer) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentFilterForExplorer) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentFilterForExplorer) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentFilterForExplorer) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForExplorer) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForExplorer) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentFilterForFirefox : the factory of component: com33-filters0x5d53d8.IntentFilterForFirefox
type comFactory4pComIntentFilterForFirefox struct {

    mPrototype * filters0x5d53d8.IntentFilterForFirefox

	

}

func (inst * comFactory4pComIntentFilterForFirefox) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentFilterForFirefox) newObject() * filters0x5d53d8.IntentFilterForFirefox {
	return & filters0x5d53d8.IntentFilterForFirefox {}
}

func (inst * comFactory4pComIntentFilterForFirefox) castObject(instance application.ComponentInstance) * filters0x5d53d8.IntentFilterForFirefox {
	return instance.Get().(*filters0x5d53d8.IntentFilterForFirefox)
}

func (inst * comFactory4pComIntentFilterForFirefox) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentFilterForFirefox) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentFilterForFirefox) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentFilterForFirefox) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForFirefox) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForFirefox) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentFilterForMsEdge : the factory of component: com34-filters0x5d53d8.IntentFilterForMsEdge
type comFactory4pComIntentFilterForMsEdge struct {

    mPrototype * filters0x5d53d8.IntentFilterForMsEdge

	

}

func (inst * comFactory4pComIntentFilterForMsEdge) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentFilterForMsEdge) newObject() * filters0x5d53d8.IntentFilterForMsEdge {
	return & filters0x5d53d8.IntentFilterForMsEdge {}
}

func (inst * comFactory4pComIntentFilterForMsEdge) castObject(instance application.ComponentInstance) * filters0x5d53d8.IntentFilterForMsEdge {
	return instance.Get().(*filters0x5d53d8.IntentFilterForMsEdge)
}

func (inst * comFactory4pComIntentFilterForMsEdge) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentFilterForMsEdge) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentFilterForMsEdge) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentFilterForMsEdge) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForMsEdge) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForMsEdge) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentFilterForNautilus : the factory of component: com35-filters0x5d53d8.IntentFilterForNautilus
type comFactory4pComIntentFilterForNautilus struct {

    mPrototype * filters0x5d53d8.IntentFilterForNautilus

	

}

func (inst * comFactory4pComIntentFilterForNautilus) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentFilterForNautilus) newObject() * filters0x5d53d8.IntentFilterForNautilus {
	return & filters0x5d53d8.IntentFilterForNautilus {}
}

func (inst * comFactory4pComIntentFilterForNautilus) castObject(instance application.ComponentInstance) * filters0x5d53d8.IntentFilterForNautilus {
	return instance.Get().(*filters0x5d53d8.IntentFilterForNautilus)
}

func (inst * comFactory4pComIntentFilterForNautilus) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentFilterForNautilus) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentFilterForNautilus) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentFilterForNautilus) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForNautilus) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForNautilus) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentFilterForPowerShell : the factory of component: com36-filters0x5d53d8.IntentFilterForPowerShell
type comFactory4pComIntentFilterForPowerShell struct {

    mPrototype * filters0x5d53d8.IntentFilterForPowerShell

	

}

func (inst * comFactory4pComIntentFilterForPowerShell) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentFilterForPowerShell) newObject() * filters0x5d53d8.IntentFilterForPowerShell {
	return & filters0x5d53d8.IntentFilterForPowerShell {}
}

func (inst * comFactory4pComIntentFilterForPowerShell) castObject(instance application.ComponentInstance) * filters0x5d53d8.IntentFilterForPowerShell {
	return instance.Get().(*filters0x5d53d8.IntentFilterForPowerShell)
}

func (inst * comFactory4pComIntentFilterForPowerShell) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentFilterForPowerShell) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentFilterForPowerShell) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentFilterForPowerShell) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForPowerShell) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForPowerShell) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentFilterForVscode : the factory of component: com37-filters0x5d53d8.IntentFilterForVscode
type comFactory4pComIntentFilterForVscode struct {

    mPrototype * filters0x5d53d8.IntentFilterForVscode

	

}

func (inst * comFactory4pComIntentFilterForVscode) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentFilterForVscode) newObject() * filters0x5d53d8.IntentFilterForVscode {
	return & filters0x5d53d8.IntentFilterForVscode {}
}

func (inst * comFactory4pComIntentFilterForVscode) castObject(instance application.ComponentInstance) * filters0x5d53d8.IntentFilterForVscode {
	return instance.Get().(*filters0x5d53d8.IntentFilterForVscode)
}

func (inst * comFactory4pComIntentFilterForVscode) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentFilterForVscode) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentFilterForVscode) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentFilterForVscode) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForVscode) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentFilterForVscode) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFilterManagerImpl : the factory of component: intent-filter-manager
type comFactory4pComFilterManagerImpl struct {

    mPrototype * intents0xec84e7.FilterManagerImpl

	
	mFiltersSelector config.InjectionSelector

}

func (inst * comFactory4pComFilterManagerImpl) init() application.ComponentFactory {

	
	inst.mFiltersSelector = config.NewInjectionSelector(".intent-filter-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFilterManagerImpl) newObject() * intents0xec84e7.FilterManagerImpl {
	return & intents0xec84e7.FilterManagerImpl {}
}

func (inst * comFactory4pComFilterManagerImpl) castObject(instance application.ComponentInstance) * intents0xec84e7.FilterManagerImpl {
	return instance.Get().(*intents0xec84e7.FilterManagerImpl)
}

func (inst * comFactory4pComFilterManagerImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFilterManagerImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFilterManagerImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFilterManagerImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFilterManagerImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFilterManagerImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Filters = inst.getterForFieldFiltersSelector(context)
	return context.LastError()
}

//getterForFieldFiltersSelector
func (inst * comFactory4pComFilterManagerImpl) getterForFieldFiltersSelector (context application.InstanceContext) []intents0xec84e7.FilterRegistry {
	list1 := inst.mFiltersSelector.GetList(context)
	list2 := make([]intents0xec84e7.FilterRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(intents0xec84e7.FilterRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAboutController : the factory of component: com39-controller0x9dc399.AboutController
type comFactory4pComAboutController struct {

    mPrototype * controller0x9dc399.AboutController

	
	mAboutServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComAboutController) init() application.ComponentFactory {

	
	inst.mAboutServiceSelector = config.NewInjectionSelector("#AboutService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAboutController) newObject() * controller0x9dc399.AboutController {
	return & controller0x9dc399.AboutController {}
}

func (inst * comFactory4pComAboutController) castObject(instance application.ComponentInstance) * controller0x9dc399.AboutController {
	return instance.Get().(*controller0x9dc399.AboutController)
}

func (inst * comFactory4pComAboutController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAboutController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAboutController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAboutController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAboutController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAboutController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AboutService = inst.getterForFieldAboutServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldAboutServiceSelector
func (inst * comFactory4pComAboutController) getterForFieldAboutServiceSelector (context application.InstanceContext) service0x3e063d.AboutService {

	o1 := inst.mAboutServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AboutService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com39-controller0x9dc399.AboutController")
		eb.Set("field", "AboutService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AboutService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComAboutController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com39-controller0x9dc399.AboutController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAuthController : the factory of component: com40-controller0x9dc399.AuthController
type comFactory4pComAuthController struct {

    mPrototype * controller0x9dc399.AuthController

	
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComAuthController) init() application.ComponentFactory {

	
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAuthController) newObject() * controller0x9dc399.AuthController {
	return & controller0x9dc399.AuthController {}
}

func (inst * comFactory4pComAuthController) castObject(instance application.ComponentInstance) * controller0x9dc399.AuthController {
	return instance.Get().(*controller0x9dc399.AuthController)
}

func (inst * comFactory4pComAuthController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAuthController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAuthController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAuthController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAuthController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAuthController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldResponderSelector
func (inst * comFactory4pComAuthController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com40-controller0x9dc399.AuthController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExampleController : the factory of component: com41-controller0x9dc399.ExampleController
type comFactory4pComExampleController struct {

    mPrototype * controller0x9dc399.ExampleController

	
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComExampleController) init() application.ComponentFactory {

	
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExampleController) newObject() * controller0x9dc399.ExampleController {
	return & controller0x9dc399.ExampleController {}
}

func (inst * comFactory4pComExampleController) castObject(instance application.ComponentInstance) * controller0x9dc399.ExampleController {
	return instance.Get().(*controller0x9dc399.ExampleController)
}

func (inst * comFactory4pComExampleController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExampleController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExampleController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExampleController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldResponderSelector
func (inst * comFactory4pComExampleController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com41-controller0x9dc399.ExampleController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableController : the factory of component: com42-controller0x9dc399.ExecutableController
type comFactory4pComExecutableController struct {

    mPrototype * controller0x9dc399.ExecutableController

	
	mExecutableServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableController) init() application.ComponentFactory {

	
	inst.mExecutableServiceSelector = config.NewInjectionSelector("#ExecutableService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableController) newObject() * controller0x9dc399.ExecutableController {
	return & controller0x9dc399.ExecutableController {}
}

func (inst * comFactory4pComExecutableController) castObject(instance application.ComponentInstance) * controller0x9dc399.ExecutableController {
	return instance.Get().(*controller0x9dc399.ExecutableController)
}

func (inst * comFactory4pComExecutableController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExecutableController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExecutableController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExecutableController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ExecutableService = inst.getterForFieldExecutableServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldExecutableServiceSelector
func (inst * comFactory4pComExecutableController) getterForFieldExecutableServiceSelector (context application.InstanceContext) service0x3e063d.ExecutableService {

	o1 := inst.mExecutableServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ExecutableService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com42-controller0x9dc399.ExecutableController")
		eb.Set("field", "ExecutableService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ExecutableService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComExecutableController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com42-controller0x9dc399.ExecutableController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableImportController : the factory of component: com43-controller0x9dc399.ExecutableImportController
type comFactory4pComExecutableImportController struct {

    mPrototype * controller0x9dc399.ExecutableImportController

	
	mExecutableImportServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableImportController) init() application.ComponentFactory {

	
	inst.mExecutableImportServiceSelector = config.NewInjectionSelector("#ExecutableImportService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableImportController) newObject() * controller0x9dc399.ExecutableImportController {
	return & controller0x9dc399.ExecutableImportController {}
}

func (inst * comFactory4pComExecutableImportController) castObject(instance application.ComponentInstance) * controller0x9dc399.ExecutableImportController {
	return instance.Get().(*controller0x9dc399.ExecutableImportController)
}

func (inst * comFactory4pComExecutableImportController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExecutableImportController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExecutableImportController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExecutableImportController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableImportController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableImportController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ExecutableImportService = inst.getterForFieldExecutableImportServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldExecutableImportServiceSelector
func (inst * comFactory4pComExecutableImportController) getterForFieldExecutableImportServiceSelector (context application.InstanceContext) service0x3e063d.ExecutableImportService {

	o1 := inst.mExecutableImportServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ExecutableImportService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com43-controller0x9dc399.ExecutableImportController")
		eb.Set("field", "ExecutableImportService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ExecutableImportService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComExecutableImportController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com43-controller0x9dc399.ExecutableImportController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFileQueryController : the factory of component: com44-controller0x9dc399.FileQueryController
type comFactory4pComFileQueryController struct {

    mPrototype * controller0x9dc399.FileQueryController

	
	mFileQueryServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComFileQueryController) init() application.ComponentFactory {

	
	inst.mFileQueryServiceSelector = config.NewInjectionSelector("#FileQueryService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFileQueryController) newObject() * controller0x9dc399.FileQueryController {
	return & controller0x9dc399.FileQueryController {}
}

func (inst * comFactory4pComFileQueryController) castObject(instance application.ComponentInstance) * controller0x9dc399.FileQueryController {
	return instance.Get().(*controller0x9dc399.FileQueryController)
}

func (inst * comFactory4pComFileQueryController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFileQueryController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFileQueryController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFileQueryController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileQueryController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileQueryController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.FileQueryService = inst.getterForFieldFileQueryServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldFileQueryServiceSelector
func (inst * comFactory4pComFileQueryController) getterForFieldFileQueryServiceSelector (context application.InstanceContext) service0x3e063d.FileQueryService {

	o1 := inst.mFileQueryServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.FileQueryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com44-controller0x9dc399.FileQueryController")
		eb.Set("field", "FileQueryService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.FileQueryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComFileQueryController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com44-controller0x9dc399.FileQueryController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRunIntentController : the factory of component: com45-controller0x9dc399.RunIntentController
type comFactory4pComRunIntentController struct {

    mPrototype * controller0x9dc399.RunIntentController

	
	mIntentServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComRunIntentController) init() application.ComponentFactory {

	
	inst.mIntentServiceSelector = config.NewInjectionSelector("#IntentService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRunIntentController) newObject() * controller0x9dc399.RunIntentController {
	return & controller0x9dc399.RunIntentController {}
}

func (inst * comFactory4pComRunIntentController) castObject(instance application.ComponentInstance) * controller0x9dc399.RunIntentController {
	return instance.Get().(*controller0x9dc399.RunIntentController)
}

func (inst * comFactory4pComRunIntentController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRunIntentController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRunIntentController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRunIntentController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunIntentController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunIntentController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.IntentService = inst.getterForFieldIntentServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldIntentServiceSelector
func (inst * comFactory4pComRunIntentController) getterForFieldIntentServiceSelector (context application.InstanceContext) service0x3e063d.IntentService {

	o1 := inst.mIntentServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.IntentService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com45-controller0x9dc399.RunIntentController")
		eb.Set("field", "IntentService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.IntentService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComRunIntentController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com45-controller0x9dc399.RunIntentController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentTemplateController : the factory of component: com46-controller0x9dc399.IntentTemplateController
type comFactory4pComIntentTemplateController struct {

    mPrototype * controller0x9dc399.IntentTemplateController

	
	mIntentTemplateServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComIntentTemplateController) init() application.ComponentFactory {

	
	inst.mIntentTemplateServiceSelector = config.NewInjectionSelector("#IntentTemplateService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentTemplateController) newObject() * controller0x9dc399.IntentTemplateController {
	return & controller0x9dc399.IntentTemplateController {}
}

func (inst * comFactory4pComIntentTemplateController) castObject(instance application.ComponentInstance) * controller0x9dc399.IntentTemplateController {
	return instance.Get().(*controller0x9dc399.IntentTemplateController)
}

func (inst * comFactory4pComIntentTemplateController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentTemplateController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentTemplateController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentTemplateController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentTemplateController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentTemplateController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.IntentTemplateService = inst.getterForFieldIntentTemplateServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldIntentTemplateServiceSelector
func (inst * comFactory4pComIntentTemplateController) getterForFieldIntentTemplateServiceSelector (context application.InstanceContext) service0x3e063d.IntentTemplateService {

	o1 := inst.mIntentTemplateServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.IntentTemplateService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com46-controller0x9dc399.IntentTemplateController")
		eb.Set("field", "IntentTemplateService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.IntentTemplateService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComIntentTemplateController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com46-controller0x9dc399.IntentTemplateController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalRepositoryController : the factory of component: com47-controller0x9dc399.LocalRepositoryController
type comFactory4pComLocalRepositoryController struct {

    mPrototype * controller0x9dc399.LocalRepositoryController

	
	mRepoServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComLocalRepositoryController) init() application.ComponentFactory {

	
	inst.mRepoServiceSelector = config.NewInjectionSelector("#LocalRepositoryService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalRepositoryController) newObject() * controller0x9dc399.LocalRepositoryController {
	return & controller0x9dc399.LocalRepositoryController {}
}

func (inst * comFactory4pComLocalRepositoryController) castObject(instance application.ComponentInstance) * controller0x9dc399.LocalRepositoryController {
	return instance.Get().(*controller0x9dc399.LocalRepositoryController)
}

func (inst * comFactory4pComLocalRepositoryController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalRepositoryController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalRepositoryController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalRepositoryController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.RepoService = inst.getterForFieldRepoServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldRepoServiceSelector
func (inst * comFactory4pComLocalRepositoryController) getterForFieldRepoServiceSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryService {

	o1 := inst.mRepoServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com47-controller0x9dc399.LocalRepositoryController")
		eb.Set("field", "RepoService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComLocalRepositoryController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com47-controller0x9dc399.LocalRepositoryController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMediaController : the factory of component: com48-controller0x9dc399.MediaController
type comFactory4pComMediaController struct {

    mPrototype * controller0x9dc399.MediaController

	
	mMediaServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComMediaController) init() application.ComponentFactory {

	
	inst.mMediaServiceSelector = config.NewInjectionSelector("#MediaService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMediaController) newObject() * controller0x9dc399.MediaController {
	return & controller0x9dc399.MediaController {}
}

func (inst * comFactory4pComMediaController) castObject(instance application.ComponentInstance) * controller0x9dc399.MediaController {
	return instance.Get().(*controller0x9dc399.MediaController)
}

func (inst * comFactory4pComMediaController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMediaController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMediaController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMediaController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMediaController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMediaController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.MediaService = inst.getterForFieldMediaServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldMediaServiceSelector
func (inst * comFactory4pComMediaController) getterForFieldMediaServiceSelector (context application.InstanceContext) service0x3e063d.MediaService {

	o1 := inst.mMediaServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.MediaService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com48-controller0x9dc399.MediaController")
		eb.Set("field", "MediaService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.MediaService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComMediaController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com48-controller0x9dc399.MediaController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPlatformController : the factory of component: com49-controller0x9dc399.PlatformController
type comFactory4pComPlatformController struct {

    mPrototype * controller0x9dc399.PlatformController

	
	mPlatformServiceSelector config.InjectionSelector
	mProfileServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComPlatformController) init() application.ComponentFactory {

	
	inst.mPlatformServiceSelector = config.NewInjectionSelector("#PlatformService",nil)
	inst.mProfileServiceSelector = config.NewInjectionSelector("#ProfileService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPlatformController) newObject() * controller0x9dc399.PlatformController {
	return & controller0x9dc399.PlatformController {}
}

func (inst * comFactory4pComPlatformController) castObject(instance application.ComponentInstance) * controller0x9dc399.PlatformController {
	return instance.Get().(*controller0x9dc399.PlatformController)
}

func (inst * comFactory4pComPlatformController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPlatformController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPlatformController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPlatformController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPlatformController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPlatformController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.PlatformService = inst.getterForFieldPlatformServiceSelector(context)
	obj.ProfileService = inst.getterForFieldProfileServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldPlatformServiceSelector
func (inst * comFactory4pComPlatformController) getterForFieldPlatformServiceSelector (context application.InstanceContext) service0x3e063d.PlatformService {

	o1 := inst.mPlatformServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.PlatformService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com49-controller0x9dc399.PlatformController")
		eb.Set("field", "PlatformService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.PlatformService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProfileServiceSelector
func (inst * comFactory4pComPlatformController) getterForFieldProfileServiceSelector (context application.InstanceContext) service0x3e063d.ProfileService {

	o1 := inst.mProfileServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProfileService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com49-controller0x9dc399.PlatformController")
		eb.Set("field", "ProfileService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProfileService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComPlatformController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com49-controller0x9dc399.PlatformController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectController : the factory of component: com50-controller0x9dc399.ProjectController
type comFactory4pComProjectController struct {

    mPrototype * controller0x9dc399.ProjectController

	
	mProjectServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectController) init() application.ComponentFactory {

	
	inst.mProjectServiceSelector = config.NewInjectionSelector("#ProjectService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectController) newObject() * controller0x9dc399.ProjectController {
	return & controller0x9dc399.ProjectController {}
}

func (inst * comFactory4pComProjectController) castObject(instance application.ComponentInstance) * controller0x9dc399.ProjectController {
	return instance.Get().(*controller0x9dc399.ProjectController)
}

func (inst * comFactory4pComProjectController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ProjectService = inst.getterForFieldProjectServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldProjectServiceSelector
func (inst * comFactory4pComProjectController) getterForFieldProjectServiceSelector (context application.InstanceContext) service0x3e063d.ProjectService {

	o1 := inst.mProjectServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProjectService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com50-controller0x9dc399.ProjectController")
		eb.Set("field", "ProjectService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProjectService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComProjectController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com50-controller0x9dc399.ProjectController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectImportController : the factory of component: com51-controller0x9dc399.ProjectImportController
type comFactory4pComProjectImportController struct {

    mPrototype * controller0x9dc399.ProjectImportController

	
	mProjectImportServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectImportController) init() application.ComponentFactory {

	
	inst.mProjectImportServiceSelector = config.NewInjectionSelector("#ProjectImportService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectImportController) newObject() * controller0x9dc399.ProjectImportController {
	return & controller0x9dc399.ProjectImportController {}
}

func (inst * comFactory4pComProjectImportController) castObject(instance application.ComponentInstance) * controller0x9dc399.ProjectImportController {
	return instance.Get().(*controller0x9dc399.ProjectImportController)
}

func (inst * comFactory4pComProjectImportController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectImportController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectImportController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectImportController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectImportController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectImportController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ProjectImportService = inst.getterForFieldProjectImportServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldProjectImportServiceSelector
func (inst * comFactory4pComProjectImportController) getterForFieldProjectImportServiceSelector (context application.InstanceContext) service0x3e063d.ProjectImportService {

	o1 := inst.mProjectImportServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProjectImportService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com51-controller0x9dc399.ProjectImportController")
		eb.Set("field", "ProjectImportService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProjectImportService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComProjectImportController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com51-controller0x9dc399.ProjectImportController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRemoteRepositoryController : the factory of component: com52-controller0x9dc399.RemoteRepositoryController
type comFactory4pComRemoteRepositoryController struct {

    mPrototype * controller0x9dc399.RemoteRepositoryController

	
	mRepoServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComRemoteRepositoryController) init() application.ComponentFactory {

	
	inst.mRepoServiceSelector = config.NewInjectionSelector("#RemoteRepositoryService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRemoteRepositoryController) newObject() * controller0x9dc399.RemoteRepositoryController {
	return & controller0x9dc399.RemoteRepositoryController {}
}

func (inst * comFactory4pComRemoteRepositoryController) castObject(instance application.ComponentInstance) * controller0x9dc399.RemoteRepositoryController {
	return instance.Get().(*controller0x9dc399.RemoteRepositoryController)
}

func (inst * comFactory4pComRemoteRepositoryController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRemoteRepositoryController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRemoteRepositoryController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRemoteRepositoryController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRemoteRepositoryController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRemoteRepositoryController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.RepoService = inst.getterForFieldRepoServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldRepoServiceSelector
func (inst * comFactory4pComRemoteRepositoryController) getterForFieldRepoServiceSelector (context application.InstanceContext) service0x3e063d.RemoteRepositoryService {

	o1 := inst.mRepoServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.RemoteRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com52-controller0x9dc399.RemoteRepositoryController")
		eb.Set("field", "RepoService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.RemoteRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComRemoteRepositoryController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com52-controller0x9dc399.RemoteRepositoryController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryImportController : the factory of component: com53-controller0x9dc399.RepositoryImportController
type comFactory4pComRepositoryImportController struct {

    mPrototype * controller0x9dc399.RepositoryImportController

	
	mRepoStateLoaderSelector config.InjectionSelector
	mImportServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryImportController) init() application.ComponentFactory {

	
	inst.mRepoStateLoaderSelector = config.NewInjectionSelector("#LocalRepositoryStateLoader",nil)
	inst.mImportServiceSelector = config.NewInjectionSelector("#RepositoryImportService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryImportController) newObject() * controller0x9dc399.RepositoryImportController {
	return & controller0x9dc399.RepositoryImportController {}
}

func (inst * comFactory4pComRepositoryImportController) castObject(instance application.ComponentInstance) * controller0x9dc399.RepositoryImportController {
	return instance.Get().(*controller0x9dc399.RepositoryImportController)
}

func (inst * comFactory4pComRepositoryImportController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositoryImportController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositoryImportController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositoryImportController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryImportController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryImportController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.RepoStateLoader = inst.getterForFieldRepoStateLoaderSelector(context)
	obj.ImportService = inst.getterForFieldImportServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldRepoStateLoaderSelector
func (inst * comFactory4pComRepositoryImportController) getterForFieldRepoStateLoaderSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryStateLoader {

	o1 := inst.mRepoStateLoaderSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryStateLoader)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com53-controller0x9dc399.RepositoryImportController")
		eb.Set("field", "RepoStateLoader")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryStateLoader")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldImportServiceSelector
func (inst * comFactory4pComRepositoryImportController) getterForFieldImportServiceSelector (context application.InstanceContext) service0x3e063d.RepositoryImportService {

	o1 := inst.mImportServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.RepositoryImportService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com53-controller0x9dc399.RepositoryImportController")
		eb.Set("field", "ImportService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.RepositoryImportService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComRepositoryImportController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com53-controller0x9dc399.RepositoryImportController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHostFilter : the factory of component: com54-filter0x8aa8f6.HostFilter
type comFactory4pComHostFilter struct {

    mPrototype * filter0x8aa8f6.HostFilter

	

}

func (inst * comFactory4pComHostFilter) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComHostFilter) newObject() * filter0x8aa8f6.HostFilter {
	return & filter0x8aa8f6.HostFilter {}
}

func (inst * comFactory4pComHostFilter) castObject(instance application.ComponentInstance) * filter0x8aa8f6.HostFilter {
	return instance.Get().(*filter0x8aa8f6.HostFilter)
}

func (inst * comFactory4pComHostFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComHostFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComHostFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComHostFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHostFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHostFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHTTP404Filter : the factory of component: com55-filter0x8aa8f6.HTTP404Filter
type comFactory4pComHTTP404Filter struct {

    mPrototype * filter0x8aa8f6.HTTP404Filter

	
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComHTTP404Filter) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComHTTP404Filter) newObject() * filter0x8aa8f6.HTTP404Filter {
	return & filter0x8aa8f6.HTTP404Filter {}
}

func (inst * comFactory4pComHTTP404Filter) castObject(instance application.ComponentInstance) * filter0x8aa8f6.HTTP404Filter {
	return instance.Get().(*filter0x8aa8f6.HTTP404Filter)
}

func (inst * comFactory4pComHTTP404Filter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComHTTP404Filter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComHTTP404Filter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComHTTP404Filter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTP404Filter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTP404Filter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComHTTP404Filter) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPlatformServiceImpl : the factory of component: PlatformService
type comFactory4pComPlatformServiceImpl struct {

    mPrototype * implservice0x5a8f41.PlatformServiceImpl

	
	mProvidersSelector config.InjectionSelector

}

func (inst * comFactory4pComPlatformServiceImpl) init() application.ComponentFactory {

	
	inst.mProvidersSelector = config.NewInjectionSelector(".PlatformProviderRegistry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPlatformServiceImpl) newObject() * implservice0x5a8f41.PlatformServiceImpl {
	return & implservice0x5a8f41.PlatformServiceImpl {}
}

func (inst * comFactory4pComPlatformServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5a8f41.PlatformServiceImpl {
	return instance.Get().(*implservice0x5a8f41.PlatformServiceImpl)
}

func (inst * comFactory4pComPlatformServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPlatformServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPlatformServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPlatformServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPlatformServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPlatformServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Providers = inst.getterForFieldProvidersSelector(context)
	return context.LastError()
}

//getterForFieldProvidersSelector
func (inst * comFactory4pComPlatformServiceImpl) getterForFieldProvidersSelector (context application.InstanceContext) []service0x3e063d.PlatformProviderRegistry {
	list1 := inst.mProvidersSelector.GetList(context)
	list2 := make([]service0x3e063d.PlatformProviderRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(service0x3e063d.PlatformProviderRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLinuxPlatformServiceImpl : the factory of component: com57-implservice0x5a8f41.LinuxPlatformServiceImpl
type comFactory4pComLinuxPlatformServiceImpl struct {

    mPrototype * implservice0x5a8f41.LinuxPlatformServiceImpl

	

}

func (inst * comFactory4pComLinuxPlatformServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) newObject() * implservice0x5a8f41.LinuxPlatformServiceImpl {
	return & implservice0x5a8f41.LinuxPlatformServiceImpl {}
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5a8f41.LinuxPlatformServiceImpl {
	return instance.Get().(*implservice0x5a8f41.LinuxPlatformServiceImpl)
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComWindowsPlatformServiceImpl : the factory of component: com58-implservice0x5a8f41.WindowsPlatformServiceImpl
type comFactory4pComWindowsPlatformServiceImpl struct {

    mPrototype * implservice0x5a8f41.WindowsPlatformServiceImpl

	

}

func (inst * comFactory4pComWindowsPlatformServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) newObject() * implservice0x5a8f41.WindowsPlatformServiceImpl {
	return & implservice0x5a8f41.WindowsPlatformServiceImpl {}
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5a8f41.WindowsPlatformServiceImpl {
	return instance.Get().(*implservice0x5a8f41.WindowsPlatformServiceImpl)
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProfileServiceImpl : the factory of component: ProfileService
type comFactory4pComProfileServiceImpl struct {

    mPrototype * implservice0x5a8f41.ProfileServiceImpl

	
	mPlatformServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComProfileServiceImpl) init() application.ComponentFactory {

	
	inst.mPlatformServiceSelector = config.NewInjectionSelector("#PlatformService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProfileServiceImpl) newObject() * implservice0x5a8f41.ProfileServiceImpl {
	return & implservice0x5a8f41.ProfileServiceImpl {}
}

func (inst * comFactory4pComProfileServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5a8f41.ProfileServiceImpl {
	return instance.Get().(*implservice0x5a8f41.ProfileServiceImpl)
}

func (inst * comFactory4pComProfileServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProfileServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProfileServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProfileServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProfileServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProfileServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.PlatformService = inst.getterForFieldPlatformServiceSelector(context)
	return context.LastError()
}

//getterForFieldPlatformServiceSelector
func (inst * comFactory4pComProfileServiceImpl) getterForFieldPlatformServiceSelector (context application.InstanceContext) service0x3e063d.PlatformService {

	o1 := inst.mPlatformServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.PlatformService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProfileService")
		eb.Set("field", "PlatformService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.PlatformService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}




