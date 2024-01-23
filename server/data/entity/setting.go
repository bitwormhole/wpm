package entity

import "github.com/bitwormhole/wpm/common/objects/dxo"

// Setting  表示一个设置项
type Setting struct {
	ID dxo.SettingID `gorm:"primaryKey"`
	Base

	Name        string `gorm:"index:,unique"`
	Value       string
	Description string
	Type        string
}

// TableName 。。。
func (Setting) TableName() string {
	return getTableName("settings")
}
