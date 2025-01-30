package entity

import (
	"github.com/bitwormhole/wpm/app/data/dxo"
)

// Location ... 表示一个文件或者目录的位置
type Location struct {
	ID dxo.LocationID

	Base

	Label       string
	Description string
	Tags        string // 以逗号分隔的一组标签

	Category   dxo.Category
	RawURI     dxo.URI
	RegularURI dxo.URI `gorm:"unique"`
	MediaType  dxo.MediaType

	AsFile   bool
	AsFolder bool
}
