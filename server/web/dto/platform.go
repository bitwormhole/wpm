package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// Platform 表示平台信息
type Platform struct {
	ID dxo.PlatformID `json:"id"`
	Base

	OS   string `json:"os"`
	Arch string `json:"arch"`
	User string `json:"user"`
	Home string `json:"home"`
}
