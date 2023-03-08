package entity

import (
	"time"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"gorm.io/gorm"
)

// Base ...
type Base struct {
	UUID dxo.UUID `gorm:"index:,unique"`

	Committer dxo.UserID
	Creator   dxo.UserID
	Owner     dxo.UserID

	// ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
