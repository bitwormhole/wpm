package init

import (
	"context"

	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
)

type myInitGetHandler struct {
	// serv            service.InitService
	ExecutableSer  service.ExecutableService
	ProjectTypeSer service.ContentTypeService
	CheckUpdateSer service.CheckUpdateService
	SetupSer       service.SetupService
}

func (inst *myInitGetHandler) handle(ctx context.Context) (*vo.Init, error) {
	view := &vo.Init{}
	steps := make([]func(ctx context.Context, view *vo.Init) error, 0)

	steps = append(steps, inst.doSetup)
	steps = append(steps, inst.doChcekUpdate)
	steps = append(steps, inst.doExecutable)
	steps = append(steps, inst.doProjectType)

	for _, step := range steps {
		err := step(ctx, view)
		if err != nil {
			return nil, err
		}
	}
	return view, nil
}

func (inst *myInitGetHandler) doChcekUpdate(ctx context.Context, view *vo.Init) error {
	view2 := &vo.AboutCheckUpdate{}
	err := inst.CheckUpdateSer.Check(ctx, view2)
	if err != nil {
		return err
	}
	view.CheckUpdate.CheckUpdate = view2
	return nil
}

func (inst *myInitGetHandler) doSetup(ctx context.Context, view *vo.Init) error {
	req, err := inst.SetupSer.IsSetupReqiured(ctx)
	if err != nil {
		return err
	}
	view.Setup.SetupRequired = req
	return nil
}

func (inst *myInitGetHandler) doProjectType(ctx context.Context, view *vo.Init) error {
	list, err := inst.ProjectTypeSer.ListAll(ctx)
	if err != nil {
		return err
	}
	view.ContentType.ContentTypes = list
	return nil
}

func (inst *myInitGetHandler) doExecutable(ctx context.Context, view *vo.Init) error {
	list, err := inst.ExecutableSer.ListAll(ctx)
	if err != nil {
		return err
	}
	view.Executable.Executables = list
	return nil
}
