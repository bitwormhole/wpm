package implservice

import (
	"context"
	"time"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/pipe"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// PipeServiceImpl ...
type PipeServiceImpl struct {
	markup.Component `id:"PipeService" initMethod:"Init"`

	pm pipe.Manager
}

func (inst *PipeServiceImpl) _Impl() service.PipeService {
	return inst
}

// Init ...
func (inst *PipeServiceImpl) Init() error {
	return inst.pm.Init()
}

// Insert ...
func (inst *PipeServiceImpl) Insert(ctx context.Context, o *dto.PipeInfo) (*dto.PipeInfo, error) {
	return inst.pm.Insert(o)
}

// Pull ...
func (inst *PipeServiceImpl) Pull(ctx context.Context, id dxo.PipeID, view *vo.Pipe, timeout int) error {
	timeout2 := time.Duration(timeout) * time.Millisecond
	return inst.pm.WaitAndPull(id, view, timeout2)
}

// Push ...
func (inst *PipeServiceImpl) Push(ctx context.Context, id dxo.PipeID, view *vo.Pipe) error {
	return inst.pm.Push(id, view)
}

// Remove ...
func (inst *PipeServiceImpl) Remove(ctx context.Context, id dxo.PipeID) error {
	return inst.pm.Remove(id)
}

// ListAll ...
func (inst *PipeServiceImpl) ListAll(ctx context.Context) ([]*dto.PipeInfo, error) {
	all := inst.pm.ListPipes()
	return all, nil
}
