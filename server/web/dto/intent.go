package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// Intent 表示一个命令实例
type Intent struct {
	ID dxo.IntentID `json:"id"`

	Base
	ActionRequest

	Properties map[string]string `json:"properties"`
	Template   *IntentTemplate   `json:"template"`
	Want       *IntentMessage    `json:"want"`
	Have       *IntentMessage    `json:"have"`

	// targets
	Executable *Executable      `json:"executable"`
	Repository *LocalRepository `json:"repository"`
	Worktree   *Worktree        `json:"worktree"`
	Submodule  *Submodule       `json:"submodule"`
	Project    *Project         `json:"project"`
	File       *File            `json:"file"`
	Folder     *File            `json:"folder"`
	Web        *WebRequest      `json:"web"`

	// as cli
	CLI *CommandRequest `json:"cli"`

	// result
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  int    `json:"status"` // use http.Status
}

type IntentMessage struct {
	Properties []*IntentPropertyDescriptor `json:"properties"`
	Templates  []*IntentTemplate           `json:"templates"`
}

// IntentTemplate 表示一个命令模板
type IntentTemplate struct {
	ID dxo.IntentTemplateID `json:"id"`
	Base

	Name        string `json:"name"`
	IconURL     string `json:"icon"`
	Title       string `json:"title"`
	Group       string `json:"group"`
	Description string `json:"description"`

	// Selector   dxo.IntentTemplateSelector `json:"selector"`
	// Action     string                     `json:"action"`
	// Target     string                     `json:"target"`     // the type of target
	// Executable dxo.ExecutableURN          `json:"executable"` // the URN of exe
	ActionRequest // `json:"action"`

	WantProperties []*IntentPropertyDescriptor `json:"want_properties"`

	// as cli
	Command   string             `json:"command"`
	Arguments dxo.StringListCRLF `json:"args"`
	Env       dxo.StringMap      `json:"env"`
	WD        string             `json:"wd"`
}

// IntentQueue 表示一个命令消息队列
type IntentQueue struct {
	ID dxo.IntentQueueID `json:"id"`
	Base

	UserHomeDir string    `json:"home"`
	Intents     []*Intent `json:"intents"`
}

// WebRequest ...
type WebRequest struct {
	Method string `json:"method"`
	URL    string `json:"url"`
}

// ActionRequest ...
type ActionRequest struct {
	Method      string                     `json:"method"`
	Name        string                     `json:"name"` // example: file.name|repository.name|...
	Label       string                     `json:"label"`
	Location    string                     `json:"location"` // a url or local-path
	ContentType string                     `json:"type"`     // project-type | content-type | target-type (file|folder|repository|worktree|submodule|project|web)
	With        dxo.ExecutableURN          `json:"with"`
	Selector    dxo.IntentTemplateSelector `json:"selector"`
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
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Value       string   `json:"value"`
	Type        string   `json:"type"`    // [string|int|bool|float|enum|...]
	Options     []string `json:"options"` // for enum
	Required    bool     `json:"required"`
}
