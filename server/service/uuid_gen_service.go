package service

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
)

// UUIDGenService ...
type UUIDGenService interface {
	GenerateUUID(seed string) dxo.UUID
}
