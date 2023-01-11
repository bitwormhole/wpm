package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// Example 表示一个命令模板
type Example struct {
	ID dxo.ExampleID `json:"id"`
	Base
}
