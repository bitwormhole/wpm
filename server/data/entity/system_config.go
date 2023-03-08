package entity

type SystemConfig struct {
	Name        string `gorm:"index:,unique"`
	ContentType string
}
