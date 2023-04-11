package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Executable ...
type Executable struct {
	Base

	Executables []*dto.Executable `json:"executables"`
}
