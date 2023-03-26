package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// Worktree 表示一颗本地的工作树
type Worktree struct {
	ID dxo.WorktreeID `gorm:"primaryKey"`
	Base

	Name string

	DotGitPath       string
	WorkingDirectory string

	Path     string         // this.Path == Location.Path == DotGitPath
	Location dxo.LocationID `gorm:"index:,unique"`
	Class    dxo.LocationClass

	OwnerRepository dxo.LocalRepositoryID
}

// ListPathFields ...
func (Worktree) ListPathFields() []string {
	return []string{"path", "dot_git_path", "working_directory"}
}
