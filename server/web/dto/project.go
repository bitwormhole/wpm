package dto

import (
	"sort"

	"github.com/bitwormhole/wpm/server/data/dxo"
)

// Project ...
type Project struct {
	ID dxo.ProjectID `json:"id"`
	Base

	Name            string                `json:"name"`
	Description     string                `json:"description"`
	ProjectType     string                `json:"project_type"`
	Path            string                `json:"path"`
	OwnerRepository dxo.LocalRepositoryID `json:"owner_repository"`
	Group           dxo.ProjectGroupName  `json:"group"`
	State           dxo.FileState         `json:"state"`
	Tags            dxo.StringList        `json:"tags"`
}

type ProjectLessFunc func(o1, o2 *Project) bool

type ProjectSorter struct {
	items []*Project
	cmp   ProjectLessFunc
}

func (inst *ProjectSorter) Sort(items []*Project, fn ProjectLessFunc) {
	inst.cmp = fn
	inst.items = items
	sort.Sort(inst)
}

func (inst *ProjectSorter) Len() int {
	return len(inst.items)
}

func (inst *ProjectSorter) Less(i1, i2 int) bool {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	return inst.cmp(o1, o2)
}

func (inst *ProjectSorter) Swap(i1, i2 int) {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	inst.items[i1] = o2
	inst.items[i2] = o1
}
