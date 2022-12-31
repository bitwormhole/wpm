package dto

import "github.com/bitwormhole/wpm/app/data/dxo"

// Base   所有DTO 的基本结构
type Base struct {
	// ID   int
	// Name string

	UUID dxo.UUID `json:"uuid"`
}
