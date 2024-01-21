package cli4wpm
import (
    p613e7ffd2 "github.com/bitwormhole/wpm/cli"
     "github.com/starter-go/application"
)

// type p613e7ffd2.Example in package:github.com/bitwormhole/wpm/cli
//
// id:com-613e7ffd2c3d3925-cli-Example
// class:
// alias:
// scope:singleton
//
type p613e7ffd2c_cli_Example struct {
}

func (inst* p613e7ffd2c_cli_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-613e7ffd2c3d3925-cli-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p613e7ffd2c_cli_Example) new() any {
    return &p613e7ffd2.Example{}
}

func (inst* p613e7ffd2c_cli_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p613e7ffd2.Example)
	nop(ie, com)

	


    return nil
}


