package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// Intent 表示一个命令模板
type Intent struct {
	ID dxo.IntentID `json:"id"`
	Base

	Pipe *PipeInfo `json:"pipe"`

	Exe *IntentExecutable `json:"exe"`
	Web *IntentWeb        `json:"web"`
	CLI *IntentCLI        `json:"cli"`
}

// IntentExecutable ...
type IntentExecutable struct {
	Executable dxo.ExecutableID `json:"executable"`
}

// IntentWeb ...
type IntentWeb struct {
	Method string `json:"method"`
	URL    string `json:"url"`
}

// IntentCLI ...
type IntentCLI struct {
	Command   string            `json:"command"`
	Arguments []string          `json:"args"`
	Env       map[string]string `json:"env"`
	WD        string            `json:"wd"`
}
