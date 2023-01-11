package impldao

import (
	"errors"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
)

// IntentDaoImpl ...
type IntentDaoImpl struct {
	markup.Component `id:"IntentDAO"`

	Agent          GormDBAgent            `inject:"#GormDBAgent"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *IntentDaoImpl) _Impl() dao.IntentDAO {
	return inst
}

func (inst *IntentDaoImpl) model() *entity.Intent {
	return &entity.Intent{}
}

func (inst *IntentDaoImpl) modelList() []*entity.Intent {
	return make([]*entity.Intent, 0)
}

func (inst *IntentDaoImpl) Find(id dxo.IntentID) (*entity.Intent, error) {
	return nil, errors.New("no impl")
}

func (inst *IntentDaoImpl) ListAll() ([]*entity.Intent, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

func (inst *IntentDaoImpl) Insert(o *entity.Intent) (*entity.Intent, error) {
	return nil, errors.New("no impl")
}

func (inst *IntentDaoImpl) Update(id dxo.IntentID, o *entity.Intent) (*entity.Intent, error) {
	return nil, errors.New("no impl")
}

func (inst *IntentDaoImpl) Remove(id dxo.IntentID) error {
	return errors.New("no impl")
}
