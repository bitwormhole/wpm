package entity

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
)

// Executable ...
type Executable struct {
	ID  dxo.ExecutableID  `gorm:"primaryKey"`
	URN dxo.ExecutableURN // 同名，不同location的多个实体表示 同一个应用的不同版本

	Base

	Name             string
	Aliases          dxo.StringList
	Namespace        string
	Title            string
	IconURL          string
	Description      string
	Size             int64
	SHA256SUM        util.Hex
	OpenWithPriority int // 如果 value<=0, 表示 disable

	OS      string
	Arch    string
	Version string

	Path string `gorm:"index:,unique"`

	// this.Path == Location.Path
	// Location dxo.LocationID `gorm:"index:,unique"`
	// Class    dxo.LocationClass
}
