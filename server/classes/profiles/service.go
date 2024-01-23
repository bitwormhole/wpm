package profiles

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
)

// Service 表示该类型的服务接口
type Service interface {
	Insert(ctx context.Context, o *dto.Profile) (*dto.Profile, error)
	Update(ctx context.Context, id dxo.ProfileID, o *dto.Profile) (*dto.Profile, error)

	Remove(ctx context.Context, id dxo.ProfileID) error

	Find(ctx context.Context, id dxo.ProfileID) (*dto.Profile, error)
	List(ctx context.Context, q *Query) ([]*dto.Profile, error)
}
