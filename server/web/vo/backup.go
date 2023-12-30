package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Backup ...
type Backup struct {
	Base

	Backups []*dto.Backup `json:"backups"`
}
