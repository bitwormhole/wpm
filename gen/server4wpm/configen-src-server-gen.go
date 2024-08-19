package server4wpm
import (
    paeb460c7d "github.com/bitwormhole/gitlib"
    p663158af9 "github.com/bitwormhole/wpm/common/classes/about"
    p9ae32fb86 "github.com/bitwormhole/wpm/common/classes/buckets"
    p765def9d2 "github.com/bitwormhole/wpm/common/classes/caches"
    p080073581 "github.com/bitwormhole/wpm/common/objects/dxo"
    pce44af903 "github.com/bitwormhole/wpm/server/classes/auths"
    p7b01405af "github.com/bitwormhole/wpm/server/classes/backups"
    p6021e9d7f "github.com/bitwormhole/wpm/server/classes/contenttypes"
    pcc7a88d45 "github.com/bitwormhole/wpm/server/classes/examples"
    p97b2b30ad "github.com/bitwormhole/wpm/server/classes/executables"
    peafe07768 "github.com/bitwormhole/wpm/server/classes/filequery"
    p709b0834a "github.com/bitwormhole/wpm/server/classes/intents"
    p6db2388a6 "github.com/bitwormhole/wpm/server/classes/intenttemplates"
    p3c68bd3f6 "github.com/bitwormhole/wpm/server/classes/locations"
    p677240472 "github.com/bitwormhole/wpm/server/classes/media"
    pdb29a3bec "github.com/bitwormhole/wpm/server/classes/packages"
    paa275b61f "github.com/bitwormhole/wpm/server/classes/projects"
    p4a2c02e71 "github.com/bitwormhole/wpm/server/classes/repositories"
    p53a0e6a53 "github.com/bitwormhole/wpm/server/classes/repositoryworktreeproject"
    p353d73ec3 "github.com/bitwormhole/wpm/server/classes/softwaresets"
    p3320c87fb "github.com/bitwormhole/wpm/server/classes/statistics"
    p69e560419 "github.com/bitwormhole/wpm/server/classes/treeroots"
    p3ededc14a "github.com/bitwormhole/wpm/server/data/database"
    pfba77a8fb "github.com/bitwormhole/wpm/server/implements/example"
    pfcb564d0d "github.com/bitwormhole/wpm/server/implements/iauths"
    p0d7c61d39 "github.com/bitwormhole/wpm/server/implements/ibackups"
    pa4eaf0bff "github.com/bitwormhole/wpm/server/implements/icontenttypes"
    pbaa181ebe "github.com/bitwormhole/wpm/server/implements/iexecutables"
    p9b527161a "github.com/bitwormhole/wpm/server/implements/ifilequery"
    p42fcb1358 "github.com/bitwormhole/wpm/server/implements/iintentqueues"
    pa521fb5ff "github.com/bitwormhole/wpm/server/implements/iintenttemplates"
    p3220b15b1 "github.com/bitwormhole/wpm/server/implements/ilocations"
    p35246d3d2 "github.com/bitwormhole/wpm/server/implements/imedia"
    p7974a74f9 "github.com/bitwormhole/wpm/server/implements/iprojects"
    pac9dcde88 "github.com/bitwormhole/wpm/server/implements/irepositories"
    p853e9edea "github.com/bitwormhole/wpm/server/implements/irepositoryworktreeproject"
    p17f7c6f87 "github.com/bitwormhole/wpm/server/implements/isettings"
    pf04311b57 "github.com/bitwormhole/wpm/server/implements/isoftware"
    pa50f35ce6 "github.com/bitwormhole/wpm/server/implements/istatistics"
    p3e97fbda8 "github.com/bitwormhole/wpm/server/implements/itreeroots"
    p72b9b4e1d "github.com/bitwormhole/wpm/server/web/controllers"
    p415d59ee4 "github.com/bitwormhole/wpm/server/web/controllers/admin"
    p0d2a11d16 "github.com/starter-go/afs"
    p0ef6f2938 "github.com/starter-go/application"
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
    p9621e8b71 "github.com/starter-go/security/random"
     "github.com/starter-go/application"
)

// type p3ededc14a.ThisGroup in package:github.com/bitwormhole/wpm/server/data/database
//
// id:com-3ededc14a5c38f08-database-ThisGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:alias-08007358192e7a966d957c7686a43d06-DatabaseAgent
// scope:singleton
//
type p3ededc14a5_database_ThisGroup struct {
}

func (inst* p3ededc14a5_database_ThisGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3ededc14a5c38f08-database-ThisGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = "alias-08007358192e7a966d957c7686a43d06-DatabaseAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3ededc14a5_database_ThisGroup) new() any {
    return &p3ededc14a.ThisGroup{}
}

func (inst* p3ededc14a5_database_ThisGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3ededc14a.ThisGroup)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Alias = inst.getAlias(ie)
    com.URI = inst.getURI(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)
    com.SourceManager = inst.getSourceManager(ie)


    return nil
}


func (inst*p3ededc14a5_database_ThisGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.wpm.enabled}")
}


func (inst*p3ededc14a5_database_ThisGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.wpm.alias}")
}


func (inst*p3ededc14a5_database_ThisGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.wpm.uri}")
}


func (inst*p3ededc14a5_database_ThisGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.wpm.table-name-prefix}")
}


func (inst*p3ededc14a5_database_ThisGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.wpm.datasource}")
}


func (inst*p3ededc14a5_database_ThisGroup) getSourceManager(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}



// type pfba77a8fb.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/example
//
// id:com-fba77a8fb5da6105-example-DaoImpl
// class:
// alias:alias-cc7a88d4518f4d0c9704940596344e7e-DAO
// scope:singleton
//
type pfba77a8fb5_example_DaoImpl struct {
}

func (inst* pfba77a8fb5_example_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fba77a8fb5da6105-example-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-cc7a88d4518f4d0c9704940596344e7e-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfba77a8fb5_example_DaoImpl) new() any {
    return &pfba77a8fb.DaoImpl{}
}

