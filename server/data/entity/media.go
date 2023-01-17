package entity

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
)

// Media ...
type Media struct {
	ID dxo.MediaID `gorm:"primaryKey"`
	Base

	Path        string
	SHA256SUM   util.Hex
	FileSize    int64
	ContentType string
}
