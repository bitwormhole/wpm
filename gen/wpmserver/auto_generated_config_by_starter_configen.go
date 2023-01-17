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
	dao0x5af8d0 "github.com/bitwormhole/wpm/server/data/dao"
	service0x3e063d "github.com/bitwormhole/wpm/server/service"
	impldao0x73998b "github.com/bitwormhole/wpm/server/support/impldao"
	implservice0x22327c "github.com/bitwormhole/wpm/server/support/implservice"
	controller0x9dc399 "github.com/bitwormhole/wpm/server/web/controller"
    
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

	// component: RepositoryService
	cominfobuilder.Next()
	cominfobuilder.ID("RepositoryService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryServiceImpl{}).init())
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

	// component: SystemMainRepositoryService
	cominfobuilder.Next()
	cominfobuilder.ID("SystemMainRepositoryService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComSystemMainRepositoryServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: UserMainRepositoryService
	cominfobuilder.Next()
	cominfobuilder.ID("UserMainRepositoryService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUserMainRepositoryServiceImpl{}).init())
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

	// component: com24-controller0x9dc399.AuthController
	cominfobuilder.Next()
	cominfobuilder.ID("com24-controller0x9dc399.AuthController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAuthController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com25-controller0x9dc399.ExampleController
	cominfobuilder.Next()
	cominfobuilder.ID("com25-controller0x9dc399.ExampleController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com26-controller0x9dc399.ExecutableController
	cominfobuilder.Next()
	cominfobuilder.ID("com26-controller0x9dc399.ExecutableController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com27-controller0x9dc399.ExecutableImportController
	cominfobuilder.Next()
	cominfobuilder.ID("com27-controller0x9dc399.ExecutableImportController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableImportController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com28-controller0x9dc399.RunIntentController
	cominfobuilder.Next()
	cominfobuilder.ID("com28-controller0x9dc399.RunIntentController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRunIntentController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com29-controller0x9dc399.IntentTemplateController
	cominfobuilder.Next()
	cominfobuilder.ID("com29-controller0x9dc399.IntentTemplateController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentTemplateController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com30-controller0x9dc399.LocalRepositoryController
	cominfobuilder.Next()
	cominfobuilder.ID("com30-controller0x9dc399.LocalRepositoryController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalRepositoryController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com31-controller0x9dc399.MediaController
	cominfobuilder.Next()
	cominfobuilder.ID("com31-controller0x9dc399.MediaController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMediaController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com32-controller0x9dc399.PlatformController
	cominfobuilder.Next()
	cominfobuilder.ID("com32-controller0x9dc399.PlatformController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPlatformController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com33-controller0x9dc399.ProjectController
	cominfobuilder.Next()
	cominfobuilder.ID("com33-controller0x9dc399.ProjectController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com34-controller0x9dc399.ProjectImportController
	cominfobuilder.Next()
	cominfobuilder.ID("com34-controller0x9dc399.ProjectImportController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectImportController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com35-controller0x9dc399.RemoteRepositoryController
	cominfobuilder.Next()
	cominfobuilder.ID("com35-controller0x9dc399.RemoteRepositoryController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRemoteRepositoryController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com36-controller0x9dc399.RepositoryImportController
	cominfobuilder.Next()
	cominfobuilder.ID("com36-controller0x9dc399.RepositoryImportController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryImportController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com37-controller0x9dc399.UserMainRepositoryController
	cominfobuilder.Next()
	cominfobuilder.ID("com37-controller0x9dc399.UserMainRepositoryController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUserMainRepositoryController{}).init())
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

}

func (inst * comFactory4pComExecutableServiceImpl) init() application.ComponentFactory {

	
	inst.mExecutableDAOSelector = config.NewInjectionSelector("#ExecutableDAO",nil)


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

// comFactory4pComRepositoryServiceImpl : the factory of component: RepositoryService
type comFactory4pComRepositoryServiceImpl struct {

    mPrototype * implservice0x22327c.RepositoryServiceImpl

	
	mUUIDGenServiceSelector config.InjectionSelector
	mRepoFinderSelector config.InjectionSelector
	mLocalRepositoryDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryServiceImpl) init() application.ComponentFactory {

	
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)
	inst.mRepoFinderSelector = config.NewInjectionSelector("#LocalRepositoryFinder",nil)
	inst.mLocalRepositoryDAOSelector = config.NewInjectionSelector("#LocalRepositoryDAO",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryServiceImpl) newObject() * implservice0x22327c.RepositoryServiceImpl {
	return & implservice0x22327c.RepositoryServiceImpl {}
}

func (inst * comFactory4pComRepositoryServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.RepositoryServiceImpl {
	return instance.Get().(*implservice0x22327c.RepositoryServiceImpl)
}

func (inst * comFactory4pComRepositoryServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositoryServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositoryServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositoryServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	obj.RepoFinder = inst.getterForFieldRepoFinderSelector(context)
	obj.LocalRepositoryDAO = inst.getterForFieldLocalRepositoryDAOSelector(context)
	return context.LastError()
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComRepositoryServiceImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryService")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldRepoFinderSelector
func (inst * comFactory4pComRepositoryServiceImpl) getterForFieldRepoFinderSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryFinder {

	o1 := inst.mRepoFinderSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryFinder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryService")
		eb.Set("field", "RepoFinder")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryFinder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLocalRepositoryDAOSelector
func (inst * comFactory4pComRepositoryServiceImpl) getterForFieldLocalRepositoryDAOSelector (context application.InstanceContext) dao0x5af8d0.LocalRepositoryDAO {

	o1 := inst.mLocalRepositoryDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.LocalRepositoryDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryService")
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

	
	mDaoSelector config.InjectionSelector
	mGitLibAgentSelector config.InjectionSelector

}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) init() application.ComponentFactory {

	
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
	obj.Dao = inst.getterForFieldDaoSelector(context)
	obj.GitLibAgent = inst.getterForFieldGitLibAgentSelector(context)
	return context.LastError()
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

// comFactory4pComMediaServiceImpl : the factory of component: MediaService
type comFactory4pComMediaServiceImpl struct {

    mPrototype * implservice0x22327c.MediaServiceImpl

	
	mMediaDAOSelector config.InjectionSelector
	mSysMainRepoServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComMediaServiceImpl) init() application.ComponentFactory {

	
	inst.mMediaDAOSelector = config.NewInjectionSelector("#MediaDAO",nil)
	inst.mSysMainRepoServiceSelector = config.NewInjectionSelector("#SystemMainRepositoryService",nil)


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
func (inst * comFactory4pComMediaServiceImpl) getterForFieldSysMainRepoServiceSelector (context application.InstanceContext) service0x3e063d.SystemMainRepositoryService {

	o1 := inst.mSysMainRepoServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.SystemMainRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MediaService")
		eb.Set("field", "SysMainRepoService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.SystemMainRepositoryService")
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

	
	inst.mRepositoryServiceSelector = config.NewInjectionSelector("#RepositoryService",nil)
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

	
	mExecutableServiceSelector config.InjectionSelector
	mIntentHandlerServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComRunIntentServiceImpl) init() application.ComponentFactory {

	
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
	obj.ExecutableService = inst.getterForFieldExecutableServiceSelector(context)
	obj.IntentHandlerService = inst.getterForFieldIntentHandlerServiceSelector(context)
	return context.LastError()
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

// comFactory4pComSystemMainRepositoryServiceImpl : the factory of component: SystemMainRepositoryService
type comFactory4pComSystemMainRepositoryServiceImpl struct {

    mPrototype * implservice0x22327c.SystemMainRepositoryServiceImpl

	
	mSystemMainRepoPathSelector config.InjectionSelector
	mGitLASelector config.InjectionSelector

}

func (inst * comFactory4pComSystemMainRepositoryServiceImpl) init() application.ComponentFactory {

	
	inst.mSystemMainRepoPathSelector = config.NewInjectionSelector("${wpm.system-main-repository.path}",nil)
	inst.mGitLASelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComSystemMainRepositoryServiceImpl) newObject() * implservice0x22327c.SystemMainRepositoryServiceImpl {
	return & implservice0x22327c.SystemMainRepositoryServiceImpl {}
}

func (inst * comFactory4pComSystemMainRepositoryServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.SystemMainRepositoryServiceImpl {
	return instance.Get().(*implservice0x22327c.SystemMainRepositoryServiceImpl)
}

func (inst * comFactory4pComSystemMainRepositoryServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComSystemMainRepositoryServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComSystemMainRepositoryServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComSystemMainRepositoryServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSystemMainRepositoryServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSystemMainRepositoryServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.SystemMainRepoPath = inst.getterForFieldSystemMainRepoPathSelector(context)
	obj.GitLA = inst.getterForFieldGitLASelector(context)
	return context.LastError()
}

//getterForFieldSystemMainRepoPathSelector
func (inst * comFactory4pComSystemMainRepositoryServiceImpl) getterForFieldSystemMainRepoPathSelector (context application.InstanceContext) string {
    return inst.mSystemMainRepoPathSelector.GetString(context)
}

//getterForFieldGitLASelector
func (inst * comFactory4pComSystemMainRepositoryServiceImpl) getterForFieldGitLASelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mGitLASelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SystemMainRepositoryService")
		eb.Set("field", "GitLA")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUserMainRepositoryServiceImpl : the factory of component: UserMainRepositoryService
type comFactory4pComUserMainRepositoryServiceImpl struct {

    mPrototype * implservice0x22327c.UserMainRepositoryServiceImpl

	

}

func (inst * comFactory4pComUserMainRepositoryServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUserMainRepositoryServiceImpl) newObject() * implservice0x22327c.UserMainRepositoryServiceImpl {
	return & implservice0x22327c.UserMainRepositoryServiceImpl {}
}

func (inst * comFactory4pComUserMainRepositoryServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.UserMainRepositoryServiceImpl {
	return instance.Get().(*implservice0x22327c.UserMainRepositoryServiceImpl)
}

func (inst * comFactory4pComUserMainRepositoryServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComUserMainRepositoryServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComUserMainRepositoryServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComUserMainRepositoryServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUserMainRepositoryServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUserMainRepositoryServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
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

// comFactory4pComAuthController : the factory of component: com24-controller0x9dc399.AuthController
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
		eb.Set("com", "com24-controller0x9dc399.AuthController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExampleController : the factory of component: com25-controller0x9dc399.ExampleController
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
		eb.Set("com", "com25-controller0x9dc399.ExampleController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableController : the factory of component: com26-controller0x9dc399.ExecutableController
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
		eb.Set("com", "com26-controller0x9dc399.ExecutableController")
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
		eb.Set("com", "com26-controller0x9dc399.ExecutableController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableImportController : the factory of component: com27-controller0x9dc399.ExecutableImportController
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
		eb.Set("com", "com27-controller0x9dc399.ExecutableImportController")
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
		eb.Set("com", "com27-controller0x9dc399.ExecutableImportController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRunIntentController : the factory of component: com28-controller0x9dc399.RunIntentController
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
		eb.Set("com", "com28-controller0x9dc399.RunIntentController")
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
		eb.Set("com", "com28-controller0x9dc399.RunIntentController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentTemplateController : the factory of component: com29-controller0x9dc399.IntentTemplateController
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
		eb.Set("com", "com29-controller0x9dc399.IntentTemplateController")
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
		eb.Set("com", "com29-controller0x9dc399.IntentTemplateController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalRepositoryController : the factory of component: com30-controller0x9dc399.LocalRepositoryController
type comFactory4pComLocalRepositoryController struct {

    mPrototype * controller0x9dc399.LocalRepositoryController

	
	mRepoServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComLocalRepositoryController) init() application.ComponentFactory {

	
	inst.mRepoServiceSelector = config.NewInjectionSelector("#RepositoryService",nil)
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
		eb.Set("com", "com30-controller0x9dc399.LocalRepositoryController")
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
		eb.Set("com", "com30-controller0x9dc399.LocalRepositoryController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMediaController : the factory of component: com31-controller0x9dc399.MediaController
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
		eb.Set("com", "com31-controller0x9dc399.MediaController")
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
		eb.Set("com", "com31-controller0x9dc399.MediaController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPlatformController : the factory of component: com32-controller0x9dc399.PlatformController
type comFactory4pComPlatformController struct {

    mPrototype * controller0x9dc399.PlatformController

	
	mPlatformServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComPlatformController) init() application.ComponentFactory {

	
	inst.mPlatformServiceSelector = config.NewInjectionSelector("#PlatformService",nil)
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
		eb.Set("com", "com32-controller0x9dc399.PlatformController")
		eb.Set("field", "PlatformService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.PlatformService")
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
		eb.Set("com", "com32-controller0x9dc399.PlatformController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectController : the factory of component: com33-controller0x9dc399.ProjectController
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
		eb.Set("com", "com33-controller0x9dc399.ProjectController")
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
		eb.Set("com", "com33-controller0x9dc399.ProjectController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectImportController : the factory of component: com34-controller0x9dc399.ProjectImportController
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
		eb.Set("com", "com34-controller0x9dc399.ProjectImportController")
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
		eb.Set("com", "com34-controller0x9dc399.ProjectImportController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRemoteRepositoryController : the factory of component: com35-controller0x9dc399.RemoteRepositoryController
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
		eb.Set("com", "com35-controller0x9dc399.RemoteRepositoryController")
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
		eb.Set("com", "com35-controller0x9dc399.RemoteRepositoryController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryImportController : the factory of component: com36-controller0x9dc399.RepositoryImportController
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
		eb.Set("com", "com36-controller0x9dc399.RepositoryImportController")
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
		eb.Set("com", "com36-controller0x9dc399.RepositoryImportController")
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
		eb.Set("com", "com36-controller0x9dc399.RepositoryImportController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUserMainRepositoryController : the factory of component: com37-controller0x9dc399.UserMainRepositoryController
type comFactory4pComUserMainRepositoryController struct {

    mPrototype * controller0x9dc399.UserMainRepositoryController

	
	mRepoServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComUserMainRepositoryController) init() application.ComponentFactory {

	
	inst.mRepoServiceSelector = config.NewInjectionSelector("#UserMainRepositoryService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUserMainRepositoryController) newObject() * controller0x9dc399.UserMainRepositoryController {
	return & controller0x9dc399.UserMainRepositoryController {}
}

func (inst * comFactory4pComUserMainRepositoryController) castObject(instance application.ComponentInstance) * controller0x9dc399.UserMainRepositoryController {
	return instance.Get().(*controller0x9dc399.UserMainRepositoryController)
}

func (inst * comFactory4pComUserMainRepositoryController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComUserMainRepositoryController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComUserMainRepositoryController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComUserMainRepositoryController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUserMainRepositoryController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUserMainRepositoryController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.RepoService = inst.getterForFieldRepoServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldRepoServiceSelector
func (inst * comFactory4pComUserMainRepositoryController) getterForFieldRepoServiceSelector (context application.InstanceContext) service0x3e063d.UserMainRepositoryService {

	o1 := inst.mRepoServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UserMainRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com37-controller0x9dc399.UserMainRepositoryController")
		eb.Set("field", "RepoService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UserMainRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComUserMainRepositoryController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com37-controller0x9dc399.UserMainRepositoryController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}




