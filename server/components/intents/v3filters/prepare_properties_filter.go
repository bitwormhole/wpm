package v3filters

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/components/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// PreparePropertiesFilter ...
type PreparePropertiesFilter struct {
	markup.Component `class:"wpm-intent-filter"`
}

func (inst *PreparePropertiesFilter) _Impl() (intents.FilterRegistry, intents.Filter) {
	return inst, inst
}

// GetFilterRegistrationList ...
func (inst *PreparePropertiesFilter) GetFilterRegistrationList() []*intents.FilterRegistration {
	fr := &intents.FilterRegistration{
		Enabled: true,
		Order:   orderPrepareProperties,
		Filter:  inst,
	}
	return []*intents.FilterRegistration{fr}
}

// Handle ...
func (inst *PreparePropertiesFilter) Handle(c context.Context, i *dto.Intent, next intents.FilterChain) error {

	src := i.Properties
	dst := make(map[string]string)

	inst.forExecutable(i.Executable, dst)
	inst.forFile(i.File, dst)
	inst.forFolder(i.Folder, dst)
	inst.forProject(i.Project, dst)
	inst.forRepository(i.Repository, dst)
	inst.forSubmodule(i.Submodule, dst)
	inst.forWorktree(i.Worktree, dst)
	inst.forWeb(i.Web, dst)

	for k, v := range src {
		dst[k] = v
	}

	i.Properties = dst
	return next.Handle(c, i)
}

func (inst *PreparePropertiesFilter) forExecutable(o *dto.Executable, dst map[string]string) {
	if o == nil {
		return
	}
	dst[intents.ExecutableName] = o.Name
	dst[intents.ExecutablePath] = o.Path
	dst[intents.ExecutableTitle] = o.Title
}

func (inst *PreparePropertiesFilter) forRepository(o *dto.LocalRepository, dst map[string]string) {
	if o == nil {
		return
	}
	dst[intents.RepositoryConfigFile] = o.ConfigFile
	dst[intents.RepositoryCoreDir] = o.RepositoryPath
	dst[intents.RepositoryDotGitPath] = o.DotGitPath
	dst[intents.RepositoryName] = o.Name
	dst[intents.RepositoryPath] = o.Path
	dst[intents.RepositoryWorkDir] = o.WorkingPath
}

func (inst *PreparePropertiesFilter) forWorktree(o *dto.Worktree, dst map[string]string) {
	if o == nil {
		return
	}
	dst[intents.WorktreeName] = o.Name
	dst[intents.WorktreeWorkDir] = o.WorkingDir
	dst[intents.WorktreeDotGitPath] = o.DotGitPath
}

func (inst *PreparePropertiesFilter) forSubmodule(o *dto.Submodule, dst map[string]string) {
	if o == nil {
		return
	}
	dst[intents.SubmoduleName] = o.Name
	dst[intents.SubmoduleDotGit] = o.DotGitPath
	dst[intents.SubmoduleWorkDir] = o.WorkingDir
}

func (inst *PreparePropertiesFilter) forProject(o *dto.Project, dst map[string]string) {
	if o == nil {
		return
	}
	dst[intents.ProjectName] = o.Name
	dst[intents.ProjectFullPath] = o.Path
}

func (inst *PreparePropertiesFilter) forFile(o *dto.File, dst map[string]string) {
	if o == nil {
		return
	}
	dst[intents.FileName] = o.Name
	dst[intents.FilePath] = o.Path
}

func (inst *PreparePropertiesFilter) forFolder(o *dto.File, dst map[string]string) {
	if o == nil {
		return
	}
	dst[intents.FolderName] = o.Name
	dst[intents.FolderPath] = o.Path
}

func (inst *PreparePropertiesFilter) forWeb(o *dto.WebRequest, dst map[string]string) {
	if o == nil {
		return
	}
	dst[intents.WebMethod] = o.Method
	dst[intents.WebURL] = o.URL
}
