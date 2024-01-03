package packages

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// Service 表示该类型的服务接口
type Service interface {
	Insert(ctx context.Context, o *dto.SoftwarePackage) (*dto.SoftwarePackage, error)
	Update(ctx context.Context, id dxo.SoftwarePackageID, o *dto.SoftwarePackage) (*dto.SoftwarePackage, error)

	Remove(ctx context.Context, id dxo.SoftwarePackageID) error

	Find(ctx context.Context, id dxo.SoftwarePackageID) (*dto.SoftwarePackage, error)
	List(ctx context.Context, q *Query) ([]*dto.SoftwarePackage, error)
}
