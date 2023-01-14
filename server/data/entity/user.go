package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// User ...
type User struct {
	ID dxo.UserID
	Base

	Name     dxo.UserName
	Avatar   string
	Nickname string
	Home     string // the home dir path
}
