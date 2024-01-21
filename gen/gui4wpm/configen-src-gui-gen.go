package gui4wpm
import (
    p0f9140652 "github.com/bitwormhole/wpm/gui"
     "github.com/starter-go/application"
)

// type p0f9140652.Example in package:github.com/bitwormhole/wpm/gui
//
// id:com-0f91406523dc70d7-gui-Example
// class:
// alias:
// scope:singleton
//
type p0f91406523_gui_Example struct {
}

func (inst* p0f91406523_gui_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0f91406523dc70d7-gui-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0f91406523_gui_Example) new() any {
    return &p0f9140652.Example{}
}

func (inst* p0f91406523_gui_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0f9140652.Example)
	nop(ie, com)

	


    return nil
}


