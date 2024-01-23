package repositories

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
)

// LocalRepositoryService 表示该类型的服务接口
type LocalRepositoryService interface {
	Insert(ctx context.Context, o *dto.LocalRepository) (*dto.LocalRepository, error)
	Update(ctx context.Context, id dxo.LocalRepositoryID, o *dto.LocalRepository) (*dto.LocalRepository, error)

	Remove(ctx context.Context, id dxo.LocalRepositoryID) error

	Find(ctx context.Context, id dxo.LocalRepositoryID) (*dto.LocalRepository, error)
	List(ctx context.Context, q *Query) ([]*dto.LocalRepository, error)
}
