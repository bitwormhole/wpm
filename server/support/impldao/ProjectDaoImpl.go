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

func (inst *ProjectDaoImpl) Insert(o *entity.Project) (*entity.Project, error) {
	return nil, errors.New("no impl")
}

func (inst *ProjectDaoImpl) Update(id dxo.ProjectID, o *entity.Project) (*entity.Project, error) {
	return nil, errors.New("no impl")
}

func (inst *ProjectDaoImpl) Remove(id dxo.ProjectID) error {
	return errors.New("no impl")
}
