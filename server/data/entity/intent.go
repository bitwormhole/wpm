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
	IconURL     string

	Selector    dxo.IntentTemplateSelector `gorm:"index:,unique"` // 根据（Action + TargetType + ExecutableName）生成
	Method      string
	ContentType string            // project-type|content-type
	Target      string            // the type of target: file|folder|repository|worktree|...|
	Executable  dxo.ExecutableURN // the URN of exe

	Command   string
	Arguments dxo.StringListCRLF // as []string
	Env       dxo.StringMap      // as map[string]string
	WD        string

	// Executable dxo.ExecutableID
}
