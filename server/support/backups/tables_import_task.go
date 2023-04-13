package backups

import (
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/backup"
	"gorm.io/gorm"
)

type tablesImportTask struct {
	db   *gorm.DB
	view *backup.StoreVO
}

func (inst *tablesImportTask) importAll() error {

	all := make([]func() error, 0)

	all = append(all, inst.importExecutables)
	all = append(all, inst.importLocalRepositories)
	all = append(all, inst.importLocations)
	all = append(all, inst.importMediae)
	all = append(all, inst.importContentTypes)
	all = append(all, inst.importProjects)
	all = append(all, inst.importIntentTemplate)

	for _, fn := range all {
		err := fn()
		if err != nil {
			vlog.Warn(err)
		}
	}

	return nil
}

func (inst *tablesImportTask) handleError(err error) {
	if err == nil {
		return
	}
	vlog.Warn(err)
}

func (inst *tablesImportTask) importProjects() error {
	db := inst.db
	list := inst.view.ProjectTable
	for _, item := range list {
		res := db.FirstOrCreate(item, item.ID)
		inst.handleError(res.Error)
	}
	return nil
}

func (inst *tablesImportTask) importIntentTemplate() error {
	db := inst.db
	list := inst.view.IntentTemplateTable
	for _, item := range list {
		res := db.FirstOrCreate(item, item.ID)
		inst.handleError(res.Error)
	}
	return nil
}

func (inst *tablesImportTask) importContentTypes() error {
	db := inst.db
	list := inst.view.ContentTypeTable
	for _, item := range list {
		res := db.FirstOrCreate(item, item.ID)
		inst.handleError(res.Error)
	}
	return nil
}

func (inst *tablesImportTask) importLocalRepositories() error {
	db := inst.db
	list := inst.view.LocalRepositoryTable
	for _, item := range list {
		res := db.FirstOrCreate(item, item.ID)
		inst.handleError(res.Error)
	}
	return nil
}

func (inst *tablesImportTask) importLocations() error {
	db := inst.db
	list := inst.view.LocationTable
	for _, item := range list {
		res := db.FirstOrCreate(item, item.ID)
		inst.handleError(res.Error)
	}
	return nil
}

func (inst *tablesImportTask) importMediae() error {
	db := inst.db
	list := inst.view.MediaTable
	for _, item := range list {
		res := db.FirstOrCreate(item, item.ID)
		inst.handleError(res.Error)
	}
	return nil
}

func (inst *tablesImportTask) importExecutables() error {
	db := inst.db
	list := inst.view.ExecutableTable
	for _, item := range list {
		res := db.FirstOrCreate(item, item.ID)
		inst.handleError(res.Error)
	}
	return nil
}