func (inst* pfba77a8fb5_example_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfba77a8fb.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pfba77a8fb5_example_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*pfba77a8fb5_example_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pfba77a8fb.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/example
//
// id:com-fba77a8fb5da6105-example-ServiceImpl
// class:
// alias:alias-cc7a88d4518f4d0c9704940596344e7e-Service
// scope:singleton
//
type pfba77a8fb5_example_ServiceImpl struct {
}

func (inst* pfba77a8fb5_example_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fba77a8fb5da6105-example-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-cc7a88d4518f4d0c9704940596344e7e-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfba77a8fb5_example_ServiceImpl) new() any {
    return &pfba77a8fb.ServiceImpl{}
}

func (inst* pfba77a8fb5_example_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfba77a8fb.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pfba77a8fb5_example_ServiceImpl) getDao(ie application.InjectionExt)pcc7a88d45.DAO{
    return ie.GetComponent("#alias-cc7a88d4518f4d0c9704940596344e7e-DAO").(pcc7a88d45.DAO)
}



// type pfcb564d0d.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/iauths
//
// id:com-fcb564d0dfb58c27-iauths-DaoImpl
// class:
// alias:alias-ce44af9033c872f977eba59d8437c564-DAO
// scope:singleton
//
type pfcb564d0df_iauths_DaoImpl struct {
}

func (inst* pfcb564d0df_iauths_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fcb564d0dfb58c27-iauths-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-ce44af9033c872f977eba59d8437c564-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfcb564d0df_iauths_DaoImpl) new() any {
    return &pfcb564d0d.DaoImpl{}
}

func (inst* pfcb564d0df_iauths_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfcb564d0d.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pfcb564d0df_iauths_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*pfcb564d0df_iauths_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pfcb564d0d.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/iauths
//
// id:com-fcb564d0dfb58c27-iauths-ServiceImpl
// class:
// alias:alias-ce44af9033c872f977eba59d8437c564-Service
// scope:singleton
//
type pfcb564d0df_iauths_ServiceImpl struct {
}

func (inst* pfcb564d0df_iauths_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fcb564d0dfb58c27-iauths-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-ce44af9033c872f977eba59d8437c564-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfcb564d0df_iauths_ServiceImpl) new() any {
    return &pfcb564d0d.ServiceImpl{}
}

func (inst* pfcb564d0df_iauths_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfcb564d0d.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pfcb564d0df_iauths_ServiceImpl) getDao(ie application.InjectionExt)pce44af903.DAO{
    return ie.GetComponent("#alias-ce44af9033c872f977eba59d8437c564-DAO").(pce44af903.DAO)
}



// type p0d7c61d39.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/ibackups
//
// id:com-0d7c61d3987c67f2-ibackups-DaoImpl
// class:
// alias:alias-7b01405afb0538c075c5775b29d3562e-DAO
// scope:singleton
//
type p0d7c61d398_ibackups_DaoImpl struct {
}

func (inst* p0d7c61d398_ibackups_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0d7c61d3987c67f2-ibackups-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-7b01405afb0538c075c5775b29d3562e-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0d7c61d398_ibackups_DaoImpl) new() any {
    return &p0d7c61d39.DaoImpl{}
}

func (inst* p0d7c61d398_ibackups_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0d7c61d39.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p0d7c61d398_ibackups_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*p0d7c61d398_ibackups_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p0d7c61d39.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/ibackups
//
// id:com-0d7c61d3987c67f2-ibackups-ServiceImpl
// class:
// alias:alias-7b01405afb0538c075c5775b29d3562e-Service
// scope:singleton
//
type p0d7c61d398_ibackups_ServiceImpl struct {
}

func (inst* p0d7c61d398_ibackups_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0d7c61d3987c67f2-ibackups-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-7b01405afb0538c075c5775b29d3562e-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0d7c61d398_ibackups_ServiceImpl) new() any {
    return &p0d7c61d39.ServiceImpl{}
}

func (inst* p0d7c61d398_ibackups_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0d7c61d39.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p0d7c61d398_ibackups_ServiceImpl) getDao(ie application.InjectionExt)p7b01405af.DAO{
    return ie.GetComponent("#alias-7b01405afb0538c075c5775b29d3562e-DAO").(p7b01405af.DAO)
}



// type pa4eaf0bff.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/icontenttypes
//
// id:com-a4eaf0bfffcbf446-icontenttypes-DaoImpl
// class:
// alias:alias-6021e9d7f5afc3355549f2f1fec8b3e7-DAO
// scope:singleton
//
type pa4eaf0bfff_icontenttypes_DaoImpl struct {
}

func (inst* pa4eaf0bfff_icontenttypes_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a4eaf0bfffcbf446-icontenttypes-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-6021e9d7f5afc3355549f2f1fec8b3e7-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa4eaf0bfff_icontenttypes_DaoImpl) new() any {
    return &pa4eaf0bff.DaoImpl{}
}

func (inst* pa4eaf0bfff_icontenttypes_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa4eaf0bff.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pa4eaf0bfff_icontenttypes_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*pa4eaf0bfff_icontenttypes_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pa4eaf0bff.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/icontenttypes
//
// id:com-a4eaf0bfffcbf446-icontenttypes-ServiceImpl
// class:
// alias:alias-6021e9d7f5afc3355549f2f1fec8b3e7-Service
// scope:singleton
//
type pa4eaf0bfff_icontenttypes_ServiceImpl struct {
}

func (inst* pa4eaf0bfff_icontenttypes_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a4eaf0bfffcbf446-icontenttypes-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-6021e9d7f5afc3355549f2f1fec8b3e7-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa4eaf0bfff_icontenttypes_ServiceImpl) new() any {
    return &pa4eaf0bff.ServiceImpl{}
}

func (inst* pa4eaf0bfff_icontenttypes_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa4eaf0bff.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pa4eaf0bfff_icontenttypes_ServiceImpl) getDao(ie application.InjectionExt)p6021e9d7f.DAO{
    return ie.GetComponent("#alias-6021e9d7f5afc3355549f2f1fec8b3e7-DAO").(p6021e9d7f.DAO)
}



// type pbaa181ebe.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/iexecutables
//
// id:com-baa181ebecab5f84-iexecutables-DaoImpl
// class:
// alias:alias-97b2b30ad7df904c32bf0f040e5527d8-DAO
// scope:singleton
//
type pbaa181ebec_iexecutables_DaoImpl struct {
}

func (inst* pbaa181ebec_iexecutables_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-baa181ebecab5f84-iexecutables-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-97b2b30ad7df904c32bf0f040e5527d8-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbaa181ebec_iexecutables_DaoImpl) new() any {
    return &pbaa181ebe.DaoImpl{}
}

func (inst* pbaa181ebec_iexecutables_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbaa181ebe.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pbaa181ebec_iexecutables_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*pbaa181ebec_iexecutables_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pbaa181ebe.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/iexecutables
//
// id:com-baa181ebecab5f84-iexecutables-ServiceImpl
// class:
// alias:alias-97b2b30ad7df904c32bf0f040e5527d8-Service
// scope:singleton
//
type pbaa181ebec_iexecutables_ServiceImpl struct {
}

func (inst* pbaa181ebec_iexecutables_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-baa181ebecab5f84-iexecutables-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-97b2b30ad7df904c32bf0f040e5527d8-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbaa181ebec_iexecutables_ServiceImpl) new() any {
    return &pbaa181ebe.ServiceImpl{}
}

func (inst* pbaa181ebec_iexecutables_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbaa181ebe.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.StateLoader = inst.getStateLoader(ie)
    com.FS = inst.getFS(ie)


    return nil
}


func (inst*pbaa181ebec_iexecutables_ServiceImpl) getDao(ie application.InjectionExt)p97b2b30ad.DAO{
    return ie.GetComponent("#alias-97b2b30ad7df904c32bf0f040e5527d8-DAO").(p97b2b30ad.DAO)
}


func (inst*pbaa181ebec_iexecutables_ServiceImpl) getStateLoader(ie application.InjectionExt)p97b2b30ad.StateLoader{
    return ie.GetComponent("#alias-97b2b30ad7df904c32bf0f040e5527d8-StateLoader").(p97b2b30ad.StateLoader)
}


func (inst*pbaa181ebec_iexecutables_ServiceImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}



// type pbaa181ebe.ExecutableStateLoaderImpl in package:github.com/bitwormhole/wpm/server/implements/iexecutables
//
// id:com-baa181ebecab5f84-iexecutables-ExecutableStateLoaderImpl
// class:
// alias:alias-97b2b30ad7df904c32bf0f040e5527d8-StateLoader
// scope:singleton
//
type pbaa181ebec_iexecutables_ExecutableStateLoaderImpl struct {
}

func (inst* pbaa181ebec_iexecutables_ExecutableStateLoaderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-baa181ebecab5f84-iexecutables-ExecutableStateLoaderImpl"
	r.Classes = ""
	r.Aliases = "alias-97b2b30ad7df904c32bf0f040e5527d8-StateLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbaa181ebec_iexecutables_ExecutableStateLoaderImpl) new() any {
    return &pbaa181ebe.ExecutableStateLoaderImpl{}
}

func (inst* pbaa181ebec_iexecutables_ExecutableStateLoaderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbaa181ebe.ExecutableStateLoaderImpl)
	nop(ie, com)

	
    com.CacheService = inst.getCacheService(ie)
    com.FS = inst.getFS(ie)


    return nil
}


func (inst*pbaa181ebec_iexecutables_ExecutableStateLoaderImpl) getCacheService(ie application.InjectionExt)p765def9d2.Service{
    return ie.GetComponent("#alias-765def9d27b81648c9336ef7a9348c53-Service").(p765def9d2.Service)
}


func (inst*pbaa181ebec_iexecutables_ExecutableStateLoaderImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}



// type p9b527161a.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/ifilequery
//
// id:com-9b527161a50fcafd-ifilequery-ServiceImpl
// class:
// alias:alias-eafe077682a78bde34843d43b78b6a30-Service
// scope:singleton
//
type p9b527161a5_ifilequery_ServiceImpl struct {
}

func (inst* p9b527161a5_ifilequery_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9b527161a50fcafd-ifilequery-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-eafe077682a78bde34843d43b78b6a30-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9b527161a5_ifilequery_ServiceImpl) new() any {
    return &p9b527161a.ServiceImpl{}
}

func (inst* p9b527161a5_ifilequery_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9b527161a.ServiceImpl)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)
    com.CTypeService = inst.getCTypeService(ie)


    return nil
}


func (inst*p9b527161a5_ifilequery_ServiceImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


func (inst*p9b527161a5_ifilequery_ServiceImpl) getCTypeService(ie application.InjectionExt)p6021e9d7f.Service{
    return ie.GetComponent("#alias-6021e9d7f5afc3355549f2f1fec8b3e7-Service").(p6021e9d7f.Service)
}



// type p42fcb1358.IntentQueueService in package:github.com/bitwormhole/wpm/server/implements/iintentqueues
//
// id:com-42fcb1358a3cc9a4-iintentqueues-IntentQueueService
// class:
// alias:alias-709b0834aef1394ca72aa585afa48882-Queues
// scope:singleton
//
type p42fcb1358a_iintentqueues_IntentQueueService struct {
}

func (inst* p42fcb1358a_iintentqueues_IntentQueueService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-42fcb1358a3cc9a4-iintentqueues-IntentQueueService"
	r.Classes = ""
	r.Aliases = "alias-709b0834aef1394ca72aa585afa48882-Queues"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p42fcb1358a_iintentqueues_IntentQueueService) new() any {
    return &p42fcb1358.IntentQueueService{}
}

func (inst* p42fcb1358a_iintentqueues_IntentQueueService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p42fcb1358.IntentQueueService)
	nop(ie, com)

	


    return nil
}



// type pa521fb5ff.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/iintenttemplates
//
// id:com-a521fb5ff6f1a969-iintenttemplates-DaoImpl
// class:
// alias:alias-6db2388a63f83ff930eeb1226ef48eb6-DAO
// scope:singleton
//
type pa521fb5ff6_iintenttemplates_DaoImpl struct {
}

func (inst* pa521fb5ff6_iintenttemplates_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a521fb5ff6f1a969-iintenttemplates-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-6db2388a63f83ff930eeb1226ef48eb6-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa521fb5ff6_iintenttemplates_DaoImpl) new() any {
    return &pa521fb5ff.DaoImpl{}
}

func (inst* pa521fb5ff6_iintenttemplates_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa521fb5ff.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pa521fb5ff6_iintenttemplates_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*pa521fb5ff6_iintenttemplates_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pa521fb5ff.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/iintenttemplates
//
// id:com-a521fb5ff6f1a969-iintenttemplates-ServiceImpl
// class:
// alias:alias-6db2388a63f83ff930eeb1226ef48eb6-Service
// scope:singleton
//
type pa521fb5ff6_iintenttemplates_ServiceImpl struct {
}

func (inst* pa521fb5ff6_iintenttemplates_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a521fb5ff6f1a969-iintenttemplates-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-6db2388a63f83ff930eeb1226ef48eb6-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa521fb5ff6_iintenttemplates_ServiceImpl) new() any {
    return &pa521fb5ff.ServiceImpl{}
}

func (inst* pa521fb5ff6_iintenttemplates_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa521fb5ff.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pa521fb5ff6_iintenttemplates_ServiceImpl) getDao(ie application.InjectionExt)p6db2388a6.DAO{
    return ie.GetComponent("#alias-6db2388a63f83ff930eeb1226ef48eb6-DAO").(p6db2388a6.DAO)
}



// type p3220b15b1.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/ilocations
//
// id:com-3220b15b16d77d2d-ilocations-DaoImpl
// class:
// alias:alias-3c68bd3f6d1fbc95a5a70cc9cbd6d4ea-DAO
// scope:singleton
//
type p3220b15b16_ilocations_DaoImpl struct {
}

func (inst* p3220b15b16_ilocations_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3220b15b16d77d2d-ilocations-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-3c68bd3f6d1fbc95a5a70cc9cbd6d4ea-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3220b15b16_ilocations_DaoImpl) new() any {
    return &p3220b15b1.DaoImpl{}
}

func (inst* p3220b15b16_ilocations_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3220b15b1.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p3220b15b16_ilocations_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*p3220b15b16_ilocations_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p3220b15b1.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/ilocations
//
// id:com-3220b15b16d77d2d-ilocations-ServiceImpl
// class:
// alias:alias-3c68bd3f6d1fbc95a5a70cc9cbd6d4ea-Service
// scope:singleton
//
type p3220b15b16_ilocations_ServiceImpl struct {
}

func (inst* p3220b15b16_ilocations_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3220b15b16d77d2d-ilocations-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-3c68bd3f6d1fbc95a5a70cc9cbd6d4ea-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3220b15b16_ilocations_ServiceImpl) new() any {
    return &p3220b15b1.ServiceImpl{}
}

func (inst* p3220b15b16_ilocations_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3220b15b1.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p3220b15b16_ilocations_ServiceImpl) getDao(ie application.InjectionExt)p3c68bd3f6.DAO{
    return ie.GetComponent("#alias-3c68bd3f6d1fbc95a5a70cc9cbd6d4ea-DAO").(p3c68bd3f6.DAO)
}



// type p35246d3d2.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/imedia
//
// id:com-35246d3d244a735d-imedia-DaoImpl
// class:
// alias:alias-67724047202291d9335f729c0f271c46-DAO
// scope:singleton
//
type p35246d3d24_imedia_DaoImpl struct {
}

func (inst* p35246d3d24_imedia_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-35246d3d244a735d-imedia-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-67724047202291d9335f729c0f271c46-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p35246d3d24_imedia_DaoImpl) new() any {
    return &p35246d3d2.DaoImpl{}
}

func (inst* p35246d3d24_imedia_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p35246d3d2.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p35246d3d24_imedia_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*p35246d3d24_imedia_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p35246d3d2.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/imedia
//
// id:com-35246d3d244a735d-imedia-ServiceImpl
// class:
// alias:alias-67724047202291d9335f729c0f271c46-Service
// scope:singleton
//
type p35246d3d24_imedia_ServiceImpl struct {
}

func (inst* p35246d3d24_imedia_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-35246d3d244a735d-imedia-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-67724047202291d9335f729c0f271c46-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p35246d3d24_imedia_ServiceImpl) new() any {
    return &p35246d3d2.ServiceImpl{}
}

func (inst* p35246d3d24_imedia_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p35246d3d2.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.Buckets = inst.getBuckets(ie)


    return nil
}


func (inst*p35246d3d24_imedia_ServiceImpl) getDao(ie application.InjectionExt)p677240472.DAO{
    return ie.GetComponent("#alias-67724047202291d9335f729c0f271c46-DAO").(p677240472.DAO)
}


func (inst*p35246d3d24_imedia_ServiceImpl) getBuckets(ie application.InjectionExt)p9ae32fb86.BucketPool{
    return ie.GetComponent("#alias-9ae32fb866160a2b2e9745348187d238-BucketPool").(p9ae32fb86.BucketPool)
}



// type p7974a74f9.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/iprojects
//
// id:com-7974a74f95fe6a2a-iprojects-DaoImpl
// class:
// alias:alias-aa275b61f0bcca40e77fdda2bbfc393c-DAO
// scope:singleton
//
type p7974a74f95_iprojects_DaoImpl struct {
}

func (inst* p7974a74f95_iprojects_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7974a74f95fe6a2a-iprojects-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-aa275b61f0bcca40e77fdda2bbfc393c-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7974a74f95_iprojects_DaoImpl) new() any {
    return &p7974a74f9.DaoImpl{}
}

func (inst* p7974a74f95_iprojects_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7974a74f9.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p7974a74f95_iprojects_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*p7974a74f95_iprojects_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p7974a74f9.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/iprojects
//
// id:com-7974a74f95fe6a2a-iprojects-ServiceImpl
// class:
// alias:alias-aa275b61f0bcca40e77fdda2bbfc393c-Service
// scope:singleton
//
type p7974a74f95_iprojects_ServiceImpl struct {
}

func (inst* p7974a74f95_iprojects_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7974a74f95fe6a2a-iprojects-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-aa275b61f0bcca40e77fdda2bbfc393c-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7974a74f95_iprojects_ServiceImpl) new() any {
    return &p7974a74f9.ServiceImpl{}
}

func (inst* p7974a74f95_iprojects_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7974a74f9.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p7974a74f95_iprojects_ServiceImpl) getDao(ie application.InjectionExt)paa275b61f.DAO{
    return ie.GetComponent("#alias-aa275b61f0bcca40e77fdda2bbfc393c-DAO").(paa275b61f.DAO)
}



// type pac9dcde88.ImportRepositoryServiceImpl in package:github.com/bitwormhole/wpm/server/implements/irepositories
//
// id:com-ac9dcde88744146b-irepositories-ImportRepositoryServiceImpl
// class:
// alias:alias-4a2c02e71bff15e74f38c4d3caa746b4-ImportRepositoryService
// scope:singleton
//
type pac9dcde887_irepositories_ImportRepositoryServiceImpl struct {
}

func (inst* pac9dcde887_irepositories_ImportRepositoryServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ac9dcde88744146b-irepositories-ImportRepositoryServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-4a2c02e71bff15e74f38c4d3caa746b4-ImportRepositoryService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pac9dcde887_irepositories_ImportRepositoryServiceImpl) new() any {
    return &pac9dcde88.ImportRepositoryServiceImpl{}
}

func (inst* pac9dcde887_irepositories_ImportRepositoryServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pac9dcde88.ImportRepositoryServiceImpl)
	nop(ie, com)

	
    com.LibAgent = inst.getLibAgent(ie)
    com.RepoService = inst.getRepoService(ie)


    return nil
}


func (inst*pac9dcde887_irepositories_ImportRepositoryServiceImpl) getLibAgent(ie application.InjectionExt)paeb460c7d.Agent{
    return ie.GetComponent("#alias-aeb460c7d339df24b0b38a0d65e30102-Agent").(paeb460c7d.Agent)
}


func (inst*pac9dcde887_irepositories_ImportRepositoryServiceImpl) getRepoService(ie application.InjectionExt)p4a2c02e71.LocalRepositoryService{
    return ie.GetComponent("#alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryService").(p4a2c02e71.LocalRepositoryService)
}



// type pac9dcde88.LocalRepositoryDao in package:github.com/bitwormhole/wpm/server/implements/irepositories
//
// id:com-ac9dcde88744146b-irepositories-LocalRepositoryDao
// class:
// alias:alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryDAO
// scope:singleton
//
type pac9dcde887_irepositories_LocalRepositoryDao struct {
}

func (inst* pac9dcde887_irepositories_LocalRepositoryDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ac9dcde88744146b-irepositories-LocalRepositoryDao"
	r.Classes = ""
	r.Aliases = "alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pac9dcde887_irepositories_LocalRepositoryDao) new() any {
    return &pac9dcde88.LocalRepositoryDao{}
}

func (inst* pac9dcde887_irepositories_LocalRepositoryDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pac9dcde88.LocalRepositoryDao)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pac9dcde887_irepositories_LocalRepositoryDao) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*pac9dcde887_irepositories_LocalRepositoryDao) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pac9dcde88.LocalRepositoryService in package:github.com/bitwormhole/wpm/server/implements/irepositories
//
// id:com-ac9dcde88744146b-irepositories-LocalRepositoryService
// class:
// alias:alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryService
// scope:singleton
//
type pac9dcde887_irepositories_LocalRepositoryService struct {
}

func (inst* pac9dcde887_irepositories_LocalRepositoryService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ac9dcde88744146b-irepositories-LocalRepositoryService"
	r.Classes = ""
	r.Aliases = "alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pac9dcde887_irepositories_LocalRepositoryService) new() any {
    return &pac9dcde88.LocalRepositoryService{}
}

func (inst* pac9dcde887_irepositories_LocalRepositoryService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pac9dcde88.LocalRepositoryService)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.LibAgent = inst.getLibAgent(ie)


    return nil
}


func (inst*pac9dcde887_irepositories_LocalRepositoryService) getDao(ie application.InjectionExt)p4a2c02e71.LocalRepositoryDAO{
    return ie.GetComponent("#alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryDAO").(p4a2c02e71.LocalRepositoryDAO)
}


func (inst*pac9dcde887_irepositories_LocalRepositoryService) getLibAgent(ie application.InjectionExt)paeb460c7d.Agent{
    return ie.GetComponent("#alias-aeb460c7d339df24b0b38a0d65e30102-Agent").(paeb460c7d.Agent)
}



// type pac9dcde88.RemoteRepositoryDao in package:github.com/bitwormhole/wpm/server/implements/irepositories
//
// id:com-ac9dcde88744146b-irepositories-RemoteRepositoryDao
// class:
// alias:alias-4a2c02e71bff15e74f38c4d3caa746b4-RemoteRepositoryDAO
// scope:singleton
//
type pac9dcde887_irepositories_RemoteRepositoryDao struct {
}

func (inst* pac9dcde887_irepositories_RemoteRepositoryDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ac9dcde88744146b-irepositories-RemoteRepositoryDao"
	r.Classes = ""
	r.Aliases = "alias-4a2c02e71bff15e74f38c4d3caa746b4-RemoteRepositoryDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pac9dcde887_irepositories_RemoteRepositoryDao) new() any {
    return &pac9dcde88.RemoteRepositoryDao{}
}

func (inst* pac9dcde887_irepositories_RemoteRepositoryDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pac9dcde88.RemoteRepositoryDao)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pac9dcde887_irepositories_RemoteRepositoryDao) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*pac9dcde887_irepositories_RemoteRepositoryDao) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pac9dcde88.RemoteRepositoryService in package:github.com/bitwormhole/wpm/server/implements/irepositories
//
// id:com-ac9dcde88744146b-irepositories-RemoteRepositoryService
// class:
// alias:alias-4a2c02e71bff15e74f38c4d3caa746b4-RemoteRepositoryService
// scope:singleton
//
type pac9dcde887_irepositories_RemoteRepositoryService struct {
}

