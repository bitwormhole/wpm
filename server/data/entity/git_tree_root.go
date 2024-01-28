package entity

import "github.com/bitwormhole/wpm/common/objects/dxo"

// TreeRoot 表示一颗本地树根目录（它可能是默认工作区、工作树、或者子模块的根目录）
type TreeRoot struct {
	ID dxo.TreeRootID `gorm:"primaryKey"`
	Base
	URN dxo.URN

	Name string

	Path       string // this.Path == DotGitPath.parentDir
	DotGitPath string // '.git' （文件|文件夹） 的路径。

	RegularPath dxo.RegularPath `gorm:"index:,unique"` // = ConfigFile/..

	// WorkingDirectory string
	// RegularPath      string `gorm:"index:,unique"` // = DotGitPath
	// Location dxo.LocationID `gorm:"index:,unique"`

	Class           dxo.LocationClass
	OwnerRepository dxo.LocalRepositoryID
}

// ListPathFields ...
func (TreeRoot) ListPathFields() []string {
	return []string{"path", "dot_git_path", "working_directory", "regular_path"}
}

// TableName 。。。
func (TreeRoot) TableName() string {
	return getTableName("git_tree_roots")
}

// ComputeRegularPath 。。。
func (inst *TreeRoot) ComputeRegularPath() dxo.RegularPath {
	path := inst.Path
	return dxo.NewRegularPath(path)
}
