package vo

import "github.com/bitwormhole/wpm/server/data/entity"

// Trash ...
type Trash struct {
	Base

	TrashItems []*entity.Holder `json:"trash_items"`
}
