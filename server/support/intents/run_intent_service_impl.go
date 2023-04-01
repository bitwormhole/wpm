package intents

import (
	"context"
	"net/http"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/components/intents"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// RunIntentServiceImpl ...
type RunIntentServiceImpl struct {
	markup.Component `id:"IntentService"`

	GitLibAgent store.LibAgent `inject:"#git-lib-agent"`

	IntentFilterManager intents.FilterManager `inject:"#wpm-intent-filter-manager"`

	LocalRepositoryService service.LocalRepositoryService `inject:"#LocalRepositoryService"`
	ExecutableService      service.ExecutableService      `inject:"#ExecutableService"`
	IntentHandlerService   service.IntentHandlerService   `inject:"#IntentHandlerService"`
}

func (inst *RunIntentServiceImpl) _Impl() service.IntentService {
	return inst
}

// Run ...
func (inst *RunIntentServiceImpl) Run(ctx context.Context, o *dto.Intent) (*dto.Intent, error) {
	ch := inst.IntentFilterManager.Chain()
	err := ch.Handle(ctx, o)
	if err != nil {
		return nil, err
	}
	if o.Status == 0 {
		o.Status = http.StatusOK
	}
	return o, nil
}

// func (inst *RunIntentServiceImpl) isReadyForRun(ctx context.Context, o1 *dto.Intent) bool {
// 	if o1.CLI == nil {
// 		return false
// 	}
// 	return true
// }

// func (inst *RunIntentServiceImpl) run(ctx context.Context, intent1 *dto.Intent) (*dto.Intent, error) {

// 	ready := inst.isReadyForRun(ctx, intent1)
// 	if !ready {
// 		e := intent1.Error
// 		if e == "" {
// 			intent1.Status = 0
// 			return intent1, nil
// 		}
// 		return nil, fmt.Errorf(e)
// 	}

// 	err := inst.IntentHandlerService.HandleIntent(intent1)
// 	if err != nil {
// 		return nil, err
// 	}

// 	intent1.Status = http.StatusOK
// 	return intent1, nil
// }

////////////////////////////////////////////////////////////////////////////////
