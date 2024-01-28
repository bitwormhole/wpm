package caches

import (
	"time"

	"github.com/starter-go/base/lang"
)

// Loader ...
type Loader interface {
	Pool(pool Pool) Loader

	Class(class string) Loader

	ID(id string) Loader

	OnLoad(fn OnLoadFunc) Loader

	MaxAge(ma time.Duration) Loader

	NotBefore(t lang.Time) Loader

	Load() (any, error)
}

////////////////////////////////////////////////////////////////////////////////

type myLoader struct {
	// agent *Agent
	want Want
}

func (inst *myLoader) _impl() Loader {
	return inst
}

func (inst *myLoader) Pool(pool Pool) Loader {
	inst.want.Pool = pool
	return inst
}

func (inst *myLoader) Class(class string) Loader {
	inst.want.Class = class
	return inst
}

func (inst *myLoader) ID(id string) Loader {
	inst.want.ID = id
	return inst
}

func (inst *myLoader) OnLoad(fn OnLoadFunc) Loader {
	inst.want.OnLoad = fn
	return inst
}

func (inst *myLoader) MaxAge(ma time.Duration) Loader {
	ms := lang.Time(ma / time.Millisecond)
	now := lang.Now()
	inst.want.NotBefore = now - ms
	return inst
}

func (inst *myLoader) NotBefore(ts lang.Time) Loader {
	inst.want.NotBefore = ts
	return inst
}

func (inst *myLoader) Load() (any, error) {
	want := &inst.want
	item := want.Item
	if item == nil {
		id := want.ID
		item = inst.getClass().GetItem(id)
		want.Item = item
	}
	return item.GetData(want)
}

func (inst *myLoader) getClass() Class {
	pool := inst.want.Pool
	if pool == nil {
		panic("want.pool is nil")
	}
	cname := inst.want.Class
	return pool.GetClass(cname)
}
