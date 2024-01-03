package impintentqueues

import (
	"context"

	"github.com/bitwormhole/wpm/server/classes/intents"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// IntentQueueService ...
type IntentQueueService struct {

	//starter:component
	_as func(intents.Queues) //starter:as("#")

	iqContext iContext
}

func (inst *IntentQueueService) _impl() intents.Queues {
	return inst
}

// Open ...
func (inst *IntentQueueService) Open(ctx context.Context, o *dto.IntentQueue) (*dto.IntentQueue, error) {

	home := o.UserHomeDir
	session, err := inst.iqContext.getSession(home, true)
	if err != nil {
		return nil, err
	}
	ep := session.openEndpoint()

	result := new(dto.IntentQueue)
	result.UserHomeDir = home
	result.ID = ep.id
	return result, nil
}

// Listen ...
func (inst *IntentQueueService) Listen(ctx context.Context, id dxo.IntentQueueID) (*dto.IntentQueue, error) {

	ep, err := inst.iqContext.findEndpoint(id)
	if err != nil {
		return nil, err
	}

	iq, err := ep.fetch()
	if err != nil {
		return nil, err
	}

	result := new(dto.IntentQueue)
	result.Intents = iq.Queue.Intents
	return result, nil
}

// Close ...
func (inst *IntentQueueService) Close(ctx context.Context, id dxo.IntentQueueID) error {
	inst.iqContext.closeEndpoint(id)
	return nil
}

// Push ...
func (inst *IntentQueueService) Push(ctx context.Context, o *dto.IntentQueue) error {

	home := o.UserHomeDir
	ep, err := inst.iqContext.findEndpointByHome(home)
	if err != nil {
		return err
	}

	v := new(vo.IntentQueue)
	v.Queue = *o
	return ep.push(v)
}
