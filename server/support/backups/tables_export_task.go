package backups

import (
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/backup"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

type tablesExportTask struct {
	db   *gorm.DB
	view *backup.VO
}

func (inst *tablesExportTask) exportAll() error {

	all := make([]func() error, 0)

	all = append(all, inst.exportExecutableTable)
	all = append(all, inst.exportIntentTemplateTable)
	all = append(all, inst.exportLocalRepositoryTable)
	all = append(all, inst.exportLocationTable)
	all = append(all, inst.exportMediaTable)
	all = append(all, inst.exportProjectTable)
	all = append(all, inst.exportProjectTypeTable)
	all = append(all, inst.exportRemoteRepositoryTable)
	all = append(all, inst.exportSettingTable)

	for _, fn := range all {
		err := fn()
		if err != nil {
			vlog.Warn(err)
		}
	}

	return nil
}

func (inst *tablesExportTask) exportExecutableTable() error {

	db := inst.db
	view := inst.view
	table := view.ExecutableTable
	table = make([]*entity.Executable, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.ExecutableTable = table
	}
	return err
}

func (inst *tablesExportTask) exportIntentTemplateTable() error {
	db := inst.db
	view := inst.view
	table := view.IntentTemplateTable
	table = make([]*entity.IntentTemplate, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.IntentTemplateTable = table
	}
	return err
}

func (inst *tablesExportTask) exportRemoteRepositoryTable() error {

	db := inst.db
	view := inst.view
	table := view.RemoteRepositoryTable
	table = make([]*entity.RemoteRepository, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.RemoteRepositoryTable = table
	}
	return err
}

func (inst *tablesExportTask) exportLocalRepositoryTable() error {

	db := inst.db
	view := inst.view
	table := view.LocalRepositoryTable
	table = make([]*entity.LocalRepository, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.LocalRepositoryTable = table
	}
	return err
}

func (inst *tablesExportTask) exportLocationTable() error {

	db := inst.db
	view := inst.view
	table := view.LocationTable
	table = make([]*entity.Location, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.LocationTable = table
	}
	return err
}

func (inst *tablesExportTask) exportProjectTable() error {

	db := inst.db
	view := inst.view
	table := view.ProjectTable
	table = make([]*entity.Project, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.ProjectTable = table
	}
	return err
}

func (inst *tablesExportTask) exportMediaTable() error {

	db := inst.db
	view := inst.view
	table := view.MediaTable
	table = make([]*entity.Media, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.MediaTable = table
	}
	return err
}

func (inst *tablesExportTask) exportSettingTable() error {

	db := inst.db
	view := inst.view
	table := view.SettingTable
	table = make([]*entity.Setting, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.SettingTable = table
	}
	return err
}

func (inst *tablesExportTask) exportProjectTypeTable() error {

	db := inst.db
	view := inst.view
	table := view.ContentTypeTable
	table = make([]*entity.ContentType, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.ContentTypeTable = table
	}
	return err
}
