package impsoftware

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/packages"
	"github.com/bitwormhole/wpm/server/classes/softwaresets"
)

// SoftwareSetService ...
type SoftwareSetService struct {

	//starter:component
	_as func(softwaresets.Service) //starter:as("#")

	PackageService packages.Service //starter:inject("#")
}

func (inst *SoftwareSetService) _impl() softwaresets.Service {
	return inst
}

// Insert ...
func (inst *SoftwareSetService) Insert(ctx context.Context, o *dto.SoftwareSet) (*dto.SoftwareSet, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *SoftwareSetService) Update(ctx context.Context, id dxo.SoftwareSetID, o *dto.SoftwareSet) (*dto.SoftwareSet, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *SoftwareSetService) Remove(ctx context.Context, id dxo.SoftwareSetID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *SoftwareSetService) Find(ctx context.Context, id dxo.SoftwareSetID) (*dto.SoftwareSet, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *SoftwareSetService) List(ctx context.Context, q *softwaresets.Query) ([]*dto.SoftwareSet, error) {
	q2 := &packages.Query{All: true}
	list1, err := inst.PackageService.List(ctx, q2)
	if err != nil {
		return nil, err
	}
	builder := new(softwareSetBuilder)
	builder.append(list1...)
	list2 := builder.create()
	return list2, nil
}
