package boot4wpm
import (
    pd8d753d11 "github.com/bitwormhole/wpm/boot"
     "github.com/starter-go/application"
)

// type pd8d753d11.Example in package:github.com/bitwormhole/wpm/boot
//
// id:com-d8d753d11225b234-boot-Example
// class:
// alias:
// scope:singleton
//
type pd8d753d112_boot_Example struct {
}

func (inst* pd8d753d112_boot_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d8d753d11225b234-boot-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd8d753d112_boot_Example) new() any {
    return &pd8d753d11.Example{}
}

func (inst* pd8d753d112_boot_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd8d753d11.Example)
	nop(ie, com)

	


    return nil
}


