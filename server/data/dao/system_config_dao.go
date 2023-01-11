package dao

import (
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/wpm/server/data/entity"
)

type SystemConfig interface {
	Get(name string) (*entity.SystemConfig, error)
	Save(name string, cfg *entity.SystemConfig) error

	GetProperties(name string) (collection.Properties, error)
	SaveProperties(name string, cfg collection.Properties) error
}
