package entity

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
)

// Executable ...
type Executable struct {
	ID dxo.ExecutableID `gorm:"primaryKey"`
	Base

	Name        string
	Title       string
	Description string
	Path        string
	Size        int64
	SHA256SUM   util.Hex
}
