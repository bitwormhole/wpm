package vo

import "github.com/bitwormhole/wpm/server/web/dto"

type Intent struct {
	Base

	Intents []*dto.Intent `json:"intents"`
}
