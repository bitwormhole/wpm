package impldao

import (
	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// GormDBAgent ...
type GormDBAgent interface {
	DB() *gorm.DB
}

////////////////////////////////////////////////////////////////////////////////

// GormDBAgentImpl ...
type GormDBAgentImpl struct {
	markup.Component `id:"GormDBAgent" class:"life"`

	Sources datasource.SourceManager `inject:"#starter-gorm-source-manager"`

	cached *gorm.DB
}

func (inst *GormDBAgentImpl) _Impl() (GormDBAgent, application.LifeRegistry) {
	return inst, inst
}

// GetLifeRegistration ...
func (inst *GormDBAgentImpl) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnInit: inst.onInit,
	}
}

// Init ...
func (inst *GormDBAgentImpl) onInit() error {
	_, err := inst.get()
	return err
}

// DB ...
func (inst *GormDBAgentImpl) DB() *gorm.DB {
	db, err := inst.get()
	if err != nil {
		panic(db)
	}
	return db
}

func (inst *GormDBAgentImpl) load() (*gorm.DB, error) {

	src, err := inst.Sources.GetSource("wpm")
	if err != nil {
		return nil, err
	}

	db, err := src.DB()
	if err != nil {
		return nil, err
	}

	err = inst.autoMakeTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (inst *GormDBAgentImpl) get() (*gorm.DB, error) {
	db := inst.cached
	if db == nil {
		o, err := inst.load()
		if err != nil {
			return nil, err
		}
		db = o
		inst.cached = o
	}
	return db, nil
}

func (inst *GormDBAgentImpl) autoMakeTables(db *gorm.DB) error {
	var e error
	models := entity.ListPrototypes()
	for _, mod := range models {
		err := db.AutoMigrate(mod)
		if err != nil {
			vlog.Warn(err)
			e = err
		}
	}
	return e
}
