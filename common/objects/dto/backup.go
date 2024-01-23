package dto

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
)

// Backup 表示一个备份文件
type Backup struct {
	ID dxo.BackupID `json:"id"`
	Base

	BackupFileName string `json:"backup_file_name"`
	BackupFilePath string `json:"backup_file_path"`
}
