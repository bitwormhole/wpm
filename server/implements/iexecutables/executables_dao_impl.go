package iexecutables

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/executables"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// DaoImpl ...
type DaoImpl struct {

	//starter:component
	_as func(executables.DAO) //starter:as("#")

	Agent       dxo.DatabaseAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *DaoImpl) _impl() executables.DAO {
	return inst
}

func (inst *DaoImpl) model() *entity.Executable {
	return new(entity.Executable)
}

func (inst *DaoImpl) modelList() []*entity.Executable {
	return make([]*entity.Executable, 0)
}

func (inst *DaoImpl) makeResult(o *entity.Executable, res *gorm.DB) (*entity.Executable, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Find ...
func (inst *DaoImpl) Find(db *gorm.DB, id dxo.ExecutableID) (*entity.Executable, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// Find(db *gorm.DB, id dxo.ExecutableID) (*entity.Executable, error)

// Insert ...
func (inst *DaoImpl) Insert(db *gorm.DB, o *entity.Executable) (*entity.Executable, error) {

	uuid := inst.UUIDService.Build().Class("entity.Executable").Generate()

	o.ID = 0
	o.UUID = uuid

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res)
}

// Update ...
func (inst *DaoImpl) Update(db *gorm.DB, id dxo.ExecutableID, fn func(*entity.Executable)) (*entity.Executable, error) {
	db = inst.Agent.DB(db)
	item1, err := inst.Find(db, id)
	if err != nil {
		return nil, err
	}
	fn(item1)
	res := db.Save(item1)
	return inst.makeResult(item1, res)
}

// Remove ...
func (inst *DaoImpl) Remove(db *gorm.DB, id dxo.ExecutableID) error {
	item := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(item, id)
	return res.Error
}

// List ...
func (inst *DaoImpl) List(db *gorm.DB, q *executables.Query) ([]*entity.Executable, error) {
	item := inst.model()
	list := inst.modelList()
	db = inst.Agent.DB(db)

	finder := new(dao.Finder)
	finder.DB = db
	finder.ItemModel = item
	finder.ListModel = &list
	finder.Pagination = &q.Pagination
	finder.WantAll = q.All

	err := finder.Find()
	if err != nil {
		return nil, err
	}
	return list, nil
}
