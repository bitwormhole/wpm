package irepositories

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/repositories"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// LocalRepositoryDao ...
type LocalRepositoryDao struct {

	//starter:component
	_as func(repositories.LocalRepositoryDAO) //starter:as("#")

	Agent       dxo.DatabaseAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *LocalRepositoryDao) _impl() repositories.LocalRepositoryDAO {
	return inst
}

func (inst *LocalRepositoryDao) model() *entity.LocalRepository {
	return new(entity.LocalRepository)
}

func (inst *LocalRepositoryDao) modelList() []*entity.LocalRepository {
	return make([]*entity.LocalRepository, 0)
}

func (inst *LocalRepositoryDao) makeResult(o *entity.LocalRepository, res *gorm.DB) (*entity.LocalRepository, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Find ...
func (inst *LocalRepositoryDao) Find(db *gorm.DB, id dxo.LocalRepositoryID) (*entity.LocalRepository, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// Insert ...
func (inst *LocalRepositoryDao) Insert(db *gorm.DB, o *entity.LocalRepository) (*entity.LocalRepository, error) {

	o.ID = 0
	o.UUID = inst.UUIDService.Build().Class("entity.LocalRepository").Generate()

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res)
}

// Update ...
func (inst *LocalRepositoryDao) Update(db *gorm.DB, id dxo.LocalRepositoryID, fn func(*entity.LocalRepository)) (*entity.LocalRepository, error) {
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
func (inst *LocalRepositoryDao) Remove(db *gorm.DB, id dxo.LocalRepositoryID) error {
	item := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(item, id)
	return res.Error
}

// List ...
func (inst *LocalRepositoryDao) List(db *gorm.DB, q *repositories.Query) ([]*entity.LocalRepository, error) {

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
