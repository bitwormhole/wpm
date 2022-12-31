package dto

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/app/data/dxo"
)

// Executable ...
type Executable struct {
	ID dxo.ExecutableID `json:"id"`
	Base

	Name      string                  `json:"name"`
	Title     string                  `json:"title"`
	Path      string                  `json:"path"`
	SHA256SUM util.Hex                `json:"sha256sum"`
	Remark    string                  `json:"remark"`
	Size      int64                   `json:"size"`
	Ready     bool                    `json:"ready"`
	Group     dxo.ExecutableGroupName `json:"group"`
	State     dxo.FileState           `json:"state"`
	Tags      dxo.StringList          `json:"tags"`
}
