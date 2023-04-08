// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmboot

import (
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
	boot0x8dc527 "github.com/bitwormhole/wpm/boot"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-boot0x8dc527.InfoLoader
	cominfobuilder.Next()
	cominfobuilder.ID("com0-boot0x8dc527.InfoLoader").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComInfoLoader{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComInfoLoader : the factory of component: com0-boot0x8dc527.InfoLoader
type comFactory4pComInfoLoader struct {

    mPrototype * boot0x8dc527.InfoLoader

	
	mACSelector config.InjectionSelector
	mDoBackupExeSelector config.InjectionSelector
	mDoDebugSelector config.InjectionSelector
	mDoDumpSelector config.InjectionSelector
	mDoLogOptionsSelector config.InjectionSelector
	mDoRunWithGUISelector config.InjectionSelector

}

func (inst * comFactory4pComInfoLoader) init() application.ComponentFactory {

	
	inst.mACSelector = config.NewInjectionSelector("context",nil)
	inst.mDoBackupExeSelector = config.NewInjectionSelector("${wpm.options.backup-this-exe}",nil)
	inst.mDoDebugSelector = config.NewInjectionSelector("${wpm.options.debug}",nil)
	inst.mDoDumpSelector = config.NewInjectionSelector("${wpm.options.dump}",nil)
	inst.mDoLogOptionsSelector = config.NewInjectionSelector("${wpm.options.log-options}",nil)
	inst.mDoRunWithGUISelector = config.NewInjectionSelector("${wpm.options.run-with-gui}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComInfoLoader) newObject() * boot0x8dc527.InfoLoader {
	return & boot0x8dc527.InfoLoader {}
}

func (inst * comFactory4pComInfoLoader) castObject(instance application.ComponentInstance) * boot0x8dc527.InfoLoader {
	return instance.Get().(*boot0x8dc527.InfoLoader)
}

func (inst * comFactory4pComInfoLoader) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComInfoLoader) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComInfoLoader) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComInfoLoader) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComInfoLoader) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComInfoLoader) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AC = inst.getterForFieldACSelector(context)
	obj.DoBackupExe = inst.getterForFieldDoBackupExeSelector(context)
	obj.DoDebug = inst.getterForFieldDoDebugSelector(context)
	obj.DoDump = inst.getterForFieldDoDumpSelector(context)
	obj.DoLogOptions = inst.getterForFieldDoLogOptionsSelector(context)
	obj.DoRunWithGUI = inst.getterForFieldDoRunWithGUISelector(context)
	return context.LastError()
}

//getterForFieldACSelector
func (inst * comFactory4pComInfoLoader) getterForFieldACSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldDoBackupExeSelector
func (inst * comFactory4pComInfoLoader) getterForFieldDoBackupExeSelector (context application.InstanceContext) bool {
    return inst.mDoBackupExeSelector.GetBool(context)
}

//getterForFieldDoDebugSelector
func (inst * comFactory4pComInfoLoader) getterForFieldDoDebugSelector (context application.InstanceContext) bool {
    return inst.mDoDebugSelector.GetBool(context)
}

//getterForFieldDoDumpSelector
func (inst * comFactory4pComInfoLoader) getterForFieldDoDumpSelector (context application.InstanceContext) bool {
    return inst.mDoDumpSelector.GetBool(context)
}

//getterForFieldDoLogOptionsSelector
func (inst * comFactory4pComInfoLoader) getterForFieldDoLogOptionsSelector (context application.InstanceContext) bool {
    return inst.mDoLogOptionsSelector.GetBool(context)
}

//getterForFieldDoRunWithGUISelector
func (inst * comFactory4pComInfoLoader) getterForFieldDoRunWithGUISelector (context application.InstanceContext) bool {
    return inst.mDoRunWithGUISelector.GetBool(context)
}




