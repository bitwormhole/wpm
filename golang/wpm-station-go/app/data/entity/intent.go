package entity

import "github.com/bitwormhole/wpm/app/data/dxo"

// Intent ...
type Intent struct {
	ID dxo.IntentID `gorm:"primaryKey"`
	Base

	Executable dxo.ExecutableID
}
