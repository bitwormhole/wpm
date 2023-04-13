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
	table := view.ExecutableTable
	table = make([]*entity.Executable, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.ExecutableTable = table
	}
	return err
}

func (inst *tablesExportTask) exportIntentTemplates() error {
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

func (inst *tablesExportTask) exportRemoteRepositories() error {

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

func (inst *tablesExportTask) exportLocalRepositories() error {

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

func (inst *tablesExportTask) exportLocations() error {

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

func (inst *tablesExportTask) exportProjects() error {

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

func (inst *tablesExportTask) exportMediae() error {

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

func (inst *tablesExportTask) exportSettings() error {

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

func (inst *tablesExportTask) exportContentTypes() error {

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

func (inst *tablesExportTask) exportNamespaces() error {

	db := inst.db
	view := inst.view
	table := view.SourceTable
	table = make([]*entity.Namespace, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.SourceTable = table
	}
	return err
}

func (inst *tablesExportTask) exportPackages() error {

	db := inst.db
	view := inst.view
	table := view.PackageTable
	table = make([]*entity.SoftwarePackage, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.PackageTable = table
	}
	return err
}

func (inst *tablesExportTask) exportUsers() error {

	db := inst.db
	view := inst.view
	table := view.UserTable
	table = make([]*entity.User, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.UserTable = table
	}
	return err
}

func (inst *tablesExportTask) exportWorktrees() error {

	db := inst.db
	view := inst.view
	table := view.WorktreeTable
	table = make([]*entity.Worktree, 0)

	res := db.Find(&table)
	err := res.Error
	if err == nil {
		view.WorktreeTable = table
	}
	return err
}
