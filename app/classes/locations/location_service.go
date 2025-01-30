package locations

import (
	"context"

	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/bitwormhole/wpm/app/web/dto"
	"github.com/starter-go/rbac"
)

// Query ...
type Query struct {
	Pagination rbac.Pagination
	All        bool
}

// Service ...
type Service interface {
	Find(c context.Context, id dxo.LocationID) (*dto.Location, error)

	ListWithCategory(c context.Context, category dxo.Category, q *Query) ([]*dto.Location, error)

	Insert(c context.Context, item *dto.Location) (*dto.Location, error)

	Update(c context.Context, id dxo.LocationID, item *dto.Location) (*dto.Location, error)

	Remove(c context.Context, id dxo.LocationID) error
}
