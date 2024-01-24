package impexecutables

import (
	"github.com/bitwormhole/wpm/common/classes/caches"
)

type myStateCache struct {
	class caches.Class
}

func (inst *myStateCache) init(ser caches.Service, ldr caches.Loader) {
	inst.class = ser.GetPool().NewClass("dto.Executable", ldr)
}

func (inst *myStateCache) load(want *caches.Want) (any, error) {
	want.ClassName = inst.class.Name()
	want.Class = inst.class
	item := inst.class.GetItem(want.ID)
	return item.GetData(want)
}
