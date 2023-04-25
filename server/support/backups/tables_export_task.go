package backups

import (
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/backup"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

type tablesExportTask struct {
	db   *gorm.DB
	view *backup.StoreVO
}

func (inst *tablesExportTask) exportAll() error {

	all := make([]func() error, 0)

	all = append(all, inst.exportContentTypes)
	all = append(all, inst.exportExecutables)
	all = append(all, inst.exportIntentTemplates)
	all = append(all, inst.exportLocalRepositories)
	all = append(all, inst.exportLocations)
	all = append(all, inst.exportMediae)
	all = append(all, inst.exportNamespaces)
	all = append(all, inst.exportPackages)
	all = append(all, inst.exportProjects)
	all = append(all, inst.exportRemoteRepositories)
	all = append(all, inst.exportSettings)
	all = append(all, inst.exportUsers)
	all = append(all, inst.exportWorktrees)

	for _, fn := range all {
		err := fn()
		if err != nil {
			vlog.Warn(err)
		}
	}

	return nil
}

func (inst *tablesExportTask) exportExecutables() error {

	db := inst.db
	view := inst.view
	table := view.Tables.ExecutableTable
	table = make([]*entity.Executable, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.Tables.ExecutableTable = table
	}
	return err
}

func (inst *tablesExportTask) exportIntentTemplates() error {
	db := inst.db
	view := inst.view
	table := view.Tables.IntentTemplateTable
	table = make([]*entity.IntentTemplate, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.Tables.IntentTemplateTable = table
	}
	return err
}

func (inst *tablesExportTask) exportRemoteRepositories() error {

	db := inst.db
	view := inst.view
	table := view.Tables.RemoteRepositoryTable
	table = make([]*entity.RemoteRepository, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.Tables.RemoteRepositoryTable = table
	}
	return err
}

func (inst *tablesExportTask) exportLocalRepositories() error {

	db := inst.db
	view := inst.view
	table := view.Tables.LocalRepositoryTable
	table = make([]*entity.LocalRepository, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.Tables.LocalRepositoryTable = table
	}
	return err
}

func (inst *tablesExportTask) exportLocations() error {

	db := inst.db
	view := inst.view
	table := view.Tables.LocationTable
	table = make([]*entity.Location, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.Tables.LocationTable = table
	}
	return err
}

func (inst *tablesExportTask) exportProjects() error {

	db := inst.db
	view := inst.view
	table := view.Tables.ProjectTable
	table = make([]*entity.Project, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.Tables.ProjectTable = table
	}
	return err
}

func (inst *tablesExportTask) exportMediae() error {

	db := inst.db
	view := inst.view
	table := view.Tables.MediaTable
	table = make([]*entity.Media, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.Tables.MediaTable = table
	}
	return err
}

func (inst *tablesExportTask) exportSettings() error {

	db := inst.db
	view := inst.view
	table := view.Tables.SettingTable
	table = make([]*entity.Setting, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.Tables.SettingTable = table
	}
	return err
}

func (inst *tablesExportTask) exportContentTypes() error {

	db := inst.db
	view := inst.view
	table := view.Tables.ContentTypeTable
	table = make([]*entity.ContentType, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.Tables.ContentTypeTable = table
	}
	return err
}

func (inst *tablesExportTask) exportNamespaces() error {

	db := inst.db
	view := inst.view
	table := view.Tables.SourceTable
	table = make([]*entity.Namespace, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.Tables.SourceTable = table
	}
	return err
}

func (inst *tablesExportTask) exportPackages() error {

	db := inst.db
	view := inst.view
	table := view.Tables.PackageTable
	table = make([]*entity.SoftwarePackage, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.Tables.PackageTable = table
	}
	return err
}

func (inst *tablesExportTask) exportUsers() error {

	db := inst.db
	view := inst.view
	table := view.Tables.UserTable
	table = make([]*entity.User, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.Tables.UserTable = table
	}
	return err
}

func (inst *tablesExportTask) exportWorktrees() error {

	db := inst.db
	view := inst.view
	table := view.Tables.WorktreeTable
	table = make([]*entity.Worktree, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.Tables.WorktreeTable = table
	}
	return err
}
