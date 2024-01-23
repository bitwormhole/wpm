package entity

import "github.com/bitwormhole/wpm/common/objects/dxo"

// Example ...
type Example struct {
	ID dxo.ExampleID

	Base

	Foo string
	Bar int
}

// TableName 。。。
func (Example) TableName() string {
	return getTableName("example")
}
