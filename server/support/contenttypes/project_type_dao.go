package contenttypes

import (
	"strings"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
)

// ProjectTypeDaoImpl ...
type ProjectTypeDaoImpl struct {
	markup.Component `id:"ContentTypeDAO"`

	Agent          dbagent.GormDBAgent    `inject:"#GormDBAgent"`
	TrashService   service.TrashService   `inject:"#TrashService"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *ProjectTypeDaoImpl) _Impl() dao.ContentTypeDAO {
	return inst
}

func (inst *ProjectTypeDaoImpl) model() *entity.ContentType {
	return &entity.ContentType{}
}

func (inst *ProjectTypeDaoImpl) modelList() []*entity.ContentType {
	return make([]*entity.ContentType, 0)
}

// Find ...
func (inst *ProjectTypeDaoImpl) Find(id dxo.ContentTypeID) (*entity.ContentType, error) {
	m := inst.model()
	db := inst.Agent.DB()
	res := db.First(m, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return m, nil
}

// FindByURN ...
func (inst *ProjectTypeDaoImpl) FindByURN(urn dxo.ContentTypeURN) (*entity.ContentType, error) {
	m := inst.model()
	db := inst.Agent.DB()
	res := db.Where("urn = ?", urn).First(m)
	if res.Error != nil {
		return nil, res.Error
	}
	return m, nil
}

// FindByName ...
func (inst *ProjectTypeDaoImpl) FindByName(name dxo.ContentTypeName) (*entity.ContentType, error) {
	m := inst.model()
	db := inst.Agent.DB()
	res := db.Where("name = ?", name).First(m)
	if res.Error != nil {
		return nil, res.Error
	}
	return m, nil
}

// ListAll ...
func (inst *ProjectTypeDaoImpl) ListAll() ([]*entity.ContentType, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// ListByPattern ...
func (inst *ProjectTypeDaoImpl) ListByPattern(pattern string) ([]*entity.ContentType, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Where("patterns LIKE ?", pattern).Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	list = inst.filterByPattern(pattern, list)
	return list, nil
}

func (inst *ProjectTypeDaoImpl) filterByPattern(want string, src []*entity.ContentType) []*entity.ContentType {
	dst := make([]*entity.ContentType, 0)
	for _, item := range src {
		patterns := item.Patterns.Array()
		for _, have := range patterns {
			have = strings.TrimSpace(have)
			have = strings.ToLower(have)
			if have == "" {
				continue
			}
			if have == want {
				dst = append(dst, item)
			}
		}
	}
	return dst
}

// Insert ...
func (inst *ProjectTypeDaoImpl) Insert(o *entity.ContentType) (*entity.ContentType, error) {

	inst.TrashService.OnInsert()

	// compute fields
	strKey := o.Patterns
	strName := o.Name.String()

	o.ID = 0
	o.UUID = inst.UUIDGenService.GenerateUUID(strName + "|entity.ProjectType|" + strKey.String())
	o.Base.PrepareInsert()

	db := inst.Agent.DB()
	res := db.Create(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// Update ....
func (inst *ProjectTypeDaoImpl) Update(id dxo.ContentTypeID, o1 *entity.ContentType) (*entity.ContentType, error) {
	o2 := inst.model()
	db := inst.Agent.DB()
	res := db.First(o2, id)
	if res.Error != nil {
		return nil, res.Error
	}

	o2.AsDir = o1.AsDir
	o2.AsFile = o1.AsFile
	o2.AsProject = o1.AsProject
	o2.Priority = o1.Priority

	o2.Patterns = o1.Patterns
	o2.Name = o1.Name
	o2.Icon = o1.Icon
	o2.Label = o1.Label
	o2.Description = o1.Description

	res = db.Save(o2)
	if res.Error != nil {
		return nil, res.Error
	}
	return o2, nil
}

// Remove ...
func (inst *ProjectTypeDaoImpl) Remove(id dxo.ContentTypeID) error {

	inst.TrashService.OnDelete()

	m := inst.model()
	db := inst.Agent.DB()
	res := db.Delete(m, id)
	return res.Error
}
