package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/web/dto"
)

// DatabaseBackupService ...
type DatabaseBackupService interface {
	Export(c context.Context, o *dto.Backup) (*dto.Backup, error)

	Import(c context.Context, o *dto.Backup) (*dto.Backup, error)

	ListAll(c context.Context) ([]*dto.Backup, error)
}
