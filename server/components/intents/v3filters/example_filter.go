package v3filters

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/components/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ExampleFilter ...
type ExampleFilter struct {
	markup.Component `class:"wpm-intent-filter"`
}

func (inst *ExampleFilter) _Impl() (intents.FilterRegistry, intents.Filter) {
	return inst, inst
}

// GetFilterRegistrationList ...
func (inst *ExampleFilter) GetFilterRegistrationList() []*intents.FilterRegistration {
	fr := &intents.FilterRegistration{
		Enabled: true,
		Order:   order0,
		Filter:  inst,
	}
	return []*intents.FilterRegistration{fr}
}

// Handle ...
func (inst *ExampleFilter) Handle(c context.Context, i *dto.Intent, next intents.FilterChain) error {
	return next.Handle(c, i)
}
