package filters

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/utils/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentFilterForVscode ...
type IntentFilterForVscode struct {
	markup.Component `class:"intent-filter-registry"`
}

func (inst *IntentFilterForVscode) _Impl() (intents.Filter, intents.FilterRegistry) {
	return inst, inst
}

// GetRegistrationList ...
func (inst *IntentFilterForVscode) GetRegistrationList() []*intents.FilterRegistration {
	reg := &intents.FilterRegistration{
		Filter: inst,
	}
	return []*intents.FilterRegistration{reg}
}

// Filter ...
func (inst *IntentFilterForVscode) Filter(c context.Context, i *dto.Intent) (*dto.Intent, error) {
	if !inst.hit(i) {
		return i, nil // bypass
	}
	wd := i.WD
	i.Arguments = []string{wd}
	return i, nil
}

func (inst *IntentFilterForVscode) hit(i *dto.Intent) bool {
	return hit(i, []string{"code", "code.exe", "code.cmd"})
}
