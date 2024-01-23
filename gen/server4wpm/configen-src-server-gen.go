package server4wpm
import (
    paeb460c7d "github.com/bitwormhole/gitlib"
    p663158af9 "github.com/bitwormhole/wpm/common/classes/about"
    p9ae32fb86 "github.com/bitwormhole/wpm/common/classes/buckets"
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
    p353d73ec3 "github.com/bitwormhole/wpm/server/classes/softwaresets"
    p3ededc14a "github.com/bitwormhole/wpm/server/data/database"
    pfba77a8fb "github.com/bitwormhole/wpm/server/implements/example"
    p9b527161a "github.com/bitwormhole/wpm/server/implements/ifilequery"
    pd11deb9a0 "github.com/bitwormhole/wpm/server/implements/impauths"
    pc347f939e "github.com/bitwormhole/wpm/server/implements/impbackups"
    pd5537a67b "github.com/bitwormhole/wpm/server/implements/impcontenttypes"
    p7fca5d6e6 "github.com/bitwormhole/wpm/server/implements/impexecutables"
    p31f30d9ac "github.com/bitwormhole/wpm/server/implements/impintentqueues"
    p1eb0a0660 "github.com/bitwormhole/wpm/server/implements/impintenttemplates"
    pb5fa9685f "github.com/bitwormhole/wpm/server/implements/implocations"
    p8cd642664 "github.com/bitwormhole/wpm/server/implements/impmedia"
    pf441dbb74 "github.com/bitwormhole/wpm/server/implements/impprojects"
    p233917231 "github.com/bitwormhole/wpm/server/implements/imprepositories"
    p337e784ad "github.com/bitwormhole/wpm/server/implements/impsettings"
    p8f4406e5f "github.com/bitwormhole/wpm/server/implements/impsoftware"
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


    return nil
}


func (inst*p9b527161a5_ifilequery_ServiceImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}



// type pd11deb9a0.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/impauths
//
// id:com-d11deb9a07724f35-impauths-DaoImpl
// class:
// alias:alias-ce44af9033c872f977eba59d8437c564-DAO
// scope:singleton
//
type pd11deb9a07_impauths_DaoImpl struct {
}

func (inst* pd11deb9a07_impauths_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d11deb9a07724f35-impauths-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-ce44af9033c872f977eba59d8437c564-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd11deb9a07_impauths_DaoImpl) new() any {
    return &pd11deb9a0.DaoImpl{}
}

