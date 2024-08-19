package irepositories

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/repositories"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// RemoteRepositoryDao ...
type RemoteRepositoryDao struct {

	//starter:component
	_as func(repositories.RemoteRepositoryDAO) //starter:as("#")

	Agent       dxo.DatabaseAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *RemoteRepositoryDao) _impl() repositories.RemoteRepositoryDAO {
	return inst
}

func (inst *RemoteRepositoryDao) model() *entity.RemoteRepository {
	return new(entity.RemoteRepository)
}

func (inst *RemoteRepositoryDao) modelList() []*entity.RemoteRepository {
	return make([]*entity.RemoteRepository, 0)
}

func (inst *RemoteRepositoryDao) makeResult(o *entity.RemoteRepository, res *gorm.DB) (*entity.RemoteRepository, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Find ...
func (inst *RemoteRepositoryDao) Find(db *gorm.DB, id dxo.RemoteRepositoryID) (*entity.RemoteRepository, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// Insert ...
func (inst *RemoteRepositoryDao) Insert(db *gorm.DB, o *entity.RemoteRepository) (*entity.RemoteRepository, error) {

	uuid := inst.UUIDService.Build().Class("entity.RemoteRepository").Generate()

	o.ID = 0
	o.UUID = uuid

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res)
}

// Update ...
func (inst *RemoteRepositoryDao) Update(db *gorm.DB, id dxo.RemoteRepositoryID, fn func(*entity.RemoteRepository)) (*entity.RemoteRepository, error) {
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
func (inst *RemoteRepositoryDao) Remove(db *gorm.DB, id dxo.RemoteRepositoryID) error {
	item := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(item, id)
	return res.Error
}

// List ...
func (inst *RemoteRepositoryDao) List(db *gorm.DB, q *repositories.Query) ([]*entity.RemoteRepository, error) {
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
