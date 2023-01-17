package implservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// RunIntentServiceImpl ...
type RunIntentServiceImpl struct {
	markup.Component `id:"IntentService"`

	ExecutableService service.ExecutableService `inject:"#ExecutableService"`

	IntentHandlerService service.IntentHandlerService `inject:"#IntentHandlerService"`
}

func (inst *RunIntentServiceImpl) _Impl() service.IntentService {
	return inst
}

// Run ...
func (inst *RunIntentServiceImpl) Run(ctx context.Context, o *dto.Intent) (*dto.Intent, error) {

	cli := o.CLI
	exe := o.Exe
	web := o.Web

	if cli != nil {
		return inst.runAsCLI(ctx, o)
	} else if exe != nil {
		return inst.runAsExe(ctx, o)
	} else if web != nil {
		return inst.runAsWeb(ctx, o)
	}

	return nil, errors.New("bad request body data")
}

func (inst *RunIntentServiceImpl) runAsCLI(ctx context.Context, intent1 *dto.Intent) (*dto.Intent, error) {

	err := inst.IntentHandlerService.HandleIntent(intent1)
	if err != nil {
		return nil, err
	}
	return intent1, nil
}

func (inst *RunIntentServiceImpl) runAsWeb(ctx context.Context, o *dto.Intent) (*dto.Intent, error) {

	return nil, fmt.Errorf("no impl: RunIntentServiceImpl.runAsWeb")
}

func (inst *RunIntentServiceImpl) runAsExe(ctx context.Context, intent1 *dto.Intent) (*dto.Intent, error) {

	exe1 := intent1.Exe
	exe2, err := inst.ExecutableService.Find(ctx, exe1.Executable)
	if err != nil {
		return nil, err
	}

	intent1.CLI = &dto.IntentCLI{
		Command: exe2.Path,
	}

	err = inst.IntentHandlerService.HandleIntent(intent1)
	if err != nil {
		return nil, err
	}

	return intent1, nil
}

////////////////////////////////////////////////////////////////////////////////
