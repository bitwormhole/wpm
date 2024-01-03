package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

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
