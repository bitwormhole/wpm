package icaches

import (
	"fmt"
	"time"

	"github.com/bitwormhole/wpm/common/classes/caches"
	"github.com/starter-go/base/lang"
)

type itemDataHolder struct {
	t2   lang.Time // updated_at
	data any
}

type itemImpl struct {
	context *cacheContext
	id      string
	class   caches.Class
	t1      lang.Time // created_at
	holder  *itemDataHolder
}

func (inst *itemImpl) _impl() caches.Item {
	return inst
}

func (inst *itemImpl) ID() string {
	return inst.id
}

func (inst *itemImpl) Class() caches.Class {
	return inst.class
}

func (inst *itemImpl) CreatedAt() lang.Time {
	return inst.t1
}

func (inst *itemImpl) UpdatedAt() lang.Time {
	h := inst.holder
	if h == nil {
		return 0
	}
	return h.t2
}

// 把这个 item 注册到 context
func (inst *itemImpl) putSelf() {

	ctx := inst.context
	ctx.mutex.Lock()
	defer func() {
		ctx.mutex.Unlock()
	}()

	key := ctx.keyForItem(inst.class, inst.id)
	ctx.items[key] = inst
}

// 更新缓存数据
func (inst *itemImpl) Update() {
	inst.holder = nil
}

func (inst *itemImpl) GetData(want *caches.Want) (any, error) {

	h := inst.holder
	if inst.isExpired(h, want) {
		h = nil
	}

	if h == nil {
		h2, err := inst.load(want)
		if err != nil {
			return nil, err
		}
		h = h2
		inst.holder = h2
		inst.putSelf()
	}

	return h.data, nil
}

func (inst *itemImpl) isExpired(h *itemDataHolder, want *caches.Want) bool {

	if h == nil {
		return true
	}

	if h.data == nil {
		return true
	}

	if h.t2 < want.NotBefore {
		return true
	}

	now := lang.Now()
	age := time.Duration(now-h.t2) * time.Millisecond
	max := inst.class.MaxAge()
	if max < age {
		return true
	}

	return false
}

func (inst *itemImpl) load(want *caches.Want) (*itemDataHolder, error) {

	if want == nil {
		return nil, fmt.Errorf("no param: want")
	}
	fn := want.OnLoad
	if fn == nil {
		return nil, fmt.Errorf("no param: OnLoad")
	}

	data, err := fn(want)
	if err != nil {
		return nil, err
	}

	h := new(itemDataHolder)
	h.data = data
	h.t2 = lang.Now()
	return h, nil
}
