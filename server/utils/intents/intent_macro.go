package intents

// executable
const (
	ExecutablePrefix = "executable."

	ExecutableName  = ExecutablePrefix + "name"
	ExecutablePath  = ExecutablePrefix + "path"
	ExecutableTitle = ExecutablePrefix + "title"
)

// repository
const (
	RepositoryPrefix = "repository."

	RepositoryName       = RepositoryPrefix + "name"
	RepositoryPath       = RepositoryPrefix + "path"
	RepositoryWorkDir    = RepositoryPrefix + "wkdir"
	RepositoryConfigFile = RepositoryPrefix + "config-file-path"
	RepositoryDotGitPath = RepositoryPrefix + "dot-git-path"
	RepositoryCoreDir    = RepositoryPrefix + "git-core-dir"
)

// worktree
const (
	WorktreePrefix     = "worktree."
	WorktreeName       = WorktreePrefix + "name"
	WorktreeDotGitPath = WorktreePrefix + "dot-git-path"
	WorktreeWorkDir    = WorktreePrefix + "wkdir"
)

// project
const (
	ProjectPrefix   = "project."
	ProjectName     = ProjectPrefix + "name"
	ProjectFullPath = ProjectPrefix + "fullpath"
)

// location
const (
	LocationPrefix = "location."
	LocationName   = LocationPrefix + "name"
	LocationPath   = LocationPrefix + "path"
	LocationType   = LocationPrefix + "type"
)

// ListMacroNames ...
func ListMacroNames() []string {

	list := make([]string, 0)

	list = append(list, ExecutableName)
	list = append(list, ExecutablePath)
	list = append(list, ExecutableTitle)

	list = append(list, LocationName)
	list = append(list, LocationPath)
	list = append(list, LocationType)

	list = append(list, ProjectName)
	list = append(list, ProjectFullPath)

	list = append(list, RepositoryName)
	list = append(list, RepositoryPath)
	list = append(list, RepositoryWorkDir)
	list = append(list, RepositoryConfigFile)
	list = append(list, RepositoryDotGitPath)
	list = append(list, RepositoryCoreDir)

	list = append(list, WorktreeName)
	list = append(list, WorktreeDotGitPath)
	list = append(list, WorktreeWorkDir)

	return list
}
