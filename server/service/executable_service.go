package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// ExecutableOptions ...
type ExecutableOptions struct {
	SkipFileChecking bool
	IgnoreException  bool
}

// ExecutableService ...
type ExecutableService interface {

	// getters

	Find(ctx context.Context, id dxo.ExecutableID, opt *ExecutableOptions) (*dto.Executable, error)

	FindByPath(ctx context.Context, path string, opt *ExecutableOptions) (*dto.Executable, error)

	FindByName(ctx context.Context, name string, opt *ExecutableOptions) (*dto.Executable, error)

	ListAll(ctx context.Context, opt *ExecutableOptions) ([]*dto.Executable, error)

	// setters

	Insert(ctx context.Context, o *dto.Executable, opt *ExecutableOptions) (*dto.Executable, error)

	Update(ctx context.Context, id dxo.ExecutableID, o *dto.Executable, opt *ExecutableOptions) (*dto.Executable, error)

	Remove(ctx context.Context, id dxo.ExecutableID, opt *ExecutableOptions) error
}

// ExecutableImportService ...
type ExecutableImportService interface {
	Save(ctx context.Context, o *vo.ExecutableImport) (*vo.ExecutableImport, error)

	Locate(ctx context.Context, o *vo.ExecutableImport) (*vo.ExecutableImport, error)

	ImportPresets(ctx context.Context) error
}
