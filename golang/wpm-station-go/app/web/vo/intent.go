package vo

import "github.com/bitwormhole/wpm/app/web/dto"

type Intent struct {
	Base

	Intents []*dto.Intent `json:"intents"`
}
