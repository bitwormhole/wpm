package implservice

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/app/data/dao"
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/bitwormhole/wpm/app/data/entity"
	"github.com/bitwormhole/wpm/app/service"
	"github.com/bitwormhole/wpm/app/web/dto"
)

// IntentServiceImpl ...
type IntentServiceImpl struct {
	markup.Component `id:"IntentService"`

	IntentDAO dao.IntentDAO `inject:"#IntentDAO"`
}

func (inst *IntentServiceImpl) _Impl() service.IntentService {
	return inst
}

func (inst *IntentServiceImpl) dto2entity(o1 *dto.Intent) (*entity.Intent, error) {
	o2 := &entity.Intent{}
	o2.ID = o1.ID
	o2.Executable = o1.Executable
	// todo ...
	return o2, nil
}

func (inst *IntentServiceImpl) entity2dto(o1 *entity.Intent) (*dto.Intent, error) {
	o2 := &dto.Intent{}
	o2.ID = o1.ID
	o2.Executable = o1.Executable
	// todo ...
	return o2, nil
}

func (inst *IntentServiceImpl) ListAll(ctx context.Context) ([]*dto.Intent, error) {
	src, err := inst.IntentDAO.ListAll()
	if err != nil {
		return nil, err
	}
	dst := make([]*dto.Intent, 0)
	for _, item1 := range src {
		item2, err := inst.entity2dto(item1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}

func (inst *IntentServiceImpl) Find(ctx context.Context, id dxo.IntentID) (*dto.Intent, error) {
	return nil, errors.New("no impl")
}

func (inst *IntentServiceImpl) Insert(ctx context.Context, o *dto.Intent) (*dto.Intent, error) {
	return nil, errors.New("no impl")
}

func (inst *IntentServiceImpl) Update(ctx context.Context, id dxo.IntentID, o *dto.Intent) (*dto.Intent, error) {
	return nil, errors.New("no impl")
}

func (inst *IntentServiceImpl) Remove(ctx context.Context, id dxo.IntentID) error {
	return errors.New("no impl")
}
