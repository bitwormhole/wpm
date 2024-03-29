package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// LocalRepositoryServiceImpl ...
type LocalRepositoryServiceImpl struct {
	markup.Component `id:"LocalRepositoryService"`

	GitLibAgent store.LibAgent `inject:"#git-lib-agent"`

	LocalRepositoryDAO dao.LocalRepositoryDAO `inject:"#LocalRepositoryDAO"`

	UUIDGenService    service.UUIDGenService             `inject:"#UUIDGenService"`
	RepoFinder        service.LocalRepositoryFinder      `inject:"#LocalRepositoryFinder"`
	LrStateLoader     service.LocalRepositoryStateLoader `inject:"#LocalRepositoryStateLoader"`
	FileSystemService service.FileSystemService          `inject:"#FileSystemService"`
	LocationService   service.LocationService            `inject:"#LocationService"`
	ProjectService    service.ProjectService             `inject:"#ProjectService"`
}

func (inst *LocalRepositoryServiceImpl) _Impl() service.LocalRepositoryService {
	return inst
}

func (inst *LocalRepositoryServiceImpl) prepareBeforeWrite(ctx context.Context, o1 *entity.LocalRepository) error {

	err := inst.prepareEntity(ctx, o1)
	if err != nil {
		return err
	}

	inst.LocationService.PutRepository(ctx, o1)

	return nil
}

func (inst *LocalRepositoryServiceImpl) dto2entity(c context.Context, o1 *dto.LocalRepository) (*entity.LocalRepository, error) {

	// convert

	o2 := &entity.LocalRepository{}

	// todo ... fields
	o2.ID = o1.ID
	o2.Name = o1.Name
	o2.DisplayName = o1.DisplayName
	o2.Description = o1.Description

	o2.Path = o1.Path
	o2.RegularPath = o1.RegularPath

	o2.ConfigFile = o1.ConfigFile
	o2.DotGitPath = o1.DotGitPath
	o2.RepositoryPath = o1.RepositoryPath
	o2.WorkingPath = o1.WorkingPath

	o2.Bare = o1.Bare

	return o2, nil
}

func (inst *LocalRepositoryServiceImpl) entity2dto(ctx context.Context, o1 *entity.LocalRepository, opt *service.LocalRepositoryOptions) (*dto.LocalRepository, error) {

	opt = inst.normalizeOptions(opt)

	o2 := &dto.LocalRepository{}
	o2.ID = o1.ID
	o2.UUID = o1.UUID
	o2.CreatedAt = util.NewTime(o1.CreatedAt)
	o2.UpdatedAt = util.NewTime(o1.UpdatedAt)

	o2.Name = o1.Name
	o2.DisplayName = o1.DisplayName
	o2.Description = o1.Description

	o2.ConfigFile = o1.ConfigFile
	o2.DotGitPath = o1.DotGitPath
	o2.RepositoryPath = o1.RepositoryPath
	o2.WorkingPath = o1.WorkingPath

	o2.Path = o1.Path
	o2.RegularPath = o1.RegularPath

	o2.Bare = o1.Bare

	if opt.WithFileState {
		inst.LrStateLoader.LoadState(ctx, o2)
	}

	if opt.WithProjects {
		inst.loadProjects(ctx, o2)
	}
	if opt.WithSubmodules {
		inst.loadSubmodules(ctx, o2)
	}
	if opt.WithWorktrees {
		inst.loadWorktrees(ctx, o2)
	}

	if opt.WithGitStatus {
		// todo ...
	}

	return o2, nil
}

func (inst *LocalRepositoryServiceImpl) normalizeOptions(opt *service.LocalRepositoryOptions) *service.LocalRepositoryOptions {
	if opt == nil {
		opt = &service.LocalRepositoryOptions{}
	}
	return opt
}

// ConvertEntityToDto ...
func (inst *LocalRepositoryServiceImpl) ConvertEntityToDto(ctx context.Context, o1 *entity.LocalRepository) (*dto.LocalRepository, error) {
	// ctx := context.Background()
	return inst.entity2dto(ctx, o1, nil)
}

// ConvertDtoToEntity ...
func (inst *LocalRepositoryServiceImpl) ConvertDtoToEntity(ctx context.Context, o1 *dto.LocalRepository) (*entity.LocalRepository, error) {
	return inst.dto2entity(ctx, o1)
}

func (inst *LocalRepositoryServiceImpl) prepareEntity(ctx context.Context, o1 *entity.LocalRepository) error {

	layout, err := inst.RepoFinder.LocateLayout(ctx, o1.Path)
	if err != nil {
		return err
	}

	dotgit := layout.DotGit()
	config := layout.Config()

	if config == nil {
		return fmt.Errorf("it's not a git repository: " + o1.Path)
	}
	if !config.IsFile() {
		return fmt.Errorf("it's not a git repository: " + o1.Path)
	}

	bare := true // default is bare
	repodir := config.GetParent()

	if dotgit != nil {
		if dotgit.Exists() {
			bare = false
		}
	}

	o1.Bare = bare
	o1.ConfigFile = config.GetPath()
	o1.RepositoryPath = repodir.GetPath()
	o1.Path = repodir.GetPath()
	o1.RegularPath = config.GetPath()
	if !bare {
		o1.DotGitPath = dotgit.GetPath()
		o1.WorkingPath = dotgit.GetParent().GetPath()
	}

	return nil
}

