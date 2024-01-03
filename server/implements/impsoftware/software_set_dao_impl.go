package impsoftware

import (
	"github.com/bitwormhole/wpm/server/classes/softwaresets"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// SoftwareSetDaoImpl ...
type SoftwareSetDaoImpl struct {

	//starter:component
	_as func(softwaresets.DAO) //starter:as("#")

	Agent       dxo.DatabaseAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *SoftwareSetDaoImpl) _impl() softwaresets.DAO {
	return inst
}

func (inst *SoftwareSetDaoImpl) model() *entity.SoftwareSet {
	return new(entity.SoftwareSet)
}

func (inst *SoftwareSetDaoImpl) modelList() []*entity.SoftwareSet {
	return make([]*entity.SoftwareSet, 0)
}

func (inst *SoftwareSetDaoImpl) makeResult(o *entity.SoftwareSet, res *gorm.DB) (*entity.SoftwareSet, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Find ...
func (inst *SoftwareSetDaoImpl) Find(db *gorm.DB, id dxo.SoftwareSetID) (*entity.SoftwareSet, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// Insert ...
func (inst *SoftwareSetDaoImpl) Insert(db *gorm.DB, o *entity.SoftwareSet) (*entity.SoftwareSet, error) {

	uuid := inst.UUIDService.Build().Class("entity.SoftwareSet").Generate()

	o.ID = 0
	o.UUID = uuid

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res)
}

// Update ...
func (inst *SoftwareSetDaoImpl) Update(db *gorm.DB, id dxo.SoftwareSetID, fn func(*entity.SoftwareSet)) (*entity.SoftwareSet, error) {
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
func (inst *SoftwareSetDaoImpl) Remove(db *gorm.DB, id dxo.SoftwareSetID) error {
	item := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(item, id)
	return res.Error
}

// List ...
func (inst *SoftwareSetDaoImpl) List(db *gorm.DB, q *softwaresets.Query) ([]*entity.SoftwareSet, error) {
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
