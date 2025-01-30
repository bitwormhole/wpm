package entity

import "github.com/bitwormhole/wpm/app/data/dxo"

// GetDataGroupInfo  ...
func GetDataGroupInfo() dxo.DataGroupInfo {
	return new(dgInfo)
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamePrefix string

type dgInfo struct{}

func (inst *dgInfo) Prototypes() []any {

	list := make([]any, 0)

	list = append(list, new(Example))
	list = append(list, new(Location))

	return list
}

func (inst *dgInfo) SetTableNamePrefix(prefix string) {
	if theTableNamePrefix == "" {
		theTableNamePrefix = prefix
	}
}

////////////////////////////////////////////////////////////////////////////////

// TableName 。。。
func (Example) TableName() string {
	return theTableNamePrefix + "examples"
}

// TableName 。。。
func (Location) TableName() string {
	return theTableNamePrefix + "locations"
}
