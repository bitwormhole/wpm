package impsoftware

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/packages"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// SoftwarePackageDaoImpl ...
type SoftwarePackageDaoImpl struct {

	//starter:component
	_as func(packages.DAO) //starter:as("#")

	Agent       dxo.DatabaseAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *SoftwarePackageDaoImpl) _impl() packages.DAO {
	return inst
}

func (inst *SoftwarePackageDaoImpl) model() *entity.SoftwarePackage {
	return new(entity.SoftwarePackage)
}

func (inst *SoftwarePackageDaoImpl) modelList() []*entity.SoftwarePackage {
	return make([]*entity.SoftwarePackage, 0)
}

func (inst *SoftwarePackageDaoImpl) makeResult(o *entity.SoftwarePackage, res *gorm.DB) (*entity.SoftwarePackage, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Find ...
func (inst *SoftwarePackageDaoImpl) Find(db *gorm.DB, id dxo.SoftwarePackageID) (*entity.SoftwarePackage, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}

// Insert ...
func (inst *SoftwarePackageDaoImpl) Insert(db *gorm.DB, o *entity.SoftwarePackage) (*entity.SoftwarePackage, error) {

	uuid := inst.UUIDService.Build().Class("entity.SoftwarePackage").Generate()

	o.ID = 0
	o.UUID = uuid

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(o, res)
}

// Update ...
func (inst *SoftwarePackageDaoImpl) Update(db *gorm.DB, id dxo.SoftwarePackageID, fn func(*entity.SoftwarePackage)) (*entity.SoftwarePackage, error) {
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
func (inst *SoftwarePackageDaoImpl) Remove(db *gorm.DB, id dxo.SoftwarePackageID) error {
	item := inst.model()
	db = inst.Agent.DB(db)
	res := db.Delete(item, id)
	return res.Error
}

// List ...
func (inst *SoftwarePackageDaoImpl) List(db *gorm.DB, q *packages.Query) ([]*entity.SoftwarePackage, error) {
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
