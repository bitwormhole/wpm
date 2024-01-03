package entity

// SystemConfig ...
type SystemConfig struct {
	Name        string `gorm:"index:,unique"`
	ContentType string
}

// TableName 。。。
func (SystemConfig) TableName() string {
	return getTableName("system_config")
}
