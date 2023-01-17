package impldao

import (
	"errors"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
)

// MediaDaoImpl ...
type MediaDaoImpl struct {
	markup.Component `id:"MediaDAO"`

	Agent          GormDBAgent            `inject:"#GormDBAgent"`
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
	res := db.Find(&o, id)
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

// Insert ...
func (inst *MediaDaoImpl) Insert(o *entity.Media) (*entity.Media, error) {
	return nil, errors.New("no impl")
}

// Update ....
func (inst *MediaDaoImpl) Update(id dxo.MediaID, o *entity.Media) (*entity.Media, error) {
	return nil, errors.New("no impl")
}

// Remove ...
func (inst *MediaDaoImpl) Remove(id dxo.MediaID) error {
	return errors.New("no impl")
}
