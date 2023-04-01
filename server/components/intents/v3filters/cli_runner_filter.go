package v3filters

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/components/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// CLIRunnerFilter ...
type CLIRunnerFilter struct {
	markup.Component `class:"wpm-intent-filter"`
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
	return next.Handle(c, i)
}
