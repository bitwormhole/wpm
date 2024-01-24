package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// IntentTemplate ...
type IntentTemplate struct {
	Base

	Templates  []*dto.IntentTemplate `json:"intent_templates"`
	MacroProps map[string]string     `json:"macro_properties"`
}

// Intent ...
type Intent struct {
	Base

	Intents []*dto.Intent `json:"intents"`
}

// IntentQueue ...
type IntentQueue struct {
	Base

	Queue dto.IntentQueue `json:"queue"`
}