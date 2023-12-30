package server4wpm
import (
    p69775efb7 "github.com/bitwormhole/wpm/server/data/dao"
    p3ededc14a "github.com/bitwormhole/wpm/server/data/database"
    p410b3e070 "github.com/bitwormhole/wpm/server/data/dxo"
    pfba77a8fb "github.com/bitwormhole/wpm/server/implements/example"
    p9c7a9f180 "github.com/bitwormhole/wpm/server/web/controllers/apiv1"
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
     "github.com/starter-go/application"
)

// type p3ededc14a.ThisGroup in package:github.com/bitwormhole/wpm/server/data/database
//
// id:com-3ededc14a5c38f08-database-ThisGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:alias-410b3e0705d26ea4e345ca7cbbc8388f-DatabaseAgent
// scope:singleton
//
type p3ededc14a5_database_ThisGroup struct {
}

func (inst* p3ededc14a5_database_ThisGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3ededc14a5c38f08-database-ThisGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = "alias-410b3e0705d26ea4e345ca7cbbc8388f-DatabaseAgent"
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
// alias:alias-69775efb768ffae09347d9efede6552d-ExampleDAO
// scope:singleton
//
type pfba77a8fb5_example_DaoImpl struct {
}

func (inst* pfba77a8fb5_example_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fba77a8fb5da6105-example-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-69775efb768ffae09347d9efede6552d-ExampleDAO"
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


    return nil
}


func (inst*pfba77a8fb5_example_DaoImpl) getAgent(ie application.InjectionExt)p410b3e070.DatabaseAgent{
    return ie.GetComponent("#alias-410b3e0705d26ea4e345ca7cbbc8388f-DatabaseAgent").(p410b3e070.DatabaseAgent)
}



// type p9c7a9f180.ExampleController in package:github.com/bitwormhole/wpm/server/web/controllers/apiv1
//
// id:com-9c7a9f180b2185ff-apiv1-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9c7a9f180b_apiv1_ExampleController struct {
}

func (inst* p9c7a9f180b_apiv1_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9c7a9f180b2185ff-apiv1-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9c7a9f180b_apiv1_ExampleController) new() any {
    return &p9c7a9f180.ExampleController{}
}

func (inst* p9c7a9f180b_apiv1_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9c7a9f180.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p9c7a9f180b_apiv1_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p9c7a9f180b_apiv1_ExampleController) getDao(ie application.InjectionExt)p69775efb7.ExampleDAO{
    return ie.GetComponent("#alias-69775efb768ffae09347d9efede6552d-ExampleDAO").(p69775efb7.ExampleDAO)
}


