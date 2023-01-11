package convert

import (
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// LocalRepository ...
type LocalRepository interface {
	EntityToView(o1 *entity.LocalRepository) (*dto.LocalRepository, error)
	ViewToEntity(o1 *dto.LocalRepository) (*entity.LocalRepository, error)
}
