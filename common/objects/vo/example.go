package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// Example ... VO
type Example struct {
	Base

	Items []*dto.Example `json:"items"`
}
