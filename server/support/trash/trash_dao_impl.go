package trash

import (
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/vo"
	"gorm.io/gorm"
)

// ImpTrashDao ...
type ImpTrashDao struct {
	markup.Component `id:"TrashDAO"`

	Agent dbagent.GormDBAgent `inject:"#GormDBAgent"`
}

func (inst *ImpTrashDao) _Impl() dao.TrashDAO {
	return inst
}

// ListAll ...
func (inst *ImpTrashDao) ListAll() ([]*entity.Holder, error) {

	steps := make([]func(t *vo.Trash) error, 0)
	trash := &vo.Trash{}
	elist := &utils.ErrorList{}
	trash.TrashItems = make([]*entity.Holder, 0)

	steps = append(steps, inst.listTrashedPackages)
	steps = append(steps, inst.listTrashedProjects)
	steps = append(steps, inst.listTrashedRepositories)

	for _, fn := range steps {
		err := fn(trash)
		elist.Append(err)
	}

	return trash.TrashItems, elist.First()
}

func (inst *ImpTrashDao) listTrashedPackages(t *vo.Trash) error {
	db := inst.Agent.DB()
	list := make([]*entity.SoftwarePackage, 0)
	limit := inst.getTrashAtLimit()
	res := db.Unscoped().Where("deleted_at > ?", limit).Find(&list)
	for _, item := range list {
		h := &entity.Holder{}
		h.SetSoftwarePackage(item)
		t.TrashItems = append(t.TrashItems, h)
	}
	return res.Error
}

func (inst *ImpTrashDao) listTrashedProjects(t *vo.Trash) error {
	db := inst.Agent.DB()
	list := make([]*entity.Project, 0)
	limit := inst.getTrashAtLimit()
	res := db.Unscoped().Where("deleted_at > ?", limit).Find(&list)
	for _, item := range list {
		h := &entity.Holder{}
		h.SetProject(item)
		t.TrashItems = append(t.TrashItems, h)
	}
	return res.Error
}

func (inst *ImpTrashDao) listTrashedRepositories(t *vo.Trash) error {
	db := inst.Agent.DB()
	list := make([]*entity.LocalRepository, 0)
	limit := inst.getTrashAtLimit()
	res := db.Unscoped().Where("deleted_at > ?", limit).Find(&list)
	for _, item := range list {
		h := &entity.Holder{}
		h.SetLocalRepository(item)
		t.TrashItems = append(t.TrashItems, h)
	}
	return res.Error
}

func (inst *ImpTrashDao) getTrashAtLimit() gorm.DeletedAt {
	t := util.NewTimeWithInt64(7 * 24 * 3600 * 1000)
	limit := gorm.DeletedAt{}
	limit.Time = t.GetTime()
	limit.Valid = true
	return limit
}

// Delete ...
func (inst *ImpTrashDao) Delete(holders ...*entity.Holder) error {
	db := inst.Agent.DB()
	elist := utils.ErrorList{}
	for _, h := range holders {
		id := h.ID
		o := h.Get()
		if id == 0 || o == nil {
			continue
		}
		vlog.Info("delete in-trash object:", h.TableName, ":", id)
		res := db.Unscoped().Delete(o, id)
		elist.Append(res.Error)
	}
	return elist.First()
}

// Recover ...
func (inst *ImpTrashDao) Recover(h ...*entity.Holder) error {
	return fmt.Errorf("no impl")
}