func (inst* pac9dcde887_irepositories_RemoteRepositoryService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ac9dcde88744146b-irepositories-RemoteRepositoryService"
	r.Classes = ""
	r.Aliases = "alias-4a2c02e71bff15e74f38c4d3caa746b4-RemoteRepositoryService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pac9dcde887_irepositories_RemoteRepositoryService) new() any {
    return &pac9dcde88.RemoteRepositoryService{}
}

func (inst* pac9dcde887_irepositories_RemoteRepositoryService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pac9dcde88.RemoteRepositoryService)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pac9dcde887_irepositories_RemoteRepositoryService) getDao(ie application.InjectionExt)pcc7a88d45.DAO{
    return ie.GetComponent("#alias-cc7a88d4518f4d0c9704940596344e7e-DAO").(pcc7a88d45.DAO)
}



// type p853e9edea.RepositoryWorktreeProjectServiceImpl in package:github.com/bitwormhole/wpm/server/implements/irepositoryworktreeproject
//
// id:com-853e9edeaaebf60b-irepositoryworktreeproject-RepositoryWorktreeProjectServiceImpl
// class:
// alias:alias-53a0e6a5358524b4150f206572d109d1-Service
// scope:singleton
//
type p853e9edeaa_irepositoryworktreeproject_RepositoryWorktreeProjectServiceImpl struct {
}

