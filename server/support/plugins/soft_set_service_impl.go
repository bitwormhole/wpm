package plugins

import (
	"context"
	"fmt"
	"sort"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ImpSoftwareSetService ...
type ImpSoftwareSetService struct {
	markup.Component `id:"SoftwareSetService"`

	SoftwarePackageService service.SoftwarePackageService `inject:"#SoftwarePackageService"`
}

func (inst *ImpSoftwareSetService) _Impl() service.SoftwareSetService {
	return inst
}

// GetOne ...
func (inst *ImpSoftwareSetService) GetOne(ctx context.Context, id dxo.SoftwarePackageID) (*dto.SoftwareSet, error) {

	ser := inst.SoftwarePackageService
	pack1, err := ser.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	src, err := ser.ListByModuleName(ctx, pack1.ModuleName)
	if err != nil {
		return nil, err
	}

	b := &mySoftwareSetBuilder{}
	sets, err := b.Create(src)
	if err != nil {
		return nil, err
	}

	if len(sets) == 1 {
		ps := sets[0]
		ps.ID = id
		return ps, nil
	}

	return nil, fmt.Errorf("the result is not only one")
}

// ListAll ...
func (inst *ImpSoftwareSetService) ListAll(ctx context.Context) ([]*dto.SoftwareSet, error) {

	src, err := inst.SoftwarePackageService.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	b := &mySoftwareSetBuilder{}
	return b.Create(src)
}

func (inst *ImpSoftwareSetService) loadPackagesForSet(ctx context.Context, ss *dto.SoftwareSet) ([]*dto.SoftwarePackage, error) {

	ser := inst.SoftwarePackageService
	all, err := ser.ListByModuleName(ctx, ss.ModuleName)
	if err != nil {
		return nil, err
	}

	// 排序
	sa := &utils.SortAdapter{}
	sa.OnLen = func() int { return len(all) }
	sa.OnSwap = func(i1, i2 int) {
		o1 := all[i1]
		o2 := all[i2]
		all[i1] = o2
		all[i2] = o1
	}
	sa.OnLess = func(i1, i2 int) bool {
		o1 := all[i1]
		o2 := all[i2]
		return o1.Revision < o2.Revision
	}
	sort.Sort(sa)

	// ss.Packages = all
	return all, nil
}

// Install ... 安装指定的软件集合
func (inst *ImpSoftwareSetService) Install(ctx context.Context, ss *dto.SoftwareSet) error {

	// find list
	list, err := inst.loadPackagesForSet(ctx, ss)
	if err != nil {
		return err
	}

	// check
	countInstalled := 0
	countNotInstalled := 0
	var wantID dxo.SoftwarePackageID = 0
	for _, p := range list {
		if p.Installed {
			countInstalled++
		} else {
			countNotInstalled++
		}
		wantID = p.ID
	}

	if countInstalled > 0 {
		return fmt.Errorf("package has been installed")
	}

	if countNotInstalled <= 0 {
		return fmt.Errorf("no package to install")
	}

	// 安装最新版本
	return inst.SoftwarePackageService.Install(ctx, wantID)
}

// Uninstall ... 卸载指定的软件集合
func (inst *ImpSoftwareSetService) Uninstall(ctx context.Context, ss *dto.SoftwareSet) error {

	list, err := inst.loadPackagesForSet(ctx, ss)
	if err != nil {
		return err
	}

	// 卸载所有已经安装的版本
	countDone := 0
	errlist := utils.ErrorList{}
	ser := inst.SoftwarePackageService
	for _, item := range list {
		if item.Installed {
			err = ser.Uninstall(ctx, item.ID)
			if err == nil {
				countDone++
			}
			errlist.Append(err)
		}
	}

	if countDone == 0 {
		return fmt.Errorf("no installed package")
	}

	return errlist.First()
}

// ReInstall ... 重新安装指定的软件集合
func (inst *ImpSoftwareSetService) ReInstall(ctx context.Context, ss *dto.SoftwareSet) error {

	list, err := inst.loadPackagesForSet(ctx, ss)
	if err != nil {
		return err
	}

	// 重新安装所有已经安装的版本
	errlist := utils.ErrorList{}
	ser := inst.SoftwarePackageService
	for _, item := range list {
		if item.Installed {
			err = ser.Uninstall(ctx, item.ID)
			errlist.Append(err)
			err = ser.Install(ctx, item.ID)
			errlist.Append(err)
		}
	}

	return errlist.First()
}

// Upgrade ... 升级指定的软件集合
func (inst *ImpSoftwareSetService) Upgrade(ctx context.Context, ss *dto.SoftwareSet) error {

	list, err := inst.loadPackagesForSet(ctx, ss)
	if err != nil {
		return err
	}

	if len(list) < 1 {
		return fmt.Errorf("no available version for the package set")
	}
	latest := list[len(list)-1]
	if latest.Installed {
		return fmt.Errorf("latest version has been installed")
	}

	// 卸载所有已经安装的 老 版本
	ser := inst.SoftwarePackageService
	for _, item := range list {
		if item.ID != latest.ID && item.Installed {
			err = ser.Uninstall(ctx, item.ID)
			if err != nil {
				return err
			}
		}
	}

	// 安装最新的版本
	return ser.Install(ctx, latest.ID)
}

////////////////////////////////////////////////////////////////////////////////

type mySoftwareSetBuilder struct {
	packs []*dto.SoftwarePackage
	sets  map[string]*dto.SoftwareSet
}

func (inst *mySoftwareSetBuilder) Create(src []*dto.SoftwarePackage) ([]*dto.SoftwareSet, error) {
	inst.packs = src
	inst.sets = make(map[string]*dto.SoftwareSet)
	inst.sortPacks()
	src = inst.packs
	for _, p := range src {
		inst.add(p)
	}
	return inst.result()
}

func (inst *mySoftwareSetBuilder) result() ([]*dto.SoftwareSet, error) {
	src := inst.sets
	dst := make([]*dto.SoftwareSet, 0)
	for _, item := range src {
		dst = append(dst, item)
		item.State = inst.computeHasNewVersion(item)
	}
	return dst, nil
}

func (inst *mySoftwareSetBuilder) computeHasNewVersion(ss *dto.SoftwareSet) dxo.SoftwarePackageState {
	if ss == nil {
		return dxo.SoftPackStateNA
	}
	var latestPack *dto.SoftwarePackage = nil
	latestRev := -1
	countInstalled := 0
	list := ss.Packages
	for _, p := range list {
		if p.Installed {
			countInstalled++
			p.State = dxo.SoftPackStateInstalled
		} else {
			p.State = dxo.SoftPackStateAvailable
		}
		if p.Revision > latestRev {
			latestPack = p
			latestRev = p.Revision
		}
	}
	if latestPack == nil {
		return dxo.SoftPackStateNA
	} else if countInstalled > 0 {
		if latestPack.Installed {
			return dxo.SoftPackStateInstalled
		}
		return dxo.SoftPackStateNewVersion
	}
	return dxo.SoftPackStateAvailable
}

func (inst *mySoftwareSetBuilder) add(sp *dto.SoftwarePackage) {
	set := inst.getOrCreateSet(sp)
	if sp.Installed {
		set.SoftwarePackage = *sp
	}
	set.Packages = append(set.Packages, sp)
}

func (inst *mySoftwareSetBuilder) getOrCreateSet(sp *dto.SoftwarePackage) *dto.SoftwareSet {
	key := sp.Namespace + "#" + sp.Name
	ss := inst.sets[key]
	if ss == nil {
		ss = &dto.SoftwareSet{}
		ss.SoftwarePackage = *sp
		inst.sets[key] = ss
	}
	return ss
}

func (inst *mySoftwareSetBuilder) sortPacks() {
	list := inst.packs
	sa := &utils.SortAdapter{}
	sa.OnLen = func() int { return len(list) }
	sa.OnLess = func(i1, i2 int) bool {
		o1 := list[i1]
		o2 := list[i2]
		return o1.Revision < o2.Revision
	}
	sa.OnSwap = func(i1, i2 int) {
		o1 := list[i1]
		o2 := list[i2]
		list[i1] = o2
		list[i2] = o1
	}
	sa.Sort()
	inst.packs = list
}
