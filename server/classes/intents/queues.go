package intents

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
)

// Queues 表示该类型的服务接口
type Queues interface {
	Open(ctx context.Context, o *dto.IntentQueue) (*dto.IntentQueue, error)
	Listen(ctx context.Context, id dxo.IntentQueueID) (*dto.IntentQueue, error)
	Close(ctx context.Context, id dxo.IntentQueueID) error
	Push(ctx context.Context, o *dto.IntentQueue) error
}
