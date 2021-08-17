// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package etcwpm

import(
	errors "errors"
	repository_d59845 "github.com/bitwormhole/gitlib/repository"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	dao_3942a8 "github.com/bitwormhole/wpm/data/dao"
	service_21941a "github.com/bitwormhole/wpm/service"
	controller_9ef436 "github.com/bitwormhole/wpm/web/controller"
)


func autoGenConfig(configbuilder application.ConfigBuilder) error {

	cominfobuilder := &config.ComInfoBuilder{}
	err := errors.New("OK")

    
	// theRepositoryDAO
	cominfobuilder.Reset()
	cominfobuilder.ID("data-repository-dao").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &dao_3942a8.RepositoryImpl{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theRepositoryDAO{}
		adapter.instance = o.(*dao_3942a8.RepositoryImpl)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }

	// theRepositoryService
	cominfobuilder.Reset()
	cominfobuilder.ID("repository-service").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &service_21941a.RepositoryServiceImpl{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theRepositoryService{}
		adapter.instance = o.(*service_21941a.RepositoryServiceImpl)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }

	// theRestControllerAuth
	cominfobuilder.Reset()
	cominfobuilder.ID("theRestControllerAuth").Class("rest-controller").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &controller_9ef436.AuthController{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theRestControllerAuth{}
		adapter.instance = o.(*controller_9ef436.AuthController)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }

	// theRestControllerDebug
	cominfobuilder.Reset()
	cominfobuilder.ID("theRestControllerDebug").Class("rest-controller").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &controller_9ef436.DebugController{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theRestControllerDebug{}
		adapter.instance = o.(*controller_9ef436.DebugController)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }

	// theRestControllerRepository
	cominfobuilder.Reset()
	cominfobuilder.ID("theRestControllerRepository").Class("rest-controller").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &controller_9ef436.RepositoryController{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theRestControllerRepository{}
		adapter.instance = o.(*controller_9ef436.RepositoryController)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }

	// theStore
	cominfobuilder.Reset()
	cominfobuilder.ID("store").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &dao_3942a8.Store{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theStore{}
		adapter.instance = o.(*dao_3942a8.Store)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }

	// theSystemConfigDAO
	cominfobuilder.Reset()
	cominfobuilder.ID("data-system-config-dao").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &dao_3942a8.SystemConfigImpl{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theSystemConfigDAO{}
		adapter.instance = o.(*dao_3942a8.SystemConfigImpl)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }

	// theTestGitController
	cominfobuilder.Reset()
	cominfobuilder.ID("theTestGitController").Class("rest-controller").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &controller_9ef436.TestGitController{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theTestGitController{}
		adapter.instance = o.(*controller_9ef436.TestGitController)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }

	// theUserConfigDAO
	cominfobuilder.Reset()
	cominfobuilder.ID("data-user-config-dao").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &dao_3942a8.UserConfigImpl{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theUserConfigDAO{}
		adapter.instance = o.(*dao_3942a8.UserConfigImpl)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }


	return nil
}


////////////////////////////////////////////////////////////////////////////////
// type theRepositoryDAO struct

func (inst *theRepositoryDAO) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters


	// to instance


	// invoke custom inject method


	return injection.Close()
}

////////////////////////////////////////////////////////////////////////////////
// type theRepositoryService struct

func (inst *theRepositoryService) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters
	inst.Store=inst.__get_Store__(injection, "#store")


	// to instance
	instance.Store=inst.Store


	// invoke custom inject method


	return injection.Close()
}

func (inst * theRepositoryService) __get_Store__(injection application.Injection,selector string) *dao_3942a8.Store {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(*dao_3942a8.Store)
	if !ok {
		err := errors.New("cannot cast component instance to type: *dao_3942a8.Store")
		injection.OnError(err)
		return nil
	}

	return o2

}

////////////////////////////////////////////////////////////////////////////////
// type theRestControllerAuth struct

func (inst *theRestControllerAuth) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters


	// to instance


	// invoke custom inject method


	return injection.Close()
}

