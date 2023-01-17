package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// Profile ...
type Profile struct {
	ID dxo.ProfileID
	Base

	User string `json:"user"`
	Home string `json:"home"`
}
