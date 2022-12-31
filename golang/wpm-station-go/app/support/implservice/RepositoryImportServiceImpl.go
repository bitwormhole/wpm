package implservice

import (
	"context"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/bitwormhole/wpm/app/service"
	"github.com/bitwormhole/wpm/app/web/dto"
	"github.com/bitwormhole/wpm/app/web/vo"
)

// RepositoryImportServiceImpl ...
type RepositoryImportServiceImpl struct {
	markup.Component `id:"RepositoryImportService"`

	LocateService     service.RepositoryLocateService `inject:"#RepositoryLocateService"`
	SearchService     service.RepositorySearchService `inject:"#RepositorySearchService"`
	RepositoryService service.RepositoryService       `inject:"#RepositoryService"`
}

func (inst *RepositoryImportServiceImpl) _Impl() service.RepositoryImportService {
	return inst
}

// Find ...
func (inst *RepositoryImportServiceImpl) Find(ctx context.Context, o1 *vo.RepositoryImport) (*vo.RepositoryImport, error) {
	src := o1.Items
	rb := &myRepositoryImportResultBuilder{}
	for _, item := range src {
		path := fs.Default().GetPath(item.Path)
		repolist, err := inst.SearchService.Search(ctx, path)
		if err == nil {
			rb.addItems(repolist, nil)
		} else {
			rb.add(item, err)
		}
	}
	return rb.create()
}

// Save ...
func (inst *RepositoryImportServiceImpl) Save(ctx context.Context, vo1 *vo.RepositoryImport) (*vo.RepositoryImport, error) {
	ser := inst.RepositoryService
	src := vo1.Items
	rb := &myRepositoryImportResultBuilder{}
	for _, item := range src {
		repo, err := ser.Insert(ctx, item)
		if repo == nil {
			repo = item
		}
		rb.add(repo, err)
	}
	return rb.create()
}

// Locate ...
func (inst *RepositoryImportServiceImpl) Locate(ctx context.Context, o1 *vo.RepositoryImport) (*vo.RepositoryImport, error) {
	src := o1.Items
	rb := &myRepositoryImportResultBuilder{}
	for _, item := range src {
		path := fs.Default().GetPath(item.Path)
		repo, err := inst.LocateService.Locate(ctx, path)
		if repo == nil {
			repo = item
		}
		rb.add(repo, err)
	}
	return rb.create()
}

// FindOrLocate ...
func (inst *RepositoryImportServiceImpl) FindOrLocate(ctx context.Context, o1 *vo.RepositoryImport) (*vo.RepositoryImport, error) {
	src := o1.Items
	rb := &myRepositoryImportResultBuilder{}
	for _, item := range src {
		path := fs.Default().GetPath(item.Path)
		// try locate
		repo, err := inst.LocateService.Locate(ctx, path)
		if err == nil {
			rb.add(repo, nil)
			continue
		}
		// try search
		repolist, err := inst.SearchService.Search(ctx, path)
		if err == nil {
			rb.addItems(repolist, err)
			continue
		}
		// deep error
		rb.add(item, err)
	}
	return rb.create()
}

////////////////////////////////////////////////////////////////////////////////

type myRepositoryImportResultBuilder struct {
	items []*dto.Repository
}

func (inst *myRepositoryImportResultBuilder) add(item *dto.Repository, err error) {
	if item == nil {
		return
	}
	if err != nil {
		item.Ready = false
		item.State = dxo.FileStateOffline
	} else {
		item.Ready = true
	}
	inst.items = append(inst.items, item)
}

func (inst *myRepositoryImportResultBuilder) addItems(items []*dto.Repository, err error) {
	for _, item := range items {
		inst.add(item, err)
	}
}

func (inst *myRepositoryImportResultBuilder) create() (*vo.RepositoryImport, error) {
	items := inst.items
	if items == nil {
		items = make([]*dto.Repository, 0)
	}
	o := &vo.RepositoryImport{}
	o.Items = items
	return o, nil
}
