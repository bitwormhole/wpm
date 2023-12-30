package entity

type UserConfig struct {
	Name        string `gorm:"index:,unique"`
	ContentType string
}