func (inst* p853e9edeaa_irepositoryworktreeproject_RepositoryWorktreeProjectServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-853e9edeaaebf60b-irepositoryworktreeproject-RepositoryWorktreeProjectServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-53a0e6a5358524b4150f206572d109d1-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p853e9edeaa_irepositoryworktreeproject_RepositoryWorktreeProjectServiceImpl) new() any {
    return &p853e9edea.RepositoryWorktreeProjectServiceImpl{}
}

func (inst* p853e9edeaa_irepositoryworktreeproject_RepositoryWorktreeProjectServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p853e9edea.RepositoryWorktreeProjectServiceImpl)
	nop(ie, com)

	
    com.GitLibAgent = inst.getGitLibAgent(ie)


    return nil
}


func (inst*p853e9edeaa_irepositoryworktreeproject_RepositoryWorktreeProjectServiceImpl) getGitLibAgent(ie application.InjectionExt)paeb460c7d.Agent{
    return ie.GetComponent("#alias-aeb460c7d339df24b0b38a0d65e30102-Agent").(paeb460c7d.Agent)
}



// type p17f7c6f87.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/isettings
//
// id:com-17f7c6f8755961b7-isettings-DaoImpl
// class:
// alias:alias-4ab661a34cf8539ff7a5d51208b7f52a-DAO
// scope:singleton
//
type p17f7c6f875_isettings_DaoImpl struct {
}

