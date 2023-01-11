package client

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentHandler ...
type IntentHandler interface {
	HandleIntent(i *dto.Intent) error
}

////////////////////////////////////////////////////////////////////////////////

// IntentHandlerImpl ...
type IntentHandlerImpl struct {
	markup.Component `id:"IntentHandler"`
}

func (inst *IntentHandlerImpl) _Impl() IntentHandler {
	return inst
}

// HandleIntent ...
func (inst *IntentHandlerImpl) HandleIntent(i *dto.Intent) error {
	if i == nil {
		return nil
	}
	vlog.Info("exec intent: ", i.Executable)
	return nil
}
