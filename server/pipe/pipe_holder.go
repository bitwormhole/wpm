package pipe

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type Holder struct {
	deletedAt util.Time

	Info    dto.PipeInfo
	Packets []*dto.PipePacket

	Listener chan int
}
