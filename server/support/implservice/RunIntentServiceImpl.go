package implservice

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// RunIntentServiceImpl ...
type RunIntentServiceImpl struct {
	markup.Component `id:"IntentService"`

	GitLibAgent store.LibAgent `inject:"#git-lib-agent"`

	IntentFilterManager intents.FilterManager `inject:"#intent-filter-manager"`

	LocalRepositoryService service.LocalRepositoryService `inject:"#LocalRepositoryService"`
	ExecutableService      service.ExecutableService      `inject:"#ExecutableService"`
	IntentHandlerService   service.IntentHandlerService   `inject:"#IntentHandlerService"`
}

func (inst *RunIntentServiceImpl) _Impl() service.IntentService {
	return inst
}

// Run ...
func (inst *RunIntentServiceImpl) Run(ctx context.Context, o *dto.Intent) (*dto.Intent, error) {
	o, err := inst.preprocessWithFilters(ctx, o)
	if err != nil {
		return nil, err
	}
	return inst.runAsCLI(ctx, o)
}

func (inst *RunIntentServiceImpl) preprocessWithFilters(ctx context.Context, o1 *dto.Intent) (*dto.Intent, error) {
	return inst.IntentFilterManager.Filter(o1)
}

func (inst *RunIntentServiceImpl) runAsCLI(ctx context.Context, intent1 *dto.Intent) (*dto.Intent, error) {

	err := inst.IntentHandlerService.HandleIntent(intent1)
	if err != nil {
		return nil, err
	}
	return intent1, nil
}

// func (inst *RunIntentServiceImpl) runAsWeb(ctx context.Context, o *dto.Intent) (*dto.Intent, error) {
// 	return nil, fmt.Errorf("no impl: RunIntentServiceImpl.runAsWeb")
// }

// func (inst *RunIntentServiceImpl) runAsExe(ctx context.Context, intent *dto.Intent) (*dto.Intent, error) {
// 	runner := &intents.ObjectIntentRunner{
// 		GitLibAgent:            inst.GitLibAgent,
// 		IntentHandlerService:   inst.IntentHandlerService,
// 		ExecutableService:      inst.ExecutableService,
// 		LocalRepositoryService: inst.LocalRepositoryService,
// 		IntentFilters:          inst.IntentFilterManager,
// 	}
// 	err := runner.Run(ctx, intent)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return intent, nil
// }

////////////////////////////////////////////////////////////////////////////////
