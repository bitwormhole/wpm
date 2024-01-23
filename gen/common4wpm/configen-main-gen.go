package common4wpm

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p09212bb029_common_Example{})
    inst.register(&pbd5ab6f4e0_iabout_ServiceImpl{})
    inst.register(&pe0ab501628_ienv_EnvironmentImpl{})
    inst.register(&pe4ab0a3957_ibuckets_MediaBucketPool{})


    return nil
}
