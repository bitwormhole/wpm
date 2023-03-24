package v2filters

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/utils/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ExampleFilter ...
type ExampleFilter struct {
	markup.Component `class:"intent-filter-registry"`
}

func (inst *ExampleFilter) _Impl() intents.FilterRegistry {
	return inst
}

// GetRegistrationList ...
func (inst *ExampleFilter) GetRegistrationList() []*intents.FilterRegistration {
	reg := &intents.FilterRegistration{
		Enabled: false,
		Filter:  inst,
		Order:   0,
	}
	return []*intents.FilterRegistration{reg}
}

// Filter ...
func (inst *ExampleFilter) Filter(c context.Context, o *dto.Intent) (*dto.Intent, error) {
	return nil, fmt.Errorf("no impl")
}
