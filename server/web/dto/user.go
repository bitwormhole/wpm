package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// User 表示一个用户
type User struct {
	ID dxo.UserID `json:"id"`
	Base

	Name dxo.UserName `json:"name"`
	Home string       `json:"home"` // the home dir path for this user
}
