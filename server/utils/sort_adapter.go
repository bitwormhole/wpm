package utils

import "sort"

// SortFn ... 排序函数组
type SortFn struct {
	OnLess func(i1, i2 int) bool
	OnSwap func(i1, i2 int)
	OnLen  func() int
}

func (inst *SortFn) ready() bool {
	a := inst.OnLen != nil
	b := inst.OnLess != nil
	c := inst.OnSwap != nil
	return (a && b && c)
}

func (inst *SortFn) setNOP() {
	nop := &sortNOP{}
	inst.OnLess = nop.less
	inst.OnSwap = nop.swap
	inst.OnLen = nop.len
}

////////////////////////////////////////////////////////////////////////////////

type sortNOP struct {
}

func (inst *sortNOP) less(i1, i2 int) bool {
	return false
}

func (inst *sortNOP) swap(i1, i2 int) {}

func (inst *sortNOP) len() int {
	return 0
}

////////////////////////////////////////////////////////////////////////////////

// SortAdapter ...排序适配器
type SortAdapter struct {
	SortFn
	Quietly bool // 出现异常时，不要抛出 panic
	cached  *SortFn
}

func (inst *SortAdapter) forCache() *SortFn {
	c := inst.cached
	if c != nil {
		return c
	}
	c = &SortFn{
		OnLen:  inst.OnLen,
		OnSwap: inst.OnSwap,
		OnLess: inst.OnLess,
	}
	if !c.ready() {
		if inst.Quietly {
			c.setNOP()
		} else {
			panic("the sort adapter is not ready (with 3 functions)")
		}
	}
	inst.cached = c
	return c
}

// Sort 执行排序
func (inst *SortAdapter) Sort() {
	sort.Sort(inst)
}

func (inst *SortAdapter) Less(i1, i2 int) bool {
	f := inst.forCache()
	return f.OnLess(i1, i2)
}

func (inst *SortAdapter) Swap(i1, i2 int) {
	f := inst.forCache()
	f.OnSwap(i1, i2)
}

func (inst *SortAdapter) Len() int {
	f := inst.forCache()
	return f.OnLen()
}
