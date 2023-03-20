package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// SetupService ...
type SetupService interface {

	// Find(ctx context.Context, id dxo.ExampleID) (*dto.Example, error)
	// Insert(ctx context.Context, o *dto.Example) (*dto.Example, error)
	// Update(ctx context.Context, id dxo.ExampleID, o *dto.Example) (*dto.Example, error)
	// Remove(ctx context.Context, id dxo.ExampleID) error

	IsSetupReqiured(ctx context.Context) (bool, error)

	ListAll(ctx context.Context) ([]*dto.Setup, error)

	Apply(ctx context.Context, items []*dto.Setup) error

	SkipAll(ctx context.Context) error
}

// SetupRegistration ...
type SetupRegistration struct {
	ID        dxo.SetupID
	Name      string
	Prototype *dto.Setup
	Handler   func(c context.Context, item *dto.Setup) error
}
