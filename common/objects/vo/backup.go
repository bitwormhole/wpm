package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// Backup ...
type Backup struct {
	Base

	Backups []*dto.Backup `json:"backups"`
}
