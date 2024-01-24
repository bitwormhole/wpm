package common4wpm
import (
    pae1817f17 "github.com/bitwormhole/wpm"
    p09212bb02 "github.com/bitwormhole/wpm/common"
    pbd5ab6f4e "github.com/bitwormhole/wpm/common/implements/iabout"
    pe4ab0a395 "github.com/bitwormhole/wpm/common/implements/ibuckets"
    p7244b39c6 "github.com/bitwormhole/wpm/common/implements/icaches"
    pe0ab50162 "github.com/bitwormhole/wpm/common/implements/ienv"
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



// type pbd5ab6f4e.ServiceImpl in package:github.com/bitwormhole/wpm/common/implements/iabout
//
// id:com-bd5ab6f4e02a4c67-iabout-ServiceImpl
// class:
// alias:alias-663158af9b1a72a79c266e906db07157-Service
// scope:singleton
//
type pbd5ab6f4e0_iabout_ServiceImpl struct {
}

func (inst* pbd5ab6f4e0_iabout_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-bd5ab6f4e02a4c67-iabout-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-663158af9b1a72a79c266e906db07157-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbd5ab6f4e0_iabout_ServiceImpl) new() any {
    return &pbd5ab6f4e.ServiceImpl{}
}

func (inst* pbd5ab6f4e0_iabout_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbd5ab6f4e.ServiceImpl)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.Env = inst.getEnv(ie)
    com.FS = inst.getFS(ie)
    com.Name = inst.getName(ie)
    com.Title = inst.getTitle(ie)
    com.Copyright = inst.getCopyright(ie)
    com.Profile = inst.getProfile(ie)


    return nil
}


func (inst*pbd5ab6f4e0_iabout_ServiceImpl) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pbd5ab6f4e0_iabout_ServiceImpl) getEnv(ie application.InjectionExt)pae1817f17.Environment{
    return ie.GetComponent("#alias-ae1817f178ee7043d385955ec6d6a87b-Environment").(pae1817f17.Environment)
}


func (inst*pbd5ab6f4e0_iabout_ServiceImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


func (inst*pbd5ab6f4e0_iabout_ServiceImpl) getName(ie application.InjectionExt)string{
    return ie.GetString("${application.name}")
}


func (inst*pbd5ab6f4e0_iabout_ServiceImpl) getTitle(ie application.InjectionExt)string{
    return ie.GetString("${application.title}")
}


func (inst*pbd5ab6f4e0_iabout_ServiceImpl) getCopyright(ie application.InjectionExt)string{
    return ie.GetString("${application.copyright}")
}


func (inst*pbd5ab6f4e0_iabout_ServiceImpl) getProfile(ie application.InjectionExt)string{
    return ie.GetString("${application.profiles.active}")
}



// type pe4ab0a395.MediaBucketPool in package:github.com/bitwormhole/wpm/common/implements/ibuckets
//
// id:com-e4ab0a39574adce4-ibuckets-MediaBucketPool
// class:
// alias:alias-9ae32fb866160a2b2e9745348187d238-BucketPool
// scope:singleton
//
type pe4ab0a3957_ibuckets_MediaBucketPool struct {
}

func (inst* pe4ab0a3957_ibuckets_MediaBucketPool) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e4ab0a39574adce4-ibuckets-MediaBucketPool"
	r.Classes = ""
	r.Aliases = "alias-9ae32fb866160a2b2e9745348187d238-BucketPool"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe4ab0a3957_ibuckets_MediaBucketPool) new() any {
    return &pe4ab0a395.MediaBucketPool{}
}

func (inst* pe4ab0a3957_ibuckets_MediaBucketPool) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe4ab0a395.MediaBucketPool)
	nop(ie, com)

	
    com.Env = inst.getEnv(ie)


    return nil
}


func (inst*pe4ab0a3957_ibuckets_MediaBucketPool) getEnv(ie application.InjectionExt)pae1817f17.Environment{
    return ie.GetComponent("#alias-ae1817f178ee7043d385955ec6d6a87b-Environment").(pae1817f17.Environment)
}



// type p7244b39c6.CacheServiceImpl in package:github.com/bitwormhole/wpm/common/implements/icaches
//
// id:com-7244b39c6dffc187-icaches-CacheServiceImpl
// class:
// alias:alias-765def9d27b81648c9336ef7a9348c53-Service
// scope:singleton
//
type p7244b39c6d_icaches_CacheServiceImpl struct {
}

func (inst* p7244b39c6d_icaches_CacheServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7244b39c6dffc187-icaches-CacheServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-765def9d27b81648c9336ef7a9348c53-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7244b39c6d_icaches_CacheServiceImpl) new() any {
    return &p7244b39c6.CacheServiceImpl{}
}

func (inst* p7244b39c6d_icaches_CacheServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7244b39c6.CacheServiceImpl)
	nop(ie, com)

	


    return nil
}



// type pe0ab50162.EnvironmentImpl in package:github.com/bitwormhole/wpm/common/implements/ienv
//
// id:com-e0ab501628e20b12-ienv-EnvironmentImpl
// class:
// alias:alias-ae1817f178ee7043d385955ec6d6a87b-Environment
// scope:singleton
//
type pe0ab501628_ienv_EnvironmentImpl struct {
}

func (inst* pe0ab501628_ienv_EnvironmentImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e0ab501628e20b12-ienv-EnvironmentImpl"
	r.Classes = ""
	r.Aliases = "alias-ae1817f178ee7043d385955ec6d6a87b-Environment"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe0ab501628_ienv_EnvironmentImpl) new() any {
    return &pe0ab50162.EnvironmentImpl{}
}

func (inst* pe0ab501628_ienv_EnvironmentImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe0ab50162.EnvironmentImpl)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)
    com.ServerPort = inst.getServerPort(ie)
    com.ServerPortNumMin = inst.getServerPortNumMin(ie)
    com.ServerPortNumMax = inst.getServerPortNumMax(ie)
    com.ServerName = inst.getServerName(ie)
    com.ServerHost = inst.getServerHost(ie)
    com.ServerProtocol = inst.getServerProtocol(ie)


    return nil
}


func (inst*pe0ab501628_ienv_EnvironmentImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


func (inst*pe0ab501628_ienv_EnvironmentImpl) getServerPort(ie application.InjectionExt)int{
    return ie.GetInt("${server.port}")
}


func (inst*pe0ab501628_ienv_EnvironmentImpl) getServerPortNumMin(ie application.InjectionExt)int{
    return ie.GetInt("${server.port}")
}


func (inst*pe0ab501628_ienv_EnvironmentImpl) getServerPortNumMax(ie application.InjectionExt)int{
    return ie.GetInt("${server.port.max}")
}


func (inst*pe0ab501628_ienv_EnvironmentImpl) getServerName(ie application.InjectionExt)string{
    return ie.GetString("${server.default}")
}


func (inst*pe0ab501628_ienv_EnvironmentImpl) getServerHost(ie application.InjectionExt)string{
    return ie.GetString("${server.host}")
}


func (inst*pe0ab501628_ienv_EnvironmentImpl) getServerProtocol(ie application.InjectionExt)string{
    return ie.GetString("${server.protocol}")
}

