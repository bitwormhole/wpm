package dto

import "github.com/bitwormhole/wpm/common/objects/dxo"

// Platform 表示平台信息
type Platform struct {
	ID dxo.PlatformID `json:"id"`
	Base

	OS   string `json:"os"`
	Arch string `json:"arch"`
}
