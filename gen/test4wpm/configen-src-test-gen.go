package test4wpm
import (
    p0a3016105 "github.com/bitwormhole/wpm/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type p0a3016105.DemoUnit in package:github.com/bitwormhole/wpm/src/test/golang/unit
//
// id:com-0a30161050d8351a-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p0a30161050_unit_DemoUnit struct {
}

func (inst* p0a30161050_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0a30161050d8351a-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0a30161050_unit_DemoUnit) new() any {
    return &p0a3016105.DemoUnit{}
}

func (inst* p0a30161050_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0a3016105.DemoUnit)
	nop(ie, com)

	


    return nil
}


