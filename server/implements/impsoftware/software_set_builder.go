package impsoftware

import (
	"sort"
	"strings"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/starter-go/base/util"
)

// softwaresets 是由若干个 packages 构成的虚拟对象，不具有对应的实体结构

type softwareSetKey string

// softwareSetBuilder ...
type softwareSetBuilder struct {
	all map[softwareSetKey]*softwareSetHolder
}

func (inst *softwareSetBuilder) append(plist ...*dto.SoftwarePackage) {
	for _, pack := range plist {
		holder := inst.getHolderFor(pack)
		holder.add(pack)
	}
}

func (inst *softwareSetBuilder) create() []*dto.SoftwareSet {

	src := inst.all
	dst := make([]*dto.SoftwareSet, 0)

	for _, holder := range src {
		ss := holder.makeSet()
		dst = append(dst, ss)
	}

	s := new(util.Sorter)
	s.OnLen = func() int { return len(dst) }
	s.OnLess = func(i1, i2 int) bool {
		v1 := dst[i1].Name
		v2 := dst[i2].Name
		n := strings.Compare(v1, v2)
		return n > 0
	}
	s.OnSwap = func(i1, i2 int) { dst[i1], dst[i2] = dst[i2], dst[i1] }
	s.Sort()

	return dst
}

func (inst *softwareSetBuilder) keyOf(pack *dto.SoftwarePackage) softwareSetKey {
	ns := pack.Namespace
	name := pack.ModuleName
	str := ns + "###" + name
	return softwareSetKey(str)
}

func (inst *softwareSetBuilder) getHolderFor(pack *dto.SoftwarePackage) *softwareSetHolder {
	all := inst.all
	if all == nil {
		all = make(map[softwareSetKey]*softwareSetHolder)
		inst.all = all
	}
	key := inst.keyOf(pack)
	h := all[key]
	if h == nil {
		h = new(softwareSetHolder)
		h.key = key
		all[key] = h
	}
	return h
}

////////////////////////////////////////////////////////////////////////////////

// softwareSetHolder ...
type softwareSetHolder struct {
	key softwareSetKey

	packs []*dto.SoftwarePackage
}

func (inst *softwareSetHolder) add(pack *dto.SoftwarePackage) {
	inst.packs = append(inst.packs, pack)
}

func (inst *softwareSetHolder) Len() int {
	return len(inst.packs)
}
func (inst *softwareSetHolder) Less(i1, i2 int) bool {
	v1 := inst.packs[i1].Revision
	v2 := inst.packs[i2].Revision
	return v1 < v2
}
func (inst *softwareSetHolder) Swap(i1, i2 int) {
	l := inst.packs
	l[i1], l[i2] = l[i2], l[i1]
}

func (inst *softwareSetHolder) makeSet() *dto.SoftwareSet {
	sort.Sort(inst)
	src := inst.packs
	dst := new(dto.SoftwareSet)
	for _, p := range src {
		dst.SoftwarePackage = *p
		break
	}
	dst.Packages = src
	return dst
}
