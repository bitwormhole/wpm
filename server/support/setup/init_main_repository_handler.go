package setup

import (
	"fmt"
	"strings"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type myInitMainRepositoryHandler struct {
	context *Context
}

func (inst *myInitMainRepositoryHandler) Run() error {

	adSer := inst.context.AppDataService
	fsSer := inst.context.FileSystemService
	mainRepoPath := adSer.GetMainRepositoryPath()
	mainRepo := fsSer.Path(mainRepoPath)
	dotgit := mainRepo

	name := dotgit.GetName()
	if name != ".git" {
		dotgit = dotgit.GetChild(".git")
	}

	head := "ref: refs/heads/main\n"
	description := "WPM Main Repository"
	config := inst.makeConfigText()

	list := make([]*dto.File, 0)
	list = inst.addDir(list, dotgit, "objects/info")
	list = inst.addDir(list, dotgit, "objects/pack")
	list = inst.addDir(list, dotgit, "refs/heads")
	list = inst.addDir(list, dotgit, "refs/tags")
	list = inst.addDir(list, dotgit, "hooks")
	list = inst.addDir(list, dotgit, "info")
	list = inst.addFile(list, dotgit, "HEAD", head)
	list = inst.addFile(list, dotgit, "description", description)
	list = inst.addFile(list, dotgit, "config", config)

	return inst.tryMakeAll(list, fsSer.FS())
}

func (inst *myInitMainRepositoryHandler) makeConfigText() string {

	const nl = "\n"
	const tab = "\t"

	builder := strings.Builder{}
	builder.WriteString("[core]" + nl)
	builder.WriteString(tab + "repositoryformatversion = 0" + nl)
	builder.WriteString(tab + "filemode = false" + nl)
	builder.WriteString(tab + "bare = false" + nl)
	builder.WriteString(tab + "logallrefupdates = true" + nl)
	builder.WriteString(tab + "symlinks = false" + nl)
	builder.WriteString(tab + "ignorecase = true" + nl)

	return builder.String()
}

func (inst *myInitMainRepositoryHandler) addFile(list []*dto.File, parent afs.Path, name string, text string) []*dto.File {

	child := parent.GetChild(name)

	item := &dto.File{}
	item.Name = child.GetName()
	item.Path = child.GetPath()
	item.IsFile = true
	item.Type = text // 用 Type 字段代替 text 承载文本内容

	list = append(list, item)
	return list
}

func (inst *myInitMainRepositoryHandler) addDir(list []*dto.File, parent afs.Path, name string) []*dto.File {

	child := parent.GetChild(name)

	item := &dto.File{}
	item.Name = child.GetName()
	item.Path = child.GetPath()
	item.IsDir = true

	list = append(list, item)
	return list
}

func (inst *myInitMainRepositoryHandler) tryMakeAll(list []*dto.File, fs afs.FS) error {
	el := utils.ErrorList{}
	for _, item := range list {
		if item.IsDir {
			err := inst.makeDir(item, fs)
			el.Append(err)
		} else if item.IsFile {
			err := inst.makeFile(item, fs)
			el.Append(err)
		}
	}
	return el.First()
}

func (inst *myInitMainRepositoryHandler) makeFile(item *dto.File, fs afs.FS) error {
	file := fs.NewPath(item.Path)
	if file.IsFile() {
		return nil // skip
	}
	if file.IsDirectory() {
		return fmt.Errorf("want a file, have a dir, path=" + item.Path)
	}
	text := item.Type
	opt := &afs.Options{Create: true, Mkdirs: true}
	return file.GetIO().WriteText(text, opt)
}

func (inst *myInitMainRepositoryHandler) makeDir(item *dto.File, fs afs.FS) error {
	dir := fs.NewPath(item.Path)
	if dir.IsFile() {
		return fmt.Errorf("want a dir, have a file, path=" + item.Path)
	}
	if dir.IsDirectory() {
		return nil // skip
	}
	opt := &afs.Options{Create: true, Mkdirs: true}
	return dir.Mkdirs(opt)
}
