package namespaces

import (
	"fmt"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ImpNamespaceDao ... 目前，entity.Namespace 实际上是从资源文件里加载的
type ImpNamespaceDao struct {
	markup.Component `id:"NamespaceDAO"`

	AC            application.Context   `inject:"context"`
	PresetService service.PresetService `inject:"#PresetService"`
}

func (inst *ImpNamespaceDao) _Impl() dao.NamespaceDAO {
	return inst
}

// Find ...
func (inst *ImpNamespaceDao) Find(id dxo.NamespaceID) (*entity.Namespace, error) {
	list, err := inst.ListAll()
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		if item.ID == id {
			return item, nil
		}
	}
	return nil, fmt.Errorf("namespace not found, by id:%v", id)
}

// ListAll ...
func (inst *ImpNamespaceDao) ListAll() ([]*entity.Namespace, error) {
	all, err := inst.PresetService.GetPresets()
	if err != nil {
		return nil, err
	}
	src := all.Namespaces
	dst := make([]*entity.Namespace, 0)
	for _, item1 := range src {
		item2 := inst.dto2entity(item1)
		dst = append(dst, item2)
	}
	return dst, nil
}

func (inst *ImpNamespaceDao) dto2entity(src *dto.Namespace) *entity.Namespace {

	dst := &entity.Namespace{}

	dst.ID = src.ID
	dst.UUID = src.UUID
	dst.URN = src.URN
	dst.Name = src.Name
	dst.URL = src.URL
	dst.OS = src.OS
	dst.Arch = src.Arch

	return dst
}

func (inst *ImpNamespaceDao) Insert(o *entity.Namespace) (*entity.Namespace, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *ImpNamespaceDao) Update(id dxo.NamespaceID, o *entity.Namespace) (*entity.Namespace, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *ImpNamespaceDao) Remove(id dxo.NamespaceID) error {
	return fmt.Errorf("no impl")
}
