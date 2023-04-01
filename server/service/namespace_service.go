package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// NamespaceService ...
type NamespaceService interface {
	Find(ctx context.Context, id dxo.NamespaceID) (*dto.Namespace, error)

	ListAll(ctx context.Context) ([]*dto.Namespace, error)

	Insert(ctx context.Context, o *dto.Namespace) (*dto.Namespace, error)
	Update(ctx context.Context, id dxo.NamespaceID, o *dto.Namespace) (*dto.Namespace, error)
	Remove(ctx context.Context, id dxo.NamespaceID) error
}
