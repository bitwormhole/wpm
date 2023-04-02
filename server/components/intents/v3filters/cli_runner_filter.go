package v3filters

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/components/intents"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// CLIRunnerFilter ...
type CLIRunnerFilter struct {
	markup.Component `class:"wpm-intent-filter"`

	IntentHandlerService service.IntentHandlerService `inject:"#IntentHandlerService"`
}

func (inst *CLIRunnerFilter) _Impl() (intents.FilterRegistry, intents.Filter) {
	return inst, inst
}

// GetFilterRegistrationList ...
func (inst *CLIRunnerFilter) GetFilterRegistrationList() []*intents.FilterRegistration {
	fr := &intents.FilterRegistration{
		Enabled: true,
		Order:   orderCLIRunner,
		Filter:  inst,
	}
	return []*intents.FilterRegistration{fr}
}

// Handle ...
func (inst *CLIRunnerFilter) Handle(c context.Context, i *dto.Intent, next intents.FilterChain) error {

	err := inst.IntentHandlerService.HandleIntent(i)
	if err != nil {
		return err
	}
	return next.Handle(c, i)
}
