package implservice

import (
	"context"
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils/filequery"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// FileQueryServiceImpl ...
type FileQueryServiceImpl struct {
	markup.Component `id:"FileQueryService"`

	ProfileService      service.ProfileService      `inject:"#ProfileService"`
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

	err := inst.preprocessQuery(ctx, q)
	if err != nil {
		return nil, err
	}

	h, err := inst.findHandler(q)
	if err != nil {
		return nil, err
	}

	return h.Query(q)
}

func (inst *FileQueryServiceImpl) preprocessQuery(ctx context.Context, q *vo.FileQuery) error {
	path := q.Path
	if strings.HasPrefix(path, "~") {
		p1 := inst.getUserHomeDir()
		p2 := path[1:]
		q.Path = p1 + "/" + p2
	}
	return nil
}

func (inst *FileQueryServiceImpl) getUserHomeDir() string {
	p, err := inst.ProfileService.GetProfile()
	if err == nil && p != nil {
		return p.Home
	}
	return "/"
}
