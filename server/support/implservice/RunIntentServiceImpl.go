package implservice

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// RunIntentServiceImpl ...
type RunIntentServiceImpl struct {
	markup.Component `id:"RunIntentService"`

	ExecutableService service.ExecutableService `inject:"#ExecutableService"`
	PipeService       service.PipeService       `inject:"#PipeService"`
}

func (inst *RunIntentServiceImpl) _Impl() service.RunIntentService {
	return inst
}

// Run ...
func (inst *RunIntentServiceImpl) Run(ctx context.Context, o *dto.Intent) (*dto.Intent, error) {

	cli := o.CLI
	exe := o.Exe
	web := o.Web

	if cli != nil {
		return inst.runAsCLI(ctx, o)
	} else if exe != nil {
		return inst.runAsExe(ctx, o)
	} else if web != nil {
		return inst.runAsWeb(ctx, o)
	}

	return nil, errors.New("bad request body data")
}

func (inst *RunIntentServiceImpl) runAsCLI(ctx context.Context, intent1 *dto.Intent) (*dto.Intent, error) {

	intent2 := &dto.Intent{}
	intent2.CLI = intent1.CLI
	packet := &dto.PipePacket{
		Intent: intent2,
	}
	pipe := &vo.Pipe{
		Packets: []*dto.PipePacket{packet},
	}

	pi := intent1.Pipe
	pi, err := inst.selectPipe(ctx, pi)
	if err != nil {
		return nil, err
	}

	err = inst.PipeService.Push(ctx, pi.ID, pipe)
	if err != nil {
		return nil, err
	}

	return intent2, nil
}

func (inst *RunIntentServiceImpl) runAsWeb(ctx context.Context, o *dto.Intent) (*dto.Intent, error) {

	return nil, fmt.Errorf("no impl: RunIntentServiceImpl.runAsWeb")
}

func (inst *RunIntentServiceImpl) runAsExe(ctx context.Context, intent1 *dto.Intent) (*dto.Intent, error) {

	exe1 := intent1.Exe
	exe2, err := inst.ExecutableService.Find(ctx, exe1.Executable)
	if err != nil {
		return nil, err
	}

	intent2 := &dto.Intent{}
	intent2.CLI = &dto.IntentCLI{
		Command: exe2.Path,
	}
	packet := &dto.PipePacket{
		Intent: intent2,
	}
	pipe := &vo.Pipe{
		Packets: []*dto.PipePacket{packet},
	}

	pi := intent1.Pipe
	pi, err = inst.selectPipe(ctx, pi)
	if err != nil {
		return nil, err
	}

	err = inst.PipeService.Push(ctx, pi.ID, pipe)
	if err != nil {
		return nil, err
	}

	return intent2, nil
}

func (inst *RunIntentServiceImpl) selectPipe(ctx context.Context, want *dto.PipeInfo) (*dto.PipeInfo, error) {

	if want == nil {
		return nil, fmt.Errorf("param is nil")
	}

	all, err := inst.PipeService.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	(&myRunIntentServicePipeSorter{items: all}).sort()

	for _, item := range all {
		if item == nil {
			continue
		}
		if item.ID == want.ID {
			return item, nil
		} else if item.UUID == want.UUID {
			return item, nil
		} else if item.Name == want.Name {
			return item, nil
		} else if item.DesktopSessionID == want.DesktopSessionID {
			return item, nil
		} else if item.DesktopSessionUser == want.DesktopSessionUser {
			return item, nil
		}
	}

	return nil, fmt.Errorf("no wanted pipe")
}

////////////////////////////////////////////////////////////////////////////////

type myRunIntentServicePipeSorter struct {
	items []*dto.PipeInfo
}

func (inst *myRunIntentServicePipeSorter) sort() {
	sort.Sort(inst)
}

func (inst *myRunIntentServicePipeSorter) Len() int {
	return len(inst.items)
}

func (inst *myRunIntentServicePipeSorter) Less(i1, i2 int) bool {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	return o1.AttachedAt > o2.AttachedAt
}

func (inst *myRunIntentServicePipeSorter) Swap(i1, i2 int) {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	inst.items[i1] = o2
	inst.items[i2] = o1
}

////////////////////////////////////////////////////////////////////////////////
