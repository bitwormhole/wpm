package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// Intent 表示一个命令模板
type Intent struct {
	ID dxo.IntentID `json:"id"`
	Base

	Executable dxo.ExecutableID `json:"executable"`
}
