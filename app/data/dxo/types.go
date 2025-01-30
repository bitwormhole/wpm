package dxo

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/mimetypes"
)

// URI ...
type URI = afs.URI

// MediaType ...
type MediaType = mimetypes.Type

// Category 表示 Location 的类型
type Category string

// 定义各种 Location 的类型
const (
	CategoryFile              Category = "file"
	CategoryFolder            Category = "folder"
	CategoryGitRepository     Category = "git_repository"      // git 本地仓库
	CategoryGitRemote         Category = "git_remote"          // git 远程仓库
	CategoryGitWorkingDir     Category = "git_working_dir"     // git 默认工作目录
	CategoryGitSubmoduleInner Category = "git_submodule_inner" // git 子模块内部目录
	CategoryGitSubmoduleWkdir Category = "git_submodule_wkdir" // git 子模块工作目录
	CategoryGitWorktreeInner  Category = "git_worktree_inner"  // git 工作树内部
	CategoryGitWorktreeWkdir  Category = "git_worktree_wkdir"  // git 工作树目录
	CategoryExecutable        Category = "executable"
	CategoryProject           Category = "project"
	CategoryMedia             Category = "media"
)
