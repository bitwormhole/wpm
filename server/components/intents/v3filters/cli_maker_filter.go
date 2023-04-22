package v3filters

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/components/intents"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// CLIMakerFilter ...
type CLIMakerFilter struct {
	markup.Component `class:"wpm-intent-filter"`

	FileSystemService service.FileSystemService `inject:"#FileSystemService"`
}

func (inst *CLIMakerFilter) _Impl() (intents.FilterRegistry, intents.Filter) {
	return inst, inst
}

// GetFilterRegistrationList ...
func (inst *CLIMakerFilter) GetFilterRegistrationList() []*intents.FilterRegistration {
	fr := &intents.FilterRegistration{
		Enabled: true,
		Order:   orderCLIMaker,
		Filter:  inst,
	}
	return []*intents.FilterRegistration{fr}
}

// Handle ...
func (inst *CLIMakerFilter) Handle(c context.Context, i *dto.Intent, next intents.FilterChain) error {

	cli := i.CLI
	if cli != nil {
		return next.Handle(c, i)
	}
	cli = &dto.CommandRequest{}

	temp := i.Templates[0]
	cli.Command = temp.Command
	cli.Arguments = temp.Arguments.Array()
	cli.Env = temp.Env.Map()
	cli.WD = temp.WD

	mr := intents.MacroResolver{}
	mr.Init(i.Properties)
	cli.Command = mr.ResolveString(cli.Command)
	cli.Arguments = mr.ResolveStringArray(cli.Arguments)
	cli.Env = mr.ResolveStringMap(cli.Env)
	cli.WD = mr.ResolveString(cli.WD)
	err := mr.Error()
	if err != nil {
		return err
	}

	wd, err := inst.normalizeWorkingDir(cli.WD)
	if err != nil {
		return err
	}
	cli.WD = wd

	i.CLI = cli
	return next.Handle(c, i)
}

func (inst *CLIMakerFilter) normalizeWorkingDir(path string) (string, error) {
	dir := inst.FileSystemService.Path(path)
	if !dir.IsDirectory() {
		return "", fmt.Errorf("no directory at path [%v]", path)
	}
	return dir.GetPath(), nil
}

////////////////////////////////////////////////////////////////////////////////
