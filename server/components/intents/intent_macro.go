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

// file
const (
	FilePrefix = "file."
	FileName   = FilePrefix + "name"
	FilePath   = FilePrefix + "path"
	FileType   = FilePrefix + "type"
)

// folder
const (
	FolderPrefix = "folder."
	FolderName   = FolderPrefix + "name"
	FolderPath   = FolderPrefix + "path"
	FolderType   = FolderPrefix + "type"
)

// submodule
const (
	SubmodulePrefix  = "submodule."
	SubmoduleName    = SubmodulePrefix + "name"
	SubmoduleWorkDir = SubmodulePrefix + "wkdir"
	SubmoduleDotGit  = SubmodulePrefix + "dotgit"
)

// ListMacroNames ...
func ListMacroNames() []string {

	list := make([]string, 0)

	list = append(list, ExecutableName)
	list = append(list, ExecutablePath)
	list = append(list, ExecutableTitle)

	list = append(list, FileName)
	list = append(list, FilePath)
	list = append(list, FileType)

	list = append(list, FolderName)
	list = append(list, FolderPath)
	list = append(list, FolderType)

	list = append(list, SubmoduleName)
	list = append(list, SubmoduleDotGit)
	list = append(list, SubmoduleWorkDir)

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
