package intenttemplates

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// Service 表示该类型的服务接口
type Service interface {
	Insert(ctx context.Context, o *dto.IntentTemplate) (*dto.IntentTemplate, error)
	Update(ctx context.Context, id dxo.IntentTemplateID, o *dto.IntentTemplate) (*dto.IntentTemplate, error)

	Remove(ctx context.Context, id dxo.IntentTemplateID) error

	Find(ctx context.Context, id dxo.IntentTemplateID) (*dto.IntentTemplate, error)
	List(ctx context.Context, q *Query) ([]*dto.IntentTemplate, error)
}
