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
}
