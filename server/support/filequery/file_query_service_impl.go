package filequery

import (
	"context"
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils/filequery"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// FileQueryServiceImpl ...
type FileQueryServiceImpl struct {
	markup.Component `id:"FileQueryService"`

	ContentTypeService  service.ContentTypeService  `inject:"#ContentTypeService"`
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
func (inst *FileQueryServiceImpl) Query(ctx context.Context, q *vo.FileQuery, ops *service.FileQueryOptions) (*vo.FileQuery, error) {

	if ops == nil {
		ops = &service.FileQueryOptions{}
	}

	err := inst.preprocessQuery(ctx, q)
	if err != nil {
		return nil, err
	}

	h, err := inst.findHandler(q)
	if err != nil {
		return nil, err
	}

	res, err := h.Query(q)
	if err != nil {
		return nil, err
	}

	if ops.WithContentType {
		err = inst.loadContentTypeInfo(ctx, res)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
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

func (inst *FileQueryServiceImpl) loadContentTypeInfo(ctx context.Context, q *vo.FileQuery) error {
	task := &myContentTypeLoadingTask{
		context:            ctx,
		ContentTypeService: inst.ContentTypeService,
	}
	task.init(q.Items)
	err := task.Load()
	if err != nil {
		vlog.Warn(err)
	}
	if task.types == nil {
		task.types = make([]*dto.ContentType, 0)
	}
	q.Types = task.types
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type myContentTypeInfo struct {
	suffix      string // like '.md'
	contenttype string // like 'text/markdown'
	urn         dxo.URN
	target      *dto.ContentType
}

type myContentTypeLoadingTask struct {
	context            context.Context
	ContentTypeService service.ContentTypeService
	table              map[dxo.URN]*myContentTypeInfo // map[urn]info
	items              []*dto.File
	types              []*dto.ContentType
}

func (inst *myContentTypeLoadingTask) init(items []*dto.File) {
	inst.items = items
	inst.table = make(map[dxo.URN]*myContentTypeInfo)
}

func (inst *myContentTypeLoadingTask) Load() error {

	for _, item := range inst.items {
		err := inst.prepareForItem(item)
		if err != nil {
			vlog.Warn(err)
		}
	}

	ctx := inst.context
	patterns := inst.listPatterns()
	typelist, err := inst.ContentTypeService.ListByPatterns(ctx, patterns)
	if err != nil {
		vlog.Warn(err)
	}

	for _, ct := range typelist {
		// vlog.Info("content.type:", x.TypeName)
		inst.applyContentType(ct)
	}

	inst.applyToItems()
	return nil
}

func (inst *myContentTypeLoadingTask) applyToItems() {
	for _, item := range inst.items {
		name := item.Name
		urn := inst.getURN("file", name)
		info := inst.table[urn]
		if info == nil {
			continue
		}
		ct := info.target
		if ct == nil {
			continue
		}
		item.Type = ct.TypeName.String()
	}
}

func (inst *myContentTypeLoadingTask) applyContentType(ct *dto.ContentType) {
	if !ct.AsFile {
		return
	}
	plist := ct.Patterns.StringArray().Normalize().Array()
	for _, p := range plist {
		suffix := inst.getSuffix(p)
		urn := inst.getURN("suffix", suffix)
		info := inst.table[urn]
		if info == nil {
			continue
		}
		if info.target == nil {
			info.contenttype = ct.TypeName.String()
			info.target = ct
			inst.types = append(inst.types, ct)
		}
	}
}

func (inst *myContentTypeLoadingTask) listPatterns() []string {
	list := make([]string, 0)
	for k, v := range inst.table {
		if k == v.urn {
			list = append(list, "*"+v.suffix)
		}
	}
	return list
}

func (inst *myContentTypeLoadingTask) prepareForItem(item *dto.File) error {

	if !item.IsFile {
		return nil
	}

	name := item.Name
	suffix := inst.getSuffix(name)
	urnF := inst.getURN("file", name)
	urnS := inst.getURN("suffix", suffix)
	table := inst.table

	info := table[urnS]
	if info == nil {
		info = &myContentTypeInfo{
			suffix: suffix,
			urn:    urnS,
		}
		table[urnS] = info
	}
	table[urnF] = info

	return nil
}

func (inst *myContentTypeLoadingTask) getSuffix(filename string) string {
	i := strings.LastIndexByte(filename, '.')
	if i < 0 {
		return "[no_suffix]"
	}
	suffix := filename[i:]
	return strings.ToLower(suffix)
}

func (inst *myContentTypeLoadingTask) getURN(kind, id string) dxo.URN {
	return dxo.NewURN(kind, id)
}
