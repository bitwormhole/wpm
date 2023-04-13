package plugins

import (
	"context"
	"runtime"

	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type myPakcageListUpdater struct {
	namespaceSer  service.NamespaceService
	packageSer    service.SoftwarePackageService
	httpclientSer service.HTTPClientExService
}

func (inst *myPakcageListUpdater) Update(c context.Context) error {
	t := &myPakcageListUpdateTask{
		updater: inst,
		context: c,
	}
	steps := make([]func() error, 0)

	steps = append(steps, t.listSources)
	steps = append(steps, t.fetchSources)
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

func (inst *myPakcageListUpdateTask) listSources() error {
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

func (inst *myPakcageListUpdateTask) fetchSources() error {
	ctx := inst.context
	htc := inst.updater.httpclientSer
	nslist := inst.namespaces
	for _, item := range nslist {
		url := item.URL
		// o := &vo.Online{}
		// _, err := htc.FetchJSON(ctx, url, o, nil)
		o, err := htc.FetchOnlineDoc(ctx, url, nil)
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
