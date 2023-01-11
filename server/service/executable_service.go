package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// ExecutableService ...
type ExecutableService interface {
	Find(ctx context.Context, id dxo.ExecutableID) (*dto.Executable, error)
	FindByPath(ctx context.Context, path string) (*dto.Executable, error)

	ListAll(ctx context.Context) ([]*dto.Executable, error)

	Insert(ctx context.Context, o *dto.Executable) (*dto.Executable, error)
	Update(ctx context.Context, id dxo.ExecutableID, o *dto.Executable) (*dto.Executable, error)
	Remove(ctx context.Context, id dxo.ExecutableID) error
}

// ExecutableImportService ...
type ExecutableImportService interface {
	Save(ctx context.Context, o *vo.ExecutableImport) (*vo.ExecutableImport, error)
	Locate(ctx context.Context, o *vo.ExecutableImport) (*vo.ExecutableImport, error)
}
