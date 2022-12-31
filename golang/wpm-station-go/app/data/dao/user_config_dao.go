package dao

import (
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/wpm/app/data/entity"
)

type UserConfig interface {
	Get(name string) (*entity.UserConfig, error)
	Save(name string, cfg *entity.UserConfig) error

	GetProperties(name string) (collection.Properties, error)
	SaveProperties(name string, cfg collection.Properties) error
}
