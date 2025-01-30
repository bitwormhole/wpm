package main4wpm
import (
    pfdd8fdb8d "github.com/bitwormhole/wpm/app/classes/locations"
    pfc892d171 "github.com/bitwormhole/wpm/app/classes/repositories"
    p02d112be4 "github.com/bitwormhole/wpm/app/data/dao"
    p70207ce47 "github.com/bitwormhole/wpm/app/data/dxo"
    p57f9b3839 "github.com/bitwormhole/wpm/src/main/golang/controllers/apiv1"
    p8aac3fe08 "github.com/bitwormhole/wpm/src/main/golang/implements/database"
    p213a6f747 "github.com/bitwormhole/wpm/src/main/golang/implements/database/idao"
    p636ab1bcf "github.com/bitwormhole/wpm/src/main/golang/implements/example"
    pb25d04dc6 "github.com/bitwormhole/wpm/src/main/golang/implements/ilocations"
    pd7f15ffc3 "github.com/bitwormhole/wpm/src/main/golang/implements/irepositories"
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
    p9621e8b71 "github.com/starter-go/security/random"
     "github.com/starter-go/application"
)

// type p57f9b3839.ExampleController in package:github.com/bitwormhole/wpm/src/main/golang/controllers/apiv1
//
// id:com-57f9b38391fc605a-apiv1-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p57f9b38391_apiv1_ExampleController struct {
}

func (inst* p57f9b38391_apiv1_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-57f9b38391fc605a-apiv1-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p57f9b38391_apiv1_ExampleController) new() any {
    return &p57f9b3839.ExampleController{}
}

func (inst* p57f9b38391_apiv1_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p57f9b3839.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p57f9b38391_apiv1_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p57f9b38391_apiv1_ExampleController) getDao(ie application.InjectionExt)p02d112be4.ExampleDAO{
    return ie.GetComponent("#alias-02d112be4878b8e01234c65940ea773b-ExampleDAO").(p02d112be4.ExampleDAO)
}



// type p57f9b3839.LocalRepositoryController in package:github.com/bitwormhole/wpm/src/main/golang/controllers/apiv1
//
// id:com-57f9b38391fc605a-apiv1-LocalRepositoryController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p57f9b38391_apiv1_LocalRepositoryController struct {
}

func (inst* p57f9b38391_apiv1_LocalRepositoryController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-57f9b38391fc605a-apiv1-LocalRepositoryController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p57f9b38391_apiv1_LocalRepositoryController) new() any {
    return &p57f9b3839.LocalRepositoryController{}
}

func (inst* p57f9b38391_apiv1_LocalRepositoryController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p57f9b3839.LocalRepositoryController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p57f9b38391_apiv1_LocalRepositoryController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p57f9b38391_apiv1_LocalRepositoryController) getService(ie application.InjectionExt)pfc892d171.LocalRepositoryService{
    return ie.GetComponent("#alias-fc892d171029572157d3f320608a6038-LocalRepositoryService").(pfc892d171.LocalRepositoryService)
}



// type p8aac3fe08.ThisGroup in package:github.com/bitwormhole/wpm/src/main/golang/implements/database
//
// id:com-8aac3fe082dc6e6e-database-ThisGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:alias-70207ce47214c3f405cba821c7b07d52-DatabaseAgent
// scope:singleton
//
type p8aac3fe082_database_ThisGroup struct {
}

func (inst* p8aac3fe082_database_ThisGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8aac3fe082dc6e6e-database-ThisGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = "alias-70207ce47214c3f405cba821c7b07d52-DatabaseAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8aac3fe082_database_ThisGroup) new() any {
    return &p8aac3fe08.ThisGroup{}
}

func (inst* p8aac3fe082_database_ThisGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8aac3fe08.ThisGroup)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Alias = inst.getAlias(ie)
    com.URI = inst.getURI(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)
    com.SourceManager = inst.getSourceManager(ie)


    return nil
}


func (inst*p8aac3fe082_database_ThisGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.default.enabled}")
}


func (inst*p8aac3fe082_database_ThisGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.alias}")
}


func (inst*p8aac3fe082_database_ThisGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.uri}")
}


func (inst*p8aac3fe082_database_ThisGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.table-name-prefix}")
}


func (inst*p8aac3fe082_database_ThisGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.datasource}")
}


func (inst*p8aac3fe082_database_ThisGroup) getSourceManager(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}



// type p213a6f747.LocationDaoImpl in package:github.com/bitwormhole/wpm/src/main/golang/implements/database/idao
//
// id:com-213a6f747acb9b55-idao-LocationDaoImpl
// class:
// alias:alias-fdd8fdb8d991bf9dc6045188ca258fc6-LocationDAO
// scope:singleton
//
type p213a6f747a_idao_LocationDaoImpl struct {
}

