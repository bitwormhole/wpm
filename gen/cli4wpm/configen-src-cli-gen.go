package cli4wpm
import (
    pabc442002 "github.com/bitwormhole/wpm/src/cli/golang/code"
     "github.com/starter-go/application"
)

// type pabc442002.Com in package:github.com/bitwormhole/wpm/src/cli/golang/code
//
// id:com-abc442002fb3b565-code-Com
// class:
// alias:
// scope:singleton
//
type pabc442002f_code_Com struct {
}

func (inst* pabc442002f_code_Com) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-abc442002fb3b565-code-Com"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pabc442002f_code_Com) new() any {
    return &pabc442002.Com{}
}

func (inst* pabc442002f_code_Com) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pabc442002.Com)
	nop(ie, com)

	


    return nil
}


