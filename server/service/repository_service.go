package service

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
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

// SystemMainRepositoryService ...
type SystemMainRepositoryService interface {
	GetInfo(ctx context.Context) (*dto.SystemMainRepository, error)

	GetRepository(ctx context.Context) (store.Repository, error)
}

// UserMainRepositoryService ...
type UserMainRepositoryService interface {
	Find(ctx context.Context, id dxo.UserMainRepositoryID) (*dto.UserMainRepository, error)
	FindByName(ctx context.Context, name string) (*dto.UserMainRepository, error)

	ListAll(ctx context.Context) ([]*dto.UserMainRepository, error)

	Insert(ctx context.Context, o *dto.UserMainRepository) (*dto.UserMainRepository, error)
	Update(ctx context.Context, id dxo.UserMainRepositoryID, o *dto.UserMainRepository) (*dto.UserMainRepository, error)
	Remove(ctx context.Context, id dxo.UserMainRepositoryID) error
}

// RepositoryImportService ...
type RepositoryImportService interface {
	Find(ctx context.Context, o *vo.RepositoryImport) (*vo.RepositoryImport, error)
	Locate(ctx context.Context, o *vo.RepositoryImport) (*vo.RepositoryImport, error)
	FindOrLocate(ctx context.Context, o *vo.RepositoryImport) (*vo.RepositoryImport, error)
	Save(ctx context.Context, o *vo.RepositoryImport) (*vo.RepositoryImport, error)
}

// LocalRepositoryFinder 。。。
type LocalRepositoryFinder interface {
	Search(ctx context.Context, path string, depthLimit int) ([]*dto.LocalRepository, error)
	Locate(ctx context.Context, path string) (*dto.LocalRepository, error)
}

// LocalRepositoryStateLoader 。。。
type LocalRepositoryStateLoader interface {
	LoadState(ctx context.Context, repo *dto.LocalRepository) error
}
