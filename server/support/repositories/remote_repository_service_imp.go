package repositories

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// RemoteRepositoryServiceImpl ...
type RemoteRepositoryServiceImpl struct {
	markup.Component `id:"RemoteRepositoryService"`
}

func (inst *RemoteRepositoryServiceImpl) _Impl() service.RemoteRepositoryService {
	return inst
}

// Find ...
func (inst *RemoteRepositoryServiceImpl) Find(ctx context.Context, id dxo.RemoteRepositoryID) (*dto.RemoteRepository, error) {
	return nil, fmt.Errorf("no impl")
}

// FindByName ...
func (inst *RemoteRepositoryServiceImpl) FindByName(ctx context.Context, name string) (*dto.RemoteRepository, error) {
	return nil, fmt.Errorf("no impl")
}

// ListAll ...
func (inst *RemoteRepositoryServiceImpl) ListAll(ctx context.Context) ([]*dto.RemoteRepository, error) {
	return nil, fmt.Errorf("no impl")
}

// Insert ...
func (inst *RemoteRepositoryServiceImpl) Insert(ctx context.Context, o *dto.RemoteRepository) (*dto.RemoteRepository, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *RemoteRepositoryServiceImpl) Update(ctx context.Context, id dxo.RemoteRepositoryID, o *dto.RemoteRepository) (*dto.RemoteRepository, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *RemoteRepositoryServiceImpl) Remove(ctx context.Context, id dxo.RemoteRepositoryID) error {
	return fmt.Errorf("no impl")
}
