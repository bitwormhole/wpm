package gui4wpm
import (
    p7c219bb63 "github.com/bitwormhole/wpm/src/gui/golang/code"
     "github.com/starter-go/application"
)

// type p7c219bb63.Com in package:github.com/bitwormhole/wpm/src/gui/golang/code
//
// id:com-7c219bb636aa295a-code-Com
// class:
// alias:
// scope:singleton
//
type p7c219bb636_code_Com struct {
}

func (inst* p7c219bb636_code_Com) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7c219bb636aa295a-code-Com"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7c219bb636_code_Com) new() any {
    return &p7c219bb63.Com{}
}

func (inst* p7c219bb636_code_Com) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7c219bb63.Com)
	nop(ie, com)

	


    return nil
}


