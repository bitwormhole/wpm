package contenttypes

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
)

// Service 表示该类型的服务接口
type Service interface {
	Insert(ctx context.Context, o *dto.ContentType) (*dto.ContentType, error)
	Update(ctx context.Context, id dxo.ContentTypeID, o *dto.ContentType) (*dto.ContentType, error)

	Remove(ctx context.Context, id dxo.ContentTypeID) error

	Find(ctx context.Context, id dxo.ContentTypeID) (*dto.ContentType, error)
	List(ctx context.Context, q *Query) ([]*dto.ContentType, error)
}
