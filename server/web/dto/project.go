package dto

import (
	"sort"

	"github.com/bitwormhole/wpm/server/data/dxo"
)

// Project ...
type Project struct {
	ID dxo.ProjectID `json:"id"`
	Base

	Name        string              `json:"name"`
	Description string              `json:"description"`
	Type        dxo.ContentTypeName `json:"type"`

	ConfigFileName string `json:"config_file_name"`
	PathInWorktree string `json:"path_in_worktree"`
	Path           string `json:"path"`
	ProjectDir     string `json:"project_dir"`

	Location dxo.LocationID    `json:"location"`
	Class    dxo.LocationClass `json:"location_class"`

	OwnerRepository dxo.LocalRepositoryID `json:"owner_repository"`
	Group           dxo.ProjectGroupName  `json:"group"`
	State           dxo.FileState         `json:"state"`
	IsDir           bool                  `json:"is_dir"`
	IsFile          bool                  `json:"is_file"`
	Tags            dxo.StringList        `json:"tags"`
}

////////////////////////////////////////////////////////////////////////////////

// ProjectLessFunc ...
type ProjectLessFunc func(o1, o2 *Project) bool

// ProjectSorter ...
type ProjectSorter struct {
	items []*Project
	cmp   ProjectLessFunc
}

// Sort ...
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

////////////////////////////////////////////////////////////////////////////////
