package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Pipe ...
type Pipe struct {
	Base

	Pipes []*dto.Pipe `json:"pipes"`

	// Intents []*dto.Intent `json:"intents"`
}
