package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// ContentTypeDAO ...
type ContentTypeDAO interface {
	Find(id dxo.ContentTypeID) (*entity.ContentType, error)

	FindByURN(urn dxo.ContentTypeURN) (*entity.ContentType, error)

	FindByName(name dxo.ContentTypeName) (*entity.ContentType, error)

	ListAll() ([]*entity.ContentType, error)

	ListByPattern(pattern string) ([]*entity.ContentType, error)

	Insert(o *entity.ContentType) (*entity.ContentType, error)

	Update(id dxo.ContentTypeID, o *entity.ContentType) (*entity.ContentType, error)

	Remove(id dxo.ContentTypeID) error
}