func (inst* p213a6f747a_idao_LocationDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-213a6f747acb9b55-idao-LocationDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-fdd8fdb8d991bf9dc6045188ca258fc6-LocationDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p213a6f747a_idao_LocationDaoImpl) new() any {
    return &p213a6f747.LocationDaoImpl{}
}

func (inst* p213a6f747a_idao_LocationDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p213a6f747.LocationDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGen = inst.getUUIDGen(ie)


    return nil
}


func (inst*p213a6f747a_idao_LocationDaoImpl) getAgent(ie application.InjectionExt)p70207ce47.DatabaseAgent{
    return ie.GetComponent("#alias-70207ce47214c3f405cba821c7b07d52-DatabaseAgent").(p70207ce47.DatabaseAgent)
}


func (inst*p213a6f747a_idao_LocationDaoImpl) getUUIDGen(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p636ab1bcf.DaoImpl in package:github.com/bitwormhole/wpm/src/main/golang/implements/example
//
// id:com-636ab1bcf16fcb5d-example-DaoImpl
// class:
// alias:alias-02d112be4878b8e01234c65940ea773b-ExampleDAO
// scope:singleton
//
type p636ab1bcf1_example_DaoImpl struct {
}

func (inst* p636ab1bcf1_example_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-636ab1bcf16fcb5d-example-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-02d112be4878b8e01234c65940ea773b-ExampleDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p636ab1bcf1_example_DaoImpl) new() any {
    return &p636ab1bcf.DaoImpl{}
}

func (inst* p636ab1bcf1_example_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p636ab1bcf.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)


    return nil
}


func (inst*p636ab1bcf1_example_DaoImpl) getAgent(ie application.InjectionExt)p70207ce47.DatabaseAgent{
    return ie.GetComponent("#alias-70207ce47214c3f405cba821c7b07d52-DatabaseAgent").(p70207ce47.DatabaseAgent)
}



// type pb25d04dc6.LocationServiceImpl in package:github.com/bitwormhole/wpm/src/main/golang/implements/ilocations
//
// id:com-b25d04dc62eced83-ilocations-LocationServiceImpl
// class:
// alias:alias-fdd8fdb8d991bf9dc6045188ca258fc6-Service
// scope:singleton
//
type pb25d04dc62_ilocations_LocationServiceImpl struct {
}

func (inst* pb25d04dc62_ilocations_LocationServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b25d04dc62eced83-ilocations-LocationServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-fdd8fdb8d991bf9dc6045188ca258fc6-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb25d04dc62_ilocations_LocationServiceImpl) new() any {
    return &pb25d04dc6.LocationServiceImpl{}
}

func (inst* pb25d04dc62_ilocations_LocationServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb25d04dc6.LocationServiceImpl)
	nop(ie, com)

	
    com.LocationDAO = inst.getLocationDAO(ie)


    return nil
}


func (inst*pb25d04dc62_ilocations_LocationServiceImpl) getLocationDAO(ie application.InjectionExt)pfdd8fdb8d.LocationDAO{
    return ie.GetComponent("#alias-fdd8fdb8d991bf9dc6045188ca258fc6-LocationDAO").(pfdd8fdb8d.LocationDAO)
}



// type pd7f15ffc3.LocalRepositoryServiceImpl in package:github.com/bitwormhole/wpm/src/main/golang/implements/irepositories
//
// id:com-d7f15ffc3700ffe8-irepositories-LocalRepositoryServiceImpl
// class:
// alias:alias-fc892d171029572157d3f320608a6038-LocalRepositoryService
// scope:singleton
//
type pd7f15ffc37_irepositories_LocalRepositoryServiceImpl struct {
}

func (inst* pd7f15ffc37_irepositories_LocalRepositoryServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d7f15ffc3700ffe8-irepositories-LocalRepositoryServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-fc892d171029572157d3f320608a6038-LocalRepositoryService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd7f15ffc37_irepositories_LocalRepositoryServiceImpl) new() any {
    return &pd7f15ffc3.LocalRepositoryServiceImpl{}
}

func (inst* pd7f15ffc37_irepositories_LocalRepositoryServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd7f15ffc3.LocalRepositoryServiceImpl)
	nop(ie, com)

	
    com.LocationService = inst.getLocationService(ie)


    return nil
}


func (inst*pd7f15ffc37_irepositories_LocalRepositoryServiceImpl) getLocationService(ie application.InjectionExt)pfdd8fdb8d.Service{
    return ie.GetComponent("#alias-fdd8fdb8d991bf9dc6045188ca258fc6-Service").(pfdd8fdb8d.Service)
}


