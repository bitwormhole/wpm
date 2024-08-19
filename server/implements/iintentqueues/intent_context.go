package iintentqueues

import (
	"fmt"
	"sync"
	"time"

	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/starter-go/base/lang"
)

type iContext struct {
	selist []*iSession
	eplist []*iQueueEP
	mutex  sync.Mutex
}

func (inst *iContext) getSession(home string, create bool) (*iSession, error) {

	inst.mutex.Lock()
	defer func() {
		inst.mutex.Unlock()
	}()

	list := inst.selist
	for _, item := range list {
		if item.home == home {
			return item, nil
		}
	}

	if !create {
		return nil, fmt.Errorf("no intent-session with home path: %s", home)
	}

	session := new(iSession)
	session.context = inst
	session.home = home

	list = append(list, session)
	inst.selist = list
	return session, nil
}

func (inst *iContext) nextQueueID() dxo.IntentQueueID {
	inst.mutex.Lock()
	defer func() {
		inst.mutex.Unlock()
	}()
	time.Sleep(time.Millisecond * 3)
	now := lang.Now()
	n := now.Int()
	n = n ^ 0xfffffffffffffff
	return dxo.IntentQueueID(n)
}

func (inst *iContext) putEndpoint(ep *iQueueEP) {

	inst.mutex.Lock()
	defer func() {
		inst.mutex.Unlock()
	}()

	if ep == nil {
		return
	}
	inst.eplist = append(inst.eplist, ep)
}

func (inst *iContext) findEndpoint(id dxo.IntentQueueID) (*iQueueEP, error) {

	inst.mutex.Lock()
	defer func() {
		inst.mutex.Unlock()
	}()

	all := inst.eplist
	for _, item := range all {
		if item.id == id {
			return item, nil
		}
	}
	return nil, fmt.Errorf("no endpoint with id: %d", id)
}

func (inst *iContext) findEndpointByHome(home string) (*iQueueEP, error) {

	inst.mutex.Lock()
	defer func() {
		inst.mutex.Unlock()
	}()

	all := inst.eplist
	for _, item := range all {
		if item.session.home == home {
			return item, nil
		}
	}
	return nil, fmt.Errorf("no endpoint with home path: %s", home)
}

func (inst *iContext) closeEndpoint(id dxo.IntentQueueID) {

	inst.mutex.Lock()
	defer func() {
		inst.mutex.Unlock()
	}()

	list1 := inst.eplist
	list2 := make([]*iQueueEP, 0)
	for _, item := range list1 {
		if item.id == id {
			// close this ep
			item.close()
		} else {
			list2 = append(list2, item)
		}
	}
	inst.eplist = list2
}
