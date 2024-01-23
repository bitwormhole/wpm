package locations

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
)

// Service 表示该类型的服务接口
type Service interface {
	Insert(ctx context.Context, o *dto.Location) (*dto.Location, error)
	Update(ctx context.Context, id dxo.LocationID, o *dto.Location) (*dto.Location, error)

	Remove(ctx context.Context, id dxo.LocationID) error

	Find(ctx context.Context, id dxo.LocationID) (*dto.Location, error)
	List(ctx context.Context, q *Query) ([]*dto.Location, error)
}
