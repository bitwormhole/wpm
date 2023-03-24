package intenttemplates

import (
	"errors"
	"fmt"

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

// ListByAET ...
func (inst *IntentTemplateDaoImpl) ListByAET(sel *entity.IntentTemplate) ([]*entity.IntentTemplate, error) {

	if sel == nil {
		return nil, fmt.Errorf("no param")
	}

	// select by exe
	list1 := inst.modelList()
	db := inst.Agent.DB()
	res := db.Where("executable = ?", sel.Executable).Find(&list1)
	if res.Error != nil {
		return nil, res.Error
	}

	// filter by action & target
	a1 := sel.Action
	t1 := sel.Target
	list2 := inst.modelList()
	for _, item := range list1 {
		if a1 != "" && a1 != item.Action {
			continue
		}
		if t1 != "" && t1 != item.Target {
			continue
		}
		list2 = append(list2, item)
	}
	return list2, nil
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
	m.Title = o.Title
	m.Description = o.Description

	m.Action = o.Action
	m.Executable = o.Executable
	m.Target = o.Target
	m.Selector = o.Selector

	m.Command = o.Command
	m.Arguments = o.Arguments
	m.Env = o.Env
	m.WD = o.WD

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
