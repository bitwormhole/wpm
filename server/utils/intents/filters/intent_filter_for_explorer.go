package filters

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/utils/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentFilterForExplorer ...
type IntentFilterForExplorer struct {
	markup.Component `class:"intent-filter-registry"`
}

func (inst *IntentFilterForExplorer) _Impl() (intents.Filter, intents.FilterRegistry) {
	return inst, inst
}

// GetRegistrationList ...
func (inst *IntentFilterForExplorer) GetRegistrationList() []*intents.FilterRegistration {
	reg := &intents.FilterRegistration{
		Filter: inst,
	}
	return []*intents.FilterRegistration{reg}
}

// Filter ...
func (inst *IntentFilterForExplorer) Filter(i *dto.Intent) (*dto.Intent, error) {
	if !inst.hit(i) {
		return i, nil // bypass
	}
	wd := i.WD
	i.Arguments = []string{wd}
	return i, nil
}

func (inst *IntentFilterForExplorer) hit(i *dto.Intent) bool {
	return hit(i, []string{"explorer", "explorer.exe"})
}
