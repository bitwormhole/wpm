package main4wpm

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

    
    inst.register(&p213a6f747a_idao_LocationDaoImpl{})
    inst.register(&p57f9b38391_apiv1_ExampleController{})
    inst.register(&p57f9b38391_apiv1_LocalRepositoryController{})
    inst.register(&p636ab1bcf1_example_DaoImpl{})
    inst.register(&p8aac3fe082_database_ThisGroup{})
    inst.register(&pb25d04dc62_ilocations_LocationServiceImpl{})
    inst.register(&pd7f15ffc37_irepositories_LocalRepositoryServiceImpl{})


    return nil
}
