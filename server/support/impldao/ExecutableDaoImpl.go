package impldao

import (
	"errors"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
)

// ExecutableDaoImpl ...
type ExecutableDaoImpl struct {
	markup.Component `id:"ExecutableDAO"`

	Agent          GormDBAgent            `inject:"#GormDBAgent"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *ExecutableDaoImpl) _Impl() dao.ExecutableDAO {
	return inst
}

func (inst *ExecutableDaoImpl) model() *entity.Executable {
	return &entity.Executable{}
}

func (inst *ExecutableDaoImpl) modelList() []*entity.Executable {
	return make([]*entity.Executable, 0)
}

func (inst *ExecutableDaoImpl) Find(id dxo.ExecutableID) (*entity.Executable, error) {
	return nil, errors.New("no impl")
}

func (inst *ExecutableDaoImpl) ListAll() ([]*entity.Executable, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// Insert ...
func (inst *ExecutableDaoImpl) Insert(o *entity.Executable) (*entity.Executable, error) {
	uuid := inst.UUIDGenService.GenerateUUID("entity.Executable,path=" + o.Path)
	o.UUID = uuid
	db := inst.Agent.DB()
	res := db.Create(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

func (inst *ExecutableDaoImpl) Update(id dxo.ExecutableID, o *entity.Executable) (*entity.Executable, error) {
	return nil, errors.New("no impl")
}

func (inst *ExecutableDaoImpl) Remove(id dxo.ExecutableID) error {
	return errors.New("no impl")
}

// FindByPath ...
func (inst *ExecutableDaoImpl) FindByPath(path string) (*entity.Executable, error) {
	db := inst.Agent.DB()
	o := inst.model()
	res := db.Where("path=?", path).First(&o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}
