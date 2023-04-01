package intents

import (
	"context"
	"sort"

	"github.com/bitwormhole/wpm/server/web/dto"
)

// FilterChainBuilder ...
type FilterChainBuilder struct {
	items []*FilterRegistration
}

// AddRegistry ...
func (inst *FilterChainBuilder) AddRegistry(list ...FilterRegistry) *FilterChainBuilder {
	for _, item1 := range list {
		if item1 == nil {
			return inst
		}
		list2 := item1.GetFilterRegistrationList()
		if list2 == nil {
			return inst
		}
		inst.AddRegistration(list2...)
	}
	return inst
}

// AddRegistration ...
func (inst *FilterChainBuilder) AddRegistration(list ...*FilterRegistration) *FilterChainBuilder {
	inst.items = append(inst.items, list...)
	return inst
}

// Create ...
func (inst *FilterChainBuilder) Create() FilterChain {
	inst.trim()
	sort.Sort(inst)
	src := inst.items
	chain0 := &filterChainEnding{}
	chain := FilterChain(chain0)
	for _, item := range src {
		node := &filterChainNode{
			next:     chain,
			myFilter: item.Filter,
		}
		chain = node
	}
	return chain
}

// 移除无效的项
func (inst *FilterChainBuilder) trim() {
	src := inst.items
	dst := make([]*FilterRegistration, 0)
	for _, item := range src {
		if inst.isok(item) {
			dst = append(dst, item)
		}
	}
	inst.items = dst
}

func (inst *FilterChainBuilder) isok(fr *FilterRegistration) bool {

	if fr == nil {
		return false
	}

	if !fr.Enabled {
		return false
	}

	if fr.Filter == nil {
		return false
	}

	return true
}

func (inst *FilterChainBuilder) Len() int {
	return len(inst.items)
}

func (inst *FilterChainBuilder) Less(i1, i2 int) bool {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	return o1.Order < o2.Order
}

func (inst *FilterChainBuilder) Swap(i1, i2 int) {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	inst.items[i1] = o2
	inst.items[i2] = o1
}

////////////////////////////////////////////////////////////////////////////////

type filterChainNode struct {
	next     FilterChain
	myFilter Filter
}

func (inst *filterChainNode) _Impl() FilterChain {
	return inst
}

func (inst *filterChainNode) Handle(c context.Context, i *dto.Intent) error {
	return inst.myFilter.Handle(c, i, inst.next)
}

////////////////////////////////////////////////////////////////////////////////

type filterChainEnding struct{}

func (inst *filterChainEnding) _Impl() FilterChain {
	return inst
}

func (inst *filterChainEnding) Handle(c context.Context, i *dto.Intent) error {
	return nil // nop for ending
}

////////////////////////////////////////////////////////////////////////////////
