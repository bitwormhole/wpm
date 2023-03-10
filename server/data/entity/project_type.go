package entity

import (
	"sort"

	"github.com/bitwormhole/wpm/server/data/dxo"
)

// ProjectType ...
type ProjectType struct {
	ID dxo.ProjectTypeID `gorm:"primaryKey"`
	Base

	Name        string `gorm:"index:,unique"`
	Type        string
	Label       string
	Description string

	Priority int // 优先级，数值越高越先处理
	AsFile   bool
	AsDir    bool
}

////////////////////////////////////////////////////////////////////////////////

// ProjectTypeLessFn ...
type ProjectTypeLessFn func(o1, o2 *ProjectType) bool

type myProjectTypeSorter struct {
	list []*ProjectType
	fn   ProjectTypeLessFn
}

func (inst *myProjectTypeSorter) Len() int {
	return len(inst.list)
}

func (inst *myProjectTypeSorter) Less(i1, i2 int) bool {
	o1 := inst.list[i1]
	o2 := inst.list[i2]
	if o1 == nil || o2 == nil {
		return false
	}
	return inst.fn(o1, o2)
}

func (inst *myProjectTypeSorter) Swap(i1, i2 int) {
	o1 := inst.list[i1]
	o2 := inst.list[i2]
	inst.list[i1] = o2
	inst.list[i2] = o1
}

// SortProjectTypeArray 对 ProjectType 数组进行排序
func SortProjectTypeArray(list []*ProjectType, fn ProjectTypeLessFn) {
	if list == nil || fn == nil {
		return
	}
	sorter := &myProjectTypeSorter{
		list: list,
		fn:   fn,
	}
	sort.Sort(sorter)
}

////////////////////////////////////////////////////////////////////////////////
