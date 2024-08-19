package entity

import "github.com/bitwormhole/wpm/common/objects/dxo"

// GitRepositoryContainer （directorry） 表示一个包含了若干 git 仓库的文件夹
type GitRepositoryContainer struct {
	ID dxo.GitRepositoryContainerID `gorm:"primaryKey"`

	Base

	Path     string
	Location dxo.LocationID `gorm:"unique"`
}

// TableName 。。。
func (GitRepositoryContainer) TableName() string {
	return getTableName("git_repo_container_dirs")
}
