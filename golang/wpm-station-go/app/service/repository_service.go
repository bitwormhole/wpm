package service

import (
	"context"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/bitwormhole/wpm/app/web/dto"
	"github.com/bitwormhole/wpm/app/web/vo"
)

// RepositoryService ...
type RepositoryService interface {
	Find(ctx context.Context, id dxo.RepositoryID) (*dto.Repository, error)
	FindByName(ctx context.Context, name string) (*dto.Repository, error)

	ListAll(ctx context.Context) ([]*dto.Repository, error)

	Insert(ctx context.Context, o *dto.Repository) (*dto.Repository, error)
	Update(ctx context.Context, id dxo.RepositoryID, o *dto.Repository) (*dto.Repository, error)
	Remove(ctx context.Context, id dxo.RepositoryID) error
}

// RepositoryImportService ...
type RepositoryImportService interface {
	Find(ctx context.Context, o *vo.RepositoryImport) (*vo.RepositoryImport, error)
	Locate(ctx context.Context, o *vo.RepositoryImport) (*vo.RepositoryImport, error)
	FindOrLocate(ctx context.Context, o *vo.RepositoryImport) (*vo.RepositoryImport, error)
	Save(ctx context.Context, o *vo.RepositoryImport) (*vo.RepositoryImport, error)
}

// RepositoryLocateService ...
type RepositoryLocateService interface {
	Locate(ctx context.Context, path fs.Path) (*dto.Repository, error)
	IsRepositoryDirectory(ctx context.Context, path fs.Path) bool
}

// RepositorySearchService ...
type RepositorySearchService interface {
	Search(ctx context.Context, path fs.Path) ([]*dto.Repository, error)
}
