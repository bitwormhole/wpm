package ifilequery

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/bitwormhole/wpm/server/classes/contenttypes"
	"github.com/bitwormhole/wpm/server/classes/filequery"
	"github.com/starter-go/afs"
	"github.com/starter-go/vlog"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(filequery.Service) //starter:as("#")

	FS           afs.FS               //starter:inject("#")
	CTypeService contenttypes.Service //starter:inject("#")

}

func (inst *ServiceImpl) _impl() filequery.Service {
	return inst
}

// Query ...
func (inst *ServiceImpl) Query(ctx context.Context, want *vo.FileQuery) (*vo.FileQuery, error) {
	finder := &fsDirItemsFinder{
		parent: inst,
	}
	result, err := finder.find(want)
	if err != nil {
		return nil, err
	}

	err = inst.loadContentTypeInfo(ctx, result)
	if err != nil {
		vlog.Warn(err.Error())
	}

	return result, nil
}

func (inst *ServiceImpl) loadContentTypeInfo(ctx context.Context, dst *vo.FileQuery) error {

	loader := &fileContentTypeInfoLoader{
		context:      ctx,
		ctypeService: inst.CTypeService,
	}

	items := dst.Items
	for _, item := range items {

		if !item.IsFile {
			continue
		}
		err := loader.loadForFile(item)
		if err != nil {
			vlog.Warn(err.Error())
		}
	}

	dst.Types = loader.listUsedTypes()
	return nil
}
