package backups

import (
	"encoding/json"
	"fmt"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/backup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/service"
)

// ImpBackupServiceDAO ...
type ImpBackupServiceDAO struct {
	markup.Component `id:"wpm-database-backup-dao" class:""`

	Agent dbagent.GormDBAgent `inject:"#GormDBAgent"`
}

func (inst *ImpBackupServiceDAO) _Impl() dao.Backup {
	return inst
}

func (inst *ImpBackupServiceDAO) exportThisExeInfo(o *backup.StoreVO) error {
	m := service.GetAppModule()
	o.WPM.Name = m.GetName()
	o.WPM.Version = m.GetVersion()
	o.WPM.Revision = m.GetRevision()
	return nil
}

func (inst *ImpBackupServiceDAO) exportTables(o *backup.StoreVO) error {
	db := inst.Agent.DB()
	task := &tablesExportTask{
		view: o,
		db:   db,
	}
	return task.exportAll()
}

// Export ...
func (inst *ImpBackupServiceDAO) Export(file afs.Path) error {

	o := &backup.StoreVO{}
	err := inst.exportTables(o)
	if err != nil {
		return err
	}

	err = inst.exportThisExeInfo(o)
	if err != nil {
		return err
	}

	data, err := json.Marshal(o)
	if err != nil {
		return err
	}

	opt := &afs.Options{
		Create: true,
		Mkdirs: false,
	}
	if file.Exists() {
		return fmt.Errorf("the export target file is exists. path = " + file.GetPath())
	}
	return file.GetIO().WriteBinary(data, opt)
}

// Import ...
func (inst *ImpBackupServiceDAO) Import(file afs.Path) error {
	view := &backup.StoreVO{}
	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, view)
	if err != nil {
		return err
	}
	db := inst.Agent.DB()
	task := &tablesImportTask{
		db:   db,
		view: view,
	}
	return task.importAll()
}
