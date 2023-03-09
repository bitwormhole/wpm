package impldao

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
	return nil, errors.New("no impl")
}

// Update ...
func (inst *IntentTemplateDaoImpl) Update(id dxo.IntentTemplateID, o *entity.IntentTemplate) (*entity.IntentTemplate, error) {
	return nil, errors.New("no impl")
}

// Remove ...
func (inst *IntentTemplateDaoImpl) Remove(id dxo.IntentTemplateID) error {
	return errors.New("no impl")
}
