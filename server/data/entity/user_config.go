package entity

// UserConfig ...
type UserConfig struct {
	Name        string `gorm:"index:,unique"`
	ContentType string
}

// TableName 。。。
func (UserConfig) TableName() string {
	return getTableName("user_config")
}
