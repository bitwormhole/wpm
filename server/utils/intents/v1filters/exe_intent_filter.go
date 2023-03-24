package filters

import (
	"context"

	"bitwormhole.com/starter/afs/files"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/utils/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ExecuteIntentFilter ...
type ExecuteIntentFilter struct {
	markup.Component `class:"intent-filter-registry"`

	GitLibAgent store.LibAgent `inject:"#git-lib-agent"`
}

func (inst *ExecuteIntentFilter) _Impl() (intents.Filter, intents.FilterRegistry) {
	return inst, inst
}

// GetRegistrationList ...
func (inst *ExecuteIntentFilter) GetRegistrationList() []*intents.FilterRegistration {
	reg := &intents.FilterRegistration{
		Filter: inst,
		Order:  -1,
	}
	return []*intents.FilterRegistration{reg}
}

// Filter ...
func (inst *ExecuteIntentFilter) Filter(c context.Context, i *dto.Intent) (*dto.Intent, error) {

	err := inst.processCommand(i)
	if err != nil {
		return nil, err
	}

	err = inst.processWD(i)
	if err != nil {
		return nil, err
	}

	err = inst.processWDwithURL(i)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (inst *ExecuteIntentFilter) processCommand(i *dto.Intent) error {

	cmd := i.Command
	exe := i.Executable

	if cmd != "" {
		return nil
	}
	if exe == nil {
		return nil
	}

	cmd = exe.Path

	i.Command = cmd
	return nil
}

func (inst *ExecuteIntentFilter) processWD(i *dto.Intent) error {

	wd := i.WD
	repo := i.Repository

	if wd != "" {
		return nil
	}
	if repo == nil {
		return nil
	}

	lib, err := inst.GitLibAgent.GetLib()
	if err != nil {
		return err
	}

	repoPath := lib.FS().NewPath(repo.Path)
	layout, err := lib.RepositoryLocator().Locate(repoPath)
	if err != nil {
		return err
	}

	dir := layout.Workspace()
	if dir != nil {
		i.WD = dir.GetPath()
		return nil
	}

	dir = layout.Repository()
	if dir != nil {
		i.WD = dir.GetPath()
		return nil
	}

	return nil
}

func (inst *ExecuteIntentFilter) processWDwithURL(i *dto.Intent) error {

	wd := i.WD
	url := i.URL

	if url != "" {
		path := files.FS().NewPath(wd + "/" + url)
		if path.IsDirectory() {
			wd = path.GetPath()
		}
	}

	i.WD = wd
	return nil
}
