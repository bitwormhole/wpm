package irepositories

import (
	"context"

	"github.com/bitwormhole/wpm-api/v1/dto"
	"github.com/bitwormhole/wpm/app/classes/locations"
	"github.com/bitwormhole/wpm/app/classes/repositories"
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/starter-go/vlog"
)

// LocalRepositoryServiceImpl ...
type LocalRepositoryServiceImpl struct {

	//starter:component

	_as func(repositories.LocalRepositoryService) //starter:as("#")

	LocationService locations.Service //starter:inject("#")

}

func (inst *LocalRepositoryServiceImpl) _impl() repositories.LocalRepositoryService {
	return inst
}

// Find ...
func (inst *LocalRepositoryServiceImpl) Find(c context.Context, id dxo.LocalRepositoryID) (*dto.LocalRepository, error) {

	o1, err := inst.LocationService.Find(c, dxo.LocationID(id))
	if err != nil {
		return nil, err
	}
	cvt := &repositories.Convertor{}
	return cvt.ConvertL2D(o1)
}

// List ...
func (inst *LocalRepositoryServiceImpl) List(c context.Context, q *repositories.Query) ([]*dto.LocalRepository, error) {
	q2 := &locations.Query{
		All:        q.All,
		Pagination: q.Pagination,
	}
	cate := dxo.CategoryGitRepository
	ser := inst.LocationService
	list1, err := ser.ListWithCategory(c, cate, q2)
	if err != nil {
		return nil, err
	}
	list2 := make([]*dto.LocalRepository, 0)
	cvt := &repositories.Convertor{}
	for _, item1 := range list1 {
		item2, err := cvt.ConvertL2D(item1)
		if err != nil {
			vlog.Warn(err.Error())
			continue
		}
		list2 = append(list2, item2)
	}
	return list2, nil
}
