package vo

import "github.com/bitwormhole/wpm/server/web/dto"

type Executable struct {
	Base

	Executables []*dto.Executable `json:"executables"`
}
