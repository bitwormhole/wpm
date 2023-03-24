package executables

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
)

// ExecutableDaoImpl ...
type ExecutableDaoImpl struct {
	markup.Component `id:"ExecutableDAO"`

	Agent          dbagent.GormDBAgent    `inject:"#GormDBAgent"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *ExecutableDaoImpl) _Impl() dao.ExecutableDAO {
	return inst
}

func (inst *ExecutableDaoImpl) model() *entity.Executable {
	return &entity.Executable{}
}

func (inst *ExecutableDaoImpl) modelList() []*entity.Executable {
	return make([]*entity.Executable, 0)
}

// Find ...
func (inst *ExecutableDaoImpl) Find(id dxo.ExecutableID) (*entity.Executable, error) {
	db := inst.Agent.DB()
	o := inst.model()
	res := db.First(&o, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// ListAll ...
func (inst *ExecutableDaoImpl) ListAll() ([]*entity.Executable, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// Insert ...
func (inst *ExecutableDaoImpl) Insert(o *entity.Executable) (*entity.Executable, error) {
	uuid := inst.UUIDGenService.GenerateUUID("entity.Executable,path=" + o.Path)
	o.UUID = uuid
	db := inst.Agent.DB()
	res := db.Create(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// Update ...
func (inst *ExecutableDaoImpl) Update(id dxo.ExecutableID, o1 *entity.Executable) (*entity.Executable, error) {
	o2, err := inst.Find(id)
	if err != nil {
		return nil, err
	}

	o2.Name = o1.Name
	o2.Title = o1.Title
	o2.IconURL = o1.IconURL
	o2.Description = o1.Description
	o2.Path = o1.Path
	o2.Size = o1.Size
	o2.SHA256SUM = o1.SHA256SUM
	o2.OpenWithPriority = o1.OpenWithPriority

	db := inst.Agent.DB()
	res := db.Save(o2)
	if res.Error != nil {
		return nil, res.Error
	}
	return o2, nil
}

// Remove ...
func (inst *ExecutableDaoImpl) Remove(id dxo.ExecutableID) error {
	o1, err := inst.Find(id)
	if err != nil {
		return err
	}
	db := inst.Agent.DB()
	res := db.Unscoped().Delete(o1, id)
	return res.Error
}

// FindByPath ...
func (inst *ExecutableDaoImpl) FindByPath(path string) (*entity.Executable, error) {
	db := inst.Agent.DB()
	o := inst.model()
	res := db.Where("path=?", path).First(&o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// FindByName ...
func (inst *ExecutableDaoImpl) FindByName(name string) (*entity.Executable, error) {
	db := inst.Agent.DB()
	o := inst.model()
	res := db.Where("name=?", name).First(&o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}
