package trash

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/service"
)

// ImpTrashService ...
//starter:component id:"TrashService" class:""
type ImpTrashService struct {
	markup.Component `id:"TrashService" initMethod:"Init"`

	TrashDao dao.TrashDAO `inject:"#TrashDAO"`

	AutoClean   bool `inject:"${wpm.auto-clean-trash.enabled}"`
	countRemove int
}

func (inst *ImpTrashService) _Impl() service.TrashService {
	return inst
}

// Init ...
func (inst *ImpTrashService) Init() error {
	inst.countRemove = 1
	return nil
}

// Clean ...
func (inst *ImpTrashService) Clean() error {
	all, err := inst.TrashDao.ListAll()
	if err != nil {
		return err
	}
	return inst.TrashDao.Delete(all...)
}

// OnInsert ...
func (inst *ImpTrashService) OnInsert() {
	if inst.AutoClean && (inst.countRemove > 0) {
		inst.countRemove = 0
		err := inst.Clean()
		if err != nil {
			vlog.Warn(err)
		}
	}
}

// OnDelete ...
func (inst *ImpTrashService) OnDelete() {
	inst.countRemove++
}

// EnableAutoCleanBeforeInsert ...
func (inst *ImpTrashService) EnableAutoCleanBeforeInsert(en bool) {
	inst.AutoClean = en
}
