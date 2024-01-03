package media

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// Service 表示该类型的服务接口
type Service interface {
	Insert(ctx context.Context, o *dto.Media) (*dto.Media, error)
	Update(ctx context.Context, id dxo.MediaID, o *dto.Media) (*dto.Media, error)

	Remove(ctx context.Context, id dxo.MediaID) error

	Find(ctx context.Context, id dxo.MediaID) (*dto.Media, error)
	List(ctx context.Context, q *Query) ([]*dto.Media, error)
}