// ListAll ...
func (inst *LocalRepositoryServiceImpl) ListAll(ctx context.Context, opt *service.LocalRepositoryOptions) ([]*dto.LocalRepository, error) {
	src, err := inst.LocalRepositoryDAO.ListAll()
	if err != nil {
		return nil, err
	}
	dst := make([]*dto.LocalRepository, 0)
	for _, o1 := range src {
		o2, err := inst.entity2dto(ctx, o1, opt)
		if err != nil {
			return nil, err
		}
		dst = append(dst, o2)
	}
	return dst, nil
}

// ListByIds ...
func (inst *LocalRepositoryServiceImpl) ListByIds(ctx context.Context, ids []dxo.LocalRepositoryID, opt *service.LocalRepositoryOptions) ([]*dto.LocalRepository, error) {
	src, err := inst.LocalRepositoryDAO.ListByIds(ids)
	if err != nil {
		return nil, err
	}
	dst := make([]*dto.LocalRepository, 0)
	for _, o1 := range src {
		o2, err := inst.entity2dto(ctx, o1, opt)
		if err != nil {
			return nil, err
		}
		dst = append(dst, o2)
	}
	return dst, nil
}

// Find ...
func (inst *LocalRepositoryServiceImpl) Find(ctx context.Context, id dxo.LocalRepositoryID, opt *service.LocalRepositoryOptions) (*dto.LocalRepository, error) {
	o1, err := inst.LocalRepositoryDAO.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(ctx, o1, opt)
}

// FindByPath ...
func (inst *LocalRepositoryServiceImpl) FindByPath(ctx context.Context, path string, opt *service.LocalRepositoryOptions) (*dto.LocalRepository, error) {

	lrdao := inst.LocalRepositoryDAO
	ent, err := lrdao.FindByPath(path)
	if err != nil {
		return nil, err
	}

	return inst.entity2dto(ctx, ent, opt)
}

// FindByName ...
func (inst *LocalRepositoryServiceImpl) FindByName(ctx context.Context, name string, opt *service.LocalRepositoryOptions) (*dto.LocalRepository, error) {
	o1, err := inst.LocalRepositoryDAO.FindByName(name)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(ctx, o1, opt)
}

// InsertOrFetch ...
func (inst *LocalRepositoryServiceImpl) InsertOrFetch(ctx context.Context, o1 *dto.LocalRepository, opt *service.LocalRepositoryOptions) (*dto.LocalRepository, error) {

	o2, err := inst.Insert(ctx, o1)
	if o2 != nil && err == nil {
		return o2, nil
	}

	errlist := &utils.ErrorList{}
	errlist.Append(err)

	paths := make([]string, 0)
	paths = append(paths, o1.ConfigFile)
	paths = append(paths, o1.RepositoryPath)
	paths = append(paths, o1.DotGitPath)
	paths = append(paths, o1.WorkingPath)
	paths = append(paths, o1.Path)

	for _, wantPath := range paths {
		wantPath = strings.TrimSpace(wantPath)
		if wantPath == "" {
			continue
		}
		o2, err := inst.FindByPath(ctx, wantPath, opt)
		if o2 != nil && err == nil {
			return o2, nil
		}
		errlist.Append(err)
	}

	return nil, errlist.First()
}

// Insert ...
func (inst *LocalRepositoryServiceImpl) Insert(ctx context.Context, o1 *dto.LocalRepository) (*dto.LocalRepository, error) {

	o2, err := inst.dto2entity(ctx, o1)
	if err != nil {
		return nil, err
	}

	err = inst.prepareBeforeWrite(ctx, o2)
	if err != nil {
		return nil, err
	}

	o3, err := inst.LocalRepositoryDAO.Insert(o2)
	if err != nil {
		return nil, err
	}

	return inst.entity2dto(ctx, o3, nil)
}

// Update ...
func (inst *LocalRepositoryServiceImpl) Update(ctx context.Context, id dxo.LocalRepositoryID, o1 *dto.LocalRepository) (*dto.LocalRepository, error) {

	o2, err := inst.dto2entity(ctx, o1)
	if err != nil {
		return nil, err
	}

	err = inst.prepareBeforeWrite(ctx, o2)
	if err != nil {
		return nil, err
	}

	o3, err := inst.LocalRepositoryDAO.Update(id, o2)
	if err != nil {
		return nil, err
	}

	return inst.entity2dto(ctx, o3, nil)
}

// Remove ...
func (inst *LocalRepositoryServiceImpl) Remove(ctx context.Context, id dxo.LocalRepositoryID) error {
	return inst.LocalRepositoryDAO.Remove(id)
}

func (inst *LocalRepositoryServiceImpl) loadWorktrees(ctx context.Context, repo *dto.LocalRepository) error {

	lib, err := inst.GitLibAgent.GetLib()
	if err != nil {
		return err
	}

	loader := &myWorktreesLoader{lib}
	path := repo.Path
	owner := repo.ID
	list, err := loader.Load(path, owner)
	if err != nil {
		return err
	}

	repo.Worktrees = list
	return nil
}

func (inst *LocalRepositoryServiceImpl) loadSubmodules(ctx context.Context, repo *dto.LocalRepository) error {

	lib, err := inst.GitLibAgent.GetLib()
	if err != nil {
		return err
	}

	loader := &mySubmodulesLoader{lib}
	path := repo.Path
	owner := repo.ID
	list, err := loader.Load(path, owner)
	if err != nil {
		return err
	}

	repo.Submodules = list
	return nil
}

func (inst *LocalRepositoryServiceImpl) loadProjects(ctx context.Context, repo *dto.LocalRepository) error {
	owner := repo.ID
	src, err := inst.ProjectService.FindByOwnerRepository(ctx, owner, nil)
	if err != nil {
		return err
	}
	repo.Projects = src
	return nil
}
