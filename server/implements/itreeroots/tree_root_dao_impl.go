package itreeroots

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/treeroots"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// TreeRootDaoImpl ...
type TreeRootDaoImpl struct {

	//starter:component
	_as func(treeroots.DAO) //starter:as("#")

	Agent       dxo.DatabaseAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")

}

func (inst *TreeRootDaoImpl) _impl() treeroots.DAO {
	return inst
}

func (inst *TreeRootDaoImpl) model() *entity.TreeRoot {
	return new(entity.TreeRoot)
}

func (inst *TreeRootDaoImpl) modelList() []*entity.TreeRoot {
	return make([]*entity.TreeRoot, 0)
}

func (inst *TreeRootDaoImpl) makeResult(o *entity.TreeRoot, res *gorm.DB) (*entity.TreeRoot, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Insert ...
func (inst *TreeRootDaoImpl) Insert(db *gorm.DB, o *entity.TreeRoot) (*entity.TreeRoot, error) {

	uuid := inst.UUIDService.Build().Class("entity.TreeRoot").Generate()

	o.ID = 0
	o.UUID = uuid

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res)
}

// Update ...
func (inst *TreeRootDaoImpl) Update(db *gorm.DB, id dxo.TreeRootID, fn func(*entity.TreeRoot)) (*entity.TreeRoot, error) {

	db = inst.Agent.DB(db)
	item1, err := inst.Find(db, id)
	if err != nil {
		return nil, err
	}
	fn(item1)
	res := db.Save(item1)
	return inst.makeResult(item1, res)
}

// Remove 。。。
func (inst *TreeRootDaoImpl) Remove(db *gorm.DB, id dxo.TreeRootID) error {
	item := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(item, id)
	return res.Error
}

// Find ...
func (inst *TreeRootDaoImpl) Find(db *gorm.DB, id dxo.TreeRootID) (*entity.TreeRoot, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// List ...
func (inst *TreeRootDaoImpl) List(db *gorm.DB, q *treeroots.Query) ([]*entity.TreeRoot, error) {
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
