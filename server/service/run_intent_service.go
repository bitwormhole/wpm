package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/web/dto"
)

// RunIntentService ...
type RunIntentService interface {

	// Find(ctx context.Context, id dxo.IntentID) (*dto.Intent, error)
	// ListAll(ctx context.Context) ([]*dto.Intent, error)
	// Insert(ctx context.Context, o *dto.Intent) (*dto.Intent, error)
	// Update(ctx context.Context, id dxo.IntentID, o *dto.Intent) (*dto.Intent, error)
	// Remove(ctx context.Context, id dxo.IntentID) error

	Run(ctx context.Context, o *dto.Intent) (*dto.Intent, error)
}