func (inst* p17f7c6f875_isettings_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-17f7c6f8755961b7-isettings-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-4ab661a34cf8539ff7a5d51208b7f52a-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p17f7c6f875_isettings_DaoImpl) new() any {
    return &p17f7c6f87.DaoImpl{}
}

func (inst* p17f7c6f875_isettings_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p17f7c6f87.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p17f7c6f875_isettings_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*p17f7c6f875_isettings_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p17f7c6f87.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/isettings
//
// id:com-17f7c6f8755961b7-isettings-ServiceImpl
// class:
// alias:alias-cc7a88d4518f4d0c9704940596344e7e-Service
// scope:singleton
//
type p17f7c6f875_isettings_ServiceImpl struct {
}

func (inst* p17f7c6f875_isettings_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-17f7c6f8755961b7-isettings-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-cc7a88d4518f4d0c9704940596344e7e-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p17f7c6f875_isettings_ServiceImpl) new() any {
    return &p17f7c6f87.ServiceImpl{}
}

func (inst* p17f7c6f875_isettings_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p17f7c6f87.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p17f7c6f875_isettings_ServiceImpl) getDao(ie application.InjectionExt)pcc7a88d45.DAO{
    return ie.GetComponent("#alias-cc7a88d4518f4d0c9704940596344e7e-DAO").(pcc7a88d45.DAO)
}



// type pf04311b57.SoftwarePackageDaoImpl in package:github.com/bitwormhole/wpm/server/implements/isoftware
//
// id:com-f04311b572f9a4e8-isoftware-SoftwarePackageDaoImpl
// class:
// alias:alias-db29a3bec1e3d63283b5b71ee9ef4989-DAO
// scope:singleton
//
type pf04311b572_isoftware_SoftwarePackageDaoImpl struct {
}

func (inst* pf04311b572_isoftware_SoftwarePackageDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f04311b572f9a4e8-isoftware-SoftwarePackageDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-db29a3bec1e3d63283b5b71ee9ef4989-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf04311b572_isoftware_SoftwarePackageDaoImpl) new() any {
    return &pf04311b57.SoftwarePackageDaoImpl{}
}

func (inst* pf04311b572_isoftware_SoftwarePackageDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf04311b57.SoftwarePackageDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pf04311b572_isoftware_SoftwarePackageDaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*pf04311b572_isoftware_SoftwarePackageDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pf04311b57.SoftwarePackageService in package:github.com/bitwormhole/wpm/server/implements/isoftware
//
// id:com-f04311b572f9a4e8-isoftware-SoftwarePackageService
// class:
// alias:alias-db29a3bec1e3d63283b5b71ee9ef4989-Service
// scope:singleton
//
type pf04311b572_isoftware_SoftwarePackageService struct {
}

func (inst* pf04311b572_isoftware_SoftwarePackageService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f04311b572f9a4e8-isoftware-SoftwarePackageService"
	r.Classes = ""
	r.Aliases = "alias-db29a3bec1e3d63283b5b71ee9ef4989-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf04311b572_isoftware_SoftwarePackageService) new() any {
    return &pf04311b57.SoftwarePackageService{}
}

func (inst* pf04311b572_isoftware_SoftwarePackageService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf04311b57.SoftwarePackageService)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pf04311b572_isoftware_SoftwarePackageService) getDao(ie application.InjectionExt)pdb29a3bec.DAO{
    return ie.GetComponent("#alias-db29a3bec1e3d63283b5b71ee9ef4989-DAO").(pdb29a3bec.DAO)
}



// type pf04311b57.SoftwareSetService in package:github.com/bitwormhole/wpm/server/implements/isoftware
//
// id:com-f04311b572f9a4e8-isoftware-SoftwareSetService
// class:
// alias:alias-353d73ec3edc2e74f8181282d9e87a31-Service
// scope:singleton
//
type pf04311b572_isoftware_SoftwareSetService struct {
}

func (inst* pf04311b572_isoftware_SoftwareSetService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f04311b572f9a4e8-isoftware-SoftwareSetService"
	r.Classes = ""
	r.Aliases = "alias-353d73ec3edc2e74f8181282d9e87a31-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf04311b572_isoftware_SoftwareSetService) new() any {
    return &pf04311b57.SoftwareSetService{}
}

func (inst* pf04311b572_isoftware_SoftwareSetService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf04311b57.SoftwareSetService)
	nop(ie, com)

	
    com.PackageService = inst.getPackageService(ie)


    return nil
}


