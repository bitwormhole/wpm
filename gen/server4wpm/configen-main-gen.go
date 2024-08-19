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

    
    inst.register(&p0d7c61d398_ibackups_DaoImpl{})
    inst.register(&p0d7c61d398_ibackups_ServiceImpl{})
    inst.register(&p17f7c6f875_isettings_DaoImpl{})
    inst.register(&p17f7c6f875_isettings_ServiceImpl{})
    inst.register(&p3220b15b16_ilocations_DaoImpl{})
    inst.register(&p3220b15b16_ilocations_ServiceImpl{})
    inst.register(&p35246d3d24_imedia_DaoImpl{})
    inst.register(&p35246d3d24_imedia_ServiceImpl{})
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
    inst.register(&p415d59ee4c_admin_RepositoryWorktreeProjectController{})
    inst.register(&p415d59ee4c_admin_SoftwarePackageController{})
    inst.register(&p415d59ee4c_admin_SoftwareSetController{})
    inst.register(&p415d59ee4c_admin_StatisticsController{})
    inst.register(&p415d59ee4c_admin_UploadController{})
    inst.register(&p42fcb1358a_iintentqueues_IntentQueueService{})
    inst.register(&p72b9b4e1d8_controllers_HTTP404Controller{})
    inst.register(&p7974a74f95_iprojects_DaoImpl{})
    inst.register(&p7974a74f95_iprojects_ServiceImpl{})
    inst.register(&p853e9edeaa_irepositoryworktreeproject_RepositoryWorktreeProjectServiceImpl{})
    inst.register(&p9b527161a5_ifilequery_ServiceImpl{})
    inst.register(&pa4eaf0bfff_icontenttypes_DaoImpl{})
    inst.register(&pa4eaf0bfff_icontenttypes_ServiceImpl{})
    inst.register(&pa50f35ce65_istatistics_StatisticServiceImpl{})
    inst.register(&pa521fb5ff6_iintenttemplates_DaoImpl{})
    inst.register(&pa521fb5ff6_iintenttemplates_ServiceImpl{})
    inst.register(&pac9dcde887_irepositories_ImportRepositoryServiceImpl{})
    inst.register(&pac9dcde887_irepositories_LocalRepositoryDao{})
    inst.register(&pac9dcde887_irepositories_LocalRepositoryService{})
    inst.register(&pac9dcde887_irepositories_RemoteRepositoryDao{})
    inst.register(&pac9dcde887_irepositories_RemoteRepositoryService{})
    inst.register(&pbaa181ebec_iexecutables_DaoImpl{})
    inst.register(&pbaa181ebec_iexecutables_ExecutableStateLoaderImpl{})
    inst.register(&pbaa181ebec_iexecutables_ServiceImpl{})
    inst.register(&pf04311b572_isoftware_SoftwarePackageDaoImpl{})
    inst.register(&pf04311b572_isoftware_SoftwarePackageService{})
    inst.register(&pf04311b572_isoftware_SoftwareSetService{})
    inst.register(&pfba77a8fb5_example_DaoImpl{})
    inst.register(&pfba77a8fb5_example_ServiceImpl{})
    inst.register(&pfcb564d0df_iauths_DaoImpl{})
    inst.register(&pfcb564d0df_iauths_ServiceImpl{})


    return nil
}
