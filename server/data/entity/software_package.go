package entity

import (
	"time"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/starter-go/base/lang"
)

// SoftwarePackage ...
type SoftwarePackage struct {
	ID  dxo.SoftwarePackageID  `gorm:"primaryKey"`
	URN dxo.SoftwarePackageURN `gorm:"index:,unique"`

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

	OS        string
	Arch      string
	Version   string
	Revision  int
	ReleaseAt time.Time

	// Installed bool // 【已废弃】用 Installation  代替, >0 表示已安装
}
