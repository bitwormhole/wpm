package filters

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/utils/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentFilterForMsEdge ...
type IntentFilterForMsEdge struct {
	markup.Component `class:"intent-filter-registry"`
}

func (inst *IntentFilterForMsEdge) _Impl() (intents.Filter, intents.FilterRegistry) {
	return inst, inst
}

// GetRegistrationList ...
func (inst *IntentFilterForMsEdge) GetRegistrationList() []*intents.FilterRegistration {
	reg := &intents.FilterRegistration{
		Filter: inst,
	}
	return []*intents.FilterRegistration{reg}
}

// Filter ...
func (inst *IntentFilterForMsEdge) Filter(i *dto.Intent) (*dto.Intent, error) {
	if !inst.hit(i) {
		return i, nil // bypass
	}
	// todo ...
	return i, nil
}

func (inst *IntentFilterForMsEdge) hit(i *dto.Intent) bool {
	return hit(i, []string{"a", "b"})
}
