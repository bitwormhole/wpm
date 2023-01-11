package implservice

import (
	"context"
	"errors"
	"sort"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// RepositorySearchServiceImpl ...
type RepositorySearchServiceImpl struct {
	markup.Component `id:"RepositorySearchService"`

	LocateService service.RepositoryLocateService `inject:"#RepositoryLocateService"`
}

func (inst *RepositorySearchServiceImpl) _Impl() service.RepositorySearchService {
	return inst
}

// Search ...
func (inst *RepositorySearchServiceImpl) Search(ctx context.Context, path fs.Path) ([]*dto.LocalRepository, error) {
	task := myRepositorySearchTask{context: ctx, parent: inst}
	task.init()
	err := task.searchInDir(path, 0)
	if err != nil {
		return nil, err
	}
	return task.results, nil
}

////////////////////////////////////////////////////////////////////////////////

type myRepositorySearchTask struct {
	context  context.Context
	parent   *RepositorySearchServiceImpl
	maxDepth int
	results  []*dto.LocalRepository
}

func (inst *myRepositorySearchTask) init() {
	inst.maxDepth = 32
	inst.results = make([]*dto.LocalRepository, 0)
}

func (inst *myRepositorySearchTask) searchInDir(dir fs.Path, depth int) error {

	if depth > inst.maxDepth {
		return errors.New("the path is too deep, path=" + dir.String())
	}

	dotgit := dir.GetChild(".git")
	if dotgit.Exists() {
		err := inst.handleDotGit(dotgit)
		if err != nil {
			vlog.Warn(err)
		}
		return nil
	}

	if !dir.IsDir() {
		return nil
	}

	// as dir
	// items
	namelist := dir.ListNames()
	sort.Strings(namelist)
	for _, name := range namelist {
		child := dir.GetChild(name)
		if child.IsDir() {
			err := inst.searchInDir(child, depth+1)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (inst *myRepositorySearchTask) handleDotGit(p fs.Path) error {
	ctx := inst.context
	repo, err := inst.parent.LocateService.Locate(ctx, p)
	if err != nil {
		return err
	}
	inst.results = append(inst.results, repo)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
