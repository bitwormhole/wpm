package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// Pipe 表示一个命令模板
type Pipe struct {
	ID dxo.PipeID `json:"id"`
	Base

	DesktopSessionID   string `json:"desktop_session_id"`
	DesktopSessionUser string `json:"desktop_session_user"`
	DesktopSessionHome string `json:"desktop_session_home"`

	Intent *Intent `json:"intent"`
}