func (inst*pf04311b572_isoftware_SoftwareSetService) getPackageService(ie application.InjectionExt)pdb29a3bec.Service{
    return ie.GetComponent("#alias-db29a3bec1e3d63283b5b71ee9ef4989-Service").(pdb29a3bec.Service)
}



// type pa50f35ce6.StatisticServiceImpl in package:github.com/bitwormhole/wpm/server/implements/istatistics
//
// id:com-a50f35ce65518e9f-istatistics-StatisticServiceImpl
// class:
// alias:alias-3320c87fb39933a17055dcfd34498251-Service
// scope:singleton
//
type pa50f35ce65_istatistics_StatisticServiceImpl struct {
}

func (inst* pa50f35ce65_istatistics_StatisticServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a50f35ce65518e9f-istatistics-StatisticServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-3320c87fb39933a17055dcfd34498251-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa50f35ce65_istatistics_StatisticServiceImpl) new() any {
    return &pa50f35ce6.StatisticServiceImpl{}
}

func (inst* pa50f35ce65_istatistics_StatisticServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa50f35ce6.StatisticServiceImpl)
	nop(ie, com)

	


    return nil
}



// type p3e97fbda8.TreeRootDaoImpl in package:github.com/bitwormhole/wpm/server/implements/itreeroots
//
// id:com-3e97fbda8f36e61e-itreeroots-TreeRootDaoImpl
// class:
// alias:alias-69e5604197df6fca537ce9f81594a863-DAO
// scope:singleton
//
type p3e97fbda8f_itreeroots_TreeRootDaoImpl struct {
}

func (inst* p3e97fbda8f_itreeroots_TreeRootDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3e97fbda8f36e61e-itreeroots-TreeRootDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-69e5604197df6fca537ce9f81594a863-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3e97fbda8f_itreeroots_TreeRootDaoImpl) new() any {
    return &p3e97fbda8.TreeRootDaoImpl{}
}

func (inst* p3e97fbda8f_itreeroots_TreeRootDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3e97fbda8.TreeRootDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p3e97fbda8f_itreeroots_TreeRootDaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*p3e97fbda8f_itreeroots_TreeRootDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p3e97fbda8.TreeRootServiceImpl in package:github.com/bitwormhole/wpm/server/implements/itreeroots
//
// id:com-3e97fbda8f36e61e-itreeroots-TreeRootServiceImpl
// class:
// alias:alias-69e5604197df6fca537ce9f81594a863-Service
// scope:singleton
//
type p3e97fbda8f_itreeroots_TreeRootServiceImpl struct {
}

func (inst* p3e97fbda8f_itreeroots_TreeRootServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3e97fbda8f36e61e-itreeroots-TreeRootServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-69e5604197df6fca537ce9f81594a863-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3e97fbda8f_itreeroots_TreeRootServiceImpl) new() any {
    return &p3e97fbda8.TreeRootServiceImpl{}
}

func (inst* p3e97fbda8f_itreeroots_TreeRootServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3e97fbda8.TreeRootServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.FS = inst.getFS(ie)
    com.StateLoader = inst.getStateLoader(ie)


    return nil
}


func (inst*p3e97fbda8f_itreeroots_TreeRootServiceImpl) getDao(ie application.InjectionExt)p69e560419.DAO{
    return ie.GetComponent("#alias-69e5604197df6fca537ce9f81594a863-DAO").(p69e560419.DAO)
}


func (inst*p3e97fbda8f_itreeroots_TreeRootServiceImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


func (inst*p3e97fbda8f_itreeroots_TreeRootServiceImpl) getStateLoader(ie application.InjectionExt)p69e560419.StateLoader{
    return ie.GetComponent("#alias-69e5604197df6fca537ce9f81594a863-StateLoader").(p69e560419.StateLoader)
}



// type p3e97fbda8.TreeRootStateLoaderImpl in package:github.com/bitwormhole/wpm/server/implements/itreeroots
//
// id:com-3e97fbda8f36e61e-itreeroots-TreeRootStateLoaderImpl
// class:
// alias:alias-69e5604197df6fca537ce9f81594a863-StateLoader
// scope:singleton
//
type p3e97fbda8f_itreeroots_TreeRootStateLoaderImpl struct {
}

func (inst* p3e97fbda8f_itreeroots_TreeRootStateLoaderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3e97fbda8f36e61e-itreeroots-TreeRootStateLoaderImpl"
	r.Classes = ""
	r.Aliases = "alias-69e5604197df6fca537ce9f81594a863-StateLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3e97fbda8f_itreeroots_TreeRootStateLoaderImpl) new() any {
    return &p3e97fbda8.TreeRootStateLoaderImpl{}
}

func (inst* p3e97fbda8f_itreeroots_TreeRootStateLoaderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3e97fbda8.TreeRootStateLoaderImpl)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)


    return nil
}


func (inst*p3e97fbda8f_itreeroots_TreeRootStateLoaderImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}



// type p72b9b4e1d.HTTP404Controller in package:github.com/bitwormhole/wpm/server/web/controllers
//
// id:com-72b9b4e1d80225c1-controllers-HTTP404Controller
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p72b9b4e1d8_controllers_HTTP404Controller struct {
}

