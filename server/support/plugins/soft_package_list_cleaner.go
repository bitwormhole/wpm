package plugins

import (
	"context"

	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type myPakcageListCleaner struct {
	namespaceSer  service.NamespaceService
	packageSer    service.SoftwarePackageService
	httpclientSer service.HTTPClientExService
}

func (inst *myPakcageListCleaner) Clean(c context.Context) error {

	all, err := inst.packageSer.ListAll(c)
	if err != nil {
		return err
	}

	// list2rm := make([]*dto.SoftwarePackage, 0)
	// list2keep := make([]*dto.SoftwarePackage, 0)
	// for _, item := range all {
	// if item.Installed {
	// 	list2keep = append(list2keep, item)
	// } else {
	// 	list2rm = append(list2rm, item)
	// }
	// }

	inst.removePacks(c, all)
	return nil
}

func (inst *myPakcageListCleaner) removePacks(c context.Context, list []*dto.SoftwarePackage) {
	for _, item := range list {
		err := inst.removePack(c, item)
		if err != nil {
			vlog.Warn(err)
		}
	}
}

func (inst *myPakcageListCleaner) removePack(c context.Context, p *dto.SoftwarePackage) error {
	if p.Installed {
		return nil // skip if installed
	}
	return inst.packageSer.Remove(c, p.ID)
}
