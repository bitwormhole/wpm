package entity

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
)

// Executable ...
type Executable struct {
	ID dxo.ExecutableID `gorm:"primaryKey"`
	Base

	Name             string
	Title            string
	IconURL          string
	Description      string
	Size             int64
	SHA256SUM        util.Hex
	OpenWithPriority int // 如果 value<=0, 表示 disable

	Path     string         // this.Path == Location.Path
	Location dxo.LocationID `gorm:"index:,unique"`
	Class    dxo.LocationClass
}
