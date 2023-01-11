package implservice

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// MainRepositoryServiceImpl ...
type MainRepositoryServiceImpl struct {
	markup.Component `id:"MainRepositoryService"`
}

func (inst *MainRepositoryServiceImpl) _Impl() service.MainRepositoryService {
	return inst
}

func (inst *MainRepositoryServiceImpl) Find(ctx context.Context, id dxo.MainRepositoryID) (*dto.MainRepository, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *MainRepositoryServiceImpl) FindByName(ctx context.Context, name string) (*dto.MainRepository, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *MainRepositoryServiceImpl) ListAll(ctx context.Context) ([]*dto.MainRepository, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *MainRepositoryServiceImpl) Insert(ctx context.Context, o *dto.MainRepository) (*dto.MainRepository, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *MainRepositoryServiceImpl) Update(ctx context.Context, id dxo.MainRepositoryID, o *dto.MainRepository) (*dto.MainRepository, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *MainRepositoryServiceImpl) Remove(ctx context.Context, id dxo.MainRepositoryID) error {
	return fmt.Errorf("no impl")
}
