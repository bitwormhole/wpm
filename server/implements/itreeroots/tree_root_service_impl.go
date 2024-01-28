package itreeroots

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/treeroots"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/starter-go/afs"
)

// TreeRootServiceImpl ...
type TreeRootServiceImpl struct {

	//starter:component
	_as func(treeroots.Service) //starter:as("#")

	Dao         treeroots.DAO         //starter:inject("#")
	FS          afs.FS                //starter:inject("#")
	StateLoader treeroots.StateLoader //starter:inject("#")

	convertor treeroots.Convertor
}

func (inst *TreeRootServiceImpl) _impl() treeroots.Service {
	return inst
}

// Insert ...
func (inst *TreeRootServiceImpl) Insert(ctx context.Context, o1 *dto.TreeRoot) (*dto.TreeRoot, error) {

	// o1.URN = inst.computeURN(o1)
	// inst.loadExeMeta(o1)

	o2 := inst.convertor.ConvertD2E(o1)
	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}
	o4 := inst.convertor.ConvertE2D(o3)
	return o4, nil
}

// Update ...
func (inst *TreeRootServiceImpl) Update(ctx context.Context, id dxo.TreeRootID, o1 *dto.TreeRoot) (*dto.TreeRoot, error) {

	o2, err := inst.Dao.Update(nil, id, func(ent *entity.TreeRoot) {

		ent.Path = o1.Path
		ent.DotGitPath = o1.DotGitPath
		ent.Class = o1.Class

	})
	if err != nil {
		return nil, err
	}
	o3 := inst.convertor.ConvertE2D(o2)
	return o3, nil
}

// Remove ...
func (inst *TreeRootServiceImpl) Remove(ctx context.Context, id dxo.TreeRootID) error {
	return inst.Dao.Remove(nil, id)
}

// Find ...
func (inst *TreeRootServiceImpl) Find(ctx context.Context, id dxo.TreeRootID) (*dto.TreeRoot, error) {
	o1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}

	o2 := inst.convertor.ConvertE2D(o1)
	err = inst.StateLoader.Load(ctx, id, o2)
	if err != nil {
		return nil, err
	}

	return o2, nil
}

// List ...
func (inst *TreeRootServiceImpl) List(ctx context.Context, q *treeroots.Query) ([]*dto.TreeRoot, error) {
	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}
	list2 := inst.convertor.ConvertListE2D(list1)

	for _, item := range list2 {
		err = inst.StateLoader.Load(ctx, item.ID, item)
		if err != nil {
			return nil, err
		}
	}

	return list2, nil
}
