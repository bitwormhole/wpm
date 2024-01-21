package softwaresets

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// softwaresets 是由若干个 packages 构成的虚拟对象，不具有对应的实体结构

// Service 表示该类型的服务接口
type Service interface {
	Insert(ctx context.Context, o *dto.SoftwareSet) (*dto.SoftwareSet, error)
	Update(ctx context.Context, id dxo.SoftwareSetID, o *dto.SoftwareSet) (*dto.SoftwareSet, error)

	Remove(ctx context.Context, id dxo.SoftwareSetID) error

	Find(ctx context.Context, id dxo.SoftwareSetID) (*dto.SoftwareSet, error)
	List(ctx context.Context, q *Query) ([]*dto.SoftwareSet, error)
}
