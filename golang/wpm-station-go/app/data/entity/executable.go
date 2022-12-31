package entity

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/app/data/dxo"
)

// Executable ...
type Executable struct {
	ID dxo.ExecutableID `gorm:"primaryKey"`
	Base

	Name      string
	Title     string
	Path      string
	Size      int64
	SHA256SUM util.Hex
}
