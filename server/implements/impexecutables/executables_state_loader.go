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

	cache caches.Agent
}

func (inst *ExecutableStateLoaderImpl) _impl() executables.StateLoader {
	return inst
}

func (inst *ExecutableStateLoaderImpl) getCacheLoader(ctx context.Context) caches.Loader {
	loader := inst.cache.NewLoader(ctx)
	loader.Class("ExecutableStateLoaderImpl::dto.Executable")
	loader.Pool(inst.CacheService.GetPool())
	return loader
}

// Load ...
func (inst *ExecutableStateLoaderImpl) Load(ctx context.Context, id dxo.ExecutableID, o *dto.Executable) error {

	if ctx == nil || o == nil {
		return fmt.Errorf("param is nil")
	}

	loader := inst.getCacheLoader(ctx)
	loader.ID(o.UUID.String())
	loader.NotBefore(o.UpdatedAt)
	loader.OnLoad(func(want *caches.Want) (any, error) {
		return inst.reload(want, o)
	})

	x, err := loader.Load()
	if err != nil {
		return err
	}

	h := x.(*myStateItemHolder)
	o.State = h.state

	return nil
}

func (inst *ExecutableStateLoaderImpl) reload(want *caches.Want, obj *dto.Executable) (any, error) {

	// obj := want.GetAttr("exe").(*dto.Executable)
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
