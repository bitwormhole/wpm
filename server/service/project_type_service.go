package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ContentTypeService ...
type ContentTypeService interface {
	Find(ctx context.Context, id dxo.ContentTypeID) (*dto.ContentType, error)

	ListAll(ctx context.Context) ([]*dto.ContentType, error)

	ListByPattern(ctx context.Context, pattern string) ([]*dto.ContentType, error)
	ListByPatterns(ctx context.Context, patterns []string) ([]*dto.ContentType, error)

	Insert(ctx context.Context, o *dto.ContentType) (*dto.ContentType, error)
	Update(ctx context.Context, id dxo.ContentTypeID, o *dto.ContentType) (*dto.ContentType, error)
	Remove(ctx context.Context, id dxo.ContentTypeID) error

	// 根据文件名（或路径）查询对应的mime类型
	GetContentType(ctx context.Context, name string) (string, error)

	// LocateProject: 如果没有给出有效的参数 path，那么就用 o.FullPath 代替
	LocateProject(ctx context.Context, o *dto.Project, path string) error
}
