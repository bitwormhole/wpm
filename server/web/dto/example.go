package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// Example ... DTO
type Example struct {
	ID dxo.ExampleID `json:"id"`
	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
