package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Example ...
type Example struct {
	Base

	Examples []*dto.Example `json:"examples"`
}
