package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// MediaService ...
type MediaService interface {
	Find(ctx context.Context, id dxo.MediaID, opt *MediaOptions) (*dto.Media, error)

	FindByIDs(ctx context.Context, ids []dxo.MediaID, opt *MediaOptions) ([]*dto.Media, error)

	FindByPath(ctx context.Context, path string, opt *MediaOptions) (*dto.Media, error)

	PrepareForDownload(ctx context.Context, me *dto.Media) (*dto.Media, error)

	ListAll(ctx context.Context, opt *MediaOptions) ([]*dto.Media, error)

	Insert(ctx context.Context, o *dto.Media) (*dto.Media, error)
	Update(ctx context.Context, id dxo.MediaID, o *dto.Media) (*dto.Media, error)
	Remove(ctx context.Context, id dxo.MediaID) error
}

// MediaOptions ...
type MediaOptions struct {
	All           bool
	WithFileState bool
}
