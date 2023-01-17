package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// Intent ...
// type Intent struct {
// 	ID dxo.IntentID `gorm:"primaryKey"`
// 	Base

// 	Executable dxo.ExecutableID
// }

// IntentTemplate ...
type IntentTemplate struct {
	ID dxo.IntentTemplateID `gorm:"primaryKey"`
	Base

	Executable dxo.ExecutableID
}
