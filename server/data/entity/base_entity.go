package entity

import (
	"time"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"gorm.io/gorm"
)

// Base ...
type Base struct {
	UUID dxo.UUID `gorm:"index:,unique"`

	// ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Committer dxo.UserID
	Creator   dxo.UserID
	Owner     dxo.UserID

	Referer string // a URL, refer to owner document of this entity

	Installation dxo.InstallationID // 用来跟踪软件包安装资源项
}
