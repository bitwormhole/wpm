package test4wpm
import (
    p70bc378de "github.com/bitwormhole/wpm/src/test/golang/code"
     "github.com/starter-go/application"
)

// type p70bc378de.DemoUnit in package:github.com/bitwormhole/wpm/src/test/golang/code
//
// id:com-70bc378dedb9cae3-code-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p70bc378ded_code_DemoUnit struct {
}

func (inst* p70bc378ded_code_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-70bc378dedb9cae3-code-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p70bc378ded_code_DemoUnit) new() any {
    return &p70bc378de.DemoUnit{}
}

func (inst* p70bc378ded_code_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p70bc378de.DemoUnit)
	nop(ie, com)

	


    return nil
}


