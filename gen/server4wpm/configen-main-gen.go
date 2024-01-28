package server4wpm

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p1eb0a0660b_impintenttemplates_DaoImpl{})
    inst.register(&p1eb0a0660b_impintenttemplates_ServiceImpl{})
    inst.register(&p233917231e_imprepositories_ImportRepositoryServiceImpl{})
    inst.register(&p233917231e_imprepositories_LocalRepositoryDao{})
    inst.register(&p233917231e_imprepositories_LocalRepositoryService{})
    inst.register(&p233917231e_imprepositories_RemoteRepositoryDao{})
    inst.register(&p233917231e_imprepositories_RemoteRepositoryService{})
    inst.register(&p31f30d9aca_impintentqueues_IntentQueueService{})
    inst.register(&p337e784ada_impsettings_DaoImpl{})
    inst.register(&p337e784ada_impsettings_ServiceImpl{})
    inst.register(&p3e97fbda8f_itreeroots_TreeRootDaoImpl{})
    inst.register(&p3e97fbda8f_itreeroots_TreeRootServiceImpl{})
    inst.register(&p3e97fbda8f_itreeroots_TreeRootStateLoaderImpl{})
    inst.register(&p3ededc14a5_database_ThisGroup{})
    inst.register(&p415d59ee4c_admin_AboutController{})
    inst.register(&p415d59ee4c_admin_AuthsController{})
    inst.register(&p415d59ee4c_admin_BackupsController{})
    inst.register(&p415d59ee4c_admin_ContentTypesController{})
    inst.register(&p415d59ee4c_admin_ExampleController{})
    inst.register(&p415d59ee4c_admin_ExecutablesController{})
    inst.register(&p415d59ee4c_admin_FileQueryController{})
    inst.register(&p415d59ee4c_admin_ImportRepositoryController{})
    inst.register(&p415d59ee4c_admin_InitController{})
    inst.register(&p415d59ee4c_admin_IntentQueueController{})
    inst.register(&p415d59ee4c_admin_IntentTemplateController{})
    inst.register(&p415d59ee4c_admin_LocalRepositoryController{})
    inst.register(&p415d59ee4c_admin_MediaController{})
    inst.register(&p415d59ee4c_admin_ProjectController{})
    inst.register(&p415d59ee4c_admin_RemoteRepositoryController{})
    inst.register(&p415d59ee4c_admin_SoftwarePackageController{})
    inst.register(&p415d59ee4c_admin_SoftwareSetController{})
    inst.register(&p415d59ee4c_admin_StatisticsController{})
    inst.register(&p415d59ee4c_admin_UploadController{})
    inst.register(&p72b9b4e1d8_controllers_HTTP404Controller{})
    inst.register(&p7fca5d6e66_impexecutables_DaoImpl{})
    inst.register(&p7fca5d6e66_impexecutables_ExecutableStateLoaderImpl{})
    inst.register(&p7fca5d6e66_impexecutables_ServiceImpl{})
    inst.register(&p8cd6426648_impmedia_DaoImpl{})
    inst.register(&p8cd6426648_impmedia_ServiceImpl{})
    inst.register(&p8f4406e5f7_impsoftware_SoftwarePackageDaoImpl{})
    inst.register(&p8f4406e5f7_impsoftware_SoftwarePackageService{})
    inst.register(&p8f4406e5f7_impsoftware_SoftwareSetService{})
    inst.register(&p9b527161a5_ifilequery_ServiceImpl{})
    inst.register(&pa50f35ce65_istatistics_StatisticServiceImpl{})
    inst.register(&pb5fa9685f8_implocations_DaoImpl{})
    inst.register(&pb5fa9685f8_implocations_ServiceImpl{})
    inst.register(&pc347f939ed_impbackups_DaoImpl{})
    inst.register(&pc347f939ed_impbackups_ServiceImpl{})
    inst.register(&pd11deb9a07_impauths_DaoImpl{})
    inst.register(&pd11deb9a07_impauths_ServiceImpl{})
    inst.register(&pd5537a67b8_impcontenttypes_DaoImpl{})
    inst.register(&pd5537a67b8_impcontenttypes_ServiceImpl{})
    inst.register(&pf441dbb74d_impprojects_DaoImpl{})
    inst.register(&pf441dbb74d_impprojects_ServiceImpl{})
    inst.register(&pfba77a8fb5_example_DaoImpl{})
    inst.register(&pfba77a8fb5_example_ServiceImpl{})


    return nil
}
