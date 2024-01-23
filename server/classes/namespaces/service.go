package namespaces

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
)

// Service 表示该类型的服务接口
type Service interface {
	Insert(ctx context.Context, o *dto.Namespace) (*dto.Namespace, error)
	Update(ctx context.Context, id dxo.NamespaceID, o *dto.Namespace) (*dto.Namespace, error)

	Remove(ctx context.Context, id dxo.NamespaceID) error

	Find(ctx context.Context, id dxo.NamespaceID) (*dto.Namespace, error)
	List(ctx context.Context, q *Query) ([]*dto.Namespace, error)
}
