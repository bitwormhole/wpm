package dto

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
)

// Base   所有DTO 的基本结构
type Base struct {
	// ID   int
	// Name string

	UUID      dxo.UUID  `json:"uuid"`
	CreatedAt util.Time `json:"created_at"`
	UpdatedAt util.Time `json:"updated_at"`

	Referer      string             `json:"referer"`      // the owner document URL
	Installation dxo.InstallationID `json:"installation"` // 用来跟踪软件包安装资源项
}
