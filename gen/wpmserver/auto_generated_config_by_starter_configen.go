// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmserver

import (
	store0x8467b3 "github.com/bitwormhole/gitlib/git/store"
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
	caches0x9d186b "github.com/bitwormhole/wpm/server/components/caches"
	intents0x8557f3 "github.com/bitwormhole/wpm/server/components/intents"
	v3filters0xa6552a "github.com/bitwormhole/wpm/server/components/intents/v3filters"
	dao0x5af8d0 "github.com/bitwormhole/wpm/server/data/dao"
	dbagent0x9f90fb "github.com/bitwormhole/wpm/server/data/dbagent"
	service0x3e063d "github.com/bitwormhole/wpm/server/service"
	support0xf47d7f "github.com/bitwormhole/wpm/server/support"
	appruntime0x8dfe0a "github.com/bitwormhole/wpm/server/support/appruntime"
	backups0xe44d54 "github.com/bitwormhole/wpm/server/support/backups"
	caches0xd7996d "github.com/bitwormhole/wpm/server/support/caches"
	checkupdate0xea1855 "github.com/bitwormhole/wpm/server/support/checkupdate"
	contenttypes0x61ca37 "github.com/bitwormhole/wpm/server/support/contenttypes"
	executables0xd3773a "github.com/bitwormhole/wpm/server/support/executables"
	httpclient0xf20fe2 "github.com/bitwormhole/wpm/server/support/httpclient"
	impldao0x73998b "github.com/bitwormhole/wpm/server/support/impldao"
	implservice0x22327c "github.com/bitwormhole/wpm/server/support/implservice"
	init0xc984bc "github.com/bitwormhole/wpm/server/support/init"
	intents0x8ee0e0 "github.com/bitwormhole/wpm/server/support/intents"
	intenttemplates0x2e3dcf "github.com/bitwormhole/wpm/server/support/intenttemplates"
	locations0xb36349 "github.com/bitwormhole/wpm/server/support/locations"
	mediae0xf005e2 "github.com/bitwormhole/wpm/server/support/mediae"
	namespaces0xceefcf "github.com/bitwormhole/wpm/server/support/namespaces"
	platforms0xb539c0 "github.com/bitwormhole/wpm/server/support/platforms"
	plugins0x82e34b "github.com/bitwormhole/wpm/server/support/plugins"
	presets0x875f8b "github.com/bitwormhole/wpm/server/support/presets"
	projects0x4d85c7 "github.com/bitwormhole/wpm/server/support/projects"
	repositories0x637d5e "github.com/bitwormhole/wpm/server/support/repositories"
	repositoryworktreeproject0x399028 "github.com/bitwormhole/wpm/server/support/repositoryworktreeproject"
	settings0x19237d "github.com/bitwormhole/wpm/server/support/settings"
	setup0xd9ff02 "github.com/bitwormhole/wpm/server/support/setup"
	worktrees0xa762f3 "github.com/bitwormhole/wpm/server/support/worktrees"
	filequery0xca51d2 "github.com/bitwormhole/wpm/server/utils/filequery"
	handlers0x162741 "github.com/bitwormhole/wpm/server/utils/filequery/handlers"
	controller0x9dc399 "github.com/bitwormhole/wpm/server/web/controller"
	filter0x8aa8f6 "github.com/bitwormhole/wpm/server/web/filter"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-v3filters0xa6552a.CLIMakerFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com0-v3filters0xa6552a.CLIMakerFilter").Class("wpm-intent-filter").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCLIMakerFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-v3filters0xa6552a.CLIRunnerFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com1-v3filters0xa6552a.CLIRunnerFilter").Class("wpm-intent-filter").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCLIRunnerFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-v3filters0xa6552a.ExampleFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com2-v3filters0xa6552a.ExampleFilter").Class("wpm-intent-filter").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-v3filters0xa6552a.PrepareActionFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com3-v3filters0xa6552a.PrepareActionFilter").Class("wpm-intent-filter").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPrepareActionFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com4-v3filters0xa6552a.PreparePropertiesFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com4-v3filters0xa6552a.PreparePropertiesFilter").Class("wpm-intent-filter").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPreparePropertiesFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com5-v3filters0xa6552a.CheckTemplateFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com5-v3filters0xa6552a.CheckTemplateFilter").Class("wpm-intent-filter").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCheckTemplateFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com6-v3filters0xa6552a.FindTemplateFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com6-v3filters0xa6552a.FindTemplateFilter").Class("wpm-intent-filter").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFindTemplateFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: GormDBAgent
	cominfobuilder.Next()
	cominfobuilder.ID("GormDBAgent").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComGormDBAgentImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: AppRuntimeService
	cominfobuilder.Next()
	cominfobuilder.ID("AppRuntimeService").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpAppRuntimeService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com9-backups0xe44d54.WpmBackupController
	cominfobuilder.Next()
	cominfobuilder.ID("com9-backups0xe44d54.WpmBackupController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComWpmBackupController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: wpm-database-backup-dao
	cominfobuilder.Next()
	cominfobuilder.ID("wpm-database-backup-dao").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpBackupServiceDAO{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: DatabaseBackupService
	cominfobuilder.Next()
	cominfobuilder.ID("DatabaseBackupService").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpBackupService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: CacheService
	cominfobuilder.Next()
	cominfobuilder.ID("CacheService").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpCacheManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: CheckUpdateService
	cominfobuilder.Next()
	cominfobuilder.ID("CheckUpdateService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTheCheckUpdateServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com14-contenttypes0x61ca37.ContentTypeController
	cominfobuilder.Next()
	cominfobuilder.ID("com14-contenttypes0x61ca37.ContentTypeController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComContentTypeController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ContentTypeDAO
	cominfobuilder.Next()
	cominfobuilder.ID("ContentTypeDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectTypeDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com16-contenttypes0x61ca37.ProjectTypeImportController
	cominfobuilder.Next()
	cominfobuilder.ID("com16-contenttypes0x61ca37.ProjectTypeImportController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectTypeImportController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ProjectTypeImportService
	cominfobuilder.Next()
	cominfobuilder.ID("ProjectTypeImportService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectTypeImportServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ContentTypeService
	cominfobuilder.Next()
	cominfobuilder.ID("ContentTypeService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectTypeServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com19-executables0xd3773a.ExecutableController
	cominfobuilder.Next()
	cominfobuilder.ID("com19-executables0xd3773a.ExecutableController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ExecutableDAO
	cominfobuilder.Next()
	cominfobuilder.ID("ExecutableDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com21-executables0xd3773a.ExecutableImportController
	cominfobuilder.Next()
	cominfobuilder.ID("com21-executables0xd3773a.ExecutableImportController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableImportController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ExecutableImportService
	cominfobuilder.Next()
	cominfobuilder.ID("ExecutableImportService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableImportServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ExecutableService
	cominfobuilder.Next()
	cominfobuilder.ID("ExecutableService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExecutableServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: HTTPClientService
	cominfobuilder.Next()
	cominfobuilder.ID("HTTPClientService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpHTTPClientService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ExampleDAO
	cominfobuilder.Next()
	cominfobuilder.ID("ExampleDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: AboutService
	cominfobuilder.Next()
	cominfobuilder.ID("AboutService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAboutServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: AppDataService
	cominfobuilder.Next()
	cominfobuilder.ID("AppDataService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAppDataServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ExampleService
	cominfobuilder.Next()
	cominfobuilder.ID("ExampleService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: FileQueryService
	cominfobuilder.Next()
	cominfobuilder.ID("FileQueryService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFileQueryServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: FileSystemService
	cominfobuilder.Next()
	cominfobuilder.ID("FileSystemService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFileSystemServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: UUIDGenService
	cominfobuilder.Next()
	cominfobuilder.ID("UUIDGenService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUUIDGenServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com32-init0xc984bc.WpmInitController
	cominfobuilder.Next()
	cominfobuilder.ID("com32-init0xc984bc.WpmInitController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComWpmInitController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: InitService
	cominfobuilder.Next()
	cominfobuilder.ID("InitService").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpInitService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com34-intents0x8ee0e0.RunIntentController
	cominfobuilder.Next()
	cominfobuilder.ID("com34-intents0x8ee0e0.RunIntentController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRunIntentController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: wpm-intent-filter-manager
	cominfobuilder.Next()
	cominfobuilder.ID("wpm-intent-filter-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFilterManagerImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: IntentHandlerService
	cominfobuilder.Next()
	cominfobuilder.ID("IntentHandlerService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentHandlerImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: IntentService
	cominfobuilder.Next()
	cominfobuilder.ID("IntentService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRunIntentServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com38-intenttemplates0x2e3dcf.IntentTemplateController
	cominfobuilder.Next()
	cominfobuilder.ID("com38-intenttemplates0x2e3dcf.IntentTemplateController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentTemplateController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: IntentTemplateDAO
	cominfobuilder.Next()
	cominfobuilder.ID("IntentTemplateDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentTemplateDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: IntentTemplateService
	cominfobuilder.Next()
	cominfobuilder.ID("IntentTemplateService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComIntentTemplateServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: LocationDAO
	cominfobuilder.Next()
	cominfobuilder.ID("LocationDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpLocationDao{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: LocationService
	cominfobuilder.Next()
	cominfobuilder.ID("LocationService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpLocationService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com43-mediae0xf005e2.MediaController
	cominfobuilder.Next()
	cominfobuilder.ID("com43-mediae0xf005e2.MediaController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMediaController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: MediaDAO
	cominfobuilder.Next()
	cominfobuilder.ID("MediaDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMediaDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: MediaService
	cominfobuilder.Next()
	cominfobuilder.ID("MediaService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMediaServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com46-namespaces0xceefcf.NamespaceController
	cominfobuilder.Next()
	cominfobuilder.ID("com46-namespaces0xceefcf.NamespaceController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComNamespaceController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: NamespaceDAO
	cominfobuilder.Next()
	cominfobuilder.ID("NamespaceDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpNamespaceDao{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: NamespaceService
	cominfobuilder.Next()
	cominfobuilder.ID("NamespaceService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpNamespaceService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com49-platforms0xb539c0.LinuxPlatformServiceImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com49-platforms0xb539c0.LinuxPlatformServiceImpl").Class("PlatformProviderRegistry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLinuxPlatformServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: PlatformService
	cominfobuilder.Next()
	cominfobuilder.ID("PlatformService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPlatformServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ProfileService
	cominfobuilder.Next()
	cominfobuilder.ID("ProfileService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProfileServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com52-platforms0xb539c0.WindowsPlatformServiceImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com52-platforms0xb539c0.WindowsPlatformServiceImpl").Class("PlatformProviderRegistry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComWindowsPlatformServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com53-plugins0x82e34b.SoftwarePackageController
	cominfobuilder.Next()
	cominfobuilder.ID("com53-plugins0x82e34b.SoftwarePackageController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComSoftwarePackageController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: SoftwarePackageDAO
	cominfobuilder.Next()
	cominfobuilder.ID("SoftwarePackageDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPluginDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: SoftwarePackageService
	cominfobuilder.Next()
	cominfobuilder.ID("SoftwarePackageService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPluginServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com56-plugins0x82e34b.SoftwareSetController
	cominfobuilder.Next()
	cominfobuilder.ID("com56-plugins0x82e34b.SoftwareSetController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComSoftwareSetController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: SoftwareSetService
	cominfobuilder.Next()
	cominfobuilder.ID("SoftwareSetService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpSoftwareSetService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: PresetCache
	cominfobuilder.Next()
	cominfobuilder.ID("PresetCache").Class("wpm-cache-provider").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCacheProvider{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: PresetService
	cominfobuilder.Next()
	cominfobuilder.ID("PresetService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpPresetService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com60-projects0x4d85c7.ProjectController
	cominfobuilder.Next()
	cominfobuilder.ID("com60-projects0x4d85c7.ProjectController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ProjectDAO
	cominfobuilder.Next()
	cominfobuilder.ID("ProjectDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com62-projects0x4d85c7.ProjectImportController
	cominfobuilder.Next()
	cominfobuilder.ID("com62-projects0x4d85c7.ProjectImportController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectImportController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ProjectImportService
	cominfobuilder.Next()
	cominfobuilder.ID("ProjectImportService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectImportServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: ProjectService
	cominfobuilder.Next()
	cominfobuilder.ID("ProjectService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComProjectServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com65-repositories0x637d5e.LocalRepositoryController
	cominfobuilder.Next()
	cominfobuilder.ID("com65-repositories0x637d5e.LocalRepositoryController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalRepositoryController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: LocalRepositoryDAO
	cominfobuilder.Next()
	cominfobuilder.ID("LocalRepositoryDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: LocalRepositoryFinder
	cominfobuilder.Next()
	cominfobuilder.ID("LocalRepositoryFinder").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalRepositoryFinderImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: LocalRepositoryService
	cominfobuilder.Next()
	cominfobuilder.ID("LocalRepositoryService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalRepositoryServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: LocalRepositoryStateLoader
	cominfobuilder.Next()
	cominfobuilder.ID("LocalRepositoryStateLoader").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLocalRepositoryStateLoaderImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: MainRepositoryService
	cominfobuilder.Next()
	cominfobuilder.ID("MainRepositoryService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMainRepositoryServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com71-repositories0x637d5e.RemoteRepositoryController
	cominfobuilder.Next()
	cominfobuilder.ID("com71-repositories0x637d5e.RemoteRepositoryController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRemoteRepositoryController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: RemoteRepositoryService
	cominfobuilder.Next()
	cominfobuilder.ID("RemoteRepositoryService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRemoteRepositoryServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com73-repositories0x637d5e.RepositoryImportController
	cominfobuilder.Next()
	cominfobuilder.ID("com73-repositories0x637d5e.RepositoryImportController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryImportController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: RepositoryImportService
	cominfobuilder.Next()
	cominfobuilder.ID("RepositoryImportService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepositoryImportServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com75-repositoryworktreeproject0x399028.RepoWorktreeProjectController
	cominfobuilder.Next()
	cominfobuilder.ID("com75-repositoryworktreeproject0x399028.RepoWorktreeProjectController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRepoWorktreeProjectController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: RepositoryWorktreeProjectService
	cominfobuilder.Next()
	cominfobuilder.ID("RepositoryWorktreeProjectService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRWPServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com77-settings0x19237d.SettingController
	cominfobuilder.Next()
	cominfobuilder.ID("com77-settings0x19237d.SettingController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComSettingController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: SettingDAO
	cominfobuilder.Next()
	cominfobuilder.ID("SettingDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComSettingDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: SettingService
	cominfobuilder.Next()
	cominfobuilder.ID("SettingService").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComSettingServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com80-setup0xd9ff02.WpmSetupController
	cominfobuilder.Next()
	cominfobuilder.ID("com80-setup0xd9ff02.WpmSetupController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComWpmSetupController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: SetupService
	cominfobuilder.Next()
	cominfobuilder.ID("SetupService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpSetupService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: WorktreeDAO
	cominfobuilder.Next()
	cominfobuilder.ID("WorktreeDAO").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpWorktreeDao{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: WorktreeService
	cominfobuilder.Next()
	cominfobuilder.ID("WorktreeService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComImpWorktreeService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com84-support0xf47d7f.WpmDataSource
	cominfobuilder.Next()
	cominfobuilder.ID("com84-support0xf47d7f.WpmDataSource").Class("starter-gorm-source-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComWpmDataSource{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com85-handlers0x162741.FileSystemHandler
	cominfobuilder.Next()
	cominfobuilder.ID("com85-handlers0x162741.FileSystemHandler").Class("filequery-handler-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFileSystemHandler{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com86-controller0x9dc399.AboutController
	cominfobuilder.Next()
	cominfobuilder.ID("com86-controller0x9dc399.AboutController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAboutController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com87-controller0x9dc399.AuthController
	cominfobuilder.Next()
	cominfobuilder.ID("com87-controller0x9dc399.AuthController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAuthController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com88-controller0x9dc399.ExampleController
	cominfobuilder.Next()
	cominfobuilder.ID("com88-controller0x9dc399.ExampleController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com89-controller0x9dc399.FileQueryController
	cominfobuilder.Next()
	cominfobuilder.ID("com89-controller0x9dc399.FileQueryController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFileQueryController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com90-controller0x9dc399.OnlineDocumentExampleController
	cominfobuilder.Next()
	cominfobuilder.ID("com90-controller0x9dc399.OnlineDocumentExampleController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComOnlineDocumentExampleController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com91-controller0x9dc399.PlatformController
	cominfobuilder.Next()
	cominfobuilder.ID("com91-controller0x9dc399.PlatformController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPlatformController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com92-controller0x9dc399.UploadController
	cominfobuilder.Next()
	cominfobuilder.ID("com92-controller0x9dc399.UploadController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUploadController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com93-filter0x8aa8f6.HostFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com93-filter0x8aa8f6.HostFilter").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHostFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com94-filter0x8aa8f6.HTTP404Filter
	cominfobuilder.Next()
	cominfobuilder.ID("com94-filter0x8aa8f6.HTTP404Filter").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHTTP404Filter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCLIMakerFilter : the factory of component: com0-v3filters0xa6552a.CLIMakerFilter
type comFactory4pComCLIMakerFilter struct {

    mPrototype * v3filters0xa6552a.CLIMakerFilter

	

}

func (inst * comFactory4pComCLIMakerFilter) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCLIMakerFilter) newObject() * v3filters0xa6552a.CLIMakerFilter {
	return & v3filters0xa6552a.CLIMakerFilter {}
}

func (inst * comFactory4pComCLIMakerFilter) castObject(instance application.ComponentInstance) * v3filters0xa6552a.CLIMakerFilter {
	return instance.Get().(*v3filters0xa6552a.CLIMakerFilter)
}

func (inst * comFactory4pComCLIMakerFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCLIMakerFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCLIMakerFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCLIMakerFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCLIMakerFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCLIMakerFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCLIRunnerFilter : the factory of component: com1-v3filters0xa6552a.CLIRunnerFilter
type comFactory4pComCLIRunnerFilter struct {

    mPrototype * v3filters0xa6552a.CLIRunnerFilter

	
	mIntentHandlerServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComCLIRunnerFilter) init() application.ComponentFactory {

	
	inst.mIntentHandlerServiceSelector = config.NewInjectionSelector("#IntentHandlerService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCLIRunnerFilter) newObject() * v3filters0xa6552a.CLIRunnerFilter {
	return & v3filters0xa6552a.CLIRunnerFilter {}
}

func (inst * comFactory4pComCLIRunnerFilter) castObject(instance application.ComponentInstance) * v3filters0xa6552a.CLIRunnerFilter {
	return instance.Get().(*v3filters0xa6552a.CLIRunnerFilter)
}

func (inst * comFactory4pComCLIRunnerFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCLIRunnerFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCLIRunnerFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCLIRunnerFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCLIRunnerFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCLIRunnerFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.IntentHandlerService = inst.getterForFieldIntentHandlerServiceSelector(context)
	return context.LastError()
}

//getterForFieldIntentHandlerServiceSelector
func (inst * comFactory4pComCLIRunnerFilter) getterForFieldIntentHandlerServiceSelector (context application.InstanceContext) service0x3e063d.IntentHandlerService {

	o1 := inst.mIntentHandlerServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.IntentHandlerService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com1-v3filters0xa6552a.CLIRunnerFilter")
		eb.Set("field", "IntentHandlerService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.IntentHandlerService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExampleFilter : the factory of component: com2-v3filters0xa6552a.ExampleFilter
type comFactory4pComExampleFilter struct {

    mPrototype * v3filters0xa6552a.ExampleFilter

	

}

func (inst * comFactory4pComExampleFilter) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExampleFilter) newObject() * v3filters0xa6552a.ExampleFilter {
	return & v3filters0xa6552a.ExampleFilter {}
}

func (inst * comFactory4pComExampleFilter) castObject(instance application.ComponentInstance) * v3filters0xa6552a.ExampleFilter {
	return instance.Get().(*v3filters0xa6552a.ExampleFilter)
}

func (inst * comFactory4pComExampleFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExampleFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExampleFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExampleFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPrepareActionFilter : the factory of component: com3-v3filters0xa6552a.PrepareActionFilter
type comFactory4pComPrepareActionFilter struct {

    mPrototype * v3filters0xa6552a.PrepareActionFilter

	

}

func (inst * comFactory4pComPrepareActionFilter) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPrepareActionFilter) newObject() * v3filters0xa6552a.PrepareActionFilter {
	return & v3filters0xa6552a.PrepareActionFilter {}
}

func (inst * comFactory4pComPrepareActionFilter) castObject(instance application.ComponentInstance) * v3filters0xa6552a.PrepareActionFilter {
	return instance.Get().(*v3filters0xa6552a.PrepareActionFilter)
}

func (inst * comFactory4pComPrepareActionFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPrepareActionFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPrepareActionFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPrepareActionFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPrepareActionFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPrepareActionFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPreparePropertiesFilter : the factory of component: com4-v3filters0xa6552a.PreparePropertiesFilter
type comFactory4pComPreparePropertiesFilter struct {

    mPrototype * v3filters0xa6552a.PreparePropertiesFilter

	

}

func (inst * comFactory4pComPreparePropertiesFilter) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPreparePropertiesFilter) newObject() * v3filters0xa6552a.PreparePropertiesFilter {
	return & v3filters0xa6552a.PreparePropertiesFilter {}
}

func (inst * comFactory4pComPreparePropertiesFilter) castObject(instance application.ComponentInstance) * v3filters0xa6552a.PreparePropertiesFilter {
	return instance.Get().(*v3filters0xa6552a.PreparePropertiesFilter)
}

func (inst * comFactory4pComPreparePropertiesFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPreparePropertiesFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPreparePropertiesFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPreparePropertiesFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPreparePropertiesFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPreparePropertiesFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCheckTemplateFilter : the factory of component: com5-v3filters0xa6552a.CheckTemplateFilter
type comFactory4pComCheckTemplateFilter struct {

    mPrototype * v3filters0xa6552a.CheckTemplateFilter

	

}

func (inst * comFactory4pComCheckTemplateFilter) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCheckTemplateFilter) newObject() * v3filters0xa6552a.CheckTemplateFilter {
	return & v3filters0xa6552a.CheckTemplateFilter {}
}

func (inst * comFactory4pComCheckTemplateFilter) castObject(instance application.ComponentInstance) * v3filters0xa6552a.CheckTemplateFilter {
	return instance.Get().(*v3filters0xa6552a.CheckTemplateFilter)
}

func (inst * comFactory4pComCheckTemplateFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCheckTemplateFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCheckTemplateFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCheckTemplateFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCheckTemplateFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCheckTemplateFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFindTemplateFilter : the factory of component: com6-v3filters0xa6552a.FindTemplateFilter
type comFactory4pComFindTemplateFilter struct {

    mPrototype * v3filters0xa6552a.FindTemplateFilter

	
	mIntentTemplateServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComFindTemplateFilter) init() application.ComponentFactory {

	
	inst.mIntentTemplateServiceSelector = config.NewInjectionSelector("#IntentTemplateService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFindTemplateFilter) newObject() * v3filters0xa6552a.FindTemplateFilter {
	return & v3filters0xa6552a.FindTemplateFilter {}
}

func (inst * comFactory4pComFindTemplateFilter) castObject(instance application.ComponentInstance) * v3filters0xa6552a.FindTemplateFilter {
	return instance.Get().(*v3filters0xa6552a.FindTemplateFilter)
}

func (inst * comFactory4pComFindTemplateFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFindTemplateFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFindTemplateFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFindTemplateFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFindTemplateFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFindTemplateFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.IntentTemplateService = inst.getterForFieldIntentTemplateServiceSelector(context)
	return context.LastError()
}

//getterForFieldIntentTemplateServiceSelector
func (inst * comFactory4pComFindTemplateFilter) getterForFieldIntentTemplateServiceSelector (context application.InstanceContext) service0x3e063d.IntentTemplateService {

	o1 := inst.mIntentTemplateServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.IntentTemplateService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com6-v3filters0xa6552a.FindTemplateFilter")
		eb.Set("field", "IntentTemplateService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.IntentTemplateService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComGormDBAgentImpl : the factory of component: GormDBAgent
type comFactory4pComGormDBAgentImpl struct {

    mPrototype * dbagent0x9f90fb.GormDBAgentImpl

	
	mSourcesSelector config.InjectionSelector

}

func (inst * comFactory4pComGormDBAgentImpl) init() application.ComponentFactory {

	
	inst.mSourcesSelector = config.NewInjectionSelector("#starter-gorm-source-manager",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComGormDBAgentImpl) newObject() * dbagent0x9f90fb.GormDBAgentImpl {
	return & dbagent0x9f90fb.GormDBAgentImpl {}
}

func (inst * comFactory4pComGormDBAgentImpl) castObject(instance application.ComponentInstance) * dbagent0x9f90fb.GormDBAgentImpl {
	return instance.Get().(*dbagent0x9f90fb.GormDBAgentImpl)
}

func (inst * comFactory4pComGormDBAgentImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComGormDBAgentImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComGormDBAgentImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComGormDBAgentImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGormDBAgentImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComGormDBAgentImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Sources = inst.getterForFieldSourcesSelector(context)
	return context.LastError()
}

//getterForFieldSourcesSelector
func (inst * comFactory4pComGormDBAgentImpl) getterForFieldSourcesSelector (context application.InstanceContext) datasource0x68a737.SourceManager {

	o1 := inst.mSourcesSelector.GetOne(context)
	o2, ok := o1.(datasource0x68a737.SourceManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "GormDBAgent")
		eb.Set("field", "Sources")
		eb.Set("type1", "?")
		eb.Set("type2", "datasource0x68a737.SourceManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpAppRuntimeService : the factory of component: AppRuntimeService
type comFactory4pComImpAppRuntimeService struct {

    mPrototype * appruntime0x8dfe0a.ImpAppRuntimeService

	
	mFileSystemServiceSelector config.InjectionSelector
	mAppDataServiceSelector config.InjectionSelector
	mMediaServiceSelector config.InjectionSelector
	mEnableBackupSelfSelector config.InjectionSelector

}

func (inst * comFactory4pComImpAppRuntimeService) init() application.ComponentFactory {

	
	inst.mFileSystemServiceSelector = config.NewInjectionSelector("#FileSystemService",nil)
	inst.mAppDataServiceSelector = config.NewInjectionSelector("#AppDataService",nil)
	inst.mMediaServiceSelector = config.NewInjectionSelector("#MediaService",nil)
	inst.mEnableBackupSelfSelector = config.NewInjectionSelector("${wpm.backup-this-exe.enabled}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpAppRuntimeService) newObject() * appruntime0x8dfe0a.ImpAppRuntimeService {
	return & appruntime0x8dfe0a.ImpAppRuntimeService {}
}

func (inst * comFactory4pComImpAppRuntimeService) castObject(instance application.ComponentInstance) * appruntime0x8dfe0a.ImpAppRuntimeService {
	return instance.Get().(*appruntime0x8dfe0a.ImpAppRuntimeService)
}

func (inst * comFactory4pComImpAppRuntimeService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpAppRuntimeService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpAppRuntimeService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpAppRuntimeService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpAppRuntimeService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpAppRuntimeService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.FileSystemService = inst.getterForFieldFileSystemServiceSelector(context)
	obj.AppDataService = inst.getterForFieldAppDataServiceSelector(context)
	obj.MediaService = inst.getterForFieldMediaServiceSelector(context)
	obj.EnableBackupSelf = inst.getterForFieldEnableBackupSelfSelector(context)
	return context.LastError()
}

//getterForFieldFileSystemServiceSelector
func (inst * comFactory4pComImpAppRuntimeService) getterForFieldFileSystemServiceSelector (context application.InstanceContext) service0x3e063d.FileSystemService {

	o1 := inst.mFileSystemServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.FileSystemService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "AppRuntimeService")
		eb.Set("field", "FileSystemService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.FileSystemService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldAppDataServiceSelector
func (inst * comFactory4pComImpAppRuntimeService) getterForFieldAppDataServiceSelector (context application.InstanceContext) service0x3e063d.AppDataService {

	o1 := inst.mAppDataServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AppDataService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "AppRuntimeService")
		eb.Set("field", "AppDataService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AppDataService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldMediaServiceSelector
func (inst * comFactory4pComImpAppRuntimeService) getterForFieldMediaServiceSelector (context application.InstanceContext) service0x3e063d.MediaService {

	o1 := inst.mMediaServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.MediaService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "AppRuntimeService")
		eb.Set("field", "MediaService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.MediaService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldEnableBackupSelfSelector
func (inst * comFactory4pComImpAppRuntimeService) getterForFieldEnableBackupSelfSelector (context application.InstanceContext) bool {
    return inst.mEnableBackupSelfSelector.GetBool(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComWpmBackupController : the factory of component: com9-backups0xe44d54.WpmBackupController
type comFactory4pComWpmBackupController struct {

    mPrototype * backups0xe44d54.WpmBackupController

	
	mBackupServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComWpmBackupController) init() application.ComponentFactory {

	
	inst.mBackupServiceSelector = config.NewInjectionSelector("#DatabaseBackupService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComWpmBackupController) newObject() * backups0xe44d54.WpmBackupController {
	return & backups0xe44d54.WpmBackupController {}
}

func (inst * comFactory4pComWpmBackupController) castObject(instance application.ComponentInstance) * backups0xe44d54.WpmBackupController {
	return instance.Get().(*backups0xe44d54.WpmBackupController)
}

func (inst * comFactory4pComWpmBackupController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComWpmBackupController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComWpmBackupController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComWpmBackupController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWpmBackupController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWpmBackupController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.BackupService = inst.getterForFieldBackupServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldBackupServiceSelector
func (inst * comFactory4pComWpmBackupController) getterForFieldBackupServiceSelector (context application.InstanceContext) service0x3e063d.DatabaseBackupService {

	o1 := inst.mBackupServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.DatabaseBackupService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com9-backups0xe44d54.WpmBackupController")
		eb.Set("field", "BackupService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.DatabaseBackupService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComWpmBackupController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com9-backups0xe44d54.WpmBackupController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpBackupServiceDAO : the factory of component: wpm-database-backup-dao
type comFactory4pComImpBackupServiceDAO struct {

    mPrototype * backups0xe44d54.ImpBackupServiceDAO

	
	mAgentSelector config.InjectionSelector

}

func (inst * comFactory4pComImpBackupServiceDAO) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpBackupServiceDAO) newObject() * backups0xe44d54.ImpBackupServiceDAO {
	return & backups0xe44d54.ImpBackupServiceDAO {}
}

func (inst * comFactory4pComImpBackupServiceDAO) castObject(instance application.ComponentInstance) * backups0xe44d54.ImpBackupServiceDAO {
	return instance.Get().(*backups0xe44d54.ImpBackupServiceDAO)
}

func (inst * comFactory4pComImpBackupServiceDAO) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpBackupServiceDAO) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpBackupServiceDAO) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpBackupServiceDAO) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpBackupServiceDAO) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpBackupServiceDAO) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComImpBackupServiceDAO) getterForFieldAgentSelector (context application.InstanceContext) dbagent0x9f90fb.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(dbagent0x9f90fb.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "wpm-database-backup-dao")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "dbagent0x9f90fb.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpBackupService : the factory of component: DatabaseBackupService
type comFactory4pComImpBackupService struct {

    mPrototype * backups0xe44d54.ImpBackupService

	
	mAppDataServiceSelector config.InjectionSelector
	mFilesysServiceSelector config.InjectionSelector
	mBackupDaoSelector config.InjectionSelector
	mDoDumpSelector config.InjectionSelector

}

func (inst * comFactory4pComImpBackupService) init() application.ComponentFactory {

	
	inst.mAppDataServiceSelector = config.NewInjectionSelector("#AppDataService",nil)
	inst.mFilesysServiceSelector = config.NewInjectionSelector("#FileSystemService",nil)
	inst.mBackupDaoSelector = config.NewInjectionSelector("#wpm-database-backup-dao",nil)
	inst.mDoDumpSelector = config.NewInjectionSelector("${wpm.db.dump.enabled}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpBackupService) newObject() * backups0xe44d54.ImpBackupService {
	return & backups0xe44d54.ImpBackupService {}
}

func (inst * comFactory4pComImpBackupService) castObject(instance application.ComponentInstance) * backups0xe44d54.ImpBackupService {
	return instance.Get().(*backups0xe44d54.ImpBackupService)
}

func (inst * comFactory4pComImpBackupService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpBackupService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpBackupService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpBackupService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpBackupService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpBackupService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AppDataService = inst.getterForFieldAppDataServiceSelector(context)
	obj.FilesysService = inst.getterForFieldFilesysServiceSelector(context)
	obj.BackupDao = inst.getterForFieldBackupDaoSelector(context)
	obj.DoDump = inst.getterForFieldDoDumpSelector(context)
	return context.LastError()
}

//getterForFieldAppDataServiceSelector
func (inst * comFactory4pComImpBackupService) getterForFieldAppDataServiceSelector (context application.InstanceContext) service0x3e063d.AppDataService {

	o1 := inst.mAppDataServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AppDataService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "DatabaseBackupService")
		eb.Set("field", "AppDataService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AppDataService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldFilesysServiceSelector
func (inst * comFactory4pComImpBackupService) getterForFieldFilesysServiceSelector (context application.InstanceContext) service0x3e063d.FileSystemService {

	o1 := inst.mFilesysServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.FileSystemService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "DatabaseBackupService")
		eb.Set("field", "FilesysService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.FileSystemService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldBackupDaoSelector
func (inst * comFactory4pComImpBackupService) getterForFieldBackupDaoSelector (context application.InstanceContext) dao0x5af8d0.Backup {

	o1 := inst.mBackupDaoSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.Backup)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "DatabaseBackupService")
		eb.Set("field", "BackupDao")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.Backup")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldDoDumpSelector
func (inst * comFactory4pComImpBackupService) getterForFieldDoDumpSelector (context application.InstanceContext) bool {
    return inst.mDoDumpSelector.GetBool(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpCacheManager : the factory of component: CacheService
type comFactory4pComImpCacheManager struct {

    mPrototype * caches0xd7996d.ImpCacheManager

	
	mProviderRegistryListSelector config.InjectionSelector

}

func (inst * comFactory4pComImpCacheManager) init() application.ComponentFactory {

	
	inst.mProviderRegistryListSelector = config.NewInjectionSelector(".wpm-cache-provider",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpCacheManager) newObject() * caches0xd7996d.ImpCacheManager {
	return & caches0xd7996d.ImpCacheManager {}
}

func (inst * comFactory4pComImpCacheManager) castObject(instance application.ComponentInstance) * caches0xd7996d.ImpCacheManager {
	return instance.Get().(*caches0xd7996d.ImpCacheManager)
}

func (inst * comFactory4pComImpCacheManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpCacheManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpCacheManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpCacheManager) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpCacheManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpCacheManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ProviderRegistryList = inst.getterForFieldProviderRegistryListSelector(context)
	return context.LastError()
}

//getterForFieldProviderRegistryListSelector
func (inst * comFactory4pComImpCacheManager) getterForFieldProviderRegistryListSelector (context application.InstanceContext) []caches0x9d186b.ProviderRegistry {
	list1 := inst.mProviderRegistryListSelector.GetList(context)
	list2 := make([]caches0x9d186b.ProviderRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(caches0x9d186b.ProviderRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTheCheckUpdateServiceImpl : the factory of component: CheckUpdateService
type comFactory4pComTheCheckUpdateServiceImpl struct {

    mPrototype * checkupdate0xea1855.TheCheckUpdateServiceImpl

	
	mPackagesURLSelector config.InjectionSelector
	mAboutServiceSelector config.InjectionSelector
	mSettingServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComTheCheckUpdateServiceImpl) init() application.ComponentFactory {

	
	inst.mPackagesURLSelector = config.NewInjectionSelector("${wpm.check-update.url}",nil)
	inst.mAboutServiceSelector = config.NewInjectionSelector("#AboutService",nil)
	inst.mSettingServiceSelector = config.NewInjectionSelector("#SettingService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTheCheckUpdateServiceImpl) newObject() * checkupdate0xea1855.TheCheckUpdateServiceImpl {
	return & checkupdate0xea1855.TheCheckUpdateServiceImpl {}
}

func (inst * comFactory4pComTheCheckUpdateServiceImpl) castObject(instance application.ComponentInstance) * checkupdate0xea1855.TheCheckUpdateServiceImpl {
	return instance.Get().(*checkupdate0xea1855.TheCheckUpdateServiceImpl)
}

func (inst * comFactory4pComTheCheckUpdateServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTheCheckUpdateServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTheCheckUpdateServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTheCheckUpdateServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheCheckUpdateServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheCheckUpdateServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.PackagesURL = inst.getterForFieldPackagesURLSelector(context)
	obj.AboutService = inst.getterForFieldAboutServiceSelector(context)
	obj.SettingService = inst.getterForFieldSettingServiceSelector(context)
	return context.LastError()
}

//getterForFieldPackagesURLSelector
func (inst * comFactory4pComTheCheckUpdateServiceImpl) getterForFieldPackagesURLSelector (context application.InstanceContext) string {
    return inst.mPackagesURLSelector.GetString(context)
}

//getterForFieldAboutServiceSelector
func (inst * comFactory4pComTheCheckUpdateServiceImpl) getterForFieldAboutServiceSelector (context application.InstanceContext) service0x3e063d.AboutService {

	o1 := inst.mAboutServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AboutService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "CheckUpdateService")
		eb.Set("field", "AboutService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AboutService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSettingServiceSelector
func (inst * comFactory4pComTheCheckUpdateServiceImpl) getterForFieldSettingServiceSelector (context application.InstanceContext) service0x3e063d.SettingService {

	o1 := inst.mSettingServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.SettingService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "CheckUpdateService")
		eb.Set("field", "SettingService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.SettingService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComContentTypeController : the factory of component: com14-contenttypes0x61ca37.ContentTypeController
type comFactory4pComContentTypeController struct {

    mPrototype * contenttypes0x61ca37.ContentTypeController

	
	mContentTypeServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComContentTypeController) init() application.ComponentFactory {

	
	inst.mContentTypeServiceSelector = config.NewInjectionSelector("#ContentTypeService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComContentTypeController) newObject() * contenttypes0x61ca37.ContentTypeController {
	return & contenttypes0x61ca37.ContentTypeController {}
}

func (inst * comFactory4pComContentTypeController) castObject(instance application.ComponentInstance) * contenttypes0x61ca37.ContentTypeController {
	return instance.Get().(*contenttypes0x61ca37.ContentTypeController)
}

func (inst * comFactory4pComContentTypeController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComContentTypeController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComContentTypeController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComContentTypeController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComContentTypeController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComContentTypeController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ContentTypeService = inst.getterForFieldContentTypeServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldContentTypeServiceSelector
func (inst * comFactory4pComContentTypeController) getterForFieldContentTypeServiceSelector (context application.InstanceContext) service0x3e063d.ContentTypeService {

	o1 := inst.mContentTypeServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ContentTypeService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com14-contenttypes0x61ca37.ContentTypeController")
		eb.Set("field", "ContentTypeService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ContentTypeService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComContentTypeController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com14-contenttypes0x61ca37.ContentTypeController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectTypeDaoImpl : the factory of component: ContentTypeDAO
type comFactory4pComProjectTypeDaoImpl struct {

    mPrototype * contenttypes0x61ca37.ProjectTypeDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectTypeDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectTypeDaoImpl) newObject() * contenttypes0x61ca37.ProjectTypeDaoImpl {
	return & contenttypes0x61ca37.ProjectTypeDaoImpl {}
}

func (inst * comFactory4pComProjectTypeDaoImpl) castObject(instance application.ComponentInstance) * contenttypes0x61ca37.ProjectTypeDaoImpl {
	return instance.Get().(*contenttypes0x61ca37.ProjectTypeDaoImpl)
}

func (inst * comFactory4pComProjectTypeDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectTypeDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectTypeDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectTypeDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectTypeDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectTypeDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComProjectTypeDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) dbagent0x9f90fb.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(dbagent0x9f90fb.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ContentTypeDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "dbagent0x9f90fb.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComProjectTypeDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ContentTypeDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectTypeImportController : the factory of component: com16-contenttypes0x61ca37.ProjectTypeImportController
type comFactory4pComProjectTypeImportController struct {

    mPrototype * contenttypes0x61ca37.ProjectTypeImportController

	
	mProjectTypeImportServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectTypeImportController) init() application.ComponentFactory {

	
	inst.mProjectTypeImportServiceSelector = config.NewInjectionSelector("#ProjectTypeImportService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectTypeImportController) newObject() * contenttypes0x61ca37.ProjectTypeImportController {
	return & contenttypes0x61ca37.ProjectTypeImportController {}
}

func (inst * comFactory4pComProjectTypeImportController) castObject(instance application.ComponentInstance) * contenttypes0x61ca37.ProjectTypeImportController {
	return instance.Get().(*contenttypes0x61ca37.ProjectTypeImportController)
}

func (inst * comFactory4pComProjectTypeImportController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectTypeImportController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectTypeImportController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectTypeImportController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectTypeImportController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectTypeImportController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ProjectTypeImportService = inst.getterForFieldProjectTypeImportServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldProjectTypeImportServiceSelector
func (inst * comFactory4pComProjectTypeImportController) getterForFieldProjectTypeImportServiceSelector (context application.InstanceContext) service0x3e063d.ProjectTypeImportService {

	o1 := inst.mProjectTypeImportServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProjectTypeImportService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com16-contenttypes0x61ca37.ProjectTypeImportController")
		eb.Set("field", "ProjectTypeImportService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProjectTypeImportService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComProjectTypeImportController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com16-contenttypes0x61ca37.ProjectTypeImportController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectTypeImportServiceImpl : the factory of component: ProjectTypeImportService
type comFactory4pComProjectTypeImportServiceImpl struct {

    mPrototype * contenttypes0x61ca37.ProjectTypeImportServiceImpl

	
	mACSelector config.InjectionSelector
	mProjectTypeServiceSelector config.InjectionSelector
	mPresetServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectTypeImportServiceImpl) init() application.ComponentFactory {

	
	inst.mACSelector = config.NewInjectionSelector("context",nil)
	inst.mProjectTypeServiceSelector = config.NewInjectionSelector("#ContentTypeService",nil)
	inst.mPresetServiceSelector = config.NewInjectionSelector("#PresetService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectTypeImportServiceImpl) newObject() * contenttypes0x61ca37.ProjectTypeImportServiceImpl {
	return & contenttypes0x61ca37.ProjectTypeImportServiceImpl {}
}

func (inst * comFactory4pComProjectTypeImportServiceImpl) castObject(instance application.ComponentInstance) * contenttypes0x61ca37.ProjectTypeImportServiceImpl {
	return instance.Get().(*contenttypes0x61ca37.ProjectTypeImportServiceImpl)
}

func (inst * comFactory4pComProjectTypeImportServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectTypeImportServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectTypeImportServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectTypeImportServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectTypeImportServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectTypeImportServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AC = inst.getterForFieldACSelector(context)
	obj.ProjectTypeService = inst.getterForFieldProjectTypeServiceSelector(context)
	obj.PresetService = inst.getterForFieldPresetServiceSelector(context)
	return context.LastError()
}

//getterForFieldACSelector
func (inst * comFactory4pComProjectTypeImportServiceImpl) getterForFieldACSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldProjectTypeServiceSelector
func (inst * comFactory4pComProjectTypeImportServiceImpl) getterForFieldProjectTypeServiceSelector (context application.InstanceContext) service0x3e063d.ContentTypeService {

	o1 := inst.mProjectTypeServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ContentTypeService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectTypeImportService")
		eb.Set("field", "ProjectTypeService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ContentTypeService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldPresetServiceSelector
func (inst * comFactory4pComProjectTypeImportServiceImpl) getterForFieldPresetServiceSelector (context application.InstanceContext) service0x3e063d.PresetService {

	o1 := inst.mPresetServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.PresetService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectTypeImportService")
		eb.Set("field", "PresetService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.PresetService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectTypeServiceImpl : the factory of component: ContentTypeService
type comFactory4pComProjectTypeServiceImpl struct {

    mPrototype * contenttypes0x61ca37.ProjectTypeServiceImpl

	
	mProjectTypeDAOSelector config.InjectionSelector
	mFileSystemServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectTypeServiceImpl) init() application.ComponentFactory {

	
	inst.mProjectTypeDAOSelector = config.NewInjectionSelector("#ContentTypeDAO",nil)
	inst.mFileSystemServiceSelector = config.NewInjectionSelector("#FileSystemService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectTypeServiceImpl) newObject() * contenttypes0x61ca37.ProjectTypeServiceImpl {
	return & contenttypes0x61ca37.ProjectTypeServiceImpl {}
}

func (inst * comFactory4pComProjectTypeServiceImpl) castObject(instance application.ComponentInstance) * contenttypes0x61ca37.ProjectTypeServiceImpl {
	return instance.Get().(*contenttypes0x61ca37.ProjectTypeServiceImpl)
}

func (inst * comFactory4pComProjectTypeServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectTypeServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectTypeServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectTypeServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectTypeServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectTypeServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ProjectTypeDAO = inst.getterForFieldProjectTypeDAOSelector(context)
	obj.FileSystemService = inst.getterForFieldFileSystemServiceSelector(context)
	return context.LastError()
}

//getterForFieldProjectTypeDAOSelector
func (inst * comFactory4pComProjectTypeServiceImpl) getterForFieldProjectTypeDAOSelector (context application.InstanceContext) dao0x5af8d0.ContentTypeDAO {

	o1 := inst.mProjectTypeDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.ContentTypeDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ContentTypeService")
		eb.Set("field", "ProjectTypeDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.ContentTypeDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldFileSystemServiceSelector
func (inst * comFactory4pComProjectTypeServiceImpl) getterForFieldFileSystemServiceSelector (context application.InstanceContext) service0x3e063d.FileSystemService {

	o1 := inst.mFileSystemServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.FileSystemService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ContentTypeService")
		eb.Set("field", "FileSystemService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.FileSystemService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableController : the factory of component: com19-executables0xd3773a.ExecutableController
type comFactory4pComExecutableController struct {

    mPrototype * executables0xd3773a.ExecutableController

	
	mExecutableServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableController) init() application.ComponentFactory {

	
	inst.mExecutableServiceSelector = config.NewInjectionSelector("#ExecutableService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableController) newObject() * executables0xd3773a.ExecutableController {
	return & executables0xd3773a.ExecutableController {}
}

func (inst * comFactory4pComExecutableController) castObject(instance application.ComponentInstance) * executables0xd3773a.ExecutableController {
	return instance.Get().(*executables0xd3773a.ExecutableController)
}

func (inst * comFactory4pComExecutableController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExecutableController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExecutableController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExecutableController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ExecutableService = inst.getterForFieldExecutableServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldExecutableServiceSelector
func (inst * comFactory4pComExecutableController) getterForFieldExecutableServiceSelector (context application.InstanceContext) service0x3e063d.ExecutableService {

	o1 := inst.mExecutableServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ExecutableService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com19-executables0xd3773a.ExecutableController")
		eb.Set("field", "ExecutableService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ExecutableService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComExecutableController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com19-executables0xd3773a.ExecutableController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableDaoImpl : the factory of component: ExecutableDAO
type comFactory4pComExecutableDaoImpl struct {

    mPrototype * executables0xd3773a.ExecutableDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableDaoImpl) newObject() * executables0xd3773a.ExecutableDaoImpl {
	return & executables0xd3773a.ExecutableDaoImpl {}
}

func (inst * comFactory4pComExecutableDaoImpl) castObject(instance application.ComponentInstance) * executables0xd3773a.ExecutableDaoImpl {
	return instance.Get().(*executables0xd3773a.ExecutableDaoImpl)
}

func (inst * comFactory4pComExecutableDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExecutableDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExecutableDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExecutableDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComExecutableDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) dbagent0x9f90fb.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(dbagent0x9f90fb.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "dbagent0x9f90fb.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComExecutableDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableImportController : the factory of component: com21-executables0xd3773a.ExecutableImportController
type comFactory4pComExecutableImportController struct {

    mPrototype * executables0xd3773a.ExecutableImportController

	
	mExecutableImportServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableImportController) init() application.ComponentFactory {

	
	inst.mExecutableImportServiceSelector = config.NewInjectionSelector("#ExecutableImportService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableImportController) newObject() * executables0xd3773a.ExecutableImportController {
	return & executables0xd3773a.ExecutableImportController {}
}

func (inst * comFactory4pComExecutableImportController) castObject(instance application.ComponentInstance) * executables0xd3773a.ExecutableImportController {
	return instance.Get().(*executables0xd3773a.ExecutableImportController)
}

func (inst * comFactory4pComExecutableImportController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExecutableImportController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExecutableImportController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExecutableImportController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableImportController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableImportController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ExecutableImportService = inst.getterForFieldExecutableImportServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldExecutableImportServiceSelector
func (inst * comFactory4pComExecutableImportController) getterForFieldExecutableImportServiceSelector (context application.InstanceContext) service0x3e063d.ExecutableImportService {

	o1 := inst.mExecutableImportServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ExecutableImportService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com21-executables0xd3773a.ExecutableImportController")
		eb.Set("field", "ExecutableImportService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ExecutableImportService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComExecutableImportController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com21-executables0xd3773a.ExecutableImportController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableImportServiceImpl : the factory of component: ExecutableImportService
type comFactory4pComExecutableImportServiceImpl struct {

    mPrototype * executables0xd3773a.ExecutableImportServiceImpl

	
	mACSelector config.InjectionSelector
	mExecutableServiceSelector config.InjectionSelector
	mFileSystemServiceSelector config.InjectionSelector
	mPresetServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableImportServiceImpl) init() application.ComponentFactory {

	
	inst.mACSelector = config.NewInjectionSelector("context",nil)
	inst.mExecutableServiceSelector = config.NewInjectionSelector("#ExecutableService",nil)
	inst.mFileSystemServiceSelector = config.NewInjectionSelector("#FileSystemService",nil)
	inst.mPresetServiceSelector = config.NewInjectionSelector("#PresetService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableImportServiceImpl) newObject() * executables0xd3773a.ExecutableImportServiceImpl {
	return & executables0xd3773a.ExecutableImportServiceImpl {}
}

func (inst * comFactory4pComExecutableImportServiceImpl) castObject(instance application.ComponentInstance) * executables0xd3773a.ExecutableImportServiceImpl {
	return instance.Get().(*executables0xd3773a.ExecutableImportServiceImpl)
}

func (inst * comFactory4pComExecutableImportServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExecutableImportServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExecutableImportServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExecutableImportServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableImportServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableImportServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AC = inst.getterForFieldACSelector(context)
	obj.ExecutableService = inst.getterForFieldExecutableServiceSelector(context)
	obj.FileSystemService = inst.getterForFieldFileSystemServiceSelector(context)
	obj.PresetService = inst.getterForFieldPresetServiceSelector(context)
	return context.LastError()
}

//getterForFieldACSelector
func (inst * comFactory4pComExecutableImportServiceImpl) getterForFieldACSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldExecutableServiceSelector
func (inst * comFactory4pComExecutableImportServiceImpl) getterForFieldExecutableServiceSelector (context application.InstanceContext) service0x3e063d.ExecutableService {

	o1 := inst.mExecutableServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ExecutableService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableImportService")
		eb.Set("field", "ExecutableService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ExecutableService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldFileSystemServiceSelector
func (inst * comFactory4pComExecutableImportServiceImpl) getterForFieldFileSystemServiceSelector (context application.InstanceContext) service0x3e063d.FileSystemService {

	o1 := inst.mFileSystemServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.FileSystemService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableImportService")
		eb.Set("field", "FileSystemService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.FileSystemService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldPresetServiceSelector
func (inst * comFactory4pComExecutableImportServiceImpl) getterForFieldPresetServiceSelector (context application.InstanceContext) service0x3e063d.PresetService {

	o1 := inst.mPresetServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.PresetService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableImportService")
		eb.Set("field", "PresetService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.PresetService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExecutableServiceImpl : the factory of component: ExecutableService
type comFactory4pComExecutableServiceImpl struct {

    mPrototype * executables0xd3773a.ExecutableServiceImpl

	
	mExecutableDAOSelector config.InjectionSelector
	mFileSystemServiceSelector config.InjectionSelector
	mLocationServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComExecutableServiceImpl) init() application.ComponentFactory {

	
	inst.mExecutableDAOSelector = config.NewInjectionSelector("#ExecutableDAO",nil)
	inst.mFileSystemServiceSelector = config.NewInjectionSelector("#FileSystemService",nil)
	inst.mLocationServiceSelector = config.NewInjectionSelector("#LocationService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExecutableServiceImpl) newObject() * executables0xd3773a.ExecutableServiceImpl {
	return & executables0xd3773a.ExecutableServiceImpl {}
}

func (inst * comFactory4pComExecutableServiceImpl) castObject(instance application.ComponentInstance) * executables0xd3773a.ExecutableServiceImpl {
	return instance.Get().(*executables0xd3773a.ExecutableServiceImpl)
}

func (inst * comFactory4pComExecutableServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExecutableServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExecutableServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExecutableServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExecutableServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ExecutableDAO = inst.getterForFieldExecutableDAOSelector(context)
	obj.FileSystemService = inst.getterForFieldFileSystemServiceSelector(context)
	obj.LocationService = inst.getterForFieldLocationServiceSelector(context)
	return context.LastError()
}

//getterForFieldExecutableDAOSelector
func (inst * comFactory4pComExecutableServiceImpl) getterForFieldExecutableDAOSelector (context application.InstanceContext) dao0x5af8d0.ExecutableDAO {

	o1 := inst.mExecutableDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.ExecutableDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableService")
		eb.Set("field", "ExecutableDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.ExecutableDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldFileSystemServiceSelector
func (inst * comFactory4pComExecutableServiceImpl) getterForFieldFileSystemServiceSelector (context application.InstanceContext) service0x3e063d.FileSystemService {

	o1 := inst.mFileSystemServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.FileSystemService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableService")
		eb.Set("field", "FileSystemService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.FileSystemService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLocationServiceSelector
func (inst * comFactory4pComExecutableServiceImpl) getterForFieldLocationServiceSelector (context application.InstanceContext) service0x3e063d.LocationService {

	o1 := inst.mLocationServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocationService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExecutableService")
		eb.Set("field", "LocationService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocationService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpHTTPClientService : the factory of component: HTTPClientService
type comFactory4pComImpHTTPClientService struct {

    mPrototype * httpclient0xf20fe2.ImpHTTPClientService

	
	mMaxContentLengthSelector config.InjectionSelector

}

func (inst * comFactory4pComImpHTTPClientService) init() application.ComponentFactory {

	
	inst.mMaxContentLengthSelector = config.NewInjectionSelector("${wpm.httpclient.max-content-length}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpHTTPClientService) newObject() * httpclient0xf20fe2.ImpHTTPClientService {
	return & httpclient0xf20fe2.ImpHTTPClientService {}
}

func (inst * comFactory4pComImpHTTPClientService) castObject(instance application.ComponentInstance) * httpclient0xf20fe2.ImpHTTPClientService {
	return instance.Get().(*httpclient0xf20fe2.ImpHTTPClientService)
}

func (inst * comFactory4pComImpHTTPClientService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpHTTPClientService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpHTTPClientService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpHTTPClientService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpHTTPClientService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpHTTPClientService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.MaxContentLength = inst.getterForFieldMaxContentLengthSelector(context)
	return context.LastError()
}

//getterForFieldMaxContentLengthSelector
func (inst * comFactory4pComImpHTTPClientService) getterForFieldMaxContentLengthSelector (context application.InstanceContext) int {
    return inst.mMaxContentLengthSelector.GetInt(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExampleDaoImpl : the factory of component: ExampleDAO
type comFactory4pComExampleDaoImpl struct {

    mPrototype * impldao0x73998b.ExampleDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComExampleDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExampleDaoImpl) newObject() * impldao0x73998b.ExampleDaoImpl {
	return & impldao0x73998b.ExampleDaoImpl {}
}

func (inst * comFactory4pComExampleDaoImpl) castObject(instance application.ComponentInstance) * impldao0x73998b.ExampleDaoImpl {
	return instance.Get().(*impldao0x73998b.ExampleDaoImpl)
}

func (inst * comFactory4pComExampleDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExampleDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExampleDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExampleDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComExampleDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) dbagent0x9f90fb.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(dbagent0x9f90fb.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExampleDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "dbagent0x9f90fb.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComExampleDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ExampleDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAboutServiceImpl : the factory of component: AboutService
type comFactory4pComAboutServiceImpl struct {

    mPrototype * implservice0x22327c.AboutServiceImpl

	
	mProfileSelector config.InjectionSelector
	mNameSelector config.InjectionSelector
	mTitleSelector config.InjectionSelector
	mCopyrightSelector config.InjectionSelector
	mServerPortSelector config.InjectionSelector
	mEnableDebugSelector config.InjectionSelector
	mPlatformServiceSelector config.InjectionSelector
	mProfileServiceSelector config.InjectionSelector
	mAppRuntimeServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComAboutServiceImpl) init() application.ComponentFactory {

	
	inst.mProfileSelector = config.NewInjectionSelector("${application.profiles.active}",nil)
	inst.mNameSelector = config.NewInjectionSelector("${application.about.name}",nil)
	inst.mTitleSelector = config.NewInjectionSelector("${application.about.title}",nil)
	inst.mCopyrightSelector = config.NewInjectionSelector("${application.about.copyright}",nil)
	inst.mServerPortSelector = config.NewInjectionSelector("${server.port}",nil)
	inst.mEnableDebugSelector = config.NewInjectionSelector("${wpm.debug.enabled}",nil)
	inst.mPlatformServiceSelector = config.NewInjectionSelector("#PlatformService",nil)
	inst.mProfileServiceSelector = config.NewInjectionSelector("#ProfileService",nil)
	inst.mAppRuntimeServiceSelector = config.NewInjectionSelector("#AppRuntimeService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAboutServiceImpl) newObject() * implservice0x22327c.AboutServiceImpl {
	return & implservice0x22327c.AboutServiceImpl {}
}

func (inst * comFactory4pComAboutServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.AboutServiceImpl {
	return instance.Get().(*implservice0x22327c.AboutServiceImpl)
}

func (inst * comFactory4pComAboutServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAboutServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAboutServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAboutServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAboutServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAboutServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Profile = inst.getterForFieldProfileSelector(context)
	obj.Name = inst.getterForFieldNameSelector(context)
	obj.Title = inst.getterForFieldTitleSelector(context)
	obj.Copyright = inst.getterForFieldCopyrightSelector(context)
	obj.ServerPort = inst.getterForFieldServerPortSelector(context)
	obj.EnableDebug = inst.getterForFieldEnableDebugSelector(context)
	obj.PlatformService = inst.getterForFieldPlatformServiceSelector(context)
	obj.ProfileService = inst.getterForFieldProfileServiceSelector(context)
	obj.AppRuntimeService = inst.getterForFieldAppRuntimeServiceSelector(context)
	return context.LastError()
}

//getterForFieldProfileSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldProfileSelector (context application.InstanceContext) string {
    return inst.mProfileSelector.GetString(context)
}

//getterForFieldNameSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldNameSelector (context application.InstanceContext) string {
    return inst.mNameSelector.GetString(context)
}

//getterForFieldTitleSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldTitleSelector (context application.InstanceContext) string {
    return inst.mTitleSelector.GetString(context)
}

//getterForFieldCopyrightSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldCopyrightSelector (context application.InstanceContext) string {
    return inst.mCopyrightSelector.GetString(context)
}

//getterForFieldServerPortSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldServerPortSelector (context application.InstanceContext) int {
    return inst.mServerPortSelector.GetInt(context)
}

//getterForFieldEnableDebugSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldEnableDebugSelector (context application.InstanceContext) bool {
    return inst.mEnableDebugSelector.GetBool(context)
}

//getterForFieldPlatformServiceSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldPlatformServiceSelector (context application.InstanceContext) service0x3e063d.PlatformService {

	o1 := inst.mPlatformServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.PlatformService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "AboutService")
		eb.Set("field", "PlatformService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.PlatformService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProfileServiceSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldProfileServiceSelector (context application.InstanceContext) service0x3e063d.ProfileService {

	o1 := inst.mProfileServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProfileService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "AboutService")
		eb.Set("field", "ProfileService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProfileService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldAppRuntimeServiceSelector
func (inst * comFactory4pComAboutServiceImpl) getterForFieldAppRuntimeServiceSelector (context application.InstanceContext) service0x3e063d.AppRuntimeService {

	o1 := inst.mAppRuntimeServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AppRuntimeService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "AboutService")
		eb.Set("field", "AppRuntimeService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AppRuntimeService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAppDataServiceImpl : the factory of component: AppDataService
type comFactory4pComAppDataServiceImpl struct {

    mPrototype * implservice0x22327c.AppDataServiceImpl

	
	mProfileServiceSelector config.InjectionSelector
	mAppRuntimeServiceSelector config.InjectionSelector
	mSQLiteDatabaseNameWithAppVersionSelector config.InjectionSelector

}

func (inst * comFactory4pComAppDataServiceImpl) init() application.ComponentFactory {

	
	inst.mProfileServiceSelector = config.NewInjectionSelector("#ProfileService",nil)
	inst.mAppRuntimeServiceSelector = config.NewInjectionSelector("#AppRuntimeService",nil)
	inst.mSQLiteDatabaseNameWithAppVersionSelector = config.NewInjectionSelector("${datasource.wpm.database-name-with-version}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAppDataServiceImpl) newObject() * implservice0x22327c.AppDataServiceImpl {
	return & implservice0x22327c.AppDataServiceImpl {}
}

func (inst * comFactory4pComAppDataServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.AppDataServiceImpl {
	return instance.Get().(*implservice0x22327c.AppDataServiceImpl)
}

func (inst * comFactory4pComAppDataServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAppDataServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAppDataServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAppDataServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAppDataServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAppDataServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ProfileService = inst.getterForFieldProfileServiceSelector(context)
	obj.AppRuntimeService = inst.getterForFieldAppRuntimeServiceSelector(context)
	obj.SQLiteDatabaseNameWithAppVersion = inst.getterForFieldSQLiteDatabaseNameWithAppVersionSelector(context)
	return context.LastError()
}

//getterForFieldProfileServiceSelector
func (inst * comFactory4pComAppDataServiceImpl) getterForFieldProfileServiceSelector (context application.InstanceContext) service0x3e063d.ProfileService {

	o1 := inst.mProfileServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProfileService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "AppDataService")
		eb.Set("field", "ProfileService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProfileService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldAppRuntimeServiceSelector
func (inst * comFactory4pComAppDataServiceImpl) getterForFieldAppRuntimeServiceSelector (context application.InstanceContext) service0x3e063d.AppRuntimeService {

	o1 := inst.mAppRuntimeServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AppRuntimeService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "AppDataService")
		eb.Set("field", "AppRuntimeService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AppRuntimeService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSQLiteDatabaseNameWithAppVersionSelector
func (inst * comFactory4pComAppDataServiceImpl) getterForFieldSQLiteDatabaseNameWithAppVersionSelector (context application.InstanceContext) bool {
    return inst.mSQLiteDatabaseNameWithAppVersionSelector.GetBool(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExampleServiceImpl : the factory of component: ExampleService
type comFactory4pComExampleServiceImpl struct {

    mPrototype * implservice0x22327c.ExampleServiceImpl

	

}

func (inst * comFactory4pComExampleServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExampleServiceImpl) newObject() * implservice0x22327c.ExampleServiceImpl {
	return & implservice0x22327c.ExampleServiceImpl {}
}

func (inst * comFactory4pComExampleServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.ExampleServiceImpl {
	return instance.Get().(*implservice0x22327c.ExampleServiceImpl)
}

func (inst * comFactory4pComExampleServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExampleServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExampleServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExampleServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFileQueryServiceImpl : the factory of component: FileQueryService
type comFactory4pComFileQueryServiceImpl struct {

    mPrototype * implservice0x22327c.FileQueryServiceImpl

	
	mHandlerRegistryListSelector config.InjectionSelector

}

func (inst * comFactory4pComFileQueryServiceImpl) init() application.ComponentFactory {

	
	inst.mHandlerRegistryListSelector = config.NewInjectionSelector(".filequery-handler-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFileQueryServiceImpl) newObject() * implservice0x22327c.FileQueryServiceImpl {
	return & implservice0x22327c.FileQueryServiceImpl {}
}

func (inst * comFactory4pComFileQueryServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.FileQueryServiceImpl {
	return instance.Get().(*implservice0x22327c.FileQueryServiceImpl)
}

func (inst * comFactory4pComFileQueryServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFileQueryServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFileQueryServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFileQueryServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileQueryServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileQueryServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.HandlerRegistryList = inst.getterForFieldHandlerRegistryListSelector(context)
	return context.LastError()
}

//getterForFieldHandlerRegistryListSelector
func (inst * comFactory4pComFileQueryServiceImpl) getterForFieldHandlerRegistryListSelector (context application.InstanceContext) []filequery0xca51d2.HandlerRegistry {
	list1 := inst.mHandlerRegistryListSelector.GetList(context)
	list2 := make([]filequery0xca51d2.HandlerRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(filequery0xca51d2.HandlerRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFileSystemServiceImpl : the factory of component: FileSystemService
type comFactory4pComFileSystemServiceImpl struct {

    mPrototype * implservice0x22327c.FileSystemServiceImpl

	

}

func (inst * comFactory4pComFileSystemServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFileSystemServiceImpl) newObject() * implservice0x22327c.FileSystemServiceImpl {
	return & implservice0x22327c.FileSystemServiceImpl {}
}

func (inst * comFactory4pComFileSystemServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.FileSystemServiceImpl {
	return instance.Get().(*implservice0x22327c.FileSystemServiceImpl)
}

func (inst * comFactory4pComFileSystemServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFileSystemServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFileSystemServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFileSystemServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileSystemServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileSystemServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUUIDGenServiceImpl : the factory of component: UUIDGenService
type comFactory4pComUUIDGenServiceImpl struct {

    mPrototype * implservice0x22327c.UUIDGenServiceImpl

	

}

func (inst * comFactory4pComUUIDGenServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUUIDGenServiceImpl) newObject() * implservice0x22327c.UUIDGenServiceImpl {
	return & implservice0x22327c.UUIDGenServiceImpl {}
}

func (inst * comFactory4pComUUIDGenServiceImpl) castObject(instance application.ComponentInstance) * implservice0x22327c.UUIDGenServiceImpl {
	return instance.Get().(*implservice0x22327c.UUIDGenServiceImpl)
}

func (inst * comFactory4pComUUIDGenServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComUUIDGenServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComUUIDGenServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComUUIDGenServiceImpl) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComUUIDGenServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUUIDGenServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComWpmInitController : the factory of component: com32-init0xc984bc.WpmInitController
type comFactory4pComWpmInitController struct {

    mPrototype * init0xc984bc.WpmInitController

	
	mInitServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComWpmInitController) init() application.ComponentFactory {

	
	inst.mInitServiceSelector = config.NewInjectionSelector("#InitService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComWpmInitController) newObject() * init0xc984bc.WpmInitController {
	return & init0xc984bc.WpmInitController {}
}

func (inst * comFactory4pComWpmInitController) castObject(instance application.ComponentInstance) * init0xc984bc.WpmInitController {
	return instance.Get().(*init0xc984bc.WpmInitController)
}

func (inst * comFactory4pComWpmInitController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComWpmInitController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComWpmInitController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComWpmInitController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWpmInitController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWpmInitController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.InitService = inst.getterForFieldInitServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldInitServiceSelector
func (inst * comFactory4pComWpmInitController) getterForFieldInitServiceSelector (context application.InstanceContext) service0x3e063d.InitService {

	o1 := inst.mInitServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.InitService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com32-init0xc984bc.WpmInitController")
		eb.Set("field", "InitService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.InitService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComWpmInitController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com32-init0xc984bc.WpmInitController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpInitService : the factory of component: InitService
type comFactory4pComImpInitService struct {

    mPrototype * init0xc984bc.ImpInitService

	
	mAboutServiceSelector config.InjectionSelector
	mProjectTypeServiceSelector config.InjectionSelector
	mExecutableServiceSelector config.InjectionSelector
	mCheckUpdateServiceSelector config.InjectionSelector
	mSetupServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComImpInitService) init() application.ComponentFactory {

	
	inst.mAboutServiceSelector = config.NewInjectionSelector("#AboutService",nil)
	inst.mProjectTypeServiceSelector = config.NewInjectionSelector("#ContentTypeService",nil)
	inst.mExecutableServiceSelector = config.NewInjectionSelector("#ExecutableService",nil)
	inst.mCheckUpdateServiceSelector = config.NewInjectionSelector("#CheckUpdateService",nil)
	inst.mSetupServiceSelector = config.NewInjectionSelector("#SetupService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpInitService) newObject() * init0xc984bc.ImpInitService {
	return & init0xc984bc.ImpInitService {}
}

func (inst * comFactory4pComImpInitService) castObject(instance application.ComponentInstance) * init0xc984bc.ImpInitService {
	return instance.Get().(*init0xc984bc.ImpInitService)
}

func (inst * comFactory4pComImpInitService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpInitService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpInitService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpInitService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpInitService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpInitService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AboutService = inst.getterForFieldAboutServiceSelector(context)
	obj.ProjectTypeService = inst.getterForFieldProjectTypeServiceSelector(context)
	obj.ExecutableService = inst.getterForFieldExecutableServiceSelector(context)
	obj.CheckUpdateService = inst.getterForFieldCheckUpdateServiceSelector(context)
	obj.SetupService = inst.getterForFieldSetupServiceSelector(context)
	return context.LastError()
}

//getterForFieldAboutServiceSelector
func (inst * comFactory4pComImpInitService) getterForFieldAboutServiceSelector (context application.InstanceContext) service0x3e063d.AboutService {

	o1 := inst.mAboutServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AboutService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "InitService")
		eb.Set("field", "AboutService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AboutService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProjectTypeServiceSelector
func (inst * comFactory4pComImpInitService) getterForFieldProjectTypeServiceSelector (context application.InstanceContext) service0x3e063d.ContentTypeService {

	o1 := inst.mProjectTypeServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ContentTypeService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "InitService")
		eb.Set("field", "ProjectTypeService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ContentTypeService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldExecutableServiceSelector
func (inst * comFactory4pComImpInitService) getterForFieldExecutableServiceSelector (context application.InstanceContext) service0x3e063d.ExecutableService {

	o1 := inst.mExecutableServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ExecutableService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "InitService")
		eb.Set("field", "ExecutableService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ExecutableService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldCheckUpdateServiceSelector
func (inst * comFactory4pComImpInitService) getterForFieldCheckUpdateServiceSelector (context application.InstanceContext) service0x3e063d.CheckUpdateService {

	o1 := inst.mCheckUpdateServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.CheckUpdateService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "InitService")
		eb.Set("field", "CheckUpdateService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.CheckUpdateService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSetupServiceSelector
func (inst * comFactory4pComImpInitService) getterForFieldSetupServiceSelector (context application.InstanceContext) service0x3e063d.SetupService {

	o1 := inst.mSetupServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.SetupService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "InitService")
		eb.Set("field", "SetupService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.SetupService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRunIntentController : the factory of component: com34-intents0x8ee0e0.RunIntentController
type comFactory4pComRunIntentController struct {

    mPrototype * intents0x8ee0e0.RunIntentController

	
	mIntentServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComRunIntentController) init() application.ComponentFactory {

	
	inst.mIntentServiceSelector = config.NewInjectionSelector("#IntentService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRunIntentController) newObject() * intents0x8ee0e0.RunIntentController {
	return & intents0x8ee0e0.RunIntentController {}
}

func (inst * comFactory4pComRunIntentController) castObject(instance application.ComponentInstance) * intents0x8ee0e0.RunIntentController {
	return instance.Get().(*intents0x8ee0e0.RunIntentController)
}

func (inst * comFactory4pComRunIntentController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRunIntentController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRunIntentController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRunIntentController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunIntentController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunIntentController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.IntentService = inst.getterForFieldIntentServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldIntentServiceSelector
func (inst * comFactory4pComRunIntentController) getterForFieldIntentServiceSelector (context application.InstanceContext) service0x3e063d.IntentService {

	o1 := inst.mIntentServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.IntentService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com34-intents0x8ee0e0.RunIntentController")
		eb.Set("field", "IntentService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.IntentService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComRunIntentController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com34-intents0x8ee0e0.RunIntentController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFilterManagerImpl : the factory of component: wpm-intent-filter-manager
type comFactory4pComFilterManagerImpl struct {

    mPrototype * intents0x8ee0e0.FilterManagerImpl

	
	mFilterRegistryListSelector config.InjectionSelector

}

func (inst * comFactory4pComFilterManagerImpl) init() application.ComponentFactory {

	
	inst.mFilterRegistryListSelector = config.NewInjectionSelector(".wpm-intent-filter",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFilterManagerImpl) newObject() * intents0x8ee0e0.FilterManagerImpl {
	return & intents0x8ee0e0.FilterManagerImpl {}
}

func (inst * comFactory4pComFilterManagerImpl) castObject(instance application.ComponentInstance) * intents0x8ee0e0.FilterManagerImpl {
	return instance.Get().(*intents0x8ee0e0.FilterManagerImpl)
}

func (inst * comFactory4pComFilterManagerImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFilterManagerImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFilterManagerImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFilterManagerImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFilterManagerImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFilterManagerImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.FilterRegistryList = inst.getterForFieldFilterRegistryListSelector(context)
	return context.LastError()
}

//getterForFieldFilterRegistryListSelector
func (inst * comFactory4pComFilterManagerImpl) getterForFieldFilterRegistryListSelector (context application.InstanceContext) []intents0x8557f3.FilterRegistry {
	list1 := inst.mFilterRegistryListSelector.GetList(context)
	list2 := make([]intents0x8557f3.FilterRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(intents0x8557f3.FilterRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentHandlerImpl : the factory of component: IntentHandlerService
type comFactory4pComIntentHandlerImpl struct {

    mPrototype * intents0x8ee0e0.IntentHandlerImpl

	

}

func (inst * comFactory4pComIntentHandlerImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentHandlerImpl) newObject() * intents0x8ee0e0.IntentHandlerImpl {
	return & intents0x8ee0e0.IntentHandlerImpl {}
}

func (inst * comFactory4pComIntentHandlerImpl) castObject(instance application.ComponentInstance) * intents0x8ee0e0.IntentHandlerImpl {
	return instance.Get().(*intents0x8ee0e0.IntentHandlerImpl)
}

func (inst * comFactory4pComIntentHandlerImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentHandlerImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentHandlerImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentHandlerImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentHandlerImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentHandlerImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRunIntentServiceImpl : the factory of component: IntentService
type comFactory4pComRunIntentServiceImpl struct {

    mPrototype * intents0x8ee0e0.RunIntentServiceImpl

	
	mGitLibAgentSelector config.InjectionSelector
	mIntentFilterManagerSelector config.InjectionSelector
	mLocalRepositoryServiceSelector config.InjectionSelector
	mExecutableServiceSelector config.InjectionSelector
	mIntentHandlerServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComRunIntentServiceImpl) init() application.ComponentFactory {

	
	inst.mGitLibAgentSelector = config.NewInjectionSelector("#git-lib-agent",nil)
	inst.mIntentFilterManagerSelector = config.NewInjectionSelector("#wpm-intent-filter-manager",nil)
	inst.mLocalRepositoryServiceSelector = config.NewInjectionSelector("#LocalRepositoryService",nil)
	inst.mExecutableServiceSelector = config.NewInjectionSelector("#ExecutableService",nil)
	inst.mIntentHandlerServiceSelector = config.NewInjectionSelector("#IntentHandlerService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRunIntentServiceImpl) newObject() * intents0x8ee0e0.RunIntentServiceImpl {
	return & intents0x8ee0e0.RunIntentServiceImpl {}
}

func (inst * comFactory4pComRunIntentServiceImpl) castObject(instance application.ComponentInstance) * intents0x8ee0e0.RunIntentServiceImpl {
	return instance.Get().(*intents0x8ee0e0.RunIntentServiceImpl)
}

func (inst * comFactory4pComRunIntentServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRunIntentServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRunIntentServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRunIntentServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunIntentServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunIntentServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.GitLibAgent = inst.getterForFieldGitLibAgentSelector(context)
	obj.IntentFilterManager = inst.getterForFieldIntentFilterManagerSelector(context)
	obj.LocalRepositoryService = inst.getterForFieldLocalRepositoryServiceSelector(context)
	obj.ExecutableService = inst.getterForFieldExecutableServiceSelector(context)
	obj.IntentHandlerService = inst.getterForFieldIntentHandlerServiceSelector(context)
	return context.LastError()
}

//getterForFieldGitLibAgentSelector
func (inst * comFactory4pComRunIntentServiceImpl) getterForFieldGitLibAgentSelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mGitLibAgentSelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentService")
		eb.Set("field", "GitLibAgent")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldIntentFilterManagerSelector
func (inst * comFactory4pComRunIntentServiceImpl) getterForFieldIntentFilterManagerSelector (context application.InstanceContext) intents0x8557f3.FilterManager {

	o1 := inst.mIntentFilterManagerSelector.GetOne(context)
	o2, ok := o1.(intents0x8557f3.FilterManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentService")
		eb.Set("field", "IntentFilterManager")
		eb.Set("type1", "?")
		eb.Set("type2", "intents0x8557f3.FilterManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLocalRepositoryServiceSelector
func (inst * comFactory4pComRunIntentServiceImpl) getterForFieldLocalRepositoryServiceSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryService {

	o1 := inst.mLocalRepositoryServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentService")
		eb.Set("field", "LocalRepositoryService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldExecutableServiceSelector
func (inst * comFactory4pComRunIntentServiceImpl) getterForFieldExecutableServiceSelector (context application.InstanceContext) service0x3e063d.ExecutableService {

	o1 := inst.mExecutableServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ExecutableService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentService")
		eb.Set("field", "ExecutableService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ExecutableService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldIntentHandlerServiceSelector
func (inst * comFactory4pComRunIntentServiceImpl) getterForFieldIntentHandlerServiceSelector (context application.InstanceContext) service0x3e063d.IntentHandlerService {

	o1 := inst.mIntentHandlerServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.IntentHandlerService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentService")
		eb.Set("field", "IntentHandlerService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.IntentHandlerService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentTemplateController : the factory of component: com38-intenttemplates0x2e3dcf.IntentTemplateController
type comFactory4pComIntentTemplateController struct {

    mPrototype * intenttemplates0x2e3dcf.IntentTemplateController

	
	mIntentTemplateServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComIntentTemplateController) init() application.ComponentFactory {

	
	inst.mIntentTemplateServiceSelector = config.NewInjectionSelector("#IntentTemplateService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentTemplateController) newObject() * intenttemplates0x2e3dcf.IntentTemplateController {
	return & intenttemplates0x2e3dcf.IntentTemplateController {}
}

func (inst * comFactory4pComIntentTemplateController) castObject(instance application.ComponentInstance) * intenttemplates0x2e3dcf.IntentTemplateController {
	return instance.Get().(*intenttemplates0x2e3dcf.IntentTemplateController)
}

func (inst * comFactory4pComIntentTemplateController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentTemplateController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentTemplateController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentTemplateController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentTemplateController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentTemplateController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.IntentTemplateService = inst.getterForFieldIntentTemplateServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldIntentTemplateServiceSelector
func (inst * comFactory4pComIntentTemplateController) getterForFieldIntentTemplateServiceSelector (context application.InstanceContext) service0x3e063d.IntentTemplateService {

	o1 := inst.mIntentTemplateServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.IntentTemplateService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com38-intenttemplates0x2e3dcf.IntentTemplateController")
		eb.Set("field", "IntentTemplateService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.IntentTemplateService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComIntentTemplateController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com38-intenttemplates0x2e3dcf.IntentTemplateController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentTemplateDaoImpl : the factory of component: IntentTemplateDAO
type comFactory4pComIntentTemplateDaoImpl struct {

    mPrototype * intenttemplates0x2e3dcf.IntentTemplateDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComIntentTemplateDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentTemplateDaoImpl) newObject() * intenttemplates0x2e3dcf.IntentTemplateDaoImpl {
	return & intenttemplates0x2e3dcf.IntentTemplateDaoImpl {}
}

func (inst * comFactory4pComIntentTemplateDaoImpl) castObject(instance application.ComponentInstance) * intenttemplates0x2e3dcf.IntentTemplateDaoImpl {
	return instance.Get().(*intenttemplates0x2e3dcf.IntentTemplateDaoImpl)
}

func (inst * comFactory4pComIntentTemplateDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentTemplateDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentTemplateDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentTemplateDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentTemplateDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentTemplateDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComIntentTemplateDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) dbagent0x9f90fb.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(dbagent0x9f90fb.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentTemplateDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "dbagent0x9f90fb.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComIntentTemplateDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentTemplateDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComIntentTemplateServiceImpl : the factory of component: IntentTemplateService
type comFactory4pComIntentTemplateServiceImpl struct {

    mPrototype * intenttemplates0x2e3dcf.IntentTemplateServiceImpl

	
	mACSelector config.InjectionSelector
	mIntentTempDAOSelector config.InjectionSelector
	mIntentFilterManagerSelector config.InjectionSelector
	mPresetServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComIntentTemplateServiceImpl) init() application.ComponentFactory {

	
	inst.mACSelector = config.NewInjectionSelector("context",nil)
	inst.mIntentTempDAOSelector = config.NewInjectionSelector("#IntentTemplateDAO",nil)
	inst.mIntentFilterManagerSelector = config.NewInjectionSelector("#wpm-intent-filter-manager",nil)
	inst.mPresetServiceSelector = config.NewInjectionSelector("#PresetService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComIntentTemplateServiceImpl) newObject() * intenttemplates0x2e3dcf.IntentTemplateServiceImpl {
	return & intenttemplates0x2e3dcf.IntentTemplateServiceImpl {}
}

func (inst * comFactory4pComIntentTemplateServiceImpl) castObject(instance application.ComponentInstance) * intenttemplates0x2e3dcf.IntentTemplateServiceImpl {
	return instance.Get().(*intenttemplates0x2e3dcf.IntentTemplateServiceImpl)
}

func (inst * comFactory4pComIntentTemplateServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComIntentTemplateServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComIntentTemplateServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComIntentTemplateServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentTemplateServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComIntentTemplateServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AC = inst.getterForFieldACSelector(context)
	obj.IntentTempDAO = inst.getterForFieldIntentTempDAOSelector(context)
	obj.IntentFilterManager = inst.getterForFieldIntentFilterManagerSelector(context)
	obj.PresetService = inst.getterForFieldPresetServiceSelector(context)
	return context.LastError()
}

//getterForFieldACSelector
func (inst * comFactory4pComIntentTemplateServiceImpl) getterForFieldACSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldIntentTempDAOSelector
func (inst * comFactory4pComIntentTemplateServiceImpl) getterForFieldIntentTempDAOSelector (context application.InstanceContext) dao0x5af8d0.IntentTemplateDAO {

	o1 := inst.mIntentTempDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.IntentTemplateDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentTemplateService")
		eb.Set("field", "IntentTempDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.IntentTemplateDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldIntentFilterManagerSelector
func (inst * comFactory4pComIntentTemplateServiceImpl) getterForFieldIntentFilterManagerSelector (context application.InstanceContext) intents0x8557f3.FilterManager {

	o1 := inst.mIntentFilterManagerSelector.GetOne(context)
	o2, ok := o1.(intents0x8557f3.FilterManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentTemplateService")
		eb.Set("field", "IntentFilterManager")
		eb.Set("type1", "?")
		eb.Set("type2", "intents0x8557f3.FilterManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldPresetServiceSelector
func (inst * comFactory4pComIntentTemplateServiceImpl) getterForFieldPresetServiceSelector (context application.InstanceContext) service0x3e063d.PresetService {

	o1 := inst.mPresetServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.PresetService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "IntentTemplateService")
		eb.Set("field", "PresetService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.PresetService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpLocationDao : the factory of component: LocationDAO
type comFactory4pComImpLocationDao struct {

    mPrototype * locations0xb36349.ImpLocationDao

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComImpLocationDao) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpLocationDao) newObject() * locations0xb36349.ImpLocationDao {
	return & locations0xb36349.ImpLocationDao {}
}

func (inst * comFactory4pComImpLocationDao) castObject(instance application.ComponentInstance) * locations0xb36349.ImpLocationDao {
	return instance.Get().(*locations0xb36349.ImpLocationDao)
}

func (inst * comFactory4pComImpLocationDao) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpLocationDao) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpLocationDao) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpLocationDao) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpLocationDao) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpLocationDao) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComImpLocationDao) getterForFieldAgentSelector (context application.InstanceContext) dbagent0x9f90fb.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(dbagent0x9f90fb.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocationDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "dbagent0x9f90fb.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComImpLocationDao) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocationDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpLocationService : the factory of component: LocationService
type comFactory4pComImpLocationService struct {

    mPrototype * locations0xb36349.ImpLocationService

	
	mDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComImpLocationService) init() application.ComponentFactory {

	
	inst.mDAOSelector = config.NewInjectionSelector("#LocationDAO",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpLocationService) newObject() * locations0xb36349.ImpLocationService {
	return & locations0xb36349.ImpLocationService {}
}

func (inst * comFactory4pComImpLocationService) castObject(instance application.ComponentInstance) * locations0xb36349.ImpLocationService {
	return instance.Get().(*locations0xb36349.ImpLocationService)
}

func (inst * comFactory4pComImpLocationService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpLocationService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpLocationService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpLocationService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpLocationService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpLocationService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.DAO = inst.getterForFieldDAOSelector(context)
	return context.LastError()
}

//getterForFieldDAOSelector
func (inst * comFactory4pComImpLocationService) getterForFieldDAOSelector (context application.InstanceContext) dao0x5af8d0.LocationDAO {

	o1 := inst.mDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.LocationDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocationService")
		eb.Set("field", "DAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.LocationDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMediaController : the factory of component: com43-mediae0xf005e2.MediaController
type comFactory4pComMediaController struct {

    mPrototype * mediae0xf005e2.MediaController

	
	mMediaServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComMediaController) init() application.ComponentFactory {

	
	inst.mMediaServiceSelector = config.NewInjectionSelector("#MediaService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMediaController) newObject() * mediae0xf005e2.MediaController {
	return & mediae0xf005e2.MediaController {}
}

func (inst * comFactory4pComMediaController) castObject(instance application.ComponentInstance) * mediae0xf005e2.MediaController {
	return instance.Get().(*mediae0xf005e2.MediaController)
}

func (inst * comFactory4pComMediaController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMediaController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMediaController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMediaController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMediaController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMediaController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.MediaService = inst.getterForFieldMediaServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldMediaServiceSelector
func (inst * comFactory4pComMediaController) getterForFieldMediaServiceSelector (context application.InstanceContext) service0x3e063d.MediaService {

	o1 := inst.mMediaServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.MediaService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com43-mediae0xf005e2.MediaController")
		eb.Set("field", "MediaService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.MediaService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComMediaController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com43-mediae0xf005e2.MediaController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMediaDaoImpl : the factory of component: MediaDAO
type comFactory4pComMediaDaoImpl struct {

    mPrototype * mediae0xf005e2.MediaDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComMediaDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMediaDaoImpl) newObject() * mediae0xf005e2.MediaDaoImpl {
	return & mediae0xf005e2.MediaDaoImpl {}
}

func (inst * comFactory4pComMediaDaoImpl) castObject(instance application.ComponentInstance) * mediae0xf005e2.MediaDaoImpl {
	return instance.Get().(*mediae0xf005e2.MediaDaoImpl)
}

func (inst * comFactory4pComMediaDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMediaDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMediaDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMediaDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMediaDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMediaDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComMediaDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) dbagent0x9f90fb.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(dbagent0x9f90fb.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MediaDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "dbagent0x9f90fb.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComMediaDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MediaDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMediaServiceImpl : the factory of component: MediaService
type comFactory4pComMediaServiceImpl struct {

    mPrototype * mediae0xf005e2.MediaServiceImpl

	
	mACSelector config.InjectionSelector
	mMediaDAOSelector config.InjectionSelector
	mSysMainRepoServiceSelector config.InjectionSelector
	mFileSystemServiceSelector config.InjectionSelector
	mContentTypeServiceSelector config.InjectionSelector
	mResPathPrefixSelector config.InjectionSelector
	mWebPathPrefixSelector config.InjectionSelector

}

func (inst * comFactory4pComMediaServiceImpl) init() application.ComponentFactory {

	
	inst.mACSelector = config.NewInjectionSelector("context",nil)
	inst.mMediaDAOSelector = config.NewInjectionSelector("#MediaDAO",nil)
	inst.mSysMainRepoServiceSelector = config.NewInjectionSelector("#MainRepositoryService",nil)
	inst.mFileSystemServiceSelector = config.NewInjectionSelector("#FileSystemService",nil)
	inst.mContentTypeServiceSelector = config.NewInjectionSelector("#ContentTypeService",nil)
	inst.mResPathPrefixSelector = config.NewInjectionSelector("${wpm.presets.res-path-prefix}",nil)
	inst.mWebPathPrefixSelector = config.NewInjectionSelector("${wpm.presets.web-path-prefix}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMediaServiceImpl) newObject() * mediae0xf005e2.MediaServiceImpl {
	return & mediae0xf005e2.MediaServiceImpl {}
}

func (inst * comFactory4pComMediaServiceImpl) castObject(instance application.ComponentInstance) * mediae0xf005e2.MediaServiceImpl {
	return instance.Get().(*mediae0xf005e2.MediaServiceImpl)
}

func (inst * comFactory4pComMediaServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMediaServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMediaServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMediaServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMediaServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMediaServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AC = inst.getterForFieldACSelector(context)
	obj.MediaDAO = inst.getterForFieldMediaDAOSelector(context)
	obj.SysMainRepoService = inst.getterForFieldSysMainRepoServiceSelector(context)
	obj.FileSystemService = inst.getterForFieldFileSystemServiceSelector(context)
	obj.ContentTypeService = inst.getterForFieldContentTypeServiceSelector(context)
	obj.ResPathPrefix = inst.getterForFieldResPathPrefixSelector(context)
	obj.WebPathPrefix = inst.getterForFieldWebPathPrefixSelector(context)
	return context.LastError()
}

//getterForFieldACSelector
func (inst * comFactory4pComMediaServiceImpl) getterForFieldACSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldMediaDAOSelector
func (inst * comFactory4pComMediaServiceImpl) getterForFieldMediaDAOSelector (context application.InstanceContext) dao0x5af8d0.MediaDAO {

	o1 := inst.mMediaDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.MediaDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MediaService")
		eb.Set("field", "MediaDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.MediaDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSysMainRepoServiceSelector
func (inst * comFactory4pComMediaServiceImpl) getterForFieldSysMainRepoServiceSelector (context application.InstanceContext) service0x3e063d.MainRepositoryService {

	o1 := inst.mSysMainRepoServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.MainRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MediaService")
		eb.Set("field", "SysMainRepoService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.MainRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldFileSystemServiceSelector
func (inst * comFactory4pComMediaServiceImpl) getterForFieldFileSystemServiceSelector (context application.InstanceContext) service0x3e063d.FileSystemService {

	o1 := inst.mFileSystemServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.FileSystemService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MediaService")
		eb.Set("field", "FileSystemService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.FileSystemService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldContentTypeServiceSelector
func (inst * comFactory4pComMediaServiceImpl) getterForFieldContentTypeServiceSelector (context application.InstanceContext) service0x3e063d.ContentTypeService {

	o1 := inst.mContentTypeServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ContentTypeService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MediaService")
		eb.Set("field", "ContentTypeService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ContentTypeService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResPathPrefixSelector
func (inst * comFactory4pComMediaServiceImpl) getterForFieldResPathPrefixSelector (context application.InstanceContext) string {
    return inst.mResPathPrefixSelector.GetString(context)
}

//getterForFieldWebPathPrefixSelector
func (inst * comFactory4pComMediaServiceImpl) getterForFieldWebPathPrefixSelector (context application.InstanceContext) string {
    return inst.mWebPathPrefixSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComNamespaceController : the factory of component: com46-namespaces0xceefcf.NamespaceController
type comFactory4pComNamespaceController struct {

    mPrototype * namespaces0xceefcf.NamespaceController

	
	mNamespaceServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComNamespaceController) init() application.ComponentFactory {

	
	inst.mNamespaceServiceSelector = config.NewInjectionSelector("#NamespaceService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComNamespaceController) newObject() * namespaces0xceefcf.NamespaceController {
	return & namespaces0xceefcf.NamespaceController {}
}

func (inst * comFactory4pComNamespaceController) castObject(instance application.ComponentInstance) * namespaces0xceefcf.NamespaceController {
	return instance.Get().(*namespaces0xceefcf.NamespaceController)
}

func (inst * comFactory4pComNamespaceController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComNamespaceController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComNamespaceController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComNamespaceController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComNamespaceController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComNamespaceController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.NamespaceService = inst.getterForFieldNamespaceServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldNamespaceServiceSelector
func (inst * comFactory4pComNamespaceController) getterForFieldNamespaceServiceSelector (context application.InstanceContext) service0x3e063d.NamespaceService {

	o1 := inst.mNamespaceServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.NamespaceService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com46-namespaces0xceefcf.NamespaceController")
		eb.Set("field", "NamespaceService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.NamespaceService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComNamespaceController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com46-namespaces0xceefcf.NamespaceController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpNamespaceDao : the factory of component: NamespaceDAO
type comFactory4pComImpNamespaceDao struct {

    mPrototype * namespaces0xceefcf.ImpNamespaceDao

	
	mACSelector config.InjectionSelector

}

func (inst * comFactory4pComImpNamespaceDao) init() application.ComponentFactory {

	
	inst.mACSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpNamespaceDao) newObject() * namespaces0xceefcf.ImpNamespaceDao {
	return & namespaces0xceefcf.ImpNamespaceDao {}
}

func (inst * comFactory4pComImpNamespaceDao) castObject(instance application.ComponentInstance) * namespaces0xceefcf.ImpNamespaceDao {
	return instance.Get().(*namespaces0xceefcf.ImpNamespaceDao)
}

func (inst * comFactory4pComImpNamespaceDao) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpNamespaceDao) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpNamespaceDao) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpNamespaceDao) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpNamespaceDao) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpNamespaceDao) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AC = inst.getterForFieldACSelector(context)
	return context.LastError()
}

//getterForFieldACSelector
func (inst * comFactory4pComImpNamespaceDao) getterForFieldACSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpNamespaceService : the factory of component: NamespaceService
type comFactory4pComImpNamespaceService struct {

    mPrototype * namespaces0xceefcf.ImpNamespaceService

	
	mMyDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComImpNamespaceService) init() application.ComponentFactory {

	
	inst.mMyDAOSelector = config.NewInjectionSelector("#NamespaceDAO",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpNamespaceService) newObject() * namespaces0xceefcf.ImpNamespaceService {
	return & namespaces0xceefcf.ImpNamespaceService {}
}

func (inst * comFactory4pComImpNamespaceService) castObject(instance application.ComponentInstance) * namespaces0xceefcf.ImpNamespaceService {
	return instance.Get().(*namespaces0xceefcf.ImpNamespaceService)
}

func (inst * comFactory4pComImpNamespaceService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpNamespaceService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpNamespaceService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpNamespaceService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpNamespaceService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpNamespaceService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.MyDAO = inst.getterForFieldMyDAOSelector(context)
	return context.LastError()
}

//getterForFieldMyDAOSelector
func (inst * comFactory4pComImpNamespaceService) getterForFieldMyDAOSelector (context application.InstanceContext) dao0x5af8d0.NamespaceDAO {

	o1 := inst.mMyDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.NamespaceDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "NamespaceService")
		eb.Set("field", "MyDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.NamespaceDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLinuxPlatformServiceImpl : the factory of component: com49-platforms0xb539c0.LinuxPlatformServiceImpl
type comFactory4pComLinuxPlatformServiceImpl struct {

    mPrototype * platforms0xb539c0.LinuxPlatformServiceImpl

	

}

func (inst * comFactory4pComLinuxPlatformServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) newObject() * platforms0xb539c0.LinuxPlatformServiceImpl {
	return & platforms0xb539c0.LinuxPlatformServiceImpl {}
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) castObject(instance application.ComponentInstance) * platforms0xb539c0.LinuxPlatformServiceImpl {
	return instance.Get().(*platforms0xb539c0.LinuxPlatformServiceImpl)
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLinuxPlatformServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPlatformServiceImpl : the factory of component: PlatformService
type comFactory4pComPlatformServiceImpl struct {

    mPrototype * platforms0xb539c0.PlatformServiceImpl

	
	mProvidersSelector config.InjectionSelector

}

func (inst * comFactory4pComPlatformServiceImpl) init() application.ComponentFactory {

	
	inst.mProvidersSelector = config.NewInjectionSelector(".PlatformProviderRegistry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPlatformServiceImpl) newObject() * platforms0xb539c0.PlatformServiceImpl {
	return & platforms0xb539c0.PlatformServiceImpl {}
}

func (inst * comFactory4pComPlatformServiceImpl) castObject(instance application.ComponentInstance) * platforms0xb539c0.PlatformServiceImpl {
	return instance.Get().(*platforms0xb539c0.PlatformServiceImpl)
}

func (inst * comFactory4pComPlatformServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPlatformServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPlatformServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPlatformServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPlatformServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPlatformServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Providers = inst.getterForFieldProvidersSelector(context)
	return context.LastError()
}

//getterForFieldProvidersSelector
func (inst * comFactory4pComPlatformServiceImpl) getterForFieldProvidersSelector (context application.InstanceContext) []service0x3e063d.PlatformProviderRegistry {
	list1 := inst.mProvidersSelector.GetList(context)
	list2 := make([]service0x3e063d.PlatformProviderRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(service0x3e063d.PlatformProviderRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProfileServiceImpl : the factory of component: ProfileService
type comFactory4pComProfileServiceImpl struct {

    mPrototype * platforms0xb539c0.ProfileServiceImpl

	
	mPlatformServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComProfileServiceImpl) init() application.ComponentFactory {

	
	inst.mPlatformServiceSelector = config.NewInjectionSelector("#PlatformService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProfileServiceImpl) newObject() * platforms0xb539c0.ProfileServiceImpl {
	return & platforms0xb539c0.ProfileServiceImpl {}
}

func (inst * comFactory4pComProfileServiceImpl) castObject(instance application.ComponentInstance) * platforms0xb539c0.ProfileServiceImpl {
	return instance.Get().(*platforms0xb539c0.ProfileServiceImpl)
}

func (inst * comFactory4pComProfileServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProfileServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProfileServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProfileServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProfileServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProfileServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.PlatformService = inst.getterForFieldPlatformServiceSelector(context)
	return context.LastError()
}

//getterForFieldPlatformServiceSelector
func (inst * comFactory4pComProfileServiceImpl) getterForFieldPlatformServiceSelector (context application.InstanceContext) service0x3e063d.PlatformService {

	o1 := inst.mPlatformServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.PlatformService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProfileService")
		eb.Set("field", "PlatformService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.PlatformService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComWindowsPlatformServiceImpl : the factory of component: com52-platforms0xb539c0.WindowsPlatformServiceImpl
type comFactory4pComWindowsPlatformServiceImpl struct {

    mPrototype * platforms0xb539c0.WindowsPlatformServiceImpl

	

}

func (inst * comFactory4pComWindowsPlatformServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) newObject() * platforms0xb539c0.WindowsPlatformServiceImpl {
	return & platforms0xb539c0.WindowsPlatformServiceImpl {}
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) castObject(instance application.ComponentInstance) * platforms0xb539c0.WindowsPlatformServiceImpl {
	return instance.Get().(*platforms0xb539c0.WindowsPlatformServiceImpl)
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWindowsPlatformServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComSoftwarePackageController : the factory of component: com53-plugins0x82e34b.SoftwarePackageController
type comFactory4pComSoftwarePackageController struct {

    mPrototype * plugins0x82e34b.SoftwarePackageController

	
	mSoftwarePackageServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComSoftwarePackageController) init() application.ComponentFactory {

	
	inst.mSoftwarePackageServiceSelector = config.NewInjectionSelector("#SoftwarePackageService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComSoftwarePackageController) newObject() * plugins0x82e34b.SoftwarePackageController {
	return & plugins0x82e34b.SoftwarePackageController {}
}

func (inst * comFactory4pComSoftwarePackageController) castObject(instance application.ComponentInstance) * plugins0x82e34b.SoftwarePackageController {
	return instance.Get().(*plugins0x82e34b.SoftwarePackageController)
}

func (inst * comFactory4pComSoftwarePackageController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComSoftwarePackageController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComSoftwarePackageController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComSoftwarePackageController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSoftwarePackageController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSoftwarePackageController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.SoftwarePackageService = inst.getterForFieldSoftwarePackageServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldSoftwarePackageServiceSelector
func (inst * comFactory4pComSoftwarePackageController) getterForFieldSoftwarePackageServiceSelector (context application.InstanceContext) service0x3e063d.SoftwarePackageService {

	o1 := inst.mSoftwarePackageServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.SoftwarePackageService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com53-plugins0x82e34b.SoftwarePackageController")
		eb.Set("field", "SoftwarePackageService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.SoftwarePackageService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComSoftwarePackageController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com53-plugins0x82e34b.SoftwarePackageController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPluginDaoImpl : the factory of component: SoftwarePackageDAO
type comFactory4pComPluginDaoImpl struct {

    mPrototype * plugins0x82e34b.PluginDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComPluginDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPluginDaoImpl) newObject() * plugins0x82e34b.PluginDaoImpl {
	return & plugins0x82e34b.PluginDaoImpl {}
}

func (inst * comFactory4pComPluginDaoImpl) castObject(instance application.ComponentInstance) * plugins0x82e34b.PluginDaoImpl {
	return instance.Get().(*plugins0x82e34b.PluginDaoImpl)
}

func (inst * comFactory4pComPluginDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPluginDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPluginDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPluginDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPluginDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPluginDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComPluginDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) dbagent0x9f90fb.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(dbagent0x9f90fb.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SoftwarePackageDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "dbagent0x9f90fb.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComPluginDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SoftwarePackageDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPluginServiceImpl : the factory of component: SoftwarePackageService
type comFactory4pComPluginServiceImpl struct {

    mPrototype * plugins0x82e34b.PluginServiceImpl

	
	mSoftwarePackageDAOSelector config.InjectionSelector
	mNamespaceServiceSelector config.InjectionSelector
	mHTTPClientServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComPluginServiceImpl) init() application.ComponentFactory {

	
	inst.mSoftwarePackageDAOSelector = config.NewInjectionSelector("#SoftwarePackageDAO",nil)
	inst.mNamespaceServiceSelector = config.NewInjectionSelector("#NamespaceService",nil)
	inst.mHTTPClientServiceSelector = config.NewInjectionSelector("#HTTPClientService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPluginServiceImpl) newObject() * plugins0x82e34b.PluginServiceImpl {
	return & plugins0x82e34b.PluginServiceImpl {}
}

func (inst * comFactory4pComPluginServiceImpl) castObject(instance application.ComponentInstance) * plugins0x82e34b.PluginServiceImpl {
	return instance.Get().(*plugins0x82e34b.PluginServiceImpl)
}

func (inst * comFactory4pComPluginServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPluginServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPluginServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPluginServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPluginServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPluginServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.SoftwarePackageDAO = inst.getterForFieldSoftwarePackageDAOSelector(context)
	obj.NamespaceService = inst.getterForFieldNamespaceServiceSelector(context)
	obj.HTTPClientService = inst.getterForFieldHTTPClientServiceSelector(context)
	return context.LastError()
}

//getterForFieldSoftwarePackageDAOSelector
func (inst * comFactory4pComPluginServiceImpl) getterForFieldSoftwarePackageDAOSelector (context application.InstanceContext) dao0x5af8d0.SoftwarePackageDAO {

	o1 := inst.mSoftwarePackageDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.SoftwarePackageDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SoftwarePackageService")
		eb.Set("field", "SoftwarePackageDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.SoftwarePackageDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldNamespaceServiceSelector
func (inst * comFactory4pComPluginServiceImpl) getterForFieldNamespaceServiceSelector (context application.InstanceContext) service0x3e063d.NamespaceService {

	o1 := inst.mNamespaceServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.NamespaceService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SoftwarePackageService")
		eb.Set("field", "NamespaceService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.NamespaceService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldHTTPClientServiceSelector
func (inst * comFactory4pComPluginServiceImpl) getterForFieldHTTPClientServiceSelector (context application.InstanceContext) service0x3e063d.HTTPClientService {

	o1 := inst.mHTTPClientServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.HTTPClientService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SoftwarePackageService")
		eb.Set("field", "HTTPClientService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.HTTPClientService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComSoftwareSetController : the factory of component: com56-plugins0x82e34b.SoftwareSetController
type comFactory4pComSoftwareSetController struct {

    mPrototype * plugins0x82e34b.SoftwareSetController

	
	mSoftwareSetServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComSoftwareSetController) init() application.ComponentFactory {

	
	inst.mSoftwareSetServiceSelector = config.NewInjectionSelector("#SoftwareSetService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComSoftwareSetController) newObject() * plugins0x82e34b.SoftwareSetController {
	return & plugins0x82e34b.SoftwareSetController {}
}

func (inst * comFactory4pComSoftwareSetController) castObject(instance application.ComponentInstance) * plugins0x82e34b.SoftwareSetController {
	return instance.Get().(*plugins0x82e34b.SoftwareSetController)
}

func (inst * comFactory4pComSoftwareSetController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComSoftwareSetController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComSoftwareSetController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComSoftwareSetController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSoftwareSetController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSoftwareSetController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.SoftwareSetService = inst.getterForFieldSoftwareSetServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldSoftwareSetServiceSelector
func (inst * comFactory4pComSoftwareSetController) getterForFieldSoftwareSetServiceSelector (context application.InstanceContext) service0x3e063d.SoftwareSetService {

	o1 := inst.mSoftwareSetServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.SoftwareSetService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com56-plugins0x82e34b.SoftwareSetController")
		eb.Set("field", "SoftwareSetService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.SoftwareSetService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComSoftwareSetController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com56-plugins0x82e34b.SoftwareSetController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpSoftwareSetService : the factory of component: SoftwareSetService
type comFactory4pComImpSoftwareSetService struct {

    mPrototype * plugins0x82e34b.ImpSoftwareSetService

	
	mSoftwarePackageServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComImpSoftwareSetService) init() application.ComponentFactory {

	
	inst.mSoftwarePackageServiceSelector = config.NewInjectionSelector("#SoftwarePackageService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpSoftwareSetService) newObject() * plugins0x82e34b.ImpSoftwareSetService {
	return & plugins0x82e34b.ImpSoftwareSetService {}
}

func (inst * comFactory4pComImpSoftwareSetService) castObject(instance application.ComponentInstance) * plugins0x82e34b.ImpSoftwareSetService {
	return instance.Get().(*plugins0x82e34b.ImpSoftwareSetService)
}

func (inst * comFactory4pComImpSoftwareSetService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpSoftwareSetService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpSoftwareSetService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpSoftwareSetService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpSoftwareSetService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpSoftwareSetService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.SoftwarePackageService = inst.getterForFieldSoftwarePackageServiceSelector(context)
	return context.LastError()
}

//getterForFieldSoftwarePackageServiceSelector
func (inst * comFactory4pComImpSoftwareSetService) getterForFieldSoftwarePackageServiceSelector (context application.InstanceContext) service0x3e063d.SoftwarePackageService {

	o1 := inst.mSoftwarePackageServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.SoftwarePackageService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SoftwareSetService")
		eb.Set("field", "SoftwarePackageService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.SoftwarePackageService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCacheProvider : the factory of component: PresetCache
type comFactory4pComCacheProvider struct {

    mPrototype * presets0x875f8b.CacheProvider

	
	mACSelector config.InjectionSelector
	mCSSelector config.InjectionSelector
	mListFileNameSelector config.InjectionSelector

}

func (inst * comFactory4pComCacheProvider) init() application.ComponentFactory {

	
	inst.mACSelector = config.NewInjectionSelector("context",nil)
	inst.mCSSelector = config.NewInjectionSelector("#CacheService",nil)
	inst.mListFileNameSelector = config.NewInjectionSelector("${wpm.presets.list-file-name}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCacheProvider) newObject() * presets0x875f8b.CacheProvider {
	return & presets0x875f8b.CacheProvider {}
}

func (inst * comFactory4pComCacheProvider) castObject(instance application.ComponentInstance) * presets0x875f8b.CacheProvider {
	return instance.Get().(*presets0x875f8b.CacheProvider)
}

func (inst * comFactory4pComCacheProvider) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCacheProvider) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCacheProvider) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCacheProvider) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCacheProvider) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCacheProvider) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AC = inst.getterForFieldACSelector(context)
	obj.CS = inst.getterForFieldCSSelector(context)
	obj.ListFileName = inst.getterForFieldListFileNameSelector(context)
	return context.LastError()
}

//getterForFieldACSelector
func (inst * comFactory4pComCacheProvider) getterForFieldACSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldCSSelector
func (inst * comFactory4pComCacheProvider) getterForFieldCSSelector (context application.InstanceContext) service0x3e063d.CacheService {

	o1 := inst.mCSSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.CacheService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "PresetCache")
		eb.Set("field", "CS")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.CacheService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldListFileNameSelector
func (inst * comFactory4pComCacheProvider) getterForFieldListFileNameSelector (context application.InstanceContext) string {
    return inst.mListFileNameSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpPresetService : the factory of component: PresetService
type comFactory4pComImpPresetService struct {

    mPrototype * presets0x875f8b.ImpPresetService

	
	mCacheSelector config.InjectionSelector

}

func (inst * comFactory4pComImpPresetService) init() application.ComponentFactory {

	
	inst.mCacheSelector = config.NewInjectionSelector("#PresetCache",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpPresetService) newObject() * presets0x875f8b.ImpPresetService {
	return & presets0x875f8b.ImpPresetService {}
}

func (inst * comFactory4pComImpPresetService) castObject(instance application.ComponentInstance) * presets0x875f8b.ImpPresetService {
	return instance.Get().(*presets0x875f8b.ImpPresetService)
}

func (inst * comFactory4pComImpPresetService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpPresetService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpPresetService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpPresetService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpPresetService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpPresetService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Cache = inst.getterForFieldCacheSelector(context)
	return context.LastError()
}

//getterForFieldCacheSelector
func (inst * comFactory4pComImpPresetService) getterForFieldCacheSelector (context application.InstanceContext) presets0x875f8b.Cache {

	o1 := inst.mCacheSelector.GetOne(context)
	o2, ok := o1.(presets0x875f8b.Cache)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "PresetService")
		eb.Set("field", "Cache")
		eb.Set("type1", "?")
		eb.Set("type2", "presets0x875f8b.Cache")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectController : the factory of component: com60-projects0x4d85c7.ProjectController
type comFactory4pComProjectController struct {

    mPrototype * projects0x4d85c7.ProjectController

	
	mProjectServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectController) init() application.ComponentFactory {

	
	inst.mProjectServiceSelector = config.NewInjectionSelector("#ProjectService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectController) newObject() * projects0x4d85c7.ProjectController {
	return & projects0x4d85c7.ProjectController {}
}

func (inst * comFactory4pComProjectController) castObject(instance application.ComponentInstance) * projects0x4d85c7.ProjectController {
	return instance.Get().(*projects0x4d85c7.ProjectController)
}

func (inst * comFactory4pComProjectController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ProjectService = inst.getterForFieldProjectServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldProjectServiceSelector
func (inst * comFactory4pComProjectController) getterForFieldProjectServiceSelector (context application.InstanceContext) service0x3e063d.ProjectService {

	o1 := inst.mProjectServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProjectService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com60-projects0x4d85c7.ProjectController")
		eb.Set("field", "ProjectService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProjectService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComProjectController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com60-projects0x4d85c7.ProjectController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectDaoImpl : the factory of component: ProjectDAO
type comFactory4pComProjectDaoImpl struct {

    mPrototype * projects0x4d85c7.ProjectDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectDaoImpl) newObject() * projects0x4d85c7.ProjectDaoImpl {
	return & projects0x4d85c7.ProjectDaoImpl {}
}

func (inst * comFactory4pComProjectDaoImpl) castObject(instance application.ComponentInstance) * projects0x4d85c7.ProjectDaoImpl {
	return instance.Get().(*projects0x4d85c7.ProjectDaoImpl)
}

func (inst * comFactory4pComProjectDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComProjectDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) dbagent0x9f90fb.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(dbagent0x9f90fb.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "dbagent0x9f90fb.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComProjectDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectImportController : the factory of component: com62-projects0x4d85c7.ProjectImportController
type comFactory4pComProjectImportController struct {

    mPrototype * projects0x4d85c7.ProjectImportController

	
	mProjectImportServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectImportController) init() application.ComponentFactory {

	
	inst.mProjectImportServiceSelector = config.NewInjectionSelector("#ProjectImportService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectImportController) newObject() * projects0x4d85c7.ProjectImportController {
	return & projects0x4d85c7.ProjectImportController {}
}

func (inst * comFactory4pComProjectImportController) castObject(instance application.ComponentInstance) * projects0x4d85c7.ProjectImportController {
	return instance.Get().(*projects0x4d85c7.ProjectImportController)
}

func (inst * comFactory4pComProjectImportController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectImportController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectImportController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectImportController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectImportController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectImportController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ProjectImportService = inst.getterForFieldProjectImportServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldProjectImportServiceSelector
func (inst * comFactory4pComProjectImportController) getterForFieldProjectImportServiceSelector (context application.InstanceContext) service0x3e063d.ProjectImportService {

	o1 := inst.mProjectImportServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProjectImportService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com62-projects0x4d85c7.ProjectImportController")
		eb.Set("field", "ProjectImportService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProjectImportService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComProjectImportController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com62-projects0x4d85c7.ProjectImportController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectImportServiceImpl : the factory of component: ProjectImportService
type comFactory4pComProjectImportServiceImpl struct {

    mPrototype * projects0x4d85c7.ProjectImportServiceImpl

	

}

func (inst * comFactory4pComProjectImportServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectImportServiceImpl) newObject() * projects0x4d85c7.ProjectImportServiceImpl {
	return & projects0x4d85c7.ProjectImportServiceImpl {}
}

func (inst * comFactory4pComProjectImportServiceImpl) castObject(instance application.ComponentInstance) * projects0x4d85c7.ProjectImportServiceImpl {
	return instance.Get().(*projects0x4d85c7.ProjectImportServiceImpl)
}

func (inst * comFactory4pComProjectImportServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectImportServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectImportServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectImportServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectImportServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectImportServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComProjectServiceImpl : the factory of component: ProjectService
type comFactory4pComProjectServiceImpl struct {

    mPrototype * projects0x4d85c7.ProjectServiceImpl

	
	mUUIDGenServiceSelector config.InjectionSelector
	mLocalRepoServiceSelector config.InjectionSelector
	mProjectTypeServiceSelector config.InjectionSelector
	mFileSystemServiceSelector config.InjectionSelector
	mLocationServiceSelector config.InjectionSelector
	mProjectDAOSelector config.InjectionSelector
	mLocalRepoDAOSelector config.InjectionSelector
	mGitLibAgentSelector config.InjectionSelector

}

func (inst * comFactory4pComProjectServiceImpl) init() application.ComponentFactory {

	
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)
	inst.mLocalRepoServiceSelector = config.NewInjectionSelector("#LocalRepositoryService",nil)
	inst.mProjectTypeServiceSelector = config.NewInjectionSelector("#ContentTypeService",nil)
	inst.mFileSystemServiceSelector = config.NewInjectionSelector("#FileSystemService",nil)
	inst.mLocationServiceSelector = config.NewInjectionSelector("#LocationService",nil)
	inst.mProjectDAOSelector = config.NewInjectionSelector("#ProjectDAO",nil)
	inst.mLocalRepoDAOSelector = config.NewInjectionSelector("#LocalRepositoryDAO",nil)
	inst.mGitLibAgentSelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComProjectServiceImpl) newObject() * projects0x4d85c7.ProjectServiceImpl {
	return & projects0x4d85c7.ProjectServiceImpl {}
}

func (inst * comFactory4pComProjectServiceImpl) castObject(instance application.ComponentInstance) * projects0x4d85c7.ProjectServiceImpl {
	return instance.Get().(*projects0x4d85c7.ProjectServiceImpl)
}

func (inst * comFactory4pComProjectServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComProjectServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComProjectServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComProjectServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComProjectServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	obj.LocalRepoService = inst.getterForFieldLocalRepoServiceSelector(context)
	obj.ProjectTypeService = inst.getterForFieldProjectTypeServiceSelector(context)
	obj.FileSystemService = inst.getterForFieldFileSystemServiceSelector(context)
	obj.LocationService = inst.getterForFieldLocationServiceSelector(context)
	obj.ProjectDAO = inst.getterForFieldProjectDAOSelector(context)
	obj.LocalRepoDAO = inst.getterForFieldLocalRepoDAOSelector(context)
	obj.GitLibAgent = inst.getterForFieldGitLibAgentSelector(context)
	return context.LastError()
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComProjectServiceImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectService")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLocalRepoServiceSelector
func (inst * comFactory4pComProjectServiceImpl) getterForFieldLocalRepoServiceSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryService {

	o1 := inst.mLocalRepoServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectService")
		eb.Set("field", "LocalRepoService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProjectTypeServiceSelector
func (inst * comFactory4pComProjectServiceImpl) getterForFieldProjectTypeServiceSelector (context application.InstanceContext) service0x3e063d.ContentTypeService {

	o1 := inst.mProjectTypeServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ContentTypeService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectService")
		eb.Set("field", "ProjectTypeService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ContentTypeService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldFileSystemServiceSelector
func (inst * comFactory4pComProjectServiceImpl) getterForFieldFileSystemServiceSelector (context application.InstanceContext) service0x3e063d.FileSystemService {

	o1 := inst.mFileSystemServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.FileSystemService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectService")
		eb.Set("field", "FileSystemService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.FileSystemService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLocationServiceSelector
func (inst * comFactory4pComProjectServiceImpl) getterForFieldLocationServiceSelector (context application.InstanceContext) service0x3e063d.LocationService {

	o1 := inst.mLocationServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocationService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectService")
		eb.Set("field", "LocationService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocationService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProjectDAOSelector
func (inst * comFactory4pComProjectServiceImpl) getterForFieldProjectDAOSelector (context application.InstanceContext) dao0x5af8d0.ProjectDAO {

	o1 := inst.mProjectDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.ProjectDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectService")
		eb.Set("field", "ProjectDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.ProjectDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLocalRepoDAOSelector
func (inst * comFactory4pComProjectServiceImpl) getterForFieldLocalRepoDAOSelector (context application.InstanceContext) dao0x5af8d0.LocalRepositoryDAO {

	o1 := inst.mLocalRepoDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.LocalRepositoryDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectService")
		eb.Set("field", "LocalRepoDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.LocalRepositoryDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldGitLibAgentSelector
func (inst * comFactory4pComProjectServiceImpl) getterForFieldGitLibAgentSelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mGitLibAgentSelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "ProjectService")
		eb.Set("field", "GitLibAgent")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalRepositoryController : the factory of component: com65-repositories0x637d5e.LocalRepositoryController
type comFactory4pComLocalRepositoryController struct {

    mPrototype * repositories0x637d5e.LocalRepositoryController

	
	mRepoServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComLocalRepositoryController) init() application.ComponentFactory {

	
	inst.mRepoServiceSelector = config.NewInjectionSelector("#LocalRepositoryService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalRepositoryController) newObject() * repositories0x637d5e.LocalRepositoryController {
	return & repositories0x637d5e.LocalRepositoryController {}
}

func (inst * comFactory4pComLocalRepositoryController) castObject(instance application.ComponentInstance) * repositories0x637d5e.LocalRepositoryController {
	return instance.Get().(*repositories0x637d5e.LocalRepositoryController)
}

func (inst * comFactory4pComLocalRepositoryController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalRepositoryController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalRepositoryController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalRepositoryController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.RepoService = inst.getterForFieldRepoServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldRepoServiceSelector
func (inst * comFactory4pComLocalRepositoryController) getterForFieldRepoServiceSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryService {

	o1 := inst.mRepoServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com65-repositories0x637d5e.LocalRepositoryController")
		eb.Set("field", "RepoService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComLocalRepositoryController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com65-repositories0x637d5e.LocalRepositoryController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryDaoImpl : the factory of component: LocalRepositoryDAO
type comFactory4pComRepositoryDaoImpl struct {

    mPrototype * repositories0x637d5e.RepositoryDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryDaoImpl) newObject() * repositories0x637d5e.RepositoryDaoImpl {
	return & repositories0x637d5e.RepositoryDaoImpl {}
}

func (inst * comFactory4pComRepositoryDaoImpl) castObject(instance application.ComponentInstance) * repositories0x637d5e.RepositoryDaoImpl {
	return instance.Get().(*repositories0x637d5e.RepositoryDaoImpl)
}

func (inst * comFactory4pComRepositoryDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositoryDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositoryDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositoryDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComRepositoryDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) dbagent0x9f90fb.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(dbagent0x9f90fb.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "dbagent0x9f90fb.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComRepositoryDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalRepositoryFinderImpl : the factory of component: LocalRepositoryFinder
type comFactory4pComLocalRepositoryFinderImpl struct {

    mPrototype * repositories0x637d5e.LocalRepositoryFinderImpl

	
	mGitLibAgentSelector config.InjectionSelector

}

func (inst * comFactory4pComLocalRepositoryFinderImpl) init() application.ComponentFactory {

	
	inst.mGitLibAgentSelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) newObject() * repositories0x637d5e.LocalRepositoryFinderImpl {
	return & repositories0x637d5e.LocalRepositoryFinderImpl {}
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) castObject(instance application.ComponentInstance) * repositories0x637d5e.LocalRepositoryFinderImpl {
	return instance.Get().(*repositories0x637d5e.LocalRepositoryFinderImpl)
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryFinderImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.GitLibAgent = inst.getterForFieldGitLibAgentSelector(context)
	return context.LastError()
}

//getterForFieldGitLibAgentSelector
func (inst * comFactory4pComLocalRepositoryFinderImpl) getterForFieldGitLibAgentSelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mGitLibAgentSelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryFinder")
		eb.Set("field", "GitLibAgent")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalRepositoryServiceImpl : the factory of component: LocalRepositoryService
type comFactory4pComLocalRepositoryServiceImpl struct {

    mPrototype * repositories0x637d5e.LocalRepositoryServiceImpl

	
	mGitLibAgentSelector config.InjectionSelector
	mLocalRepositoryDAOSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector
	mRepoFinderSelector config.InjectionSelector
	mLrStateLoaderSelector config.InjectionSelector
	mFileSystemServiceSelector config.InjectionSelector
	mLocationServiceSelector config.InjectionSelector
	mProjectServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComLocalRepositoryServiceImpl) init() application.ComponentFactory {

	
	inst.mGitLibAgentSelector = config.NewInjectionSelector("#git-lib-agent",nil)
	inst.mLocalRepositoryDAOSelector = config.NewInjectionSelector("#LocalRepositoryDAO",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)
	inst.mRepoFinderSelector = config.NewInjectionSelector("#LocalRepositoryFinder",nil)
	inst.mLrStateLoaderSelector = config.NewInjectionSelector("#LocalRepositoryStateLoader",nil)
	inst.mFileSystemServiceSelector = config.NewInjectionSelector("#FileSystemService",nil)
	inst.mLocationServiceSelector = config.NewInjectionSelector("#LocationService",nil)
	inst.mProjectServiceSelector = config.NewInjectionSelector("#ProjectService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) newObject() * repositories0x637d5e.LocalRepositoryServiceImpl {
	return & repositories0x637d5e.LocalRepositoryServiceImpl {}
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) castObject(instance application.ComponentInstance) * repositories0x637d5e.LocalRepositoryServiceImpl {
	return instance.Get().(*repositories0x637d5e.LocalRepositoryServiceImpl)
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.GitLibAgent = inst.getterForFieldGitLibAgentSelector(context)
	obj.LocalRepositoryDAO = inst.getterForFieldLocalRepositoryDAOSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	obj.RepoFinder = inst.getterForFieldRepoFinderSelector(context)
	obj.LrStateLoader = inst.getterForFieldLrStateLoaderSelector(context)
	obj.FileSystemService = inst.getterForFieldFileSystemServiceSelector(context)
	obj.LocationService = inst.getterForFieldLocationServiceSelector(context)
	obj.ProjectService = inst.getterForFieldProjectServiceSelector(context)
	return context.LastError()
}

//getterForFieldGitLibAgentSelector
func (inst * comFactory4pComLocalRepositoryServiceImpl) getterForFieldGitLibAgentSelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mGitLibAgentSelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryService")
		eb.Set("field", "GitLibAgent")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLocalRepositoryDAOSelector
func (inst * comFactory4pComLocalRepositoryServiceImpl) getterForFieldLocalRepositoryDAOSelector (context application.InstanceContext) dao0x5af8d0.LocalRepositoryDAO {

	o1 := inst.mLocalRepositoryDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.LocalRepositoryDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryService")
		eb.Set("field", "LocalRepositoryDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.LocalRepositoryDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComLocalRepositoryServiceImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryService")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldRepoFinderSelector
func (inst * comFactory4pComLocalRepositoryServiceImpl) getterForFieldRepoFinderSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryFinder {

	o1 := inst.mRepoFinderSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryFinder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryService")
		eb.Set("field", "RepoFinder")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryFinder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLrStateLoaderSelector
func (inst * comFactory4pComLocalRepositoryServiceImpl) getterForFieldLrStateLoaderSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryStateLoader {

	o1 := inst.mLrStateLoaderSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryStateLoader)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryService")
		eb.Set("field", "LrStateLoader")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryStateLoader")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldFileSystemServiceSelector
func (inst * comFactory4pComLocalRepositoryServiceImpl) getterForFieldFileSystemServiceSelector (context application.InstanceContext) service0x3e063d.FileSystemService {

	o1 := inst.mFileSystemServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.FileSystemService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryService")
		eb.Set("field", "FileSystemService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.FileSystemService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLocationServiceSelector
func (inst * comFactory4pComLocalRepositoryServiceImpl) getterForFieldLocationServiceSelector (context application.InstanceContext) service0x3e063d.LocationService {

	o1 := inst.mLocationServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocationService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryService")
		eb.Set("field", "LocationService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocationService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProjectServiceSelector
func (inst * comFactory4pComLocalRepositoryServiceImpl) getterForFieldProjectServiceSelector (context application.InstanceContext) service0x3e063d.ProjectService {

	o1 := inst.mProjectServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProjectService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryService")
		eb.Set("field", "ProjectService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProjectService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLocalRepositoryStateLoaderImpl : the factory of component: LocalRepositoryStateLoader
type comFactory4pComLocalRepositoryStateLoaderImpl struct {

    mPrototype * repositories0x637d5e.LocalRepositoryStateLoaderImpl

	
	mLocalRepoServiceSelector config.InjectionSelector
	mDaoSelector config.InjectionSelector
	mGitLibAgentSelector config.InjectionSelector

}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) init() application.ComponentFactory {

	
	inst.mLocalRepoServiceSelector = config.NewInjectionSelector("#LocalRepositoryService",nil)
	inst.mDaoSelector = config.NewInjectionSelector("#LocalRepositoryDAO",nil)
	inst.mGitLibAgentSelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) newObject() * repositories0x637d5e.LocalRepositoryStateLoaderImpl {
	return & repositories0x637d5e.LocalRepositoryStateLoaderImpl {}
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) castObject(instance application.ComponentInstance) * repositories0x637d5e.LocalRepositoryStateLoaderImpl {
	return instance.Get().(*repositories0x637d5e.LocalRepositoryStateLoaderImpl)
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.LocalRepoService = inst.getterForFieldLocalRepoServiceSelector(context)
	obj.Dao = inst.getterForFieldDaoSelector(context)
	obj.GitLibAgent = inst.getterForFieldGitLibAgentSelector(context)
	return context.LastError()
}

//getterForFieldLocalRepoServiceSelector
func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) getterForFieldLocalRepoServiceSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryService {

	o1 := inst.mLocalRepoServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryStateLoader")
		eb.Set("field", "LocalRepoService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldDaoSelector
func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) getterForFieldDaoSelector (context application.InstanceContext) dao0x5af8d0.LocalRepositoryDAO {

	o1 := inst.mDaoSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.LocalRepositoryDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryStateLoader")
		eb.Set("field", "Dao")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.LocalRepositoryDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldGitLibAgentSelector
func (inst * comFactory4pComLocalRepositoryStateLoaderImpl) getterForFieldGitLibAgentSelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mGitLibAgentSelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "LocalRepositoryStateLoader")
		eb.Set("field", "GitLibAgent")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMainRepositoryServiceImpl : the factory of component: MainRepositoryService
type comFactory4pComMainRepositoryServiceImpl struct {

    mPrototype * repositories0x637d5e.MainRepositoryServiceImpl

	
	mAppDataServiceSelector config.InjectionSelector
	mGitLASelector config.InjectionSelector

}

func (inst * comFactory4pComMainRepositoryServiceImpl) init() application.ComponentFactory {

	
	inst.mAppDataServiceSelector = config.NewInjectionSelector("#AppDataService",nil)
	inst.mGitLASelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMainRepositoryServiceImpl) newObject() * repositories0x637d5e.MainRepositoryServiceImpl {
	return & repositories0x637d5e.MainRepositoryServiceImpl {}
}

func (inst * comFactory4pComMainRepositoryServiceImpl) castObject(instance application.ComponentInstance) * repositories0x637d5e.MainRepositoryServiceImpl {
	return instance.Get().(*repositories0x637d5e.MainRepositoryServiceImpl)
}

func (inst * comFactory4pComMainRepositoryServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMainRepositoryServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMainRepositoryServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMainRepositoryServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMainRepositoryServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMainRepositoryServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AppDataService = inst.getterForFieldAppDataServiceSelector(context)
	obj.GitLA = inst.getterForFieldGitLASelector(context)
	return context.LastError()
}

//getterForFieldAppDataServiceSelector
func (inst * comFactory4pComMainRepositoryServiceImpl) getterForFieldAppDataServiceSelector (context application.InstanceContext) service0x3e063d.AppDataService {

	o1 := inst.mAppDataServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AppDataService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MainRepositoryService")
		eb.Set("field", "AppDataService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AppDataService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldGitLASelector
func (inst * comFactory4pComMainRepositoryServiceImpl) getterForFieldGitLASelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mGitLASelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "MainRepositoryService")
		eb.Set("field", "GitLA")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRemoteRepositoryController : the factory of component: com71-repositories0x637d5e.RemoteRepositoryController
type comFactory4pComRemoteRepositoryController struct {

    mPrototype * repositories0x637d5e.RemoteRepositoryController

	
	mRepoServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComRemoteRepositoryController) init() application.ComponentFactory {

	
	inst.mRepoServiceSelector = config.NewInjectionSelector("#RemoteRepositoryService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRemoteRepositoryController) newObject() * repositories0x637d5e.RemoteRepositoryController {
	return & repositories0x637d5e.RemoteRepositoryController {}
}

func (inst * comFactory4pComRemoteRepositoryController) castObject(instance application.ComponentInstance) * repositories0x637d5e.RemoteRepositoryController {
	return instance.Get().(*repositories0x637d5e.RemoteRepositoryController)
}

func (inst * comFactory4pComRemoteRepositoryController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRemoteRepositoryController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRemoteRepositoryController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRemoteRepositoryController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRemoteRepositoryController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRemoteRepositoryController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.RepoService = inst.getterForFieldRepoServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldRepoServiceSelector
func (inst * comFactory4pComRemoteRepositoryController) getterForFieldRepoServiceSelector (context application.InstanceContext) service0x3e063d.RemoteRepositoryService {

	o1 := inst.mRepoServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.RemoteRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com71-repositories0x637d5e.RemoteRepositoryController")
		eb.Set("field", "RepoService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.RemoteRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComRemoteRepositoryController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com71-repositories0x637d5e.RemoteRepositoryController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRemoteRepositoryServiceImpl : the factory of component: RemoteRepositoryService
type comFactory4pComRemoteRepositoryServiceImpl struct {

    mPrototype * repositories0x637d5e.RemoteRepositoryServiceImpl

	

}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) newObject() * repositories0x637d5e.RemoteRepositoryServiceImpl {
	return & repositories0x637d5e.RemoteRepositoryServiceImpl {}
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) castObject(instance application.ComponentInstance) * repositories0x637d5e.RemoteRepositoryServiceImpl {
	return instance.Get().(*repositories0x637d5e.RemoteRepositoryServiceImpl)
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRemoteRepositoryServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryImportController : the factory of component: com73-repositories0x637d5e.RepositoryImportController
type comFactory4pComRepositoryImportController struct {

    mPrototype * repositories0x637d5e.RepositoryImportController

	
	mRepoStateLoaderSelector config.InjectionSelector
	mImportServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryImportController) init() application.ComponentFactory {

	
	inst.mRepoStateLoaderSelector = config.NewInjectionSelector("#LocalRepositoryStateLoader",nil)
	inst.mImportServiceSelector = config.NewInjectionSelector("#RepositoryImportService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryImportController) newObject() * repositories0x637d5e.RepositoryImportController {
	return & repositories0x637d5e.RepositoryImportController {}
}

func (inst * comFactory4pComRepositoryImportController) castObject(instance application.ComponentInstance) * repositories0x637d5e.RepositoryImportController {
	return instance.Get().(*repositories0x637d5e.RepositoryImportController)
}

func (inst * comFactory4pComRepositoryImportController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositoryImportController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositoryImportController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositoryImportController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryImportController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryImportController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.RepoStateLoader = inst.getterForFieldRepoStateLoaderSelector(context)
	obj.ImportService = inst.getterForFieldImportServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldRepoStateLoaderSelector
func (inst * comFactory4pComRepositoryImportController) getterForFieldRepoStateLoaderSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryStateLoader {

	o1 := inst.mRepoStateLoaderSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryStateLoader)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com73-repositories0x637d5e.RepositoryImportController")
		eb.Set("field", "RepoStateLoader")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryStateLoader")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldImportServiceSelector
func (inst * comFactory4pComRepositoryImportController) getterForFieldImportServiceSelector (context application.InstanceContext) service0x3e063d.RepositoryImportService {

	o1 := inst.mImportServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.RepositoryImportService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com73-repositories0x637d5e.RepositoryImportController")
		eb.Set("field", "ImportService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.RepositoryImportService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComRepositoryImportController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com73-repositories0x637d5e.RepositoryImportController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepositoryImportServiceImpl : the factory of component: RepositoryImportService
type comFactory4pComRepositoryImportServiceImpl struct {

    mPrototype * repositories0x637d5e.RepositoryImportServiceImpl

	
	mRepositoryServiceSelector config.InjectionSelector
	mRepoFinderSelector config.InjectionSelector

}

func (inst * comFactory4pComRepositoryImportServiceImpl) init() application.ComponentFactory {

	
	inst.mRepositoryServiceSelector = config.NewInjectionSelector("#LocalRepositoryService",nil)
	inst.mRepoFinderSelector = config.NewInjectionSelector("#LocalRepositoryFinder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepositoryImportServiceImpl) newObject() * repositories0x637d5e.RepositoryImportServiceImpl {
	return & repositories0x637d5e.RepositoryImportServiceImpl {}
}

func (inst * comFactory4pComRepositoryImportServiceImpl) castObject(instance application.ComponentInstance) * repositories0x637d5e.RepositoryImportServiceImpl {
	return instance.Get().(*repositories0x637d5e.RepositoryImportServiceImpl)
}

func (inst * comFactory4pComRepositoryImportServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepositoryImportServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepositoryImportServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepositoryImportServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryImportServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepositoryImportServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.RepositoryService = inst.getterForFieldRepositoryServiceSelector(context)
	obj.RepoFinder = inst.getterForFieldRepoFinderSelector(context)
	return context.LastError()
}

//getterForFieldRepositoryServiceSelector
func (inst * comFactory4pComRepositoryImportServiceImpl) getterForFieldRepositoryServiceSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryService {

	o1 := inst.mRepositoryServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryImportService")
		eb.Set("field", "RepositoryService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldRepoFinderSelector
func (inst * comFactory4pComRepositoryImportServiceImpl) getterForFieldRepoFinderSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryFinder {

	o1 := inst.mRepoFinderSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryFinder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryImportService")
		eb.Set("field", "RepoFinder")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryFinder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRepoWorktreeProjectController : the factory of component: com75-repositoryworktreeproject0x399028.RepoWorktreeProjectController
type comFactory4pComRepoWorktreeProjectController struct {

    mPrototype * repositoryworktreeproject0x399028.RepoWorktreeProjectController

	
	mRWPServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComRepoWorktreeProjectController) init() application.ComponentFactory {

	
	inst.mRWPServiceSelector = config.NewInjectionSelector("#RepositoryWorktreeProjectService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRepoWorktreeProjectController) newObject() * repositoryworktreeproject0x399028.RepoWorktreeProjectController {
	return & repositoryworktreeproject0x399028.RepoWorktreeProjectController {}
}

func (inst * comFactory4pComRepoWorktreeProjectController) castObject(instance application.ComponentInstance) * repositoryworktreeproject0x399028.RepoWorktreeProjectController {
	return instance.Get().(*repositoryworktreeproject0x399028.RepoWorktreeProjectController)
}

func (inst * comFactory4pComRepoWorktreeProjectController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRepoWorktreeProjectController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRepoWorktreeProjectController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRepoWorktreeProjectController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepoWorktreeProjectController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRepoWorktreeProjectController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.RWPService = inst.getterForFieldRWPServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldRWPServiceSelector
func (inst * comFactory4pComRepoWorktreeProjectController) getterForFieldRWPServiceSelector (context application.InstanceContext) service0x3e063d.RepositoryWorktreeProjectService {

	o1 := inst.mRWPServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.RepositoryWorktreeProjectService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com75-repositoryworktreeproject0x399028.RepoWorktreeProjectController")
		eb.Set("field", "RWPService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.RepositoryWorktreeProjectService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComRepoWorktreeProjectController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com75-repositoryworktreeproject0x399028.RepoWorktreeProjectController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRWPServiceImpl : the factory of component: RepositoryWorktreeProjectService
type comFactory4pComRWPServiceImpl struct {

    mPrototype * repositoryworktreeproject0x399028.RWPServiceImpl

	
	mRepositoriesSelector config.InjectionSelector
	mWorktreesSelector config.InjectionSelector
	mProjectsSelector config.InjectionSelector
	mGitLibSelector config.InjectionSelector

}

func (inst * comFactory4pComRWPServiceImpl) init() application.ComponentFactory {

	
	inst.mRepositoriesSelector = config.NewInjectionSelector("#LocalRepositoryService",nil)
	inst.mWorktreesSelector = config.NewInjectionSelector("#WorktreeService",nil)
	inst.mProjectsSelector = config.NewInjectionSelector("#ProjectService",nil)
	inst.mGitLibSelector = config.NewInjectionSelector("#git-lib-agent",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRWPServiceImpl) newObject() * repositoryworktreeproject0x399028.RWPServiceImpl {
	return & repositoryworktreeproject0x399028.RWPServiceImpl {}
}

func (inst * comFactory4pComRWPServiceImpl) castObject(instance application.ComponentInstance) * repositoryworktreeproject0x399028.RWPServiceImpl {
	return instance.Get().(*repositoryworktreeproject0x399028.RWPServiceImpl)
}

func (inst * comFactory4pComRWPServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRWPServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRWPServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRWPServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRWPServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRWPServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Repositories = inst.getterForFieldRepositoriesSelector(context)
	obj.Worktrees = inst.getterForFieldWorktreesSelector(context)
	obj.Projects = inst.getterForFieldProjectsSelector(context)
	obj.GitLib = inst.getterForFieldGitLibSelector(context)
	return context.LastError()
}

//getterForFieldRepositoriesSelector
func (inst * comFactory4pComRWPServiceImpl) getterForFieldRepositoriesSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryService {

	o1 := inst.mRepositoriesSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryWorktreeProjectService")
		eb.Set("field", "Repositories")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldWorktreesSelector
func (inst * comFactory4pComRWPServiceImpl) getterForFieldWorktreesSelector (context application.InstanceContext) service0x3e063d.WorktreeService {

	o1 := inst.mWorktreesSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.WorktreeService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryWorktreeProjectService")
		eb.Set("field", "Worktrees")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.WorktreeService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProjectsSelector
func (inst * comFactory4pComRWPServiceImpl) getterForFieldProjectsSelector (context application.InstanceContext) service0x3e063d.ProjectService {

	o1 := inst.mProjectsSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProjectService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryWorktreeProjectService")
		eb.Set("field", "Projects")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProjectService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldGitLibSelector
func (inst * comFactory4pComRWPServiceImpl) getterForFieldGitLibSelector (context application.InstanceContext) store0x8467b3.LibAgent {

	o1 := inst.mGitLibSelector.GetOne(context)
	o2, ok := o1.(store0x8467b3.LibAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "RepositoryWorktreeProjectService")
		eb.Set("field", "GitLib")
		eb.Set("type1", "?")
		eb.Set("type2", "store0x8467b3.LibAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComSettingController : the factory of component: com77-settings0x19237d.SettingController
type comFactory4pComSettingController struct {

    mPrototype * settings0x19237d.SettingController

	
	mSettingServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComSettingController) init() application.ComponentFactory {

	
	inst.mSettingServiceSelector = config.NewInjectionSelector("#SettingService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComSettingController) newObject() * settings0x19237d.SettingController {
	return & settings0x19237d.SettingController {}
}

func (inst * comFactory4pComSettingController) castObject(instance application.ComponentInstance) * settings0x19237d.SettingController {
	return instance.Get().(*settings0x19237d.SettingController)
}

func (inst * comFactory4pComSettingController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComSettingController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComSettingController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComSettingController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSettingController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSettingController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.SettingService = inst.getterForFieldSettingServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldSettingServiceSelector
func (inst * comFactory4pComSettingController) getterForFieldSettingServiceSelector (context application.InstanceContext) service0x3e063d.SettingService {

	o1 := inst.mSettingServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.SettingService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com77-settings0x19237d.SettingController")
		eb.Set("field", "SettingService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.SettingService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComSettingController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com77-settings0x19237d.SettingController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComSettingDaoImpl : the factory of component: SettingDAO
type comFactory4pComSettingDaoImpl struct {

    mPrototype * settings0x19237d.SettingDaoImpl

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComSettingDaoImpl) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComSettingDaoImpl) newObject() * settings0x19237d.SettingDaoImpl {
	return & settings0x19237d.SettingDaoImpl {}
}

func (inst * comFactory4pComSettingDaoImpl) castObject(instance application.ComponentInstance) * settings0x19237d.SettingDaoImpl {
	return instance.Get().(*settings0x19237d.SettingDaoImpl)
}

func (inst * comFactory4pComSettingDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComSettingDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComSettingDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComSettingDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSettingDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSettingDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComSettingDaoImpl) getterForFieldAgentSelector (context application.InstanceContext) dbagent0x9f90fb.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(dbagent0x9f90fb.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SettingDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "dbagent0x9f90fb.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComSettingDaoImpl) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SettingDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComSettingServiceImpl : the factory of component: SettingService
type comFactory4pComSettingServiceImpl struct {

    mPrototype * settings0x19237d.SettingServiceImpl

	
	mSettingDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComSettingServiceImpl) init() application.ComponentFactory {

	
	inst.mSettingDAOSelector = config.NewInjectionSelector("#SettingDAO",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComSettingServiceImpl) newObject() * settings0x19237d.SettingServiceImpl {
	return & settings0x19237d.SettingServiceImpl {}
}

func (inst * comFactory4pComSettingServiceImpl) castObject(instance application.ComponentInstance) * settings0x19237d.SettingServiceImpl {
	return instance.Get().(*settings0x19237d.SettingServiceImpl)
}

func (inst * comFactory4pComSettingServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComSettingServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComSettingServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComSettingServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSettingServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSettingServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.SettingDAO = inst.getterForFieldSettingDAOSelector(context)
	return context.LastError()
}

//getterForFieldSettingDAOSelector
func (inst * comFactory4pComSettingServiceImpl) getterForFieldSettingDAOSelector (context application.InstanceContext) dao0x5af8d0.SettingDAO {

	o1 := inst.mSettingDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.SettingDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SettingService")
		eb.Set("field", "SettingDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.SettingDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComWpmSetupController : the factory of component: com80-setup0xd9ff02.WpmSetupController
type comFactory4pComWpmSetupController struct {

    mPrototype * setup0xd9ff02.WpmSetupController

	
	mResponderSelector config.InjectionSelector
	mSetupServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComWpmSetupController) init() application.ComponentFactory {

	
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)
	inst.mSetupServiceSelector = config.NewInjectionSelector("#SetupService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComWpmSetupController) newObject() * setup0xd9ff02.WpmSetupController {
	return & setup0xd9ff02.WpmSetupController {}
}

func (inst * comFactory4pComWpmSetupController) castObject(instance application.ComponentInstance) * setup0xd9ff02.WpmSetupController {
	return instance.Get().(*setup0xd9ff02.WpmSetupController)
}

func (inst * comFactory4pComWpmSetupController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComWpmSetupController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComWpmSetupController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComWpmSetupController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWpmSetupController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWpmSetupController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	obj.SetupService = inst.getterForFieldSetupServiceSelector(context)
	return context.LastError()
}

//getterForFieldResponderSelector
func (inst * comFactory4pComWpmSetupController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com80-setup0xd9ff02.WpmSetupController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSetupServiceSelector
func (inst * comFactory4pComWpmSetupController) getterForFieldSetupServiceSelector (context application.InstanceContext) service0x3e063d.SetupService {

	o1 := inst.mSetupServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.SetupService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com80-setup0xd9ff02.WpmSetupController")
		eb.Set("field", "SetupService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.SetupService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpSetupService : the factory of component: SetupService
type comFactory4pComImpSetupService struct {

    mPrototype * setup0xd9ff02.ImpSetupService

	
	mAppDataServiceSelector config.InjectionSelector
	mFileSystemServiceSelector config.InjectionSelector
	mExecutableServiceSelector config.InjectionSelector
	mExecutableImportServiceSelector config.InjectionSelector
	mIntentTemplateServiceSelector config.InjectionSelector
	mMediaServiceSelector config.InjectionSelector
	mProjectTypeImportServiceSelector config.InjectionSelector
	mSettingServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComImpSetupService) init() application.ComponentFactory {

	
	inst.mAppDataServiceSelector = config.NewInjectionSelector("#AppDataService",nil)
	inst.mFileSystemServiceSelector = config.NewInjectionSelector("#FileSystemService",nil)
	inst.mExecutableServiceSelector = config.NewInjectionSelector("#ExecutableService",nil)
	inst.mExecutableImportServiceSelector = config.NewInjectionSelector("#ExecutableImportService",nil)
	inst.mIntentTemplateServiceSelector = config.NewInjectionSelector("#IntentTemplateService",nil)
	inst.mMediaServiceSelector = config.NewInjectionSelector("#MediaService",nil)
	inst.mProjectTypeImportServiceSelector = config.NewInjectionSelector("#ProjectTypeImportService",nil)
	inst.mSettingServiceSelector = config.NewInjectionSelector("#SettingService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpSetupService) newObject() * setup0xd9ff02.ImpSetupService {
	return & setup0xd9ff02.ImpSetupService {}
}

func (inst * comFactory4pComImpSetupService) castObject(instance application.ComponentInstance) * setup0xd9ff02.ImpSetupService {
	return instance.Get().(*setup0xd9ff02.ImpSetupService)
}

func (inst * comFactory4pComImpSetupService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpSetupService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpSetupService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpSetupService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpSetupService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpSetupService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AppDataService = inst.getterForFieldAppDataServiceSelector(context)
	obj.FileSystemService = inst.getterForFieldFileSystemServiceSelector(context)
	obj.ExecutableService = inst.getterForFieldExecutableServiceSelector(context)
	obj.ExecutableImportService = inst.getterForFieldExecutableImportServiceSelector(context)
	obj.IntentTemplateService = inst.getterForFieldIntentTemplateServiceSelector(context)
	obj.MediaService = inst.getterForFieldMediaServiceSelector(context)
	obj.ProjectTypeImportService = inst.getterForFieldProjectTypeImportServiceSelector(context)
	obj.SettingService = inst.getterForFieldSettingServiceSelector(context)
	return context.LastError()
}

//getterForFieldAppDataServiceSelector
func (inst * comFactory4pComImpSetupService) getterForFieldAppDataServiceSelector (context application.InstanceContext) service0x3e063d.AppDataService {

	o1 := inst.mAppDataServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AppDataService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SetupService")
		eb.Set("field", "AppDataService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AppDataService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldFileSystemServiceSelector
func (inst * comFactory4pComImpSetupService) getterForFieldFileSystemServiceSelector (context application.InstanceContext) service0x3e063d.FileSystemService {

	o1 := inst.mFileSystemServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.FileSystemService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SetupService")
		eb.Set("field", "FileSystemService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.FileSystemService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldExecutableServiceSelector
func (inst * comFactory4pComImpSetupService) getterForFieldExecutableServiceSelector (context application.InstanceContext) service0x3e063d.ExecutableService {

	o1 := inst.mExecutableServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ExecutableService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SetupService")
		eb.Set("field", "ExecutableService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ExecutableService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldExecutableImportServiceSelector
func (inst * comFactory4pComImpSetupService) getterForFieldExecutableImportServiceSelector (context application.InstanceContext) service0x3e063d.ExecutableImportService {

	o1 := inst.mExecutableImportServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ExecutableImportService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SetupService")
		eb.Set("field", "ExecutableImportService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ExecutableImportService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldIntentTemplateServiceSelector
func (inst * comFactory4pComImpSetupService) getterForFieldIntentTemplateServiceSelector (context application.InstanceContext) service0x3e063d.IntentTemplateService {

	o1 := inst.mIntentTemplateServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.IntentTemplateService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SetupService")
		eb.Set("field", "IntentTemplateService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.IntentTemplateService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldMediaServiceSelector
func (inst * comFactory4pComImpSetupService) getterForFieldMediaServiceSelector (context application.InstanceContext) service0x3e063d.MediaService {

	o1 := inst.mMediaServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.MediaService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SetupService")
		eb.Set("field", "MediaService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.MediaService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProjectTypeImportServiceSelector
func (inst * comFactory4pComImpSetupService) getterForFieldProjectTypeImportServiceSelector (context application.InstanceContext) service0x3e063d.ProjectTypeImportService {

	o1 := inst.mProjectTypeImportServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProjectTypeImportService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SetupService")
		eb.Set("field", "ProjectTypeImportService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProjectTypeImportService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSettingServiceSelector
func (inst * comFactory4pComImpSetupService) getterForFieldSettingServiceSelector (context application.InstanceContext) service0x3e063d.SettingService {

	o1 := inst.mSettingServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.SettingService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "SetupService")
		eb.Set("field", "SettingService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.SettingService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpWorktreeDao : the factory of component: WorktreeDAO
type comFactory4pComImpWorktreeDao struct {

    mPrototype * worktrees0xa762f3.ImpWorktreeDao

	
	mAgentSelector config.InjectionSelector
	mUUIDGenServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComImpWorktreeDao) init() application.ComponentFactory {

	
	inst.mAgentSelector = config.NewInjectionSelector("#GormDBAgent",nil)
	inst.mUUIDGenServiceSelector = config.NewInjectionSelector("#UUIDGenService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpWorktreeDao) newObject() * worktrees0xa762f3.ImpWorktreeDao {
	return & worktrees0xa762f3.ImpWorktreeDao {}
}

func (inst * comFactory4pComImpWorktreeDao) castObject(instance application.ComponentInstance) * worktrees0xa762f3.ImpWorktreeDao {
	return instance.Get().(*worktrees0xa762f3.ImpWorktreeDao)
}

func (inst * comFactory4pComImpWorktreeDao) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpWorktreeDao) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpWorktreeDao) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpWorktreeDao) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpWorktreeDao) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpWorktreeDao) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Agent = inst.getterForFieldAgentSelector(context)
	obj.UUIDGenService = inst.getterForFieldUUIDGenServiceSelector(context)
	return context.LastError()
}

//getterForFieldAgentSelector
func (inst * comFactory4pComImpWorktreeDao) getterForFieldAgentSelector (context application.InstanceContext) dbagent0x9f90fb.GormDBAgent {

	o1 := inst.mAgentSelector.GetOne(context)
	o2, ok := o1.(dbagent0x9f90fb.GormDBAgent)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "WorktreeDAO")
		eb.Set("field", "Agent")
		eb.Set("type1", "?")
		eb.Set("type2", "dbagent0x9f90fb.GormDBAgent")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGenServiceSelector
func (inst * comFactory4pComImpWorktreeDao) getterForFieldUUIDGenServiceSelector (context application.InstanceContext) service0x3e063d.UUIDGenService {

	o1 := inst.mUUIDGenServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.UUIDGenService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "WorktreeDAO")
		eb.Set("field", "UUIDGenService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.UUIDGenService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComImpWorktreeService : the factory of component: WorktreeService
type comFactory4pComImpWorktreeService struct {

    mPrototype * worktrees0xa762f3.ImpWorktreeService

	
	mFileSystemServiceSelector config.InjectionSelector
	mLocationServiceSelector config.InjectionSelector
	mRepoFinderSelector config.InjectionSelector
	mDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComImpWorktreeService) init() application.ComponentFactory {

	
	inst.mFileSystemServiceSelector = config.NewInjectionSelector("#FileSystemService",nil)
	inst.mLocationServiceSelector = config.NewInjectionSelector("#LocationService",nil)
	inst.mRepoFinderSelector = config.NewInjectionSelector("#LocalRepositoryFinder",nil)
	inst.mDAOSelector = config.NewInjectionSelector("#WorktreeDAO",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComImpWorktreeService) newObject() * worktrees0xa762f3.ImpWorktreeService {
	return & worktrees0xa762f3.ImpWorktreeService {}
}

func (inst * comFactory4pComImpWorktreeService) castObject(instance application.ComponentInstance) * worktrees0xa762f3.ImpWorktreeService {
	return instance.Get().(*worktrees0xa762f3.ImpWorktreeService)
}

func (inst * comFactory4pComImpWorktreeService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComImpWorktreeService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComImpWorktreeService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComImpWorktreeService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpWorktreeService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComImpWorktreeService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.FileSystemService = inst.getterForFieldFileSystemServiceSelector(context)
	obj.LocationService = inst.getterForFieldLocationServiceSelector(context)
	obj.RepoFinder = inst.getterForFieldRepoFinderSelector(context)
	obj.DAO = inst.getterForFieldDAOSelector(context)
	return context.LastError()
}

//getterForFieldFileSystemServiceSelector
func (inst * comFactory4pComImpWorktreeService) getterForFieldFileSystemServiceSelector (context application.InstanceContext) service0x3e063d.FileSystemService {

	o1 := inst.mFileSystemServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.FileSystemService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "WorktreeService")
		eb.Set("field", "FileSystemService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.FileSystemService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLocationServiceSelector
func (inst * comFactory4pComImpWorktreeService) getterForFieldLocationServiceSelector (context application.InstanceContext) service0x3e063d.LocationService {

	o1 := inst.mLocationServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocationService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "WorktreeService")
		eb.Set("field", "LocationService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocationService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldRepoFinderSelector
func (inst * comFactory4pComImpWorktreeService) getterForFieldRepoFinderSelector (context application.InstanceContext) service0x3e063d.LocalRepositoryFinder {

	o1 := inst.mRepoFinderSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.LocalRepositoryFinder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "WorktreeService")
		eb.Set("field", "RepoFinder")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.LocalRepositoryFinder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldDAOSelector
func (inst * comFactory4pComImpWorktreeService) getterForFieldDAOSelector (context application.InstanceContext) dao0x5af8d0.WorktreeDAO {

	o1 := inst.mDAOSelector.GetOne(context)
	o2, ok := o1.(dao0x5af8d0.WorktreeDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "WorktreeService")
		eb.Set("field", "DAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0x5af8d0.WorktreeDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComWpmDataSource : the factory of component: com84-support0xf47d7f.WpmDataSource
type comFactory4pComWpmDataSource struct {

    mPrototype * support0xf47d7f.WpmDataSource

	
	mDMSelector config.InjectionSelector
	mAppDataServiceSelector config.InjectionSelector
	mAboutServiceSelector config.InjectionSelector
	mDriverSelector config.InjectionSelector
	mHostSelector config.InjectionSelector
	mPortSelector config.InjectionSelector
	mUserNameSelector config.InjectionSelector
	mPasswordSelector config.InjectionSelector
	mDatabaseSelector config.InjectionSelector

}

func (inst * comFactory4pComWpmDataSource) init() application.ComponentFactory {

	
	inst.mDMSelector = config.NewInjectionSelector("#starter-gorm-driver-manager",nil)
	inst.mAppDataServiceSelector = config.NewInjectionSelector("#AppDataService",nil)
	inst.mAboutServiceSelector = config.NewInjectionSelector("#AboutService",nil)
	inst.mDriverSelector = config.NewInjectionSelector("${datasource.wpm.driver}",nil)
	inst.mHostSelector = config.NewInjectionSelector("${datasource.wpm.host}",nil)
	inst.mPortSelector = config.NewInjectionSelector("${datasource.wpm.port}",nil)
	inst.mUserNameSelector = config.NewInjectionSelector("${datasource.wpm.username}",nil)
	inst.mPasswordSelector = config.NewInjectionSelector("${datasource.wpm.password}",nil)
	inst.mDatabaseSelector = config.NewInjectionSelector("${datasource.wpm.database}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComWpmDataSource) newObject() * support0xf47d7f.WpmDataSource {
	return & support0xf47d7f.WpmDataSource {}
}

func (inst * comFactory4pComWpmDataSource) castObject(instance application.ComponentInstance) * support0xf47d7f.WpmDataSource {
	return instance.Get().(*support0xf47d7f.WpmDataSource)
}

func (inst * comFactory4pComWpmDataSource) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComWpmDataSource) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComWpmDataSource) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComWpmDataSource) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWpmDataSource) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWpmDataSource) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.DM = inst.getterForFieldDMSelector(context)
	obj.AppDataService = inst.getterForFieldAppDataServiceSelector(context)
	obj.AboutService = inst.getterForFieldAboutServiceSelector(context)
	obj.Driver = inst.getterForFieldDriverSelector(context)
	obj.Host = inst.getterForFieldHostSelector(context)
	obj.Port = inst.getterForFieldPortSelector(context)
	obj.UserName = inst.getterForFieldUserNameSelector(context)
	obj.Password = inst.getterForFieldPasswordSelector(context)
	obj.Database = inst.getterForFieldDatabaseSelector(context)
	return context.LastError()
}

//getterForFieldDMSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldDMSelector (context application.InstanceContext) datasource0x68a737.DriverManager {

	o1 := inst.mDMSelector.GetOne(context)
	o2, ok := o1.(datasource0x68a737.DriverManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com84-support0xf47d7f.WpmDataSource")
		eb.Set("field", "DM")
		eb.Set("type1", "?")
		eb.Set("type2", "datasource0x68a737.DriverManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldAppDataServiceSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldAppDataServiceSelector (context application.InstanceContext) service0x3e063d.AppDataService {

	o1 := inst.mAppDataServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AppDataService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com84-support0xf47d7f.WpmDataSource")
		eb.Set("field", "AppDataService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AppDataService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldAboutServiceSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldAboutServiceSelector (context application.InstanceContext) service0x3e063d.AboutService {

	o1 := inst.mAboutServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AboutService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com84-support0xf47d7f.WpmDataSource")
		eb.Set("field", "AboutService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AboutService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldDriverSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldDriverSelector (context application.InstanceContext) string {
    return inst.mDriverSelector.GetString(context)
}

//getterForFieldHostSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldHostSelector (context application.InstanceContext) string {
    return inst.mHostSelector.GetString(context)
}

//getterForFieldPortSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldPortSelector (context application.InstanceContext) int {
    return inst.mPortSelector.GetInt(context)
}

//getterForFieldUserNameSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldUserNameSelector (context application.InstanceContext) string {
    return inst.mUserNameSelector.GetString(context)
}

//getterForFieldPasswordSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldPasswordSelector (context application.InstanceContext) string {
    return inst.mPasswordSelector.GetString(context)
}

//getterForFieldDatabaseSelector
func (inst * comFactory4pComWpmDataSource) getterForFieldDatabaseSelector (context application.InstanceContext) string {
    return inst.mDatabaseSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFileSystemHandler : the factory of component: com85-handlers0x162741.FileSystemHandler
type comFactory4pComFileSystemHandler struct {

    mPrototype * handlers0x162741.FileSystemHandler

	

}

func (inst * comFactory4pComFileSystemHandler) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFileSystemHandler) newObject() * handlers0x162741.FileSystemHandler {
	return & handlers0x162741.FileSystemHandler {}
}

func (inst * comFactory4pComFileSystemHandler) castObject(instance application.ComponentInstance) * handlers0x162741.FileSystemHandler {
	return instance.Get().(*handlers0x162741.FileSystemHandler)
}

func (inst * comFactory4pComFileSystemHandler) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFileSystemHandler) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFileSystemHandler) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFileSystemHandler) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileSystemHandler) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileSystemHandler) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAboutController : the factory of component: com86-controller0x9dc399.AboutController
type comFactory4pComAboutController struct {

    mPrototype * controller0x9dc399.AboutController

	
	mAboutServiceSelector config.InjectionSelector
	mUpdateServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector
	mProfileSelector config.InjectionSelector

}

func (inst * comFactory4pComAboutController) init() application.ComponentFactory {

	
	inst.mAboutServiceSelector = config.NewInjectionSelector("#AboutService",nil)
	inst.mUpdateServiceSelector = config.NewInjectionSelector("#CheckUpdateService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)
	inst.mProfileSelector = config.NewInjectionSelector("${application.profiles.active}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAboutController) newObject() * controller0x9dc399.AboutController {
	return & controller0x9dc399.AboutController {}
}

func (inst * comFactory4pComAboutController) castObject(instance application.ComponentInstance) * controller0x9dc399.AboutController {
	return instance.Get().(*controller0x9dc399.AboutController)
}

func (inst * comFactory4pComAboutController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAboutController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAboutController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAboutController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAboutController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAboutController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AboutService = inst.getterForFieldAboutServiceSelector(context)
	obj.UpdateService = inst.getterForFieldUpdateServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	obj.Profile = inst.getterForFieldProfileSelector(context)
	return context.LastError()
}

//getterForFieldAboutServiceSelector
func (inst * comFactory4pComAboutController) getterForFieldAboutServiceSelector (context application.InstanceContext) service0x3e063d.AboutService {

	o1 := inst.mAboutServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.AboutService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com86-controller0x9dc399.AboutController")
		eb.Set("field", "AboutService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.AboutService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUpdateServiceSelector
func (inst * comFactory4pComAboutController) getterForFieldUpdateServiceSelector (context application.InstanceContext) service0x3e063d.CheckUpdateService {

	o1 := inst.mUpdateServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.CheckUpdateService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com86-controller0x9dc399.AboutController")
		eb.Set("field", "UpdateService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.CheckUpdateService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComAboutController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com86-controller0x9dc399.AboutController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProfileSelector
func (inst * comFactory4pComAboutController) getterForFieldProfileSelector (context application.InstanceContext) string {
    return inst.mProfileSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAuthController : the factory of component: com87-controller0x9dc399.AuthController
type comFactory4pComAuthController struct {

    mPrototype * controller0x9dc399.AuthController

	
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComAuthController) init() application.ComponentFactory {

	
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAuthController) newObject() * controller0x9dc399.AuthController {
	return & controller0x9dc399.AuthController {}
}

func (inst * comFactory4pComAuthController) castObject(instance application.ComponentInstance) * controller0x9dc399.AuthController {
	return instance.Get().(*controller0x9dc399.AuthController)
}

func (inst * comFactory4pComAuthController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAuthController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAuthController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAuthController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAuthController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAuthController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldResponderSelector
func (inst * comFactory4pComAuthController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com87-controller0x9dc399.AuthController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExampleController : the factory of component: com88-controller0x9dc399.ExampleController
type comFactory4pComExampleController struct {

    mPrototype * controller0x9dc399.ExampleController

	
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComExampleController) init() application.ComponentFactory {

	
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExampleController) newObject() * controller0x9dc399.ExampleController {
	return & controller0x9dc399.ExampleController {}
}

func (inst * comFactory4pComExampleController) castObject(instance application.ComponentInstance) * controller0x9dc399.ExampleController {
	return instance.Get().(*controller0x9dc399.ExampleController)
}

func (inst * comFactory4pComExampleController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExampleController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExampleController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExampleController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldResponderSelector
func (inst * comFactory4pComExampleController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com88-controller0x9dc399.ExampleController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFileQueryController : the factory of component: com89-controller0x9dc399.FileQueryController
type comFactory4pComFileQueryController struct {

    mPrototype * controller0x9dc399.FileQueryController

	
	mFileQueryServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComFileQueryController) init() application.ComponentFactory {

	
	inst.mFileQueryServiceSelector = config.NewInjectionSelector("#FileQueryService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFileQueryController) newObject() * controller0x9dc399.FileQueryController {
	return & controller0x9dc399.FileQueryController {}
}

func (inst * comFactory4pComFileQueryController) castObject(instance application.ComponentInstance) * controller0x9dc399.FileQueryController {
	return instance.Get().(*controller0x9dc399.FileQueryController)
}

func (inst * comFactory4pComFileQueryController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFileQueryController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFileQueryController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFileQueryController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileQueryController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFileQueryController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.FileQueryService = inst.getterForFieldFileQueryServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldFileQueryServiceSelector
func (inst * comFactory4pComFileQueryController) getterForFieldFileQueryServiceSelector (context application.InstanceContext) service0x3e063d.FileQueryService {

	o1 := inst.mFileQueryServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.FileQueryService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com89-controller0x9dc399.FileQueryController")
		eb.Set("field", "FileQueryService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.FileQueryService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComFileQueryController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com89-controller0x9dc399.FileQueryController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComOnlineDocumentExampleController : the factory of component: com90-controller0x9dc399.OnlineDocumentExampleController
type comFactory4pComOnlineDocumentExampleController struct {

    mPrototype * controller0x9dc399.OnlineDocumentExampleController

	
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComOnlineDocumentExampleController) init() application.ComponentFactory {

	
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComOnlineDocumentExampleController) newObject() * controller0x9dc399.OnlineDocumentExampleController {
	return & controller0x9dc399.OnlineDocumentExampleController {}
}

func (inst * comFactory4pComOnlineDocumentExampleController) castObject(instance application.ComponentInstance) * controller0x9dc399.OnlineDocumentExampleController {
	return instance.Get().(*controller0x9dc399.OnlineDocumentExampleController)
}

func (inst * comFactory4pComOnlineDocumentExampleController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComOnlineDocumentExampleController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComOnlineDocumentExampleController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComOnlineDocumentExampleController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComOnlineDocumentExampleController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComOnlineDocumentExampleController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldResponderSelector
func (inst * comFactory4pComOnlineDocumentExampleController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com90-controller0x9dc399.OnlineDocumentExampleController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPlatformController : the factory of component: com91-controller0x9dc399.PlatformController
type comFactory4pComPlatformController struct {

    mPrototype * controller0x9dc399.PlatformController

	
	mPlatformServiceSelector config.InjectionSelector
	mProfileServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComPlatformController) init() application.ComponentFactory {

	
	inst.mPlatformServiceSelector = config.NewInjectionSelector("#PlatformService",nil)
	inst.mProfileServiceSelector = config.NewInjectionSelector("#ProfileService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPlatformController) newObject() * controller0x9dc399.PlatformController {
	return & controller0x9dc399.PlatformController {}
}

func (inst * comFactory4pComPlatformController) castObject(instance application.ComponentInstance) * controller0x9dc399.PlatformController {
	return instance.Get().(*controller0x9dc399.PlatformController)
}

func (inst * comFactory4pComPlatformController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPlatformController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPlatformController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPlatformController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPlatformController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPlatformController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.PlatformService = inst.getterForFieldPlatformServiceSelector(context)
	obj.ProfileService = inst.getterForFieldProfileServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldPlatformServiceSelector
func (inst * comFactory4pComPlatformController) getterForFieldPlatformServiceSelector (context application.InstanceContext) service0x3e063d.PlatformService {

	o1 := inst.mPlatformServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.PlatformService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com91-controller0x9dc399.PlatformController")
		eb.Set("field", "PlatformService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.PlatformService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldProfileServiceSelector
func (inst * comFactory4pComPlatformController) getterForFieldProfileServiceSelector (context application.InstanceContext) service0x3e063d.ProfileService {

	o1 := inst.mProfileServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.ProfileService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com91-controller0x9dc399.PlatformController")
		eb.Set("field", "ProfileService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.ProfileService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComPlatformController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com91-controller0x9dc399.PlatformController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUploadController : the factory of component: com92-controller0x9dc399.UploadController
type comFactory4pComUploadController struct {

    mPrototype * controller0x9dc399.UploadController

	
	mFileSystemServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComUploadController) init() application.ComponentFactory {

	
	inst.mFileSystemServiceSelector = config.NewInjectionSelector("#FileSystemService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUploadController) newObject() * controller0x9dc399.UploadController {
	return & controller0x9dc399.UploadController {}
}

func (inst * comFactory4pComUploadController) castObject(instance application.ComponentInstance) * controller0x9dc399.UploadController {
	return instance.Get().(*controller0x9dc399.UploadController)
}

func (inst * comFactory4pComUploadController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComUploadController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComUploadController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComUploadController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUploadController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUploadController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.FileSystemService = inst.getterForFieldFileSystemServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldFileSystemServiceSelector
func (inst * comFactory4pComUploadController) getterForFieldFileSystemServiceSelector (context application.InstanceContext) service0x3e063d.FileSystemService {

	o1 := inst.mFileSystemServiceSelector.GetOne(context)
	o2, ok := o1.(service0x3e063d.FileSystemService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com92-controller0x9dc399.UploadController")
		eb.Set("field", "FileSystemService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x3e063d.FileSystemService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComUploadController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com92-controller0x9dc399.UploadController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHostFilter : the factory of component: com93-filter0x8aa8f6.HostFilter
type comFactory4pComHostFilter struct {

    mPrototype * filter0x8aa8f6.HostFilter

	

}

func (inst * comFactory4pComHostFilter) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComHostFilter) newObject() * filter0x8aa8f6.HostFilter {
	return & filter0x8aa8f6.HostFilter {}
}

func (inst * comFactory4pComHostFilter) castObject(instance application.ComponentInstance) * filter0x8aa8f6.HostFilter {
	return instance.Get().(*filter0x8aa8f6.HostFilter)
}

func (inst * comFactory4pComHostFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComHostFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComHostFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComHostFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHostFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHostFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHTTP404Filter : the factory of component: com94-filter0x8aa8f6.HTTP404Filter
type comFactory4pComHTTP404Filter struct {

    mPrototype * filter0x8aa8f6.HTTP404Filter

	
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComHTTP404Filter) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComHTTP404Filter) newObject() * filter0x8aa8f6.HTTP404Filter {
	return & filter0x8aa8f6.HTTP404Filter {}
}

func (inst * comFactory4pComHTTP404Filter) castObject(instance application.ComponentInstance) * filter0x8aa8f6.HTTP404Filter {
	return instance.Get().(*filter0x8aa8f6.HTTP404Filter)
}

func (inst * comFactory4pComHTTP404Filter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComHTTP404Filter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComHTTP404Filter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComHTTP404Filter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTP404Filter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTP404Filter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComHTTP404Filter) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}




