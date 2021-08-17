package dao

import (
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/wpm/data/entity"
)

type SystemConfig interface {
	Get(name string) (*entity.SystemConfig, error)
	Save(name string, cfg *entity.SystemConfig) error

	GetProperties(name string) (collection.Properties, error)
	SaveProperties(name string, cfg collection.Properties) error
}

type SystemConfigImpl struct {
}

func (inst *SystemConfigImpl) _Impl() SystemConfig {
	return inst
}

func (inst *SystemConfigImpl) Get(name string) (*entity.SystemConfig, error) {
	return nil, nil
}

func (inst *SystemConfigImpl) Save(name string, cfg *entity.SystemConfig) error {
	return nil
}

func (inst *SystemConfigImpl) GetProperties(name string) (collection.Properties, error) {
	return nil, nil
}

func (inst *SystemConfigImpl) SaveProperties(name string, cfg collection.Properties) error {
	return nil
}
