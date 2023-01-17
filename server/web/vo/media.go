package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Media ...
type Media struct {
	Base

	Mediae []*dto.Media `json:"mediae"`
}
