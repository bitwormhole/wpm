package dto

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/starter-go/base/lang"
)

// Media 表示一个命令模板
type Media struct {
	ID dxo.MediaID `json:"id"`
	Base

	FileSize  int64    `json:"size"`
	SHA256SUM lang.Hex `json:"sha256sum"`

	Bucket      string `json:"bucket"`
	ContentType string `json:"content_type"`
	Label       string `json:"label"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Source      string `json:"source"` // 一个URL，file:// 或者 https://

	// LocalFilePath string `json:"local_file_path"` // 已经废弃，用 Source 代替

}
