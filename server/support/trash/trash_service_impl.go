package trash

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/service"
)

// ImpTrashService ...
//starter:component id:"TrashService" class:""
type ImpTrashService struct {
	markup.Component `id:"TrashService"`

	TrashDao dao.TrashDAO `inject:"#TrashDAO"`
}

func (inst *ImpTrashService) _Impl() service.TrashService {
	return inst
}

// Clean ...
func (inst *ImpTrashService) Clean() error {
	all, err := inst.TrashDao.ListAll()
	if err != nil {
		return err
	}
	return inst.TrashDao.Delete(all...)
}