func (inst* pd11deb9a07_impauths_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd11deb9a0.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pd11deb9a07_impauths_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*pd11deb9a07_impauths_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pd11deb9a0.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/impauths
//
// id:com-d11deb9a07724f35-impauths-ServiceImpl
// class:
// alias:alias-ce44af9033c872f977eba59d8437c564-Service
// scope:singleton
//
type pd11deb9a07_impauths_ServiceImpl struct {
}

func (inst* pd11deb9a07_impauths_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d11deb9a07724f35-impauths-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-ce44af9033c872f977eba59d8437c564-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd11deb9a07_impauths_ServiceImpl) new() any {
    return &pd11deb9a0.ServiceImpl{}
}

func (inst* pd11deb9a07_impauths_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd11deb9a0.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pd11deb9a07_impauths_ServiceImpl) getDao(ie application.InjectionExt)pce44af903.DAO{
    return ie.GetComponent("#alias-ce44af9033c872f977eba59d8437c564-DAO").(pce44af903.DAO)
}



// type pc347f939e.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/impbackups
//
// id:com-c347f939edc8147e-impbackups-DaoImpl
// class:
// alias:alias-7b01405afb0538c075c5775b29d3562e-DAO
// scope:singleton
//
type pc347f939ed_impbackups_DaoImpl struct {
}

func (inst* pc347f939ed_impbackups_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c347f939edc8147e-impbackups-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-7b01405afb0538c075c5775b29d3562e-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc347f939ed_impbackups_DaoImpl) new() any {
    return &pc347f939e.DaoImpl{}
}

func (inst* pc347f939ed_impbackups_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc347f939e.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pc347f939ed_impbackups_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*pc347f939ed_impbackups_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pc347f939e.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/impbackups
//
// id:com-c347f939edc8147e-impbackups-ServiceImpl
// class:
// alias:alias-7b01405afb0538c075c5775b29d3562e-Service
// scope:singleton
//
type pc347f939ed_impbackups_ServiceImpl struct {
}

func (inst* pc347f939ed_impbackups_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c347f939edc8147e-impbackups-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-7b01405afb0538c075c5775b29d3562e-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc347f939ed_impbackups_ServiceImpl) new() any {
    return &pc347f939e.ServiceImpl{}
}

func (inst* pc347f939ed_impbackups_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc347f939e.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pc347f939ed_impbackups_ServiceImpl) getDao(ie application.InjectionExt)p7b01405af.DAO{
    return ie.GetComponent("#alias-7b01405afb0538c075c5775b29d3562e-DAO").(p7b01405af.DAO)
}



// type pd5537a67b.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/impcontenttypes
//
// id:com-d5537a67b8cf1421-impcontenttypes-DaoImpl
// class:
// alias:alias-6021e9d7f5afc3355549f2f1fec8b3e7-DAO
// scope:singleton
//
type pd5537a67b8_impcontenttypes_DaoImpl struct {
}

func (inst* pd5537a67b8_impcontenttypes_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d5537a67b8cf1421-impcontenttypes-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-6021e9d7f5afc3355549f2f1fec8b3e7-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd5537a67b8_impcontenttypes_DaoImpl) new() any {
    return &pd5537a67b.DaoImpl{}
}

func (inst* pd5537a67b8_impcontenttypes_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd5537a67b.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pd5537a67b8_impcontenttypes_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*pd5537a67b8_impcontenttypes_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pd5537a67b.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/impcontenttypes
//
// id:com-d5537a67b8cf1421-impcontenttypes-ServiceImpl
// class:
// alias:alias-6021e9d7f5afc3355549f2f1fec8b3e7-Service
// scope:singleton
//
type pd5537a67b8_impcontenttypes_ServiceImpl struct {
}

func (inst* pd5537a67b8_impcontenttypes_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d5537a67b8cf1421-impcontenttypes-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-6021e9d7f5afc3355549f2f1fec8b3e7-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd5537a67b8_impcontenttypes_ServiceImpl) new() any {
    return &pd5537a67b.ServiceImpl{}
}

func (inst* pd5537a67b8_impcontenttypes_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd5537a67b.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pd5537a67b8_impcontenttypes_ServiceImpl) getDao(ie application.InjectionExt)p6021e9d7f.DAO{
    return ie.GetComponent("#alias-6021e9d7f5afc3355549f2f1fec8b3e7-DAO").(p6021e9d7f.DAO)
}



// type p7fca5d6e6.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/impexecutables
//
// id:com-7fca5d6e6670ce3b-impexecutables-DaoImpl
// class:
// alias:alias-97b2b30ad7df904c32bf0f040e5527d8-DAO
// scope:singleton
//
type p7fca5d6e66_impexecutables_DaoImpl struct {
}

func (inst* p7fca5d6e66_impexecutables_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7fca5d6e6670ce3b-impexecutables-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-97b2b30ad7df904c32bf0f040e5527d8-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7fca5d6e66_impexecutables_DaoImpl) new() any {
    return &p7fca5d6e6.DaoImpl{}
}

func (inst* p7fca5d6e66_impexecutables_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7fca5d6e6.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p7fca5d6e66_impexecutables_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*p7fca5d6e66_impexecutables_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p7fca5d6e6.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/impexecutables
//
// id:com-7fca5d6e6670ce3b-impexecutables-ServiceImpl
// class:
// alias:alias-97b2b30ad7df904c32bf0f040e5527d8-Service
// scope:singleton
//
type p7fca5d6e66_impexecutables_ServiceImpl struct {
}

func (inst* p7fca5d6e66_impexecutables_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7fca5d6e6670ce3b-impexecutables-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-97b2b30ad7df904c32bf0f040e5527d8-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7fca5d6e66_impexecutables_ServiceImpl) new() any {
    return &p7fca5d6e6.ServiceImpl{}
}

func (inst* p7fca5d6e66_impexecutables_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7fca5d6e6.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p7fca5d6e66_impexecutables_ServiceImpl) getDao(ie application.InjectionExt)p97b2b30ad.DAO{
    return ie.GetComponent("#alias-97b2b30ad7df904c32bf0f040e5527d8-DAO").(p97b2b30ad.DAO)
}



// type p31f30d9ac.IntentQueueService in package:github.com/bitwormhole/wpm/server/implements/impintentqueues
//
// id:com-31f30d9aca437348-impintentqueues-IntentQueueService
// class:
// alias:alias-709b0834aef1394ca72aa585afa48882-Queues
// scope:singleton
//
type p31f30d9aca_impintentqueues_IntentQueueService struct {
}

func (inst* p31f30d9aca_impintentqueues_IntentQueueService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-31f30d9aca437348-impintentqueues-IntentQueueService"
	r.Classes = ""
	r.Aliases = "alias-709b0834aef1394ca72aa585afa48882-Queues"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p31f30d9aca_impintentqueues_IntentQueueService) new() any {
    return &p31f30d9ac.IntentQueueService{}
}

func (inst* p31f30d9aca_impintentqueues_IntentQueueService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p31f30d9ac.IntentQueueService)
	nop(ie, com)

	


    return nil
}



// type p1eb0a0660.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/impintenttemplates
//
// id:com-1eb0a0660bb35862-impintenttemplates-DaoImpl
// class:
// alias:alias-6db2388a63f83ff930eeb1226ef48eb6-DAO
// scope:singleton
//
type p1eb0a0660b_impintenttemplates_DaoImpl struct {
}

func (inst* p1eb0a0660b_impintenttemplates_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1eb0a0660bb35862-impintenttemplates-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-6db2388a63f83ff930eeb1226ef48eb6-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1eb0a0660b_impintenttemplates_DaoImpl) new() any {
    return &p1eb0a0660.DaoImpl{}
}

func (inst* p1eb0a0660b_impintenttemplates_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1eb0a0660.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p1eb0a0660b_impintenttemplates_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*p1eb0a0660b_impintenttemplates_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p1eb0a0660.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/impintenttemplates
//
// id:com-1eb0a0660bb35862-impintenttemplates-ServiceImpl
// class:
// alias:alias-6db2388a63f83ff930eeb1226ef48eb6-Service
// scope:singleton
//
type p1eb0a0660b_impintenttemplates_ServiceImpl struct {
}

func (inst* p1eb0a0660b_impintenttemplates_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1eb0a0660bb35862-impintenttemplates-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-6db2388a63f83ff930eeb1226ef48eb6-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1eb0a0660b_impintenttemplates_ServiceImpl) new() any {
    return &p1eb0a0660.ServiceImpl{}
}

func (inst* p1eb0a0660b_impintenttemplates_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1eb0a0660.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p1eb0a0660b_impintenttemplates_ServiceImpl) getDao(ie application.InjectionExt)p6db2388a6.DAO{
    return ie.GetComponent("#alias-6db2388a63f83ff930eeb1226ef48eb6-DAO").(p6db2388a6.DAO)
}



// type pb5fa9685f.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/implocations
//
// id:com-b5fa9685f808298c-implocations-DaoImpl
// class:
// alias:alias-3c68bd3f6d1fbc95a5a70cc9cbd6d4ea-DAO
// scope:singleton
//
type pb5fa9685f8_implocations_DaoImpl struct {
}

func (inst* pb5fa9685f8_implocations_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b5fa9685f808298c-implocations-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-3c68bd3f6d1fbc95a5a70cc9cbd6d4ea-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb5fa9685f8_implocations_DaoImpl) new() any {
    return &pb5fa9685f.DaoImpl{}
}

func (inst* pb5fa9685f8_implocations_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb5fa9685f.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pb5fa9685f8_implocations_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*pb5fa9685f8_implocations_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pb5fa9685f.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/implocations
//
// id:com-b5fa9685f808298c-implocations-ServiceImpl
// class:
// alias:alias-3c68bd3f6d1fbc95a5a70cc9cbd6d4ea-Service
// scope:singleton
//
type pb5fa9685f8_implocations_ServiceImpl struct {
}

func (inst* pb5fa9685f8_implocations_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b5fa9685f808298c-implocations-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-3c68bd3f6d1fbc95a5a70cc9cbd6d4ea-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb5fa9685f8_implocations_ServiceImpl) new() any {
    return &pb5fa9685f.ServiceImpl{}
}

func (inst* pb5fa9685f8_implocations_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb5fa9685f.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pb5fa9685f8_implocations_ServiceImpl) getDao(ie application.InjectionExt)p3c68bd3f6.DAO{
    return ie.GetComponent("#alias-3c68bd3f6d1fbc95a5a70cc9cbd6d4ea-DAO").(p3c68bd3f6.DAO)
}



// type p8cd642664.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/impmedia
//
// id:com-8cd64266481c28e5-impmedia-DaoImpl
// class:
// alias:alias-67724047202291d9335f729c0f271c46-DAO
// scope:singleton
//
type p8cd6426648_impmedia_DaoImpl struct {
}

func (inst* p8cd6426648_impmedia_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8cd64266481c28e5-impmedia-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-67724047202291d9335f729c0f271c46-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8cd6426648_impmedia_DaoImpl) new() any {
    return &p8cd642664.DaoImpl{}
}

func (inst* p8cd6426648_impmedia_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8cd642664.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p8cd6426648_impmedia_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*p8cd6426648_impmedia_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p8cd642664.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/impmedia
//
// id:com-8cd64266481c28e5-impmedia-ServiceImpl
// class:
// alias:alias-67724047202291d9335f729c0f271c46-Service
// scope:singleton
//
type p8cd6426648_impmedia_ServiceImpl struct {
}

func (inst* p8cd6426648_impmedia_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8cd64266481c28e5-impmedia-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-67724047202291d9335f729c0f271c46-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8cd6426648_impmedia_ServiceImpl) new() any {
    return &p8cd642664.ServiceImpl{}
}

func (inst* p8cd6426648_impmedia_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8cd642664.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.Buckets = inst.getBuckets(ie)


    return nil
}


func (inst*p8cd6426648_impmedia_ServiceImpl) getDao(ie application.InjectionExt)p677240472.DAO{
    return ie.GetComponent("#alias-67724047202291d9335f729c0f271c46-DAO").(p677240472.DAO)
}


func (inst*p8cd6426648_impmedia_ServiceImpl) getBuckets(ie application.InjectionExt)p9ae32fb86.BucketPool{
    return ie.GetComponent("#alias-9ae32fb866160a2b2e9745348187d238-BucketPool").(p9ae32fb86.BucketPool)
}



// type pf441dbb74.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/impprojects
//
// id:com-f441dbb74d09b5d7-impprojects-DaoImpl
// class:
// alias:alias-aa275b61f0bcca40e77fdda2bbfc393c-DAO
// scope:singleton
//
type pf441dbb74d_impprojects_DaoImpl struct {
}

func (inst* pf441dbb74d_impprojects_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f441dbb74d09b5d7-impprojects-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-aa275b61f0bcca40e77fdda2bbfc393c-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf441dbb74d_impprojects_DaoImpl) new() any {
    return &pf441dbb74.DaoImpl{}
}

func (inst* pf441dbb74d_impprojects_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf441dbb74.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pf441dbb74d_impprojects_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*pf441dbb74d_impprojects_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pf441dbb74.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/impprojects
//
// id:com-f441dbb74d09b5d7-impprojects-ServiceImpl
// class:
// alias:alias-aa275b61f0bcca40e77fdda2bbfc393c-Service
// scope:singleton
//
type pf441dbb74d_impprojects_ServiceImpl struct {
}

func (inst* pf441dbb74d_impprojects_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f441dbb74d09b5d7-impprojects-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-aa275b61f0bcca40e77fdda2bbfc393c-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf441dbb74d_impprojects_ServiceImpl) new() any {
    return &pf441dbb74.ServiceImpl{}
}

func (inst* pf441dbb74d_impprojects_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf441dbb74.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pf441dbb74d_impprojects_ServiceImpl) getDao(ie application.InjectionExt)paa275b61f.DAO{
    return ie.GetComponent("#alias-aa275b61f0bcca40e77fdda2bbfc393c-DAO").(paa275b61f.DAO)
}



// type p233917231.ImportRepositoryServiceImpl in package:github.com/bitwormhole/wpm/server/implements/imprepositories
//
// id:com-233917231e1bcc32-imprepositories-ImportRepositoryServiceImpl
// class:
// alias:alias-4a2c02e71bff15e74f38c4d3caa746b4-ImportRepositoryService
// scope:singleton
//
type p233917231e_imprepositories_ImportRepositoryServiceImpl struct {
}

func (inst* p233917231e_imprepositories_ImportRepositoryServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-233917231e1bcc32-imprepositories-ImportRepositoryServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-4a2c02e71bff15e74f38c4d3caa746b4-ImportRepositoryService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p233917231e_imprepositories_ImportRepositoryServiceImpl) new() any {
    return &p233917231.ImportRepositoryServiceImpl{}
}

func (inst* p233917231e_imprepositories_ImportRepositoryServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p233917231.ImportRepositoryServiceImpl)
	nop(ie, com)

	
    com.LibAgent = inst.getLibAgent(ie)
    com.RepoService = inst.getRepoService(ie)


    return nil
}


func (inst*p233917231e_imprepositories_ImportRepositoryServiceImpl) getLibAgent(ie application.InjectionExt)paeb460c7d.Agent{
    return ie.GetComponent("#alias-aeb460c7d339df24b0b38a0d65e30102-Agent").(paeb460c7d.Agent)
}


func (inst*p233917231e_imprepositories_ImportRepositoryServiceImpl) getRepoService(ie application.InjectionExt)p4a2c02e71.LocalRepositoryService{
    return ie.GetComponent("#alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryService").(p4a2c02e71.LocalRepositoryService)
}



// type p233917231.LocalRepositoryDao in package:github.com/bitwormhole/wpm/server/implements/imprepositories
//
// id:com-233917231e1bcc32-imprepositories-LocalRepositoryDao
// class:
// alias:alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryDAO
// scope:singleton
//
type p233917231e_imprepositories_LocalRepositoryDao struct {
}

func (inst* p233917231e_imprepositories_LocalRepositoryDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-233917231e1bcc32-imprepositories-LocalRepositoryDao"
	r.Classes = ""
	r.Aliases = "alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p233917231e_imprepositories_LocalRepositoryDao) new() any {
    return &p233917231.LocalRepositoryDao{}
}

func (inst* p233917231e_imprepositories_LocalRepositoryDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p233917231.LocalRepositoryDao)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p233917231e_imprepositories_LocalRepositoryDao) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*p233917231e_imprepositories_LocalRepositoryDao) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p233917231.LocalRepositoryService in package:github.com/bitwormhole/wpm/server/implements/imprepositories
//
// id:com-233917231e1bcc32-imprepositories-LocalRepositoryService
// class:
// alias:alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryService
// scope:singleton
//
type p233917231e_imprepositories_LocalRepositoryService struct {
}

func (inst* p233917231e_imprepositories_LocalRepositoryService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-233917231e1bcc32-imprepositories-LocalRepositoryService"
	r.Classes = ""
	r.Aliases = "alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p233917231e_imprepositories_LocalRepositoryService) new() any {
    return &p233917231.LocalRepositoryService{}
}

func (inst* p233917231e_imprepositories_LocalRepositoryService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p233917231.LocalRepositoryService)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.LibAgent = inst.getLibAgent(ie)


    return nil
}


func (inst*p233917231e_imprepositories_LocalRepositoryService) getDao(ie application.InjectionExt)p4a2c02e71.LocalRepositoryDAO{
    return ie.GetComponent("#alias-4a2c02e71bff15e74f38c4d3caa746b4-LocalRepositoryDAO").(p4a2c02e71.LocalRepositoryDAO)
}


func (inst*p233917231e_imprepositories_LocalRepositoryService) getLibAgent(ie application.InjectionExt)paeb460c7d.Agent{
    return ie.GetComponent("#alias-aeb460c7d339df24b0b38a0d65e30102-Agent").(paeb460c7d.Agent)
}



// type p233917231.RemoteRepositoryDao in package:github.com/bitwormhole/wpm/server/implements/imprepositories
//
// id:com-233917231e1bcc32-imprepositories-RemoteRepositoryDao
// class:
// alias:alias-4a2c02e71bff15e74f38c4d3caa746b4-RemoteRepositoryDAO
// scope:singleton
//
type p233917231e_imprepositories_RemoteRepositoryDao struct {
}

func (inst* p233917231e_imprepositories_RemoteRepositoryDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-233917231e1bcc32-imprepositories-RemoteRepositoryDao"
	r.Classes = ""
	r.Aliases = "alias-4a2c02e71bff15e74f38c4d3caa746b4-RemoteRepositoryDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p233917231e_imprepositories_RemoteRepositoryDao) new() any {
    return &p233917231.RemoteRepositoryDao{}
}

func (inst* p233917231e_imprepositories_RemoteRepositoryDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p233917231.RemoteRepositoryDao)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p233917231e_imprepositories_RemoteRepositoryDao) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*p233917231e_imprepositories_RemoteRepositoryDao) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p233917231.RemoteRepositoryService in package:github.com/bitwormhole/wpm/server/implements/imprepositories
//
// id:com-233917231e1bcc32-imprepositories-RemoteRepositoryService
// class:
// alias:alias-4a2c02e71bff15e74f38c4d3caa746b4-RemoteRepositoryService
// scope:singleton
//
type p233917231e_imprepositories_RemoteRepositoryService struct {
}

func (inst* p233917231e_imprepositories_RemoteRepositoryService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-233917231e1bcc32-imprepositories-RemoteRepositoryService"
	r.Classes = ""
	r.Aliases = "alias-4a2c02e71bff15e74f38c4d3caa746b4-RemoteRepositoryService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p233917231e_imprepositories_RemoteRepositoryService) new() any {
    return &p233917231.RemoteRepositoryService{}
}

func (inst* p233917231e_imprepositories_RemoteRepositoryService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p233917231.RemoteRepositoryService)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p233917231e_imprepositories_RemoteRepositoryService) getDao(ie application.InjectionExt)pcc7a88d45.DAO{
    return ie.GetComponent("#alias-cc7a88d4518f4d0c9704940596344e7e-DAO").(pcc7a88d45.DAO)
}



