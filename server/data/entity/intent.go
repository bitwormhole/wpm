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

	Name        string
	Title       string
	Description string

	Selector   dxo.IntentTemplateSelector // 已废弃，用（Action + TargetType + ExecutableName）
	Action     string
	Executable string // the name of exe
	Target     string // the type of target

	Command   string
	Arguments dxo.StringListCRLF // as []string
	Env       dxo.StringMap      // as map[string]string
	WD        string

	// Executable dxo.ExecutableID
}
