package istatistics

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/bitwormhole/wpm/server/classes/statistics"
)

// StatisticServiceImpl ...
type StatisticServiceImpl struct {

	//starter:component
	_as func(statistics.Service) //starter:as("#")

}

func (inst *StatisticServiceImpl) _impl() statistics.Service {
	return inst
}

// GetResult ...
func (inst *StatisticServiceImpl) GetResult(ctx context.Context, result *vo.Statistic) error {

	return nil
}
