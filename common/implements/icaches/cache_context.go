package icaches

import (
	"sync"

	"github.com/bitwormhole/wpm/common/classes/caches"
)

type myItemKey string

type cacheContext struct {
	service caches.Service
	pool    caches.Pool

	mutex   sync.Mutex
	classes map[string]caches.Class
	items   map[myItemKey]caches.Item
}

func (inst *cacheContext) keyForItem(class caches.Class, itemID string) myItemKey {
	s1 := class.UUID().String()
	s2 := itemID
	str := s1 + "/" + s2
	return myItemKey(str)
}
