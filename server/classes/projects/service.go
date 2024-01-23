package projects

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
)

// Service 表示该类型的服务接口
type Service interface {
	Insert(ctx context.Context, o *dto.Project) (*dto.Project, error)
	Update(ctx context.Context, id dxo.ProjectID, o *dto.Project) (*dto.Project, error)

	Remove(ctx context.Context, id dxo.ProjectID) error

	Find(ctx context.Context, id dxo.ProjectID) (*dto.Project, error)
	List(ctx context.Context, q *Query) ([]*dto.Project, error)
}
