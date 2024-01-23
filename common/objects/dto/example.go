package dto

import "github.com/bitwormhole/wpm/common/objects/dxo"

// Example ... DTO
type Example struct {
	ID dxo.ExampleID `json:"id"`
	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
