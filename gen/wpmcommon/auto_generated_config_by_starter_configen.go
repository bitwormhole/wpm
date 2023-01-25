// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmcommon

import (
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
	implservice0x5a8f41 "github.com/bitwormhole/wpm/common/implservice"
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

	// component: PlatformService
	cominfobuilder.Next()
	cominfobuilder.ID("PlatformService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPlatformServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-implservice0x5a8f41.LinuxPlatformServiceImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com1-implservice0x5a8f41.LinuxPlatformServiceImpl").Class("PlatformProviderRegistry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLinuxPlatformServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-implservice0x5a8f41.WindowsPlatformServiceImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com2-implservice0x5a8f41.WindowsPlatformServiceImpl").Class("PlatformProviderRegistry").Aliases("").Scope("")
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

// comFactory4pComLinuxPlatformServiceImpl : the factory of component: com1-implservice0x5a8f41.LinuxPlatformServiceImpl
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

// comFactory4pComWindowsPlatformServiceImpl : the factory of component: com2-implservice0x5a8f41.WindowsPlatformServiceImpl
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




