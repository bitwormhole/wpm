package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

type ExecutableDAO interface {
	Find(id dxo.ExecutableID) (*entity.Executable, error)
	FindByPath(path string) (*entity.Executable, error)

	ListAll() ([]*entity.Executable, error)

	Insert(o *entity.Executable) (*entity.Executable, error)
	Update(id dxo.ExecutableID, o *entity.Executable) (*entity.Executable, error)
	Remove(id dxo.ExecutableID) error
}
