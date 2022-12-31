package entity

import (
	"time"

	"github.com/bitwormhole/wpm/app/data/dxo"
	"gorm.io/gorm"
)

// Base ...
type Base struct {
	UUID dxo.UUID

	Committer dxo.UserID
	Creator   dxo.UserID
	Owner     dxo.UserID

	// ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
