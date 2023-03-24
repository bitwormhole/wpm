package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/web/vo"
)

// RepositoryWorktreeProjectService ...
type RepositoryWorktreeProjectService interface {
	Find(c context.Context, o *vo.RepositoryWorktreeProject) (*vo.RepositoryWorktreeProject, error)
	Save(c context.Context, o *vo.RepositoryWorktreeProject) (*vo.RepositoryWorktreeProject, error)
}
