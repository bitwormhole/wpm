package implservice

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils/filequery"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// FileQueryServiceImpl ...
type FileQueryServiceImpl struct {
	markup.Component `id:"FileQueryService"`

	HandlerRegistryList []filequery.HandlerRegistry `inject:".filequery-handler-registry"`

	cachedHandlers []filequery.Handler
}

func (inst *FileQueryServiceImpl) _Impl() service.FileQueryService {
	return inst
}

func (inst *FileQueryServiceImpl) loadHandlers() []filequery.Handler {
	src := inst.HandlerRegistryList
	dst := make([]filequery.Handler, 0)
	for _, r1 := range src {
		list := r1.ListRegistrations()
		for _, r2 := range list {
			h := r2.Handler
			if r2.Enabled && h != nil {
				dst = append(dst, h)
			}
		}
	}
	return dst
}

func (inst *FileQueryServiceImpl) getHandlers() []filequery.Handler {
	dst := inst.cachedHandlers
	if dst == nil {
		dst = inst.loadHandlers()
		inst.cachedHandlers = dst
	}
	return dst
}

func (inst *FileQueryServiceImpl) findHandler(q *vo.FileQuery) (filequery.Handler, error) {
	all := inst.getHandlers()
	for _, h := range all {
		if h.Accept(q) {
			return h, nil
		}
	}
	return nil, fmt.Errorf("no wanted handler for the file-query")
}

// Query ...
func (inst *FileQueryServiceImpl) Query(ctx context.Context, q *vo.FileQuery) (*vo.FileQuery, error) {
	h, err := inst.findHandler(q)
	if err != nil {
		return nil, err
	}
	return h.Query(q)
}
