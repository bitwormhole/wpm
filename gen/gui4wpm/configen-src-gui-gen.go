package gui4wpm
import (
    p663158af9 "github.com/bitwormhole/wpm/common/classes/about"
    p0f9140652 "github.com/bitwormhole/wpm/gui"
    pa1bfd2596 "github.com/starter-go/browsers"
     "github.com/starter-go/application"
)

// type p0f9140652.BrowserLauncher in package:github.com/bitwormhole/wpm/gui
//
// id:com-0f91406523dc70d7-gui-BrowserLauncher
// class:
// alias:
// scope:singleton
//
type p0f91406523_gui_BrowserLauncher struct {
}

func (inst* p0f91406523_gui_BrowserLauncher) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0f91406523dc70d7-gui-BrowserLauncher"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0f91406523_gui_BrowserLauncher) new() any {
    return &p0f9140652.BrowserLauncher{}
}

func (inst* p0f91406523_gui_BrowserLauncher) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0f9140652.BrowserLauncher)
	nop(ie, com)

	
    com.WebService = inst.getWebService(ie)
    com.AboutService = inst.getAboutService(ie)


    return nil
}


func (inst*p0f91406523_gui_BrowserLauncher) getWebService(ie application.InjectionExt)pa1bfd2596.Service{
    return ie.GetComponent("#alias-a1bfd25967b1d55d99511da18bc92d78-Service").(pa1bfd2596.Service)
}


func (inst*p0f91406523_gui_BrowserLauncher) getAboutService(ie application.InjectionExt)p663158af9.Service{
    return ie.GetComponent("#alias-663158af9b1a72a79c266e906db07157-Service").(p663158af9.Service)
}


