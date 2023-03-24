package filters

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/utils/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentFilterForPowerShell ...
type IntentFilterForPowerShell struct {
	markup.Component `class:"intent-filter-registry"`
}

func (inst *IntentFilterForPowerShell) _Impl() (intents.Filter, intents.FilterRegistry) {
	return inst, inst
}

// GetRegistrationList ...
func (inst *IntentFilterForPowerShell) GetRegistrationList() []*intents.FilterRegistration {
	reg := &intents.FilterRegistration{
		Filter: inst,
	}
	return []*intents.FilterRegistration{reg}
}

// Filter ...
func (inst *IntentFilterForPowerShell) Filter(c context.Context, i *dto.Intent) (*dto.Intent, error) {
	if !inst.hit(i) {
		return i, nil // bypass
	}
	i.Command = "conhost"
	i.Arguments = []string{"powershell"}
	return i, nil
}

func (inst *IntentFilterForPowerShell) hit(i *dto.Intent) bool {
	return hit(i, []string{"powershell", "powershell.exe"})
}
