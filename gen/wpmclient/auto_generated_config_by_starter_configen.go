// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmclient

import (
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
	client0x4f5e0e "github.com/bitwormhole/wpm/client"
	service0x3e063d "github.com/bitwormhole/wpm/server/service"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-client0x4f5e0e.Client
	cominfobuilder.Next()
	cominfobuilder.ID("com0-client0x4f5e0e.Client").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComClient{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: IntentHandler
	cominfobuilder.Next()
	cominfobuilder.ID("IntentHandler").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentHandlerImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComClient : the factory of component: com0-client0x4f5e0e.Client
type comFactory4pComClient struct {

    mPrototype * client0x4f5e0e.Client

	
	mIntentHandlerSelector config.InjectionSelector
	mPlatformServiceSelector config.InjectionSelector
	mProtocolSelector config.InjectionSelector
	mHostSelector config.InjectionSelector
	mPortSelector config.InjectionSelector

}

func (inst * comFactory4pComClient) init() application.ComponentFactory {

	
	inst.mIntentHandlerSelector = config.NewInjectionSelector("#IntentHandler",nil)
	inst.mPlatformServiceSelector = config.NewInjectionSelector("#PlatformService",nil)
	inst.mProtocolSelector = config.NewInjectionSelector("${wpm.server.protocol}",nil)
	inst.mHostSelector = config.NewInjectionSelector("${wpm.server.host}",nil)
	inst.mPortSelector = config.NewInjectionSelector("${wpm.server.port}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComClient) newObject() * client0x4f5e0e.Client {
	return & client0x4f5e0e.Client {}
}

func (inst * comFactory4pComClient) castObject(instance application.ComponentInstance) * client0x4f5e0e.Client {
	return instance.Get().(*client0x4f5e0e.Client)
}

func (inst * comFactory4pComClient) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComClient) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComClient) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComClient) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComClient) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComClient) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.IntentHandler = inst.getterForFieldIntentHandlerSelector(context)
	obj.PlatformService = inst.getterForFieldPlatformServiceSelector(context)
	obj.Protocol = inst.getterForFieldProtocolSelector(context)
	obj.Host = inst.getterForFieldHostSelector(context)
	obj.Port = inst.getterForFieldPortSelector(context)
	return context.LastError()
}

//getterForFieldIntentHandlerSelector
func (inst * comFactory4pComClient) getterForFieldIntentHandlerSelector (context application.InstanceContext) client0x4f5e0e.IntentHandler {

	o1 := inst.mIntentHandlerSelector.GetOne(context)
	o2, ok := o1.(client0x4f5e0e.IntentHandler)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-client0x4f5e0e.Client")
		eb.Set("field", "IntentHandler")
		eb.Set("type1", "?")
		eb.Set("type2", "client0x4f5e0e.IntentHandler")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldPlatformServiceSelector
func (inst * comFactory4pComClient) getterForFieldPlatformServiceSelector (context application.InstanceContext) service0x3e063d.PlatformService {

	o1 := inst.mPlatformServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.PlatformService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-client0x4f5e0e.Client")
		eb.Set("field", "PlatformService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.PlatformService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProtocolSelector
func (inst * comFactory4pComClient) getterForFieldProtocolSelector (context application.InstanceContext) string {
    return inst.mProtocolSelector.GetString(context)
}

//getterForFieldHostSelector
func (inst * comFactory4pComClient) getterForFieldHostSelector (context application.InstanceContext) string {
    return inst.mHostSelector.GetString(context)
}

//getterForFieldPortSelector
func (inst * comFactory4pComClient) getterForFieldPortSelector (context application.InstanceContext) int {
    return inst.mPortSelector.GetInt(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentHandlerImpl : the factory of component: IntentHandler
type comFactory4pComIntentHandlerImpl struct {

    mPrototype * client0x4f5e0e.IntentHandlerImpl

	

}

func (inst * comFactory4pComIntentHandlerImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentHandlerImpl) newObject() * client0x4f5e0e.IntentHandlerImpl {
	return & client0x4f5e0e.IntentHandlerImpl {}
}

func (inst * comFactory4pComIntentHandlerImpl) castObject(instance application.ComponentInstance) * client0x4f5e0e.IntentHandlerImpl {
	return instance.Get().(*client0x4f5e0e.IntentHandlerImpl)
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




