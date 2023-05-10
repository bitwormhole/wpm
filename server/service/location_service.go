package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// LocationOptions ...
type LocationOptions struct {
	All           bool
	WithFileState bool
	WithGitStatus bool
}

// LocationService ...
type LocationService interface {

	// getters

	Find(ctx context.Context, id dxo.LocationID, opt *LocationOptions) (*dto.Location, error)

	FindByPath(ctx context.Context, path string, opt *LocationOptions) (*dto.Location, error)

	ListAll(ctx context.Context, opt *LocationOptions) ([]*dto.Location, error)

	//setters

	Insert(ctx context.Context, o *dto.Location) (*dto.Location, error)
	Update(ctx context.Context, id dxo.LocationID, o *dto.Location) (*dto.Location, error)
	Remove(ctx context.Context, id dxo.LocationID) error

	// InsertOrFetch(ctx context.Context, o *dto.Location, options *LocationOptions) (*dto.Location, error)

	PutExecutable(ctx context.Context, o *entity.Executable) error
	PutProject(ctx context.Context, o *entity.Project) error
	PutRepository(ctx context.Context, o *entity.LocalRepository) error
	PutWorktree(ctx context.Context, o *entity.Worktree) error
	// PutSubmodule(ctx context.Context, o *entity.PutSubmodule) error
}
