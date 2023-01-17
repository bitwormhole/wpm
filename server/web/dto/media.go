package dto

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
)

// Media 表示一个命令模板
type Media struct {
	ID dxo.MediaID `json:"id"`
	Base

	FileSize      int64    `json:"size"`
	ContentType   string   `json:"content_type"`
	SHA256SUM     util.Hex `json:"sha256sum"`
	Path          string   `json:"path"`
	WebPath       string   `json:"web_path"`
	LocalFilePath string   `json:"local_file_path"`
}
