package service

import (
	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/util"
)

// AppDataService ...
type AppDataService interface {
	GetAppDataDirectory() string

	GetSQLiteDBFile() string

	GetMainRepositoryPath() string

	GetBackupDumpDirectory() string

	GetBackupExecutableFile(sum util.Hex) string

	// get root dir
	GetRoot() afs.Path

	GetPath(opt *GetPathOptions) afs.Path

	Ready() bool

	Setup() error
}

// GetPathOptions 是用于 AppDataService.GetPath 的选项
type GetPathOptions struct {
	Type string
	ID   string
	Name string
}