////////////////////////////////////////////////////////////////////////////////
// type theRestControllerDebug struct

func (inst *theRestControllerDebug) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters
	inst.AppContext=inst.__get_AppContext__(injection, "context")


	// to instance
	instance.AppContext=inst.AppContext


	// invoke custom inject method


	return injection.Close()
}

func (inst * theRestControllerDebug) __get_AppContext__(injection application.Injection,selector string) application.Context {
	return injection.Context()
}

////////////////////////////////////////////////////////////////////////////////
// type theRestControllerRepository struct

func (inst *theRestControllerRepository) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters
	inst.RepoService=inst.__get_RepoService__(injection, "#repository-service")


	// to instance
	instance.RepoService=inst.RepoService


	// invoke custom inject method


	return injection.Close()
}

func (inst * theRestControllerRepository) __get_RepoService__(injection application.Injection,selector string) service_21941a.RepositoryService {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(service_21941a.RepositoryService)
	if !ok {
		err := errors.New("cannot cast component instance to type: service_21941a.RepositoryService")
		injection.OnError(err)
		return nil
	}

	return o2

}

////////////////////////////////////////////////////////////////////////////////
// type theStore struct

func (inst *theStore) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters
	inst.RepositoryDAO=inst.__get_RepositoryDAO__(injection, "#data-repository-dao")
	inst.SystemConfigDAO=inst.__get_SystemConfigDAO__(injection, "#data-system-config-dao")
	inst.UserConfigDAO=inst.__get_UserConfigDAO__(injection, "#data-user-config-dao")


	// to instance
	instance.RepositoryDAO=inst.RepositoryDAO
	instance.SystemConfigDAO=inst.SystemConfigDAO
	instance.UserConfigDAO=inst.UserConfigDAO


	// invoke custom inject method


	return injection.Close()
}

func (inst * theStore) __get_RepositoryDAO__(injection application.Injection,selector string) dao_3942a8.Repository {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(dao_3942a8.Repository)
	if !ok {
		err := errors.New("cannot cast component instance to type: dao_3942a8.Repository")
		injection.OnError(err)
		return nil
	}

	return o2

}

func (inst * theStore) __get_SystemConfigDAO__(injection application.Injection,selector string) dao_3942a8.SystemConfig {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(dao_3942a8.SystemConfig)
	if !ok {
		err := errors.New("cannot cast component instance to type: dao_3942a8.SystemConfig")
		injection.OnError(err)
		return nil
	}

	return o2

}

func (inst * theStore) __get_UserConfigDAO__(injection application.Injection,selector string) dao_3942a8.UserConfig {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(dao_3942a8.UserConfig)
	if !ok {
		err := errors.New("cannot cast component instance to type: dao_3942a8.UserConfig")
		injection.OnError(err)
		return nil
	}

	return o2

}

////////////////////////////////////////////////////////////////////////////////
// type theSystemConfigDAO struct

func (inst *theSystemConfigDAO) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters


	// to instance


	// invoke custom inject method


	return injection.Close()
}

////////////////////////////////////////////////////////////////////////////////
// type theTestGitController struct

func (inst *theTestGitController) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters
	inst.AppContext=inst.__get_AppContext__(injection, "context")
	inst.RM=inst.__get_RM__(injection, "#git-repository-manager")


	// to instance
	instance.AppContext=inst.AppContext
	instance.RM=inst.RM


	// invoke custom inject method


	return injection.Close()
}

func (inst * theTestGitController) __get_AppContext__(injection application.Injection,selector string) application.Context {
	return injection.Context()
}

func (inst * theTestGitController) __get_RM__(injection application.Injection,selector string) repository_d59845.Manager {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(repository_d59845.Manager)
	if !ok {
		err := errors.New("cannot cast component instance to type: repository_d59845.Manager")
		injection.OnError(err)
		return nil
	}

	return o2

}

////////////////////////////////////////////////////////////////////////////////
// type theUserConfigDAO struct

func (inst *theUserConfigDAO) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters


	// to instance


	// invoke custom inject method


	return injection.Close()
}

