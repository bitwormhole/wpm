package entity

import "github.com/bitwormhole/wpm/common/objects/dxo"

// Profile ...
type Profile struct {
	ID dxo.ProfileID `gorm:"primaryKey"`
	Base

	// URN dxo.URN

	Name        string
	Description string
}

// TableName 。。。
func (Profile) TableName() string {
	return getTableName("profiles")
}
