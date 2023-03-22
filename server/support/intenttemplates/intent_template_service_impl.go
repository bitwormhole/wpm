package intenttemplates

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
	o2.UUID = o1.UUID

	o2.Name = o1.Name
	o2.Description = o1.Description

	o2.Arguments = dxo.NewStringListCRLF(o1.Arguments)
	o2.Command = o1.Command
	o2.Env = dxo.NewStringMap(o1.Env)
	o2.WD = o1.WD

	// todo ...
	return o2, nil
}

func (inst *IntentTemplateServiceImpl) entity2dto(o1 *entity.IntentTemplate) (*dto.IntentTemplate, error) {
	o2 := &dto.IntentTemplate{}
	o2.ID = o1.ID
	o2.UUID = o1.UUID

	o2.Name = o1.Name
	o2.Description = o1.Description

	o2.Arguments = o1.Arguments.Array()
	o2.Command = o1.Command
	o2.Env = o1.Env.Map()
	o2.WD = o1.WD

	// todo ...
	return o2, nil
}

func (inst *IntentTemplateServiceImpl) entity2dtoList(src []*entity.IntentTemplate) ([]*dto.IntentTemplate, error) {
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

// ListAll ...
func (inst *IntentTemplateServiceImpl) ListAll(ctx context.Context) ([]*dto.IntentTemplate, error) {
	src, err := inst.IntentTempDAO.ListAll()
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(src)
}

func (inst *IntentTemplateServiceImpl) Find(ctx context.Context, id dxo.IntentTemplateID) (*dto.IntentTemplate, error) {
	return nil, errors.New("no impl")
}

// Insert ...
func (inst *IntentTemplateServiceImpl) Insert(ctx context.Context, o1 *dto.IntentTemplate) (*dto.IntentTemplate, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.IntentTempDAO.Insert(o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3)
}

// Update ...
func (inst *IntentTemplateServiceImpl) Update(ctx context.Context, id dxo.IntentTemplateID, o1 *dto.IntentTemplate) (*dto.IntentTemplate, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.IntentTempDAO.Update(id, o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3)
}

// Remove ...
func (inst *IntentTemplateServiceImpl) Remove(ctx context.Context, id dxo.IntentTemplateID) error {
	return inst.IntentTempDAO.Remove(id)
}
