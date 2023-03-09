package projecttypes

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ProjectTypeServiceImpl ...
type ProjectTypeServiceImpl struct {
	markup.Component `id:"ProjectTypeService"`

	ProjectTypeDAO    dao.ProjectTypeDAO        `inject:"#ProjectTypeDAO"`
	FileSystemService service.FileSystemService `inject:"#FileSystemService"`
}

func (inst *ProjectTypeServiceImpl) _Impl() service.ProjectTypeService {
	return inst
}

func (inst *ProjectTypeServiceImpl) dto2entity(o1 *dto.ProjectType) (*entity.ProjectType, error) {
	o2 := &entity.ProjectType{}
	o2.ID = o1.ID

	// todo ...
	o2.Name = o1.Name
	o2.Type = o1.Type
	o2.Label = o1.Label
	o2.Description = o1.Description
	o2.AsDir = o1.AsDir
	o2.AsFile = o1.AsFile

	return o2, nil
}

func (inst *ProjectTypeServiceImpl) entity2dto(o1 *entity.ProjectType) (*dto.ProjectType, error) {
	o2 := &dto.ProjectType{}
	o2.ID = o1.ID

	// todo ...
	o2.Name = o1.Name
	o2.Type = o1.Type
	o2.Label = o1.Label
	o2.Description = o1.Description
	o2.AsDir = o1.AsDir
	o2.AsFile = o1.AsFile

	return o2, nil
}

// ListAll ...
func (inst *ProjectTypeServiceImpl) ListAll(ctx context.Context) ([]*dto.ProjectType, error) {
	src, err := inst.ProjectTypeDAO.ListAll()
	if err != nil {
		return nil, err
	}
	dst := make([]*dto.ProjectType, 0)
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
func (inst *ProjectTypeServiceImpl) Find(ctx context.Context, id dxo.ProjectTypeID) (*dto.ProjectType, error) {
	o1, err := inst.ProjectTypeDAO.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1)
}

// Insert ...
func (inst *ProjectTypeServiceImpl) Insert(ctx context.Context, o1 *dto.ProjectType) (*dto.ProjectType, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.ProjectTypeDAO.Insert(o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3)
}

// Update ...
func (inst *ProjectTypeServiceImpl) Update(ctx context.Context, id dxo.ProjectTypeID, o1 *dto.ProjectType) (*dto.ProjectType, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.ProjectTypeDAO.Update(id, o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3)
}

// Remove ...
func (inst *ProjectTypeServiceImpl) Remove(ctx context.Context, id dxo.ProjectTypeID) error {
	return inst.ProjectTypeDAO.Remove(id)
}

// LocateProject ...
func (inst *ProjectTypeServiceImpl) LocateProject(ctx context.Context, o *dto.Project, path string) error {
	locator := &myProjectLocatorWithTypes{
		parent:  inst,
		path:    path,
		project: o,
		context: ctx,
	}
	return locator.Locate()
}
