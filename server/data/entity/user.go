package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// User ...
type User struct {
	ID dxo.UserID
	Base

	Name     dxo.UserName `gorm:"index:,unique"`
	Avatar   string
	Nickname string
	Home     string // the home dir path
}

// TableName 。。。
func (User) TableName() string {
	return getTableName("users")
}
