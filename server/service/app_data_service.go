package service

import "github.com/bitwormhole/starter/util"

// AppDataService ...
type AppDataService interface {
	GetAppDataDirectory() string

	GetSQLiteDBFile() string

	GetMainRepositoryPath() string

	GetBackupDumpDirectory() string

	GetBackupExecutableFile(sum util.Hex) string

	Ready() bool

	Setup() error
}
