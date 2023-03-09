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

// ExampleDaoImpl ...
type ExampleDaoImpl struct {
	markup.Component `id:"ExampleDAO"`

	Agent          dbagent.GormDBAgent    `inject:"#GormDBAgent"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *ExampleDaoImpl) _Impl() dao.ExampleDAO {
	return inst
}

func (inst *ExampleDaoImpl) model() *entity.Example {
	return &entity.Example{}
}

func (inst *ExampleDaoImpl) modelList() []*entity.Example {
	return make([]*entity.Example, 0)
}

// Find ...
func (inst *ExampleDaoImpl) Find(id dxo.ExampleID) (*entity.Example, error) {
	return nil, errors.New("no impl")
}

// ListAll ...
func (inst *ExampleDaoImpl) ListAll() ([]*entity.Example, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// Insert ...
func (inst *ExampleDaoImpl) Insert(o *entity.Example) (*entity.Example, error) {
	return nil, errors.New("no impl")
}

// Update ....
func (inst *ExampleDaoImpl) Update(id dxo.ExampleID, o *entity.Example) (*entity.Example, error) {
	return nil, errors.New("no impl")
}

// Remove ...
func (inst *ExampleDaoImpl) Remove(id dxo.ExampleID) error {
	return errors.New("no impl")
}