// type p337e784ad.DaoImpl in package:github.com/bitwormhole/wpm/server/implements/impsettings
//
// id:com-337e784ada98004a-impsettings-DaoImpl
// class:
// alias:alias-4ab661a34cf8539ff7a5d51208b7f52a-DAO
// scope:singleton
//
type p337e784ada_impsettings_DaoImpl struct {
}

func (inst* p337e784ada_impsettings_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-337e784ada98004a-impsettings-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-4ab661a34cf8539ff7a5d51208b7f52a-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p337e784ada_impsettings_DaoImpl) new() any {
    return &p337e784ad.DaoImpl{}
}

func (inst* p337e784ada_impsettings_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p337e784ad.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p337e784ada_impsettings_DaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*p337e784ada_impsettings_DaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p337e784ad.ServiceImpl in package:github.com/bitwormhole/wpm/server/implements/impsettings
//
// id:com-337e784ada98004a-impsettings-ServiceImpl
// class:
// alias:alias-cc7a88d4518f4d0c9704940596344e7e-Service
// scope:singleton
//
type p337e784ada_impsettings_ServiceImpl struct {
}

func (inst* p337e784ada_impsettings_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-337e784ada98004a-impsettings-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-cc7a88d4518f4d0c9704940596344e7e-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p337e784ada_impsettings_ServiceImpl) new() any {
    return &p337e784ad.ServiceImpl{}
}

func (inst* p337e784ada_impsettings_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p337e784ad.ServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p337e784ada_impsettings_ServiceImpl) getDao(ie application.InjectionExt)pcc7a88d45.DAO{
    return ie.GetComponent("#alias-cc7a88d4518f4d0c9704940596344e7e-DAO").(pcc7a88d45.DAO)
}



// type p8f4406e5f.SoftwarePackageDaoImpl in package:github.com/bitwormhole/wpm/server/implements/impsoftware
//
// id:com-8f4406e5f78c176d-impsoftware-SoftwarePackageDaoImpl
// class:
// alias:alias-db29a3bec1e3d63283b5b71ee9ef4989-DAO
// scope:singleton
//
type p8f4406e5f7_impsoftware_SoftwarePackageDaoImpl struct {
}

func (inst* p8f4406e5f7_impsoftware_SoftwarePackageDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8f4406e5f78c176d-impsoftware-SoftwarePackageDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-db29a3bec1e3d63283b5b71ee9ef4989-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8f4406e5f7_impsoftware_SoftwarePackageDaoImpl) new() any {
    return &p8f4406e5f.SoftwarePackageDaoImpl{}
}

func (inst* p8f4406e5f7_impsoftware_SoftwarePackageDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8f4406e5f.SoftwarePackageDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*p8f4406e5f7_impsoftware_SoftwarePackageDaoImpl) getAgent(ie application.InjectionExt)p080073581.DatabaseAgent{
    return ie.GetComponent("#alias-08007358192e7a966d957c7686a43d06-DatabaseAgent").(p080073581.DatabaseAgent)
}


func (inst*p8f4406e5f7_impsoftware_SoftwarePackageDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p8f4406e5f.SoftwarePackageService in package:github.com/bitwormhole/wpm/server/implements/impsoftware
//
// id:com-8f4406e5f78c176d-impsoftware-SoftwarePackageService
// class:
// alias:alias-db29a3bec1e3d63283b5b71ee9ef4989-Service
// scope:singleton
//
type p8f4406e5f7_impsoftware_SoftwarePackageService struct {
}

func (inst* p8f4406e5f7_impsoftware_SoftwarePackageService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8f4406e5f78c176d-impsoftware-SoftwarePackageService"
	r.Classes = ""
	r.Aliases = "alias-db29a3bec1e3d63283b5b71ee9ef4989-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8f4406e5f7_impsoftware_SoftwarePackageService) new() any {
    return &p8f4406e5f.SoftwarePackageService{}
}

func (inst* p8f4406e5f7_impsoftware_SoftwarePackageService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8f4406e5f.SoftwarePackageService)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p8f4406e5f7_impsoftware_SoftwarePackageService) getDao(ie application.InjectionExt)pdb29a3bec.DAO{
    return ie.GetComponent("#alias-db29a3bec1e3d63283b5b71ee9ef4989-DAO").(pdb29a3bec.DAO)
}



// type p8f4406e5f.SoftwareSetService in package:github.com/bitwormhole/wpm/server/implements/impsoftware
//
// id:com-8f4406e5f78c176d-impsoftware-SoftwareSetService
// class:
// alias:alias-353d73ec3edc2e74f8181282d9e87a31-Service
// scope:singleton
//
type p8f4406e5f7_impsoftware_SoftwareSetService struct {
}

func (inst* p8f4406e5f7_impsoftware_SoftwareSetService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8f4406e5f78c176d-impsoftware-SoftwareSetService"
	r.Classes = ""
	r.Aliases = "alias-353d73ec3edc2e74f8181282d9e87a31-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8f4406e5f7_impsoftware_SoftwareSetService) new() any {
    return &p8f4406e5f.SoftwareSetService{}
}

func (inst* p8f4406e5f7_impsoftware_SoftwareSetService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8f4406e5f.SoftwareSetService)
	nop(ie, com)

	
    com.PackageService = inst.getPackageService(ie)


    return nil
}


func (inst*p8f4406e5f7_impsoftware_SoftwareSetService) getPackageService(ie application.InjectionExt)pdb29a3bec.Service{
    return ie.GetComponent("#alias-db29a3bec1e3d63283b5b71ee9ef4989-Service").(pdb29a3bec.Service)
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


