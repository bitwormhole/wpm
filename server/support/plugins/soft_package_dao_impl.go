package plugins

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
)

// PluginDaoImpl ...
type PluginDaoImpl struct {
	markup.Component `id:"SoftwarePackageDAO"`

	Agent          dbagent.GormDBAgent    `inject:"#GormDBAgent"`
	TrashService   service.TrashService   `inject:"#TrashService"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *PluginDaoImpl) _Impl() dao.SoftwarePackageDAO {
	return inst
}

func (inst *PluginDaoImpl) model() *entity.SoftwarePackage {
	return &entity.SoftwarePackage{}
}

func (inst *PluginDaoImpl) modelList() []*entity.SoftwarePackage {
	return make([]*entity.SoftwarePackage, 0)
}

// Find ...
func (inst *PluginDaoImpl) Find(id dxo.SoftwarePackageID) (*entity.SoftwarePackage, error) {
	o := inst.model()
	db := inst.Agent.DB()
	res := db.First(&o, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// ListAll ...
func (inst *PluginDaoImpl) ListAll() ([]*entity.SoftwarePackage, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// ListByModuleName ...
func (inst *PluginDaoImpl) ListByModuleName(mod string) ([]*entity.SoftwarePackage, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Where("module_name = ?", mod).Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// Insert ...
func (inst *PluginDaoImpl) Insert(o *entity.SoftwarePackage) (*entity.SoftwarePackage, error) {

	inst.TrashService.OnInsert()

	uuid := inst.UUIDGenService.GenerateUUID("entity.SoftwarePackage,path=" + o.Name)
	o.UUID = uuid
	o.ID = 0
	db := inst.Agent.DB()
	res := db.Create(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// Update ...
func (inst *PluginDaoImpl) Update(id dxo.SoftwarePackageID, o1 *entity.SoftwarePackage) (*entity.SoftwarePackage, error) {

	o2 := inst.model()
	db := inst.Agent.DB()
	res := db.First(&o2, id)
	if res.Error != nil {
		return nil, res.Error
	}

	o2.Name = o1.Name
	o2.Description = o1.Description
	o2.Referer = o1.Referer
	o2.Installation = o1.Installation

	res = db.Save(o2)
	if res.Error != nil {
		return nil, res.Error
	}

	return o2, nil
}

// Remove ...
func (inst *PluginDaoImpl) Remove(id dxo.SoftwarePackageID) error {

	inst.TrashService.OnDelete()

	o := inst.model()
	db := inst.Agent.DB()
	res := db.Delete(o, id)
	return res.Error
}
