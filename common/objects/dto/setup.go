package dto

import "github.com/bitwormhole/wpm/common/objects/dxo"

// SetupState 表示设置项的状态
type SetupState string

// 定义设置项的各种状态
const (
	SetupStateInit    SetupState = "init"
	SetupStateDone    SetupState = "done"
	SetupStateWant    SetupState = "want"
	SetupStateHave    SetupState = "have"
	SetupStateSuccess SetupState = "success"
	SetupStateError   SetupState = "error"
)

// Setup 表示一个安装设置项
type Setup struct {
	ID dxo.SetupID `json:"id"`
	Base

	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`

	State SetupState `json:"state"`
	Error string     `json:"error"`

	Properties []*SetupProperty `json:"properties"`
}

// SetupProperty 表示一个安装设置项的属性
type SetupProperty struct {
	Name  string
	Value string
	Type  string // [string|bool|int|float|...]
}
