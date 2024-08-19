package iintentqueues

import (
	"encoding/json"

	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/common/objects/vo"
)

type iQueueEP struct {
	session *iSession
	id      dxo.IntentQueueID
	chl     chan string
}

func (inst *iQueueEP) open() {
	const size = 32
	if inst.chl == nil {
		c := make(chan string, size)
		inst.chl = c
	}
}

func (inst *iQueueEP) close() {
	c := inst.chl
	inst.chl = nil
	if c == nil {
		return
	}
	close(c)
}

func (inst *iQueueEP) push(obj *vo.IntentQueue) error {
	if obj == nil {
		return nil
	}
	bin, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	text := string(bin)
	inst.chl <- text
	return nil
}

func (inst *iQueueEP) fetch() (*vo.IntentQueue, error) {
	obj := new(vo.IntentQueue)
	text := <-inst.chl
	bin := []byte(text)
	err := json.Unmarshal(bin, obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
