package projecttypes

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
)

// ProjectTypeDaoImpl ...
type ProjectTypeDaoImpl struct {
	markup.Component `id:"ProjectTypeDAO"`

	Agent          dbagent.GormDBAgent    `inject:"#GormDBAgent"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *ProjectTypeDaoImpl) _Impl() dao.ProjectTypeDAO {
	return inst
}

func (inst *ProjectTypeDaoImpl) model() *entity.ProjectType {
	return &entity.ProjectType{}
}

func (inst *ProjectTypeDaoImpl) modelList() []*entity.ProjectType {
	return make([]*entity.ProjectType, 0)
}

// Find ...
func (inst *ProjectTypeDaoImpl) Find(id dxo.ProjectTypeID) (*entity.ProjectType, error) {
	m := inst.model()
	db := inst.Agent.DB()
	res := db.First(m, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return m, nil
}

// ListAll ...
func (inst *ProjectTypeDaoImpl) ListAll() ([]*entity.ProjectType, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// Insert ...
func (inst *ProjectTypeDaoImpl) Insert(o *entity.ProjectType) (*entity.ProjectType, error) {

	// compute fields
	strKey := o.Key.String()
	strName := o.Name.String()
	o.ID = 0
	o.UUID = inst.UUIDGenService.GenerateUUID(strName + "|entity.ProjectType|" + strKey)

	db := inst.Agent.DB()
	res := db.Create(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// Update ....
func (inst *ProjectTypeDaoImpl) Update(id dxo.ProjectTypeID, o1 *entity.ProjectType) (*entity.ProjectType, error) {
	o2 := inst.model()
	db := inst.Agent.DB()
	res := db.First(o2, id)
	if res.Error != nil {
		return nil, res.Error
	}

	o2.AsDir = o1.AsDir
	o2.AsFile = o1.AsFile
	o2.Priority = o1.Priority

	o2.Key = o1.Key
	o2.Name = o1.Name
	o2.Label = o1.Label
	o2.Description = o1.Description

	res = db.Save(o2)
	if res.Error != nil {
		return nil, res.Error
	}
	return o2, nil
}

// Remove ...
func (inst *ProjectTypeDaoImpl) Remove(id dxo.ProjectTypeID) error {
	m := inst.model()
	db := inst.Agent.DB()
	res := db.Unscoped().Delete(m, id)
	return res.Error
}
