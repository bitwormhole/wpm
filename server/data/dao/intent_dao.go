package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

type IntentDAO interface {
	Find(id dxo.IntentID) (*entity.Intent, error)

	ListAll() ([]*entity.Intent, error)

	Insert(o *entity.Intent) (*entity.Intent, error)
	Update(id dxo.IntentID, o *entity.Intent) (*entity.Intent, error)
	Remove(id dxo.IntentID) error
}