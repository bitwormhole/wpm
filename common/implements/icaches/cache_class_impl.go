package icaches

import (
	"crypto/sha1"
	"time"

	"github.com/bitwormhole/wpm/common/classes/caches"
	"github.com/starter-go/base/lang"
)

type classImpl struct {
	context *cacheContext
	name    string
	maxAge  time.Duration
	loader  caches.Loader
	uuid    lang.UUID
}

func (inst *classImpl) _impl() caches.Class {
	return inst
}

func (inst *classImpl) init() {
	name := inst.name
	sum := sha1.Sum([]byte(name))
	uuid := lang.NewUUID(sum[:])
	inst.uuid = uuid
}

func (inst *classImpl) UUID() lang.UUID {
	return inst.uuid
}

func (inst *classImpl) Name() string {
	return inst.name
}

func (inst *classImpl) Loader() caches.Loader {
	return inst.loader
}

func (inst *classImpl) MaxAge() time.Duration {
	return inst.maxAge
}

func (inst *classImpl) GetItem(id string) caches.Item {

	ctx := inst.context
	ctx.mutex.Lock()
	defer func() {
		ctx.mutex.Unlock()
	}()

	key := ctx.keyForItem(inst, id)
	old := ctx.items[key]
	if old != nil {
		return old
	}

	now := lang.Now()
	item2 := &itemImpl{
		context: ctx,
		class:   inst,
		id:      id,
		t1:      now,
	}

	return item2
}
