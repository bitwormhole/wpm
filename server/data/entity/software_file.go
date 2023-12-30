package entity

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/starter-go/base/lang"
)

// InstalledFile ...
type InstalledFile struct {
	ID dxo.InstalledFileID `gorm:"primaryKey"`

	Base

	Name         string // the file name
	Path         string `gorm:"index:,unique"`
	Size         int64
	SHA256SUM    lang.Hex
	IsFile       bool
	IsDir        bool
	OwnerPackage dxo.SoftwarePackageURN
}
