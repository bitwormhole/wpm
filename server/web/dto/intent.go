package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// Intent 表示一个命令实例
type Intent struct {
	ID dxo.IntentID `json:"id"`
	Base

	Options    []string          `json:"options"` // used for filters
	Properties map[string]string `json:"properties"`

	// targets
	Executable *Executable      `json:"executable"`
	Repository *LocalRepository `json:"repository"`
	Worktree   *Worktree        `json:"worktree"`
	Submodule  *Submodule       `json:"submodule"`
	Project    *Project         `json:"project"`
	File       *File            `json:"file"`
	Folder     *File            `json:"folder"`

	// templates
	Templates []*IntentTemplate `json:"templates"`

	// as action
	Action ActionRequest `json:"action"`

	// as web
	Web *WebRequest `json:"web"`

	// as cli
	CLI *CommandRequest `json:"cli"`

	// result
	WantProperties []*IntentPropertyDescriptor `json:"want_properties"`
	Message        string                      `json:"message"`
	Error          string                      `json:"error"`
	Status         int                         `json:"status"` // use http.Status
}

// IntentTemplate 表示一个命令模板
type IntentTemplate struct {
	ID dxo.IntentTemplateID `json:"id"`
	Base

	Name        string `json:"name"`
	IconURL     string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`

	// Selector   dxo.IntentTemplateSelector `json:"selector"`
	// Action     string                     `json:"action"`
	// Target     string                     `json:"target"`     // the type of target
	// Executable dxo.ExecutableURN          `json:"executable"` // the URN of exe
	Action ActionRequest `json:"action"`

	WantProperties []*IntentPropertyDescriptor `json:"want_properties"`

	// as cli
	Command   string             `json:"command"`
	Arguments dxo.StringListCRLF `json:"args"`
	Env       dxo.StringMap      `json:"env"`
	WD        string             `json:"wd"`
}

// WebRequest ...
type WebRequest struct {
	Method string `json:"method"`
	URL    string `json:"url"`
}

// ActionRequest ...
type ActionRequest struct {
	Method   string                     `json:"method"`
	Target   string                     `json:"target"` // file|folder|repository|worktree|submodule|project
	Type     string                     `json:"type"`   // project-type | content-type
	With     dxo.ExecutableURN          `json:"with"`
	Selector dxo.IntentTemplateSelector `json:"selector"`
}

// CommandRequest ...
type CommandRequest struct {
	Command   string            `json:"command"`
	Arguments []string          `json:"args"`
	Env       map[string]string `json:"env"`
	WD        string            `json:"wd"`
}

// IntentPropertyDescriptor ...
type IntentPropertyDescriptor struct {
	Name         string
	Description  string
	ValueDefault string
	Type         string   // [string|int|bool|float|enum|...]
	Options      []string // for enum
	Required     bool
}
