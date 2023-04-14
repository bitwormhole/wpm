package intenttemplates

import (
	"errors"
	"fmt"
	"strings"

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

func (inst *IntentTemplateDaoImpl) normalizeText(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}

func (inst *IntentTemplateDaoImpl) makeCondsForListBySelector(sel *entity.IntentTemplate) (string, []any, error) {

	if sel == nil {
		return "", nil, fmt.Errorf("selector is nil")
	}

	table := make(map[string]string)
	table["executable"] = sel.Executable.String()
	table["target"] = sel.Target
	table["content_type"] = sel.ContentType
	table["method"] = sel.Method

	args := make([]any, 0)
	query := strings.Builder{}
	sep := ""

	for k, v := range table {
		v = inst.normalizeText(v)
		if v == "" || v == "*" {
			continue
		}
		query.WriteString(sep)
		query.WriteString(k + " = ?")
		args = append(args, v)
		sep = " AND "
	}

	return query.String(), args, nil
}

// ListBySelector ...
func (inst *IntentTemplateDaoImpl) ListBySelector(sel *entity.IntentTemplate) ([]*entity.IntentTemplate, error) {

	query, args, err := inst.makeCondsForListBySelector(sel)
	if err != nil {
		return nil, err
	}

	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Where(query, args...).Find(&list)
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
	m.Title = o.Title
	m.Group = o.Group
	m.Description = o.Description

	m.ContentType = o.ContentType
	m.Executable = o.Executable
	m.Method = o.Method
	m.Selector = o.Selector
	m.Target = o.Target

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
	res := db.Delete(&m, id)
	return res.Error
}
