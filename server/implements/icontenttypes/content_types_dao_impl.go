package icontenttypes

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/contenttypes"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// DaoImpl ...
type DaoImpl struct {

	//starter:component
	_as func(contenttypes.DAO) //starter:as("#")

	Agent       dxo.DatabaseAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *DaoImpl) _impl() contenttypes.DAO {
	return inst
}

func (inst *DaoImpl) model() *entity.ContentType {
	return new(entity.ContentType)
}

func (inst *DaoImpl) modelList() []*entity.ContentType {
	return make([]*entity.ContentType, 0)
}

func (inst *DaoImpl) makeResult(o *entity.ContentType, res *gorm.DB) (*entity.ContentType, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Find ...
func (inst *DaoImpl) Find(db *gorm.DB, id dxo.ContentTypeID) (*entity.ContentType, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// Insert ...
func (inst *DaoImpl) Insert(db *gorm.DB, o *entity.ContentType) (*entity.ContentType, error) {

	uuid := inst.UUIDService.Build().Class("entity.ContentType").Generate()

	o.ID = 0
	o.UUID = uuid

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res)
}

// Update ...
func (inst *DaoImpl) Update(db *gorm.DB, id dxo.ContentTypeID, fn func(*entity.ContentType)) (*entity.ContentType, error) {
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
func (inst *DaoImpl) Remove(db *gorm.DB, id dxo.ContentTypeID) error {
	item := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(item, id)
	return res.Error
}

// List ...
func (inst *DaoImpl) List(db *gorm.DB, q *contenttypes.Query) ([]*entity.ContentType, error) {
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
