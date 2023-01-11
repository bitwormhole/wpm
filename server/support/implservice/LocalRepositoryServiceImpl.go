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

// RepositoryServiceImpl ...
type RepositoryServiceImpl struct {
	markup.Component `id:"RepositoryService"`

	UUIDGenService     service.UUIDGenService          `inject:"#UUIDGenService"`
	LocateService      service.RepositoryLocateService `inject:"#RepositoryLocateService"`
	LocalRepositoryDAO dao.LocalRepositoryDAO          `inject:"#RepositoryDAO"`
}

func (inst *RepositoryServiceImpl) _Impl() service.LocalRepositoryService {
	return inst
}

func (inst *RepositoryServiceImpl) dto2entity(o1 *dto.LocalRepository) (*entity.LocalRepository, error) {

	o2 := &entity.LocalRepository{}

	// todo ... fields
	o2.ID = o1.ID
	o2.Name = o1.Name
	o2.DisplayName = o1.DisplayName
	o2.Description = o1.Description
	o2.IconURL = o1.IconURL
	o2.Path = o1.Path

	return o2, nil
}

func (inst *RepositoryServiceImpl) entity2dto(o1 *entity.LocalRepository) (*dto.LocalRepository, error) {

	o2 := &dto.LocalRepository{}

	// todo ... fields
	o2.ID = o1.ID
	o2.Path = o1.Path
	o2.Name = o1.Name
	o2.DisplayName = o1.DisplayName
	o2.Description = o1.Description
	o2.IconURL = o1.IconURL
	// o2.Ready =o1

	return o2, nil
}

func (inst *RepositoryServiceImpl) prepareEntity(ctx context.Context, o1 *entity.LocalRepository) error {

	deffs := fs.Default()
	path := deffs.GetPath(o1.Path)
	o2, err := inst.LocateService.Locate(ctx, path)
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
func (inst *RepositoryServiceImpl) ListAll(ctx context.Context) ([]*dto.LocalRepository, error) {
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
func (inst *RepositoryServiceImpl) Find(ctx context.Context, id dxo.LocalRepositoryID) (*dto.LocalRepository, error) {
	o1, err := inst.LocalRepositoryDAO.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1)
}

// FindByName ...
func (inst *RepositoryServiceImpl) FindByName(ctx context.Context, name string) (*dto.LocalRepository, error) {
	o1, err := inst.LocalRepositoryDAO.FindByName(name)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1)
}

// Insert ...
func (inst *RepositoryServiceImpl) Insert(ctx context.Context, o1 *dto.LocalRepository) (*dto.LocalRepository, error) {
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
func (inst *RepositoryServiceImpl) Update(ctx context.Context, id dxo.LocalRepositoryID, o1 *dto.LocalRepository) (*dto.LocalRepository, error) {
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
func (inst *RepositoryServiceImpl) Remove(ctx context.Context, id dxo.LocalRepositoryID) error {
	return inst.LocalRepositoryDAO.Remove(id)
}
