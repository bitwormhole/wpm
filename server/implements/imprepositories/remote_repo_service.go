package imprepositories

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/server/classes/examples"
	"github.com/bitwormhole/wpm/server/classes/repositories"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// RemoteRepositoryService ...
type RemoteRepositoryService struct {

	//starter:component
	_as func(repositories.RemoteRepositoryService) //starter:as("#")

	Dao examples.DAO //starter:inject("#")
}

func (inst *RemoteRepositoryService) _impl() repositories.RemoteRepositoryService {
	return inst
}

// Insert ...
func (inst *RemoteRepositoryService) Insert(ctx context.Context, o *dto.RemoteRepository) (*dto.RemoteRepository, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *RemoteRepositoryService) Update(ctx context.Context, id dxo.RemoteRepositoryID, o *dto.RemoteRepository) (*dto.RemoteRepository, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *RemoteRepositoryService) Remove(ctx context.Context, id dxo.RemoteRepositoryID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *RemoteRepositoryService) Find(ctx context.Context, id dxo.RemoteRepositoryID) (*dto.RemoteRepository, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *RemoteRepositoryService) List(ctx context.Context, q *repositories.Query) ([]*dto.RemoteRepository, error) {
	return nil, fmt.Errorf("no impl")
}
