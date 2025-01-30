package idao

import (
	"github.com/bitwormhole/wpm/app/classes/locations"
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/bitwormhole/wpm/app/data/entity"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// LocationDaoImpl ...
type LocationDaoImpl struct {

	//starter:component

	_as func(locations.LocationDAO) //starter:as("#")

	Agent   dxo.DatabaseAgent  //starter:inject("#")
	UUIDGen random.UUIDService //starter:inject("#")

}

func (inst *LocationDaoImpl) _impl() locations.LocationDAO {
	return inst
}

func (inst *LocationDaoImpl) model() *entity.Location {
	return new(entity.Location)
}

func (inst *LocationDaoImpl) modelList() []*entity.Location {
	return make([]*entity.Location, 0)
}

func (inst *LocationDaoImpl) makeResult(item *entity.Location, res *gorm.DB) (*entity.Location, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Find ...
func (inst *LocationDaoImpl) Find(db *gorm.DB, id dxo.LocationID) (*entity.Location, error) {
	db = inst.Agent.DB(db)
	m := inst.model()
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// ListWithCategory ...
func (inst *LocationDaoImpl) ListWithCategory(db *gorm.DB, category dxo.Category, q *locations.Query) ([]*entity.Location, error) {
	db = inst.Agent.DB(db)
	list := inst.modelList()
	item := inst.model()

	if q == nil {
		q = new(locations.Query)
		q.Pagination.Page = 1
		q.Pagination.Size = 5
	}

	f := &finder{
		db:   db,
		item: item,
		list: list,
		all:  q.All,
		page: &q.Pagination,
	}
	f.where = func(db2 *gorm.DB) *gorm.DB {
		return db2.Where("category = ?", category)
	}
	err := f.find()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// Insert ...
func (inst *LocationDaoImpl) Insert(db *gorm.DB, item *entity.Location) (*entity.Location, error) {

	uuid := inst.UUIDGen.Build().Class("entity.Location").Generate()

	item.ID = 0
	item.UUID = uuid

	db = inst.Agent.DB(db)
	res := db.Create(item)
	return inst.makeResult(item, res)
}

// Update ...
func (inst *LocationDaoImpl) Update(db *gorm.DB, id dxo.LocationID, fn func(*entity.Location)) (*entity.Location, error) {
	db = inst.Agent.DB(db)
	item, err := inst.Find(db, id)
	if err != nil {
		return nil, err
	}
	fn(item)
	res := db.Save(item)
	return inst.makeResult(item, res)
}

// Remove ...
func (inst *LocationDaoImpl) Remove(db *gorm.DB, id dxo.LocationID) error {
	db = inst.Agent.DB(db)
	item := inst.model()
	res := db.Delete(item, id)
	return res.Error
}
