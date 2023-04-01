package namespaces

import (
	"encoding/json"
	"fmt"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/backup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// ImpNamespaceDao ... 目前，entity.Namespace 实际上是从资源文件里加载的
type ImpNamespaceDao struct {
	markup.Component `id:"NamespaceDAO"`

	AC application.Context `inject:"context"`
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
	const (
		name = "res:///namespaces.json"
	)
	data, err := inst.AC.GetResources().GetBinary(name)
	if err != nil {
		return nil, err
	}
	vo := &backup.VO{}
	err = json.Unmarshal(data, vo)
	if err != nil {
		return nil, err
	}
	return vo.NamespaceTable, nil
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
