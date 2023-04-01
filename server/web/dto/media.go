package dto

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
)

// Media 表示一个命令模板
type Media struct {
	ID dxo.MediaID `json:"id"`
	Base

	FileSize  int64    `json:"size"`
	SHA256SUM util.Hex `json:"sha256sum"`

	Bucket        string `json:"bucket"`
	ContentType   string `json:"content_type"`
	Label         string `json:"label"`
	LocalFilePath string `json:"local_file_path"`
	Name          string `json:"name"`
	URL           string `json:"url"`
	Source        string `json:"source"`
}
