package common4wpm
import (
    p09212bb02 "github.com/bitwormhole/wpm/common"
    p534481b92 "github.com/bitwormhole/wpm/common/implements/impabout"
    p0d2a11d16 "github.com/starter-go/afs"
    p0ef6f2938 "github.com/starter-go/application"
     "github.com/starter-go/application"
)

// type p09212bb02.Example in package:github.com/bitwormhole/wpm/common
//
// id:com-09212bb029144995-common-Example
// class:
// alias:
// scope:singleton
//
type p09212bb029_common_Example struct {
}

func (inst* p09212bb029_common_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-09212bb029144995-common-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p09212bb029_common_Example) new() any {
    return &p09212bb02.Example{}
}

func (inst* p09212bb029_common_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p09212bb02.Example)
	nop(ie, com)

	


    return nil
}



// type p534481b92.ServiceImpl in package:github.com/bitwormhole/wpm/common/implements/impabout
//
// id:com-534481b9200d7ab5-impabout-ServiceImpl
// class:
// alias:alias-663158af9b1a72a79c266e906db07157-Service
// scope:singleton
//
type p534481b920_impabout_ServiceImpl struct {
}

func (inst* p534481b920_impabout_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-534481b9200d7ab5-impabout-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-663158af9b1a72a79c266e906db07157-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p534481b920_impabout_ServiceImpl) new() any {
    return &p534481b92.ServiceImpl{}
}

func (inst* p534481b920_impabout_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p534481b92.ServiceImpl)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)
    com.AC = inst.getAC(ie)
    com.Name = inst.getName(ie)
    com.Title = inst.getTitle(ie)
    com.Copyright = inst.getCopyright(ie)
    com.Profile = inst.getProfile(ie)
    com.ServerName = inst.getServerName(ie)
    com.ServerHost = inst.getServerHost(ie)
    com.ServerPort = inst.getServerPort(ie)
    com.ServerProtocol = inst.getServerProtocol(ie)


    return nil
}


func (inst*p534481b920_impabout_ServiceImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


func (inst*p534481b920_impabout_ServiceImpl) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p534481b920_impabout_ServiceImpl) getName(ie application.InjectionExt)string{
    return ie.GetString("${application.name}")
}


func (inst*p534481b920_impabout_ServiceImpl) getTitle(ie application.InjectionExt)string{
    return ie.GetString("${application.title}")
}


func (inst*p534481b920_impabout_ServiceImpl) getCopyright(ie application.InjectionExt)string{
    return ie.GetString("${application.copyright}")
}


func (inst*p534481b920_impabout_ServiceImpl) getProfile(ie application.InjectionExt)string{
    return ie.GetString("${application.profiles.active}")
}


func (inst*p534481b920_impabout_ServiceImpl) getServerName(ie application.InjectionExt)string{
    return ie.GetString("${server.default}")
}


func (inst*p534481b920_impabout_ServiceImpl) getServerHost(ie application.InjectionExt)string{
    return ie.GetString("${server.host}")
}


func (inst*p534481b920_impabout_ServiceImpl) getServerPort(ie application.InjectionExt)int{
    return ie.GetInt("${server.port}")
}


func (inst*p534481b920_impabout_ServiceImpl) getServerProtocol(ie application.InjectionExt)string{
    return ie.GetString("${server.protocol}")
}


