package service

import (
	"github.com/bitwormhole/wpm/app/data/dxo"
)

// UUIDGenService ...
type UUIDGenService interface {
	GenerateUUID(seed string) dxo.UUID
}