func (inst* p72b9b4e1d8_controllers_HTTP404Controller) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-72b9b4e1d80225c1-controllers-HTTP404Controller"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p72b9b4e1d8_controllers_HTTP404Controller) new() any {
    return &p72b9b4e1d.HTTP404Controller{}
}

func (inst* p72b9b4e1d8_controllers_HTTP404Controller) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p72b9b4e1d.HTTP404Controller)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)


    return nil
}


func (inst*p72b9b4e1d8_controllers_HTTP404Controller) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}



// type p415d59ee4.AboutController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-AboutController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_AboutController struct {
}

func (inst* p415d59ee4c_admin_AboutController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-AboutController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_AboutController) new() any {
    return &p415d59ee4.AboutController{}
}

func (inst* p415d59ee4c_admin_AboutController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.AboutController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_AboutController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_AboutController) getService(ie application.InjectionExt)p663158af9.Service{
    return ie.GetComponent("#alias-663158af9b1a72a79c266e906db07157-Service").(p663158af9.Service)
}



// type p415d59ee4.AuthsController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-AuthsController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_AuthsController struct {
}

func (inst* p415d59ee4c_admin_AuthsController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-AuthsController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_AuthsController) new() any {
    return &p415d59ee4.AuthsController{}
}

func (inst* p415d59ee4c_admin_AuthsController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.AuthsController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_AuthsController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_AuthsController) getService(ie application.InjectionExt)pce44af903.Service{
    return ie.GetComponent("#alias-ce44af9033c872f977eba59d8437c564-Service").(pce44af903.Service)
}



// type p415d59ee4.BackupsController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-BackupsController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_BackupsController struct {
}

func (inst* p415d59ee4c_admin_BackupsController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-BackupsController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_BackupsController) new() any {
    return &p415d59ee4.BackupsController{}
}

func (inst* p415d59ee4c_admin_BackupsController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.BackupsController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_BackupsController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_BackupsController) getService(ie application.InjectionExt)p7b01405af.Service{
    return ie.GetComponent("#alias-7b01405afb0538c075c5775b29d3562e-Service").(p7b01405af.Service)
}



// type p415d59ee4.ContentTypesController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-ContentTypesController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_ContentTypesController struct {
}

func (inst* p415d59ee4c_admin_ContentTypesController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-ContentTypesController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_ContentTypesController) new() any {
    return &p415d59ee4.ContentTypesController{}
}

func (inst* p415d59ee4c_admin_ContentTypesController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.ContentTypesController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_ContentTypesController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_ContentTypesController) getService(ie application.InjectionExt)p6021e9d7f.Service{
    return ie.GetComponent("#alias-6021e9d7f5afc3355549f2f1fec8b3e7-Service").(p6021e9d7f.Service)
}



// type p415d59ee4.ExampleController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_ExampleController struct {
}

func (inst* p415d59ee4c_admin_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_ExampleController) new() any {
    return &p415d59ee4.ExampleController{}
}

func (inst* p415d59ee4c_admin_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p415d59ee4c_admin_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_ExampleController) getDao(ie application.InjectionExt)pcc7a88d45.DAO{
    return ie.GetComponent("#alias-cc7a88d4518f4d0c9704940596344e7e-DAO").(pcc7a88d45.DAO)
}



// type p415d59ee4.ExecutablesController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-ExecutablesController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_ExecutablesController struct {
}

func (inst* p415d59ee4c_admin_ExecutablesController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-ExecutablesController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_ExecutablesController) new() any {
    return &p415d59ee4.ExecutablesController{}
}

func (inst* p415d59ee4c_admin_ExecutablesController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.ExecutablesController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_ExecutablesController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_ExecutablesController) getService(ie application.InjectionExt)p97b2b30ad.Service{
    return ie.GetComponent("#alias-97b2b30ad7df904c32bf0f040e5527d8-Service").(p97b2b30ad.Service)
}



// type p415d59ee4.FileQueryController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-FileQueryController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_FileQueryController struct {
}

func (inst* p415d59ee4c_admin_FileQueryController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-FileQueryController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_FileQueryController) new() any {
    return &p415d59ee4.FileQueryController{}
}

func (inst* p415d59ee4c_admin_FileQueryController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.FileQueryController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.FileQueryService = inst.getFileQueryService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_FileQueryController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_FileQueryController) getFileQueryService(ie application.InjectionExt)peafe07768.Service{
    return ie.GetComponent("#alias-eafe077682a78bde34843d43b78b6a30-Service").(peafe07768.Service)
}



// type p415d59ee4.ImportRepositoryController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-ImportRepositoryController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_ImportRepositoryController struct {
}

func (inst* p415d59ee4c_admin_ImportRepositoryController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-ImportRepositoryController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_ImportRepositoryController) new() any {
    return &p415d59ee4.ImportRepositoryController{}
}

func (inst* p415d59ee4c_admin_ImportRepositoryController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.ImportRepositoryController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_ImportRepositoryController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_ImportRepositoryController) getService(ie application.InjectionExt)p4a2c02e71.ImportRepositoryService{
    return ie.GetComponent("#alias-4a2c02e71bff15e74f38c4d3caa746b4-ImportRepositoryService").(p4a2c02e71.ImportRepositoryService)
}



// type p415d59ee4.InitController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-InitController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_InitController struct {
}

func (inst* p415d59ee4c_admin_InitController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-InitController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_InitController) new() any {
    return &p415d59ee4.InitController{}
}

func (inst* p415d59ee4c_admin_InitController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.InitController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p415d59ee4c_admin_InitController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_InitController) getDao(ie application.InjectionExt)pcc7a88d45.DAO{
    return ie.GetComponent("#alias-cc7a88d4518f4d0c9704940596344e7e-DAO").(pcc7a88d45.DAO)
}



// type p415d59ee4.IntentQueueController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-IntentQueueController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_IntentQueueController struct {
}

func (inst* p415d59ee4c_admin_IntentQueueController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-IntentQueueController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_IntentQueueController) new() any {
    return &p415d59ee4.IntentQueueController{}
}

func (inst* p415d59ee4c_admin_IntentQueueController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.IntentQueueController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Queues = inst.getQueues(ie)


    return nil
}


func (inst*p415d59ee4c_admin_IntentQueueController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_IntentQueueController) getQueues(ie application.InjectionExt)p709b0834a.Queues{
    return ie.GetComponent("#alias-709b0834aef1394ca72aa585afa48882-Queues").(p709b0834a.Queues)
}



// type p415d59ee4.IntentTemplateController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-IntentTemplateController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_IntentTemplateController struct {
}

func (inst* p415d59ee4c_admin_IntentTemplateController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-IntentTemplateController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_IntentTemplateController) new() any {
    return &p415d59ee4.IntentTemplateController{}
}

func (inst* p415d59ee4c_admin_IntentTemplateController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.IntentTemplateController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_IntentTemplateController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_IntentTemplateController) getService(ie application.InjectionExt)p6db2388a6.Service{
    return ie.GetComponent("#alias-6db2388a63f83ff930eeb1226ef48eb6-Service").(p6db2388a6.Service)
}



// type p415d59ee4.LocalRepositoryController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-LocalRepositoryController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_LocalRepositoryController struct {
}

