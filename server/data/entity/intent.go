package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// IntentTemplate ...
type IntentTemplate struct {
	ID dxo.IntentTemplateID `gorm:"primaryKey"`
	Base

	Name        string
	Title       string
	Group       string
	Description string
	IconURL     string

	Selector    dxo.IntentTemplateSelector `gorm:"index:,unique"` // 根据（Action + TargetType + ExecutableName）生成
	Method      string
	ContentType string // project-type|content-type|target-type
	// Target      string            // the type of target: file|folder|repository|worktree|...|
	Executable dxo.ExecutableURN // the URN of exe

	Command   string
	Arguments dxo.StringListCRLF // as []string
	Env       dxo.StringMap      // as map[string]string
	WD        string
}

// TableName 。。。
func (IntentTemplate) TableName() string {
	return getTableName("intent_templates")
}
