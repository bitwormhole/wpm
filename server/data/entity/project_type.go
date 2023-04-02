package entity

import (
	"sort"

	"github.com/bitwormhole/wpm/server/data/dxo"
)

// ContentType ...
type ContentType struct {
	ID  dxo.ContentTypeID  `gorm:"primaryKey"`
	URN dxo.ContentTypeURN `gorm:"index:,unique"`

	Base

	Name     dxo.ContentTypeName
	Patterns dxo.StringList

	Label       string
	Icon        string
	Description string

	Priority  int // 优先级，数值越高越先处理
	AsFile    bool
	AsDir     bool
	AsProject bool // true 表示这是一个项目类型
}

////////////////////////////////////////////////////////////////////////////////

// ContentTypeLessFn ...
type ContentTypeLessFn func(o1, o2 *ContentType) bool

type myContentTypeSorter struct {
	list []*ContentType
	fn   ContentTypeLessFn
}

func (inst *myContentTypeSorter) Len() int {
	return len(inst.list)
}

func (inst *myContentTypeSorter) Less(i1, i2 int) bool {
	o1 := inst.list[i1]
	o2 := inst.list[i2]
	if o1 == nil || o2 == nil {
		return false
	}
	return inst.fn(o1, o2)
}

func (inst *myContentTypeSorter) Swap(i1, i2 int) {
	o1 := inst.list[i1]
	o2 := inst.list[i2]
	inst.list[i1] = o2
	inst.list[i2] = o1
}

// SortContentTypeArray 对 ContentType 数组进行排序
func SortContentTypeArray(list []*ContentType, fn ContentTypeLessFn) {
	if list == nil || fn == nil {
		return
	}
	sorter := &myContentTypeSorter{
		list: list,
		fn:   fn,
	}
	sort.Sort(sorter)
}

////////////////////////////////////////////////////////////////////////////////
