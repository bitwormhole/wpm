package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ProjectTypeService ...
type ProjectTypeService interface {
	Find(ctx context.Context, id dxo.ProjectTypeID) (*dto.ProjectType, error)

	ListAll(ctx context.Context) ([]*dto.ProjectType, error)

	Insert(ctx context.Context, o *dto.ProjectType) (*dto.ProjectType, error)
	Update(ctx context.Context, id dxo.ProjectTypeID, o *dto.ProjectType) (*dto.ProjectType, error)
	Remove(ctx context.Context, id dxo.ProjectTypeID) error

	// LocateProject: 如果没有给出有效的参数 path，那么就用 o.FullPath 代替
	LocateProject(ctx context.Context, o *dto.Project, path string) error
}
