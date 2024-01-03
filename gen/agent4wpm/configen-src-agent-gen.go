package agent4wpm
import (
    p7c0694d50 "github.com/bitwormhole/wpm/agent"
    p0ef6f2938 "github.com/starter-go/application"
    pdea5a0f47 "github.com/starter-go/httpagent"
     "github.com/starter-go/application"
)

// type p7c0694d50.WPMAgent in package:github.com/bitwormhole/wpm/agent
//
// id:com-7c0694d502d27c9a-agent-WPMAgent
// class:
// alias:
// scope:singleton
//
type p7c0694d502_agent_WPMAgent struct {
}

func (inst* p7c0694d502_agent_WPMAgent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7c0694d502d27c9a-agent-WPMAgent"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7c0694d502_agent_WPMAgent) new() any {
    return &p7c0694d50.WPMAgent{}
}

func (inst* p7c0694d502_agent_WPMAgent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7c0694d50.WPMAgent)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.HTTP = inst.getHTTP(ie)
    com.ServerPort = inst.getServerPort(ie)


    return nil
}


func (inst*p7c0694d502_agent_WPMAgent) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p7c0694d502_agent_WPMAgent) getHTTP(ie application.InjectionExt)pdea5a0f47.Clients{
    return ie.GetComponent("#alias-dea5a0f47697e78c03558bf00bc7ff9c-Clients").(pdea5a0f47.Clients)
}


func (inst*p7c0694d502_agent_WPMAgent) getServerPort(ie application.InjectionExt)int{
    return ie.GetInt("${server.http.port}")
}


