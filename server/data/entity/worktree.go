package entity

import "github.com/bitwormhole/wpm/common/objects/dxo"

// Worktree 表示一颗本地的工作树
type Worktree struct {
	ID dxo.WorktreeID `gorm:"primaryKey"`
	Base
	URN dxo.URN

	Name string

	DotGitPath       string
	WorkingDirectory string
	RegularPath      string `gorm:"index:,unique"` // = DotGitPath
	Path             string // this.Path == Location.Path == DotGitPath

	// Location dxo.LocationID `gorm:"index:,unique"`
	// Class    dxo.LocationClass

	OwnerRepository dxo.LocalRepositoryID
}

// ListPathFields ...
func (Worktree) ListPathFields() []string {
	return []string{"path", "dot_git_path", "working_directory", "regular_path"}
}

// TableName 。。。
func (Worktree) TableName() string {
	return getTableName("worktrees")
}
