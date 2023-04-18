package settings

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
)

// SettingDaoImpl ...
type SettingDaoImpl struct {
	markup.Component `id:"SettingDAO"`

	Agent          dbagent.GormDBAgent    `inject:"#GormDBAgent"`
	TrashService   service.TrashService   `inject:"#TrashService"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *SettingDaoImpl) _Impl() dao.SettingDAO {
	return inst
}

func (inst *SettingDaoImpl) model() *entity.Setting {
	return &entity.Setting{}
}

func (inst *SettingDaoImpl) modelList() []*entity.Setting {
	return make([]*entity.Setting, 0)
}

// Find ...
func (inst *SettingDaoImpl) Find(id dxo.SettingID) (*entity.Setting, error) {
	o := inst.model()
	db := inst.Agent.DB()
	res := db.First(&o, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// FindByName ...
func (inst *SettingDaoImpl) FindByName(name string) (*entity.Setting, error) {
	o := inst.model()
	db := inst.Agent.DB()
	res := db.Where("name=?", name).First(&o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// ListAll ...
func (inst *SettingDaoImpl) ListAll() ([]*entity.Setting, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// Insert ...
func (inst *SettingDaoImpl) Insert(o *entity.Setting) (*entity.Setting, error) {

	inst.TrashService.OnInsert()

	o.ID = 0
	o.UUID = inst.UUIDGenService.GenerateUUID(o.Name + "|entity.Setting|")

	db := inst.Agent.DB()
	res := db.Create(o)
	if res.Error != nil {
		return nil, res.Error
	}

	return o, nil
}

// Update ...
func (inst *SettingDaoImpl) Update(id dxo.SettingID, o1 *entity.Setting) (*entity.Setting, error) {
	o2 := inst.model()
	db := inst.Agent.DB()
	res := db.First(&o2, id)
	if res.Error != nil {
		return nil, res.Error
	}
	o2.Value = o1.Value
	res = db.Save(o2)
	if res.Error != nil {
		return nil, res.Error
	}
	return o2, nil
}

// Remove ...
func (inst *SettingDaoImpl) Remove(id dxo.SettingID) error {

	inst.TrashService.OnDelete()

	o := inst.model()
	db := inst.Agent.DB()
	res := db.Delete(&o, id)
	return res.Error
}

// Put ...
func (inst *SettingDaoImpl) Put(name string, o1 *entity.Setting) {
	o2, err := inst.FindByName(name)
	if o2 != nil && err == nil {
		// update
		id := o2.ID
		inst.Update(id, o2)
		return
	}
	// insert new
	inst.Insert(o1)
}

// Get ...
func (inst *SettingDaoImpl) Get(name string, def *entity.Setting) *entity.Setting {
	o, _ := inst.FindByName(name)
	if o == nil {
		o = def
	}
	return o
}
