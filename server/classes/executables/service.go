package executables

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
)

// Service 表示该类型的服务接口
type Service interface {
	Insert(ctx context.Context, o *dto.Executable) (*dto.Executable, error)
	Update(ctx context.Context, id dxo.ExecutableID, o *dto.Executable) (*dto.Executable, error)

	Remove(ctx context.Context, id dxo.ExecutableID) error

	Find(ctx context.Context, id dxo.ExecutableID) (*dto.Executable, error)
	List(ctx context.Context, q *Query) ([]*dto.Executable, error)
}
