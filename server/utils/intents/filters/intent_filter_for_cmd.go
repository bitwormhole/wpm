package filters

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/utils/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentFilterForCmd ...
type IntentFilterForCmd struct {
	markup.Component `class:"intent-filter-registry"`
}

func (inst *IntentFilterForCmd) _Impl() (intents.Filter, intents.FilterRegistry) {
	return inst, inst
}

// GetRegistrationList ...
func (inst *IntentFilterForCmd) GetRegistrationList() []*intents.FilterRegistration {
	reg := &intents.FilterRegistration{
		Filter: inst,
	}
	return []*intents.FilterRegistration{reg}
}

// Filter ...
func (inst *IntentFilterForCmd) Filter(i *dto.Intent) (*dto.Intent, error) {
	if !inst.hit(i) {
		return i, nil // bypass
	}
	i.Command = "conhost"
	i.Arguments = []string{"cmd"}
	return i, nil
}

func (inst *IntentFilterForCmd) hit(i *dto.Intent) bool {
	return hit(i, []string{"cmd", "cmd.exe"})
}
