package wpmboot

import (
	"github.com/bitwormhole/starter/application"
)

// ExportConfig ...
func ExportConfig(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
	// return nil
}
