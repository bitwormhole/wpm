package test4wpm
import (
    p6021e9d7f "github.com/bitwormhole/wpm/server/classes/contenttypes"
    p97b2b30ad "github.com/bitwormhole/wpm/server/classes/executables"
    p6db2388a6 "github.com/bitwormhole/wpm/server/classes/intenttemplates"
    p3c68bd3f6 "github.com/bitwormhole/wpm/server/classes/locations"
    p677240472 "github.com/bitwormhole/wpm/server/classes/media"
    pdb29a3bec "github.com/bitwormhole/wpm/server/classes/packages"
    paa275b61f "github.com/bitwormhole/wpm/server/classes/projects"
    p4a2c02e71 "github.com/bitwormhole/wpm/server/classes/repositories"
    p4ab661a34 "github.com/bitwormhole/wpm/server/classes/settings"
    p70bc378de "github.com/bitwormhole/wpm/src/test/golang/code"
     "github.com/starter-go/application"
)

// type p70bc378de.DebugAllDAOs in package:github.com/bitwormhole/wpm/src/test/golang/code
//
// id:com-70bc378dedb9cae3-code-DebugAllDAOs
// class:
// alias:
// scope:singleton
//
type p70bc378ded_code_DebugAllDAOs struct {
}

func (inst* p70bc378ded_code_DebugAllDAOs) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-70bc378dedb9cae3-code-DebugAllDAOs"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p70bc378ded_code_DebugAllDAOs) new() any {
    return &p70bc378de.DebugAllDAOs{}
}

func (inst* p70bc378ded_code_DebugAllDAOs) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p70bc378de.DebugAllDAOs)
	nop(ie, com)

	
    com.ContentTypeDAO = inst.getContentTypeDAO(ie)
    com.ExecutableDAO = inst.getExecutableDAO(ie)
    com.IntentTemplateDAO = inst.getIntentTemplateDAO(ie)
    com.LocationDAO = inst.getLocationDAO(ie)
    com.MediaDAO = inst.getMediaDAO(ie)
    com.ProjectDAO = inst.getProjectDAO(ie)
    com.LocalRepoDAO = inst.getLocalRepoDAO(ie)
    com.RemoteRepoDAO = inst.getRemoteRepoDAO(ie)
    com.SettingDAO = inst.getSettingDAO(ie)
    com.PackageDAO = inst.getPackageDAO(ie)


    return nil
}


func (inst*p70bc378ded_code_DebugAllDAOs) getContentTypeDAO(ie application.InjectionExt)p6021e9d7f.DAO{
    return ie.GetComponent("#alias-6021e9d7f5afc3355549f2f1fec8b3e7-DAO").(p6021e9d7f.DAO)
}


func (inst*p70bc378ded_code_DebugAllDAOs) getExecutableDAO(ie application.InjectionExt)p97b2b30ad.DAO{
    return ie.GetComponent("#alias-97b2b30ad7df904c32bf0f040e5527d8-DAO").(p97b2b30ad.DAO)
}


func (inst*p70bc378ded_code_DebugAllDAOs) getIntentTemplateDAO(ie application.InjectionExt)p6db2388a6.DAO{
    return ie.GetComponent("#alias-6db2388a63f83ff930eeb1226ef48eb6-DAO").(p6db2388a6.DAO)
}


func (inst*p70bc378ded_code_DebugAllDAOs) getLocationDAO(ie application.InjectionExt)p3c68bd3f6.DAO{
    return ie.GetComponent("#alias-3c68bd3f6d1fbc95a5a70cc9cbd6d4ea-DAO").(p3c68bd3f6.DAO)
}


func (inst*p70bc378ded_code_DebugAllDAOs) getMediaDAO(ie application.InjectionExt)p677240472.DAO{
    return ie.GetComponent("#alias-67724047202291d9335f729c0f271c46-DAO").(p677240472.DAO)
}


func (inst*p70bc378ded_code_DebugAllDAOs) getProjectDAO(ie application.InjectionExt)paa275b61f.DAO{
    return ie.GetComponent("#alias-aa275b61f0bcca40e77fdda2bbfc393c-DAO").(paa275b61f.DAO)
}


func (inst*p70bc378ded_code_DebugAllDAOs) getLocalRepoDAO(ie application.InjectionExt)p4a2c02e71.LocalRepositoryDAO{
    return ie.GetComponent("#alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryDAO").(p4a2c02e71.LocalRepositoryDAO)
}


func (inst*p70bc378ded_code_DebugAllDAOs) getRemoteRepoDAO(ie application.InjectionExt)p4a2c02e71.RemoteRepositoryDAO{
    return ie.GetComponent("#alias-4a2c02e71bff15e74f38c4d3caa746b4-RemoteRepositoryDAO").(p4a2c02e71.RemoteRepositoryDAO)
}


func (inst*p70bc378ded_code_DebugAllDAOs) getSettingDAO(ie application.InjectionExt)p4ab661a34.DAO{
    return ie.GetComponent("#alias-4ab661a34cf8539ff7a5d51208b7f52a-DAO").(p4ab661a34.DAO)
}


func (inst*p70bc378ded_code_DebugAllDAOs) getPackageDAO(ie application.InjectionExt)pdb29a3bec.DAO{
    return ie.GetComponent("#alias-db29a3bec1e3d63283b5b71ee9ef4989-DAO").(pdb29a3bec.DAO)
}



// type p70bc378de.DemoUnit in package:github.com/bitwormhole/wpm/src/test/golang/code
//
// id:com-70bc378dedb9cae3-code-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p70bc378ded_code_DemoUnit struct {
}

func (inst* p70bc378ded_code_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-70bc378dedb9cae3-code-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p70bc378ded_code_DemoUnit) new() any {
    return &p70bc378de.DemoUnit{}
}

func (inst* p70bc378ded_code_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p70bc378de.DemoUnit)
	nop(ie, com)

	


    return nil
}


