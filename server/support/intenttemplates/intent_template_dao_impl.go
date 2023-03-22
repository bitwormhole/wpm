package intenttemplates

import (
	"errors"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
)

// IntentTemplateDaoImpl ...
type IntentTemplateDaoImpl struct {
	markup.Component `id:"IntentTemplateDAO"`

	Agent          dbagent.GormDBAgent    `inject:"#GormDBAgent"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *IntentTemplateDaoImpl) _Impl() dao.IntentTemplateDAO {
	return inst
}

func (inst *IntentTemplateDaoImpl) model() *entity.IntentTemplate {
	return &entity.IntentTemplate{}
}

func (inst *IntentTemplateDaoImpl) modelList() []*entity.IntentTemplate {
	return make([]*entity.IntentTemplate, 0)
}

// Find ...
func (inst *IntentTemplateDaoImpl) Find(id dxo.IntentTemplateID) (*entity.IntentTemplate, error) {
	return nil, errors.New("no impl")
}

// ListAll ...
func (inst *IntentTemplateDaoImpl) ListAll() ([]*entity.IntentTemplate, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// Insert ...
func (inst *IntentTemplateDaoImpl) Insert(o *entity.IntentTemplate) (*entity.IntentTemplate, error) {

	o.ID = 0
	o.UUID = inst.UUIDGenService.GenerateUUID(o.Name + "|entity.IntentTemplate|")

	db := inst.Agent.DB()
	res := db.Create(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// Update ...
func (inst *IntentTemplateDaoImpl) Update(id dxo.IntentTemplateID, o *entity.IntentTemplate) (*entity.IntentTemplate, error) {
	m := inst.model()
	db := inst.Agent.DB()
	res := db.Find(m, id)
	if res.Error != nil {
		return nil, res.Error
	}

	m.Name = o.Name
	m.Description = o.Description

	res = db.Save(m)
	return m, res.Error
}

// Remove ...
func (inst *IntentTemplateDaoImpl) Remove(id dxo.IntentTemplateID) error {
	m := inst.model()
	db := inst.Agent.DB()
	res := db.Unscoped().Delete(&m, id)
	return res.Error
}
