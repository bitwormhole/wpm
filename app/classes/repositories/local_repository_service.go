package repositories

import (
	"context"

	"github.com/bitwormhole/wpm-api/v1/dto"
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/starter-go/rbac"
)

// Query ...
type Query struct {
	All        bool
	Pagination rbac.Pagination
}

// LocalRepositoryService ...
type LocalRepositoryService interface {
	Find(c context.Context, id dxo.LocalRepositoryID) (*dto.LocalRepository, error)

	List(c context.Context, q *Query) ([]*dto.LocalRepository, error)
}
