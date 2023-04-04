package plugins

import (
	"context"
	"runtime"

	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

type myPakcageListUpdater struct {
	namespaceSer  service.NamespaceService
	packageSer    service.SoftwarePackageService
	httpclientSer service.HTTPClientService
}

func (inst *myPakcageListUpdater) Update(c context.Context) error {
	t := &myPakcageListUpdateTask{
		updater: inst,
		context: c,
	}
	steps := make([]func() error, 0)

	steps = append(steps, t.listNamespaces)
	steps = append(steps, t.fetchNamespaces)
	steps = append(steps, t.save)

	for _, step := range steps {
		err := step()
		if err != nil {
			return err
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type myPakcageListUpdateTask struct {
	context context.Context
	updater *myPakcageListUpdater

	namespaces []*dto.Namespace
	packages   []*dto.SoftwarePackage
}

func (inst *myPakcageListUpdateTask) listNamespaces() error {
	c := inst.context
	src, err := inst.updater.namespaceSer.ListAll(c)
	if err != nil {
		return err
	}
	wantArch := runtime.GOARCH
	wantOS := runtime.GOOS
	for _, item := range src {
		if item.Arch == wantArch && item.OS == wantOS {
			inst.namespaces = append(inst.namespaces, item)
		}
	}
	return nil
}

func (inst *myPakcageListUpdateTask) fetchNamespaces() error {
	ctx := inst.context
	htc := inst.updater.httpclientSer
	nslist := inst.namespaces
	for _, item := range nslist {
		url := item.URL
		o := &vo.Online{}
		_, err := htc.FetchJSON(ctx, url, o, nil)
		if err != nil {
			// return err
			vlog.Warn(err)
			continue
		}
		packs := o.Packages
		inst.packages = append(inst.packages, packs...)
	}
	return nil
}

func (inst *myPakcageListUpdateTask) save() error {
	src := inst.packages
	ser := inst.updater.packageSer
	ctx := inst.context
	for _, item := range src {
		_, err := ser.Insert(ctx, item)
		if err != nil {
			vlog.Warn(err)
		}
	}
	return nil
}
