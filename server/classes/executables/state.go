package executables

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
)

// StateLoader ...
type StateLoader interface {
	Load(ctx context.Context, id dxo.ExecutableID, o *dto.Executable) error
}
