package isoftware

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/packages"
)

// SoftwarePackageService ...
type SoftwarePackageService struct {

	//starter:component
	_as func(packages.Service) //starter:as("#")

	Dao packages.DAO //starter:inject("#")

	convertor packages.Convertor
}

func (inst *SoftwarePackageService) _impl() packages.Service {
	return inst
}

// Insert ...
func (inst *SoftwarePackageService) Insert(ctx context.Context, o *dto.SoftwarePackage) (*dto.SoftwarePackage, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *SoftwarePackageService) Update(ctx context.Context, id dxo.SoftwarePackageID, o *dto.SoftwarePackage) (*dto.SoftwarePackage, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *SoftwarePackageService) Remove(ctx context.Context, id dxo.SoftwarePackageID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *SoftwarePackageService) Find(ctx context.Context, id dxo.SoftwarePackageID) (*dto.SoftwarePackage, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *SoftwarePackageService) List(ctx context.Context, q *packages.Query) ([]*dto.SoftwarePackage, error) {
	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}
	list2 := inst.convertor.ConvertListE2D(list1)
	return list2, nil
}
