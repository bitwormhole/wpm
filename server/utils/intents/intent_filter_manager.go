package intents

import (
	"context"
	"sort"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// FilterManager ... [inject:"#intent-filter-manager"]
type FilterManager interface {
	Filter
}

////////////////////////////////////////////////////////////////////////////////

// FilterManagerImpl ...
type FilterManagerImpl struct {
	markup.Component `id:"intent-filter-manager"`

	Filters []FilterRegistry `inject:".intent-filter-registry"`

	cached  []*FilterRegistration
	loading []*FilterRegistration
}

func (inst *FilterManagerImpl) _Impl() FilterManager {
	return inst
}

// Filter ...
func (inst *FilterManagerImpl) Filter(c context.Context, intent *dto.Intent) (*dto.Intent, error) {
	p := intent
	all := inst.getFilterList()
	for _, reg := range all {
		p2, err := inst.doFilter(c, reg, p)
		if err != nil {
			return nil, err
		} else if p2 == nil {
			return nil, nil
		}
		p = p2
	}
	return p, nil
}

func (inst *FilterManagerImpl) doFilter(c context.Context, r *FilterRegistration, i *dto.Intent) (*dto.Intent, error) {
	if r == nil || i == nil {
		return i, nil
	}
	f := r.Filter
	if f == nil {
		return i, nil
	}
	return f.Filter(c, i)
}

func (inst *FilterManagerImpl) getFilterList() []*FilterRegistration {
	list := inst.cached
	if list == nil {
		list = inst.loadFilterList()
		inst.cached = list
	}
	return list
}

func (inst *FilterManagerImpl) loadFilterList() []*FilterRegistration {

	src := inst.Filters
	dst := make([]*FilterRegistration, 0)

	for _, r1 := range src {
		rlist := r1.GetRegistrationList()
		for _, r2 := range rlist {
			if r2.Enabled && (r2.Filter != nil) {
				dst = append(dst, r2)
			}
		}
	}

	inst.loading = dst
	sort.Sort(inst)
	inst.loading = nil
	return dst
}

func (inst *FilterManagerImpl) Len() int {
	return len(inst.loading)
}
func (inst *FilterManagerImpl) Less(i1, i2 int) bool {
	r1 := inst.loading[i1]
	r2 := inst.loading[i2]
	return r1.Order < r2.Order
}
func (inst *FilterManagerImpl) Swap(i1, i2 int) {
	r1 := inst.loading[i1]
	r2 := inst.loading[i2]
	inst.loading[i1] = r2
	inst.loading[i2] = r1
}
