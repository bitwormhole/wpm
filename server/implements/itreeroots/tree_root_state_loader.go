package itreeroots

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/treeroots"
	"github.com/starter-go/afs"
)

// TreeRootStateLoaderImpl ...
type TreeRootStateLoaderImpl struct {

	//starter:component
	_as func(treeroots.StateLoader) //starter:as("#")

	FS afs.FS //starter:inject("#")

	// Dao         treeroots.DAO         //starter:inject("#")
	// StateLoader treeroots.StateLoader //starter:inject("#")
	// convertor   treeroots.Convertor
}

func (inst *TreeRootStateLoaderImpl) _impl() treeroots.StateLoader {
	return inst
}

// Load  ...
func (inst *TreeRootStateLoaderImpl) Load(ctx context.Context, id dxo.TreeRootID, o *dto.TreeRoot) error {
	return nil
}
