package plugins

import (
	"context"
	"fmt"
	"sort"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
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

// ListAll ...
func (inst *ImpSoftwareSetService) ListAll(ctx context.Context) ([]*dto.SoftwareSet, error) {

	src, err := inst.SoftwarePackageService.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	b := &mySoftwareSetBuilder{}
	return b.Create(src)
}

func (inst *ImpSoftwareSetService) Install(ctx context.Context, ss *dto.SoftwareSet) error {
	return fmt.Errorf("no impl")
}

func (inst *ImpSoftwareSetService) Uninstall(ctx context.Context, ss *dto.SoftwareSet) error {
	return fmt.Errorf("no impl")
}

func (inst *ImpSoftwareSetService) ReInstall(ctx context.Context, ss *dto.SoftwareSet) error {
	return fmt.Errorf("no impl")
}

func (inst *ImpSoftwareSetService) Upgrade(ctx context.Context, ss *dto.SoftwareSet) error {
	return fmt.Errorf("no impl")
}

////////////////////////////////////////////////////////////////////////////////

type mySoftwareSetBuilder struct {
	packs []*dto.SoftwarePackage
	sets  map[string]*dto.SoftwareSet
}

func (inst *mySoftwareSetBuilder) Create(src []*dto.SoftwarePackage) ([]*dto.SoftwareSet, error) {
	inst.packs = src
	inst.sets = make(map[string]*dto.SoftwareSet)
	sort.Sort(inst)
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
		item.HasNewVersion = inst.computeHasNewVersion(item)
	}
	return dst, nil
}

func (inst *mySoftwareSetBuilder) computeHasNewVersion(ss *dto.SoftwareSet) bool {
	if ss == nil {
		return false
	}
	packs := ss.Packages
	if packs == nil {
		return false
	}
	cnt := len(packs)
	if cnt < 1 {
		return false
	}
	last := packs[cnt-1]
	if last == nil {
		return false
	}
	return !last.Installed
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

func (inst *mySoftwareSetBuilder) Len() int {
	return len(inst.packs)
}
func (inst *mySoftwareSetBuilder) Less(i1, i2 int) bool {
	o1 := inst.packs[i1]
	o2 := inst.packs[i2]
	return o1.Revision < o2.Revision
}
func (inst *mySoftwareSetBuilder) Swap(i1, i2 int) {
	o1 := inst.packs[i1]
	o2 := inst.packs[i2]
	inst.packs[i1] = o2
	inst.packs[i2] = o1
}
