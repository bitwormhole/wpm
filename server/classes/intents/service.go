package intents

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
)

// Service 表示该类型的服务接口
type Service interface {
	Insert(ctx context.Context, o *dto.Intent) (*dto.Intent, error)
	Update(ctx context.Context, id dxo.IntentID, o *dto.Intent) (*dto.Intent, error)

	Remove(ctx context.Context, id dxo.IntentID) error

	Find(ctx context.Context, id dxo.IntentID) (*dto.Intent, error)
	List(ctx context.Context, q *Query) ([]*dto.Intent, error)
}
