package entity

import (
	"time"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/starter-go/base/lang"
)

// SoftwareInfo ...
type SoftwareInfo struct {
	Base

	Namespace  string
	Name       string
	ModuleName string // = namespace + "#" + name
	FileName   string

	Icon        string
	Title       string
	Description string

	ContentType string
	Size        int64
	SHA256SUM   lang.Hex

	WebPageURL  string
	DownloadURL string
	ResourceURL string

	OS   string
	Arch string

	// Installed bool // 【已废弃】用 Installation  代替, >0 表示已安装

}

// SoftwarePackage ...
type SoftwarePackage struct {
	ID  dxo.SoftwarePackageID  `gorm:"primaryKey"`
	URN dxo.SoftwarePackageURN `gorm:"index:,unique"`

	SoftwareInfo

	Version   string
	Revision  int
	ReleaseAt time.Time
}

// SoftwareSet ...
type SoftwareSet struct {
	ID  dxo.SoftwareSetID  `gorm:"primaryKey"`
	URN dxo.SoftwareSetURN `gorm:"index:,unique"`

	SoftwareInfo
}

////////////////////////////////////////////////////////////////////////////////

// TableName 。。。
func (SoftwarePackage) TableName() string {
	return getTableName("packages")
}
