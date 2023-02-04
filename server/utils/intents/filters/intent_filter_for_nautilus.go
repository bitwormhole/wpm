package filters

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/utils/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentFilterForNautilus ...
type IntentFilterForNautilus struct {
	markup.Component `class:"intent-filter-registry"`
}

func (inst *IntentFilterForNautilus) _Impl() (intents.Filter, intents.FilterRegistry) {
	return inst, inst
}

// GetRegistrationList ...
func (inst *IntentFilterForNautilus) GetRegistrationList() []*intents.FilterRegistration {
	reg := &intents.FilterRegistration{
		Filter: inst,
	}
	return []*intents.FilterRegistration{reg}
}

// Filter ...
func (inst *IntentFilterForNautilus) Filter(i *dto.Intent) (*dto.Intent, error) {
	if !inst.hit(i) {
		return i, nil // bypass
	}
	wd := i.WD
	i.Arguments = []string{wd}
	return i, nil
}

func (inst *IntentFilterForNautilus) hit(i *dto.Intent) bool {
	return hit(i, []string{"nautilus"})
}
