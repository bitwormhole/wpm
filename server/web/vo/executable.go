package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Executable ...
type Executable struct {
	Base

	OptionSkipFileChecking bool `json:"option_skip_file_checking"`
	OptionIgnoreException  bool `json:"option_ignore_exception"`

	Executables []*dto.Executable `json:"executables"`
}
