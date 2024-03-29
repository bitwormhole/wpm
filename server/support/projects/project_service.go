package projects

import (
	"context"
	"fmt"
	"strings"

	"bitwormhole.com/starter/afs"
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

// ProjectServiceImpl ...
type ProjectServiceImpl struct {
	markup.Component `id:"ProjectService"`

	UUIDGenService     service.UUIDGenService         `inject:"#UUIDGenService"`
	LocalRepoService   service.LocalRepositoryService `inject:"#LocalRepositoryService"`
	ProjectTypeService service.ContentTypeService     `inject:"#ContentTypeService"`
	FileSystemService  service.FileSystemService      `inject:"#FileSystemService"`
	LocationService    service.LocationService        `inject:"#LocationService"`

	ProjectDAO   dao.ProjectDAO         `inject:"#ProjectDAO"`
	LocalRepoDAO dao.LocalRepositoryDAO `inject:"#LocalRepositoryDAO"`

	GitLibAgent store.LibAgent `inject:"#git-lib-agent"`
}

func (inst *ProjectServiceImpl) _Impl() service.ProjectService {
	return inst
}

func (inst *ProjectServiceImpl) dto2entity(c context.Context, o1 *dto.Project) (*entity.Project, error) {

	err := inst.ProjectTypeService.LocateProject(c, o1, "")
	if err != nil {
		return nil, err
	}

	// convert

	o2 := &entity.Project{}
	o2.ID = o1.ID

	o2.Description = o1.Description
	o2.IsDir = o1.IsDir
	o2.IsFile = o1.IsFile
	o2.Name = o1.Name
	o2.OwnerRepository = o1.OwnerRepository
	o2.PathInWorktree = o1.PathInWorktree
	o2.ProjectDir = o1.ProjectDir
	o2.ConfigFileName = o1.ConfigFileName

	o2.Path = o1.Path
	o2.RegularPath = o1.RegularPath
	o2.Type = o1.Type

	return o2, nil
}

func (inst *ProjectServiceImpl) entity2dto(o1 *entity.Project, opt *service.ProjectOptions) (*dto.Project, error) {

	if opt == nil {
		opt = &service.ProjectOptions{}
	}

	o2 := &dto.Project{}
	o2.ID = o1.ID
	o2.UUID = o1.UUID
	o2.CreatedAt = util.NewTime(o1.CreatedAt)
	o2.UpdatedAt = util.NewTime(o1.UpdatedAt)

	o2.Description = o1.Description
	o2.IsDir = o1.IsDir
	o2.IsFile = o1.IsFile
	o2.Name = o1.Name
	o2.OwnerRepository = o1.OwnerRepository

	o2.PathInWorktree = o1.PathInWorktree
	o2.ProjectDir = o1.ProjectDir
	o2.ConfigFileName = o1.ConfigFileName

	o2.Path = o1.Path
	o2.RegularPath = o1.RegularPath
	o2.Type = o1.Type

	if opt.WithFileState {
		o2.State = inst.checkProjectFileState(o1)
	}

	return o2, nil
}

func (inst *ProjectServiceImpl) entity2dtoList(src []*entity.Project, opt *service.ProjectOptions) ([]*dto.Project, error) {
	dst := make([]*dto.Project, 0)
	for _, item1 := range src {
		item2, err := inst.entity2dto(item1, opt)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}

func (inst *ProjectServiceImpl) checkProjectFileState(o1 *entity.Project) dxo.FileState {
	path := inst.FileSystemService.Path(o1.Path)
	if path.Exists() {
		return dxo.FileStateReady
	}
	return dxo.FileStateOffline
}

// ListAll ...
func (inst *ProjectServiceImpl) ListAll(ctx context.Context, options *service.ProjectOptions) ([]*dto.Project, error) {
	src, err := inst.ProjectDAO.ListAll()
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(src, options)
}

// ListByIds ...
func (inst *ProjectServiceImpl) ListByIds(ctx context.Context, ids []dxo.ProjectID, options *service.ProjectOptions) ([]*dto.Project, error) {
	src, err := inst.ProjectDAO.ListByIds(ids)
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(src, options)
}

// Find ...
func (inst *ProjectServiceImpl) Find(ctx context.Context, id dxo.ProjectID, options *service.ProjectOptions) (*dto.Project, error) {
	p, err := inst.ProjectDAO.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(p, options)
}

// FindByOwnerRepository ...
func (inst *ProjectServiceImpl) FindByOwnerRepository(ctx context.Context, id dxo.LocalRepositoryID, options *service.ProjectOptions) ([]*dto.Project, error) {
	items1, err := inst.ProjectDAO.FindByOwnerRepository(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(items1, options)
}

// Locate ...
func (inst *ProjectServiceImpl) Locate(ctx context.Context, o1 *dto.Project, options *service.ProjectOptions) (*dto.Project, error) {
	o2, err := inst.dto2entity(ctx, o1)
	if err != nil {
		return nil, err
	}
	err = (&myProjectRepoPathFinder{ser: inst}).locate(o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o2, options)
}

// Insert ...
func (inst *ProjectServiceImpl) Insert(ctx context.Context, o1 *dto.Project) (*dto.Project, error) {

	o2, err := inst.dto2entity(ctx, o1)
	if err != nil {
		return nil, err
	}

	err = (&myProjectRepoPathFinder{ser: inst}).locate(o2)
	if err != nil {
		return nil, err
	}

	inst.LocationService.PutProject(ctx, o2)

	o3, err := inst.ProjectDAO.Insert(o2)
	if err != nil {
		return nil, err
	}

	return inst.entity2dto(o3, nil)
}

// InsertOrFetch ...
func (inst *ProjectServiceImpl) InsertOrFetch(ctx context.Context, o1 *dto.Project, opt *service.ProjectOptions) (*dto.Project, error) {

	o2, err := inst.Insert(ctx, o1)
	if o2 != nil && err == nil {
		return o2, nil
	}

	errlist := &utils.ErrorList{}
	errlist.Append(err)

	paths := make([]string, 0)
	paths = append(paths, o1.Path)
	paths = append(paths, o1.ProjectDir)

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

// Update ...
func (inst *ProjectServiceImpl) Update(ctx context.Context, id dxo.ProjectID, o1 *dto.Project) (*dto.Project, error) {

	o2, err := inst.dto2entity(ctx, o1)
	if err != nil {
		return nil, err
	}

	err = (&myProjectRepoPathFinder{ser: inst}).locate(o2)
	if err != nil {
		return nil, err
	}

	inst.LocationService.PutProject(ctx, o2)

	o3, err := inst.ProjectDAO.Update(id, o2)
	if err != nil {
		return nil, err
	}

	return inst.entity2dto(o3, nil)
}

// Remove ...
func (inst *ProjectServiceImpl) Remove(ctx context.Context, id dxo.ProjectID) error {
	return inst.ProjectDAO.Remove(id)
}

// FindByPath ...
func (inst *ProjectServiceImpl) FindByPath(ctx context.Context, path string, options *service.ProjectOptions) (*dto.Project, error) {
	pdao := inst.ProjectDAO
	ent, err := pdao.FindByPath(path)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(ent, options)
}

////////////////////////////////////////////////////////////////////////////////

type myProjectRepoPathFinder struct {
	ctx context.Context
	ser *ProjectServiceImpl
	gl  store.Lib
}

func (inst *myProjectRepoPathFinder) getGitLib() (store.Lib, error) {
	gl := inst.gl
	if gl != nil {
		return gl, nil
	}
	gl, err := inst.ser.GitLibAgent.GetLib()
	if err != nil {
		return nil, err
	}
	return gl, nil
}

func (inst *myProjectRepoPathFinder) locate(p *entity.Project) error {
	fullpath := p.Path
	gl, err := inst.getGitLib()
	if err != nil {
		return err
	}
	wd := gl.FS().NewPath(fullpath)
	if !wd.Exists() {
		return fmt.Errorf("file or dir is not exists: " + fullpath)
	}

	p.IsFile = wd.IsFile()
	p.IsDir = wd.IsDirectory()
	p.RegularPath = fullpath

	layout, err := gl.RepositoryLocator().Locate(wd)
	if err == nil {
		return inst.locateProjectWithRepository(p, gl, layout)
	}
	return inst.locateProjectWithoutRepository(p)
}

func (inst *myProjectRepoPathFinder) locateProjectWithRepository(p *entity.Project, gl store.Lib, layout store.RepositoryLayout) error {
	ent, err := inst.openOrCreateRepository(gl, layout)
	if err != nil {
		return err
	}
	p.PathInWorktree = inst.computePathInWorktree(layout)
	p.OwnerRepository = ent.ID
	return nil
}

func (inst *myProjectRepoPathFinder) locateProjectWithoutRepository(p *entity.Project) error {
	p.PathInWorktree = ""
	p.OwnerRepository = 0
	return nil
}

func (inst *myProjectRepoPathFinder) computePathInWorktree(layout store.RepositoryLayout) string {
	s1 := inst.getPathString(layout.Workspace())
	s2 := inst.getPathString(layout.WD())
	len1 := len(s1)
	len2 := len(s2)
	if len1 >= len2 {
		return ""
	}
	s := s2[len1:]
	for strings.HasPrefix(s, "/") || strings.HasPrefix(s, "\\") {
		s = s[1:]
	}
	return s
}

func (inst *myProjectRepoPathFinder) openOrCreateRepository(gl store.Lib, layout store.RepositoryLayout) (*entity.LocalRepository, error) {

	dotgit := layout.DotGit()
	ent, err := inst.ser.LocalRepoDAO.FindByDotGit(dotgit.GetPath())
	if err == nil {
		// found
		return ent, nil
	}

	// make new repository
	name := inst.getPathName(layout.Workspace())
	o := &dto.LocalRepository{}

	o.DotGitPath = inst.getPathString(layout.DotGit())
	o.RepositoryPath = inst.getPathString(layout.Repository())
	o.WorkingPath = inst.getPathString(layout.Workspace())
	o.Path = inst.getPathString(layout.Workspace())

	o.Name = name
	o.DisplayName = name
	o.Description = ""

	o, err = inst.ser.LocalRepoService.Insert(inst.ctx, o)
	if err != nil {
		return nil, err
	}

	id := o.ID
	return inst.ser.LocalRepoDAO.Find(id)
}

func (inst *myProjectRepoPathFinder) getPathString(p afs.Path) string {
	if p == nil {
		return ""
	}
	return p.GetPath()
}

func (inst *myProjectRepoPathFinder) getPathName(p afs.Path) string {
	if p == nil {
		return ""
	}
	return p.GetName()
}

////////////////////////////////////////////////////////////////////////////////
