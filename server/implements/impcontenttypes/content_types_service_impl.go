package impcontenttypes

import (
	"context"
	"strings"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/contenttypes"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(contenttypes.Service) //starter:as("#")

	Dao contenttypes.DAO //starter:inject("#")

	convertor contenttypes.Convertor
}

func (inst *ServiceImpl) _impl() contenttypes.Service {
	return inst
}

func (inst *ServiceImpl) makeURN(o1 *dto.ContentType) dxo.ContentTypeURN {
	name := o1.TypeName.String()
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	urn := "urn:contenttype:" + name
	return dxo.ContentTypeURN(urn)
}

// Insert ...
func (inst *ServiceImpl) Insert(ctx context.Context, o1 *dto.ContentType) (*dto.ContentType, error) {

	o1.URN = inst.makeURN(o1)

	o2 := inst.convertor.ConvertD2E(o1)
	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}
	o4 := inst.convertor.ConvertE2D(o3)
	return o4, nil
}

// Update ...
func (inst *ServiceImpl) Update(ctx context.Context, id dxo.ContentTypeID, o1 *dto.ContentType) (*dto.ContentType, error) {

	ent, err := inst.Dao.Update(nil, id, func(ct *entity.ContentType) {

		// ct.Name = o1.TypeName

		ct.Priority = o1.Priority
		ct.AsDir = o1.AsDir
		ct.AsFile = o1.AsFile
		ct.AsProject = o1.AsProject

	})
	if err != nil {
		return nil, err
	}

	o2 := inst.convertor.ConvertE2D(ent)
	return o2, nil
}

// Remove ...
func (inst *ServiceImpl) Remove(ctx context.Context, id dxo.ContentTypeID) error {
	return inst.Dao.Remove(nil, id)
}

// Find ...
func (inst *ServiceImpl) Find(ctx context.Context, id dxo.ContentTypeID) (*dto.ContentType, error) {
	item1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}
	item2 := inst.convertor.ConvertE2D(item1)
	return item2, nil
}

// List ...
func (inst *ServiceImpl) List(ctx context.Context, q *contenttypes.Query) ([]*dto.ContentType, error) {
	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}
	list2 := inst.convertor.ConvertListE2D(list1)
	return list2, nil
}
