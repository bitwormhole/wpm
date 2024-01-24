package icaches

import "github.com/bitwormhole/wpm/common/classes/caches"

// CacheServiceImpl ...
type CacheServiceImpl struct {

	//starter:component
	_as func(caches.Service) //starter:as("#")

	context *cacheContext
}

func (inst *CacheServiceImpl) _impl() caches.Service {
	return inst
}

func (inst *CacheServiceImpl) getCtx() *cacheContext {
	ctx := inst.context
	if ctx != nil {
		return ctx
	}

	ctx = new(cacheContext)
	pool := new(cachePool)

	ctx.pool = pool
	ctx.items = make(map[myItemKey]caches.Item)
	ctx.classes = make(map[string]caches.Class)
	ctx.service = inst

	inst.context = ctx
	pool.context = ctx

	return ctx
}

// GetPool ...
func (inst *CacheServiceImpl) GetPool() caches.Pool {
	return inst.getCtx().pool
}

// func (inst *CacheServiceImpl) GetItem(want *caches.Want) (caches.Item, error) {
// 	className := want.ClassName
// 	pool := inst.GetPool()
// 	class := pool.GetClass(className)
// 	if class == nil {
// 		class = pool.
// 	}
// }
