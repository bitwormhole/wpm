package agent4wpm
import (
    p7c0694d50 "github.com/bitwormhole/wpm/agent"
     "github.com/starter-go/application"
)

// type p7c0694d50.Example in package:github.com/bitwormhole/wpm/agent
//
// id:com-7c0694d502d27c9a-agent-Example
// class:
// alias:
// scope:singleton
//
type p7c0694d502_agent_Example struct {
}

func (inst* p7c0694d502_agent_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7c0694d502d27c9a-agent-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7c0694d502_agent_Example) new() any {
    return &p7c0694d50.Example{}
}

func (inst* p7c0694d502_agent_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7c0694d50.Example)
	nop(ie, com)

	


    return nil
}


