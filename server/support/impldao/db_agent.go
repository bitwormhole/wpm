package impldao

import (
	"time"

	"github.com/bitwormhole/starter-gorm/datasource"
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
	markup.Component `id:"GormDBAgent" initMethod:"Init"`

	Source datasource.Source `inject:"#gorm-datasource-default"`
}

func (inst *GormDBAgentImpl) _Impl() GormDBAgent {
	return inst
}

// Init ...
func (inst *GormDBAgentImpl) Init() error {
	inst.startAutoMakeTables()
	return nil
}

func (inst *GormDBAgentImpl) startAutoMakeTables() {
	go func() {
		time.Sleep(time.Second * 3)
		err := inst.autoMakeTables()
		if err != nil {
			panic(err)
		}
	}()
}

// DB ...
func (inst *GormDBAgentImpl) DB() *gorm.DB {
	return inst.Source.DB()
}

func (inst *GormDBAgentImpl) autoMakeTables() error {

	var e error
	db := inst.DB()
	models := make([]interface{}, 0)

	models = append(models, &entity.Executable{})
	models = append(models, &entity.Intent{})
	models = append(models, &entity.Project{})
	models = append(models, &entity.LocalRepository{})
	models = append(models, &entity.MainRepository{})
	models = append(models, &entity.RemoteRepository{})

	for _, mod := range models {
		err := db.AutoMigrate(mod)
		if err != nil {
			vlog.Warn(err)
			e = err
		}
	}
	return e
}
