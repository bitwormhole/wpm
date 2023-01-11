package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// ProjectService ...
type ProjectService interface {
	Find(ctx context.Context, id dxo.ProjectID) (*dto.Project, error)
	FindByOwnerRepository(ctx context.Context, id dxo.LocalRepositoryID) ([]*dto.Project, error)

	ListAll(ctx context.Context) ([]*dto.Project, error)

	Insert(ctx context.Context, o *dto.Project) (*dto.Project, error)
	Update(ctx context.Context, id dxo.ProjectID, o *dto.Project) (*dto.Project, error)
	Remove(ctx context.Context, id dxo.ProjectID) error
}

// ProjectImportService ...
type ProjectImportService interface {
	Find(ctx context.Context, o *vo.ProjectImport) (*vo.ProjectImport, error)
	FindOrLocate(ctx context.Context, o *vo.ProjectImport) (*vo.ProjectImport, error)
	Locate(ctx context.Context, o *vo.ProjectImport) (*vo.ProjectImport, error)
	Save(ctx context.Context, o *vo.ProjectImport) (*vo.ProjectImport, error)
}
