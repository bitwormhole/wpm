package dto

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
)

// PipeInfo 表示一个管道信息
type PipeInfo struct {
	ID dxo.PipeID `json:"id"`
	Base

	Name       dxo.PipeName `json:"name"`
	AttachedAt util.Time    `json:"attached_at"`

	DesktopSessionID   dxo.DesktopSessionID `json:"desktop_session_id"`
	DesktopSessionUser string               `json:"desktop_session_user"`
	DesktopSessionHome string               `json:"desktop_session_home"`
}

// PipePacket 表示一个管道包
type PipePacket struct {
	Intent *Intent `json:"intent"`
}
