package ilocations

import (
	"context"

	"github.com/bitwormhole/wpm/app/classes/locations"
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/bitwormhole/wpm/app/data/entity"
	"github.com/bitwormhole/wpm/app/web/dto"
)

// LocationServiceImpl ...
type LocationServiceImpl struct {

	//starter:component

	_as func(locations.Service) //starter:as("#")

	LocationDAO locations.LocationDAO //starter:inject("#")
}

func (inst *LocationServiceImpl) _impl() locations.Service {
	return inst
}

// Find ...
func (inst *LocationServiceImpl) Find(c context.Context, id dxo.LocationID) (*dto.Location, error) {
	o1, err := inst.LocationDAO.Find(nil, id)
	if err != nil {
		return nil, err
	}
	o2 := locations.ConvertE2D(o1)
	return o2, nil
}

// ListWithCategory ...
func (inst *LocationServiceImpl) ListWithCategory(c context.Context, category dxo.Category, q *locations.Query) ([]*dto.Location, error) {
	list1, err := inst.LocationDAO.ListWithCategory(nil, category, q)
	if err != nil {
		return nil, err
	}
	list2 := locations.ConvertListE2D(list1)
	return list2, nil
}

// Insert ....
func (inst *LocationServiceImpl) Insert(c context.Context, item *dto.Location) (*dto.Location, error) {
	o2 := locations.ConvertD2E(item)
	o3, err := inst.LocationDAO.Insert(nil, o2)
	if err != nil {
		return nil, err
	}
	o4 := locations.ConvertE2D(o3)
	return o4, nil
}

// Update ...
func (inst *LocationServiceImpl) Update(c context.Context, id dxo.LocationID, item *dto.Location) (*dto.Location, error) {

	o2, err := inst.LocationDAO.Update(nil, id, func(ent *entity.Location) {

		ent.Label = item.Label
		ent.Description = item.Description

	})
	if err != nil {
		return nil, err
	}
	o3 := locations.ConvertE2D(o2)
	return o3, nil
}

// Remove ...
func (inst *LocationServiceImpl) Remove(c context.Context, id dxo.LocationID) error {
	return inst.LocationDAO.Remove(nil, id)
}
