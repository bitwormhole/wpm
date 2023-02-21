package implservice

import (
	"context"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// LocalRepositoryServiceImpl ...
type LocalRepositoryServiceImpl struct {
	markup.Component `id:"LocalRepositoryService"`

	LocalRepositoryDAO dao.LocalRepositoryDAO             `inject:"#LocalRepositoryDAO"`
	UUIDGenService     service.UUIDGenService             `inject:"#UUIDGenService"`
	RepoFinder         service.LocalRepositoryFinder      `inject:"#LocalRepositoryFinder"`
	LrStateLoader      service.LocalRepositoryStateLoader `inject:"#LocalRepositoryStateLoader"`
}

func (inst *LocalRepositoryServiceImpl) _Impl() service.LocalRepositoryService {
	return inst
}

func (inst *LocalRepositoryServiceImpl) dto2entity(o1 *dto.LocalRepository) (*entity.LocalRepository, error) {

	o2 := &entity.LocalRepository{}

	// todo ... fields
	o2.ID = o1.ID
	o2.Name = o1.Name
	o2.DisplayName = o1.DisplayName
	o2.Description = o1.Description

	o2.Path = o1.Path
	o2.DotGitPath = o1.DotGitPath
	o2.RepositoryPath = o1.RepositoryPath
	o2.WorkingPath = o1.WorkingPath

	return o2, nil
}

func (inst *LocalRepositoryServiceImpl) entity2dto(ctx context.Context, o1 *entity.LocalRepository, opt *service.LocalRepositoryOptions) (*dto.LocalRepository, error) {

	opt = inst.normalizeOptions(opt)
	o2 := &dto.LocalRepository{}

	// todo ... fields
	o2.ID = o1.ID
	o2.Name = o1.Name
	o2.DisplayName = o1.DisplayName
	o2.Description = o1.Description

	o2.Path = o1.Path
	o2.DotGitPath = o1.DotGitPath
	o2.RepositoryPath = o1.RepositoryPath
	o2.WorkingPath = o1.WorkingPath

	// o2.Ready =o1

	if opt.WithFileState {
		inst.LrStateLoader.LoadState(ctx, o2)
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
func (inst *LocalRepositoryServiceImpl) ConvertEntityToDto(o1 *entity.LocalRepository) (*dto.LocalRepository, error) {
	ctx := context.Background()
	return inst.entity2dto(ctx, o1, nil)
}

// ConvertDtoToEntity ...
func (inst *LocalRepositoryServiceImpl) ConvertDtoToEntity(o1 *dto.LocalRepository) (*entity.LocalRepository, error) {
	return inst.dto2entity(o1)
}

func (inst *LocalRepositoryServiceImpl) prepareEntity(ctx context.Context, o1 *entity.LocalRepository) error {

	deffs := fs.Default()
	path := o1.Path
	o2, err := inst.RepoFinder.Locate(ctx, path)
	if err != nil {
		return err
	}

	dotgit := deffs.GetPath(o2.Path)
	config := dotgit.GetChild("config")
	working := dotgit.Parent()

	o1.Path = dotgit.String()
	o1.DotGitPath = dotgit.String()
	o1.ConfigFile = config.Path()
	o1.WorkingPath = working.Path()
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

// FindByName ...
func (inst *LocalRepositoryServiceImpl) FindByName(ctx context.Context, name string, opt *service.LocalRepositoryOptions) (*dto.LocalRepository, error) {
	o1, err := inst.LocalRepositoryDAO.FindByName(name)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(ctx, o1, opt)
}

// Insert ...
func (inst *LocalRepositoryServiceImpl) Insert(ctx context.Context, o1 *dto.LocalRepository) (*dto.LocalRepository, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}

	err = inst.prepareEntity(ctx, o2)
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
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}

	err = inst.prepareEntity(ctx, o2)
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
