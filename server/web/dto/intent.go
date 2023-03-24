package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// Intent 表示一个命令实例
type Intent struct {
	ID dxo.IntentID `json:"id"`
	Base

	// Exe *IntentExecutable `json:"exe"`
	// Web *IntentWeb        `json:"web"`
	// CLI *IntentCLI        `json:"cli"`

	Options    []string          `json:"options"`    // used for filters
	Properties map[string]string `json:"properties"` // used for filters

	// as exe
	Executable *Executable      `json:"executable"`
	Repository *LocalRepository `json:"repository"`
	Worktree   *Worktree        `json:"worktree"`
	Project    *Project         `json:"project"`
	Location   *File            `json:"location"`
	Template   *IntentTemplate  `json:"template"`

	// as web
	Method string `json:"method"`
	URL    string `json:"url"`

	// as cli
	Command   string            `json:"command"`
	Arguments []string          `json:"args"`
	Env       map[string]string `json:"env"`
	WD        string            `json:"wd"`
}

// IntentTemplate 表示一个命令模板
type IntentTemplate struct {
	ID dxo.IntentTemplateID `json:"id"`
	Base

	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`

	Selector   dxo.IntentTemplateSelector `json:"selector"`
	Action     string                     `json:"action"`
	Target     string                     `json:"target"`     // the type of target
	Executable string                     `json:"executable"` // the name of exe

	// as cli
	Command   string             `json:"command"`
	Arguments dxo.StringListCRLF `json:"args"`
	Env       dxo.StringMap      `json:"env"`
	WD        string             `json:"wd"`

	// Exe *IntentExecutable `json:"exe"`
	// Web *IntentWeb        `json:"web"`
	// CLI *IntentCLI        `json:"cli"`
}

// // IntentExecutable ...
// type IntentExecutable struct {
// 	Executable *Executable      `json:"executable"`
// 	Repository *LocalRepository `json:"repository"`
// 	Path       string           `json:"path"` // 相对路径，基于工作区目录
// }

// // IntentWeb ...
// type IntentWeb struct {
// 	Method string `json:"method"`
// 	URL    string `json:"url"`
// }

// // IntentCLI ...
// type IntentCLI struct {
// 	Command   string            `json:"command"`
// 	Arguments []string          `json:"args"`
// 	Env       map[string]string `json:"env"`
// 	WD        string            `json:"wd"`
// }
