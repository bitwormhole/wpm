package filters

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/utils/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentFilterForChrome ...
type IntentFilterForChrome struct {
	markup.Component `class:"intent-filter-registry"`
}

func (inst *IntentFilterForChrome) _Impl() (intents.Filter, intents.FilterRegistry) {
	return inst, inst
}

// GetRegistrationList ...
func (inst *IntentFilterForChrome) GetRegistrationList() []*intents.FilterRegistration {
	reg := &intents.FilterRegistration{
		Filter: inst,
	}
	return []*intents.FilterRegistration{reg}
}

// Filter ...
func (inst *IntentFilterForChrome) Filter(c context.Context, i *dto.Intent) (*dto.Intent, error) {
	if !inst.hit(i) {
		return i, nil // bypass
	}
	// todo ...
	return i, nil
}

func (inst *IntentFilterForChrome) hit(i *dto.Intent) bool {
	return hit(i, []string{"a", "b"})
}
