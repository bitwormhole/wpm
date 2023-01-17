package implservice

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentTemplateServiceImpl ...
type IntentTemplateServiceImpl struct {
	markup.Component `id:"IntentTemplateService"`

	IntentTempDAO dao.IntentTemplateDAO `inject:"#IntentTemplateDAO"`
}

func (inst *IntentTemplateServiceImpl) _Impl() service.IntentTemplateService {
	return inst
}

func (inst *IntentTemplateServiceImpl) dto2entity(o1 *dto.IntentTemplate) (*entity.IntentTemplate, error) {
	o2 := &entity.IntentTemplate{}
	o2.ID = o1.ID

	// o2.Exe  = o1.Executable
	// o2.cli = o1.Executable
	// o2.Executable = o1.Executable
	// todo ...
	return o2, nil
}

func (inst *IntentTemplateServiceImpl) entity2dto(o1 *entity.IntentTemplate) (*dto.IntentTemplate, error) {
	o2 := &dto.IntentTemplate{}
	o2.ID = o1.ID
	// o2.Executable = o1.Executable
	// todo ...
	return o2, nil
}

func (inst *IntentTemplateServiceImpl) ListAll(ctx context.Context) ([]*dto.IntentTemplate, error) {
	src, err := inst.IntentTempDAO.ListAll()
	if err != nil {
		return nil, err
	}
	dst := make([]*dto.IntentTemplate, 0)
	for _, item1 := range src {
		item2, err := inst.entity2dto(item1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}

func (inst *IntentTemplateServiceImpl) Find(ctx context.Context, id dxo.IntentTemplateID) (*dto.IntentTemplate, error) {
	return nil, errors.New("no impl")
}

func (inst *IntentTemplateServiceImpl) Insert(ctx context.Context, o *dto.IntentTemplate) (*dto.IntentTemplate, error) {
	return nil, errors.New("no impl")
}

func (inst *IntentTemplateServiceImpl) Update(ctx context.Context, id dxo.IntentTemplateID, o *dto.IntentTemplate) (*dto.IntentTemplate, error) {
	return nil, errors.New("no impl")
}

func (inst *IntentTemplateServiceImpl) Remove(ctx context.Context, id dxo.IntentTemplateID) error {
	return errors.New("no impl")
}
