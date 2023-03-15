package dao

import "bitwormhole.com/starter/afs"

// type Store struct {
// 	RepositoryDAO   Repository   `inject:""`
// 	SystemConfigDAO SystemConfig `inject:""`
// 	UserConfigDAO   UserConfig   `inject:""`
// }

// Backup [inject:"#wpm-database-backup-dao"]
type Backup interface {
	Export(file afs.Path) error

	Import(file afs.Path) error
}
