package entity

import (
	"time"

	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
)

// SoftwarePackage ...
type SoftwarePackage struct {
	ID  dxo.SoftwarePackageID  `gorm:"primaryKey"`
	URN dxo.SoftwarePackageURN `gorm:"index:,unique"`

	Base

	Name       string
	FileName   string
	ModuleName string // = namespace + "#" + name
	Namespace  string

	Icon        string
	Title       string
	Description string

	ContentType string
	Size        int64
	SHA256SUM   util.Hex

	WebPageURL  string
	DownloadURL string
	ResourceURL string

	OS        string
	Arch      string
	Version   string
	Revision  int
	ReleaseAt time.Time

	Installed bool
}
