package about

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/vo"
)

// Service 表示该类型的服务接口
type Service interface {

	// Insert(ctx context.Context, o *dto.Example) (*dto.Example, error)
	// Update(ctx context.Context, id dxo.ExampleID, o *dto.Example) (*dto.Example, error)
	// Remove(ctx context.Context, id dxo.ExampleID) error
	// Find(ctx context.Context, id dxo.ExampleID) (*dto.Example, error)
	// List(ctx context.Context, q *Query) ([]*dto.Example, error)

	GetAboutInfo(ctx context.Context, dst *vo.About) error
}
