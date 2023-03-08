package impldao

import (
	"errors"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
)

// ProjectDaoImpl ...
type ProjectDaoImpl struct {
	markup.Component `id:"ProjectDAO"`

	Agent          GormDBAgent            `inject:"#GormDBAgent"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *ProjectDaoImpl) _Impl() dao.ProjectDAO {
	return inst
}

func (inst *ProjectDaoImpl) model() *entity.Project {
	return &entity.Project{}
}

func (inst *ProjectDaoImpl) modelList() []*entity.Project {
	return make([]*entity.Project, 0)
}

// Find ...
func (inst *ProjectDaoImpl) Find(id dxo.ProjectID) (*entity.Project, error) {
	return nil, errors.New("no impl")
}

// FindByOwnerRepository ...
func (inst *ProjectDaoImpl) FindByOwnerRepository(id dxo.LocalRepositoryID) ([]*entity.Project, error) {
	return nil, errors.New("no impl")
}

// ListAll ...
func (inst *ProjectDaoImpl) ListAll() ([]*entity.Project, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// ListByIds ...
func (inst *ProjectDaoImpl) ListByIds(ids []dxo.ProjectID) ([]*entity.Project, error) {

	if ids == nil {
		return []*entity.Project{}, nil
	}

	if len(ids) == 0 {
		return []*entity.Project{}, nil
	}

	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list, ids)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// Insert ...
func (inst *ProjectDaoImpl) Insert(o *entity.Project) (*entity.Project, error) {

	// compute fields
	o.ID = 0
	o.UUID = inst.UUIDGenService.GenerateUUID(o.FullPath + "|entity.Project|")

	// save
	db := inst.Agent.DB()
	res := db.Create(o)
	if res.Error != nil {
		return nil, res.Error
	}

	return o, nil
}

// Update ...
func (inst *ProjectDaoImpl) Update(id dxo.ProjectID, o1 *entity.Project) (*entity.Project, error) {
	db := inst.Agent.DB()
	o2 := inst.model()
	res := db.First(o2, id)
	if res.Error != nil {
		return nil, res.Error
	}

	// todo ...
	o2.Name = o1.Name
	o2.ProjectType = o1.ProjectType
	o2.PathInWorktree = o1.PathInWorktree
	o2.FullPath = o1.FullPath
	o2.Description = o1.Description
	o2.OwnerRepository = o1.OwnerRepository

	res = db.Save(o2)
	if res.Error != nil {
		return nil, res.Error
	}
	return o2, nil
}

// Remove ...
func (inst *ProjectDaoImpl) Remove(id dxo.ProjectID) error {
	db := inst.Agent.DB()
	m := inst.model()
	res := db.Unscoped().Delete(m, id)
	return res.Error
}
