package intents

// // ObjectIntentRunner ...
// type ObjectIntentRunner struct {
// 	GitLibAgent            store.LibAgent
// 	IntentHandlerService   service.IntentHandlerService
// 	LocalRepositoryService service.LocalRepositoryService
// 	ExecutableService      service.ExecutableService
// 	IntentFilters          FilterManager
// }

// // Run ...
// func (inst *ObjectIntentRunner) Run(c context.Context, intent *dto.Intent) error {

// 	err := inst.loadExecutable(c, intent)
// 	if err != nil {
// 		return err
// 	}

// 	err = inst.loadRepository(c, intent)
// 	if err != nil {
// 		return err
// 	}

// 	exe := intent.Executable
// 	repo := intent.Repository
// 	path, err := inst.getPathOfRepo(repo)
// 	if err != nil {
// 		return err
// 	}

// 	intent2 := &dto.Intent{}
// 	intent2.Exe = &dto.IntentExecutable{
// 		Repository: repo,
// 		Executable: exe,
// 	}
// 	intent2.CLI = &dto.IntentCLI{
// 		Command: exe.Path,
// 		WD:      path,
// 	}

// 	intent3, err := inst.IntentFilters.Filter(intent2)
// 	if err != nil {
// 		return err
// 	} else if intent3 == nil {
// 		return nil
// 	}

// 	return inst.IntentHandlerService.HandleIntent(intent3)
// }

// func (inst *ObjectIntentRunner) loadExecutable(c context.Context, intent *dto.Intent) error {
// 	ex := intent.Exe
// 	exeid := ex.Executable.ID
// 	ser := inst.ExecutableService
// 	o, err := ser.Find(c, exeid)
// 	if err != nil {
// 		return err
// 	}
// 	ex.Executable = o
// 	return nil
// }

// func (inst *ObjectIntentRunner) loadRepository(c context.Context, intent *dto.Intent) error {
// 	ex := intent.Exe
// 	repoid := ex.Repository.ID
// 	ser := inst.LocalRepositoryService
// 	o, err := ser.Find(c, repoid)
// 	if err != nil {
// 		return err
// 	}
// 	ex.Repository = o
// 	return nil
// }

// func (inst *ObjectIntentRunner) getPathOfRepo(repo *dto.LocalRepository) (string, error) {

// 	lib, err := inst.GitLibAgent.GetLib()
// 	if err != nil {
// 		return "", err
// 	}

// 	path := lib.FS().NewPath(repo.Path)
// 	layout, err := lib.RepositoryLocator().Locate(path)
// 	if err != nil {
// 		return "", err
// 	}

// 	dir1 := layout.Workspace()
// 	dir2 := layout.DotGit()
// 	dir3 := layout.Repository()
// 	dirs := []afs.Path{dir1, dir2, dir3}

// 	for _, dir := range dirs {
// 		if dir == nil {
// 			continue
// 		}
// 		if !dir.IsDirectory() {
// 			continue
// 		}
// 		return dir.GetPath(), nil
// 	}

// 	return "", fmt.Errorf("no dir in the repository")
// }
