package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// User ...
type User struct {
	Base

	Users []*dto.User `json:"users"`
}
