package impexecutables

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/common/classes/caches"
	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/executables"
	"github.com/starter-go/afs"
	"github.com/starter-go/vlog"
)

type myStateItemHolder struct {
	item  caches.Item
	state dxo.FileState
}

////////////////////////////////////////////////////////////////////////////////

// ExecutableStateLoaderImpl ...
type ExecutableStateLoaderImpl struct {

	//starter:component
	_as func(executables.StateLoader) //starter:as("#")

	CacheService caches.Service //starter:inject("#")
	FS           afs.FS         //starter:inject("#")

	cache *myStateCache
}

func (inst *ExecutableStateLoaderImpl) _impl() executables.StateLoader {
	return inst
}

func (inst *ExecutableStateLoaderImpl) getCache() *myStateCache {
	c := inst.cache
	if c == nil {
		c = new(myStateCache)
		c.init(inst.CacheService, inst.reload)
		inst.cache = c
	}
	return c
}

// Load ...
func (inst *ExecutableStateLoaderImpl) Load(ctx context.Context, id dxo.ExecutableID, o *dto.Executable) error {

	if ctx == nil || o == nil {
		return fmt.Errorf("param is nil")
	}

	want := &caches.Want{
		ID: o.UUID.String(),
	}
	want.SetAttr("exe", o)
	want.NotBefore = o.UpdatedAt

	x, err := inst.getCache().load(want)
	if err != nil {
		return err
	}

	h := x.(*myStateItemHolder)
	o.State = h.state

	return nil
}

func (inst *ExecutableStateLoaderImpl) reload(want *caches.Want) (any, error) {

	obj := want.GetAttr("exe").(*dto.Executable)
	file := inst.FS.NewPath(obj.Path)

	vlog.Info("reload state of Executable:%d", obj.ID)
	h := new(myStateItemHolder)
	h.item = want.Item

	if file.IsFile() {
		h.state = dxo.FileStateReady
	} else {
		h.state = dxo.FileStateOffline
	}

	return h, nil
}
