package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// User ...
type User struct {
	Base

	Users []*dto.User `json:"users"`
}
