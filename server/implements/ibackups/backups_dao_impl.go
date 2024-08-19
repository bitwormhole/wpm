package ibackups

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/backups"
	"github.com/starter-go/security/random"
)

// DaoImpl ...
type DaoImpl struct {

	//starter:component
	_as func(backups.DAO) //starter:as("#")

	Agent       dxo.DatabaseAgent  //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *DaoImpl) _impl() backups.DAO {
	return inst
}

// func (inst *DaoImpl) model() *entity.Backup {
// 	return new(entity.Backup)
// }

// func (inst *DaoImpl) modelList() []*entity.Backup {
// 	return make([]*entity.Backup, 0)
// }

// func (inst *DaoImpl) makeResult(o *entity.Backup, res *gorm.DB) (*entity.Backup, error) {
// 	err := res.Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return o, nil
// }

// // Find ...
// func (inst *DaoImpl) Find(db *gorm.DB, id dxo.BackupID) (*entity.Backup, error) {
// 	m := inst.model()
// 	db = inst.Agent.DB(db)
// 	res := db.First(m, id)
// 	return inst.makeResult(m, res)
// }

// // Find(db *gorm.DB, id dxo.BackupID) (*entity.Backup, error)

// // Insert ...
// func (inst *DaoImpl) Insert(db *gorm.DB, o *entity.Backup) (*entity.Backup, error) {
// 	return nil, fmt.Errorf("no impl")
// }

// // Update ...
// func (inst *DaoImpl) Update(db *gorm.DB, id dxo.BackupID, fn func(*entity.Backup)) (*entity.Backup, error) {
// 	return nil, fmt.Errorf("no impl")
// }

// // Remove ...
// func (inst *DaoImpl) Remove(db *gorm.DB, id dxo.BackupID) error {
// 	return fmt.Errorf("no impl")
// }

// // List ...
// func (inst *DaoImpl) List(db *gorm.DB, q *Backups.Query) ([]*entity.Backup, error) {
// 	return nil, fmt.Errorf("no impl")
// }
