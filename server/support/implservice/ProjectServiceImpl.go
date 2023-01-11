package implservice

import (
	"context"
	"errors"

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
	o2.Name = o1.Name
	o2.Path = o1.Path
	o2.OwnerRepository = o1.OwnerRepository
	// todo ...
	return o2, nil
}

func (inst *ProjectServiceImpl) entity2dto(o1 *entity.Project) (*dto.Project, error) {
	o2 := &dto.Project{}
	o2.ID = o1.ID
	o2.Name = o1.Name
	o2.Path = o1.Path
	o2.OwnerRepository = o1.OwnerRepository
	// todo ...
	return o2, nil
}

// ListAll ...
func (inst *ProjectServiceImpl) ListAll(ctx context.Context) ([]*dto.Project, error) {
	src, err := inst.ProjectDAO.ListAll()
	if err != nil {
		return nil, err
	}
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

// Find ...
func (inst *ProjectServiceImpl) Find(ctx context.Context, id dxo.ProjectID) (*dto.Project, error) {
	return nil, errors.New("no impl")
}

// FindByOwnerRepository ...
func (inst *ProjectServiceImpl) FindByOwnerRepository(ctx context.Context, id dxo.LocalRepositoryID) ([]*dto.Project, error) {
	return nil, errors.New("no impl")
}

// Insert ...
func (inst *ProjectServiceImpl) Insert(ctx context.Context, o *dto.Project) (*dto.Project, error) {
	return nil, errors.New("no impl")
}

// Update ...
func (inst *ProjectServiceImpl) Update(ctx context.Context, id dxo.ProjectID, o *dto.Project) (*dto.Project, error) {
	return nil, errors.New("no impl")
}

// Remove ...
func (inst *ProjectServiceImpl) Remove(ctx context.Context, id dxo.ProjectID) error {
	return errors.New("no impl")
}
