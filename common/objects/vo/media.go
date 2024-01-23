package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// Media ...
type Media struct {
	Base

	Mediae []*dto.Media `json:"mediae"`
}
