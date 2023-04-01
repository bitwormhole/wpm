package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// NamespaceDAO ...
type NamespaceDAO interface {
	Find(id dxo.NamespaceID) (*entity.Namespace, error)

	ListAll() ([]*entity.Namespace, error)

	Insert(o *entity.Namespace) (*entity.Namespace, error)

	Update(id dxo.NamespaceID, o *entity.Namespace) (*entity.Namespace, error)

	Remove(id dxo.NamespaceID) error
}
