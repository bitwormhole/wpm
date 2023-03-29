package service

import (
	"context"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// MediaService ...
type MediaService interface {

	// getters

	Find(ctx context.Context, id dxo.MediaID, opt *MediaOptions) (*dto.Media, error)

	FindByIDs(ctx context.Context, ids []dxo.MediaID, opt *MediaOptions) ([]*dto.Media, error)

	FindByPath(ctx context.Context, path string, opt *MediaOptions) (*dto.Media, error)

	PrepareForDownload(ctx context.Context, me *dto.Media) (*dto.Media, error)

	ListAll(ctx context.Context, opt *MediaOptions) ([]*dto.Media, error)

	ComputeMediaPath(ctx context.Context, me *dto.Media) (afs.Path, error)

	// setters

	ImportPresets(ctx context.Context) error

	Insert(ctx context.Context, o *dto.Media) (*dto.Media, error)
	Update(ctx context.Context, id dxo.MediaID, o *dto.Media) (*dto.Media, error)
	Remove(ctx context.Context, id dxo.MediaID) error
}

// MediaOptions ...
type MediaOptions struct {
	All           bool
	WithFileState bool
}
