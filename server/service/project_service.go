package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// ProjectOptions ...
type ProjectOptions struct {
	All bool

	WithFileState bool
	WithGitStatus bool
	//   WithProjects  bool
}

// ProjectService ...
type ProjectService interface {
	Find(ctx context.Context, id dxo.ProjectID, options *ProjectOptions) (*dto.Project, error)
	FindByOwnerRepository(ctx context.Context, id dxo.LocalRepositoryID, options *ProjectOptions) ([]*dto.Project, error)
	ListAll(ctx context.Context, options *ProjectOptions) ([]*dto.Project, error)
	ListByIds(ctx context.Context, ids []dxo.ProjectID, options *ProjectOptions) ([]*dto.Project, error)

	Locate(ctx context.Context, o *dto.Project) (*dto.Project, error)

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
