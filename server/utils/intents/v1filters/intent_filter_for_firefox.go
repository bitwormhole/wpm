package filters

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/utils/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentFilterForFirefox ...
type IntentFilterForFirefox struct {
	markup.Component `class:"intent-filter-registry"`
}

func (inst *IntentFilterForFirefox) _Impl() (intents.Filter, intents.FilterRegistry) {
	return inst, inst
}

// GetRegistrationList ...
func (inst *IntentFilterForFirefox) GetRegistrationList() []*intents.FilterRegistration {
	reg := &intents.FilterRegistration{
		Filter: inst,
	}
	return []*intents.FilterRegistration{reg}
}

// Filter ...
func (inst *IntentFilterForFirefox) Filter(c context.Context, i *dto.Intent) (*dto.Intent, error) {
	if !inst.hit(i) {
		return i, nil // bypass
	}
	// todo ...
	return i, nil
}

func (inst *IntentFilterForFirefox) hit(i *dto.Intent) bool {
	return hit(i, []string{"a", "b"})
}