func (inst* p415d59ee4c_admin_LocalRepositoryController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-LocalRepositoryController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_LocalRepositoryController) new() any {
    return &p415d59ee4.LocalRepositoryController{}
}

func (inst* p415d59ee4c_admin_LocalRepositoryController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.LocalRepositoryController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_LocalRepositoryController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_LocalRepositoryController) getService(ie application.InjectionExt)p4a2c02e71.LocalRepositoryService{
    return ie.GetComponent("#alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryService").(p4a2c02e71.LocalRepositoryService)
}



// type p415d59ee4.MediaController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-MediaController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_MediaController struct {
}

func (inst* p415d59ee4c_admin_MediaController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-MediaController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_MediaController) new() any {
    return &p415d59ee4.MediaController{}
}

func (inst* p415d59ee4c_admin_MediaController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.MediaController)
	nop(ie, com)

	
    com.UploadSizeMax = inst.getUploadSizeMax(ie)
    com.DownloadSizeMax = inst.getDownloadSizeMax(ie)
    com.Sender = inst.getSender(ie)
    com.MediaService = inst.getMediaService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_MediaController) getUploadSizeMax(ie application.InjectionExt)int{
    return ie.GetInt("${http.upload.content-length.max}")
}


func (inst*p415d59ee4c_admin_MediaController) getDownloadSizeMax(ie application.InjectionExt)int{
    return ie.GetInt("${http.download.content-length.max}")
}


func (inst*p415d59ee4c_admin_MediaController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_MediaController) getMediaService(ie application.InjectionExt)p677240472.Service{
    return ie.GetComponent("#alias-67724047202291d9335f729c0f271c46-Service").(p677240472.Service)
}



// type p415d59ee4.ProjectController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-ProjectController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_ProjectController struct {
}

func (inst* p415d59ee4c_admin_ProjectController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-ProjectController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_ProjectController) new() any {
    return &p415d59ee4.ProjectController{}
}

func (inst* p415d59ee4c_admin_ProjectController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.ProjectController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_ProjectController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_ProjectController) getService(ie application.InjectionExt)paa275b61f.Service{
    return ie.GetComponent("#alias-aa275b61f0bcca40e77fdda2bbfc393c-Service").(paa275b61f.Service)
}



// type p415d59ee4.RemoteRepositoryController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-RemoteRepositoryController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_RemoteRepositoryController struct {
}

func (inst* p415d59ee4c_admin_RemoteRepositoryController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-RemoteRepositoryController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_RemoteRepositoryController) new() any {
    return &p415d59ee4.RemoteRepositoryController{}
}

func (inst* p415d59ee4c_admin_RemoteRepositoryController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.RemoteRepositoryController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_RemoteRepositoryController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_RemoteRepositoryController) getService(ie application.InjectionExt)p4a2c02e71.RemoteRepositoryService{
    return ie.GetComponent("#alias-4a2c02e71bff15e74f38c4d3caa746b4-RemoteRepositoryService").(p4a2c02e71.RemoteRepositoryService)
}



// type p415d59ee4.RepositoryWorktreeProjectController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-RepositoryWorktreeProjectController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_RepositoryWorktreeProjectController struct {
}

func (inst* p415d59ee4c_admin_RepositoryWorktreeProjectController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-RepositoryWorktreeProjectController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_RepositoryWorktreeProjectController) new() any {
    return &p415d59ee4.RepositoryWorktreeProjectController{}
}

func (inst* p415d59ee4c_admin_RepositoryWorktreeProjectController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.RepositoryWorktreeProjectController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_RepositoryWorktreeProjectController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_RepositoryWorktreeProjectController) getService(ie application.InjectionExt)p53a0e6a53.Service{
    return ie.GetComponent("#alias-53a0e6a5358524b4150f206572d109d1-Service").(p53a0e6a53.Service)
}



// type p415d59ee4.SoftwarePackageController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-SoftwarePackageController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_SoftwarePackageController struct {
}

func (inst* p415d59ee4c_admin_SoftwarePackageController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-SoftwarePackageController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_SoftwarePackageController) new() any {
    return &p415d59ee4.SoftwarePackageController{}
}

func (inst* p415d59ee4c_admin_SoftwarePackageController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.SoftwarePackageController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_SoftwarePackageController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_SoftwarePackageController) getService(ie application.InjectionExt)pdb29a3bec.Service{
    return ie.GetComponent("#alias-db29a3bec1e3d63283b5b71ee9ef4989-Service").(pdb29a3bec.Service)
}



// type p415d59ee4.SoftwareSetController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-SoftwareSetController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_SoftwareSetController struct {
}

func (inst* p415d59ee4c_admin_SoftwareSetController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-SoftwareSetController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_SoftwareSetController) new() any {
    return &p415d59ee4.SoftwareSetController{}
}

func (inst* p415d59ee4c_admin_SoftwareSetController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.SoftwareSetController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_SoftwareSetController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_SoftwareSetController) getService(ie application.InjectionExt)p353d73ec3.Service{
    return ie.GetComponent("#alias-353d73ec3edc2e74f8181282d9e87a31-Service").(p353d73ec3.Service)
}



// type p415d59ee4.StatisticsController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-StatisticsController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_StatisticsController struct {
}

func (inst* p415d59ee4c_admin_StatisticsController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-StatisticsController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_StatisticsController) new() any {
    return &p415d59ee4.StatisticsController{}
}

func (inst* p415d59ee4c_admin_StatisticsController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.StatisticsController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_StatisticsController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_StatisticsController) getService(ie application.InjectionExt)p3320c87fb.Service{
    return ie.GetComponent("#alias-3320c87fb39933a17055dcfd34498251-Service").(p3320c87fb.Service)
}



// type p415d59ee4.UploadController in package:github.com/bitwormhole/wpm/server/web/controllers/admin
//
// id:com-415d59ee4c0dd6da-admin-UploadController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p415d59ee4c_admin_UploadController struct {
}

func (inst* p415d59ee4c_admin_UploadController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-415d59ee4c0dd6da-admin-UploadController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p415d59ee4c_admin_UploadController) new() any {
    return &p415d59ee4.UploadController{}
}

func (inst* p415d59ee4c_admin_UploadController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p415d59ee4.UploadController)
	nop(ie, com)

	
    com.UploadSizeMax = inst.getUploadSizeMax(ie)
    com.DownloadSizeMax = inst.getDownloadSizeMax(ie)
    com.Sender = inst.getSender(ie)
    com.MediaService = inst.getMediaService(ie)


    return nil
}


func (inst*p415d59ee4c_admin_UploadController) getUploadSizeMax(ie application.InjectionExt)int{
    return ie.GetInt("${http.upload.content-length.max}")
}


func (inst*p415d59ee4c_admin_UploadController) getDownloadSizeMax(ie application.InjectionExt)int{
    return ie.GetInt("${http.download.content-length.max}")
}


func (inst*p415d59ee4c_admin_UploadController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p415d59ee4c_admin_UploadController) getMediaService(ie application.InjectionExt)p677240472.Service{
    return ie.GetComponent("#alias-67724047202291d9335f729c0f271c46-Service").(p677240472.Service)
}


