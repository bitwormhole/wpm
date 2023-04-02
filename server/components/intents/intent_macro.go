package intents

import (
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/wpm/server/utils"
)

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

////////////////////////////////////////////////////////////////////////////////

// MacroResolver ...
type MacroResolver struct {
	properties map[string]string
	errlist    utils.ErrorList
	vkey       string
}

// Init ...
func (inst *MacroResolver) Init(p map[string]string) {
	if p == nil {
		p = make(map[string]string)
	}
	inst.properties = p
	inst.vkey = "abdd7cf871de41581402ddf2d6622ecb1dacbc38"
}

func (inst *MacroResolver) Error() error {
	return inst.errlist.First()
}

// ResolveStringMap ...
func (inst *MacroResolver) ResolveStringMap(src map[string]string) map[string]string {
	dst := make(map[string]string)
	for k, v := range src {
		k = inst.ResolveString(k)
		v = inst.ResolveString(v)
		dst[k] = v
	}
	return dst
}

// ResolveStringArray ...
func (inst *MacroResolver) ResolveStringArray(src []string) []string {
	dst := make([]string, 0)
	for _, item1 := range src {
		item2 := inst.ResolveString(item1)
		dst = append(dst, item2)
	}
	return dst
}

// ResolveString ...
func (inst *MacroResolver) ResolveString(s string) string {
	const (
		t1 = "${"
		t2 = "}"
	)
	key := inst.vkey
	props := collection.CreateProperties()
	props.Import(inst.properties)
	props.SetProperty(key, s)
	err := collection.ResolvePropertiesVarWithTokens(t1, t2, props)
	inst.errlist.Append(err)
	return props.GetProperty(key, s)
}
