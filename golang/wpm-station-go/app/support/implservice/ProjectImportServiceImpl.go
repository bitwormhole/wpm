package implservice

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/app/service"
	"github.com/bitwormhole/wpm/app/web/vo"
)

// ProjectImportServiceImpl ...
type ProjectImportServiceImpl struct {
	markup.Component `id:"ProjectImportService"`
}

func (inst *ProjectImportServiceImpl) _Impl() service.ProjectImportService {
	return inst
}

func (inst *ProjectImportServiceImpl) Find(ctx context.Context, o *vo.ProjectImport) (*vo.ProjectImport, error) {
	return nil, errors.New("no impl")
}

func (inst *ProjectImportServiceImpl) FindOrLocate(ctx context.Context, o *vo.ProjectImport) (*vo.ProjectImport, error) {
	return nil, errors.New("no impl")
}

func (inst *ProjectImportServiceImpl) Locate(ctx context.Context, o *vo.ProjectImport) (*vo.ProjectImport, error) {
	return nil, errors.New("no impl")
}

func (inst *ProjectImportServiceImpl) Save(ctx context.Context, o *vo.ProjectImport) (*vo.ProjectImport, error) {
	return nil, errors.New("no impl")
}
