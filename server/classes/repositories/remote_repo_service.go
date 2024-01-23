package repositories

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
)

// RemoteRepositoryService 表示该类型的服务接口
type RemoteRepositoryService interface {
	Insert(ctx context.Context, o *dto.RemoteRepository) (*dto.RemoteRepository, error)
	Update(ctx context.Context, id dxo.RemoteRepositoryID, o *dto.RemoteRepository) (*dto.RemoteRepository, error)

	Remove(ctx context.Context, id dxo.RemoteRepositoryID) error

	Find(ctx context.Context, id dxo.RemoteRepositoryID) (*dto.RemoteRepository, error)
	List(ctx context.Context, q *Query) ([]*dto.RemoteRepository, error)
}
