package support

import (
	"log"
	"os"
	"time"

	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// WpmDataSource ...
type WpmDataSource struct {
	markup.Component `class:"starter-gorm-source-registry"`

	DM             datasource.DriverManager `inject:"#starter-gorm-driver-manager"`
	AppDataService service.AppDataService   `inject:"#AppDataService"`
	AboutService   service.AboutService     `inject:"#AboutService"`

	Driver   string `inject:"${datasource.wpm.driver}"`
	Host     string `inject:"${datasource.wpm.host}"`
	Port     int    `inject:"${datasource.wpm.port}"`
	UserName string `inject:"${datasource.wpm.username}"`
	Password string `inject:"${datasource.wpm.password}"`
	Database string `inject:"${datasource.wpm.database}"`

	cachedDB *gorm.DB
}

func (inst *WpmDataSource) _Impl() (datasource.SourceRegistry, datasource.Source) {
	return inst, inst
}

// ListSources ...
func (inst *WpmDataSource) ListSources() []*datasource.SourceRegistration {
	reg := &datasource.SourceRegistration{
		Name:   "wpm",
		Source: inst,
	}
	return []*datasource.SourceRegistration{reg}
}

// DB ...
func (inst *WpmDataSource) DB() (*gorm.DB, error) {
	db := inst.cachedDB
	if db != nil {
		return db, nil
	}
	db, err := inst.open()
	if err != nil {
		return nil, err
	}
	inst.cachedDB = db
	return db, nil
}

func (inst *WpmDataSource) open() (*gorm.DB, error) {

	ads := inst.AppDataService
	if !ads.Ready() {
		err := ads.Setup()
		if err != nil {
			return nil, err
		}
	}

	cfg := inst.config()
	src, err := inst.DM.OpenSource(cfg)
	if err != nil {
		return nil, err
	}

	db, err := src.DB()
	if err != nil {
		return nil, err
	}

	db = inst.configLogger(db)
	return db, nil
}

func (inst *WpmDataSource) configLogger(db1 *gorm.DB) *gorm.DB {

	release := inst.AboutService.IsRelease()
	if !release {
		return db1
	}

	// as release mode
	log := inst.makeLoggerForRelease()
	db2 := db1.Session(&gorm.Session{Logger: log})
	if db2.Error != nil {
		return db1
	}
	return db2
}

func (inst *WpmDataSource) makeLoggerForRelease() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)
	return newLogger
}

func (inst *WpmDataSource) config() *datasource.Configuration {
	cfg := &datasource.Configuration{}
	cfg.Name = "wpm"
	cfg.Driver = inst.Driver
	cfg.Host = inst.Host
	cfg.Port = inst.Port
	cfg.Database = inst.Database
	cfg.Username = inst.UserName
	cfg.Password = inst.Password

	if cfg.Driver == "sqlite" {
		inst.configForSQLite(cfg)
	}

	return cfg
}

func (inst *WpmDataSource) configForSQLite(c *datasource.Configuration) {
	path := inst.AppDataService.GetSQLiteDBFile()
	c.Database = path
}

// Close ...
func (inst *WpmDataSource) Close() error {
	return nil
}
