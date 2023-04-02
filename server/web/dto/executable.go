package dto

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
)

// Executable ...
type Executable struct {
	ID  dxo.ExecutableID  `json:"id"`
	URN dxo.ExecutableURN `json:"urn"`

	Base

	Name             string                  `json:"name"`
	Aliases          dxo.StringList          `json:"aliases"`
	Namespace        string                  `json:"namespace"`
	Title            string                  `json:"title"`
	Description      string                  `json:"description"`
	IconURL          string                  `json:"icon"`
	SHA256SUM        util.Hex                `json:"sha256sum"`
	Remark           string                  `json:"remark"`
	Size             int64                   `json:"size"`
	Ready            bool                    `json:"ready"`
	Group            dxo.ExecutableGroupName `json:"group"`
	State            dxo.FileState           `json:"state"`
	Tags             dxo.StringList          `json:"tags"`
	OpenWithPriority int                     `json:"open_with_priority"` // 如果 value<=0, 表示 disable

	OS   string `json:"os"`   // 操作系统
	Arch string `json:"arch"` // 处理器架构

	Path     string            `json:"path"`
	Class    dxo.LocationClass `json:"location_class"`
	Location dxo.LocationID    `json:"location"`
}
