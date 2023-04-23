package v3filters

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/components/intents"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// PrepareActionFilter ...
type PrepareActionFilter struct {
	markup.Component `class:"wpm-intent-filter"`

	ProfileService service.ProfileService `inject:"#ProfileService"`
}

func (inst *PrepareActionFilter) _Impl() (intents.FilterRegistry, intents.Filter) {
	return inst, inst
}

// GetFilterRegistrationList ...
func (inst *PrepareActionFilter) GetFilterRegistrationList() []*intents.FilterRegistration {
	fr := &intents.FilterRegistration{
		Enabled: true,
		Order:   orderPrepareAction,
		Filter:  inst,
	}
	return []*intents.FilterRegistration{fr}
}

// Handle ...
func (inst *PrepareActionFilter) Handle(c context.Context, i *dto.Intent, next intents.FilterChain) error {

	action := &i.Action
	action.Method = inst.prepareMethod(i)
	action.Target = inst.prepareTarget(i)
	action.Type = inst.prepareContentType(i)
	action.With = inst.prepareExe(i)

	if action.Target == "none" {
		inst.prepareTargetWithWD(i)
	}

	itsb := dxo.IntentTemplateSelectorBuilder{
		Method: action.Method,
		With:   action.With.String(),
		Type:   action.Type,
		Target: action.Target,
	}
	action.Selector = itsb.Create()

	return next.Handle(c, i)
}

func (inst *PrepareActionFilter) prepareMethod(i *dto.Intent) string {
	m := i.Action.Method
	if m == "" {
		m = "open"
	}
	return m
}

func (inst *PrepareActionFilter) prepareContentType(i *dto.Intent) string {
	ct := i.Action.Type
	if ct == "" {
		ct = "*"
	}
	return ct
}

func (inst *PrepareActionFilter) prepareExe(i *dto.Intent) dxo.ExecutableURN {
	exe := i.Executable
	if exe == nil {
		return "*"
	}
	return exe.URN
}

func (inst *PrepareActionFilter) prepareTarget(i *dto.Intent) string {
	table := make(map[string]bool)
	table["file"] = i.File == nil
	table["folder"] = i.Folder == nil
	table["repository"] = i.Repository == nil
	table["worktree"] = i.Worktree == nil
	table["submodule"] = i.Submodule == nil
	table["project"] = i.Project == nil
	for name, isnil := range table {
		if !isnil {
			return name
		}
	}
	return "none"
}

func (inst *PrepareActionFilter) prepareTargetWithWD(i *dto.Intent) error {

	p, err := inst.ProfileService.GetProfile()
	if err != nil {
		return err
	}
	home := p.Home

	i.Action.Target = "folder"
	i.Action.Type = "*"
	i.Folder = &dto.File{
		Path: home,
	}

	return nil
}
