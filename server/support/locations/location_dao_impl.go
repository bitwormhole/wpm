package locations

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
)

// ImpLocationDao ...
type ImpLocationDao struct {
	markup.Component `id:"LocationDAO"`

	Agent          dbagent.GormDBAgent    `inject:"#GormDBAgent"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *ImpLocationDao) _Impl() dao.LocationDAO {
	return inst
}

func (inst *ImpLocationDao) model() *entity.Location {
	return &entity.Location{}
}

func (inst *ImpLocationDao) modelList() []*entity.Location {
	return make([]*entity.Location, 0)
}

// Find ...
func (inst *ImpLocationDao) Find(id dxo.LocationID) (*entity.Location, error) {
	db := inst.Agent.DB()
	o2 := inst.model()
	res := db.First(o2, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return o2, nil
}

// FindByPath ...
func (inst *ImpLocationDao) FindByPath(path string) (*entity.Location, error) {

	path = strings.TrimSpace(path)
	if path == "" {
		return nil, fmt.Errorf("param:`path` is empty")
	}

	erlist := &utils.ErrorList{}
	db := inst.Agent.DB()
	o := inst.model()
	fields := o.ListPathFields()

	for _, field := range fields {
		res := db.Where(field+"=?", path).First(o)
		if res.Error == nil {
			return o, nil
		}
		erlist.Append(res.Error)
	}

	return nil, erlist.First()
}

// ListAll ...
func (inst *ImpLocationDao) ListAll() ([]*entity.Location, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// Insert ...
func (inst *ImpLocationDao) Insert(o *entity.Location) (*entity.Location, error) {

	// compute fields
	o.ID = 0
	o.UUID = inst.UUIDGenService.GenerateUUID(o.Path + "|entity.Location|")

	// save
	db := inst.Agent.DB()
	res := db.Create(o)
	if res.Error != nil {
		return nil, res.Error
	}

	return o, nil
}

// Update ...
func (inst *ImpLocationDao) Update(id dxo.LocationID, o1 *entity.Location) (*entity.Location, error) {

	db := inst.Agent.DB()
	o2 := inst.model()
	res := db.First(o2, id)
	if res.Error != nil {
		return nil, res.Error
	}

	// todo ...
	o2.Path = o1.Path
	o2.Class = o1.Class
	o2.AsDir = o1.AsDir
	o2.AsFile = o1.AsFile

	// o2.TypeID = o1.TypeID
	// o2.TypeKey = o1.TypeKey
	// o2.TypeName = o1.TypeName

	res = db.Save(o2)
	if res.Error != nil {
		return nil, res.Error
	}
	return o2, nil
}

// Remove ...
func (inst *ImpLocationDao) Remove(id dxo.LocationID) error {
	db := inst.Agent.DB()
	m := inst.model()
	res := db.Delete(m, id)
	return res.Error
}
