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

	UUIDGenService     service.UUIDGenService        `inject:"#UUIDGenService"`
	RepoFinder         service.LocalRepositoryFinder `inject:"#LocalRepositoryFinder"`
	LocalRepositoryDAO dao.LocalRepositoryDAO        `inject:"#LocalRepositoryDAO"`
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

func (inst *LocalRepositoryServiceImpl) entity2dto(o1 *entity.LocalRepository) (*dto.LocalRepository, error) {

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
	o2.State = "todo..."

	return o2, nil
}

// ConvertEntityToDto ...
func (inst *LocalRepositoryServiceImpl) ConvertEntityToDto(o1 *entity.LocalRepository) (*dto.LocalRepository, error) {
	return inst.entity2dto(o1)
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
func (inst *LocalRepositoryServiceImpl) ListAll(ctx context.Context) ([]*dto.LocalRepository, error) {
	src, err := inst.LocalRepositoryDAO.ListAll()
	if err != nil {
		return nil, err
	}
	dst := make([]*dto.LocalRepository, 0)
	for _, o1 := range src {
		o2, err := inst.entity2dto(o1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, o2)
	}
	return dst, nil
}

// Find ...
func (inst *LocalRepositoryServiceImpl) Find(ctx context.Context, id dxo.LocalRepositoryID) (*dto.LocalRepository, error) {
	o1, err := inst.LocalRepositoryDAO.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1)
}

// FindByName ...
func (inst *LocalRepositoryServiceImpl) FindByName(ctx context.Context, name string) (*dto.LocalRepository, error) {
	o1, err := inst.LocalRepositoryDAO.FindByName(name)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1)
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
	return inst.entity2dto(o3)
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
	return inst.entity2dto(o3)
}

// Remove ...
func (inst *LocalRepositoryServiceImpl) Remove(ctx context.Context, id dxo.LocalRepositoryID) error {
	return inst.LocalRepositoryDAO.Remove(id)
}
