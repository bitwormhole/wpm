package v3filters

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/components/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// CLIMakerFilter ...
type CLIMakerFilter struct {
	markup.Component `class:"wpm-intent-filter"`
}

func (inst *CLIMakerFilter) _Impl() (intents.FilterRegistry, intents.Filter) {
	return inst, inst
}

// GetFilterRegistrationList ...
func (inst *CLIMakerFilter) GetFilterRegistrationList() []*intents.FilterRegistration {
	fr := &intents.FilterRegistration{
		Enabled: true,
		Order:   orderCLIMaker,
		Filter:  inst,
	}
	return []*intents.FilterRegistration{fr}
}

// Handle ...
func (inst *CLIMakerFilter) Handle(c context.Context, i *dto.Intent, next intents.FilterChain) error {
	return next.Handle(c, i)
}
