// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmappcfg

import (
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
	dao0xda26a8 "github.com/bitwormhole/wpm/app/data/dao"
	service0xe6cb35 "github.com/bitwormhole/wpm/app/service"
	impldao0x204069 "github.com/bitwormhole/wpm/app/support/impldao"
	implservice0x5e39ce "github.com/bitwormhole/wpm/app/support/implservice"
	controller0x906cab "github.com/bitwormhole/wpm/app/web/controller"
    
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
	cominfobuilder.ID("GormDBAgent").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComGormDBAgentImpl{}).init())
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

	// component: IntentDAO
	cominfobuilder.Next()
	cominfobuilder.ID("IntentDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentDaoImpl{}).init())
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

	// component: RepositoryDAO
	cominfobuilder.Next()
	cominfobuilder.ID("RepositoryDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryDaoImpl{}).init())
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

	// component: IntentService
	cominfobuilder.Next()
	cominfobuilder.ID("IntentService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentServiceImpl{}).init())
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

	// component: RepositoryImportService
	cominfobuilder.Next()
	cominfobuilder.ID("RepositoryImportService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryImportServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: RepositoryLocateService
	cominfobuilder.Next()
	cominfobuilder.ID("RepositoryLocateService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryLocateServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: RepositorySearchService
	cominfobuilder.Next()
	cominfobuilder.ID("RepositorySearchService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositorySearchServiceImpl{}).init())
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

	// component: UUIDGenService
	cominfobuilder.Next()
	cominfobuilder.ID("UUIDGenService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUUIDGenServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com15-controller0x906cab.AuthController
	cominfobuilder.Next()
	cominfobuilder.ID("com15-controller0x906cab.AuthController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAuthController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com16-controller0x906cab.ExecutableController
	cominfobuilder.Next()
	cominfobuilder.ID("com16-controller0x906cab.ExecutableController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com17-controller0x906cab.ExecutableImportController
	cominfobuilder.Next()
	cominfobuilder.ID("com17-controller0x906cab.ExecutableImportController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableImportController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com18-controller0x906cab.IntentController
	cominfobuilder.Next()
	cominfobuilder.ID("com18-controller0x906cab.IntentController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com19-controller0x906cab.ProjectController
	cominfobuilder.Next()
	cominfobuilder.ID("com19-controller0x906cab.ProjectController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com20-controller0x906cab.ProjectImportController
	cominfobuilder.Next()
	cominfobuilder.ID("com20-controller0x906cab.ProjectImportController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectImportController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com21-controller0x906cab.RepositoryController
	cominfobuilder.Next()
	cominfobuilder.ID("com21-controller0x906cab.RepositoryController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com22-controller0x906cab.RepositoryImportController
	cominfobuilder.Next()
	cominfobuilder.ID("com22-controller0x906cab.RepositoryImportController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryImportController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGormDBAgentImpl : the factory of component: GormDBAgent
type comFactory4pComGormDBAgentImpl struct {

    mPrototype * impldao0x204069.GormDBAgentImpl

	
	mSourceSelector config.InjectionSelector

}

func (inst * comFactory4pComGormDBAgentImpl) init() application.ComponentFactory {

	
	inst.mSourceSelector = config.NewInjectionSelector("#gorm-datasource-default",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGormDBAgentImpl) newObject() * impldao0x204069.GormDBAgentImpl {
	return & impldao0x204069.GormDBAgentImpl {}
}

func (inst * comFactory4pComGormDBAgentImpl) castObject(instance application.ComponentInstance) * impldao0x204069.GormDBAgentImpl {
	return instance.Get().(*impldao0x204069.GormDBAgentImpl)
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
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComGormDBAgentImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGormDBAgentImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Source = inst.getterForFieldSourceSelector(context)
	return context.LastError()
}

//getterForFieldSourceSelector
func (inst * comFactory4pComGormDBAgentImpl) getterForFieldSourceSelector (context application.InstanceContext) datasource0x68a737.Source {

	o1 := inst.mSourceSelector.GetOne(context)
	o2, ok := o1.(datasource0x68a737.Source)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "GormDBAgent")
		eb.Set("field", "Source")
		eb.Set("type1", "?")
		eb.Set("type2", "datasource0x68a737.Source")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableDaoImpl : the factory of component: ExecutableDAO
type comFactory4pComExecutableDaoImpl struct {

    mPrototype * impldao0x204069.ExecutableDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableDaoImpl) newObject() * impldao0x204069.ExecutableDaoImpl {
	return & impldao0x204069.ExecutableDaoImpl {}
}

func (inst * comFactory4pComExecutableDaoImpl) castObject(instance application.ComponentInstance) * impldao0x204069.ExecutableDaoImpl {
	return instance.Get().(*impldao0x204069.ExecutableDaoImpl)
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
func (inst * comFactory4pComExecutableDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) impldao0x204069.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(impldao0x204069.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "impldao0x204069.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComExecutableDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0xe6cb35.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentDaoImpl : the factory of component: IntentDAO
type comFactory4pComIntentDaoImpl struct {

    mPrototype * impldao0x204069.IntentDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComIntentDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentDaoImpl) newObject() * impldao0x204069.IntentDaoImpl {
	return & impldao0x204069.IntentDaoImpl {}
}

func (inst * comFactory4pComIntentDaoImpl) castObject(instance application.ComponentInstance) * impldao0x204069.IntentDaoImpl {
	return instance.Get().(*impldao0x204069.IntentDaoImpl)
}

func (inst * comFactory4pComIntentDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComIntentDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) impldao0x204069.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(impldao0x204069.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "impldao0x204069.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComIntentDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0xe6cb35.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectDaoImpl : the factory of component: ProjectDAO
type comFactory4pComProjectDaoImpl struct {

    mPrototype * impldao0x204069.ProjectDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectDaoImpl) newObject() * impldao0x204069.ProjectDaoImpl {
	return & impldao0x204069.ProjectDaoImpl {}
}

func (inst * comFactory4pComProjectDaoImpl) castObject(instance application.ComponentInstance) * impldao0x204069.ProjectDaoImpl {
	return instance.Get().(*impldao0x204069.ProjectDaoImpl)
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
func (inst * comFactory4pComProjectDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) impldao0x204069.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(impldao0x204069.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "impldao0x204069.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComProjectDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0xe6cb35.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryDaoImpl : the factory of component: RepositoryDAO
type comFactory4pComRepositoryDaoImpl struct {

    mPrototype * impldao0x204069.RepositoryDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryDaoImpl) newObject() * impldao0x204069.RepositoryDaoImpl {
	return & impldao0x204069.RepositoryDaoImpl {}
}

func (inst * comFactory4pComRepositoryDaoImpl) castObject(instance application.ComponentInstance) * impldao0x204069.RepositoryDaoImpl {
	return instance.Get().(*impldao0x204069.RepositoryDaoImpl)
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
func (inst * comFactory4pComRepositoryDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) impldao0x204069.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(impldao0x204069.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "impldao0x204069.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComRepositoryDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0xe6cb35.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableImportServiceImpl : the factory of component: ExecutableImportService
type comFactory4pComExecutableImportServiceImpl struct {

    mPrototype * implservice0x5e39ce.ExecutableImportServiceImpl

	
	mExecutableServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableImportServiceImpl) init() application.ComponentFactory {

	
	inst.mExecutableServiceSelector = config.NewInjectionSelector("#ExecutableService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableImportServiceImpl) newObject() * implservice0x5e39ce.ExecutableImportServiceImpl {
	return & implservice0x5e39ce.ExecutableImportServiceImpl {}
}

func (inst * comFactory4pComExecutableImportServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5e39ce.ExecutableImportServiceImpl {
	return instance.Get().(*implservice0x5e39ce.ExecutableImportServiceImpl)
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
func (inst * comFactory4pComExecutableImportServiceImpl) getterForFieldExecutableServiceSelector (context application.InstanceContext) service0xe6cb35.ExecutableService {

	o1 := inst.mExecutableServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.ExecutableService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableImportService")
		eb.Set("field", "ExecutableService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.ExecutableService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableServiceImpl : the factory of component: ExecutableService
type comFactory4pComExecutableServiceImpl struct {

    mPrototype * implservice0x5e39ce.ExecutableServiceImpl

	
	mExecutableDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableServiceImpl) init() application.ComponentFactory {

	
	inst.mExecutableDAOSelector = config.NewInjectionSelector("#ExecutableDAO",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableServiceImpl) newObject() * implservice0x5e39ce.ExecutableServiceImpl {
	return & implservice0x5e39ce.ExecutableServiceImpl {}
}

func (inst * comFactory4pComExecutableServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5e39ce.ExecutableServiceImpl {
	return instance.Get().(*implservice0x5e39ce.ExecutableServiceImpl)
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
func (inst * comFactory4pComExecutableServiceImpl) getterForFieldExecutableDAOSelector (context application.InstanceContext) dao0xda26a8.ExecutableDAO {

	o1 := inst.mExecutableDAOSelector.GetOne(context)
	o2, ok := o1.(dao0xda26a8.ExecutableDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableService")
		eb.Set("field", "ExecutableDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0xda26a8.ExecutableDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentServiceImpl : the factory of component: IntentService
type comFactory4pComIntentServiceImpl struct {

    mPrototype * implservice0x5e39ce.IntentServiceImpl

	
	mIntentDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComIntentServiceImpl) init() application.ComponentFactory {

	
	inst.mIntentDAOSelector = config.NewInjectionSelector("#IntentDAO",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentServiceImpl) newObject() * implservice0x5e39ce.IntentServiceImpl {
	return & implservice0x5e39ce.IntentServiceImpl {}
}

func (inst * comFactory4pComIntentServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5e39ce.IntentServiceImpl {
	return instance.Get().(*implservice0x5e39ce.IntentServiceImpl)
}

func (inst * comFactory4pComIntentServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.IntentDAO = inst.getterForFieldIntentDAOSelector(context)
	return context.LastError()
}

//getterForFieldIntentDAOSelector
func (inst * comFactory4pComIntentServiceImpl) getterForFieldIntentDAOSelector (context application.InstanceContext) dao0xda26a8.IntentDAO {

	o1 := inst.mIntentDAOSelector.GetOne(context)
	o2, ok := o1.(dao0xda26a8.IntentDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentService")
		eb.Set("field", "IntentDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0xda26a8.IntentDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectImportServiceImpl : the factory of component: ProjectImportService
type comFactory4pComProjectImportServiceImpl struct {

    mPrototype * implservice0x5e39ce.ProjectImportServiceImpl

	

}

func (inst * comFactory4pComProjectImportServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectImportServiceImpl) newObject() * implservice0x5e39ce.ProjectImportServiceImpl {
	return & implservice0x5e39ce.ProjectImportServiceImpl {}
}

func (inst * comFactory4pComProjectImportServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5e39ce.ProjectImportServiceImpl {
	return instance.Get().(*implservice0x5e39ce.ProjectImportServiceImpl)
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

    mPrototype * implservice0x5e39ce.ProjectServiceImpl

	
	mUUIDGenServiceSelector config.InjectionSelector
	mProjectDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectServiceImpl) init() application.ComponentFactory {

	
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)
	inst.mProjectDAOSelector = config.NewInjectionSelector("#ProjectDAO",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectServiceImpl) newObject() * implservice0x5e39ce.ProjectServiceImpl {
	return & implservice0x5e39ce.ProjectServiceImpl {}
}

func (inst * comFactory4pComProjectServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5e39ce.ProjectServiceImpl {
	return instance.Get().(*implservice0x5e39ce.ProjectServiceImpl)
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
func (inst * comFactory4pComProjectServiceImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0xe6cb35.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectService")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProjectDAOSelector
func (inst * comFactory4pComProjectServiceImpl) getterForFieldProjectDAOSelector (context application.InstanceContext) dao0xda26a8.ProjectDAO {

	o1 := inst.mProjectDAOSelector.GetOne(context)
	o2, ok := o1.(dao0xda26a8.ProjectDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectService")
		eb.Set("field", "ProjectDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0xda26a8.ProjectDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryImportServiceImpl : the factory of component: RepositoryImportService
type comFactory4pComRepositoryImportServiceImpl struct {

    mPrototype * implservice0x5e39ce.RepositoryImportServiceImpl

	
	mLocateServiceSelector config.InjectionSelector
	mSearchServiceSelector config.InjectionSelector
	mRepositoryServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryImportServiceImpl) init() application.ComponentFactory {

	
	inst.mLocateServiceSelector = config.NewInjectionSelector("#RepositoryLocateService",nil)
	inst.mSearchServiceSelector = config.NewInjectionSelector("#RepositorySearchService",nil)
	inst.mRepositoryServiceSelector = config.NewInjectionSelector("#RepositoryService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryImportServiceImpl) newObject() * implservice0x5e39ce.RepositoryImportServiceImpl {
	return & implservice0x5e39ce.RepositoryImportServiceImpl {}
}

func (inst * comFactory4pComRepositoryImportServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5e39ce.RepositoryImportServiceImpl {
	return instance.Get().(*implservice0x5e39ce.RepositoryImportServiceImpl)
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
	obj.LocateService = inst.getterForFieldLocateServiceSelector(context)
	obj.SearchService = inst.getterForFieldSearchServiceSelector(context)
	obj.RepositoryService = inst.getterForFieldRepositoryServiceSelector(context)
	return context.LastError()
}

//getterForFieldLocateServiceSelector
func (inst * comFactory4pComRepositoryImportServiceImpl) getterForFieldLocateServiceSelector (context application.InstanceContext) service0xe6cb35.RepositoryLocateService {

	o1 := inst.mLocateServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.RepositoryLocateService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryImportService")
		eb.Set("field", "LocateService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.RepositoryLocateService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSearchServiceSelector
func (inst * comFactory4pComRepositoryImportServiceImpl) getterForFieldSearchServiceSelector (context application.InstanceContext) service0xe6cb35.RepositorySearchService {

	o1 := inst.mSearchServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.RepositorySearchService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryImportService")
		eb.Set("field", "SearchService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.RepositorySearchService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldRepositoryServiceSelector
func (inst * comFactory4pComRepositoryImportServiceImpl) getterForFieldRepositoryServiceSelector (context application.InstanceContext) service0xe6cb35.RepositoryService {

	o1 := inst.mRepositoryServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.RepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryImportService")
		eb.Set("field", "RepositoryService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.RepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryLocateServiceImpl : the factory of component: RepositoryLocateService
type comFactory4pComRepositoryLocateServiceImpl struct {

    mPrototype * implservice0x5e39ce.RepositoryLocateServiceImpl

	

}

func (inst * comFactory4pComRepositoryLocateServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryLocateServiceImpl) newObject() * implservice0x5e39ce.RepositoryLocateServiceImpl {
	return & implservice0x5e39ce.RepositoryLocateServiceImpl {}
}

func (inst * comFactory4pComRepositoryLocateServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5e39ce.RepositoryLocateServiceImpl {
	return instance.Get().(*implservice0x5e39ce.RepositoryLocateServiceImpl)
}

func (inst * comFactory4pComRepositoryLocateServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositoryLocateServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositoryLocateServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositoryLocateServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryLocateServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryLocateServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositorySearchServiceImpl : the factory of component: RepositorySearchService
type comFactory4pComRepositorySearchServiceImpl struct {

    mPrototype * implservice0x5e39ce.RepositorySearchServiceImpl

	
	mLocateServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositorySearchServiceImpl) init() application.ComponentFactory {

	
	inst.mLocateServiceSelector = config.NewInjectionSelector("#RepositoryLocateService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositorySearchServiceImpl) newObject() * implservice0x5e39ce.RepositorySearchServiceImpl {
	return & implservice0x5e39ce.RepositorySearchServiceImpl {}
}

func (inst * comFactory4pComRepositorySearchServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5e39ce.RepositorySearchServiceImpl {
	return instance.Get().(*implservice0x5e39ce.RepositorySearchServiceImpl)
}

func (inst * comFactory4pComRepositorySearchServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositorySearchServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositorySearchServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositorySearchServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositorySearchServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositorySearchServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.LocateService = inst.getterForFieldLocateServiceSelector(context)
	return context.LastError()
}

//getterForFieldLocateServiceSelector
func (inst * comFactory4pComRepositorySearchServiceImpl) getterForFieldLocateServiceSelector (context application.InstanceContext) service0xe6cb35.RepositoryLocateService {

	o1 := inst.mLocateServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.RepositoryLocateService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositorySearchService")
		eb.Set("field", "LocateService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.RepositoryLocateService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryServiceImpl : the factory of component: RepositoryService
type comFactory4pComRepositoryServiceImpl struct {

    mPrototype * implservice0x5e39ce.RepositoryServiceImpl

	
	mUUIDGenServiceSelector config.InjectionSelector
	mLocateServiceSelector config.InjectionSelector
	mRepositoryDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryServiceImpl) init() application.ComponentFactory {

	
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)
	inst.mLocateServiceSelector = config.NewInjectionSelector("#RepositoryLocateService",nil)
	inst.mRepositoryDAOSelector = config.NewInjectionSelector("#RepositoryDAO",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryServiceImpl) newObject() * implservice0x5e39ce.RepositoryServiceImpl {
	return & implservice0x5e39ce.RepositoryServiceImpl {}
}

func (inst * comFactory4pComRepositoryServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5e39ce.RepositoryServiceImpl {
	return instance.Get().(*implservice0x5e39ce.RepositoryServiceImpl)
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
	obj.LocateService = inst.getterForFieldLocateServiceSelector(context)
	obj.RepositoryDAO = inst.getterForFieldRepositoryDAOSelector(context)
	return context.LastError()
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComRepositoryServiceImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0xe6cb35.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryService")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLocateServiceSelector
func (inst * comFactory4pComRepositoryServiceImpl) getterForFieldLocateServiceSelector (context application.InstanceContext) service0xe6cb35.RepositoryLocateService {

	o1 := inst.mLocateServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.RepositoryLocateService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryService")
		eb.Set("field", "LocateService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.RepositoryLocateService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldRepositoryDAOSelector
func (inst * comFactory4pComRepositoryServiceImpl) getterForFieldRepositoryDAOSelector (context application.InstanceContext) dao0xda26a8.RepositoryDAO {

	o1 := inst.mRepositoryDAOSelector.GetOne(context)
	o2, ok := o1.(dao0xda26a8.RepositoryDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryService")
		eb.Set("field", "RepositoryDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0xda26a8.RepositoryDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUUIDGenServiceImpl : the factory of component: UUIDGenService
type comFactory4pComUUIDGenServiceImpl struct {

    mPrototype * implservice0x5e39ce.UUIDGenServiceImpl

	

}

func (inst * comFactory4pComUUIDGenServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUUIDGenServiceImpl) newObject() * implservice0x5e39ce.UUIDGenServiceImpl {
	return & implservice0x5e39ce.UUIDGenServiceImpl {}
}

func (inst * comFactory4pComUUIDGenServiceImpl) castObject(instance application.ComponentInstance) * implservice0x5e39ce.UUIDGenServiceImpl {
	return instance.Get().(*implservice0x5e39ce.UUIDGenServiceImpl)
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

// comFactory4pComAuthController : the factory of component: com15-controller0x906cab.AuthController
type comFactory4pComAuthController struct {

    mPrototype * controller0x906cab.AuthController

	
	mRepoServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComAuthController) init() application.ComponentFactory {

	
	inst.mRepoServiceSelector = config.NewInjectionSelector("#RepositoryService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAuthController) newObject() * controller0x906cab.AuthController {
	return & controller0x906cab.AuthController {}
}

func (inst * comFactory4pComAuthController) castObject(instance application.ComponentInstance) * controller0x906cab.AuthController {
	return instance.Get().(*controller0x906cab.AuthController)
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
	obj.RepoService = inst.getterForFieldRepoServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldRepoServiceSelector
func (inst * comFactory4pComAuthController) getterForFieldRepoServiceSelector (context application.InstanceContext) service0xe6cb35.RepositoryService {

	o1 := inst.mRepoServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.RepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com15-controller0x906cab.AuthController")
		eb.Set("field", "RepoService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.RepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComAuthController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com15-controller0x906cab.AuthController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableController : the factory of component: com16-controller0x906cab.ExecutableController
type comFactory4pComExecutableController struct {

    mPrototype * controller0x906cab.ExecutableController

	
	mExecutableServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableController) init() application.ComponentFactory {

	
	inst.mExecutableServiceSelector = config.NewInjectionSelector("#ExecutableService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableController) newObject() * controller0x906cab.ExecutableController {
	return & controller0x906cab.ExecutableController {}
}

func (inst * comFactory4pComExecutableController) castObject(instance application.ComponentInstance) * controller0x906cab.ExecutableController {
	return instance.Get().(*controller0x906cab.ExecutableController)
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
func (inst * comFactory4pComExecutableController) getterForFieldExecutableServiceSelector (context application.InstanceContext) service0xe6cb35.ExecutableService {

	o1 := inst.mExecutableServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.ExecutableService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com16-controller0x906cab.ExecutableController")
		eb.Set("field", "ExecutableService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.ExecutableService")
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
		eb.Set("com", "com16-controller0x906cab.ExecutableController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableImportController : the factory of component: com17-controller0x906cab.ExecutableImportController
type comFactory4pComExecutableImportController struct {

    mPrototype * controller0x906cab.ExecutableImportController

	
	mExecutableImportServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableImportController) init() application.ComponentFactory {

	
	inst.mExecutableImportServiceSelector = config.NewInjectionSelector("#ExecutableImportService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableImportController) newObject() * controller0x906cab.ExecutableImportController {
	return & controller0x906cab.ExecutableImportController {}
}

func (inst * comFactory4pComExecutableImportController) castObject(instance application.ComponentInstance) * controller0x906cab.ExecutableImportController {
	return instance.Get().(*controller0x906cab.ExecutableImportController)
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
func (inst * comFactory4pComExecutableImportController) getterForFieldExecutableImportServiceSelector (context application.InstanceContext) service0xe6cb35.ExecutableImportService {

	o1 := inst.mExecutableImportServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.ExecutableImportService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com17-controller0x906cab.ExecutableImportController")
		eb.Set("field", "ExecutableImportService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.ExecutableImportService")
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
		eb.Set("com", "com17-controller0x906cab.ExecutableImportController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentController : the factory of component: com18-controller0x906cab.IntentController
type comFactory4pComIntentController struct {

    mPrototype * controller0x906cab.IntentController

	
	mIntentServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComIntentController) init() application.ComponentFactory {

	
	inst.mIntentServiceSelector = config.NewInjectionSelector("#IntentService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentController) newObject() * controller0x906cab.IntentController {
	return & controller0x906cab.IntentController {}
}

func (inst * comFactory4pComIntentController) castObject(instance application.ComponentInstance) * controller0x906cab.IntentController {
	return instance.Get().(*controller0x906cab.IntentController)
}

func (inst * comFactory4pComIntentController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.IntentService = inst.getterForFieldIntentServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldIntentServiceSelector
func (inst * comFactory4pComIntentController) getterForFieldIntentServiceSelector (context application.InstanceContext) service0xe6cb35.IntentService {

	o1 := inst.mIntentServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.IntentService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com18-controller0x906cab.IntentController")
		eb.Set("field", "IntentService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.IntentService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComIntentController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com18-controller0x906cab.IntentController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectController : the factory of component: com19-controller0x906cab.ProjectController
type comFactory4pComProjectController struct {

    mPrototype * controller0x906cab.ProjectController

	
	mProjectServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectController) init() application.ComponentFactory {

	
	inst.mProjectServiceSelector = config.NewInjectionSelector("#ProjectService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectController) newObject() * controller0x906cab.ProjectController {
	return & controller0x906cab.ProjectController {}
}

func (inst * comFactory4pComProjectController) castObject(instance application.ComponentInstance) * controller0x906cab.ProjectController {
	return instance.Get().(*controller0x906cab.ProjectController)
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
func (inst * comFactory4pComProjectController) getterForFieldProjectServiceSelector (context application.InstanceContext) service0xe6cb35.ProjectService {

	o1 := inst.mProjectServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.ProjectService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com19-controller0x906cab.ProjectController")
		eb.Set("field", "ProjectService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.ProjectService")
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
		eb.Set("com", "com19-controller0x906cab.ProjectController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectImportController : the factory of component: com20-controller0x906cab.ProjectImportController
type comFactory4pComProjectImportController struct {

    mPrototype * controller0x906cab.ProjectImportController

	
	mProjectImportServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectImportController) init() application.ComponentFactory {

	
	inst.mProjectImportServiceSelector = config.NewInjectionSelector("#ProjectImportService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectImportController) newObject() * controller0x906cab.ProjectImportController {
	return & controller0x906cab.ProjectImportController {}
}

func (inst * comFactory4pComProjectImportController) castObject(instance application.ComponentInstance) * controller0x906cab.ProjectImportController {
	return instance.Get().(*controller0x906cab.ProjectImportController)
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
func (inst * comFactory4pComProjectImportController) getterForFieldProjectImportServiceSelector (context application.InstanceContext) service0xe6cb35.ProjectImportService {

	o1 := inst.mProjectImportServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.ProjectImportService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com20-controller0x906cab.ProjectImportController")
		eb.Set("field", "ProjectImportService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.ProjectImportService")
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
		eb.Set("com", "com20-controller0x906cab.ProjectImportController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryController : the factory of component: com21-controller0x906cab.RepositoryController
type comFactory4pComRepositoryController struct {

    mPrototype * controller0x906cab.RepositoryController

	
	mRepoServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryController) init() application.ComponentFactory {

	
	inst.mRepoServiceSelector = config.NewInjectionSelector("#RepositoryService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryController) newObject() * controller0x906cab.RepositoryController {
	return & controller0x906cab.RepositoryController {}
}

func (inst * comFactory4pComRepositoryController) castObject(instance application.ComponentInstance) * controller0x906cab.RepositoryController {
	return instance.Get().(*controller0x906cab.RepositoryController)
}

func (inst * comFactory4pComRepositoryController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositoryController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositoryController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositoryController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.RepoService = inst.getterForFieldRepoServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldRepoServiceSelector
func (inst * comFactory4pComRepositoryController) getterForFieldRepoServiceSelector (context application.InstanceContext) service0xe6cb35.RepositoryService {

	o1 := inst.mRepoServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.RepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com21-controller0x906cab.RepositoryController")
		eb.Set("field", "RepoService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.RepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComRepositoryController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com21-controller0x906cab.RepositoryController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryImportController : the factory of component: com22-controller0x906cab.RepositoryImportController
type comFactory4pComRepositoryImportController struct {

    mPrototype * controller0x906cab.RepositoryImportController

	
	mImportServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryImportController) init() application.ComponentFactory {

	
	inst.mImportServiceSelector = config.NewInjectionSelector("#RepositoryImportService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryImportController) newObject() * controller0x906cab.RepositoryImportController {
	return & controller0x906cab.RepositoryImportController {}
}

func (inst * comFactory4pComRepositoryImportController) castObject(instance application.ComponentInstance) * controller0x906cab.RepositoryImportController {
	return instance.Get().(*controller0x906cab.RepositoryImportController)
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
	obj.ImportService = inst.getterForFieldImportServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldImportServiceSelector
func (inst * comFactory4pComRepositoryImportController) getterForFieldImportServiceSelector (context application.InstanceContext) service0xe6cb35.RepositoryImportService {

	o1 := inst.mImportServiceSelector.GetOne(context)
	o2, ok := o1.(service0xe6cb35.RepositoryImportService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com22-controller0x906cab.RepositoryImportController")
		eb.Set("field", "ImportService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xe6cb35.RepositoryImportService")
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
		eb.Set("com", "com22-controller0x906cab.RepositoryImportController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}




