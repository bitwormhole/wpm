package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// Worktree 表示一颗本地的工作树
type Worktree struct {
	ID dxo.WorktreeID `gorm:"primaryKey"`
	Base

	Name             string
	DotGitPath       string `gorm:"index:,unique"`
	WorkingDirectory string

	OwnerRepository dxo.LocalRepositoryID
}
