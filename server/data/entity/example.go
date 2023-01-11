package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// Example ...
type Example struct {
	ID dxo.ExampleID `gorm:"primaryKey"`
	Base
}
