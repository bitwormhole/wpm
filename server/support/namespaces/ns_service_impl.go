package namespaces

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ImpNamespaceService ...
type ImpNamespaceService struct {
	markup.Component `id:"NamespaceService"`

	MyDAO dao.NamespaceDAO `inject:"#NamespaceDAO"`
}

func (inst *ImpNamespaceService) _Impl() service.NamespaceService {
	return inst
}

func (inst *ImpNamespaceService) dto2entity(o1 *dto.Namespace) (*entity.Namespace, error) {

	o2 := &entity.Namespace{}

	o2.ID = o1.ID
	o2.UUID = o1.UUID
	o2.URL = o1.URL
	o2.URN = o1.URN
	o2.Name = o1.Name
	o2.OS = o1.OS
	o2.Arch = o1.Arch

	return o2, nil
}

func (inst *ImpNamespaceService) entity2dto(o1 *entity.Namespace) (*dto.Namespace, error) {

	o2 := &dto.Namespace{}

	o2.ID = o1.ID
	o2.UUID = o1.UUID
	o2.URL = o1.URL
	o2.URN = o1.URN
	o2.Name = o1.Name
	o2.OS = o1.OS
	o2.Arch = o1.Arch

	return o2, nil
}

func (inst *ImpNamespaceService) entity2dtoList(src []*entity.Namespace) ([]*dto.Namespace, error) {
	dst := make([]*dto.Namespace, 0)
	for _, o1 := range src {
		o2, err := inst.entity2dto(o1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, o2)
	}
	return dst, nil
}

// Find ...
func (inst *ImpNamespaceService) Find(ctx context.Context, id dxo.NamespaceID) (*dto.Namespace, error) {
	o1, err := inst.MyDAO.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1)
}

// ListAll ...
func (inst *ImpNamespaceService) ListAll(ctx context.Context) ([]*dto.Namespace, error) {
	list1, err := inst.MyDAO.ListAll()
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(list1)
}

func (inst *ImpNamespaceService) Insert(ctx context.Context, o *dto.Namespace) (*dto.Namespace, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *ImpNamespaceService) Update(ctx context.Context, id dxo.NamespaceID, o *dto.Namespace) (*dto.Namespace, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *ImpNamespaceService) Remove(ctx context.Context, id dxo.NamespaceID) error {
	return fmt.Errorf("no impl")
}
