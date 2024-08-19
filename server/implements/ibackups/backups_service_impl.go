package ibackups

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/backups"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(backups.Service) //starter:as("#")

	Dao backups.DAO //starter:inject("#")
}

func (inst *ServiceImpl) _impl() backups.Service {
	return inst
}

// Insert ...
func (inst *ServiceImpl) Insert(ctx context.Context, o *dto.Backup) (*dto.Backup, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *ServiceImpl) Update(ctx context.Context, id dxo.BackupID, o *dto.Backup) (*dto.Backup, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *ServiceImpl) Remove(ctx context.Context, id dxo.BackupID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *ServiceImpl) Find(ctx context.Context, id dxo.BackupID) (*dto.Backup, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *ServiceImpl) List(ctx context.Context, q *backups.Query) ([]*dto.Backup, error) {
	return nil, fmt.Errorf("no impl")
}
