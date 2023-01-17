package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Intent ...
type Intent struct {
	Base

	Intents []*dto.Intent `json:"intents"`
}

// IntentTemplate ...
type IntentTemplate struct {
	Base

	Templates []*dto.IntentTemplate `json:"intent_templates"`
}
