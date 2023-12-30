package main4wpm
import (
    pd894cb5e7 "github.com/bitwormhole/wpm/src/main/golang/code"
     "github.com/starter-go/application"
)

// type pd894cb5e7.Example in package:github.com/bitwormhole/wpm/src/main/golang/code
//
// id:com-d894cb5e7ef0c5a4-code-Example
// class:
// alias:
// scope:singleton
//
type pd894cb5e7e_code_Example struct {
}

func (inst* pd894cb5e7e_code_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d894cb5e7ef0c5a4-code-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd894cb5e7e_code_Example) new() any {
    return &pd894cb5e7.Example{}
}

func (inst* pd894cb5e7e_code_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd894cb5e7.Example)
	nop(ie, com)

	


    return nil
}


