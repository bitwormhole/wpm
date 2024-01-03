package entity

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/starter-go/base/lang"
)

// Media ...
type Media struct {
	ID dxo.MediaID `gorm:"primaryKey"`
	Base

	Name   string
	URL    string `gorm:"index:,unique"`
	Source string // Source 表示资源的原始来源URL，如果本地没有该资源，可以通过此URL下载
	Bucket string
	Label  string

	ContentType   string
	ContentLength int64

	SHA256SUM lang.Hex
}

// TableName 。。。
func (Media) TableName() string {
	return getTableName("mediae")
}
