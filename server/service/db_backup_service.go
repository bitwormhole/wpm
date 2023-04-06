package service

import (
	"context"

	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// DatabaseBackupService ...
type DatabaseBackupService interface {
	Export(c context.Context, o *dto.Backup) (*dto.Backup, error)

	Import(c context.Context, o *dto.Backup) (*dto.Backup, error)

	ListAll(c context.Context) ([]*dto.Backup, error)

	// 导入转存数据
	ImportDumpData(c context.Context, form, to util.Time) error

	// 导出转存数据
	ExportDumpData(c context.Context) error
}
