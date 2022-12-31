package vo

import "github.com/bitwormhole/wpm/app/web/dto"

type Executable struct {
	Base

	Executables []*dto.Executable `json:"executables"`
}
