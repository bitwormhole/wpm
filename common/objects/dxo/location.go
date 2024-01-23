package dxo

// LocationClass 表示 Location 的各种类型
type LocationClass string

// 定义 Location 的各种类型
const (
	LocationGitWorktree   LocationClass = "git.worktree"   // for Git Worktree
	LocationGitSubmodule  LocationClass = "git.submodule"  // for Git Submodule
	LocationGitConfig     LocationClass = "git.config"     // for Git Repository Config
	LocationGitRepository LocationClass = "git.repository" // for LocalRepository
	LocationProject       LocationClass = "project"        // the project config file or dir
	LocationFolder        LocationClass = "folder"
	LocationExecutable    LocationClass = "executable"
)

func (value LocationClass) String() string {
	return string(value)
}
