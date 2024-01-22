package media

import (
	"context"
	"io"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/starter-go/afs"
)

// Service 表示该类型的服务接口
type Service interface {
	Insert(ctx context.Context, o *dto.Media) (*dto.Media, error)
	Update(ctx context.Context, id dxo.MediaID, o *dto.Media) (*dto.Media, error)

	Remove(ctx context.Context, id dxo.MediaID) error

	Find(ctx context.Context, id dxo.MediaID) (*dto.Media, error)
	List(ctx context.Context, q *Query) ([]*dto.Media, error)
	FindMediaFile(ctx context.Context, want *dto.Media) (*dto.Media, afs.Path, error)

	ImportMediaStream(ctx context.Context, o *dto.Media, data io.Reader) (*dto.Media, error)
}
