package plugins

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
)

// SoftInstalledFileDAO ...
type SoftInstalledFileDAO struct {
	markup.Component `id:"InstalledFileDAO"`

	Agent          dbagent.GormDBAgent    `inject:"#GormDBAgent"`
	TrashService   service.TrashService   `inject:"#TrashService"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *SoftInstalledFileDAO) _Impl() dao.InstalledFileDAO {
	return inst
}

func (inst *SoftInstalledFileDAO) model() *entity.InstalledFile {
	return &entity.InstalledFile{}
}

func (inst *SoftInstalledFileDAO) modelList() []*entity.InstalledFile {
	return make([]*entity.InstalledFile, 0)
}

// Find ...
func (inst *SoftInstalledFileDAO) Find(id dxo.InstalledFileID) (*entity.InstalledFile, error) {
	db := inst.Agent.DB()
	o2 := inst.model()
	res := db.First(o2, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return o2, nil
}

// ListAll ...
func (inst *SoftInstalledFileDAO) ListAll() ([]*entity.InstalledFile, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// Insert ...
func (inst *SoftInstalledFileDAO) Insert(o *entity.InstalledFile) (*entity.InstalledFile, error) {

	inst.TrashService.OnInsert()

	// compute fields
	o.ID = 0
	o.UUID = inst.UUIDGenService.GenerateUUID(o.Path + "|entity.InstalledFile|")

	// save
	db := inst.Agent.DB()
	res := db.Create(o)
	if res.Error != nil {
		return nil, res.Error
	}

	return o, nil
}

// Update ...
func (inst *SoftInstalledFileDAO) Update(id dxo.InstalledFileID, o1 *entity.InstalledFile) (*entity.InstalledFile, error) {

	db := inst.Agent.DB()
	o2 := inst.model()
	res := db.First(o2, id)
	if res.Error != nil {
		return nil, res.Error
	}

	o2.Name = o1.Name
	o2.Path = o1.Path
	o2.IsDir = o1.IsDir
	o2.IsFile = o1.IsFile
	o2.Size = o1.Size
	o2.SHA256SUM = o1.SHA256SUM
	o2.OwnerPackage = o1.OwnerPackage
	o2.Installation = o1.Installation

	res = db.Save(o2)
	if res.Error != nil {
		return nil, res.Error
	}
	return o2, nil
}

// Remove ...
func (inst *SoftInstalledFileDAO) Remove(id dxo.InstalledFileID) error {

	inst.TrashService.OnDelete()

	db := inst.Agent.DB()
	m := inst.model()
	res := db.Delete(m, id)
	return res.Error
}
