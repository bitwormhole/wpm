package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// MediaService ...
type MediaService interface {
	Find(ctx context.Context, id dxo.MediaID) (*dto.Media, error)

	FindByPath(ctx context.Context, path string) (*dto.Media, error)

	PrepareForDownload(ctx context.Context, me *dto.Media) (*dto.Media, error)

	ListAll(ctx context.Context) ([]*dto.Media, error)

	Insert(ctx context.Context, o *dto.Media) (*dto.Media, error)
	Update(ctx context.Context, id dxo.MediaID, o *dto.Media) (*dto.Media, error)
	Remove(ctx context.Context, id dxo.MediaID) error
}
