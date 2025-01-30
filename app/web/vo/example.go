package vo

import "github.com/bitwormhole/wpm/app/web/dto"

// Example ... VO
type Example struct {
	Base

	Items []*dto.Example `json:"items"`
}
