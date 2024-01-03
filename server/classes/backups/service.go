package backups

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// Service 表示该类型的服务接口
type Service interface {
	Insert(ctx context.Context, o *dto.Backup) (*dto.Backup, error)
	Update(ctx context.Context, id dxo.BackupID, o *dto.Backup) (*dto.Backup, error)

	Remove(ctx context.Context, id dxo.BackupID) error

	Find(ctx context.Context, id dxo.BackupID) (*dto.Backup, error)
	List(ctx context.Context, q *Query) ([]*dto.Backup, error)
}
