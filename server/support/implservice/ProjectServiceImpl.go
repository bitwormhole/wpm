package implservice

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ProjectServiceImpl ...
type ProjectServiceImpl struct {
	markup.Component `id:"ProjectService"`

	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
	ProjectDAO     dao.ProjectDAO         `inject:"#ProjectDAO"`
}

func (inst *ProjectServiceImpl) _Impl() service.ProjectService {
	return inst
}

func (inst *ProjectServiceImpl) dto2entity(o1 *dto.Project) (*entity.Project, error) {
	o2 := &entity.Project{}
	o2.ID = o1.ID

	o2.Description = o1.Description
	o2.FullPath = o1.FullPath
	o2.Name = o1.Name
	o2.OwnerRepository = o1.OwnerRepository
	o2.PathInWorktree = o1.PathInWorktree
	o2.ProjectType = o1.ProjectType

	return o2, nil
}

func (inst *ProjectServiceImpl) entity2dto(o1 *entity.Project) (*dto.Project, error) {
	o2 := &dto.Project{}
	o2.ID = o1.ID

	o2.Name = o1.Name
	o2.Description = o1.Description

	o2.FullPath = o1.FullPath
	o2.PathInWorktree = o1.PathInWorktree

	o2.OwnerRepository = o1.OwnerRepository
	o2.ProjectType = o1.ProjectType

	return o2, nil
}

func (inst *ProjectServiceImpl) entity2dtoList(src []*entity.Project) ([]*dto.Project, error) {
	dst := make([]*dto.Project, 0)
	for _, item1 := range src {
		item2, err := inst.entity2dto(item1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}

// ListAll ...
func (inst *ProjectServiceImpl) ListAll(ctx context.Context, options *service.ProjectOptions) ([]*dto.Project, error) {
	src, err := inst.ProjectDAO.ListAll()
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(src)
}

// ListByIds ...
func (inst *ProjectServiceImpl) ListByIds(ctx context.Context, ids []dxo.ProjectID, options *service.ProjectOptions) ([]*dto.Project, error) {
	src, err := inst.ProjectDAO.ListByIds(ids)
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(src)
}

// Find ...
func (inst *ProjectServiceImpl) Find(ctx context.Context, id dxo.ProjectID, options *service.ProjectOptions) (*dto.Project, error) {
	p, err := inst.ProjectDAO.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(p)
}

// FindByOwnerRepository ...
func (inst *ProjectServiceImpl) FindByOwnerRepository(ctx context.Context, id dxo.LocalRepositoryID, options *service.ProjectOptions) ([]*dto.Project, error) {
	items1, err := inst.ProjectDAO.FindByOwnerRepository(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(items1)
}

// Insert ...
func (inst *ProjectServiceImpl) Insert(ctx context.Context, o1 *dto.Project) (*dto.Project, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.ProjectDAO.Insert(o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3)
}

// Update ...
func (inst *ProjectServiceImpl) Update(ctx context.Context, id dxo.ProjectID, o1 *dto.Project) (*dto.Project, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.ProjectDAO.Update(id, o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3)
}

// Remove ...
func (inst *ProjectServiceImpl) Remove(ctx context.Context, id dxo.ProjectID) error {
	return inst.ProjectDAO.Remove(id)
}
