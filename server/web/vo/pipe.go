package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Pipe ...
type Pipe struct {
	Base

	// Current dto.PipeInfo      `json:"pipe"`

	Pipes   []*dto.PipeInfo   `json:"pipes"`
	Packets []*dto.PipePacket `json:"packets"`
}
