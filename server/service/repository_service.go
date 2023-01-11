package service

import (
	"context"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// RemoteRepositoryService ...
type RemoteRepositoryService interface {
	Find(ctx context.Context, id dxo.RemoteRepositoryID) (*dto.RemoteRepository, error)
	FindByName(ctx context.Context, name string) (*dto.RemoteRepository, error)

	ListAll(ctx context.Context) ([]*dto.RemoteRepository, error)

	Insert(ctx context.Context, o *dto.RemoteRepository) (*dto.RemoteRepository, error)
	Update(ctx context.Context, id dxo.RemoteRepositoryID, o *dto.RemoteRepository) (*dto.RemoteRepository, error)
	Remove(ctx context.Context, id dxo.RemoteRepositoryID) error
}

// LocalRepositoryService ...
type LocalRepositoryService interface {
	Find(ctx context.Context, id dxo.LocalRepositoryID) (*dto.LocalRepository, error)
	FindByName(ctx context.Context, name string) (*dto.LocalRepository, error)

	ListAll(ctx context.Context) ([]*dto.LocalRepository, error)

	Insert(ctx context.Context, o *dto.LocalRepository) (*dto.LocalRepository, error)
	Update(ctx context.Context, id dxo.LocalRepositoryID, o *dto.LocalRepository) (*dto.LocalRepository, error)
	Remove(ctx context.Context, id dxo.LocalRepositoryID) error
}

// MainRepositoryService ...
type MainRepositoryService interface {
	Find(ctx context.Context, id dxo.MainRepositoryID) (*dto.MainRepository, error)
	FindByName(ctx context.Context, name string) (*dto.MainRepository, error)

	ListAll(ctx context.Context) ([]*dto.MainRepository, error)

	Insert(ctx context.Context, o *dto.MainRepository) (*dto.MainRepository, error)
	Update(ctx context.Context, id dxo.MainRepositoryID, o *dto.MainRepository) (*dto.MainRepository, error)
	Remove(ctx context.Context, id dxo.MainRepositoryID) error
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
	Locate(ctx context.Context, path fs.Path) (*dto.LocalRepository, error)
	IsRepositoryDirectory(ctx context.Context, path fs.Path) bool
}

// RepositorySearchService ...
type RepositorySearchService interface {
	Search(ctx context.Context, path fs.Path) ([]*dto.LocalRepository, error)
}
