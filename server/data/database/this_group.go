package database

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/starter-go/libgorm"
	"gorm.io/gorm"
)

// ThisGroup ...
type ThisGroup struct {

	//starter:component
	_as func(libgorm.GroupRegistry, dxo.DatabaseAgent) //starter:as(".","#")

	Enabled bool   //starter:inject("${datagroup.wpm.enabled}")
	Alias   string //starter:inject("${datagroup.wpm.alias}")
	URI     string //starter:inject("${datagroup.wpm.uri}")
	Prefix  string //starter:inject("${datagroup.wpm.table-name-prefix}")
	Source  string //starter:inject("${datagroup.wpm.datasource}")

	SourceManager libgorm.DataSourceManager //starter:inject("#")

	plist []any // the prototypes list
	inner libgorm.DataSourceAgent
	// info  dxo.DataGroupInfo
}

func (inst *ThisGroup) _impl() (libgorm.GroupRegistry, libgorm.Group, dxo.DatabaseAgent) {
	return inst, inst, inst
}

// Groups ...
func (inst *ThisGroup) Groups() []*libgorm.GroupRegistration {

	inst.plist = entity.ListAllTypes()

	r1 := &libgorm.GroupRegistration{
		Enabled: true,
		Alias:   inst.Alias,
		URI:     inst.URI,
		Prefix:  inst.Prefix,
		Source:  inst.Source,
		Group:   inst,
	}

	return []*libgorm.GroupRegistration{r1}
}

// Prototypes ...
func (inst *ThisGroup) Prototypes() []any {
	return inst.plist
}

// DB ...
func (inst *ThisGroup) DB(db *gorm.DB) *gorm.DB {
	if !inst.inner.Ready() {
		srcManager := inst.SourceManager
		srcName := inst.Source
		inst.inner.Init(srcManager, srcName)
	}
	return inst.inner.DB(db)
}
