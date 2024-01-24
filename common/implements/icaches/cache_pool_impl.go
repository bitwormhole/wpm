package icaches

import (
	"time"

	"github.com/bitwormhole/wpm/common/classes/caches"
)

type cachePool struct {
	context *cacheContext
}

func (inst *cachePool) _impl() caches.Pool {
	return inst
}

func (inst *cachePool) NewClass(name string, loader caches.Loader) caches.Class {
	ctx := inst.context
	ctx.mutex.Lock()
	defer func() {
		ctx.mutex.Unlock()
	}()

	old := ctx.classes[name]
	if old != nil {
		return old
	}

	class := new(classImpl)
	class.name = name
	class.loader = loader
	class.maxAge = 60 * time.Minute
	class.context = inst.context
	class.init()

	ctx.classes[name] = class

	return class
}

func (inst *cachePool) GetClass(name string) caches.Class {
	ctx := inst.context
	ctx.mutex.Lock()
	defer func() {
		ctx.mutex.Unlock()
	}()
	return ctx.classes[name]
}

func (inst *cachePool) Clean(want *caches.Want) {
	ctx := inst.context
	ctx.mutex.Lock()
	defer func() {
		ctx.mutex.Unlock()
	}()

	// todo ....

}
