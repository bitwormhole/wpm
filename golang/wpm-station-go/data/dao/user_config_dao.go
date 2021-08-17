package dao

import (
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/wpm/data/entity"
)

type UserConfig interface {
	Get(name string) (*entity.UserConfig, error)
	Save(name string, cfg *entity.UserConfig) error

	GetProperties(name string) (collection.Properties, error)
	SaveProperties(name string, cfg collection.Properties) error
}

type UserConfigImpl struct {
}

func (inst *UserConfigImpl) _Impl() UserConfig {
	return inst
}

func (inst *UserConfigImpl) Get(name string) (*entity.UserConfig, error) {
	return nil, nil
}

func (inst *UserConfigImpl) Save(name string, cfg *entity.UserConfig) error {
	return nil
}

func (inst *UserConfigImpl) GetProperties(name string) (collection.Properties, error) {
	return nil, nil
}

func (inst *UserConfigImpl) SaveProperties(name string, cfg collection.Properties) error {
	return nil
}
