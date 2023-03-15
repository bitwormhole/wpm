package mediae

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
)

// MediaDaoImpl ...
type MediaDaoImpl struct {
	markup.Component `id:"MediaDAO"`

	Agent          dbagent.GormDBAgent    `inject:"#GormDBAgent"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *MediaDaoImpl) _Impl() dao.MediaDAO {
	return inst
}

func (inst *MediaDaoImpl) model() *entity.Media {
	return &entity.Media{}
}

func (inst *MediaDaoImpl) modelList() []*entity.Media {
	return make([]*entity.Media, 0)
}

// Find ...
func (inst *MediaDaoImpl) Find(id dxo.MediaID) (*entity.Media, error) {
	o := inst.model()
	db := inst.Agent.DB()
	res := db.First(&o, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// FindByPath ...
func (inst *MediaDaoImpl) FindByPath(path string) (*entity.Media, error) {
	o := inst.model()
	db := inst.Agent.DB()
	res := db.Where("path = ?", path).First(&o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// ListAll ...
func (inst *MediaDaoImpl) ListAll() ([]*entity.Media, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// ListByIDs ...
func (inst *MediaDaoImpl) ListByIDs(ids []dxo.MediaID) ([]*entity.Media, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list, ids)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// Insert ...
func (inst *MediaDaoImpl) Insert(o *entity.Media) (*entity.Media, error) {

	o.ID = 0
	o.UUID = inst.UUIDGenService.GenerateUUID(o.SHA256SUM.String() + "|entity.Media|")

	db := inst.Agent.DB()
	res := db.Create(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// Update ....
func (inst *MediaDaoImpl) Update(id dxo.MediaID, o1 *entity.Media) (*entity.Media, error) {

	db := inst.Agent.DB()
	o2 := inst.model()
	res := db.Find(&o2, id)
	if res.Error != nil {
		return nil, res.Error
	}

	o2.Bucket = o1.Bucket
	o2.Label = o1.Label
	o2.Name = o1.Name

	res = db.Save(o2)
	if res.Error != nil {
		return nil, res.Error
	}

	return o2, nil
}

// Remove ...
func (inst *MediaDaoImpl) Remove(id dxo.MediaID) error {
	db := inst.Agent.DB()
	o2 := inst.model()
	res := db.Find(&o2, id)
	if res.Error != nil {
		return res.Error
	}
	res = db.Unscoped().Delete(o2, id)
	return res.Error
}
