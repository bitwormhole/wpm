package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// SettingDAO ...
type SettingDAO interface {
	Find(id dxo.SettingID) (*entity.Setting, error)

	FindByName(name string) (*entity.Setting, error)

	ListAll() ([]*entity.Setting, error)

	Insert(o *entity.Setting) (*entity.Setting, error)

	Update(id dxo.SettingID, o *entity.Setting) (*entity.Setting, error)

	Remove(id dxo.SettingID) error

	Put(name string, o *entity.Setting)

	Get(name string, def *entity.Setting) *entity.Setting
}
